package question_controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewQuestionRoutes),
	fx.Provide(NewQuestionController),
)
