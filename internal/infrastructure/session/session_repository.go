package session_infrastructure

import (
	"context"

	"main/internal/domain/session"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewSessionRepository(db pkg.MongoDatabase, logger pkg.Logger) session.SessionRepository {
	return SessionRepository{
		collection: db.Collection("sessions"),
		logger:     logger,
	}
}

func (repo SessionRepository) Save(session *session.Session) (*session.Session, error) {
	if session.ID == nil {
		return repo.create(session)
	}
	return repo.update(session)
}

func (repo SessionRepository) create(session *session.Session) (*session.Session, error) {
	res, err := repo.collection.InsertOne(context.Background(), mapToSchema(*session))
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(res.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}

	session.ID = &id

	return session, nil
}

func (repo SessionRepository) update(sess *session.Session) (*session.Session, error) {
	_, err := repo.collection.ReplaceOne(context.Background(), bson.M{"_id": sess.ID}, mapToSchema(*sess))
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (c SessionRepository) GetByChatId(chatId int64) (*session.Session, error) {
	var schema SessionSchema
	err := c.collection.FindOne(context.Background(), bson.M{"chat_id": chatId}).Decode(schema)
	if err != nil {
		return nil, err
	}
	entity := mapShema(schema)
	return &entity, nil
}

func mapShema(schema SessionSchema) session.Session {
	return session.Session{
		ID:     schema.ID,
		ChatId: schema.ChatID,
	}
}

func mapToSchema(session session.Session) SessionSchema {
	return SessionSchema{
		ID:     session.ID,
		ChatID: session.ChatId,
	}
}
