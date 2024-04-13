package categorization

import (
	"fmt"
	"strings"

	language_model_domain "main/internal/domain/languageModel"
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
	joinedString := strings.Join(tokens, ", ")
	promt := fmt.Sprintf(categoriesPromt, joinedString)
	answer, err := c.languageModel.GetAnswer(message, promt, categoriesTemperature)
	if err != nil {
		return nil, err
	}

	tokens = append(tokens, otherCategory)

	foundTokens := make([]string, 0)

	for _, token := range tokens {
		if strings.Contains(answer, token) {
			foundTokens = append(foundTokens, token)
		}
	}

	return foundTokens, nil
}
