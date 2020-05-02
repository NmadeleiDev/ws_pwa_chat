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
		user.Chats[i] = container
	}

	return err
}

func GetChatDataById(chatId string) (structs.Chat, error) {
	database := client.Database("user")
	chatCollection := database.Collection("chats")
	objectId, err := primitive.ObjectIDFromHex(chatId)
	if err != nil {
		log.Error("Error creating object id: ", err)
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
		log.Error("Error creating object id: ", err)
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

func CreateChatFromMessage(message structs.Message) (structs.Chat, error) {

	database := client.Database("user")

	newChat := structs.Chat{Usernames: []string{message.Sender, message.ChatId}, // когда приходит сообщение в новый чат, имя собеседника передается в chatId
		Admin:message.Sender,
		Name:message.Sender + " + " + message.ChatId}

	newMessagePoolId := fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(message.Date) + strconv.Itoa(rand.Int()) + strings.Join(newChat.Usernames, "hi!"))))
	newChat.MessagePoolId = newMessagePoolId

	chatsCollection := database.Collection("chats")

	res, err := chatsCollection.InsertOne(context.TODO(), newChat)
	if err != nil {
		log.Error("Error inserting chat to mongo: ", err)
		return structs.Chat{}, err
	}
	newChat.ChatId = strings.Split(res.InsertedID.(primitive.ObjectID).String(), "\"")[1]

	userCollection := database.Collection("users")

	filter := bson.D{{"$or", bson.A{bson.M{"username": newChat.Usernames[0]}, bson.M{"username": newChat.Usernames[1]}}}}
	update := bson.D{{"$push", bson.D{{"chats", bson.D{{"chatid", newChat.ChatId}, {"messagepoolid", newChat.MessagePoolId}}}}}}

	result, err := userCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating user chats: ", err)
	}

	log.Infof("Pushed new chat for %v users: %v", result.ModifiedCount, newChat.Usernames)

	return newChat, nil
}

func CreateChat(newChat structs.Chat) (structs.Chat, error) {

	database := client.Database("user")

	newMessagePoolId := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String() + strconv.Itoa(rand.Int()) + strings.Join(newChat.Usernames, "hi!"))))
	newChat.MessagePoolId = newMessagePoolId

	chatsCollection := database.Collection("chats")

	res, err := chatsCollection.InsertOne(context.TODO(), newChat)
	if err != nil {
		log.Error("Error inserting chat to mongo: ", err)
		return structs.Chat{}, err
	}
	newChat.ChatId = strings.Split(res.InsertedID.(primitive.ObjectID).String(), "\"")[1]

	userCollection := database.Collection("users")

	filter := bson.D{{"$or", bson.A{bson.M{"username": newChat.Usernames[0]}, bson.M{"username": newChat.Usernames[1]}}}}
	//filter := bson.D{{"$or", newChat.Usernames}}
	update := bson.D{{"$push", bson.D{{"chats", bson.D{{"chatid", newChat.ChatId}, {"messagepoolid", newChat.MessagePoolId}}}}}}

	result, err := userCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Error("Error updating user chats: ", err)
	}

	log.Infof("Pushed new chat for %v users: %v", result.ModifiedCount, newChat.Usernames)

	return newChat, nil
}

func CloseConnection() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Error("Error disconnect mongo: ", err)
	}
}