package handlers

import (
	"chat_backend/client"
	"chat_backend/db/mainDataStorage"
	"chat_backend/db/userKeysData"
	"chat_backend/server/utils"
	"chat_backend/structs"
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
	id, ok := utils.IdentifyWebOrMobileRequest(r)
	if !ok {
		utils.SendFailResponse(w, "Unauthorized request")
		return
	}
	user := &structs.User{Id: id}
	if !mainDataStorage.Manager.FillUserData(user) {
		log.Error("Failed to get user data from mongo.")
		return
	}
	go userKeysData.Manager.ToggleUserOnlineState(user.Id, true)

	clientStruct := client.CreateNewClient(connection, user)
	clientStruct.SubscribeToDBEvents()
	go clientStruct.ReadHub()
	go clientStruct.WriteHub()
}



