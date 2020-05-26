package utils

import (
	"chat_backend/db/postgres"
	"chat_backend/structs"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

const (
	oneDayInSeconds = 86400
)

func ValidateRequest(w http.ResponseWriter, r *http.Request) bool {
	sessionKey := GetCookieValue(r, "session_id")
	_, err := postgres.GetUserNameIdAndPool(sessionKey)
	if err != nil {
		log.Error("Error getting user data from postgres: ", err)
		SendFailResponse(w)
		return false
	}
	return true
}

func GenerateSecret(user structs.User) string {
	secretString := user.Username + "_hey_" + time.Now().String() + string(rand.Int63())
	return CalculateSha1(secretString)
}

func CalculateSha256(value string) string {
	sha256Calculate := sha256.Sum256([]byte(value))
	hash := fmt.Sprintf("%x", sha256Calculate)
	return hash
}

func CalculateSha1(value string) string {
	sha256Calculate := sha1.Sum([]byte(value))
	hash := fmt.Sprintf("%x", sha256Calculate)
	return hash
}

func SetCookie(w http.ResponseWriter, cookieName, value string) {
	c := http.Cookie{
		Name: cookieName,
		Value: value,
		Path: "/",
		//SameSite: http.SameSiteNoneMode,
		MaxAge: oneDayInSeconds * 1}
	http.SetCookie(w, &c)
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

func SendFailResponse(w http.ResponseWriter) {
	var packet	[]byte
	var err		error

	response := &structs.HttpResponse{Status: false, Data:nil}
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

func RefreshRequestSessionKeyCookie(w http.ResponseWriter, user structs.User) bool {
	sessionKey, err := postgres.IssueUserSessionKey(user)

	if err != nil {
		SendFailResponse(w)
		return false
	}

	SetCookie(w, "session_id", sessionKey)
	return true
}