package core

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
