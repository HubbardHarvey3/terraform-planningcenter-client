package core

// ************ Email ************

type EmailRoot struct {
	Data Email `json:"data,omitempty"`
}

type Email struct {
	Type          string              `json:"type"`
	ID            string              `json:"id"`
	Attributes    EmailAttributes     `json:"attributes"`
	Relationships *EmailRelationships `json:"relationships,omitempty"`
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
