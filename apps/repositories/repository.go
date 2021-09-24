package repositories

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoRepository(collection *mongo.Collection, ctx context.Context) *MongoRepository {
	return &MongoRepository{
		collection: collection,
		ctx:        ctx,
	}
}

func (mr *MongoRepository) Count(filter interface{}) (int64, error) {
	log.Println("Counting the docs")

	total, err := mr.collection.CountDocuments(mr.ctx, filter)

	return total, err
}

func (mr *MongoRepository) FindAll(filter interface{}, findOptions *options.FindOptions) (*mongo.Cursor, error) {
	log.Println("Searching the docs")

	cursor, err := mr.collection.Find(mr.ctx, filter, findOptions)

	return cursor, err
}

func (mr *MongoRepository) FindOne(filter interface{}, options *options.FindOneOptions) *mongo.SingleResult {
	log.Println("Find one doc")

	result := mr.collection.FindOne(mr.ctx, filter, options)

	return result
}

func (mr *MongoRepository) Delete(filter interface{}, deleteOptions *options.DeleteOptions) (*mongo.DeleteResult, error) {
	log.Println("delete the doc")

	result, err := mr.collection.DeleteOne(mr.ctx, filter, deleteOptions)

	return result, err
}
