package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	helpers "github.com/berksafran/go-url-shortener/helpers"
)

// Client contains mongo.Client
type Client struct {
	Client *mongo.Client
	Ctx    context.Context
}

// DBClient is main struct of configuration DB.
var DBClient = Client{}

// ConnectDB is handling connection
func ConnectDB() {
	dbURI := helpers.GetEnv("DB_URI")

	// Set client options
	clientOptions := options.Client().ApplyURI(dbURI)

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

	DBClient = Client{Client: client, Ctx: context.TODO()}
}
