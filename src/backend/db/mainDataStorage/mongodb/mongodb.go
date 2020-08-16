package mongodb

import (
	"chat_backend/structs"
	"context"
	"crypto/md5"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type MongoMainDataStorage struct {
	client *mongo.Client
}

const (
	usersDb             = "user"
	usersDataCollection = "users"
	chatsDataCollection = "chats"
)

func (db *MongoMainDataStorage) MakeConnection() {
	var err error
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	addr := os.Getenv("MONGO_ADDRESS")

	if user == "" || password == "" || addr == "" {
		log.Error("Env is empty", user, password, addr)
	}

	connStr := fmt.Sprintf("mongodb://%v:%v@%v", user, password, addr)
	log.Info("Connecting to mongo: ", connStr)
	opts := options.Client().ApplyURI(connStr).SetReplicaSet("rs0")
	db.client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error getting client mongo: ", err)
	}
	if err != nil {
		log.Fatal("Error connecting to mongo: ", err)
	}
	if err = db.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Error pinging: ", err)
	}
	log.Info("Connected.")
}

func (db *MongoMainDataStorage) CreateUser(user structs.User) (id string, ok bool) {

	database := db.client.Database(usersDb)
	userCollection := database.Collection(usersDataCollection)

	user.Chats = make([]structs.Chat, 0)

	// I don't need to store password in mongo
	user.Password = ""

	res, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Error("Error creating user in mongo: ", err)
		return "", false
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), true
}

func (db *MongoMainDataStorage) FillUserData(user *structs.User) bool {

	database := db.client.Database(usersDb)
	userCollection := database.Collection(usersDataCollection)

	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		log.Errorf("error creating objectId from user id: %v", err)
		return false
	}

	filter := bson.M{"_id": objectId}
	container := structs.User{}
	err = userCollection.FindOne(context.Background(), filter).Decode(&container)
	if err != nil {
		log.Error("Error finding user document: ", err)
		return false
	} else {
		log.Info("Got user document: ", container)
	}
	user.Username = container.Username
	user.Chats = container.Chats

	for i, chat := range user.Chats {
		container, err := db.GetChatDataById(chat.ChatId)
		if err != nil {
			continue
		}
		container.LastReadMessageId = chat.LastReadMessageId
		user.Chats[i] = container
	}

	return true
}

func (db *MongoMainDataStorage) GetChatDataById(chatId string) (structs.Chat, error) {
	database := db.client.Database(usersDb)
	chatCollection := database.Collection(chatsDataCollection)
	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Errorf("Error creating object id: %v; chat id provided: %v", err, chatId)
		return structs.Chat{}, err
	}
	filter := bson.D{{"_id", objectId}}
	container := structs.Chat{}
	err = chatCollection.FindOne(context.Background(), filter).Decode(&container)
	if err != nil {
		log.Error("Error finding chat document: ", err)
		return structs.Chat{}, err
	}
	container.ChatId = chatId
	return container, nil
}

func (db *MongoMainDataStorage) GetMessagesFromPool(chatId string) ([]structs.Message, error) {

	var chat structs.Chat
	var messages = make([]structs.Message, 0, 20)

	database := db.client.Database(usersDb)
	coll := database.Collection(chatsDataCollection)

	opts := options.FindOne()
	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Error("Error creating object id: %v; provided chat id: %v", err, chatId)
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	err = coll.FindOne(context.TODO(), filter, opts).Decode(&chat)
	if err != nil {
		log.Error("Error getting chat: ", err)
		return nil, err
	}

	database = db.client.Database("chat")
	coll = database.Collection(chat.MessagePoolId)

	optsFind := options.Find().SetSort(bson.D{{"date", 1}})
	cursor, err := coll.Find(context.TODO(), bson.D{}, optsFind)
	if err != nil {
		log.Error("error finding messages: ", err)
		return nil, err
	}
	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	for cursor.Next(context.TODO()) {
		// A new result variable should be declared for each document.
		var result structs.Message
		if err := cursor.Decode(&result); err != nil {
			log.Error("Error unmarshal message: ", err)
		} else {
			messages = append(messages, result)
		}
	}
	return messages, nil
}

func (db *MongoMainDataStorage) WriteNewMessage(messagePoolId string, message structs.Message) {

	database := db.client.Database("chat")
	coll := database.Collection(messagePoolId)

	message.ChatId = ""

	_, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		log.Error("Error saving message to mongodb: ", err)
	}
}

func (db *MongoMainDataStorage) UpdateMessage(messagePoolId string, message structs.Message) {

	database := db.client.Database("chat")
	coll := database.Collection(messagePoolId)

	message.ChatId = ""

	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"id", message.Id}}
	update := bson.D{{"$set", bson.D{{"state", message.State}, {"text", message.Text}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Error("Error updating message: ", err)
	}
	log.Infof("Updated %v document message", result.ModifiedCount)
}

func (db *MongoMainDataStorage) DeleteMessage(messagePoolId string, message structs.Message) {

	database := db.client.Database("chat")
	coll := database.Collection(messagePoolId)

	message.ChatId = ""

	opts := options.Delete()
	filter := bson.D{{"id", message.Id}}

	result, err := coll.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		log.Error("Error updating message: ", err)
	}
	log.Infof("Deleted %v document message", result.DeletedCount)
}

