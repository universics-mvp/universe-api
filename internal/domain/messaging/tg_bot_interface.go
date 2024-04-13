package messaging

import "gopkg.in/telebot.v3"

type TGBot interface {
	Run()
	SendMessage(chatId int64, message string) error
	AddMessageHandler(handler func(newMessage telebot.Message))
	HandleNextMessage(c telebot.Context) error
}
