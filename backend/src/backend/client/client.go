package client

import (
	"chat_backend/db/mongodb"
	"chat_backend/structs"
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
	ClientExitChan		chan byte
}

func	CreateNewClient(connection *websocket.Conn, user structs.User) (client Client) {

	client = Client{User:user, Connection:connection, ReadMessageChan:make(chan structs.Message), ClientExitChan: make(chan byte)}
	return client
}

func	(client *Client) SubscribeToChatEvents() {

	for _, chat := range client.User.Chats {
		go mongodb.ListenChangeStream(chat.MessagePoolId, client.ClientExitChan, client.ReadMessageChan)
		log.Printf("Subscribed %v to %v chat", client.User.Username, chat.Name)
	}
}
