package tgbot

import (
	"log"
	"main/internal/config"
	"main/pkg"
	"time"

	"gopkg.in/telebot.v3"
)

type TGBot struct {
	logger          pkg.Logger
	messageHandlers []func(newMessage telebot.Message)
	bot             telebot.Bot
}

func (bot TGBot) SendMessage(chatId int64, message string) {
	bot.bot.Send(bot.bot.ChatByID(chatId))
}

func (bot *TGBot) AddMessageHandler(handler func(newMessage telebot.Message)) {
	bot.messageHandlers = append(bot.messageHandlers, handler)
}

func (bot TGBot) HandleNextMessage(c telebot.Context) error {
	for _, handler := range bot.messageHandlers {
		handler(*c.Message())
	}
	return nil
}

func NewTGBot(env config.Env, logger pkg.Logger) TGBot {
	pref := telebot.Settings{
		Token:  env.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	instance := TGBot{
		logger:          logger,
		messageHandlers: make([]func(newMessage telebot.Message), 0),
	}

	b.Handle(telebot.OnText, instance.HandleNextMessage)

	b.Start()
	return instance
}
