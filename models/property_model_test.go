package models

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

var propertyId *bson.ObjectId

func TestProperties_CreateProperty(t *testing.T) {

	Convey("If a properties database exists", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When creating a property", func() {
			property := NewProperty()
			err := property.CreateProperty()
			So(err, ShouldBeNil)
			propertyId = property.Id
			So(propertyId, ShouldNotBeNil)

			Convey("A property should have been created", func() {
				properties, err := GetProperties()
				So(err, ShouldBeNil)
				So(properties, ShouldNotBeNil)
				found := false

				for i := range(properties) {
					if properties[i].Id.Hex() == propertyId.Hex() {
						found = true
					}
				}

				So(found, ShouldEqual, true)
			})
		})
	})
}

func TestPropertiesModel_GetProperties(t *testing.T) {

	Convey(
		"If a properties database exists with a property", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When getting properties", func() {
			properties, err := GetProperties()
			So(err, ShouldBeNil)

			Convey("A list of properties to be returned", func() {
				So(properties, ShouldNotBeNil)
				So(len(properties), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestPropertiesModel_FindPropertyById(t *testing.T) {

	Convey("If a properties database exists", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When finding a property by id", func() {
			property, err := FindPropertyById(propertyId.Hex())
			So(err, ShouldBeNil)
			So(property, ShouldNotBeNil)

			Convey(
				"The property found should have the same id", func() {
				So(property.Id, ShouldResemble, propertyId)
			})
		})
	})
}

func TestPropertiesModel_Update(t *testing.T) {

	Convey("If a properties database exists", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When updating a property to have 5 floors", func() {
			property, err := FindPropertyById(propertyId.Hex())
			property.Floors = 5
			So(err, ShouldBeNil)
			So(property, ShouldNotBeNil)

			err = property.UpdateProperty()
			So(err, ShouldBeNil)

			Convey("The updated property should have 5 floors", func() {
				property, err := FindPropertyById(propertyId.Hex())
				So(err, ShouldBeNil)
				So(property, ShouldNotBeNil)
				So(property.Floors, ShouldEqual, 5)
			})
		})
	})
}

func TestPropertiesModel_DeleteProperty(t *testing.T) {

	Convey("If a properties database exists", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When deleting an existing property", func() {
			err := DeleteProperty(propertyId.Hex())
			So(err, ShouldBeNil)

			Convey("A property should have been deleted", func() {
				_, err := FindPropertyById(propertyId.Hex())
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestPropertiesModel_RetrieveLodgixProperties(t *testing.T) {

	Convey("If properties exist in Lodgix", t, func() {

		Convey("When ingesting properties from lodgix", func() {
			properties, err := RetrieveLodgixProperties()
			So(err, ShouldBeNil)

			Convey("A list of properties should returned", func() {
				So(properties, ShouldNotBeNil)
				So(len(properties), ShouldBeGreaterThan, 0)
			})
		})
	})
}
