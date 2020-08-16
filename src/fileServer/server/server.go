package server

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start(port string) {

	router := mux.NewRouter()

	//router.HandleFunc("/connect", handlers.ChatSocketHandler)

	log.Info("Listening ", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server error: ", err)
	}
}
