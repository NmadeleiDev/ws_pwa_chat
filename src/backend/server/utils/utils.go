package utils

import (
	"chat_backend/db/userKeysData"
	"chat_backend/hashes"
	"chat_backend/structs"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	oneDayInSeconds = 86400
)

func IdentifyWebOrMobileRequest(r *http.Request) (string, bool) {
	var ok bool
	var id string

	if r.Header.Get("mobile") == "true" {
		id, ok = AuthMobileToken(r)
	} else {
		id, ok = AuthWebRequest(r)
	}
	return id, ok
}

func AuthWebRequest(r *http.Request) (string, bool) {
	sessionKey := GetCookieValue(r, "session_id")
	id, ok := userKeysData.Manager.IdentifyUserByCookie(sessionKey)
	if !ok {
		log.Warnf("Not ok authWeb. Id: %v; Cookie: %v", id, sessionKey)
		return "", false
	}
	//SetCookie(w, "session_id", newCookie)
	return id, true
}

func GenerateSecret(user structs.User) string {
	secretString := user.Username + "_hey_" + time.Now().String() + string(rand.Int63())
	return hashes.CalculateSha1(secretString)
}

func SetCookie(w *http.ResponseWriter, cookieName, value string) {
	c := http.Cookie{
		Name: cookieName,
		Value: value,
		Path: "/",
		//SameSite: http.SameSiteNoneMode,
		MaxAge: oneDayInSeconds * 1}
	http.SetCookie(*w, &c)
}

func GetCookieValue(r *http.Request, cookieName string) string {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		log.Println("Failed getting cookie", err)
		return ""
	} else {
		log.Println("Got cookie: ", cookie)
	}
	return cookie.Value
}

func SendFailResponse(w http.ResponseWriter, data interface{}) {
	var packet	[]byte
	var err		error

	response := &structs.HttpResponse{Status: false, Data:data}
	if packet, err = json.Marshal(response); err != nil {
		log.Error("Error marshalling response: ", err)
	}
	if _, err = w.Write(packet); err != nil {
		log.Error("Error sending response: ", err)
	}
}

func SendSuccessResponse(w http.ResponseWriter) {
	var packet	[]byte
	var err		error

	response := &structs.HttpResponse{Status: true, Data:nil}
	if packet, err = json.Marshal(response); err != nil {
		log.Error("Error marshalling response: ", err)
	}
	if _, err = w.Write(packet); err != nil {
		log.Error("Error sending response: ", err)
	}
}

func SendDataResponse(w http.ResponseWriter, data interface{}) {
	var packet	[]byte
	var err		error

	response := &structs.HttpResponse{Status: true, Data:data}
	if packet, err = json.Marshal(response); err != nil {
		log.Error("Error marshalling response: ", err)
	}
	if _, err = w.Write(packet); err != nil {
		log.Error("Error sending response: ", err)
	}
}

func AuthMobileToken(r *http.Request) (string, bool) {
	requestContent := struct {
		Auth	structs.MobileToken		`json:"auth"`
	}{}
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("Can't read request body for login: ", err)
		return "", false
	}
	timestamp := r.Header.Get("Event-date")
	err = json.Unmarshal(requestData, &requestContent)
	if err != nil {
		log.Error("Can't parse request body for login: ", err)
		return "", false
	}
	log.Infof("event time: %v; token: %v", timestamp, requestContent.Auth)
	id, ok := userKeysData.Manager.IdentifyUserByToken(requestContent.Auth, timestamp)
	if ok {
		return id, true
	} else {
		return "", false
	}
}
