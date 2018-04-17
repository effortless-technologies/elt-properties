package server

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/effortless-technologies/elt-properties/models"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

var propertyPayload = `
{
	"address": {
		"street_address": "123 4th St",
		"city": "Denver",
		"state": "CO",
		"zip_code": "80210"
	},
	"owner": {
		"first_name": "Jerry",
		"last_name": "Rice",
		"phone": "720-777-3432",
		"email": "jerryrice@gmail.com"
	}
}
`

var propertyId *bson.ObjectId

func TestProperties_CreateProperty(t *testing.T) {
	Convey("If a properties database exists", t, func() {
		db := []string{*mongoAddr}
		models.MongoAddr = db
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(
			echo.POST, "/", strings.NewReader(propertyPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/properties")

		Convey("When calling the property POST/properties " +
			"handler", func() {
				err := CreateProperty(c)
				So(err, ShouldBeNil)

			Convey("Then a property should be returned with a status " +
				"code of 200", func() {
					So(rec.Code, ShouldEqual, 200)

					payload, _ := ioutil.ReadAll(rec.Body)
					var property *models.Property
					err = json.Unmarshal([]byte(payload), &property)
					So(err, ShouldBeNil)
					So(property.Owner.FirstName, ShouldEqual, "Jerry")
					So(property.Owner.LastName, ShouldEqual, "Rice")

					propertyId = property.Id
					So(propertyId, ShouldNotBeNil)
			})
		})
	})
}

func TestProperties_GetProperties(t *testing.T) {
	Convey("If a properties database exists", t, func() {
		db := []string{*mongoAddr}
		models.MongoAddr = db
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/properties")

		Convey("When calling the GET/properties handler", func() {
			err := GetProperties(c)
			So(err, ShouldBeNil)

			Convey("Then a json array of properties should be " +
				"returned with a status code of 200", func() {
					So(rec.Code, ShouldEqual, 200)

					payload, _ := ioutil.ReadAll(rec.Body)
					var properties []*models.Property
					err = json.Unmarshal(payload, &properties)
					So(err, ShouldBeNil)
					So(len(properties), ShouldBeGreaterThan, 0)
				})
		})
	})
}

func TestProperties_UpdateProperty(t *testing.T) {
	Convey("If a properties database exists", t, func() {
		db := []string{*mongoAddr}
		models.MongoAddr = db
		So(models.MongoAddr, ShouldNotBeNil)

		var reqPayload = `
		{
			"shampoo": true,
			"iron": true
		}
		`

		e := echo.New()
		req := httptest.NewRequest(
			echo.PUT, "/", strings.NewReader(reqPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/properties/:id")

		c.SetParamNames("id")
		c.SetParamValues(propertyId.Hex())

		Convey("When calling the UPDATE/properties handler", func() {
			err := UpdateProperty(c)
			So(err, ShouldBeNil)

			Convey("Then a property with the update values with " +
				"a status code of 200 should be returned", func() {
					So(rec.Code, ShouldEqual, 200)

					payload, _ := ioutil.ReadAll(rec.Body)
					var property *models.Property
					err = json.Unmarshal([]byte(payload), &property)
					So(err, ShouldBeNil)
					So(property.Owner.FirstName, ShouldEqual, "Jerry")
					So(property.Owner.LastName, ShouldEqual, "Rice")
					So(property.Shampoo, ShouldEqual, true)
					So(property.Iron, ShouldEqual, true)
			})
		})
	})
}

func TestProperties_DeleteProperty(t *testing.T) {
	Convey("If a properties database exists", t, func() {
		db := []string{*mongoAddr}
		models.MongoAddr = db
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/properties/:id")

		c.SetParamNames("id")
		c.SetParamValues(propertyId.Hex())

		Convey("When calling the DELETE/properties handler", func() {
			err := DeleteProperty(c)
			So(err, ShouldBeNil)

			Convey("Then the returned status code should be " +
				"204", func() {
					So(rec.Code, ShouldEqual, 204)
			})
		})
	})
}

func TestProperties_IngestProperties(t *testing.T) {
	Convey("If a properties database exists", t, func() {
		db := []string{*mongoAddr}
		models.MongoAddr = db
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/properties/ingest")

		Convey(
			"When calling the GET/properties/ingest handler", func() {
				err := IngestProperties(c)
				So(err, ShouldBeNil)

			Convey("Then a json array of properties should be " +
				"returned with a status code of 200", func() {
					So(rec.Code, ShouldEqual, 200)

					payload, _ := ioutil.ReadAll(rec.Body)
					var properties []*models.Property
					err = json.Unmarshal(payload, &properties)
					So(err, ShouldBeNil)
					So(len(properties), ShouldBeGreaterThan, 0)
			})
		})

		ps, err := models.GetProperties()
		So(err, ShouldBeNil)
		for _, p := range ps {
			err := models.DeleteProperty(p.Id.Hex())
			if err != nil {
				panic(err)
			}
		}
	})
}
