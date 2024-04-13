package session

import (
	"main/pkg"
)

type SessionService struct {
	logger pkg.Logger
	repo   SessionRepository
}

func NewSessionService(repo SessionRepository, logger pkg.Logger) SessionService {
	return SessionService{logger: logger, repo: repo}
}

func (svc SessionService) GetOrCreateSessionForChat(chatId int64) (*Session, error) {
	session, err := svc.repo.GetByChatId(chatId)
	if err != nil {
		svc.logger.Error(err)
		return nil, err
	}
	if session != nil {
		return session, nil
	}
	session = CreateSession(chatId)
	session, err = svc.repo.Save(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
