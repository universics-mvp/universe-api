package message_infrastructure

import (
	"context"

	"main/internal/domain/message"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewMessageRepository(db pkg.MongoDatabase, logger pkg.Logger) message.MessageRepository {
	return MessageRepository{
		collection: db.Collection("messages"),
		logger:     logger,
	}
}

func (repo MessageRepository) Save(message *message.Message) (*message.Message, error) {
	if message.ID == nil {
		return repo.create(message)
	}
	return repo.update(message)
}

func (repo MessageRepository) create(message *message.Message) (*message.Message, error) {
	newId := primitive.NewObjectID()
	message.ID = &newId
	_, err := repo.collection.InsertOne(context.Background(), mapToSchema(*message))
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (repo MessageRepository) update(sess *message.Message) (*message.Message, error) {
	_, err := repo.collection.ReplaceOne(context.Background(), bson.M{"_id": sess.ID}, mapToSchema(*sess))
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (c MessageRepository) GetMessagesForChat(chatId int64, since int64) ([]message.Message, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{"chat_id": chatId, "date": bson.M{"$gte": since}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []MessageSchema

	err = cur.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	var messages = make([]message.Message, 0)
	for _, m := range result {
		messages = append(messages, mapShema(m))
	}
	return messages, nil
}

func mapShema(schema MessageSchema) message.Message {
	return message.Message{
		ID:           &schema.ID,
		ChatId:       schema.ChatID,
		MessageID:    schema.MessageID,
		Text:         schema.Text,
		SessionId:    schema.SessionID,
		UserFullName: schema.UserFullName,
		Date:         schema.Date,
		UserID:       schema.UserID,
	}
}

func mapToSchema(msg message.Message) MessageSchema {
	return MessageSchema{
		ID:           *msg.ID,
		ChatID:       msg.ChatId,
		Text:         msg.Text,
		UserFullName: msg.UserFullName,
		SessionID:    msg.SessionId,
		MessageID:    msg.MessageID,
		UserID:       msg.UserID,
		Date:         msg.Date,
	}
}
