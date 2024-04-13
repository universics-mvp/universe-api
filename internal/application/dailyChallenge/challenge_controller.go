package challenge_application

import (
	"net/http"

	"main/pkg"

	application_common "main/internal/application/common"
	challenge_domain "main/internal/domain/dailyChallenge"

	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
	challengeService challenge_domain.ChallengeService
	logger           pkg.Logger
	roleValidator    pkg.RoleValidator
}

func NewChallengeController(challengeService challenge_domain.ChallengeService, logger pkg.Logger) ChallengeController {
	return ChallengeController{
		roleValidator:    pkg.RoleValidator{},
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
	ok := controller.roleValidator.Validate(ctx.GetHeader("role"), application_common.CuratorRole)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})

		return
	}

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

type GetVariantsMessage struct {
	Message string `json:"message"`
}

func (controller *ChallengeController) GetVariantsForChallenge(ctx *gin.Context) {
	var dto GetVariantsMessage
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	result, err := controller.challengeService.GetVariantsForChallenge(dto.Message)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
