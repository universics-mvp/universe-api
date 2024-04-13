package yandex_language_model

type CompletionOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"maxTokens"`
}

func NewCompletionOptions(temperature float32, maxTokens int) CompletionOptions {
	return CompletionOptions{
		Stream:      false,
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}
}
