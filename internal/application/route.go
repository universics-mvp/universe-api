package route

import (
	"main/docs"
	challenge_answer_application "main/internal/application/challengeAnswer"
	challenge_application "main/internal/application/dailyChallenge"
	question_controller "main/internal/application/question"

	"go.uber.org/fx"
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
	challengeAnswerRoutes challenge_answer_application.ChallengeAnswerRoutes,
	docRoutes docs.SwaggerRoutes,
) Routes {
	return Routes{
		questionRoutes,
		challengeRoutes,
		challengeAnswerRoutes,
		docRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
