package challenge_answer_domain

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	return s.repo.CreateChallengeAnswer(challengeAnswer)
}
