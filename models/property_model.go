package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Notes:
//	Many duplicates
// 	Sort out Amenities
//	Sort out cleaner checks
// 	Sort out what data should be altered and read by which people
// 	Need to classify types of data
//		+ Amenities
//		+ Extras

var MongoAddr *string

type Property struct {
	Id 							*bson.ObjectId	`json:"id" bson:"_id"`
	LodgixId 					int				`json:"lodgix_id" bson:"lodgix_id"`
	Address						*Address 		`json:"address" bson:"address"`
	HouseType					string			`json:"house_type" bson:"house_type"`
	Bedrooms					[]*Bedroom		`json:"beds" bson:"beds"`			// Number of bedsrooms
	Bathrooms					int				`json:"baths" bson:"baths"`			// Number of Bathrooms
	Sleeps 						int				`json:"sleeps" bson:"sleeps"`
	Owner 						*Owner			`json:"owner" bson:"owner"`
	PutOnMarket					time.Time		`json:"date_on_market" bson:"date_on_market"`	// When is Property Open
	BlackoutDates				[]*Period		`json:"blackout_period" bson:"blackout_period"`	// Details about future dates to block
	IsShorterm					bool			`json:"is_shorterm" bson:"is_shorterm"`	// Currently Short-term Rental
	FutureReservations 			bool			`json:"future_reservations" bson:"future_reservations"`
	FirstReservation 			time.Time		`json:"first_reservation" bson:"first_reservation"`
	NumFutureReservations 		int				`json:"num_future_reservations" bson:"num_future_reservations"` 	// How Many Future Reservations
	LockType					string			`json:"lock_type" bson:"lock_type"`	// Type of lock on property
	DoorCodes 					[]*DoorCode		`json:"door_codes" bson:"door_codes"`
	CommunityCode				string			`json:"community_code" bson:"community_id"`	// Community/Gate Code
	BuildingCode				string			`json:"building_code" bson:"building_code"`
	IntercomSystem				bool 			`json:"intercom_system" bson:"intercom_system"`
	Warranties 					[]*Warranty		`json:"warranties" bson:"warranties"`
	KnownIssues					[]string		`json:"known_issues" bson:"known_issues"`
	SnowRemoval					bool			`json:"snow_removal" bson:"snow_removal"`	// Does the Client want Snow Removal, Yard, ect ; Will your property need snow removal
	Landscaping					bool			`json:"landscaping" bson:"landscaping"`		// Does the Client want Snow Removal, Yard, ect ; Will your property need lawn care
	MaxOccupancy				int				`json:"max_occupancy" bson:"max_occupancy"`
	Floors						int				`json:"floors" bson:"floors"`		// Number of Floors
	LocationOfLinens			string			`json:"location_of_linens" bson:"location_of_lines"`	// Location of Extra Pillows, Blankets, ect
	AirConditioning 			bool			`json:"air_conditioning" bson:"air_conditioning"`
	Heating						bool			`json:"heating" bson:"heating"`
	ThermostatLocation 			string			`json:"thermostat_location" bson:"thermostat_location"` // Detailed Location of Thermostat
	SecuritySystem 				bool			`json:"security_system" bson:"security_system"`	// Home Security
	SecuritySystemCode			string 			`json:"security_system_code" bson:"security_system_code"`	// Hoe Security Code
	MainWaterLocation			string			`json:"main_water_location" bson:"main_water_locatoin"`	// Detailed Location of Water Shut Off Valve
	WaterHeaterLocation 		string			`json:"water_heater_location" bson:"water_heater_location"`	// Detailed Location of Water Heater
	SmokeDetectors 				bool			`json:"smoke_detectors" bson:"smoke_detectors"`
	CarbonMonoxideDetector		bool			`json:"carbon_monoxide_detector" bson:"carbon_monoxide_detector"`
	Tvs							[]*Tv 			`json:"tvs" bson:"tvs"`
	Wifi 						*Wifi 			`json:"wifi" bson:"wifi"`
	Stereo 						bool			`json:"stereo" bson:"stereo"`
	BreakersLocations			[]*string		`json:"breakers_locations" bson:"breaker_locations"` // Detailed Location of Breakers
	FitnessCenter				bool			`json:"fitness_center" bson:"fitness_center"`	// Detailed Gym Locatoin
	FitnessCenterLocation 		string			`json:"fitness_center_location" bson:"fitness_center_location"`	// Detailed Gym Location
	CleanersCloset				*Closet			`json:"cleaners_closet" bson:"cleaners_closet"`
	OwnersCloset				*Closet			`json:"owners_closet" bson:"owners_closet"`
	RooftopDeckPatio			bool			`json:"rooftop_deck_patio" bson:"rooftop_deck_patio"` 	// Rooftop/Deck/Patio
	Grill 						bool			`json:"grill" bson:"grill"`			// BBQ/Grill
	GrillType					string			`json:"grill_type" bson:"grill_type"`	// Type of Grill
	HandicapAccessible			bool			`json:"handicap_accessible" bson:"handicap_accessible"`
	TrashPickup					string 			`json:"trash_pickup" bson:"trash_pickup"`	// Track Pickup Day
	TrashLocation				string			`json:"trash_location" bson:"trash_location"`	// Location ad Instructions for Trash
	TrashInstructions			string			`json:"trash_instructions" bson:"trash_instructions"`	// Location ad Instructions for Trash
	RecyclingPickup 			string			`json:"recycling_pickup" bson:"recycling_pickup"`	// Recycling Pickup Day
	RecyclingLocation			string			`json:"recycling_location" bson:"recycling_location"`	// Location and Instructions for Trach & Recycling
	RecyclingInstruction 		string			`json:"recycling_instruction" bson:"recycling_instruction"`	// Location and Instructions for Trach & Recycling
	Elevator 					bool			`json:"elevator" bson:"elevator"`
	Parking						string			`json:"parking" bson:"parking"`
	Vacuum						bool			`json:"vacuum" bson:"vacuum"`
	HairDryer					bool			`json:"hair_dryer" bson:"hair_dryer"`
	ExtraSheets					bool			`json:"extra_sheets" bson:"extra_sheets"`	// 3 Set of Sheets for Each Bed
	Duvets						bool			`json:"duvets" bson:"duvets"`		// Duvets for all blankets
	Iron						bool			`json:"iron" bson:"iron"`
	IroningBoard				bool			`json:"ironing_board" bson:"ironing_board"`
	Broom						bool			`json:"broom" bson:"broom"`
	ToiletPlunger				bool			`json:"toilet_plunger" bson:"toilet_plunger"`
	BathMats					bool			`json:"bath_mats" bson:"bath_mats"`
	ToiletBowlBrush				bool			`json:"toilet_bowl_brush" bson:"toilet_bowl_brush"`
	WaterproofMattressCovers	bool			`json:"waterproof_mattress_covers" bson:"waterproof_mattress_covers"`	// Waterproof Mattress Covers for all Beds
	ToiletPaper					bool			`json:"toilet_paper" bson:"toilet_paper"`
	PaperTowels					bool			`json:"paper_towels" bson:"paper_towels"`
	HandSoap					bool			`json:"hand_soap" bson:"hand_soap"`					// Hand Soap for all Sinks
	DishwasherSoap				bool			`json:"dishwasher_soap" bson:"dishwasher_soap"`
	Coffee						bool			`json:"coffee" bson:"coffee"`
	BathTrashCan				bool			`json:"bath_trash_can" bson:"bath_trash_can"`	// Is there a Trashcan in every Bathroom
	TrashBags					bool			`json:"trash_bags" bson:"trash_bags"`	// Backup Trash Bags
	DoorMats					bool			`json:"door_mats" bson:"door_mats"`
	ShowerSoapDispensers		bool			`json:"shower_soap_dispensers" bson:"shower_soap_dispensers"`	// ShowerDispensers
	Shampoo						bool			`json:"shampoo" bson:"shampoo"`
	Conditioner					bool			`json:"conditioner" bson:"conditioner"`
	BodyWash					bool			`json:"body_wash" bson:"body_wash"`
	Glasses						bool			`json:"glasses" bson:"glasses"`		// Enough drinking glasses for all guests
	Bowls						bool			`json:"bowls" bson:"bowls"`			// Bowls for all Guests
	Plates						bool			`json:"plates" bson:"players"`		// Plates for all Guests
	Silverware					bool			`json:"silverware" bson:"silberware"`	// Silverware for all Guests
	Kitchenware					bool			`json:"kitchenware" bson:"kitchenware"`	// Kitchenware for all Guests
	SaltPepper					bool			`json:"salt_pepper" bson:"salt_pepper"`
	OliveOil					bool			`json:"olive_oil" bson:"olive_oil"`
	HalfBathrooms				int				`json:"half_bedrooms" bson:"half_bedrooms"`	// Number of Half Bathrooms
	Quirks						string			`json:"quirks" bson:"quirks"`		// Any Special Quirks to know About Property
	DigitalAmenities			[]*string		`json:"digital_amenities" bson:"digital_amenities"`	// AbbleTV, Roku, Video Games Systems
	BedBlankets					bool			`json:"bed_blankets" bson:"bed_blankets"`	// Blankets for all Beds
	CoffeeMaker					bool			`json:"coffee_maker" bson:"coffee_maker"`
	ParkingInstructions 		string			`json:"parking_instructions" bson:"parking_instructions"`
	UniqueItems					string			`json:"unique_items" bson:"unique_items"`	// Any unique property characteristics/features
	SurveillanceSystem			bool			`json:"surveillance_system" bson:"surveillance_system"`	// Does your property have surveilance systems?
	SurveillanceLocations		[]*string		`json:"surveillance_locations" bson:"surveillance_locations"`	// Where are the cameras located
	HotTub						bool			`json:"hot_tub" bson:"hot_tub"`
	Pool 						bool			`json:"pool" bson:"pool"`
	SaunaSteamRoom				bool			`json:"sauna_steam_room" bson:"sauna_steam_room"`	// Sauna/Stream Room
	Conveniences 				[]*Convenience	`json:"conveniences" bson:"conveniences"`	// Are there any conveniences around your property
	OwnerNeeds					[]*Need			`json:"owner_needs" bson:"owner_needs"`	// If yes, please provide specifics about needs
	WasherDryerLocation			string			`json:"washer_dryer_location" bson:"waher_dryer_location"`	// Washer and Dryer
	WasherDryer					bool			`json:"washer_dryer" bson:"washer_dryer"`
	Hoa 						bool			`json:"hoa" bson:"hoa"`
	NumEnsuiteBathrooms			int				`json:"num_ensuite_bathrooms" bson:"num_ensuite_bathrooms"`
	BankAccount					string			`json:"bank_account" bson:"bank_account"`
	RoutingNumber				string			`json:"routing_number" bson:"routing_number"`
	TypeOfAccount				string			`json:"type_of_account" bson:"type_of_account"`
	AirbnbCreds 				*Creds			`json:"airbnb_creds" bson:"airbnb_creds"`
	VrboCreds					*Creds			`json:"vrbo_creds" bson:"vrbo_creds"`
	FireExtinguisher 			bool			`json:"fire_extinguisher" bson:"fire_extinguisher"`
	Toaster						bool			`json:"toaster" bson:"toaster"`
	Plants						[]*Plant		`json:"plants" bson:"plants"`	// Plants
	FirstAidKit					bool			`json:"first_aid_kit" bson:"first_aid_kit"`
}

