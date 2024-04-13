package yandex_language_model

import "fmt"

type RequestBody struct {
	modelUri          string            `json:"modelUri"`
	completionOptions CompletionOptions `json:"completionOptions"`
	messages          []Message         `json:"messages"`
}

func getModelUri(token string) string {
	return fmt.Sprintf("gpt://%s/yandexgpt-lite", token)
}

func NewRequestBody(token string, completionOptions CompletionOptions, messages []Message) RequestBody {
	modelUri := getModelUri(token)
	return RequestBody{
		modelUri:          modelUri,
		completionOptions: completionOptions,
		messages:          messages,
	}
}
