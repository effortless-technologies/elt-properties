package server

import (
	"encoding/json"
	"net/http"

	"github.com/effortless-technologies/elt-properties/models"

	"github.com/labstack/echo"
)

func CreateProperty(c echo.Context) error {

	p := models.NewProperty()
	if err := c.Bind(p); err != nil {
		return err
	}
	err := p.CreateProperty()
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

func IngestProperties(c echo.Context) error {

	lodgixProperties, err := models.RetrieveLodgixProperties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	properties, err := models.GetProperties()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var notFoundLodgixProperties []*interface{}
	for i := range lodgixProperties {
		lp := *lodgixProperties[i]
		for k, v := range lp.(map[string]interface{}) {
			if k == "logdix_id" {
				lpId := int(v.(float64))
				found := false
				for i := range properties {
					if lpId == properties[i].LodgixId {
						found = true
					}
				}
				if found != true {
					notFoundLodgixProperties = append(
						notFoundLodgixProperties, &lp)
				}
			}
		}
	}

	var ingestedProperties []*models.Property
	for i := range notFoundLodgixProperties {
		p := models.NewProperty()

		lp := *notFoundLodgixProperties[i]
		for k, v := range lp.(map[string]interface{}) {
			if k == "logdix_id" {
				p.LodgixId = int(v.(float64))
			} else if k == "address" {
				 address := new(models.Address)
				for k, v := range v.(map[string]interface{}) {
					if k == "street_1" {
						address.StreetAddress = v.(string)
					} else if k == "street_2" {

					} else if k == "city" {
						address.City = v.(string)
					} else if k == "state" {
						address.State = v.(string)
					} else if k == "country" {

					} else if k == "zip_code" {
						address.ZipCode = v.(string)
					}
				}
				p.Address = address
			}
		}

		ingestedProperties = append(ingestedProperties, p)

		err = p.CreateProperty()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		ingestedProperties = append(ingestedProperties, p)
	}

	return c.JSON(http.StatusOK, ingestedProperties)
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

	p.UpdateProperty()

	return c.JSON(http.StatusOK, p)
}
