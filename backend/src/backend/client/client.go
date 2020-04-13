package client

import (
	"chat_backend/db/mongodb"
	"chat_backend/structs"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 15 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	User				structs.User
	Connection			*websocket.Conn
	ReadMessageChan		chan structs.Message
}

func	CreateNewClient(connection *websocket.Conn, user structs.User) (client Client) {

	client = Client{User:user, Connection:connection, ReadMessageChan:make(chan structs.Message)}
	return client
}

func	(client *Client) SubscribeToChatEvents() {

	for _, chat := range client.User.Chats {
		go mongodb.ListenChangeStream(chat.MessagePoolId, client.User, client.ReadMessageChan)
		log.Printf("Subscribed %v to %v chat", client.User.Username, chat.Name)
	}
}

func	(client *Client) WriteHub() {

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := client.Connection.Close(); err != nil {
			log.Error("Error closing connection in write: ", err)
		}
		close(client.ReadMessageChan)
	}()

	for {
		select {
		case message := <- client.ReadMessageChan:
			log.Info("Got message from chan: ", message.Text)
			err := client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Error("Error setting write deadline: ", err)
				return
			}
			w, err := client.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Error("Error getting next writer: ", err)
				return
			}
			messageBytes, err := json.Marshal(message)
			if err != nil {
				log.Error("Error marshalling message: ", err)
				return
			}
			_, err = w.Write(messageBytes)
			if err != nil {
				log.Error("Error writing message to ws: ", err)
				return
			}
			log.Info("Message is sent.")
		case <-ticker.C:
			err := client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Error("Error setting write deadline: ", err)
				return
			}
			if err := client.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Error("Error writing ticker message to ws: ", err)
				return
			}
		}
	}
}

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
