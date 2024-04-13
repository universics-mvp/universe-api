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
	mapper     ChallengeMapper
}

func NewChallengeMongoRepository(db pkg.MongoDatabase, logger pkg.Logger) challenge_domain.ChallengeRepository {
	return ChallengeMongoRepository{
		collection: db.Collection("challenges"),
		logger:     logger,
		mapper:     NewChallengeMapper(),
	}
}

func (c ChallengeMongoRepository) GetChallenges() ([]challenge_domain.DailyChallenge, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var result []ChallengeSchema

	err = cur.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	var challenges []challenge_domain.DailyChallenge

	for _, challenge := range result {
		challenges = append(challenges, *c.mapper.SchemaToEntity(challenge))
	}
	return challenges, nil
}
