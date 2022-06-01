package persistence

import (
	"context"

	"github.com/akselarzuman/go-jaeger/internal/pkg/db"
	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence/models"
)

type UserRepositoryPostgres struct {
	client *db.PostgresClient
}

type UserRepositoryPostgresInterface interface {
	Add(ctx context.Context, user *models.UserPostgresModel) error
}

func NewUserRepositoryPostgres() *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		client: db.NewPostgresClient(),
	}
}

func (r *UserRepositoryPostgres) Add(ctx context.Context, user *models.UserPostgresModel) error {
	return r.client.GormDB.WithContext(ctx).Create(user).Error
}
