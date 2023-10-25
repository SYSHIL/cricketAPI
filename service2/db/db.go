package db

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a global mongo.Client variable
var globalMongoClient *mongo.Client
var mongoClientOnce sync.Once

// Initialize and return the global mongo.Client
func GetGlobalMongoClient() *mongo.Client {
	// Use a sync.Once to ensure that initialization only happens once
	mongoClientOnce.Do(func() {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI("mongodb+srv://shivaniniranjan30:Shivani30!@cluster0.0q2rc6i.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}

		// Send a ping to confirm a successful connection
		if err := client.Database("Cricket").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
			panic(err)
		}

		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

		globalMongoClient = client
	})

	return globalMongoClient
}
