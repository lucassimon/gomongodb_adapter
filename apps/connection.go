package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection interface {
	Close(ctx context.Context)
	Ping(ctx context.Context) error
	DB() *mongo.Database
}

type MongoDBClient struct {
	client *mongo.Client
	config *MongoConfig
}

func NewMongoConnection(ctx context.Context, setting *MongoConfig) MongoDBConnection {
	uri := setting.Dsn()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	mg := MongoDBClient{
		client: client,
		config: setting,
	}
	return &mg
}
