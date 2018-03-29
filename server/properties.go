package server

import (
	"net/http"

	"github.com/effortless-technologies/elt-properties/models"

	"github.com/labstack/echo"
	"encoding/json"
)

func CreateProperty(c echo.Context) error {

	p := models.NewProperty()
	if err := c.Bind(p); err != nil {
		return err
	}
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

func UpdateProperty(c echo.Context) error {

	id := c.Param("id")
	property, err := models.FindPropertyById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var payload map[string]interface{}
	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	propertyJson, err := json.Marshal(property)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var propertyMap map[string]interface{}
	err = json.Unmarshal(propertyJson, &propertyMap)
	for k, v := range payload {
		for kk := range propertyMap {
			if k == kk {
				propertyMap[kk] = v
			}
		}
	}

	propertyJson, err = json.Marshal(propertyMap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	p := new(models.Property)
	err = json.Unmarshal(propertyJson, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	p.Update()

	return c.JSON(http.StatusOK, p)
}
