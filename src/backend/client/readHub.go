package client

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/structs"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	NewChatMeta = 1
	InsertMessage = 1
	UpdateMessage = 2
	DeleteMessage = 3

	MessageSent = 1
	MessageDelivered = 2
	MessageRead = 3
)

func	(client *Client) ReadHub() {
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
		var messageStruct structs.ClientMessage
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
			log.Info("Got message in read hub: ", messageStruct)
			for _, chat := range client.User.Chats {
				if chat.ChatId == messageStruct.Message.ChatId {
					switch messageStruct.Type {
					case InsertMessage:
						messageStruct.Message.State = MessageDelivered
						go mainDataStorage.Manager.WriteNewMessage(chat.MessagePoolId, messageStruct.Message)
					case UpdateMessage:
						go mainDataStorage.Manager.UpdateMessage(chat.MessagePoolId, messageStruct.Message)
					case DeleteMessage:
						go mainDataStorage.Manager.DeleteMessage(chat.MessagePoolId, messageStruct.Message)
					default:
						log.Warn("Unknown message type: ", messageStruct)
					}
					log.Info("Found matching chat for message " + messageStruct.Message.Text)
				}
			}
		}
	}
}