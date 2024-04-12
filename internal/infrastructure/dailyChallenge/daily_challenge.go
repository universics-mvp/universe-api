package challenge_infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewChallengeMongoRepository),
)
