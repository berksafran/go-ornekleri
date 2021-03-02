package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/berksafran/go-url-shortener/db"
	dbhelpers "github.com/berksafran/go-url-shortener/dbhelpers"
	"github.com/berksafran/go-url-shortener/helpers"
	model "github.com/berksafran/go-url-shortener/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// URL is model of request body.
type URL struct {
	URL          string    `json:"url"`
	Path         string    `json:"path"`
	Description  string    `json:"description"`
	CreatedDate  time.Time `json:"createdDate"`
	VisitedCount int       `json:"visitedCount"`
}

// Validate is validation of request body.
func (u URL) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.URL, validation.Required, is.URL),
		validation.Field(&u.Description, validation.Required, validation.Length(5, 50)),
	)
}

// AddURLHandler creates new short URL.
func AddURLHandler(c echo.Context) error {
	urlStruct := &URL{}

	err := c.Bind(&urlStruct)
	if err != nil {
		log.Println("[ERROR]:", err.Error())
		return c.JSON(http.StatusBadRequest,
			model.ResponseError{
				Status:  false,
				Message: "Bad Request",
			})
	}

	// Validation
	err = urlStruct.Validate()
	// to get error with JSON formatted, uncomment following line and change Message: string(errJSON)
	// errJSON, _ := json.Marshal(err)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest,
			model.ResponseError{
				Status:  false,
				Message: string(err.Error()),
			})
	}

	// Default value of selected path's visited counter is 0
	urlStruct.VisitedCount = 0

	// Generate random path
	path := helpers.GeneratePath()

	isExist, err := dbhelpers.IsPathExist(path)
	if isExist == true {
		log.Println("[SERVER] New path is generating..")
		path = helpers.GeneratePath()
	}

	// Insert a new Path Object in the database
	result, err := db.InsertOne(
		bson.D{
			{Key: "url", Value: urlStruct.URL},
			{Key: "path", Value: path},
			{Key: "description", Value: urlStruct.Description},
			{Key: "createdDate", Value: time.Now().Local()},
			{Key: "visitedCount", Value: urlStruct.VisitedCount},
		})
	if err != nil {
		log.Println("ERROR:", err)
		return c.JSON(500,
			model.ResponseError{
				Status:  false,
				Message: "There is an error.",
			})
	}

	return c.JSON(http.StatusAccepted,
		model.ResponseSuccess{
			ID:      result.InsertedID,
			Path:    path,
			Status:  true,
			Message: "Success! URL has been created.",
		})
}
