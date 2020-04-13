package handlers

import (
	"chat_backend/client"
	"chat_backend/db/mongodb"
	"chat_backend/db/postgres"
	"chat_backend/server/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	//HandshakeTimeout:  10,
	//ReadBufferSize:    1024,
	//WriteBufferSize:   1024,
	//WriteBufferPool:   nil,
	Subprotocols:      []string{"chat"},
	//Error:             nil,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	//EnableCompression: false,
}

func ChatSocketHandler(w http.ResponseWriter, r *http.Request)  {

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error establishing ws connection: ", err)
		return
	}
	sessionKey := utils.GetCookieValue(r, "session_id")

	userData, err := postgres.GetUserNameAndEmail(sessionKey)
	if err != nil {
		log.Error("Error getting user data from postgres: ", err)
		return
	}
	userData, err = mongodb.FillUserData(userData)
	if err != nil {
		log.Error("Error getting user data from mongo: ", err)
		return
	}

	clientStruct := client.CreateNewClient(connection, userData)
	clientStruct.SubscribeToChatEvents()
	go clientStruct.ReadHub()
	go clientStruct.WriteHub()
}



