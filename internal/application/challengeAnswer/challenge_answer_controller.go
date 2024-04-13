package challenge_answer_application

import (
	"net/http"

	application_common "main/internal/application/common"
	challenge_answer_domain "main/internal/domain/challengeAnswer"
	"main/pkg"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChallengeAnswerController struct {
	service       challenge_answer_domain.ChallengeAnswerService
	logger        pkg.Logger
	roleValidator pkg.RoleValidator
}

func NewChallengeAnswerController(service challenge_answer_domain.ChallengeAnswerService, logger pkg.Logger) ChallengeAnswerController {
	return ChallengeAnswerController{service: service, logger: logger, roleValidator: pkg.RoleValidator{}}
}

func (controller ChallengeAnswerController) GetChallengeAnswers(ctx *gin.Context) {
	ok := controller.roleValidator.Validate(ctx.GetHeader("role"), application_common.CuratorRole)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})

		return
	}

	id, err := primitive.ObjectIDFromHex(ctx.Param("challenge_id"))

	challengeAnswers, err := controller.service.GetChallengeAnswers(id)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, challengeAnswers)
}

func (controller ChallengeAnswerController) GetUserAnswers(ctx *gin.Context) {
	userId := ctx.GetHeader("user_id")

	challengeAnswers, err := controller.service.GetChallengeAnswersByUserId(userId)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, challengeAnswers)
}

func (controller ChallengeAnswerController) AnswerToChallenge(ctx *gin.Context) {
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

func (controller ChallengeAnswerController) UpdateChallengeAnswerStatus(ctx *gin.Context) {
	ok := controller.roleValidator.Validate(ctx.GetHeader("role"), application_common.CuratorRole)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})

		return
	}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var dto UpdateAnswerStatusDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	updatedChallengeAnswer, err := controller.service.UpdateChallengeAnswerStatus(id, dto.Status, dto.Mark, &dto.Comment)
	if err != nil {
		if err.Error() == "invalid mark or status" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		controller.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, updatedChallengeAnswer)
}
