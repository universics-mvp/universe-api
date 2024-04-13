package challenge_application

import "main/pkg"

type ChallengeRoutes struct {
	challengeController ChallengeController
	handler             pkg.RequestHandler
	logger              pkg.Logger
}

func NewChallengeRoutes(
	challengeController ChallengeController, handler pkg.RequestHandler, logger pkg.Logger,
) ChallengeRoutes {
	return ChallengeRoutes{
		challengeController: challengeController,
		handler:             handler,
		logger:              logger,
	}
}

func (r ChallengeRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/challenge")
	group.GET("", r.challengeController.GetChallenges)
	group.POST("", r.challengeController.CreateChallenge)
	group.GET("/variants", r.challengeController.GetVariantsForChallenge)
}
