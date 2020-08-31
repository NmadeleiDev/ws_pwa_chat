package server

import (
	"fileServer/server/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start(port string) {

	router := mux.NewRouter()

	router.HandleFunc("/file", handlers.ManagerFileHandler)

	log.Info("Listening ", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server error: ", err)
	}
}
