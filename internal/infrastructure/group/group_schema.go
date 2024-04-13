package group_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroupSchema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CuratorID int64              `json:"curator_id" bson:"curator_id"`
	ChatID    int64              `json:"chat_id" bson:"chat_id"`
	Title     string             `json:"title" bson:"title"`
}
