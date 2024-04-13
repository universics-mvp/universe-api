package message

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/telebot.v3"
)

type Message struct {
	ID        *primitive.ObjectID
	ChatId    int64
	SessionId primitive.ObjectID
	MessageID int
	Text      string
	UserID    int64
	Date      int64
}

func CreateMessage(
	chatId int64,
	sessionId primitive.ObjectID,
	text string,
	messageID int,
	userID int64,
	date int64,
) Message {
	return Message{ChatId: chatId, SessionId: sessionId, Text: text, MessageID: messageID, UserID: userID, Date: date}
}

func FromTGMessage(msg telebot.Message, sessionID primitive.ObjectID) Message {
	return Message{
		SessionId: sessionID, ChatId: msg.Chat.ID, Text: msg.Text, MessageID: msg.ID, UserID: msg.Sender.ID, Date: msg.Unixtime,
	}
}
