package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Db     *mongo.Database
	Client *mongo.Client
}

func Connect(mongoURI, dbName string) (*MongoInstance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	result := &MongoInstance{
		Db:     db,
		Client: client,
	}

	return result, nil
}
