package db

import (
	"context"

	helpers "github.com/berksafran/go-url-shortener/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx context.Context
var database *mongo.Database
var dbName string
var collectionName string
var urlCollection *mongo.Collection

// GetConfig loads constant and variables about DB configuration.
func GetConfig() {
	// Default Database name, collection name.
	dbName = helpers.GetEnv("DB_NAME")
	collectionName = helpers.GetEnv("COLLECTION_NAME")
	client = DBClient.Client
	ctx = DBClient.Ctx
	database = client.Database(dbName)
	urlCollection = database.Collection(collectionName)
}
