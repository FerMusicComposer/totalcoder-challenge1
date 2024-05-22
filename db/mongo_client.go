package db

import (
	"context"

	"github.com/FerMusicComposer/totalcoder-challenge1/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGOURI = "mongodb://root:example@localhost:27017"
	DBNAME   = "records-db"
)

type MongoConnection struct {
	Client       *mongo.Client
	Database     *mongo.Database
	DatabaseName string
}

func NewMongoConnection(uri, dbName string) (*MongoConnection, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, utils.ConnectionError.From("Mongo Client Builder => couldn't initialize mongo client", err).Err
	}

	// Verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, utils.ConnectionError.From("Mongo Client Builder => mongo instance unreachable", err).Err
	}

	db := client.Database(dbName)
	return &MongoConnection{
		Client:       client,
		Database:     db,
		DatabaseName: dbName,
	}, nil
}
