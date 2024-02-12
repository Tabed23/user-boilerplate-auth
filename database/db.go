package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	db *mongo.Client
}

func NewDatabase() (*Mongodb, error) {

	uri := "mongodb://localhost:27017"

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &Mongodb{db}, nil

}

func (m *Mongodb) Close() {
	m.db.Disconnect(context.TODO())
}

func (m *Mongodb) GetDB() *mongo.Database {
	return m.db.Database("restaurantmanagementsystem")
}
