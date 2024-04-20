package client

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

var responsePerson = `{
	"data": {
		"type": "person",
		"attributes": {
			"accounting_administrator": false,
			"anniversary": null,
			"birthdate": "1990-01-01",
			"first_name": "UnitTestEmail",
			"gender": "male",
			"given_name": null,
			"grade": null,
			"graduation_year": null,
			"inactivated_at": null,
			"last_name": "MailDoe",
			"medical_notes": null,
			"membership": "member",
			"middle_name": null,
			"nickname": null,
			"site_administrator": false,
			"status": "active"
		}
	}
}`

var responseEmail = `{
	"data": {
		"type": "Email",
		"attributes": {
			"address": "john.doe@example.com",
			"location": "Home",
			"primary": true
		}
	} 
}
`

var emailId string
var appIdEmail = os.Getenv("PC_APP_ID")
var secretTokenEmail = os.Getenv("PC_SECRET_TOKEN")

func TestCreateEmail(t *testing.T) {
	var dataPerson PeopleRoot
	var dataEmail EmailRootNoRelationship

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into PeopleRoot
	err := json.Unmarshal([]byte(responseJSON), &dataPerson)
	if err != nil {
		t.Error(err)
	}

	client := NewPCClient(appIdEmail, secretTokenEmail, URL)

	person, err := CreatePeople(client, appIdEmail, secretTokenEmail, &dataPerson)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var responsePerson PeopleRoot
	json.Unmarshal(person, &responsePerson)

	personId = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responseEmail), &dataEmail)
	if err != nil {
		t.Error(err)
	}

	emailBytes, err := CreateEmail(client, appIdEmail, secretTokenEmail, personId, &dataEmail)

	var email EmailRootNoRelationship
	json.Unmarshal(emailBytes, &email)
	emailId = email.Data.ID

	if email.Data.Attributes.Address != "john.doe@example.com" {
		t.Errorf("Address is not john.doe@example.com, but is showing as : %v", email.Data.Attributes.Address)
	}

}

func TestGetEmail(t *testing.T) {
	var email EmailRoot

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := NewPCClient(appIdEmail, secretTokenEmail, URL)

	email, err := GetEmail(client, appIdEmail, secretTokenEmail, emailId)
	if err != nil {
		t.Errorf("GetPeople failed with an error ::: %v\n", err)
	}

	if email.Data.Attributes.Address != "john.doe@example.com" {
		t.Errorf("Address is not john.doe@example.com, but is showing as : %v", email.Data.Attributes.Address)
	}
}

func TestDeleteEmail(t *testing.T) {

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := NewPCClient(appIdEmail, secretTokenEmail, URL)

	err := DeleteEmail(client, appIdEmail, secretTokenEmail, emailId)
	if err != nil {
		t.Errorf("Error during DeleteEmail : %v\n", err)
	}

	_, err = GetEmail(client, appIdEmail, secretTokenEmail, emailId)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetEmail should be throwing a 404 after the person was deleted")
	}

	err = DeletePeople(client, appIdEmail, secretTokenEmail, personId)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}
}
