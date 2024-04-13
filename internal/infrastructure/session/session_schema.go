package session_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type SessionSchema struct {
	ID     *primitive.ObjectID `json:"id" bson:"_id"`
	ChatID int64               `json:"chat_id" bson:"chat_id"`
}
