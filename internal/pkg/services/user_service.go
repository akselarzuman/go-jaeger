package services

import (
	"context"
	"log"
	"time"

	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence"
	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository persistence.UserRepositoryInterface
}

type UserServiceInterface interface {
	Add(ctx context.Context, name, surname, email, password string) error
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: persistence.NewUserRepository(),
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
	}

	return err
}
