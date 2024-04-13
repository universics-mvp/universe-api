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
	Mark        int                `json:"mark"`
	Status      string             `json:"status"`
	Comment     string             `json:"comment"`
}

type ChallengeAnswerRepository interface {
	GetChallengeAnswers(id primitive.ObjectID) ([]ChallengeAnswer, error)
	CreateChallengeAnswer(challengeAnswer ChallengeAnswer) (*ChallengeAnswer, error)
	UpdateChallengeAnswer(challengeAnswer ChallengeAnswer) (*ChallengeAnswer, error)
	FindChallengeAnswer(id primitive.ObjectID) (*ChallengeAnswer, error)
	GetChallengeAnswersByUserId(userId string) ([]ChallengeAnswer, error)
	FindChallengeAnswerByChallengeIdAndUserId(challengeId primitive.ObjectID, userId string) (*ChallengeAnswer, error)
}

const (
	StatusAccepted = "accepted"
	StatusRejected = "rejected"
	StatusPending  = "pending"
)

var Module = fx.Options(
	fx.Provide(NewChallengeAnswerService),
)
