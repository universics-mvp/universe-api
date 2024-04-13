package challenge_answer_application

import "main/pkg"

type ChallengeAnswerRoutes struct {
	ChallengeAnswerController ChallengeAnswerController
	handler                   pkg.RequestHandler
	logger                    pkg.Logger
}

func NewChallengeAnswerRoutes(
	ChallengeAnswerController ChallengeAnswerController, handler pkg.RequestHandler, logger pkg.Logger,
) ChallengeAnswerRoutes {
	return ChallengeAnswerRoutes{
		ChallengeAnswerController: ChallengeAnswerController,
		handler:                   handler,
		logger:                    logger,
	}
}

func (r ChallengeAnswerRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/challenge_answers")
	r.handler.Gin.GET("/api/v1/challenge/:challenge_id/answers", r.ChallengeAnswerController.GetChallengeAnswers)
	group.POST("", r.ChallengeAnswerController.CreateChallenge)
	group.PUT("/:id", r.ChallengeAnswerController.UpdateChallengeAnswerStatus)
}
