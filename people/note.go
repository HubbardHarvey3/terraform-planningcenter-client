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

Endpoint = /people/v2/people/<person ID>
*/
func GetNote(client *core.PC_Client, noteId string) (core.NoteRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/notes/" + noteId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.NoteRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.NoteRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.NoteRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.NoteRoot{}, fmt.Errorf("Error unmarshalling during GetNote ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create a note object.

Assignable Attributes
  - given_name
  - note
  - created_at
  - updated_at
  - display_date
  - note_category_id *Required

Endpoint = /people/v2/people/<person ID>/notes
*/
func CreateNote(client *core.PC_Client, personId string, responseData *core.NoteRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/people/" + personId + "/notes"

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
		return nil, fmt.Errorf("Error executing create note request: %w", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove a note.

Endpoint = /people/v2/notes/<notes ID>
*/
func DeleteNote(client *core.PC_Client, noteId string) error {
	endpoint := client.Endpoint + "/people/v2/notes/" + noteId

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
  - note
  - created_at
  - updated_at
  - display_date
  - note_category_id

Endpoint = /people/v2/notes/<note id>
*/
func UpdateNote(client *core.PC_Client, noteId string, responseData *core.NoteRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/notes/" + noteId

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
