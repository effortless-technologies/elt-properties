package models

import (
	"time"
)

// Notes:
//	Many duplicates
// 	Sort out Amenities
//	Sort out cleaner checks
// 	Sort out what data should be altered and read by which people
// 	Need to classify types of data
//		+ Amenities
//		+ Extras
//

type Property struct {
	Owner 						*Owner			`json:"owner"`
	PutOnMarket					time.Time		`json:"date_on_market"`				// When is Property Open
	BlackoutDates				[]*Period		`json:"blackout_period"`			// Details about future dates to block
	IsShorterm					bool			`json:"is_shorterm"`				// Currently Short-term Rental
	FutureReservations 			bool			`json:"future_reservations"`
	FirstReservation 			time.Time		`json:"first_reservation"`
	NumFutureReservations 		int				`json:"num_future_reservations"` 	// How Many Future Reservations
	LockType					string			`json:"lock_type"`					// Type of lock on property
	DoorCodes 					[]*DoorCode		`json:"door_codes"`
	CommunityCode				string			`json:"community_code"`				// Community/Gate Code
	BuildingCode				string			`json:"building_code"`
	IntercomSystem				bool 			`json:"intercom_system"`
	HouseType					string			`json:"house_type"`
	Address						*Address 		`json:"address"`
	Warranties 					[]*Warranty		`json:"warranties"`
	KnownIssues					[]string		`json:"known_issues"`
	SnowRemoval					bool			`json:"snow_removal"`				// Does the Client want Snow Removal, Yard, ect ; Will your property need snow removal
	Landscaping					bool			`json:"landscaping"`				// Does the Client want Snow Removal, Yard, ect ; Will your property need lawn care
	MaxOccupancy				int				`json:"max_occupancy"`
	Bedrooms					[]*Bedroom		`json:"beds"`						// Number of bedsrooms
	Bathrooms					int				`json:"baths"`						// Number of Bathrooms
	Floors						int				`json:"floors"`						// Number of Floors
	LocationOfLinens			string			`json:"location_of_linens"`			// Location of Extra Pillows, Blankets, ect
	AirConditioning 			bool			`json:"air_conditioning"`
	Heating						bool			`json:"heating"`
	ThermostatLocation 			string			`json:"thermostat_location"` 		// Detailed Location of Thermostat
	SecuritySystem 				bool			`json:"security_system"`			// Home Security
	SecuritySystemCode			string 			`json:"security_system_code"`		// Hoe Security Code
	MainWaterLocation			string			`json:"main_water_location"`		// Detailed Location of Water Shut Off Valve
	WaterHeaterLocation 		string			`json:"water_heater_location"`		// Detailed Location of Water Heater
	SmokeDetectors 				bool			`json:"smoke_detectors"`
	CarbonMonoxideDetector		bool			`json:"carbon_monoxide_detector"`
	Tvs							[]*Tv 			`json:"tvs"`
	Wifi 						*Wifi 			`json:"wifi"`
	Stereo 						bool			`json:"stereo"`
	BreakersLocations			[]*string		`json:"breakers_locations"` 		// Detailed Location of Breakers
	FitnessCenter				bool			`json:"fitness_center"`				// Detailed Gym Locatoin
	FitnessCenterLocation 		string			`json:"fitness_center_location"`	// Detailed Gym Location
	CleanersCloset				*Closet			`json:"cleaners_closet"`
	OwnersCloset				*Closet			`json:"owners_closet"`
	RooftopDeckPatio			bool			`json:"rooftop_deck_patio"` 		// Rooftop/Deck/Patio
	Grill 						bool			`json:"grill"`						// BBQ/Grill
	GrillType					string			`json:"grill_type"`					// Type of Grill
	HandicapAccessible			bool			`json:"handicap_accessible"`
	TrackPickup					string 			`json:"track_pickup"`				// Track Pickup Day
	TrashLocation				string			`json:"trash_location"`				// Location ad Instructions for Trash
	TrashInstructions			string			`json:"trash_instructions"`			// Location ad Instructions for Trash
	RecyclingPickup 			string			`json:"recycling_pickup"`			// Recycling Pickup Day
	RecyclingLocation			string			`json:"recycling_location"`			// Location and Instructions for Trach & Recycling
	RecyclingInstruction 		string			`json:"recycling_instruction"`		// Location and Instructions for Trach & Recycling
	Elevator 					bool			`json:"elevator"`
	Parking						string			`json:"parking"`
	Vacuuum						bool			`json:"vacuuum"`
	HairDryer					bool			`json:"hair_dryer"`
	ExtraSheets					bool			`json:"extra_sheets"`				// 3 Set of Sheets for Each Bed
	Duvets						bool			`json:"duvets"`						// Duvets for all blankets
	Iron						bool			`json:"iron"`
	IroningBoard				bool			`json:"ironing_board"`
	Broom						bool			`json:"broom"`
	ToiletPlunger				bool			`json:"toilet_plunger"`
	BathMats					bool			`json:"bath_mats"`
	ToiletBowlBrush				bool			`json:"toilet_bowl_brush"`
	WaterproofMattressCovers	bool			`json:"waterproof_mattress_covers"`	// Waterproof Mattress Covers for all Beds
	ToiletPaper					bool			`json:"toilet_paper"`
	PaperTowels					bool			`json:"paper_towels"`
	HandSoap					bool			`json:"hand_soap"`					// Hand Soap for all Sinks
	DishwasherSoap				bool			`json:"dishwasher_soap"`
	Coffee						bool			`json:"coffee"`
	BathTrashCan				bool			`json:"bath_trash_can"`				// Is there a Trashcan in every Bathroom
	TrashBags					bool			`json:"trash_bags"`					// Backup Trash Bags
	DoorMats					bool			`json:"door_mats"`
	ShowerSoapDispensers		bool			`json:"shower_soap_dispensers"`		// ShowerDispensers
	Shampoo						bool			`json:"shampoo"`
	Conditioner					bool			`json:"conditioner"`
	BodyWash					bool			`json:"body_wash"`
	Glasses						bool			`json:"glasses"`					// Enough drinking glasses for all guests
	Bowls						bool			`json:"bowls"`						// Bowls for all Guests
	Plates						bool			`json:"plates"`						// Plates for all Guests
	Silverware					bool			`json:"silverware"`					// Silverware for all Guests
	Kitchenware					bool			`json:"kitchenware"`				// Kitchenware for all Guests
	SaltPepper					bool			`json:"salt_pepper"`
	OliveOil					bool			`json:"olive_oil"`
	HalfBathrooms				int				`json:"half_bedrooms"`				// Number of Half Bathrooms
	Quirks						string			`json:"quirks"`						// Any Special Quirks to know About Property
	DigitalAmenities			[]*string		`json:"digital_amenities"`			// AbbleTV, Roku, Video Games Systems
	BedBlankets					bool			`json:"bed_blankets"`				// Blankets for all Beds
	CoffeeMaker					bool			`json:"coffee_maker"`
	ParkingInstructions 		string			`json:"parking_instructions"`
	UniqueItems					string			`json:"unique_items"`				// Any unique property characteristics/features
	SurveillanceSystem			bool			`json:"surveillance_system"`		// Does your property have surveilance systems?
	SurveillanceLocations		[]*string		`json:"surveillance_locations"`		// Where are the cameras located
	HotTub						bool			`json:"hot_tub"`
	Pool 						bool			`json:"pool"`
	SaunaSteamRoom				bool			`json:"sauna_steam_room"`			// Sauna/Stream Room
	Conveniences 				[]*Convenience	`json:"conveniences"`				// Are there any conveniences around your property
	OwnerNeeds					[]*Need			`json:"owner_needs"`				// If yes, please provide specifics about needs
	WasherDryerLocation			string			`json:"washer_dryer_location"`		// Washer and Dryer
	WasherDryer					bool			`json:"washer_dryer"`
	Hoa 						bool			`json:"hoa"`
	NumEnsuiteBathrooms			int				`json:"num_ensuite_bathrooms"`
	BankAccount					string			`json:"bank_account"`
	RoutingNumber				string			`json:"routing_number"`
	TypeOfAccount				string			`json:"type_of_account"`
	AirbnbCreds 				*Creds			`json:"airbnb_creds"`
	VrboCreds					*Creds			`json:"vrbo_creds"`
	FireExtinguisher 			bool			`json:"fire_extinguisher"`
	Toaster						bool			`json:"toaster"`
	Plants						[]*Plant		`json:"plants"`						// Plants
	FirstAidKit					bool			`json:"first_aid_kit"`
}

