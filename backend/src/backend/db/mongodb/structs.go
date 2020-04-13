package mongodb

import "chat_backend/structs"

type UpdatedMessageData struct {
	OperationType			string		`bson:"operationType"`
	Message					structs.Message	`bson:"fullDocument"`
}
