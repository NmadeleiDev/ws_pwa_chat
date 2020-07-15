package client

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/structs"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = pongWait * 9 / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

type Client struct {
	User            *structs.User
	Connection      *websocket.Conn
	ReadMessageChan chan structs.SocketMessage
	ClientExitChan  chan byte
}

type ServerClientMessage struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

func CreateNewClient(connection *websocket.Conn, user *structs.User) (client *Client) {

	client = &Client{User: user, Connection: connection, ReadMessageChan: make(chan structs.SocketMessage), ClientExitChan: make(chan byte)}
	return client
}

func (client *Client) SubscribeToDBEvents() {

	// TODO зачем я подписываюсь на каждую коллекцию отдельно, занимая память горутинами, если можно подписаться на базу?
	for _, chat := range client.User.Chats {
		go mainDataStorage.Manager.ListenChatMessagesStream(chat.MessagePoolId, chat.ChatId, client.ClientExitChan, client.ReadMessageChan)
		log.Printf("Subscribed %v to %v chat", client.User.Username, chat.Name)
	}
	go mainDataStorage.Manager.ListenUserChatsStream(client.User, client.ClientExitChan, client.ReadMessageChan)
}
