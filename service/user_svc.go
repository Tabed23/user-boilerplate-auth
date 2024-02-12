package service

import (
	"context"
	"log"
	"time"

	"github.com/tabed23/user-boilerplate-auth/models"
	"github.com/tabed23/user-boilerplate-auth/repository"
	"github.com/tabed23/user-boilerplate-auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repo repository.UserRepositry
}

func NewUserService(s repository.UserRepositry) *UserService {
	return &UserService{repo: s}
}

func (s *UserService) CreaterUser(ctx context.Context, u models.User) (*mongo.InsertOneResult, error) {

	isExist, err := s.repo.IsExist(ctx, "email", u.Email)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	if isExist {
		// do something with
		log.Fatalf("user already exist")
		return nil, nil
	}

	isExist, err = s.repo.IsExist(ctx, "phone", u.Phone)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if isExist {
		// do something with
		return nil, nil
	}

	u.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	u.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	u.ID = primitive.NewObjectID()

	// Genrate the token for the user
	u.Token, err = utils.GenrateToken(u.Role, u.Email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Password, _ = utils.HashPassword(u.Password)

	return s.repo.CreaterUser(ctx, u)
}

func (s *UserService) UpdateUser(ctx context.Context, u models.UserUpdate) (models.User, error) {
	isExist, err := s.repo.IsExist(ctx, "phone", u.Phone)
	if err != nil {
		return models.User{}, err
	}
	if !isExist {
		return models.User{}, err
	}

	usr, err := s.repo.GetUserByValue(ctx, "phone", u.Phone)
	if err != nil {
		return models.User{}, err
	}

	return s.repo.UpdateUser(ctx, usr.Email, u)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) (bool, error) {
	return s.repo.DeleteUserByEmail(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserById(ctx context.Context, id string) {}

func (s *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.Get(ctx)
}

func (s *UserService) IsExist(ctx context.Context, value, data string) (bool, error) {
	return s.repo.IsExist(ctx, value, data)
}
