package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get a note object.

Endpoint = /people/v2/note_categories
*/
func GetNoteCategory(client *core.PC_Client, noteCategoryId string) (core.NoteCategoryRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/note_categories/" + noteCategoryId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.NoteCategoryRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.NoteCategoryRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.NoteCategoryRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.NoteCategoryRoot{}, fmt.Errorf("Error unmarshalling during GetNote ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a note object.

Assignable Attributes
  - name

Endpoint = /people/v2/note_categories
*/
func CreateNoteCategory(client *core.PC_Client, responseData *core.NoteCategoryRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/note_categories"

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
		return nil, fmt.Errorf("Error executing create note category request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove a note.

Endpoint = /people/v2/note_categories/<note_categories ID>

Warning: Deleting a Note Category will also delete all associated notes
*/
func DeleteNoteCategory(client *core.PC_Client, noteCategoryId string) error {
	endpoint := client.Endpoint + "people/v2/note_categories/" + noteCategoryId

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

Endpoint = /people/v2/note_categories/<note_category id>
*/
func UpdateNoteCategory(client *core.PC_Client, noteCategoryId string, responseData *core.NoteCategoryRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/note_categories/" + noteCategoryId

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
