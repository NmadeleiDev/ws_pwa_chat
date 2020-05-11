package handlers

import (
	"chat_backend/db/mongodb"
	"chat_backend/db/postgres"
	"chat_backend/server/utils"
	"chat_backend/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func	GetUserDataHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sessionKey := utils.GetCookieValue(r, "session_id")

		userData, err := postgres.GetUserNameAndId(sessionKey)
		if err != nil {
			log.Error("Error getting user data from postgres: ", err)
			utils.SendFailResponse(w)
			return
		}
		err = mongodb.FillUserData(&userData)
		if err != nil {
			log.Error("Error getting user data from mongo: ", err)
			utils.SendFailResponse(w)
			return
		}
		utils.SendDataResponse(w, userData)
	}
}

func	CreateChatHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var chatData structs.Chat

		if !utils.ValidateRequest(w, r) {
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}

		err = json.Unmarshal(requestData, &chatData)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		chat, err := mongodb.CreateChat(chatData)
		if err != nil {
			log.Error("Error creating chat: ", err)
			return
		}
		utils.SendDataResponse(w, chat)
	}
}

func	GetChatMessagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		if !utils.ValidateRequest(w, r) {
			return
		}
		chatId := mux.Vars(r)["chatId"]

		messages, err := mongodb.GetMessagesFromPool(chatId)
		if err != nil {
			log.Error("Error getting messages: ", err)
			utils.SendFailResponse(w)
		}
		utils.SendDataResponse(w, messages)
	}
}

func	GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		if !utils.ValidateRequest(w, r) {
			return
		}
		users, err := postgres.GetAllUsers()
		if err != nil {
			log.Error("Error getting messages: ", err)
			utils.SendFailResponse(w)
		}
		utils.SendDataResponse(w, users)
	}
}

func AddUserToChatHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		var userToAddData struct {
			User	structs.User	`json:"user"`
			Chat	structs.Chat	`json:"chat"`
		}
		if !utils.ValidateRequest(w, r) {
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		err = json.Unmarshal(requestData, &userToAddData)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if !mongodb.AddUserToChatMembers(userToAddData.Chat.ChatId, userToAddData.User) {
			utils.SendFailResponse(w)
			return
		}
		if !mongodb.AddChatToUserChats(userToAddData.Chat, userToAddData.User.Username) {
			utils.SendFailResponse(w)
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func SaveChatNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newChatData structs.Chat
		if !utils.ValidateRequest(w, r) {
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for chat name edit: ", err)
			return
		}
		err = json.Unmarshal(requestData, &newChatData)
		if err != nil {
			log.Error("Can't parse request body for chat name edit: ", err)
			return
		}
		if !mongodb.EditChatName(newChatData) {
			utils.SendFailResponse(w)
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func UpdateLastReadMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message structs.Message
		sessionKey := utils.GetCookieValue(r, "session_id")
		userData, err := postgres.GetUserNameAndId(sessionKey)
		if err != nil {
			log.Error("Error getting user data from postgres: ", err)
			utils.SendFailResponse(w)
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		err = json.Unmarshal(requestData, &message)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if !mongodb.UpdateLastReadMessageId(message, userData.Username) {
			utils.SendFailResponse(w)
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func LeaveChatHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		var chat structs.Chat
		sessionKey := utils.GetCookieValue(r, "session_id")
		userData, err := postgres.GetUserNameAndId(sessionKey)
		if err != nil {
			log.Error("Error getting user data from postgres: ", err)
			utils.SendFailResponse(w)
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		err = json.Unmarshal(requestData, &chat)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if !mongodb.DeleteChatFromUserChats(chat, userData.Username) {
			utils.SendFailResponse(w)
			return
		}
		if !mongodb.DeleteUserFromChatMembers(chat.ChatId, userData.Username) {
			utils.SendFailResponse(w)
			return
		}
		utils.SendSuccessResponse(w)
	}
}