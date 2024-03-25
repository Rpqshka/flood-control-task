package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	floodControl "task"
)

type FloodMongo struct {
	db *mongo.Database
}

func NewFloodMongo(db *mongo.Database) *FloodMongo {
	return &FloodMongo{db: db}
}

func (r *FloodMongo) Check(ctx context.Context, userID int64) (bool, error) {
	currentTime := ctx.Value("time").(int64)
	floodTime := ctx.Value("k").(int64)
	requestLimit := ctx.Value("n").(int64)
	var flood = floodControl.Flood{Id: userID, Time: currentTime}

	filter := bson.M{"time": bson.M{"$gte": floodTime}}

	count, err := r.db.Collection(floodTable).CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	if count >= requestLimit {
		return false, errors.New("stop flood. try again later")
	}

	if _, err = r.db.Collection(floodTable).InsertOne(ctx, flood); err != nil {
		return false, err
	}

	return true, nil
}
