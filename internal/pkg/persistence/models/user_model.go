package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Surname   string             `bson:"surname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Active    bool               `bson:"active"`
}

type UserPostgresModel struct {
	gorm.Model
	Name      string
	Surname   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
