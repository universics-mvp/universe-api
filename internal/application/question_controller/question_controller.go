package question_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
}

func NewQuestionController() QuestionController {
	return QuestionController{}
}

func (controller QuestionController) AskQuestion (c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message" : "ok", 
	})

	return
}
