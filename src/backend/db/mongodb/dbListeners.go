package mongodb

import (
	"chat_backend/constants"
	"chat_backend/structs"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListenChatChangeStream(messagePoolId string, chatId string, clientExitChan chan byte, writeUpdatesChan chan structs.SocketMessage) {

	database := client.Database("chat")
	collection := database.Collection(messagePoolId)
	collectionUpdateChan := make(chan structs.Message)

	matchStage := bson.D{{"$match", bson.D{{"operationType", "insert"}}}}
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), mongo.Pipeline{matchStage}, opts)
	if err != nil {
		log.Error("Error watching message pool: ", err,  " poolID: ", messagePoolId)
		return
	}

	go func(stream mongo.ChangeStream, updateChan chan structs.Message) {
		for {
			if stream.TryNext(context.TODO()) {
				container := UpdatedMessageData{}
				if err := stream.Decode(&container); err != nil {
					log.Error("Error decoding message: ", err)
				}
				updateChan <- container.Message
				continue
			}
			// If TryNext returns false, the next change is not yet available, the change stream was closed by the server,
			// or an error occurred. TryNext should only be called again for the empty batch case.
			if err := stream.Err(); err != nil {
				log.Errorf("Error reading messages stream: %v;", err)
				break
			}
			if stream.ID() == 0 {
				break
			}
		}
	}(*changeStream, collectionUpdateChan)

	log.Infof("Subscribed someone to his chats with change stream %v", changeStream.ID())

	defer func() {
		if err := changeStream.Close(context.TODO()); err != nil {
			log.Error("error closing change stream: ", err)
		}
		close(collectionUpdateChan)
	}()

	for {
		select {
		case message := <- collectionUpdateChan:
			message.ChatId = chatId
			update := structs.SocketMessage{Type: constants.MessageType, Data: message, Error: nil}
			writeUpdatesChan <- update
		case <- clientExitChan:
			return
		}
	}
}

func ListenUserChatsStream(user *structs.User, clientExitChan chan byte, writeUpdatesChan chan structs.SocketMessage) {
	database := client.Database("user")
	collection := database.Collection("users")
	collectionUpdateChan := make(chan structs.Chat)

	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{{"$and",
		bson.A{
			bson.D{{"fullDocument.username", user.Username}},
			bson.D{{"operationType", "update"}}}}}}},
		bson.D{{"$project", bson.D{{"updateDescription.updatedFields", 1}}}}}
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), pipeline, opts)
	if err != nil {
		log.Errorf("Error creating change stream %v for user %v", err, user.Username)
		return
	}

	go func(stream mongo.ChangeStream, updateChan chan structs.Chat) {
		for {
			if stream.TryNext(context.TODO()) {
				chat := structs.Chat{}
				current := stream.Current
				chatsObjName := getChatsObjName(len(user.Chats)) // Объясняю: монго возвращает вставенный чат не просто как объект, а как массив с одмин чатом, с названием по форме "chats.{номер вставленного чата}", поэтому, чтобы его найти, приходится создавать именно такое название
				log.Info("Chat obj: ", chatsObjName)
				err = current.Lookup("updateDescription", "updatedFields", chatsObjName).Unmarshal(&chat)
				if err != nil {
					log.Error("error getting updated chat: ", err)
				}
				log.Infof("Got new update chat: %v", chat)
				if stream.ID() != 0 { // проверяем, не закрыт ли stream, так как если он закрыт, то закрыт и updateChan, и знать писать не него не надо
					updateChan <- chat
				}
				continue
			}
			if err := stream.Err(); err != nil {
				log.Errorf("Error reading chats stream: %v; for user %v", err, user.Username)
				log.Info("Leaving stream goroutine")
				break
			}
			if stream.ID() == 0 {
				break
			}
		}
	}(*changeStream, collectionUpdateChan)

	log.Infof("Subscribed %v to his chats with change stream %v", user.Username, changeStream.ID())

	defer func() {
		if err := changeStream.Close(context.TODO()); err != nil {
			log.Error("error closing %v change stream: ", changeStream.ID(), err)
		}
		close(collectionUpdateChan)
		log.Infof("Closed %v chats stream for %v", changeStream.ID(), user.Username)
	}()

	for {
		select {
		case chat := <- collectionUpdateChan:
			chat, err = GetChatDataById(chat.ChatId)
			if err != nil {
				log.Error("Error getting chat data by id: ", err)
			}
			update := structs.SocketMessage{Type: constants.ChatType, Data: chat, Error: err}
			writeUpdatesChan <- update
		case <- clientExitChan:
			log.Infof("Exiting from %v chats stream", user.Username)
			return
		}
	}
}