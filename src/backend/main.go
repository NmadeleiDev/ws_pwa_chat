package main

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/db/userKeysData"
	"chat_backend/server"
	"os"
)

func main() {

	defer func() {
		userKeysData.Manager.CloseConnection()
		mainDataStorage.Manager.CloseConnection()
	}()

	port := os.Getenv("BACKEND_PORT")

	userKeysData.Init()
	mainDataStorage.Init()

	server.StartServer(port)
}
