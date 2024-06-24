package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get all phone numbers for an Organization.

Endpoint = /people/v2/phone_number/<phone number id>
*/
func GetPhoneNumber(client *core.PC_Client, phoneNumberId string) (core.PhoneNumberRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/phone_numbers/" + phoneNumberId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.PhoneNumberRoot{}, fmt.Errorf("Error creating request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.PhoneNumberRoot{}, err
	}

	//Convert from json to the struct
	var jsonBody core.PhoneNumberRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.PhoneNumberRoot{}, fmt.Errorf("Error unmarshalling during GetPhoneNumber ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a phone number object.

Assignable Attributes
  - number
  - carrier
  - location
  - primary

Endpoint = /people/v2/people/<person ID>/phone_numbers
*/
func CreatePhoneNumber(client *core.PC_Client, peopleId string, responseData *core.PhoneNumberRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId + "/phone_numbers"

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling JSON: %w", err)
	}

	// Make relationships nil so it isn't in the API payload
	responseData.Data.Relationships = nil

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
Delete HTTP Method to remove an phone number object.

- Requires the Phone Number ID you want deleted

Endpoint = /people/v2/phone_numbers/<Phone Number ID>
*/
func DeletePhoneNumber(client *core.PC_Client, phoneNumberId string) error {
	endpoint := client.Endpoint + "people/v2/phone_numbers/" + phoneNumberId

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
PATCH HTTP Method to update an phone number object.

Assignable Attributes
  - number
  - carrier
  - location
  - primary

Endpoint = /people/v2/phone_numbers/<Phone Number ID>
*/
func UpdatePhoneNumber(client *core.PC_Client, phoneNumberId string, responseData *core.PhoneNumberRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/phone_numbers/" + phoneNumberId

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
