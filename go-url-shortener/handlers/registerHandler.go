package handlers

import (
	"log"
	"net/http"

	model "github.com/berksafran/go-url-shortener/models"
	"github.com/labstack/echo/v4"
)

// RegisterHandler is handler of user registration for creating custom paths.
func RegisterHandler(c echo.Context) error {
	userStruct := &model.User{}

	err := c.Bind(&userStruct)
	if err != nil {
		log.Println("[ERROR]:", err.Error())
		return c.JSON(http.StatusBadRequest,
			model.ResponseError{
				Status:  false,
				Message: "Bad Request",
			})
	}

	return c.String(200, "Hello World. \nThis is Shorty.")
}
