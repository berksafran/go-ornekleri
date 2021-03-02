package db

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOne creates new data on selected collection.
func InsertOne(object interface{}) (*mongo.InsertOneResult, error) {
	GetConfig()
	result, err := urlCollection.InsertOne(ctx, object)
	if err != nil {
		log.Println(err)
	}

	return result, err
}

// Find finds param object in the database, then returns result if object exists.
func Find(obj interface{}) ([]primitive.M, error) {
	GetConfig()

	filterCursor, err := urlCollection.Find(ctx, obj)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []bson.M
	if err = filterCursor.All(ctx, &result); err != nil {
		log.Println(err)
		return nil, err
	}

	defer filterCursor.Close(ctx)
	return result, nil
}

// UpdateOne updates value selected key on selected item.
func UpdateOne(objID primitive.ObjectID, updatedObj primitive.D) (*mongo.UpdateResult, error) {
	GetConfig()

	result, err := urlCollection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		updatedObj,
	)
	if err != nil {
		err = errors.New("[ERROR]: Path does not exist in the database")
		return nil, err
	}

	return result, nil
}
