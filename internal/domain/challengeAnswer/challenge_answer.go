package challenge_answer_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
)

type ChallengeAnswer struct {
	ID          primitive.ObjectID `json:"id"`
	ChallengeId primitive.ObjectID `json:"challenge_id" binding:"required"`
	UserId      string             `json:"user_id"      binding:"required"`
	Answer      string             `json:"answer"       binding:"required"`
}

type ChallengeAnswerRepository interface {
	GetChallengeAnswers(id primitive.ObjectID) ([]ChallengeAnswer, error)
	CreateChallengeAnswer(challengeAnswer ChallengeAnswer) (*ChallengeAnswer, error)
}

var Module = fx.Options(
	fx.Provide(NewChallengeAnswerService),
)
