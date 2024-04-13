package session

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID     *primitive.ObjectID
	ChatId int64
}

func CreateSession(chatId int64) *Session {
	return &Session{ChatId: chatId}
}
