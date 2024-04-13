package challenge_answer_application

type UpdateAnswerStatusDTO struct {
	Status  string `json:"status" binding:"required" validate:"omitempty,oneof=accepted rejected"`
	Mark    *int   `json:"mark" binding:"required"`
	Comment string `json:"comment"`
}
