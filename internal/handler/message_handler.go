package handler

import (
	"encoding/json"
	"net/http"

	"learn-go/internal/model"
	"learn-go/internal/service"
)

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    messages, err := service.GetMessages()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(model.APIResponse{
            Message: "Failed to fetch messages",
            Status:  "error",
            Data:    nil,
        })
        return
    }

    json.NewEncoder(w).Encode(model.APIResponse{
        Message: "Success",
        Status:  "success",
        Data:    messages,
    })
}