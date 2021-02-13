package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Öncelikle Echo nesnemizi oluşturuyoruz.
	e := echo.New()
	// Terminalde sunucuyu kaldırdığımızda çıkacak Echo banner'ını kapatıyoruz.
	e.HideBanner = true

	// Tüm endpointler için çalışan Logger tanımlıyoruz.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Echo ile GET,POST,PATCH vb. methodlarını belirleyebiliyoruz.
	e.GET("/", mainHandler)

	e.GET("/app", appHandler)               // Query Param Örneği
	e.GET("/users/:id/:name", usersHandler) // Path Param Örneği
	e.GET("/data/:data", dataHandler)       // Path Param Örneği

	e.POST("addUser", addUserHandler) // Request'te JSON kabul etme

	// Grup oluşturup, grup özelinde middleware tanımlıyoruz.
	a := e.Group("/admin", middleware.Logger())
	a.GET("/dashboard", adminDashboardHandler)

	// Sunucumuzu ayağa kaldırıyoruz.
	e.Start(":8000")
}

func mainHandler(c echo.Context) error {
	/*
		c.String ile response olarak String dönebiliriz.
		200 yerine http.StatusOK 'da yazabiliriz.
	*/
	return c.String(200, "Hello World.")
}

//	QUERY PARAM Örneği - "/app?name=Veteran&version=23.0"
func appHandler(c echo.Context) error {
	// Query Param'lara c.queryParam ile ulaşıyoruz.
	name := c.QueryParam("name")
	version := c.QueryParam("version")

	// veya
	queryParams := c.QueryParams()
	fmt.Println("QP - name:", queryParams.Get("name"), queryParams["name"])
	fmt.Println("QP - version:", queryParams.Get("version"), queryParams["version"])

	if name == "" && version == "" {
		message := "There is no query param on request."
		fmt.Println(message)
		return c.String(404, message)
	}

	fmt.Printf("QueryParams: name: %s, version: %v\n", name, version)
	return c.String(200, "Done")
}

// PATH PARAM Örneği - "/users/:id/:name"
func usersHandler(c echo.Context) error {
	id := c.Param("id")
	name := c.Param("name")
	return c.String(200, "ID: "+id+" Name: "+name)
}

// PATH PARAM + QUERY PARAM Örneği - '/data/:json?username=""&password="" '
func dataHandler(c echo.Context) error {
	// Path Param
	dataType := c.Param("data")

	// Query Params
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if dataType == "string" {
		return c.String(200, "Username:"+username+" Password:"+password)
	}

	// JSON yaratmak için map oluşturuyoruz. key ve value string türlerinde olacak.
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"password": password,
		})
	}

	return c.String(404, "'/json' or 'string' paths are accepted.")
}

// REQUEST'TE JSON KABUL ETME - "/addUser"
// User is model of request body.
/*
	Dikkat! User struct'ının fieldları erişebilir olmalı.
	Bu yüzden, baş harflerini büyük yazıyoruz.
*/
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func addUserHandler(c echo.Context) error {
	log.Println("POST isteği geldi!")
	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		c.String(404, "Request body must be not empty.")
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("ERROR:", err)

		// return err ile user'a JSON olarak hata döneceğiz.
		// (Burada Internal Server Error dönecektir.)
		return err
	}

	// Gelen veriyi terminale bastıralım ve kullanıcıya success mesajı dönelim.
	fmt.Println("Incoming Request:", user)
	return c.String(200, "Success!")
}

// Group Tanımlama ve Middleware Kullanımı - "/admin/"
func adminDashboardHandler(c echo.Context) error {
	return c.String(200, "It is /admin/dashboard")
}
