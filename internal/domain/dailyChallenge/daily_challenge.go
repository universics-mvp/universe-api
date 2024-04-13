package challenge_domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
)

type ChallengeRepository interface {
	GetChallenges() ([]bson.M, error)
}

var Module = fx.Options(
	fx.Provide(NewChallengeService),
)
