package models

import (

)

type Owner struct {
	FirstName					string			`json:"first_name" bson:"first_name"`
	LastName 					string			`json:"last_name" bson:"last_name"`
	//ClientType.. maybe "time_when_enrolled"
	ReferredBy					string			`json:"referred_by" bson:"referred_by"`
	PhoneNumber					string			`json:"phone_number" bson:"phone_number"`
	Email 						string			`json:"email" bson:"email"`
	PreferredContactMethod		string			`json:"preferred_contact_method" bson:"preferred_contact_method"` 	// Preferred Form of Contact
	Availability				string			`json:"availability" bson:"availability"`	// Best Times to Contact
	ShorttermLicenseNumber		string			`json:"shortterm_license_number" bson:"shortterm_license_number"`
	Location 					string			`json:"location" bson:"location"`	// Are you local?
}

func NewOwner(
	firstName string,
	lastName string,
	phoneNumber string,
	email string) *Owner {

		o := new(Owner)
		o.FirstName = firstName
		o.LastName = lastName
		o.PhoneNumber = phoneNumber
		o.Email = email

		return o
}