type Owner struct {
	FirstName					string			`json:"first_name"`
	LastName 					string			`json:"last_name"`
	// ClientType.. maybe "time_when_enrolled"
	ReferredBy					string			`json:"referred_by"`
	Phone						string			`json:"phone"`
	Email 						string			`json:"email"`
	PreferredContactMethod		string			`json:"preferred_contact_method"` 	// Preferred Form of Contact
	Availability				string			`json:"availability"`				// Best Times to Contact
	ShorttermLicenseNumber		string			`json:"shortterm_license_number"`
	Location 					string			`json:"location"`					// Are you local?
 }

type Period struct {
	Start						time.Time		`json:"start"`
	End							time.Time		`json:"end"`
}

type DoorCode struct {
	// todo: Implement
}

type Address struct {
	// todo: Implement
}

type Warranty struct {
	//todo: Implement
}

type Bedroom struct {
	Type 						string			`json:"type"`
	Beds 						[]string		`json:"beds"`
}

type Tv struct {
	Location 					string			`json:"location"`
	Size 						string			`json:"size"`						// TC size(s) and Location(s)
	Cable						bool			`json:"cable"`						// Cable/Satellite
	Provider 					string			`json:"provider"` 					// Cable/Satellite Providor
	AccountNumber				string			`json:"account_number"`
	NameOnFile					string			`json:"name_on_file"`
	LocationOfCableBox			string			`json:"location_of_cable_box"`		// Detailed location of cable box
}

type Wifi struct {
	Name 						string			`json:"name"`						// Wifi Name
	Password					string 			`json:"password"`					// Wifi Password
	RouterLocation				string			`json:"router_location"`			// Detailed Location of Router
	Provider 					string			`json:"provider"`					// Internet Service Provider
	NameOnAccount 				string			`json:"name_on_account"`
	AccountNumber				string			`json:"account_number"`				// Internet Account Number
}

type Closet struct {
	Location 					string 			`json:"location"`					// Detailed Location of Cleaners Closet/Owners Closet
	Code 						string 			`json:"code"`						// Cleaners/Owners Closet Code
}

type Convenience struct {
	Name 						string 			`json:"name"`
	Address 					*Address 		`json:"address"`
}

type Need struct {
	//todo: Implement
}

type Creds struct {
	Username 					string			`json:"username"`
	Password					string			`json:"password"`
}

type Plant struct {
	// todo: Implement
}
