package core

// ************ Campus ************

type CampusRoot struct {
	Data Campus `json:"data,omitempty"`
}

type Campus struct {
	Type          string               `json:"type"`
	ID            string               `json:"id"`
	Attributes    CampusAttributes     `json:"attributes"`
	Relationships *CampusRelationships `json:"relationships,omitempty"`
}

type CampusAttributes struct {
	Latitude                 string `json:"latitude"`
	Longitude                string `json:"longitude"`
	Description              string `json:"description"`
	Street                   string `json:"street"`
	City                     string `json:"city"`
	State                    string `json:"state"`
	Zip                      string `json:"zip"`
	Country                  string `json:"country"`
	PhoneNumber              string `json:"phone_number"`
	Website                  string `json:"website"`
	TwentyFourHourTime       bool   `json:"twenty_four_hour_time"`
	DateFormat               int32  `json:"date_format"`
	ChurchCenterEnabled      bool   `json:"church_center_enabled"`
	ContactEmailAddress      string `json:"contact_email_address"`
	TimeZone                 string `json:"time_zone"`
	Geolocation_set_manually bool   `json:"geolocation_set_manually"`
	Name                     string `json:"name"`
}

type CampusRelationships struct {
	Organization CampusOrg `json:"organization,omitempty"`
}

type CampusOrg struct {
	Data CampusPersonData `json:"data,omitempty"`
}

type CampusPersonData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
