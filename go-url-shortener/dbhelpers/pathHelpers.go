package dbhelpers

import (
	"errors"
	"fmt"
	"log"
	"time"

	db "github.com/berksafran/go-url-shortener/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPathObject gets selected path's object from DB.
func GetPathObject(path string) (primitive.M, error) {
	result, err := db.Find(bson.M{"path": path})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if result == nil {
		err = errors.New("Error! Path does not exist in the database")
		return nil, err
	}

	return result[0], nil
}

// IsPathExist checks whether selected path exists or not
func IsPathExist(path string) (bool, error) {
	result, err := db.Find(bson.M{"path": path})
	if err != nil {
		return false, err
	}
	if len(result) > 0 {
		return true, nil
	}
	return false, nil
}

// IncreaseVisitedCounter increases visited number of selected path.
func IncreaseVisitedCounter(pathObj primitive.M) error {
	id := pathObj["_id"].(primitive.ObjectID)
	updatedObj := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "visitedCount", Value: 1},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "lastVisited", Value: time.Now().Local().String()},
		}},
	}

	_, err := db.UpdateOne(id, updatedObj)
	if err != nil {
		err = errors.New("[ERROR]: There is something error")
		return err
	}

	fmt.Println("[SERVER]:", pathObj["url"], "visited.")
	return nil
}
