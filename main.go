package main

import (
	"flag"
	"net/http"

	"github.com/effortless-technologies/elt-properties/models"
	"github.com/effortless-technologies/elt-properties/server"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

func accessible(c echo.Context) error {

	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {

	flag.Parse()

	models.MongoAddr = mongoAddr

	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", accessible)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
	r.POST("/properties", server.CreateProperty)
	r.GET("/properties", server.GetProperties)
	r.PUT("/properties/:id", server.UpdateProperty)
	r.DELETE("/properties/:id", server.DeleteProperty)

	e.Logger.Fatal(e.Start(":7001"))
}

