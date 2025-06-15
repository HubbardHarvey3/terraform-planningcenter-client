package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var phoneNumberPersonId string = "123456789"

func TestCreatePhoneNumber(t *testing.T) {
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
	var dataPhoneNumber core.PhoneNumberRoot
	var mockServer *httptest.Server = setupMockServer(responsePhoneNumber, http.StatusOK)

	//Convert json into core.PeopleRoot
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	err := json.Unmarshal([]byte(responsePhoneNumber), &dataPhoneNumber)
	if err != nil {
		t.Error(err)
	}

	phoneNumberBytes, err := CreatePhoneNumber(client, phoneNumberPersonId, &dataPhoneNumber)

	var phoneNumber core.PhoneNumberRoot
	json.Unmarshal(phoneNumberBytes, &phoneNumber)

	if phoneNumber.Data.Attributes.Location != "Mobile" {
		t.Errorf("Location is not 'Mobile', but is showing as : %v", phoneNumber.Data.Attributes.Location)
	}

}

func TestGetPhoneNumber(t *testing.T) {
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
	var mockServer *httptest.Server = setupMockServer(responsePhoneNumber, http.StatusOK)
	var phoneNumber core.PhoneNumberRoot

	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	phoneNumber, err := GetPhoneNumber(client, phoneNumberPersonId)
	if err != nil {
		t.Errorf("GetphoneNumber failed with an error ::: %v\n", err)
	}

	if phoneNumber.Data.Attributes.Number != "(123) 888-9999" {
		t.Errorf("phoneNumber is not '(123) 888-9999', but is showing as : %v", phoneNumber.Data.Attributes.Number)
	}
}

func TestUpdatePhoneNumber(t *testing.T) {
	var responsePhoneNumber = `{
		"data": {
			"type": "Address",
			"attributes": {
				"number": "(123) 888-9999",
				"carrier": null,
				"location": "Home",
				"primary": true
			}
		}
	}`
	var phoneNumber core.PhoneNumberRoot
	var mockServer *httptest.Server = setupMockServer(responsePhoneNumber, http.StatusOK)

	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	phoneNumber, err := GetPhoneNumber(client, phoneNumberPersonId)
	if err != nil {
		t.Errorf("GetPhoneNumber failed with an error ::: %v\n", err)
	}

	phoneNumber.Data.Attributes.Location = "Home"
	// Alter to without Relationships .... TODO Make this better
	var updatedPhoneNumber core.PhoneNumberRoot
	updatedPhoneNumber.Data.Attributes = phoneNumber.Data.Attributes

	response, err := UpdatePhoneNumber(client, phoneNumberPersonId, &updatedPhoneNumber)

	json.Unmarshal(response, &updatedPhoneNumber)

	if updatedPhoneNumber.Data.Attributes.Location != "Home" {
		t.Errorf("Location is not 'Home', but is showing as : %v", updatedPhoneNumber.Data.Attributes.Location)
	}

}

func TestDeletePhoneNumber(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeletePhoneNumber(client, phoneNumberPersonId)
	if response != nil {
		t.Errorf("DeletePhoneNumber response should have been nil, but returned something")
	}

}
