package messaging

import (
	"main/internal/domain/message"
	"main/internal/domain/session"
	"main/pkg"

	"gopkg.in/telebot.v3"
)

type ChatService struct {
	bot         TGBot
	logger      pkg.Logger
	msgRepo     message.MessageRepository
	sessService session.SessionService
}

func NewChatService(logger pkg.Logger, bot TGBot, msgRepo message.MessageRepository, sessService session.SessionService) ChatService {
	return ChatService{
		bot:         bot,
		logger:      logger,
		sessService: sessService,
		msgRepo:     msgRepo,
	}
}

func (svc ChatService) HandleChatMessage(receivedMessage telebot.Message) error {
	chatSession, err := svc.sessService.GetOrCreateSessionForChat(receivedMessage.Chat.ID)
	if err != nil {
		return err
	}
	message := message.FromTGMessage(receivedMessage, *chatSession.ID)
	_, err = svc.msgRepo.Save(&message)
	return err
}
