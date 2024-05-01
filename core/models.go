package core

import ()

// ************ People ************

type PeopleRoot struct {
	Links interface{} `json:"links"`
	Data  Person      `json:"data"`
}

type Person struct {
	Type       string           `json:"type"`
	ID         string           `json:"id"`
	Attributes PersonAttributes `json:"attributes"`
}

type PersonAttributes struct {
	AccountingAdministrator bool        `json:"accounting_administrator"`
	Anniversary             interface{} `json:"anniversary"`
	Avatar                  string      `json:"avatar"`
	Birthdate               string      `json:"birthdate"`
	Child                   bool        `json:"child"`
	FirstName               string      `json:"first_name"`
	Gender                  string      `json:"gender"`
	GivenName               interface{} `json:"given_name"`
	Grade                   interface{} `json:"grade"`
	GraduationYear          interface{} `json:"graduation_year"`
	InactivatedAt           interface{} `json:"inactivated_at"`
	LastName                string      `json:"last_name"`
	MedicalNotes            interface{} `json:"medical_notes"`
	Membership              string      `json:"membership"`
	MiddleName              interface{} `json:"middle_name"`
	Nickname                interface{} `json:"nickname"`
	PeoplePermissions       string      `json:"people_permissions"`
	RemoteID                interface{} `json:"remote_id"`
	SiteAdministrator       bool        `json:"site_administrator"`
	Status                  string      `json:"status"`
}

// ************ Email ************

type EmailRootNoRelationship struct {
	Data EmailNoRelationship `json:"data,omitempty"`
}

type EmailNoRelationship struct {
	Type       string          `json:"type"`
	ID         string          `json:"id"`
	Attributes EmailAttributes `json:"attributes"`
}

type EmailRoot struct {
	Data Email `json:"data,omitempty"`
}

type Email struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Attributes    EmailAttributes    `json:"attributes"`
	Relationships EmailRelationships `json:"relationships,omitempty"`
}

type EmailAttributes struct {
	Address  string `json:"address"`
	Location string `json:"location"`
	Primary  bool   `json:"primary"`
}

type EmailRelationships struct {
	Person EmailPerson `json:"person,omitempty"`
}

type EmailPerson struct {
	Data EmailPersonData `json:"data,omitempty"`
}

type EmailPersonData struct {
	Type string `json:"person,omitempty"`
	ID   string `json:"id,omitempty"`
}

// ************ Address ************

type AddressRootNoRelationship struct {
	Data AddressNoRelationship `json:"data"`
}

type AddressNoRelationship struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes AddressAttributes `json:"attributes"`
}

type AddressAttributes struct {
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	CountryCode string `json:"country_code"`
	Location    string `json:"location"`
	Primary     bool   `json:"primary"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
}

type AddressRoot struct {
	Data Address `json:"data,omitempty"`
}

type Address struct {
	Type          string               `json:"type"`
	ID            string               `json:"id"`
	Attributes    AddressAttributes    `json:"attributes"`
	Relationships AddressRelationships `json:"relationships,omitempty"`
}

type AddressRelationships struct {
	Person AddressPerson `json:"person,omitempty"`
}

type AddressPerson struct {
	Data AddressPersonData `json:"data,omitempty"`
}

type AddressPersonData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

// ************ Organization ************

type OrganizationRoot struct {
	Data Organization `json:"data"`
}

type Organization struct {
	Type       string                 `json:"type"`
	ID         string                 `json:"id"`
	Attributes OrganizationAttributes `json:"attributes"`
}

type OrganizationAttributes struct {
	Name           string `json:"name"`
	CountryCode    string `json:"country_code"`
	DateFormat     int    `json:"date_format"`
	TimeZone       string `json:"time_zone"`
	ContactWebsite string `json:"contact_website"`
}
