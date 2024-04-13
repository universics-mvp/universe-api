package challenge_application

import (
	"net/http"

	challenge_domain "main/internal/domain/dailyChallenge"
	"main/pkg"

	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
	challengeService challenge_domain.ChallengeService
	logger           pkg.Logger
}

func NewChallengeController(challengeService challenge_domain.ChallengeService, logger pkg.Logger) ChallengeController {
	return ChallengeController{
		challengeService: challengeService,
		logger:           logger,
	}
}

func (controller *ChallengeController) GetChallenges(ctx *gin.Context) {
	result, err := controller.challengeService.GetChallenges()
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, result)
}
