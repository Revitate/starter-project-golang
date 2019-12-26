package connector

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConnector struct {
	*mongo.Database
}

type MongoConnectorOption struct {
	URI      string
	Database string
}

func NewMongoConnector(option MongoConnectorOption) (*MongoConnector, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(option.URI))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &MongoConnector{client.Database(option.Database)}, nil
}
