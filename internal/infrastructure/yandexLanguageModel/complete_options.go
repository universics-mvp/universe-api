package yandex_language_model

type CompleteOptions struct {
	stream      bool    `json:"stream"`
	temperature float32 `json:"temperature"`
	maxTokens   int     `json:"maxTokens"`
}

func NewCompleteOptions(temperature float32, maxTokens int) CompleteOptions {
	return CompleteOptions{
		stream:      false,
		temperature: temperature,
		maxTokens:   maxTokens,
	}
}
