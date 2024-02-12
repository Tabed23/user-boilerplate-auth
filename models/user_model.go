package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"id"`
	FirstName string             `json:"first_name" validate:"required,min=3,max=100"`
	LastName  string             `json:"last_name" validate:"required,min=3,max=100"`
	Password  string             `json:"password" validate:"required,min=6"`
	Email     string             `json:"email" validate:"email,required"`
	Phone     string             `json:"phone" validate:"required"`
	Token     string             `json:"token"`
	Role      string             `json:"role" validate:"required,eq=ADMIN|eq=USER"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdate struct {
	FirstName string             `json:"first_name" validate:"required,min=3,max=100"`
	LastName  string             `json:"last_name" validate:"required,min=3,max=100"`
	Phone     string             `json:"phone" validate:"required"`
}