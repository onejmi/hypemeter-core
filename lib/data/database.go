package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)

//DatabaseName - Name of main database which should be accessed by the application
const DatabaseName = "core"
const DatabasePassword = "TRUk9i9pFoy4I0jE"

//DatabaseClient - mongodb client which can be used to interface with the database
var DatabaseClient *mongo.Client

//Connect - connects to database and logs any errors that may have occurred while doing so
func Connect() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://hypeadmin:" + DatabasePassword + "@main-7dq3n.mongodb.net")

	// Connect to MongoDB
	potentialClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = potentialClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DatabaseClient = potentialClient

	fmt.Println("Connected to Database!")
}

func Exists(collectionName string, filters interface{}) bool {
	err := DatabaseClient.Database(DatabaseName).Collection(collectionName).
		FindOne(context.TODO(), filters).Decode(&struct{}{})

	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return false
		} else {
			print(err)
		}
	}
	return true
}

func GetOne(collectionName string, filters interface{}, target interface{}) error {
	return DatabaseClient.Database(DatabaseName).Collection(collectionName).
		FindOne(context.TODO(), filters).Decode(target)
}

func List(collectionName string, filters bson.D) *mongo.Cursor {
	cursor, _ := DatabaseClient.Database(DatabaseName).Collection(collectionName).
		Find(context.TODO(), filters)
	return cursor
}

func Update(collectionName string, filter bson.D, update bson.M, upsert bool) (*mongo.UpdateResult, error) {
	return DatabaseClient.Database(DatabaseName).Collection(collectionName).UpdateOne(context.TODO(), filter, update,
		&options.UpdateOptions{Upsert: &upsert})
}

func Delete(collectionName string, filters bson.D) (*mongo.DeleteResult, error) {
	return DatabaseClient.Database(DatabaseName).Collection(collectionName).
		DeleteOne(context.TODO(), filters)
}

func Insert(collectionName string, entry interface{}) {
	transaction, err := DatabaseClient.Database(DatabaseName).Collection(collectionName).
		InsertOne(context.TODO(), entry)
	if err != nil {
		print(err)
	}

	fmt.Printf("Inserted %v in %s @ %d", transaction.InsertedID, collectionName, time.Now().Unix())
}
