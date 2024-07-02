package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongo() *Mongo {
	uri := os.Getenv("MONGO_URL")

	opt := options.Client().
		SetMonitor(otelmongo.NewMonitor()).
		ApplyURI(uri)

	c, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatal(err.Error())
	}

	return &Mongo{
		Client: c,
	}
}