func (db *MongoMainDataStorage) CreateChat(newChat structs.ChatWithMessages) (structs.ChatWithMessages, error) {

	database := db.client.Database(usersDb)
	chatsCollection := database.Collection(chatsDataCollection)

	newMessagePoolId := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+strconv.Itoa(rand.Int())+strings.Join(newChat.Usernames, "hi!"))))
	newChat.MessagePoolId = newMessagePoolId

	chatData := structs.Chat{MessagePoolId: newMessagePoolId,
		ChatId:            newChat.ChatId,
		Admin:             newChat.Admin,
		Usernames:         newChat.Usernames,
		Name:              newChat.Name,
		LastReadMessageId: newChat.LastReadMessageId,
		StorePeriod:       30}
	res, err := chatsCollection.InsertOne(context.TODO(), chatData)
	if err != nil {
		log.Error("Error inserting chat to mongo: ", err)
		return structs.ChatWithMessages{}, err
	}
	newChat.ChatId = strings.Split(res.InsertedID.(primitive.ObjectID).String(), "\"")[1]

	db.AddChatToUserChats(structs.ChatInfo{ChatId: newChat.ChatId, MessagePoolId: newChat.MessagePoolId}, newChat.Usernames)

	for _, message := range newChat.Messages {
		db.WriteNewMessage(newMessagePoolId, message)
	}

	return newChat, nil
}

func (db *MongoMainDataStorage) AddUserToChatMembers(chatId string, user structs.User) bool {
	database := db.client.Database(usersDb)
	chatsCollection := database.Collection(chatsDataCollection)

	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Error("Error creating object id: %v; provided chat id: %v", err, chatId)
		return false
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$addToSet", bson.D{{"usernames", user.Username}}}}

	result, err := chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating chats collection: ", err)
		return false
	}
	log.Infof("Pushed %v new user %v for chat: %v", result.ModifiedCount, user, chatId)
	return true
}

func (db *MongoMainDataStorage) DeleteUserFromChatMembers(chatId string, username string) bool {
	database := db.client.Database(usersDb)
	chatsCollection := database.Collection(chatsDataCollection)

	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Error("Error creating object id: %v; provided chat id: %v", err, chatId)
		return false
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$pull", bson.D{{"usernames", username}}}}

	result, err := chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error deleting user from chat: ", err)
		return false
	}
	log.Infof("deleted %v user %v from chat: %v", result.ModifiedCount, username, chatId)
	return true
}

func (db *MongoMainDataStorage) AddChatToUserChats(chat structs.ChatInfo, usernames []string) bool {
	database := db.client.Database(usersDb)
	userCollection := database.Collection(usersDataCollection)

	bsonUsernames := bson.A{}
	for _, name := range usernames {
		bsonUsernames = append(bsonUsernames, bson.D{{"username", name}})
	}

	filter := bson.D{{"$or", bsonUsernames}}
	update := bson.D{{"$addToSet", bson.D{{chatsDataCollection, bson.D{{"chatid", chat.ChatId}, {"messagepoolid", chat.MessagePoolId}, {"last_read_message_id", ""}}}}}}

	result, err := userCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating user chats: ", err)
		return false
	}
	log.Infof("Pushed new chat for %v user: %v", result.ModifiedCount, usernames)
	return true
}

func (db *MongoMainDataStorage) DeleteChatFromUserChats(chat structs.Chat, userId string) bool {
	database := db.client.Database(usersDb)
	userCollection := database.Collection(usersDataCollection)
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Error("Error creating object id: %v; provided user id: %v", err, chat)
		return false
	}

	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$pull", bson.D{{chatsDataCollection, bson.D{{"chatid", chat.ChatId}}}}}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error deleting user chat: ", err)
		return false
	}
	log.Infof("deleted chat for %v user: %v", result.ModifiedCount, userId)
	return true
}

func (db *MongoMainDataStorage) EditChatName(chat structs.Chat) bool {
	database := db.client.Database(usersDb)
	chatsCollection := database.Collection(chatsDataCollection)

	objectId, err := primitive.ObjectIDFromHex(chat.ChatId)
	if err != nil {
		log.Error("Error creating object id: %v; provided chat id: %v", err, chat)
		return false
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"name", chat.Name}}}}

	_, err = chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating chats collection: ", err)
		return false
	}
	return true
}

func (db *MongoMainDataStorage) EditChatStorePeriod(chat structs.Chat) bool {
	database := db.client.Database(usersDb)
	chatsCollection := database.Collection(chatsDataCollection)

	objectId, err := primitive.ObjectIDFromHex(chat.ChatId)
	if err != nil {
		log.Error("Error creating object id: %v; provided chat id: %v", err, chat)
		return false
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"store_period", chat.StorePeriod}}}}

	_, err = chatsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating chats collection (store period): ", err)
		return false
	}
	return true
}

func (db *MongoMainDataStorage) UpdateLastReadMessageId(message structs.Message, id string) bool {
	database := db.client.Database(usersDb)
	userCollection := database.Collection(usersDataCollection)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error("Error creating object id: %v; provided user id: %v", err, id)
		return false
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"chats.$[chat].last_read_message_id", message.Id}}}}
	arrayFilter := options.ArrayFilters{Filters: []interface{}{bson.M{"chat.chatid": message.ChatId}}}
	opts := options.UpdateOptions{ArrayFilters: &arrayFilter}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update, &opts)
	if err != nil {
		log.Error("Error updating user chats: ", err)
		return false
	}
	log.Infof("Pushed new chat for %v user: %v", result.ModifiedCount, id)
	return true
}

func (db *MongoMainDataStorage) CloseConnection() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Error("Error disconnect mongo: ", err)
	}
}
