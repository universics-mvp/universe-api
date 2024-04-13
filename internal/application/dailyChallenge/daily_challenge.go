package challenge_application

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewChallengeController),
	fx.Provide(NewChallengeRoutes),
)
