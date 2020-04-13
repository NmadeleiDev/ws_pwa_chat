package server

import (
	"chat_backend/server/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartServer(port string) {

	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.SignUpHandler)
	router.HandleFunc("/signin", handlers.SignInHandler)
	router.HandleFunc("/signout", handlers.SignOutHandler)
	router.HandleFunc("/unregister", handlers.UnregisterHandler)

	router.HandleFunc("/get_data", handlers.GetUserDataHandler)
	router.HandleFunc("/all_users", handlers.GetAllUsersHandler)

	router.HandleFunc("/get_messages/{chatId}", handlers.GetChatMessagesHandler)

	router.HandleFunc("/connect", handlers.ChatSocketHandler)

	log.Info("Listening ", port)
	if err := http.ListenAndServe(":" + port, router); err != nil {
		log.Fatal("Server error: ", err)
	}
}
