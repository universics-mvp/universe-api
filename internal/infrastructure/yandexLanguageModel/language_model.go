package yandex_language_model

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type LanguageModel struct {
	oauthToken  string
	directoryId string
}

func NewLanguageModel(oauthToken string, directoryId string) LanguageModel {
	return LanguageModel{
		oauthToken:  oauthToken,
		directoryId: directoryId,
	}
}

func (l *LanguageModel) getIamToken() (string, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"yandexPassportOauthToken": "%s"}`, l.oauthToken)).
		Post("https://iam.api.cloud.yandex.net/iam/v1/tokens")

	if err != nil {
		return "error", err
	}

	return resp.String(), nil
}

func (l *LanguageModel) GetAnswer(msg string, promt string, temperture float32) (string, error) {
	client := resty.New()

	iamToken, err := l.getIamToken()
	if err != nil {
		return "error", err
	}

	promtMessage := NewMessage("system", promt)
	msgMessage := NewMessage("user", msg)
	complitionOptions := NewCompletionOptions(temperture, 1000)
	request := NewRequestBody(l.directoryId, complitionOptions, []Message{promtMessage, msgMessage})

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+iamToken).
		SetBody(request).
		Post("https://llm.api.cloud.yandex.net/foundationModels/v1/completion")

	if err != nil {
		return "error", err
	}

	decodedJSON := resp.Body()

	var data interface {
	}

	err = json.Unmarshal(decodedJSON, &data)
	if err != nil {
		return "error", err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "error", err
	}

	return string(jsonData), nil
}
