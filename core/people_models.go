package core

// ************ People ************

type PeopleRoot struct {
	Links interface{} `json:"links"`
	Data  Person      `json:"data"`
}

type Person struct {
	Type          string               `json:"type"`
	ID            string               `json:"id"`
	Attributes    PersonAttributes     `json:"attributes"`
	Relationships *PersonRelationships `json:"relationships"`
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

type PersonRelationships struct {
	PrimaryCampus PersonRelationshipsData `json:"primary_Campus"`
	Gender        PersonRelationshipsData `json:"gender"`
}

type PersonRelationshipsData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
