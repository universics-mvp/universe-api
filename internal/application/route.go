package route

import (
	"go.uber.org/fx"
	challenge_application "main/internal/application/dailyChallenge"
	question_controller "main/internal/application/question"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Routes []Route

// Route interface.
type Route interface {
	Setup()
}

// here should return routers.
func NewRoutes(
	questionRoutes question_controller.QuestionRoutes,
	challengeRoutes challenge_application.ChallengeRoutes,
) Routes {
	return Routes{
		questionRoutes,
		challengeRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
