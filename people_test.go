package client

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var responseJSON = `{
	"data": {
		"type": "person",
		"attributes": {
			"accounting_administrator": true,
			"anniversary": null,
			"birthdate": "1990-01-01",
			"first_name": "UnitTest",
			"gender": "male",
			"given_name": null,
			"grade": null,
			"graduation_year": null,
			"inactivated_at": null,
			"last_name": "Doe",
			"medical_notes": null,
			"membership": "member",
			"middle_name": null,
			"nickname": null,
			"site_administrator": false,
			"status": "active"
		}
	}
}`

var personId string
var appId = os.Getenv("PC_APP_ID")
var secretToken = os.Getenv("PC_SECRET_TOKEN")

func TestCreatePeople(t *testing.T) {
	var data PeopleRoot

	if appId == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretToken == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into PeopleRoot
	json.Unmarshal([]byte(responseJSON), &data)

	fmt.Printf("%+v :::: %v", data, URL)
	client := NewPCClient(appId, secretToken, URL)
	fmt.Printf("%v", client)

	person := CreatePeople(client, appId, secretToken, &data)

	fmt.Printf("Returned []byte ::: %v", string(person))

	var response PeopleRoot
	json.Unmarshal(person, &response)

	personId = response.Data.ID
	if response.Data.Attributes.FirstName != "UnitTest" {
		t.Errorf("Expected person.Data.ATtributes.FirstName to be UnitTest, instead got %v\n", response.Data.Attributes.FirstName)
	}

}

func TestGetPeople(t *testing.T) {

	if appId == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretToken == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := NewPCClient(appId, secretToken, URL)

	person := GetPeople(client, appId, secretToken, personId)

	if person.Data.Attributes.FirstName != "UnitTest" {
		t.Errorf("Expected person.Data.ATtributes.FirstName to be UnitTest, instead got %v\n", person.Data.Attributes.FirstName)
	}

	if person.Data.Attributes.Birthdate != "1990-01-01" {
		t.Errorf("Expected person.Data.Attributes.Birthdate to be 1990-01-01, instead got %v\n", person.Data.Attributes.Birthdate)
	}
}

func TestDeletePeople(t *testing.T) {

	if appId == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretToken == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := NewPCClient(appId, secretToken, URL)

	DeletePeople(client, appId, secretToken, personId)

}
