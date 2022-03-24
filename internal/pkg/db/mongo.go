package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongo() *Mongo {
	uri := os.Getenv("MONGO_URL")

	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Connect(ctx); err != nil {
		log.Fatal(err.Error())
		return nil
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatal(err.Error())
	}

	return &Mongo{
		Client: c,
	}
}
