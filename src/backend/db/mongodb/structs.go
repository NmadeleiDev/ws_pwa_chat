package mongodb

import (
	"chat_backend/structs"
)

type UpdatedMessageData struct {
	OperationType			string		`bson:"operationType"`
	Message					structs.Message	`bson:"fullDocument"`
}

type UpdatedChatsData struct {
	OperationType			string		`bson:"-"`
	Chats					[]structs.Chat	`bson:"updateDescription.updatedFields.chats"`
}
