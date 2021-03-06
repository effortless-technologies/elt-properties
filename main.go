package main

import (
	"flag"

	"github.com/effortless-technologies/elt-properties/models"
	"github.com/effortless-technologies/elt-properties/server"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

func main() {

	flag.Parse()

	models.MongoAddr = mongoAddr

	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/properties", server.CreateProperty)
	r.GET("/properties", server.GetProperties)
	r.PUT("/properties/:id", server.UpdateProperty)
	r.DELETE("/properties/:id", server.DeleteProperty)
	r.GET("/properties/ingest", server.IngestProperties)

	e.Logger.Fatal(e.Start(":7001"))
}

