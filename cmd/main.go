package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/FerMusicComposer/totalcoder-challenge1/db"
	"github.com/FerMusicComposer/totalcoder-challenge1/handlers"
)

func main() {
	addr := ":5000"
	srv := &http.Server{
		Addr: addr,
	}

	// Initialize DB connection
	conn, err := db.NewMongoConnection(db.MONGOURI, db.DBNAME)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Could not close connection to MongoDB: %s", err)
		}

		fmt.Println("Mongo connection closed")
	}()

	fmt.Println("Connected to MongoDB")

	// Handlers initialization
	recordStore := db.NewMongoRecordStore(conn)
	recordHandler := handlers.NewRecordHandler(recordStore)

	// Routes
	mux := http.NewServeMux()

	mux.HandleFunc("/records", recordHandler.HandleGetRecords)

	srv.Handler = mux

	go func() {
		fmt.Println("Starting server...")
		fmt.Printf("Server started on port: %s\n", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s", err)
		}
	}()

	// Channel for the 'quit' signal to manage graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	fmt.Println("Server exiting")
}
