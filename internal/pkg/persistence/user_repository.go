package persistence

import (
	"context"
	"os"

	"github.com/akselarzuman/go-jaeger/internal/pkg/db"
	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserRepository struct {
	mongo                        *db.Mongo
	db                           string
	collection                   *mongo.Collection
	secondaryPreferredCollection *mongo.Collection
}

type UserRepositoryInterface interface {
	Add(ctx context.Context, user *models.User) error
}

func NewUserRepository() *UserRepository {
	m := db.NewMongo()

	if m == nil {
		return nil
	}

	return &UserRepository{
		mongo:      m,
		db:         os.Getenv("JAEGER_SERVICE_NAME"),
		collection: m.Client.Database(os.Getenv("JAEGER_SERVICE_NAME")).Collection("users"),
		secondaryPreferredCollection: m.Client.Database(os.Getenv("JAEGER_SERVICE_NAME")).Collection("users", &options.CollectionOptions{
			ReadPreference: readpref.SecondaryPreferred(),
		}),
	}
}

func (r *UserRepository) Add(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
