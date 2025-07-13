package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get a list category object.

Endpoint = /people/v2/list_categories/<List Category ID>
*/
func GetListCategory(client *core.PC_Client, listCategoryId string) (core.ListCategoryRoot, error) {
	endpoint := client.Endpoint + "/people/v2/list_categories/" + listCategoryId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.ListCategoryRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.ListCategoryRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.ListCategoryRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.ListCategoryRoot{}, fmt.Errorf("Error unmarshalling during GetListCategory ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a person object.

Assignable Attributes
  - name

Endpoint = /people/v2/list_categories/
*/
func CreateListCategory(client *core.PC_Client, responseData *core.ListCategoryRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/list_categories/"

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
		return nil, fmt.Errorf("Error executing create list category request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove a person.

Endpoint = /people/v2/people/<person ID>
*/
func DeleteListCategory(client *core.PC_Client, listCategoryId string) error {
	endpoint := client.Endpoint + "/people/v2/list_categories/" + listCategoryId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Error creating request ::: %v", err)
	}

	_, err = client.DoRequest(request)
	if err != nil {
		return fmt.Errorf("Error executing delete list category request: %w", err)
	}

	return nil
}

/*
PATCH HTTP Method to update a person object.

Assignable Attributes
  - name

Endpoint = /people/v2/people/<person id>
*/
func UpdateListCategory(client *core.PC_Client, listCategoryId string, responseData *core.ListCategoryRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/list_categories/" + listCategoryId

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
		return nil, fmt.Errorf("Error executing update list category request: %w", err)
	}

	return body, nil

}
