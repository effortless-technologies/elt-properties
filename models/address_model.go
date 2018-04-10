package models

import (

)

type Address struct {
	StreetAddress1 				string			`json:"street_address_1" bson:"street_address_1"`
	StreetAddress2 				string			`json:"street_address_2" bson:"street_address_2"`
	City 						string			`json:"city" bson:"city"`
	State 						string			`json:"state" bson:"state"`
	ZipCode 					string			`json:"zip_code" bson:"zip_code"`
}

func NewAddress(
	streetAddress1 string,
	streetAddress2 string,
	city string,
	state string,
	zipCode string) *Address {

		a := new(Address)
		a.StreetAddress1 = streetAddress1
		a.StreetAddress2 = streetAddress2
		a.City = city
		a.State = state
		a.ZipCode = zipCode

		return a
}
