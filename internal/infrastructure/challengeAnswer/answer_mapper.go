package challenge_answer_infrastructure

import challenge_answer_domain "main/internal/domain/challengeAnswer"

type answerMapper struct{}

func newAnswerMapper() answerMapper {
	return answerMapper{}
}

func (am answerMapper) EntityToSchema(challengeAnswer challenge_answer_domain.ChallengeAnswer) answerSchema {
	return answerSchema{
		ID:          challengeAnswer.ID,
		ChallengeId: challengeAnswer.ChallengeId,
		UserId:      challengeAnswer.UserId,
		Answer:      challengeAnswer.Answer,
	}
}

func (am answerMapper) SchemaToEntity(challengeAnswer answerSchema) challenge_answer_domain.ChallengeAnswer {
	return challenge_answer_domain.ChallengeAnswer{
		ID:          challengeAnswer.ID,
		ChallengeId: challengeAnswer.ChallengeId,
		UserId:      challengeAnswer.UserId,
		Answer:      challengeAnswer.Answer,
	}
}
