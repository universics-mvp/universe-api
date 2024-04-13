package challenge_domain

import (
	"main/pkg"
)

type ChallengeService struct {
	repositoty ChallengeRepository
	logger     pkg.Logger
}

func NewChallengeService(repositoty ChallengeRepository, logger pkg.Logger) ChallengeService {
	return ChallengeService{
		repositoty: repositoty,
		logger:     logger,
	}
}

func (c *ChallengeService) GetChallenges() ([]DailyChallenge, error) {
	result, err := c.repositoty.GetChallenges()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ChallengeService) CreateChallenge(challenge DailyChallenge) (*DailyChallenge, error) {
	result, err := c.repositoty.CreateChallenge(challenge)
	if err != nil {
		return nil, err
	}

	return result, nil
}
