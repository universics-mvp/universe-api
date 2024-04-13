package challenge_answer_infrastructure

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
)

type answerSchema struct {
	ID          primitive.ObjectID `bson:"_id"`
	ChallengeId primitive.ObjectID `bson:"challenge_id"`
	UserId      string             `bson:"user_id"`
	Answer      string             `bson:"answer"`
	Mark        int                `bson:"mark"`
	Status      string             `bson:"status"`
	Comment     string             `bson:"comment"`
}

var Module = fx.Options(
	fx.Provide(NewChallengeAnswerMongoRepository),
)
