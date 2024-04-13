package messaging

import (
	"main/internal/domain/message"
	"main/internal/domain/session"
	"main/pkg"

	"gopkg.in/telebot.v3"
)

type IMService struct {
	bot         TGBot
	logger      pkg.Logger
	msgRepo     message.MessageRepository
	sessService session.SessionService
}

func NewIMService(logger pkg.Logger, bot TGBot, msgRepo message.MessageRepository, sessService session.SessionService) IMService {
	return IMService{
		bot:         bot,
		logger:      logger,
		sessService: sessService,
		msgRepo:     msgRepo,
	}
}

func (svc IMService) HandleMessage(receivedMessage telebot.Message) error {
	chatSession, err := svc.sessService.GetOrCreateSessionForChat(receivedMessage.Chat.ID)
	if err != nil {
		return err
	}
	message := message.FromTGMessage(receivedMessage, *chatSession.ID)
	_, err = svc.msgRepo.Save(&message)
	svc.bot.SendMessage(receivedMessage.Chat.ID, receivedMessage.Text)
	return err
}
