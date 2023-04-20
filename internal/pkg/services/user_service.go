package services

import (
	"context"
	"log"
	"time"

	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence"
	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence/models"
	"github.com/akselarzuman/go-jaeger/internal/pkg/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRedis              redis.UserRedisInterface
	userRepository         persistence.UserRepositoryInterface
	userRepositoryPostgres persistence.UserRepositoryPostgresInterface
}

type UserServiceInterface interface {
	Add(ctx context.Context, name, surname, email, password string) error
}

func NewUserService() *UserService {
	return &UserService{
		userRedis:              redis.NewUserRedis(),
		userRepository:         persistence.NewUserRepository(),
		userRepositoryPostgres: persistence.NewUserRepositoryPostgres(),
	}
}

func (s *UserService) Add(ctx context.Context, name, surname, email, password string) error {
	err := s.userRepository.Add(ctx, &models.User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Surname:   surname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.userRepositoryPostgres.Add(ctx, &models.UserPostgresModel{
		Name:      name,
		Surname:   surname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := s.userRedis.IncrUserCount(ctx); err != nil {
		log.Println(err.Error())
	}

	return nil
}
