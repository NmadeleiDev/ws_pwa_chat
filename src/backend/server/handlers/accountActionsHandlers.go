package handlers

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/db/userKeysData"
	"chat_backend/server/utils"
	"chat_backend/structs"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for signup: ", err)
			return
		}
		defer r.Body.Close()

		dataCont := struct {
			Data structs.User `json:"data"`
		}{}
		err = json.Unmarshal(requestData, &dataCont)
		if err != nil {
			log.Error("Can't parse request body for signup: ", err)
			return
		}
		id, ok := mainDataStorage.Manager.CreateUser(dataCont.Data)
		if ok {
			dataCont.Data.Id = id
			cookie, token, success := userKeysData.Manager.CreateUserAndGenerateKeys(dataCont.Data)
			if !success {
				utils.SendFailResponse(w, "Unauthorized request")
				return
			}
			dataCont.Data.Token = token
			dataCont.Data.Password = ""
			utils.SetCookie(&w, "session_id", cookie)
			utils.SendDataResponse(w, dataCont.Data)
			//utils.SendSuccessResponse(w)
		} else {
			utils.SendFailResponse(w, "error")
		}
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		requestData, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			log.Error("Can't read request body for login: ", err)
			utils.SendFailResponse(w, "get body error")
			return
		}

		dataCont := struct {
			Data structs.User `json:"data"`
		}{}

		err = json.Unmarshal(requestData, &dataCont)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			utils.SendFailResponse(w, "body read error")
			return
		}
		userData := &dataCont.Data

		cookie, token, success := userKeysData.Manager.SignInAndRefreshMobileAndWebKeys(*userData)
		if !success {
			utils.SendFailResponse(w, "Wrong password")
			return
		}
		userData.Token = token
		userData.Password = ""
		utils.SetCookie(&w, "session_id", cookie)
		log.Infof("Set user cookie: %v; id: %v", cookie, userData.Id)
		utils.SendDataResponse(w, userData)
	}
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		sessionKey := utils.GetCookieValue(r, "session_id")

		userKeysData.Manager.DeleteCookieKey(sessionKey)

		utils.SendSuccessResponse(w)
	}
}

func UnregisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		requestData, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

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

		//if utils.RefreshRequestSessionKeyCookie(w, *userData) {
		//	userKeysData.Manager.DeleteUser(*userData)
		//	utils.SendSuccessResponse(w)
		//	// TODO mongo.DeleteUser
		//} else {
		//	log.Error("Looks like unverified attempt to delete account... User: ", userData)
		//}
	}
}
