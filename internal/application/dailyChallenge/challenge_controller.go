package challenge_application

import (
	"net/http"

	"main/pkg"

	challenge_domain "main/internal/domain/dailyChallenge"

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

func (controller *ChallengeController) CreateChallenge(ctx *gin.Context) {
	var challenge challenge_domain.DailyChallenge
	if err := ctx.ShouldBindJSON(&challenge); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	result, err := controller.challengeService.CreateChallenge(challenge)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, result)
}