func NewProperty() *Property {

	p := new(Property)
	id := bson.NewObjectId()
	p.Id = &id

	return p
}

func (p *Property) CreateProperty() error {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("properties")
	_, err = c.UpsertId(p.Id, p)
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return err
	}

	return nil
}

func DeleteProperty(id string) error {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("properties")
	err = c.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Println("Error deleteing Property: ", err.Error())
		return err
	}

	return nil
}

func FindPropertyById(id string) (*Property, error) {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}

	var property *Property
	c := session.DB("elt").C("properties")
	defer session.Close()
	err = c.FindId(bson.ObjectIdHex(id)).One(&property)
	if err != nil {
		return nil, err
	}

	return property, nil
}

func GetProperties() ([]*Property, error) {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("properties")
	var properties []*Property
	err = c.Find(bson.M{}).All(&properties)
	if err != nil {
		log.Println("Could not find properties: ", err.Error())
		return nil, err
	}

	return properties, nil
}

func RetrieveLodgixProperties() ([]*interface{}, error) {

	var client http.Client
	resp, err := client.Get("http://35.193.114.158:1323/props")
	//resp, err := client.Get("http://localhost:1323/props")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type lodgixResponse struct {
		Count 			int 				`json:"count"`
		Filters 		[]*interface{} 		`json:"filters"`
		Properties 		[]*interface{} 		`json:"properties"`
	}
	lp := new(lodgixResponse)
	err = json.Unmarshal(bodyBytes, lp)
	if err != nil {
		return nil, err
	}

	var properties []*interface{}
	for _, v := range lp.Properties {
		properties = append(properties, v)
	}

	return properties, nil
}

