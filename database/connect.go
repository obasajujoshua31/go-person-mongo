package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDatabase() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Unable to connect to MongoDB Database")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Unable to connect to MongoDB Database")
	}

	return client.Database("go-mongo")
}
