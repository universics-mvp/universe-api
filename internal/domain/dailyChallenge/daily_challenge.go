package challenge_domain

import (
	"go.uber.org/fx"
)

type ChallengeRepository interface {
	GetChallenges() ([]DailyChallenge, error)
}

var Module = fx.Options(
	fx.Provide(NewChallengeService),
)

type DailyChallenge struct {
	ID          string   `bson:"_id" json:"id"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	CreatorId   string   `bson:"creator_id" json:"creator_id"`
	Groups      []string `bson:"groups" json:"groups"`
}
