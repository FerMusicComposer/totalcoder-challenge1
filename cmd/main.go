package main

import (
	"fmt"

	"github.com/FerMusicComposer/totalcoder-challenge1/db"
)

func main() {
	conn, err := db.NewMongoConnection(db.MONGOURI, db.DBNAME)
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to mongodb")
	fmt.Println(conn)
}
