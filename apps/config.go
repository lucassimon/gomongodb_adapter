package mongodb

import (
	"fmt"
	"log"
	"os"
)

type ConfigMongo interface {
	Dsn() string
	DbName() string
}

type MongoConfig struct {
	user string
	pass string
	host string
	name string
	dsn  string
}

func NewMongoConfig(dbName string) *MongoConfig {
	var cfg MongoConfig
	cfg.user = os.Getenv("MONGODB_USER")
	cfg.pass = os.Getenv("MONGODB_PASS")
	cfg.host = os.Getenv("MONGODB_HOST")
	cfg.name = dbName

	cfg.dsn = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", cfg.user, cfg.pass, cfg.host, cfg.name)
	return &cfg
}

func (cfg *MongoConfig) Dsn() string {
	log.Printf("Get the connection uri mongodb+srv://*******:*****@%s/%s?retryWrites=true&w=majority", cfg.host, cfg.name)
	return cfg.dsn
}

func (cfg *MongoConfig) DbName() string {
	return cfg.name
}
