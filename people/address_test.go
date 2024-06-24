package people

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonAddress = `{
	"data": {
		"type": "person",
		"attributes": {
			"accounting_administrator": false,
			"anniversary": null,
			"birthdate": "1990-01-01",
			"first_name": "UnitTest",
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

var responseAddress = `{
	"data": {
		"type": "Address",
		"attributes": {
			"city": "Hometown",
			"country_code": "US",
			"location": "Home",
			"primary": true,
			"state": "GA",
			"street_line_1": "12349 Cool Name Ave",
			"street_line_2": ""
		}
	}
}`

var personIdAddress string
var addressId string
var appIdAddress = os.Getenv("PC_APP_ID")
var secretTokenAddress = os.Getenv("PC_SECRET_TOKEN")

func TestCreateAddress(t *testing.T) {
	var dataPerson core.PeopleRoot
	var dataAddress core.AddressRoot

	if appIdAddress == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenAddress == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into core.PeopleRoot
	err := json.Unmarshal([]byte(responsePersonAddress), &dataPerson)
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

	personIdAddress = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responseAddress), &dataAddress)
	if err != nil {
		t.Error(err)
	}

	addressBytes, err := CreateAddress(client, personIdAddress, &dataAddress)

	var address core.AddressRoot
	err = json.Unmarshal(addressBytes, &address)
	if err != nil {
		t.Errorf("Error unmarshalling addressBytes")
	}
	addressId = address.Data.ID

	if address.Data.Attributes.CountryCode != "US" {
		t.Errorf("Country Code is not US, but is showing as : %v", address.Data.Attributes.CountryCode)
	}

}

func TestGetAddress(t *testing.T) {
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

	if address.Data.Attributes.City != "Hometown" {
		t.Errorf("Address is not 'Hometown', but is showing as : %v", address.Data.Attributes.City)
	}
}

func TestUpdateAddress(t *testing.T) {
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
	var updatedAddress core.AddressRoot
	updatedAddress.Data.Attributes = address.Data.Attributes

	response, err := UpdateAddress(client, addressId, &updatedAddress)

	json.Unmarshal(response, &updatedAddress)

	if updatedAddress.Data.Attributes.City != "Updated" {
		t.Errorf("Address is not 'Updated', but is showing as : %v", updatedAddress.Data.Attributes.City)
	}

}

func TestDeleteAddress(t *testing.T) {
	if appIdAddress == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenAddress == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdAddress, secretTokenAddress)

	err := DeleteAddress(client, addressId)
	if err != nil {
		t.Errorf("Error during DeleteAddress : %v\n", err)
	}

	_, err = GetAddress(client, personIdAddress)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetAddress should be throwing a 404 after the person was deleted.  Error was %v", err)
	}

	err = DeletePeople(client, personIdAddress)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}

}
