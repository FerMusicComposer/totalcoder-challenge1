package db

import (
	"context"
	"fmt"
	"time"

	"github.com/FerMusicComposer/totalcoder-challenge1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const recordColl = "records"

type RecordStore interface {
	GetRecordsByFilter(context.Context, time.Time, time.Time, int, int) ([]models.Record, error)
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

func (s *MongoRecordStore) GetRecordsByFilter(ctx context.Context, startDate, endDate time.Time, minCount, maxCount int) ([]models.Record, error) {
	records := []models.Record{}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"createdAt": bson.M{
				"$gte": startDate.Format(time.RFC3339),
				"$lte": endDate.Format(time.RFC3339),
			},
		},
		}},
		{{Key: "$addFields", Value: bson.M{"totalCount": bson.M{"$sum": "$count"}}}},
		{{Key: "$match", Value: bson.M{
			"totalCount": bson.M{
				"$gte": minCount,
				"$lte": maxCount},
		},
		}},
	}

	fmt.Println("pipeline:", pipeline)
	cursor, err := s.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &records); err != nil {
		return nil, err
	}
	fmt.Println("records returned:", records)
	return records, nil
}
