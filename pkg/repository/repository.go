package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Checker interface {
	SetParam(ctx context.Context) error
	GetParam() (int64, int64, error)
}

type Repository struct {
	*FloodMongo
	Checker
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		FloodMongo: NewFloodMongo(db),
		Checker:    NewFloodMongo(db),
	}
}
