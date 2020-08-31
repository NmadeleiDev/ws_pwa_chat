package handlers

import (
	"fileServer/types"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

func parseSaveKeyToInt(key []byte) int64 {
	res, err := strconv.Atoi(string(key))
	if err != nil {
		logrus.Errorf("Error atoi savekey: %v", err)
		return 0
	}
	return int64(res)
}

func listDir(dirName string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Infof("Got files: %v", files)

	return files
}

func removeFile(file *os.File) {
	name := file.Name()
	if err := file.Close(); err != nil {
		logrus.Errorf("Error closing file: %v; name: %v", err, file.Name())
	}
	e := os.Remove(name)
	if e != nil {
		logrus.Errorf("Error deleting file: %v; name: %v", e, name)
	}
}