package router

import (
	"learn-go/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/api/messages", handler.GetMessagesHandler).Methods("GET")
    return r
}
