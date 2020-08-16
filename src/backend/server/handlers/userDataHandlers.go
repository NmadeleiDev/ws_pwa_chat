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

func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	var ok bool
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Can't read request body for create chat: ", err)
		utils.SendFailResponse(w, "read error")
		return
	}
	id, ok = utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
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

func CreateChatHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var chatData struct {
			Data structs.ChatWithMessages `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for create chat: ", err)
			utils.SendFailResponse(w, "read error")
			return
		}
		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}

		err = json.Unmarshal(requestData, &chatData)
		if err != nil {
			log.Error("Can't parse request body for create chat: ", err)
			utils.SendFailResponse(w, "parse error")
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

func GetChatMessagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var chatData struct {
			Data structs.Chat `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}

		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	var id string
	var ok bool
	// это немного необычный запрос, он гет если от веба и пост, если от мобайла.
	if r.Method == http.MethodGet {
		id, ok = utils.AuthWebRequest(utils.GetCookieValue(r, "session_id"))
	} else if r.Method == http.MethodPost {
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		id, ok = utils.AuthMobileToken(requestData, r.Header.Get("Event-date"))
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

func AddUserToChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userToAddData struct {
			Data struct {
				User structs.User `json:"user"`
				Chat structs.Chat `json:"chat"`
			} `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}

		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}

		err = json.Unmarshal(requestData, &userToAddData)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		userToAddData.Data.User.Id, ok = userKeysData.Manager.GetUserIdByName(userToAddData.Data.User.Username)
		if !ok {
			utils.SendFailResponse(w, "user to add not found")
			return
		}
		if !mainDataStorage.Manager.AddUserToChatMembers(userToAddData.Data.Chat.ChatId, userToAddData.Data.User) {
			utils.SendFailResponse(w, "error")
			return
		}
		if !mainDataStorage.Manager.AddChatToUserChats(structs.ChatInfo{ChatId: userToAddData.Data.Chat.ChatId, MessagePoolId: userToAddData.Data.Chat.MessagePoolId}, []string{userToAddData.Data.User.Username}) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func SaveChatNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newChatData struct {
			Data structs.Chat `json:"data"`
		}

		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for chat name edit: ", err)
			return
		}
		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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

func SaveChatStorePeriodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newChatData struct {
			Data structs.Chat `json:"data"`
		}

		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for chat name edit: ", err)
			return
		}
		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		err = json.Unmarshal(requestData, &newChatData)
		if err != nil {
			log.Error("Can't parse request body for chat name edit: ", err)
			return
		}
		if !mainDataStorage.Manager.EditChatStorePeriod(newChatData.Data) {
			utils.SendFailResponse(w, "error")
			return
		}
		utils.SendSuccessResponse(w)
	}
}

func UpdateLastReadMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message struct {
			Data structs.Message `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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

func LeaveChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var chat struct {
			Data structs.Chat `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
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

func ManagePoolHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // добавление пользователя в пул
		var pool struct {
			Data structs.Pool `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		err = json.Unmarshal(requestData, &pool)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if pool.Data.PoolId == "" || userKeysData.Manager.TryPoolSignIn(pool.Data) {
			if userKeysData.Manager.UpdateUserPoolId(id, pool.Data) {
				utils.SendSuccessResponse(w)
			} else {
				utils.SendFailResponse(w, "error")
			}
		} else {
			utils.SendFailResponse(w, "error")
		}
	} else if r.Method == http.MethodPut { // создание пула
		var pool struct {
			Data structs.Pool `json:"data"`
		}
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Can't read request body for login: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		err = json.Unmarshal(requestData, &pool)
		if err != nil {
			log.Error("Can't parse request body for login: ", err)
			return
		}
		if userKeysData.Manager.CreatePool(pool.Data) {
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

//func CreatePoolHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		var pool struct {
//			Data structs.Pool `json:"data"`
//		}
//		requestData, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			log.Error("Can't read request body for login: ", err)
//			return
//		}
//		_, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
//		if !ok {
//			utils.SendFailResponse(w, "Unauthorized request")
//			return
//		}
//		err = json.Unmarshal(requestData, &pool)
//		if err != nil {
//			log.Error("Can't parse request body for login: ", err)
//			return
//		}
//		if userKeysData.Manager.CreatePool(pool.Data) {
//			utils.SendSuccessResponse(w)
//		} else {
//			utils.SendFailResponse(w, "error")
//		}
//	}
//}
