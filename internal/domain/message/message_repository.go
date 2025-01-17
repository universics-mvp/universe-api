package message

type MessageRepository interface {
	Save(msg *Message) (*Message, error)
	GetMessagesForChat(chatId int64, since int64) ([]Message, error)
}
