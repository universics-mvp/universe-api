package language_model_domain

type LanguageModel interface {
	DetectTokens(msg string, tokens []string) []string
	GetAnswer(msg string, promt string) []string
}
