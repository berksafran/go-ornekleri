package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func login(c echo.Context) error {

	// Form'dan username ve password'u alıyoruz.
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(username)

	// Eğer username veya password yanlışsa Unauthorized hatası gönderiyoruz.
	if username != "berk" || password != "safran" {
		return echo.ErrUnauthorized
	}

	// Token yaratıyoruz.
	token := jwt.New(jwt.SigningMethodHS256)

	// Claims (token içerisine sakladığımız object) set ediyoruz.
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretooo"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	// "user" c değişkeni içerisinde geliyor statik olarak.
	user := c.Get("user").(*jwt.Token)
	// Claims token içerisine saklanan object anlamına geliyor.
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println("**********")
	fmt.Println("Claims:", claims)
	fmt.Println("Claims interface:", claims["name"])
	fmt.Println("Claims interface:", claims["name"].(string))
	name := claims["name"].(string) // interface{} boş kümesini string'e çeviriyoruz.
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	e.HideBanner = true

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secretooo")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
