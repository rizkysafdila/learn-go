package handler

import (
	"encoding/json"
	"learn-go/internal/service"
	"net/http"
)

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    messages := service.GetMessages()
    json.NewEncoder(w).Encode(messages)
}
