package handlers

import (
	"chat_backend/db/mongodb"
	"chat_backend/db/postgres"
	"chat_backend/server/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func	GetUserDataHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		sessionKey := utils.GetCookieValue(r, "session_id")

		userData, err := postgres.GetUserNameAndEmail(sessionKey)
		if err != nil {
			log.Error("Error getting user data from postgres: ", err)
			utils.SendFailResponse(w)
			return
		}
		userData, err = mongodb.FillUserData(userData)
		if err != nil {
			log.Error("Error getting user data from mongo: ", err)
			utils.SendFailResponse(w)
			return
		}
		utils.SendDataResponse(w, userData)
	}
}

func	GetChatMessagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

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

		users, err := postgres.GetAllUsers()
		if err != nil {
			log.Error("Error getting messages: ", err)
			utils.SendFailResponse(w)
		}
		utils.SendDataResponse(w, users)
	}
}