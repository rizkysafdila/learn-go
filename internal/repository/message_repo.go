package repository

import (
	"learn-go/internal/db"
	"learn-go/internal/model"
)

var messages = []model.Message{
    {ID: 1, Content: "Hello, World!"},
    {ID: 2, Content: "Go is awesome!"},
    {ID: 3, Content: "Go mantap parahh!"},
}

func GetAllMessages() ([]model.Message, error) {
    rows, err := db.DB.Query("SELECT id, content FROM messages")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []model.Message
    for rows.Next() {
        var msg model.Message
        if err := rows.Scan(&msg.ID, &msg.Content); err != nil {
            return nil, err
        }
        messages = append(messages, msg)
    }

    return messages, nil
}
