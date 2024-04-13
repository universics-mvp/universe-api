package challenge_infrastructure

import challenge_domain "main/internal/domain/dailyChallenge"

type challengeMapper struct{}

func NewChallengeMapper() challengeMapper {
	return challengeMapper{}
}

func (cm challengeMapper) EntityToSchema(challenge challenge_domain.DailyChallenge) challengeSchema {
	return challengeSchema{
		ID:          challenge.ID,
		Title:       challenge.Title,
		Description: challenge.Description,
		CreatorID:   challenge.CreatorId,
		Groups:      challenge.Groups,
	}
}

func (cm challengeMapper) SchemaToEntity(challenge challengeSchema) challenge_domain.DailyChallenge {
	return challenge_domain.DailyChallenge{
		ID:          challenge.ID,
		Title:       challenge.Title,
		Description: challenge.Description,
		CreatorId:   challenge.CreatorID,
		Groups:      challenge.Groups,
	}
}
