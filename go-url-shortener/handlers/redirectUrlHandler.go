package handlers

import (
	"net/http"

	"github.com/berksafran/go-url-shortener/dbhelpers"
	model "github.com/berksafran/go-url-shortener/models"
	"github.com/labstack/echo/v4"
)

// RedirectURLHandler forwards to selected saved URL in DB.
func RedirectURLHandler(c echo.Context) error {
	path := c.Param("path")

	// Getting Result as BSON Object from DB.
	result, err := dbhelpers.GetPathObject(path)
	if err != nil || result == nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound,
			model.ResponseError{
				Status:  false,
				Message: err.Error(),
			})
	}

	// If path exists in the DB, increase visitedCount of related path
	err = dbhelpers.IncreaseVisitedCounter(result)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotModified,
			model.ResponseError{
				Status:  false,
				Message: "Error! There is something wrong. Try again later.",
			})
	}
	// Then, redirect client to URL.
	url := result["url"].(string)
	return c.Redirect(http.StatusPermanentRedirect, url)
}
