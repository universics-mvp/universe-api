package messaging

import "gopkg.in/telebot.v3"

type TGBot interface {
	sendMessage(chatId int, message string)
	AddMessageHandler(handler func(newMessage telebot.Message))
	HandleNextMessage(c telebot.Context) error
}
