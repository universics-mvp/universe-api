package session

type SessionRepository interface {
	Save(session *Session) (*Session, error)
	GetByChatId(chatId int64) (*Session, error)
}
