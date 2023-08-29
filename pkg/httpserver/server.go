package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jiteshchawla1511/chat-app/pkg/redisrepo"
	"github.com/rs/cors"
)

func StartHttpServer() {
	redisClient := redisrepo.NewRedisClient()
	defer redisClient.Close()

	redisrepo.CreateFetchChatBetweenIndex()

	r := mux.NewRouter()
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	}).Methods(http.MethodGet)

	r.HandleFunc("/register", registerHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/verify-contact", verifyContactHandler).Methods(http.MethodPost)
	r.HandleFunc("/chat-history", chatHistoryHandler).Methods(http.MethodGet)
	r.HandleFunc("/contact-list", contactListHandler).Methods(http.MethodGet)

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)

}
