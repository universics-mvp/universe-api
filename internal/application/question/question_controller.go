package question_controller

import (
	"net/http"

	"main/internal/application/categorization"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	categorizer categorization.Categorizer
}

func NewQuestionController(categorizer categorization.Categorizer) QuestionController {
	return QuestionController{
		categorizer: categorizer,
	}
}

// @Summary Ask question
// @Tags Question
// @Description ask
// @Accept json
// @Produce json
// @Param question body QuestionDTO true "question"
// @Router /api/v1/question [post]
// @Success 200 {string} string
func (controller QuestionController) AskQuestion(c *gin.Context) {
	var dto *QuestionDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	msg := dto.UserMessage
	answer, err := controller.categorizer.Categorize(msg, StandardCategoies)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": answer,
	})
}
