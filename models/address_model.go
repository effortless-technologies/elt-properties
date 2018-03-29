package models

import (

)

type Address struct {
	StreetAddress 				string			`json:"street_address" bson:"street_address"`
	City 						string			`json:"city" bson:"city"`
	State 						string			`json:"state" bson:"state"`
	ZipCode 					string			`json:"zip_code" bson:"zip_code"`
}

func NewAddress(
	streetAddress string, city string, state string, zipCode string) *Address {

		a := new(Address)
		a.StreetAddress = streetAddress
		a.City = city
		a.State = state
		a.ZipCode = zipCode

		return a
}
