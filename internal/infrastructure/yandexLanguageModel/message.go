package yandex_language_model

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

func NewMessage(role string, text string) Message {
	return Message{
		Role: role,
		Text: text,
	}
}
