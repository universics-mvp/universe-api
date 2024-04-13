package challenge_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
)

type ChallengeRepository interface {
	GetChallenges() ([]DailyChallenge, error)
	CreateChallenge(challenge DailyChallenge) (*DailyChallenge, error)
}

var Module = fx.Options(
	fx.Provide(NewChallengeService),
)

type DailyChallenge struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	CreatorId   string             `json:"creator_id" binding:"required"`
	Groups      []string           `json:"groups"`
}
