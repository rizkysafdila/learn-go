package repository

import "learn-go/internal/model"

var messages = []model.Message{
    {ID: 1, Content: "Hello, World!"},
    {ID: 2, Content: "Go is awesome!"},
    {ID: 3, Content: "Go is awesome!"},
}

func GetAllMessages() []model.Message {
    return messages
}
