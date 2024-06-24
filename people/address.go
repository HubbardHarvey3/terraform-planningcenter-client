package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get an address object.

Endpoint = /people/v2/addresses/<address ID>
*/
func GetAddress(client *core.PC_Client, addressId string) (core.AddressRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/addresses/" + addressId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.AddressRoot{}, fmt.Errorf("Error creating GetAddress request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.AddressRoot{}, fmt.Errorf("Error executing GetAddress request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.AddressRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.AddressRoot{}, fmt.Errorf("Error unmarshalling during GetAddress ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create an address object.  On create, the address is automatically
assigned to the person ID listed in the endpoint

Assignable Attributes
  - city
  - state
  - zip
  - country_cde
  - location
  - primary
  - street_line_1
  - street_line_2

Endpoint = /people/v2/people/<people ID>/addresses
*/
func CreateAddress(client *core.PC_Client, peopleId string, responseData *core.AddressRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId + "/addresses"

	// Make relationships nil so it isn't in the API payload
	responseData.Data.Relationships = nil

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %w", err)
	}

	// Create a request with the JSON data
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	// Set the content type to application/json
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	body, err := client.DoRequest(request)
	if err != nil {
		return nil, fmt.Errorf("Error executing create request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove an address.

Endpoint = /people/v2/addresses/<address ID>
*/
func DeleteAddress(client *core.PC_Client, addressId string) error {
	endpoint := client.Endpoint + "people/v2/addresses/" + addressId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Error creating request ::: %v", err)
	}

	_, err = client.DoRequest(request)
	if err != nil {
		return fmt.Errorf("Error executing delete request: %w", err)
	}

	return nil
}

/*
PATCH HTTP Method to update the address for a person.

Assignable Attributes
  - city
  - state
  - zip
  - country_cde
  - location
  - primary
  - street_line_1
  - street_line_2

Endpoint = /people/v2/addresses/<address ID>
*/
func UpdateAddress(client *core.PC_Client, addressId string, responseData *core.AddressRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/addresses/" + addressId

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %w", err)
	}

	// Create a request with the JSON data
	request, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	// Set the content type to application/json
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	body, err := client.DoRequest(request)
	if err != nil {
		return nil, fmt.Errorf("Error executing update request: %w", err)
	}

	return body, nil

}
