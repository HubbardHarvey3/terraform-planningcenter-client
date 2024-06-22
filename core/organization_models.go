package core

import ()

type OrganizationRootNoRelationship struct {
	Data OrganizationNoRelationship `json:"data"`
}

type OrganizationNoRelationship struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes OrganizationAttributes `json:"attributes"`
}

type OrganizationAttributes struct {
	Name           string `json:"name"`
	CountryCode    string `json:"country_code"`
	DateFormat     string `json:"date_format"`
	TimeZone       string `json:"time_zone"`
	ContactWebsite string `json:"contact_website"`
}

type OrganizationRootAddress struct {
	Data []OrganizationAddress `json:"data"`
}

type OrganizationAddress struct {
	Type          string                           `json:"type"`
	ID            string                           `json:"id"`
	Attributes    OrganizationAttributesAddress    `json:"attributes"`
	Relationships OrganizationRelationshipsAddress `json:"relationships"`
}

type OrganizationAttributesAddress struct {
	City          string `json:"city"`
	Country_Code  string `json:"country_code"`
	Country_Name  string `json:"country_name"`
	Location      string `json:"location"`
	State         string `json:"state"`
	Street_Line_1 string `json:"street_line_1"`
	Street_Line_2 string `json:"street_line_2"`
	Zip           string `json:"zip"`
}

type OrganizationRelationshipsAddress struct {
	Person OrganizationRelationshipPerson `json:"person"`
}

type OrganizationRelationshipPerson struct {
	Data OrganizationRelationshipPersonData `json:"data"`
}

type OrganizationRelationshipPersonData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
