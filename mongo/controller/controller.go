package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://learncodeonline:hitesh@cluster0.humov.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "movies"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	fmt.Println("Controller in go")
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection success")

	collection := client.Database(dbName).Collection(colName)

}
