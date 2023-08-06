package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "poll"
	collectionName = "poll"
)

var collection *mongo.Collection

func Collection() *mongo.Collection {
	return collection
}

func Init() {
	connectionString := "mongodb://localhost:27017/"
	println("hello world")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkError(err)

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance is ready")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
