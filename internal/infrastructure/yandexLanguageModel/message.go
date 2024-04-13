package yandex_language_model

type Message struct {
	role string `json:"role"`
	text string `json:"text"`
}

func NewMessage(role string, text string) Message {
	return Message{
		role: role,
		text: text,
	}
}
