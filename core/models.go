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
	Data AddressPersonData `json:"person,omitempty"`
}

type AddressPersonData struct {
	Type string `json:"person,omitempty"`
	ID   string `json:"id,omitempty"`
}

/*
{
  "links": {
    "self": "https://api.planningcenteronline.com/people/v2/addresses"
  },
  "data": [
    {
      "type": "Address",
      "id": "112014048",
      "attributes": {
        "city": "Hometown",
        "country_code": "US",
        "country_name": "United States",
        "created_at": "2024-04-23T10:57:38Z",
        "location": "Home",
        "primary": true,
        "state": "GA",
        "street": "1234\nCool Name Avenue",
        "updated_at": "2024-04-23T10:57:38Z",
        "zip": "555555"
      },
      "relationships": {
        "person": {
          "data": {
            "type": "Person",
            "id": "138378248"
          }
        }
      },
      "links": {
        "self": "https://api.planningcenteronline.com/people/v2/addresses/112014048"
      }
    }
  ],
  "included": [],
  "meta": {
    "total_count": 1,
    "count": 1,
    "can_order_by": [
      "city",
      "state",
      "zip",
      "country_code",
      "location",
      "primary",
      "created_at",
      "updated_at",
      "street_line_1",
      "street_line_2",
      "street"
    ],
    "can_query_by": [
      "city",
      "state",
      "zip",
      "country_code",
      "location",
      "primary",
      "street"
    ],
    "parent": {
      "id": "458241",
      "type": "Organization"
    }
  }
}
*/
