package question_controller

import (
	language_model_domain "main/internal/domain/languageModel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	languageModel language_model_domain.LanguageModel
}

func NewQuestionController(languageModel language_model_domain.LanguageModel) QuestionController {
	return QuestionController{
		languageModel: languageModel,
	}
}

func (controller QuestionController) AskQuestion(c *gin.Context) {
	var dto *QuestionDTO

	err := c.BindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	msg := dto.UserMessage
	answer, err := controller.languageModel.GetAnswer(msg, standardPromt, standardTemperature)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": answer,
	})
}
