package challenge_infrastructure

import challenge_domain "main/internal/domain/dailyChallenge"

type ChallengeMapper struct{}

func NewChallengeMapper() ChallengeMapper {
	return ChallengeMapper{}
}

func (cm ChallengeMapper) EntityToSchema(challenge *challenge_domain.DailyChallenge) ChallengeSchema {
	return ChallengeSchema{
		ID:          challenge.ID,
		Title:       challenge.Title,
		Description: challenge.Description,
		CreatorID:   challenge.CreatorId,
		Groups:      challenge.Groups,
	}
}

func (cm ChallengeMapper) SchemaToEntity(challenge ChallengeSchema) *challenge_domain.DailyChallenge {
	return &challenge_domain.DailyChallenge{
		ID:          challenge.ID,
		Title:       challenge.Title,
		Description: challenge.Description,
		CreatorId:   challenge.CreatorID,
		Groups:      challenge.Groups,
	}
}
