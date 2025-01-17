package chatbot

import (
	"main/internal/domain/messaging"
	"main/pkg"

	"gopkg.in/telebot.v3"
)

type TgHandler struct {
	imService   messaging.IMService
	chatService messaging.ChatService
	logger      pkg.Logger
	bot         messaging.TGBot
}

func NewTGHandler(service messaging.IMService, chatService messaging.ChatService, tgBot messaging.TGBot, logger pkg.Logger) TgHandler {
	return TgHandler{
		imService:   service,
		chatService: chatService,
		bot:         tgBot,
		logger:      logger,
	}
}

func (handler TgHandler) Run() {
	go handler.bot.Run()
	handler.bot.AddMessageHandler(handler.handleMessage)
}

func (handler TgHandler) handleMessage(msg telebot.Message) {
	handler.logger.Debugf("%s: %s", msg.Sender.Username, msg.Text)
	if msg.Chat.Type == telebot.ChatSuperGroup {
		err := handler.chatService.HandleChatMessage(msg)
		if err != nil {
			handler.logger.Error(err)
		}
		return
	}
	err := handler.imService.HandleMessage(msg)
	if err != nil {
		handler.logger.Error(err)
	}
}
