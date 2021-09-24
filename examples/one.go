package examples

import (
	"context"
	"log"
	"os"
	"time"

	mongodb "github.com/lucassimon/gomongodb_adapter/apps"
	"github.com/lucassimon/gomongodb_adapter/apps/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type SomeService struct {
	Repository *repositories.MongoRepository
}

func NewSomeService(repository *repositories.MongoRepository) *SomeService {
	return &SomeService{
		Repository: repository,
	}
}

func (c *SomeService) count(filter interface{}) int64 {
	log.Println("Services count messages")

	total, err := c.Repository.Count(filter)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Total", total)

	return total
}

func example_one() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoConfig := mongodb.NewMongoConfig(os.Getenv("DB_NAME"))
	mongoConn := mongodb.NewMongoConnection(ctx, mongoConfig)
	defer mongoConn.Close(ctx)

	db := mongoConn.DB()
	collection := db.Collection(os.Getenv("COLLECTION_NAME"))

	repo := repositories.NewMongoRepository(collection, ctx)
	service := NewSomeService(repo)

	filter := bson.M{}
	service.count(filter)
}
