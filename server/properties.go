package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-properties/models"

	"github.com/labstack/echo"
)

func Properties(c echo.Context) error {

	p := new(models.Property)

	return c.JSON(http.StatusOK, p)
}

