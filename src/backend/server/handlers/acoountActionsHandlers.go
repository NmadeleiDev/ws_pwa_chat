package handlers

import (
	"chat_backend/db/mongodb"
	"chat_backend/db/postgres"
	"chat_backend/server/utils"
	"chat_backend/structs"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func	SignUpHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for signup: ", err)
			return
		}

		userData := &structs.User{}
		err = json.Unmarshal(requestData, userData)
		if err != nil {
			log.Error("Can't parse request body for signup: ", err)
			return
		}

		if !postgres.CreateUser(*userData) {
			utils.SendFailResponse(w)
			return
		}

		if !mongodb.CreateUser(*userData) {
			utils.SendFailResponse(w)
			return
		}

		utils.RefreshRequestSessionKeyCookie(w, *userData)
	}
}

func	SignInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}

		userData := &structs.User{}
		err = json.Unmarshal(requestData, userData)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}

		utils.RefreshRequestSessionKeyCookie(w, *userData)
	}
}

func	SignOutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		sessionKey := utils.GetCookieValue(r, "session_id")

		postgres.DeleteSessionKey(sessionKey)

		utils.SendSuccessResponse(w)
	}
}

func	UnregisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}

		userData := &structs.User{}
		err = json.Unmarshal(requestData, userData)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}

		if utils.RefreshRequestSessionKeyCookie(w, *userData) {
			postgres.DeleteUser(*userData)
			// TODO mongo.DeleteUser
		} else {
			log.Error("Looks like unverified attempt to delete account... User: ", userData)
		}
	}
}