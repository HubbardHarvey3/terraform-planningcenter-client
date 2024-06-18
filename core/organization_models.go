package core

import ()

type OrganizationRootNoRelationship struct {
	Data Organization `json:"data"`
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

type OrganizationRoot struct {
	Data []Organization `json:"data"`
}

type Organization struct {
	Type          string                    `json:"type"`
	ID            string                    `json:"id"`
	Attributes    OrganizationAttributes    `json:"attributes"`
	Relationships OrganizationRelationships `json:"relationships"`
}

type OrganizationRelationships struct {
	Data OrganizationAddress `json:"person"`
}

type OrganizationAddress struct {
	Person OrganizationAddressData `json:"data"`
}

type OrganizationAddressData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
