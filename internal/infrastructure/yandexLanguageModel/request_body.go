package yandex_language_model

import "fmt"

type RequestBody struct {
	ModelUri          string            `json:"modelUri"`
	CompletionOptions CompletionOptions `json:"completionOptions"`
	Messages          []Message         `json:"messages"`
}

func getModelUri(token string) string {
	return fmt.Sprintf("gpt://%s/yandexgpt-lite", token)
}

func NewRequestBody(token string, completionOptions CompletionOptions, messages []Message) RequestBody {
	modelUri := getModelUri(token)
	return RequestBody{
		ModelUri:          modelUri,
		CompletionOptions: completionOptions,
		Messages:          messages,
	}
}
