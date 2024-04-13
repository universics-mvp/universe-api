package tgbot

import (
	"log"
	"main/internal/config"
	"main/internal/domain/messaging"
	"main/pkg"
	"time"

	"gopkg.in/telebot.v3"
)

type TGBot struct {
	logger          pkg.Logger
	messageHandlers []func(newMessage telebot.Message)
	bot             *telebot.Bot
}

func (bot TGBot) Run() {
	bot.bot.Handle(telebot.OnText, bot.HandleNextMessage)
	bot.bot.Handle("/start", func(ctx telebot.Context) error {
		err := ctx.Reply("Hello")
		return err
	})

	go bot.bot.Start()
}

func (bot TGBot) SendMessage(chatId int64, message string) error {
	chat, err := bot.bot.ChatByID(chatId)
	if err != nil {
		return err
	}
	_, err = bot.bot.Send(chat, message)
	return err
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

func NewTGBot(env config.Env, logger pkg.Logger) messaging.TGBot {
	pref := telebot.Settings{
		Token:  env.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &TGBot{
		logger:          logger,
		messageHandlers: make([]func(newMessage telebot.Message), 0),
		bot:             b,
	}
}
