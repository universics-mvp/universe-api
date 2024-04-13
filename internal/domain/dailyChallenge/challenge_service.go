package challenge_domain

import (
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
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

func (c *ChallengeService) GetChallenges() ([]bson.M, error) {
	result, err := c.repositoty.GetChallenges()
	if err != nil {
		return nil, err
	}

	return result, nil
}
