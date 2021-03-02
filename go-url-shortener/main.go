package main

import (
	log "log"

	db "github.com/berksafran/go-url-shortener/db"
	h "github.com/berksafran/go-url-shortener/handlers"
	helpers "github.com/berksafran/go-url-shortener/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// Initial load .env config file.
	err := helpers.InitializeEnvVars()
	if err != nil {
		log.Fatal("[ERROR]: Error loading .env file")
	}
}

func main() {
	// Create new Echo instance.
	e := echo.New()
	e.HideBanner = true

	// Logrus
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[CLIENT] method=${method}, uri=${uri}, status=${status}\n",
	}))

	// DB Connection
	db.ConnectDB()
	client := db.DBClient.Client
	ctx := db.DBClient.Ctx

	log.Printf("[SERVER] DB Connection was established.\n\n")

	// Do not forget to close connection of DB.
	defer client.Disconnect(ctx)

	// Tanımladığımız Custom Middleware'ı kullanalım.
	// e.Use(ServerHeader)

	// Route handlers.
	e.GET("/", h.MainHandler)
	e.POST("/add", h.AddURLHandler)
	e.GET("/:path", h.RedirectURLHandler)

	// Start server.
	e.Start(":1234")
}

// Custom Middleware Oluşturma
// Response Header'ına Server parametresi tanımlıyoruz.

// ServerHeader is ...
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Berk/v1.0")
		return next(c)
	}
}
