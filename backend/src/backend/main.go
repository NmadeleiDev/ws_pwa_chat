package main

import (
	"chat_backend/db/mongodb"
	"chat_backend/db/postgres"
	"chat_backend/server"
	"os"
)

func main() {

	defer func() {
		postgres.CloseConnection()
		mongodb.CloseConnection()
	}()

	port := os.Getenv("BACKEND_PORT")

	postgres.MakeConnection()
	mongodb.MakeConnection()

	postgres.InitTables()

	server.StartServer(port)
}
