package client

import (
	"chat_backend/db/mongodb"
	"chat_backend/structs"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

func	(client *Client) ReadHub() {
	var messageStruct structs.Message

	defer func() {
		if err := client.Connection.Close(); err != nil {
			log.Error("Error closing connection in read: ", err)
		}
	}()

	client.Connection.SetReadLimit(maxMessageSize)
	if err := client.Connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Error("Error setting write deadline: ", err)
		return
	}
	client.Connection.SetPongHandler(func(string) error { client.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		if err := json.Unmarshal(message, &messageStruct); err != nil {
			log.Error("Error unmarshal message: ", err, " Message: ", message)
		} else {
			if messageStruct.Meta == "pong" {
				log.Info("Got pong message.")
				continue
			}
			log.Info("Got message: ", messageStruct)
			if messageStruct.ChatId == "" && len(messageStruct.Meta) > 0 {
				newChat, err := mongodb.CreateChatFromMessage(messageStruct)
				if err != nil {
					log.Error("Error creating chat from message: ", err)
					break
				}
				client.User.Chats = append(client.User.Chats, newChat)
				messageStruct.ChatId = newChat.ChatId
			}
			for _, chat := range client.User.Chats {
				if chat.ChatId == messageStruct.ChatId {
					go mongodb.WriteNewMessage(chat.MessagePoolId, messageStruct)
				}
			}
		}
	}
}

