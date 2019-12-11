package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database string
	password string
)

func init() {
	database = os.Getenv("AZURE_DATABASE")
	password = os.Getenv("AZURE_DATABASE_PASSWORD")

	if database == "" || password == "" {
		fmt.Printf("AZURE_DATABASE environment variable must be the name of the Cosmos DB database and AZURE_DATABASE_PASSWORD must be the primary password for that database.")
		os.Exit(1)
	}
}

func main() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s.documents.azure.com:10255/%s?ssl=true", database, password, database, database)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Can't create mongodb client, go error %v\n", err)
	}

	// Set a 5s timeout and ensure that it is called
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Establish a connection to Cosmos DB
	err = client.Connect(context)
	if err != nil {
		log.Fatalf("Can't connect to mongodb server %s, go error %v\n", uri, err)
	}

	// Retrieve a reference to the package collection in Cosmos DB
	collection := client.Database(database).Collection("package")

	// Write a single document into Cosmos DB
	var insertResult *mongo.InsertOneResult
	insertResult, err = collection.InsertOne(context,
		bson.M{
			"FullName":      "react",
			"Description":   "A framework for building native apps with React.",
			"ForksCount":    11392,
			"StarsCount":    48794,
			"LastUpdatedBy": "shergin",
		},
	)
	if err != nil {
		log.Fatalf("Failed to insert record: %v\n", err)
	}

	id := insertResult.InsertedID
	fmt.Printf("Inserted record id: %d\n", id)

	// Update document
	var updateResult *mongo.UpdateResult
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"fullname": "react-native",
		},
	}
	updateResult, err = collection.UpdateOne(context, filter, update)
	if err != nil {
		log.Fatalf("Error updating record %v\n", err)
	}
	fmt.Printf("Updated this many records: %d\n", updateResult.ModifiedCount)

	// Delete document by id
	var deleteResult *mongo.DeleteResult
	deleteResult, err = collection.DeleteOne(context, filter)
	if err != nil {
		log.Fatalf("Error deleting record: %v\n", err)
	}
	fmt.Printf("Deleted this many records: %d\n", deleteResult.DeletedCount)
}
