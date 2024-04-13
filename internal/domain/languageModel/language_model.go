package language_model_domain

type LanguageModel interface {
	GetAnswer(msg string, promt string) (string, error)
}
