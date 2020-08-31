package handlers

import (
	"chat_backend/db/mainDataStorage"
	"chat_backend/db/userKeysData"
	"chat_backend/hashes"
	"chat_backend/server/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func RequestFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var fileRequest struct {
			ChatId string `json:"data"`
		}

		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error("Can't read request body for chat name edit: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		name, _ := userKeysData.Manager.GetUserNameById(id)
		err = json.Unmarshal(requestData, &fileRequest)
		if err != nil {
			logrus.Error("Can't parse request body for chat name edit: ", err)
			return
		}
		if mainDataStorage.Manager.CheckIfUserIsInChat(fileRequest.ChatId, name) {
			lotId := userKeysData.Manager.CreateFileLot(fileRequest.ChatId)
			utils.SendDataResponse(w, lotId) // да, там может прийти -1, но это проверяется на клиенте
		} else {
			utils.SendFailResponse(w, "failed to check request")
		}
	} else if r.Method == http.MethodPut {
		var fileRequest struct {
			LotId string `json:"data"`
		}

		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error("Can't read request body for chat name edit: ", err)
			return
		}
		id, ok := utils.IdentifyWebOrMobileRequest(requestData, utils.GetCookieValue(r, "session_id"), r.Header.Clone())
		if !ok {
			utils.SendFailResponse(w, "Unauthorized request")
			return
		}
		name, _ := userKeysData.Manager.GetUserNameById(id)
		err = json.Unmarshal(requestData, &fileRequest)
		if err != nil {
			logrus.Error("Can't parse request body for chat name edit: ", err)
			return
		}
		chatId := userKeysData.Manager.GetFileLotChatId(fileRequest.LotId)
		if len(chatId) == 0 {
			logrus.Errorf("Lot not found, or chat is to set")
			utils.SendFailResponse(w, "file not found")
			return
		}
		if mainDataStorage.Manager.CheckIfUserIsInChat(chatId, name) {
			key := hashes.CalculateSha1(strconv.Itoa(rand.Int()) + strconv.Itoa(time.Now().Nanosecond()+122))
			if userKeysData.Manager.AddViewKeyToFileLot(fileRequest.LotId, key) {
				utils.SendDataResponse(w, key)
			} else {
				utils.SendFailResponse(w, "failed create key")
			}
		} else {
			utils.SendFailResponse(w, "failed to check request")
		}
	}
}
