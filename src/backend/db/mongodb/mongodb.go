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

var client *mongo.Client

func MakeConnection() {
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
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error getting client mongo: ", err)
	}
	if err != nil {
		log.Fatal("Error connecting to mongo: ", err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Error pinging: ", err)
	}
	log.Info("Connected.")
}

func CreateUser(user structs.User) bool {

	database := client.Database("user")
	userCollection := database.Collection("users")

	user.Chats = make([]structs.Chat, 0)

	// I don't need to store password in mongo
	user.Password = ""

	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Error("Error creating user in mongo: ", err)
		return false
	}
	return true
}

func FillUserData(user *structs.User) error {

	database := client.Database("user")
	userCollection := database.Collection("users")

	filter := bson.M{"username": user.Username}
	container := structs.User{}
	err := userCollection.FindOne(context.Background(),filter).Decode(&container)
	if  err != nil {
		log.Error("Error finding user document: ", err)
		return err
	} else {
		log.Info("Got user document: ", container)
	}
	user.Chats = container.Chats

	for i, chat := range user.Chats {
		container, err := GetChatDataById(chat.ChatId)
		if err != nil {
			continue
		}
		container.LastReadMessageId = chat.LastReadMessageId
		user.Chats[i] = container
	}

	return err
}

func GetChatDataById(chatId string) (structs.Chat, error) {
	database := client.Database("user")
	chatCollection := database.Collection("chats")
	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Errorf("Error creating object id: %v; chat id provided: %v", err, chatId)
		return structs.Chat{}, err
	}
	filter := bson.D{{"_id",  objectId}}
	container := structs.Chat{}
	err = chatCollection.FindOne(context.Background(),filter).Decode(&container)
	if  err != nil {
		log.Error("Error finding chat document: ", err)
		return structs.Chat{}, err
	}
	container.ChatId = chatId
	return container, nil
}

func GetMessagesFromPool(chatId string) ([]structs.Message, error) {

	var chat structs.Chat
	var messages = make([]structs.Message, 0, 20)

	database := client.Database("user")
	coll := database.Collection("chats")

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

	database = client.Database("chat")
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

func WriteNewMessage(messagePoolId string, message structs.Message) {

	database := client.Database("chat")
	coll := database.Collection(messagePoolId)

	message.ChatId = ""

	_, err := coll.InsertOne(context.TODO(), message)
	if err != nil {
		log.Error("Error saving message to mongodb: ", err)
	}
}

func UpdateMessage(messagePoolId string, message structs.Message) {

	database := client.Database("chat")
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

func DeleteMessage(messagePoolId string, message structs.Message) {

	database := client.Database("chat")
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

func CreateChat(newChat structs.Chat) (structs.Chat, error) {

	database := client.Database("user")
	chatsCollection := database.Collection("chats")

	newMessagePoolId := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String() + strconv.Itoa(rand.Int()) + strings.Join(newChat.Usernames, "hi!"))))
	newChat.MessagePoolId = newMessagePoolId

	res, err := chatsCollection.InsertOne(context.TODO(), newChat)
	if err != nil {
		log.Error("Error inserting chat to mongo: ", err)
		return structs.Chat{}, err
	}
	newChat.ChatId = strings.Split(res.InsertedID.(primitive.ObjectID).String(), "\"")[1]

	for _, user := range newChat.Usernames {
		AddChatToUserChats(newChat, user)
	}

	return newChat, nil
}

func AddUserToChatMembers(chatId string, user structs.User) bool {
	database := client.Database("user")
	chatsCollection := database.Collection("chats")

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

func DeleteUserFromChatMembers(chatId string, username string) bool {
	database := client.Database("user")
	chatsCollection := database.Collection("chats")

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

func AddChatToUserChats(chat structs.Chat, username string) bool {
	database := client.Database("user")
	userCollection := database.Collection("users")

	filter := bson.D{{"username", username}}
	update := bson.D{{"$addToSet", bson.D{{"chats", bson.D{{"chatid", chat.ChatId}, {"messagepoolid", chat.MessagePoolId}, {"last_read_message_id", ""}}}}}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating user chats: ", err)
		return false
	}
	log.Infof("Pushed new chat for %v user: %v", result.ModifiedCount, username)
	return true
}

func DeleteChatFromUserChats(chat structs.Chat, username string) bool {
	database := client.Database("user")
	userCollection := database.Collection("users")

	filter := bson.D{{"username", username}}
	update := bson.D{{"$pull", bson.D{{"chats", bson.D{{"chatid", chat.ChatId}}}}}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error deleting user chat: ", err)
		return false
	}
	log.Infof("deleted chat for %v user: %v", result.ModifiedCount, username)
	return true
}

func EditChatName(chat structs.Chat) bool {
	database := client.Database("user")
	chatsCollection := database.Collection("chats")

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

func UpdateLastReadMessageId(message structs.Message, username string) bool {
	database := client.Database("user")
	userCollection := database.Collection("users")

	filter := bson.D{{"username", username}}
	update := bson.D{{"$set", bson.D{{"chats.$[chat].last_read_message_id",  message.Id}}}}
	arrayFilter := options.ArrayFilters{Filters: []interface{}{bson.M{"chat.chatid": message.ChatId}}}
	opts := options.UpdateOptions{ArrayFilters: &arrayFilter}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update, &opts)
	if err != nil {
		log.Error("Error updating user chats: ", err)
		return false
	}
	log.Infof("Pushed new chat for %v user: %v", result.ModifiedCount, username)
	return true
}

func CloseConnection() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Error("Error disconnect mongo: ", err)
	}
}