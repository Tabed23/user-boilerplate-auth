package repository

import (
	"context"

	"github.com/tabed23/user-boilerplate-auth/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositry interface {
	CreaterUser(ctx context.Context, usr models.User) (*mongo.InsertOneResult, error)
	IsExist(context.Context, string, string) (bool, error)
	Get(context.Context) ([]models.User, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	DeleteUserByEmail( context.Context, string)(bool, error)
	UpdateUser(ctx context.Context, email string, user models.UserUpdate)(models.User, error)
	GetUserByValue(context.Context, string,string)(models.User, error)
}
