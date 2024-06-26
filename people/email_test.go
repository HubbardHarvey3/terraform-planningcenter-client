package people

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonEmail = `{
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
			"address": "john.doe@gmail.com",
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
	var dataPerson core.PeopleRoot
	var dataEmail core.EmailRoot

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into core.PeopleRoot
	err := json.Unmarshal([]byte(responsePersonEmail), &dataPerson)
	if err != nil {
		t.Error(err)
	}

	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	person, err := CreatePeople(client, &dataPerson)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var responsePerson core.PeopleRoot
	json.Unmarshal(person, &responsePerson)

	personId = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responseEmail), &dataEmail)
	if err != nil {
		t.Error(err)
	}

	emailBytes, err := CreateEmail(client, personId, &dataEmail)
	if err != nil {
		t.Error(err)
	}

	var email core.EmailRoot
	json.Unmarshal(emailBytes, &email)
	emailId = email.Data.ID

	if email.Data.Attributes.Address != "john.doe@gmail.com" {
		t.Errorf("Address is not john.doe@gmail.com, but is showing as : %v", email.Data.Attributes.Address)
	}

}

func TestGetEmail(t *testing.T) {
	var email core.EmailRoot

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	email, err := GetEmail(client, emailId)
	if err != nil {
		t.Errorf("GetPeople failed with an error ::: %v\n", err)
	}

	if email.Data.Attributes.Address != "john.doe@gmail.com" {
		t.Errorf("Address is not john.doe@gmail.com, but is showing as : %v", email.Data.Attributes.Address)
	}
}

func TestUpdateEmail(t *testing.T) {
	var email core.EmailRoot

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	email, err := GetEmail(client, emailId)
	if err != nil {
		t.Errorf("Getemail failed with an error ::: %v\n", err)
	}

	email.Data.Attributes.Address = "john.doe.updated@gmail.com"

	// Convert EmailRoot to EmailRoots
	var updatedEmail core.EmailRoot
	updatedEmail.Data.Attributes = email.Data.Attributes

	response, err := UpdateEmail(client, emailId, &updatedEmail)
	if err != nil {
		t.Error(err)
	}

	json.Unmarshal(response, &updatedEmail)

	if updatedEmail.Data.Attributes.Address != "john.doe.updated@gmail.com" {
		t.Errorf("email is not 'john.doe.updated@gmail.com', but is showing as : %v", updatedEmail.Data.Attributes.Address)
	}

}

func TestDeleteEmail(t *testing.T) {

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	err := DeleteEmail(client, emailId)
	if err != nil {
		t.Errorf("Error during DeleteEmail : %v\n", err)
	}

	_, err = GetEmail(client, emailId)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetEmail should be throwing a 404 after the person was deleted")
	}

	err = DeletePeople(client, personId)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}
}
