package handlers

import "github.com/labstack/echo/v4"

// MainHandler is handler of "/"
func MainHandler(c echo.Context) error {
	return c.String(200, "Hello World. \nThis is Shorty.")
}
