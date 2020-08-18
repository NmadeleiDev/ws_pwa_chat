package main

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/db/userKeysData"
	"chat_backend/server"
	"os"
	"time"
)

func main() {

	defer func() {
		userKeysData.Manager.CloseConnection()
		mainDataStorage.Manager.CloseConnection()
	}()

	port := os.Getenv("BACKEND_PORT")

	userKeysData.Init()
	mainDataStorage.Init()

	go mainDataStorage.Manager.StartCleaningMessages(time.Minute * 3)

	server.StartServer(port)
}
