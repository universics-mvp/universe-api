package message_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type MessageSchema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Text      string             `json:"text" bson:"text"`
	SessionID primitive.ObjectID `json:"session_id" bson:"session_id"`
	MessageID int                `json:"message_id" bson:"message_id"`
	ChatID    int64              `json:"chat_id" bson:"chat_id"`
	UserID    int64              `json:"user_id" bson:"user_id"`
	Date      int64              `json:"date" bson:"date"`
}
