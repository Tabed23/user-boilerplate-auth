package store

import (
	"context"
	"log"
	"time"

	"github.com/tabed23/user-boilerplate-auth/models"
	"github.com/tabed23/user-boilerplate-auth/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	coll mongo.Collection
}

func NewuserStore(coll mongo.Collection) repository.UserRepositry {
	return &UserStore{coll: coll}
}

func (s *UserStore) CreaterUser(ctx context.Context, usr models.User) (*mongo.InsertOneResult, error) {
	res, err := s.coll.InsertOne(ctx, usr)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UserStore) IsExist(ctx context.Context, value, data string) (bool, error) {
	count, err := s.coll.CountDocuments(ctx, bson.M{value: data})
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil

}

func (s *UserStore) Get(c context.Context) ([]models.User, error) {
	user := []models.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for data.Next(ctx) {
		var u models.User
		if err := data.Decode(&u); err != nil {
			log.Fatal(err)
			return nil, err
		}
		user = append(user, u)
	}

	return user, nil
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	u := models.User{}
	s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	return u, nil
}

func (s *UserStore) DeleteUserByEmail(ctx context.Context, email string) (bool, error) {
	_, err := s.coll.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *UserStore) UpdateUser(ctx context.Context, email string, user models.UserUpdate) (models.User, error) {
	update := bson.M{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"phone":      user.Phone,
	}
	_, err := s.coll.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": update})
	if err != nil {
		return models.User{}, err
	}

	return s.GetUserByEmail(ctx, email)
}

func (s *UserStore) GetUserByValue(ctx context.Context, value, data string) (models.User, error) {
	u := models.User{}
	s.coll.FindOne(ctx, bson.M{value: data}).Decode(&u)
	return u, nil
}
