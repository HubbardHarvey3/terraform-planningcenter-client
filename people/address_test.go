package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

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

var addressId = "1111111"
var personIdAddress = "123456789"

func TestCreateAddress(t *testing.T) {
	var dataAddress core.AddressRoot

	var mockServer *httptest.Server = setupMockServer(responseAddress, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	err := json.Unmarshal([]byte(responseAddress), &dataAddress)
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

	var mockServer *httptest.Server = setupMockServer(responseAddress, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

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

	var mockServer *httptest.Server = setupMockServer(responseAddress, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response, err := UpdateAddress(client, addressId, &address)
	if err != nil {
		t.Errorf("UpdateAddress failed with an error ::: %v\n", err)
	}

	json.Unmarshal(response, &address)

	if address.Data.Attributes.City != "Hometown" {
		t.Errorf("Address is not 'Hometown', but is showing as : %v", address.Data.Attributes.City)
	}

}

func TestDeleteAddress(t *testing.T) {
	var mockServer *httptest.Server = setupMockServer(responseAddress, http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeleteAddress(client, addressId)
	if response != nil {
		t.Errorf("DeleteAddress returned something, but should be nil")
	}

}
