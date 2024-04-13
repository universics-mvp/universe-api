package challenge_answer_domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChallengeAnswerService struct {
	repo ChallengeAnswerRepository
}

func NewChallengeAnswerService(repo ChallengeAnswerRepository) ChallengeAnswerService {
	return ChallengeAnswerService{repo: repo}
}

func (s ChallengeAnswerService) GetChallengeAnswers(id primitive.ObjectID) ([]ChallengeAnswer, error) {
	return s.repo.GetChallengeAnswers(id)
}

func (s ChallengeAnswerService) CreateChallengeAnswer(challengeAnswer ChallengeAnswer) (*ChallengeAnswer, error) {
	_, err := s.repo.FindChallengeAnswerByChallengeIdAndUserId(challengeAnswer.ChallengeId, challengeAnswer.UserId)
	if err == nil {
		return nil, errors.New("you have already answered this challenge")
	}

	challengeAnswer.Status = StatusPending

	return s.repo.CreateChallengeAnswer(challengeAnswer)
}

func (s ChallengeAnswerService) GetChallengeAnswersByUserId(userId string) ([]ChallengeAnswer, error) {
	return s.repo.GetChallengeAnswersByUserId(userId)
}

func (s ChallengeAnswerService) UpdateChallengeAnswerStatus(id primitive.ObjectID, status string, mark *int, comment *string) (*ChallengeAnswer, error) {
	answer, err := s.repo.FindChallengeAnswer(id)
	if err != nil {
		return nil, err
	}

	answer.Comment = *comment

	if status == StatusAccepted {
		if *mark <= 0 {
			return nil, errors.New("invalid mark")
		}
	} else if status == StatusRejected {
		*mark = 0
		if comment == nil || *comment == "" {
			return nil, errors.New("comment is required if status is rejected")
		}
	} else {
		return nil, errors.New("invalid status")
	}

	answer.Status = status
	answer.Mark = *mark
	answer.Comment = *comment

	return s.repo.UpdateChallengeAnswer(*answer)
}
