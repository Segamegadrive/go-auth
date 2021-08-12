package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {

	//Loading .env file from the directory
	err := godotenv.Load(".env")

	//Checking if .env file ever exists
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Grabbing the specific key from the .env file
	MongoDB := os.Getenv("MONGODB_URL")

	//As a required key from the .env file for db connection is ready/available now, it is the time for defining a new client and tagging the key with it.
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	//https://pkg.go.dev/context
	//Create context whilst sending request to the servers.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("user").Collection(collectionName)

	return collection
}
