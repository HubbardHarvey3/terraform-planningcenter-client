package core

// ************ PhoneNumber ************

type PhoneNumberRootNoRelationship struct {
	Data PhoneNumberNoRelationship `json:"data,omitempty"`
}

type PhoneNumberNoRelationship struct {
	Type       string                `json:"type"`
	ID         string                `json:"id"`
	Attributes PhoneNumberAttributes `json:"attributes"`
}

type PhoneNumberRoot struct {
	Data []PhoneNumber `json:"data,omitempty"`
}

type PhoneNumber struct {
	Type          string                   `json:"type"`
	ID            string                   `json:"id"`
	Attributes    PhoneNumberAttributes    `json:"attributes"`
	Relationships PhoneNumberRelationships `json:"relationships,omitempty"`
}

type PhoneNumberAttributes struct {
	Carrier  string `json:"carrier"`
	Location string `json:"location"`
	Number   string `json:"number"`
	Primary  bool   `json:"primary"`
}

type PhoneNumberRelationships struct {
	Person PhoneNumberPerson `json:"person,omitempty"`
}

type PhoneNumberPerson struct {
	Data PhoneNumberPersonData `json:"data,omitempty"`
}

type PhoneNumberPersonData struct {
	Type string `json:"person,omitempty"`
	ID   string `json:"id,omitempty"`
}
