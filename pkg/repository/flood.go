package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	floodControl "task"
	"time"
)

type FloodMongo struct {
	db *mongo.Database
}

func NewFloodMongo(db *mongo.Database) *FloodMongo {
	return &FloodMongo{db: db}
}

func (r *FloodMongo) Check(ctx context.Context, userID int64) (bool, error) {
	currentTime := ctx.Value("time").(int64)
	n := ctx.Value("n").(int64)
	k := ctx.Value("k").(int64)
	floodTime := time.Now().Add(-time.Duration(k) * time.Second).Unix()

	var flood = floodControl.Flood{Id: userID, Time: currentTime}

	filter := bson.M{"time": bson.M{"$gte": floodTime}}

	count, err := r.db.Collection(floodTable).CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	if count >= n {
		return false, errors.New("stop flood. try again later. change n or k values")
	}

	if _, err = r.db.Collection(floodTable).InsertOne(ctx, flood); err != nil {
		return false, err
	}

	return true, nil
}

func (r *FloodMongo) SetParam(ctx context.Context) error {
	n := ctx.Value("n").(int64)
	k := ctx.Value("k").(int64)
	floodTime := time.Now().Add(-time.Duration(k) * time.Second).Unix()
	filter := bson.M{"title": "param"}

	var existingDoc bson.M
	err := r.db.Collection(floodTable).FindOne(ctx, filter).Decode(&existingDoc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			param := bson.M{
				"title": "param",
				"n":     n,
				"k":     floodTime,
			}
			_, err := r.db.Collection(floodTable).InsertOne(ctx, param)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	update := bson.M{"$set": bson.M{"n": n, "k": k}}
	if _, err = r.db.Collection(floodTable).UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (r *FloodMongo) GetParam() (int64, int64, error) {
	filter := bson.M{"title": "param"}

	var result bson.M
	err := r.db.Collection(floodTable).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 1, 5, nil
		}
		return 1, 5, err
	}

	n := result["n"].(int64)
	k := result["k"].(int64)

	return n, k, nil
}
