package challenge_answer_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewChallengeAnswerController),
	fx.Provide(NewChallengeAnswerRoutes),
)
