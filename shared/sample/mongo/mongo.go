package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get a collection
	collection := client.Database("testdb").Collection("testcollection")

	// Insert a document
	doc := bson.D{
		{Key: "name", Value: "John Doe"},
		{Key: "age", Value: 30},
	}
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted a single document: %v\n", insertResult.InsertedID)

	// Find a document
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{
		{Key: "name", Value: "John Doe"},
	}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %v\n", result)

	// Disconnect from MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB.")
}
