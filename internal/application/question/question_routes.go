package question_controller

import "main/pkg"

type QuestionRoutes struct {
	logger     pkg.Logger
	controller QuestionController
	handler    pkg.RequestHandler
}

func NewQuestionRoutes(logger pkg.Logger, controller QuestionController, handler pkg.RequestHandler) QuestionRoutes {
	return QuestionRoutes{
		logger:     logger,
		controller: controller,
		handler:    handler,
	}
}

func (r QuestionRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/question")
	group.POST("", r.controller.AskQuestion)
}
