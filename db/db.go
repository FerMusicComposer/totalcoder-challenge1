package db

import "context"

const (
	MONGOURI = "mongodb://localhost:27017"
	DBNAME   = "TotalCoderChallenges"
)

type Dropper interface {
	Drop(context.Context) error
}
