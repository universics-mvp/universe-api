package challenge_answer_infrastructure

import (
	"context"

	domain "main/internal/domain/challengeAnswer"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChallengeAnswerMongoRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
	mapper     answerMapper
}

func NewChallengeAnswerMongoRepository(
	db pkg.MongoDatabase, logger pkg.Logger,
) domain.ChallengeAnswerRepository {
	return ChallengeAnswerMongoRepository{
		collection: db.Collection("challenge_answers"),
		logger:     logger,
		mapper:     newAnswerMapper(),
	}
}

func (c ChallengeAnswerMongoRepository) GetChallengeAnswers(id primitive.ObjectID) ([]domain.ChallengeAnswer, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{"challenge_id": id})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []answerSchema

	err = cur.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	var challengeAnswers []domain.ChallengeAnswer

	for _, challengeAnswer := range result {
		challengeAnswers = append(challengeAnswers, c.mapper.SchemaToEntity(challengeAnswer))
	}

	return challengeAnswers, nil
}

func (c ChallengeAnswerMongoRepository) CreateChallengeAnswer(challengeAnswer domain.ChallengeAnswer) (*domain.ChallengeAnswer, error) {
	schema := c.mapper.EntityToSchema(challengeAnswer)
	schema.ID = primitive.NewObjectID()

	_, err := c.collection.InsertOne(context.Background(), schema)
	if err != nil {
		return nil, err
	}

	challengeAnswer.ID = schema.ID

	return &challengeAnswer, nil
}
