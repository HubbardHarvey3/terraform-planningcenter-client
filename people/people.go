package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get a person object.

Endpoint = /people/v2/people/<person ID>
*/
func GetPeople(client *core.PC_Client, peopleId string) (core.PeopleRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/people/" + peopleId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.PeopleRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.PeopleRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.PeopleRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.PeopleRoot{}, fmt.Errorf("Error unmarshalling during GetPeople ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a person object.

Assignable Attributes
  - given_name
  - first_name
  - nickname
  - middle_name
  - last_name
  - birthdate
  - anniversary
  - grade
  - child
  - graduation_year
  - site_administrator
  - accounting_administrator
  - people_permissions
  - gender
  - membership
  - inactivated_at
  - status
  - medical_notes
  - avatar
  - primary_campus_id
  - gender_id
  - remote_id

Endpoint = /people/v2/people/
*/
func CreatePeople(client *core.PC_Client, responseData *core.PeopleRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/"

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
		return nil, fmt.Errorf("Error executing create people request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove a person.

Endpoint = /people/v2/people/<person ID>
*/
func DeletePeople(client *core.PC_Client, peopleId string) error {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Error creating request ::: %v", err)
	}

	_, err = client.DoRequest(request)
	if err != nil {
		return fmt.Errorf("Error executing delete people request: %w", err)
	}

	return nil
}

/*
PATCH HTTP Method to update a person object.

Assignable Attributes
  - given_name
  - first_name
  - nickname
  - middle_name
  - last_name
  - birthdate
  - anniversary
  - grade
  - child
  - graduation_year
  - site_administrator
  - accounting_administrator
  - people_permissions
  - gender
  - membership
  - inactivated_at
  - status
  - medical_notes
  - avatar
  - primary_campus_id
  - gender_id
  - remote_id

Endpoint = /people/v2/people/<person id>
*/
func UpdatePeople(client *core.PC_Client, peopleId string, responseData *core.PeopleRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId

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
		return nil, fmt.Errorf("Error executing update people request: %w", err)
	}

	return body, nil

}
