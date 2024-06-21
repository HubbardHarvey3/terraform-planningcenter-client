package people

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonPhoneNumber = `{
	"data": {
		"type": "person",
		"attributes": {
			"accounting_administrator": false,
			"anniversary": null,
			"birthdate": "1990-01-01",
			"first_name": "PhoneNumberTest",
			"gender": "male",
			"given_name": null,
			"grade": null,
			"graduation_year": null,
			"inactivated_at": null,
			"last_name": "DoeAddress",
			"medical_notes": null,
			"membership": "member",
			"middle_name": null,
			"nickname": null,
			"site_administrator": false,
			"status": "active"
		}
	}
}`

var responsePhoneNumber = `{
	"data": {
		"type": "Address",
		"attributes": {
		    "number": "(123) 888-9999",
			"carrier": null,
			"location": "Mobile",
			"primary": true
		}
	}
}`

var personIdPhoneNumber string
var phoneNumberId string
var appIdPhoneNumber = os.Getenv("PC_APP_ID")
var secretTokenPhoneNumber = os.Getenv("PC_SECRET_TOKEN")

func TestCreatePhoneNumber(t *testing.T) {
	var dataPerson core.PeopleRoot
	var dataPhoneNumber core.PhoneNumberRootNoRelationship

	if appIdPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into core.PeopleRoot
	err := json.Unmarshal([]byte(responsePersonPhoneNumber), &dataPerson)
	if err != nil {
		t.Error(err)
	}

	client := core.NewPCClient(appIdPhoneNumber, secretTokenPhoneNumber)

	person, err := CreatePeople(client, &dataPerson)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var responsePerson core.PeopleRoot
	json.Unmarshal(person, &responsePerson)

	personIdPhoneNumber = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responsePhoneNumber), &dataPhoneNumber)
	if err != nil {
		t.Error(err)
	}

	addressBytes, err := CreatePhoneNumber(client, personIdPhoneNumber, &dataPhoneNumber)

	var address core.PhoneNumberRootNoRelationship
	json.Unmarshal(addressBytes, &address)
	addressId = address.Data.ID

	if address.Data.Attributes.Location != "Mobile" {
		t.Errorf("Location is not 'Mobile', but is showing as : %v", address.Data.Attributes.Location)
	}

}

func TestGetPhoneNumber(t *testing.T) {
	var phoneNumber core.PhoneNumberRoot

	if appIdPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdPhoneNumber, secretTokenPhoneNumber)

	phoneNumber, err := GetAllPhoneNumbers(client, phoneNumberId)
	if err != nil {
		t.Errorf("GetphoneNumber failed with an error ::: %v\n", err)
	}

	if phoneNumber.Data[0].Attributes.Number != "(555) 555-5555" {
		t.Errorf("phoneNumber is not '(555) 555-5555', but is showing as : %v", phoneNumber.Data[0].Attributes.Number)
	}
}

/*
The Get request for an object returns the relationships listed in the json.
Therefore, the struct that is used with GET requests, should have the relationships.
For Updates, you get a 422 if you attempt to update using a json payload that contains relationships
For now, I am copying the attributes from the Root struct to the RootNoRelationship model

TODO - Fix Email Updates and create test
*/
func TestUpdatePhoneNumber(t *testing.T) {
	var address core.AddressRoot

	if appIdAddress == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenAddress == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdAddress, secretTokenAddress)

	address, err := GetAddress(client, addressId)
	if err != nil {
		t.Errorf("GetAddress failed with an error ::: %v\n", err)
	}

	address.Data.Attributes.City = "Updated"
	// Alter to without Relationships .... TODO Make this better
	var updatedAddress core.AddressRootNoRelationship
	updatedAddress.Data.Attributes = address.Data.Attributes

	response, err := UpdateAddress(client, addressId, &updatedAddress)

	json.Unmarshal(response, &updatedAddress)

	if updatedAddress.Data.Attributes.City != "Updated" {
		t.Errorf("Address is not 'Updated', but is showing as : %v", updatedAddress.Data.Attributes.City)
	}

}

func TestDeletePhoneNumber(t *testing.T) {
	if appIdPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenPhoneNumber == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdPhoneNumber, secretTokenPhoneNumber)

	err := DeletePhoneNumber(client, phoneNumberId)
	if err != nil {
		t.Errorf("Error during DeletePhoneNumber : %v\n", err)
	}

	// GetAllPhoneNumbers can't take a phone number ID, need to figure out a different way of testing a deleted number is gone
	// until this is resolved, you will need to confirm the phone number is gone before you delete the person...yuck
	//	_, err = GetAllPhoneNumbers(client, personIdPhoneNumber)
	//	if !strings.Contains(err.Error(), "404") {
	//		t.Errorf("GetPhoneNumber should be throwing a 404 after the person was deleted.  Error was %v", err)
	//	}

	err = DeletePeople(client, personIdPhoneNumber)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}

}
