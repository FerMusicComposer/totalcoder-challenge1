package db

import (
	"context"
	"fmt"

	"github.com/FerMusicComposer/totalcoder-challenge1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const recordColl = "records"

type RecordStore interface {
	GetRecordsByFilter(context.Context, bson.M) ([]models.Record, error)
}

type MongoRecordStore struct {
	connection *MongoConnection
	coll       *mongo.Collection
}

func NewMongoRecordStore(conn *MongoConnection) *MongoRecordStore {
	return &MongoRecordStore{
		connection: conn,
		coll:       conn.Database.Collection(recordColl),
	}
}

func (s *MongoRecordStore) GetRecordsByFilter(ctx context.Context, filter bson.M) ([]models.Record, error) {
	records := []models.Record{}
	fmt.Println("received filter:", filter)
	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// defer cursor.Close(ctx)

	if err := cursor.All(ctx, &records); err != nil {
		return nil, err
	}
	fmt.Println("records returned:", records)
	return records, nil
}
