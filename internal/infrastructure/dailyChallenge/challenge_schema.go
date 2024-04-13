package challenge_infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChallengeSchema struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatorID   string             `json:"creator_id" bson:"creator_id"`
	Groups      []string           `json:"groups" bson:"groups"`
}
