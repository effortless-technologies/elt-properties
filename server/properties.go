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

func DeleteProperty(c echo.Context) error {

	id := c.Param("id")

	err := models.DeleteProperty(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func GetProperties(c echo.Context) error {

	p, err := models.GetProperties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, p)
}

