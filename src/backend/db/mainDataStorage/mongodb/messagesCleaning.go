package mongodb

import (
	"chat_backend/structs"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (db *MongoMainDataStorage) StartCleaningMessages() {

	database := db.client.Database(usersDb)
	collection := database.Collection(chatsDataCollection)

	for {
		time.Sleep(time.Minute * 3)
		var container structs.Chat
		optsFind := options.Find()
		cursor, err := collection.Find(context.TODO(), bson.D{}, optsFind)
		if err != nil {
			logrus.Error("error finding messages: ", err)
			continue
		}
		for cursor.Next(context.TODO()) {
			if err := cursor.Decode(&container); err != nil {
				logrus.Errorf("Error decoding chat: %v", err)
				continue
			}
			dbase := db.client.Database("chat")
			coll := dbase.Collection(container.MessagePoolId)

			timeLimit := time.Now().Add(-(time.Hour * time.Duration(container.StorePeriod))).Unix()

			deleteFilter := bson.D{{"date", bson.D{{"$lt", timeLimit}}}}
			res, err := coll.DeleteMany(context.Background(), deleteFilter, options.Delete())
			if err != nil {
				logrus.Errorf("Error deleting old messages: %v", err)
			}
			logrus.Infof("Deleted %v messages", res.DeletedCount)
		}
	}
}
