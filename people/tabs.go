package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get a tab object.

Endpoint = /people/v2/tabs/<tab ID>
*/
func GetTab(client *core.PC_Client, tabId string) (core.SingleTabRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/tabs/" + tabId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.SingleTabRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.SingleTabRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.SingleTabRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.SingleTabRoot{}, fmt.Errorf("Error unmarshalling during GetPerson ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a tab object.

Assignable Attributes
  - name
  - sequence

Endpoint = /people/v2/tabs/
*/
func CreateTab(client *core.PC_Client, responseData *core.SingleTabRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/tabs/"

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
		return nil, fmt.Errorf("Error executing create people request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove a person.

Endpoint = /people/v2/people/<person ID>
*/
func DeleteTab(client *core.PC_Client, tabId string) error {
	endpoint := client.Endpoint + "/people/v2/tabs/" + tabId

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
  - name
  - sequence

Endpoint = /people/v2/tabs/<tab id>
*/
func UpdateTab(client *core.PC_Client, tabId string, responseData *core.SingleTabRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/people/" + tabId

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
