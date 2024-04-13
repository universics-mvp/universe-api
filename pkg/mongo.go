package pkg

import (
	"context"
	"fmt"

	"main/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	*mongo.Client
}

type MongoDatabase struct {
	*mongo.Database
}

func NewMongoDatabase(logger Logger, env config.Env) MongoDatabase {
	ctx := context.Background()

	cl, err := mongo.Connect(ctx, &options.ClientOptions{
		Auth: &options.Credential{
			Username: env.DBUser,
			Password: env.DBPass,
		},
		Hosts: []string{
			fmt.Sprint(env.DBHost, ":", env.DBPort),
		},
	})
	if err != nil {
		logger.Fatal(err)
	}

	err = cl.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err)
	}

	db := cl.Database(env.DBName)
	return MongoDatabase{
		db,
	}
}
