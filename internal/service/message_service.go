package service

import (
	"learn-go/internal/model"
	"learn-go/internal/repository"
)

func GetMessages() []model.Message {
    return repository.GetAllMessages()
}
