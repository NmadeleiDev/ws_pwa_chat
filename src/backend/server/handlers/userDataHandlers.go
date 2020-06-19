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

func	GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	var ok bool
	// это немного необычный запрос, он гет если от веба и пост, если от мобайла.
	if r.Method == http.MethodGet {
		id, ok = utils.AuthWebRequest(r)
	} else if r.Method == http.MethodPost {
		id, ok = utils.AuthMobileToken(r)
	}
	if !ok {
		utils.SendFailResponse(w, "Unauthorized request")
		log.Infof("Unauthorized. Id: %v", id)
		return
	}
	user := &structs.User{Id: id}

	if mainDataStorage.Manager.FillUserData(user) {
		utils.SendDataResponse(w, *user)
	} else {
		utils.SendFailResponse(w, "error")
	}
}

func	CreateChatHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var chatData struct{
			Data	structs.Chat	`json:"data"`
		}

		_, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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
		chat, err := mainDataStorage.Manager.CreateChat(chatData.Data)
		if err != nil {
			log.Error("Error creating chat: ", err)
			return
		}
		utils.SendDataResponse(w, chat)
	}
}

func	GetChatMessagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var chatData struct{
			Data	structs.Chat	`json:"data"`
		}

		_, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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

		messages, err := mainDataStorage.Manager.GetMessagesFromPool(chatData.Data.ChatId)
		if err != nil {
			log.Error("Error getting messages: ", err)
			utils.SendFailResponse(w, "error")
		}
		utils.SendDataResponse(w, messages)
	}
}

func	GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var ok bool
	// это немного необычный запрос, он гет если от веба и пост, если от мобайла.
	if r.Method == http.MethodGet {
		id, ok = utils.AuthWebRequest(r)
	} else if r.Method == http.MethodPost {
		id, ok = utils.AuthMobileToken(r)
	}
	if !ok {
		utils.SendFailResponse(w, "Unauthorized request")
		return
	}

	users, err := userKeysData.Manager.GetAllSamePoolUsers(id)
	if err != nil {
		log.Error("Error getting users: ", err)
		utils.SendFailResponse(w, "error")
	} else {
		utils.SendDataResponse(w, users)
	}
}

func AddUserToChatHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		var userToAddData struct {
			Data	struct{
				User	structs.User	`json:"user"`
				Chat	structs.Chat	`json:"chat"`
			}		`json:"data"`
		}

		id, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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
		userToAddData.Data.User.Id = id
		if !mainDataStorage.Manager.AddUserToChatMembers(userToAddData.Data.Chat.ChatId, userToAddData.Data.User) {
			utils.SendFailResponse(w, "error")
			return
		}
		if !mainDataStorage.Manager.AddChatToUserChats(userToAddData.Data.Chat, userToAddData.Data.User.Username) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func SaveChatNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newChatData struct{
			Data structs.Chat `json:"data"`
		}
		_, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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
		if !mainDataStorage.Manager.EditChatName(newChatData.Data) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func UpdateLastReadMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message struct{
			Data structs.Message `json:"data"`
		}
		id, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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
		if !mainDataStorage.Manager.UpdateLastReadMessageId(message.Data, id) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func LeaveChatHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		var chat struct {
			Data structs.Chat `json:"data"`
		}
		id, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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
		if !mainDataStorage.Manager.DeleteChatFromUserChats(chat.Data, id) {
			utils.SendFailResponse(w, "error")
			return
		}
		if !mainDataStorage.Manager.DeleteUserFromChatMembers(chat.Data.ChatId, id) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func JoinUserToPool(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var pool struct{
			Data structs.Pool `json:"data"`
		}
		id, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		err = json.Unmarshal(requestData, &pool)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if userKeysData.Manager.TryPoolSignIn(pool.Data) {
			if userKeysData.Manager.UpdateUserPoolId(id, pool.Data) {
				utils.SendSuccessResponse(w)
			} else {
				utils.SendFailResponse(w, "error")
			}
		} else {
			utils.SendFailResponse(w, "error")
		}
	}
}

func CreatePoolHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var pool struct{
			Data structs.Pool	`json:"data"`
		}
		_, ok := utils.IdentifyWebOrMobileRequest(r)
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		err = json.Unmarshal(requestData, &pool)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if userKeysData.Manager.CreatePool(pool.Data) {
			utils.SendSuccessResponse(w)
		} else {
			utils.SendFailResponse(w, "error")
		}
	}
}