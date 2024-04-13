package challenge_domain

import (
	language_model_domain "main/internal/domain/languageModel"
	"main/pkg"
)

type ChallengeService struct {
	repositoty    ChallengeRepository
	logger        pkg.Logger
	languageModel language_model_domain.LanguageModel
}

const (
	generateDailyChallengePromt = `Ты - помощник куратора первокурсников. 
	В университете реализована программа ежедневных заданий для студентов. Придумай задание на сегодняшний день. 
	Задания не должны быть слишком сложными, чтобы студенты могли их выполнить за короткое время
	Примеры заданий:
	- Подготовить краткое резюме к одной из лекций
	- найти аудиторию, в которой находится кафедра
	- найти информацию о научной деятельности одного из преподавателей
	- получить наивысший балл за тест по предмету
	- найти информацию о студенческом клубе
	- перечислить все предметы, которые входят в учебный план
	
	Описание и назначение задания писать не надо, только сами идеи и сокращенное описание`

	standardTemperature = 0.5
)

func NewChallengeService(repositoty ChallengeRepository, logger pkg.Logger, languageModel language_model_domain.LanguageModel) ChallengeService {
	return ChallengeService{
		repositoty:    repositoty,
		logger:        logger,
		languageModel: languageModel,
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

func (c *ChallengeService) GetVariantsForChallenge(message string) (string, error) {
	result, err := c.languageModel.GetAnswer(message, generateDailyChallengePromt, standardTemperature)
	if err != nil {
		return "", err
	}

	return result, nil
}
