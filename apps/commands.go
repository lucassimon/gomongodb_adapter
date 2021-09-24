package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (c *MongoDBClient) Ping(ctx context.Context) error {
	// Ping the primary
	if err := c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func (c *MongoDBClient) DB() *mongo.Database {
	return c.client.Database(c.config.DbName())
}

func (c *MongoDBClient) Close(ctx context.Context) {
	if err := c.client.Disconnect(ctx); err != nil {
		log.Panicln(err)
	}
}
