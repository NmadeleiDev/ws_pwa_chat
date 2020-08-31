package main

import (
	"fileServer/db/fileInfoManager"
	"fileServer/envParser"
	"fileServer/s3Manager"
	"fileServer/server"
)

func main() {
	fileInfoManager.Init()
	s3Manager.Init()

	defer func() {
		fileInfoManager.Manager.CloseConnection()
	}()

	server.Start(envParser.GetParser().GetServerPort())
}
