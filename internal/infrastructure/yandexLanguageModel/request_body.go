package yandex_language_model

type RequestBody struct {
	modelUri          string            `json:"modelUri"`
	completionOptions CompletionOptions `json:"completionOptions"`
}
