package challenge_infrastructure

import (
	"context"

	"main/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChallengeMongoRepository struct {
	collection *mongo.Collection
	logger     pkg.Logger
}

func NewChallengeMongoRepository(db pkg.MongoDatabase, logger pkg.Logger) ChallengeMongoRepository {
	return ChallengeMongoRepository{
		collection: db.Collection("challenges"),
		logger:     logger,
	}
}

func (c *ChallengeMongoRepository) GetChallenges() (bson.M, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.logger.Error(err)

		return nil, err
	}
	defer cur.Close(context.Background())

	var result bson.M

	err = cur.Decode(&result)
	if err != nil {
		c.logger.Error(err)

		return nil, err
	}

	return result, nil
}
