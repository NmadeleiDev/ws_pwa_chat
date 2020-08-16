package mainDataStorage

import (
	"chat_backend/db/mainDataStorage/mongodb"
	"chat_backend/structs"
)

var Manager MainDataManager

type MainDataManager interface {
	MakeConnection()
	CloseConnection()
	CreateUser(user structs.User) (id string, ok bool)
	FillUserData(user *structs.User) bool
	GetChatDataById(chatId string) (structs.Chat, error)
	GetMessagesFromPool(chatId string) ([]structs.Message, error)
	WriteNewMessage(messagePoolId string, message structs.Message)
	UpdateMessage(messagePoolId string, message structs.Message)
	DeleteMessage(messagePoolId string, message structs.Message)
	CreateChat(newChat structs.ChatWithMessages) (structs.ChatWithMessages, error)
	AddUserToChatMembers(chatId string, user structs.User) bool
	DeleteUserFromChatMembers(chatId string, username string) bool
	AddChatToUserChats(chat structs.ChatInfo, usernames []string) bool
	DeleteChatFromUserChats(chat structs.Chat, username string) bool
	EditChatName(chat structs.Chat) bool
	EditChatStorePeriod(chat structs.Chat) bool
	UpdateLastReadMessageId(message structs.Message, username string) bool

	ListenChatMessagesStream(messagePoolId string, chatId string, clientExitChan chan byte, writeUpdatesChan chan structs.SocketMessage)
	ListenUserChatsStream(user *structs.User, clientExitChan chan byte, writeUpdatesChan chan structs.SocketMessage)
}

func Init() {
	Manager = &mongodb.MongoMainDataStorage{}
	Manager.MakeConnection()
}
