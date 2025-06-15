package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get an campus object.

Endpoint = /people/v2/campuses/<campus ID>
*/
func GetCampus(client *core.PC_Client, campusId string) (core.CampusRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/campuses/" + campusId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.CampusRoot{}, fmt.Errorf("Error creating Getcampus request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.CampusRoot{}, fmt.Errorf("Error executing Getcampus request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.CampusRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.CampusRoot{}, fmt.Errorf("Error unmarshalling during Getcampus ::: %v\n", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create an campus object.  On create, the campus is automatically
assigned to the person ID listed in the endpoint

Assignable Attributes
  - Latitude
  - Longitude
  - Description
  - Street
  - City
  - State
  - Zip
  - Country
  - Phone_number
  - Website
  - Twenty_four_hour_time
  - Date_format
  - Church_center_enabled
  - Contact_email_address
  - Time_zone
  - Geolocation_set_manually
  - Name

Endpoint = /people/v2/campuses
*/
func CreateCampus(client *core.PC_Client, responseData *core.CampusRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/campuses"

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
Delete HTTP Method to remove an campus.

Endpoint = /people/v2/campuses/<campus ID>
*/
func DeleteCampus(client *core.PC_Client, campusId string) error {
	endpoint := client.Endpoint + "/people/v2/campuses/" + campusId

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
PATCH HTTP Method to update the campus

Assignable Attributes
  - Latitude
  - Longitude
  - Description
  - Street
  - City
  - State
  - Zip
  - Country
  - Phone_number
  - Website
  - Twenty_four_hour_time
  - Date_format
  - Church_center_enabled
  - Contact_email_address
  - Time_zone
  - Geolocation_set_manually
  - Name

Endpoint = /people/v2/campuses/<campus ID>
*/
func UpdateCampus(client *core.PC_Client, campusId string, responseData *core.CampusRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/campus/" + campusId

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
