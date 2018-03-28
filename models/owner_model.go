package models

import (

)

type Owner struct {
	FirstName					string			`json:"first_name"`
	LastName 					string			`json:"last_name"`
	//ClientType.. maybe "time_when_enrolled"
	ReferredBy					string			`json:"referred_by"`
	Phone						string			`json:"phone"`
	Email 						string			`json:"email"`
	PreferredContactMethod		string			`json:"preferred_contact_method"` 	// Preferred Form of Contact
	Availability				string			`json:"availability"`				// Best Times to Contact
	ShorttermLicenseNumber		string			`json:"shortterm_license_number"`
	Location 					string			`json:"location"`					// Are you local?
}

func NewOwner(
	firstName string, lastName string, phone string, email string) *Owner {

		o := new(Owner)
		o.FirstName = firstName
		o.LastName = lastName
		o.Phone = phone
		o.Email = email

		return o
}
