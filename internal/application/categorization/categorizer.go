package categorization

import (
	language_model_domain "main/internal/domain/languageModel"
	"strings"
)

type Categorizer struct {
	languageModel language_model_domain.LanguageModel
}

func NewCategorizer(languageModel language_model_domain.LanguageModel) Categorizer {
	return Categorizer{
		languageModel: languageModel,
	}
}

func (c *Categorizer) Categorize(message string, tokens []string) ([]string, error) {
	answer, err := c.languageModel.GetAnswer(message, categoriesPromt, categoriesTemperature)
	if err != nil {
		return nil, err
	}

	foundTokens := make([]string, 0)

	for _, token := range tokens {
		if strings.Contains(answer, token) {
			foundTokens = append(foundTokens, token)
		}
	}

	return foundTokens, nil
}
