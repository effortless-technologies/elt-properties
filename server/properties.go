package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-properties/models"

	"github.com/labstack/echo"
)

func CreateProperty(c echo.Context) error {

	p := models.NewProperty()
	err := p.Create()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, p)
}

func GetProperties(c echo.Context) error {

	p := new(models.Property)

	return c.JSON(http.StatusOK, p)
}

