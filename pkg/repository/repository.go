package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	*FloodMongo
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		FloodMongo: NewFloodMongo(db),
	}
}