func (p *Property) UpdateProperty() (error) {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("properties")
	_, err = c.UpsertId(p.Id, p)
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return err
	}

	return nil
}

type Period struct {
	Start						time.Time		`json:"start"`
	End							time.Time		`json:"end"`
}

type DoorCode struct {
	// todo: Implement
}

type Warranty struct {
	// todo: Implement
}

type Bedroom struct {
	Type 						string			`json:"type" bson:"type"`
	Beds 						[]string		`json:"beds" bson:"beds"`
}

type Tv struct {
	Location 					string			`json:"location" bson:"location"`
	Size 						string			`json:"size" bson:"size"`			// TC size(s) and Location(s)
	Cable						bool			`json:"cable" bson:"cable"`			// Cable/Satellite
	Provider 					string			`json:"provider" bson:"_id"`  		// Cable/Satellite Providor
	AccountNumber				string			`json:"account_number" bson:"account_number"`
	NameOnFile					string			`json:"name_on_file" bson:"name_on_file"`
	LocationOfCableBox			string			`json:"location_of_cable_box" bson:"location_of_cable_box"`	// Detailed location of cable box
}

type Wifi struct {
	Name 						string			`json:"name" bson:"name"`			// Wifi Name
	Password					string 			`json:"password" bson:"password"`	// Wifi Password
	RouterLocation				string			`json:"router_location" bson:"router_location"`	// Detailed Location of Router
	Provider 					string			`json:"provider" bson:"provider"`	// Internet Service Provider
	NameOnAccount 				string			`json:"name_on_account" bson:"name_on_account"`
	AccountNumber				string			`json:"account_number" bson:"account_number"`	// Internet Account Number
}

type Closet struct {
	Location 					string 			`json:"location" bson:"location"`	// Detailed Location of Cleaners Closet/Owners Closet
	Code 						string 			`json:"code" bson:"code"`			// Cleaners/Owners Closet Code
}

type Convenience struct {
	Name 						string 			`json:"name" bson:"name"`
	Address 					*Address 		`json:"address" bson:"address"`
}

type Need struct {
	//todo: Implement
}

type Creds struct {
	Username 					string			`json:"username" bson:"username"`
	Password					string			`json:"password" bson:"password"`
}

type Plant struct {
	// todo: Implement
}
