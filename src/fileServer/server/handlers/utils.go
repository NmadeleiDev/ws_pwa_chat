package handlers

import (
	"fileServer/types"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SendDataResponse(w http.ResponseWriter, data interface{}) {
	var packet []byte
	var err error

	response := &types.HttpResponse{Status: true, Data: data}
	if packet, err = json.Marshal(response); err != nil {
		logrus.Error("Error marshalling response: ", err)
	}
	if _, err = w.Write(packet); err != nil {
		logrus.Error("Error sending response: ", err)
	}
}

func SendFailResponse(w http.ResponseWriter, data interface{}) {
	var packet []byte
	var err error

	response := &types.HttpResponse{Status: false, Data: data}
	if packet, err = json.Marshal(response); err != nil {
		logrus.Error("Error marshalling response: ", err)
	}
	if _, err = w.Write(packet); err != nil {
		logrus.Error("Error sending response: ", err)
	}
}
