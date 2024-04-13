package yandex_language_model

type CompletionOptions struct {
	stream      bool    `json:"stream"`
	temperature float32 `json:"temperature"`
	maxTokens   int     `json:"maxTokens"`
}

func NewCompletionOptions(temperature float32, maxTokens int) CompletionOptions {
	return CompletionOptions{
		stream:      false,
		temperature: temperature,
		maxTokens:   maxTokens,
	}
}
