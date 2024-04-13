package challenge_answer_application

import (
	"net/http"

	challenge_answer_domain "main/internal/domain/challengeAnswer"
	"main/pkg"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChallengeAnswerController struct {
	service challenge_answer_domain.ChallengeAnswerService
	logger  pkg.Logger
}

func NewChallengeAnswerController(service challenge_answer_domain.ChallengeAnswerService, logger pkg.Logger) ChallengeAnswerController {
	return ChallengeAnswerController{service: service, logger: logger}
}

func (controller ChallengeAnswerController) GetChallengeAnswers(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("challenge_id"))

	challengeAnswers, err := controller.service.GetChallengeAnswers(id)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, challengeAnswers)
}

func (controller ChallengeAnswerController) CreateChallenge(ctx *gin.Context) {
	var challengeAnswer challenge_answer_domain.ChallengeAnswer
	if err := ctx.ShouldBindJSON(&challengeAnswer); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	createdChallengeAnswer, err := controller.service.CreateChallengeAnswer(challengeAnswer)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, createdChallengeAnswer)
}
