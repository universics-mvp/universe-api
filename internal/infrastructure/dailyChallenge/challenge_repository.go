package challenge_infrastructure

import (
	"context"

	challenge_domain "main/internal/domain/dailyChallenge"
	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChallengeMongoRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewChallengeMongoRepository(db pkg.MongoDatabase, logger pkg.Logger) challenge_domain.ChallengeRepository {
	return ChallengeMongoRepository{
		collection: db.Collection("challenges"),
		logger:     logger,
	}
}

func (c ChallengeMongoRepository) GetChallenges() ([]bson.M, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []bson.M

	err = cur.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
