package chatbot

import (
	"main/internal/domain/messaging"
	tg_bot "main/internal/infrastructure/tgBot"
	"main/pkg"

	"gopkg.in/telebot.v3"
)

type TgHandler struct {
	imService   messaging.IMService
	chatService messaging.ChatService
	logger      pkg.Logger
}

func NewTGHandler(service messaging.IMService, chatService messaging.ChatService, tgBot tg_bot.TGBot) TgHandler {
	handler := TgHandler{
		imService:   service,
		chatService: chatService,
	}
	tgBot.AddMessageHandler(handler.handleMessage)
	return handler
}

func (handler TgHandler) handleMessage(msg telebot.Message) {
	if msg.Chat.Type == telebot.ChatGroup {
		err := handler.chatService.HandleChatMessage(msg)
		handler.logger.Error(err)
		return
	}
	err := handler.imService.HandleMessage(msg)
	handler.logger.Error(err)
}
