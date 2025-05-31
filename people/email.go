package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get an email object.

Endpoint = /people/v2/email/<email ID>
*/
func GetEmail(client *core.PC_Client, emailId string) (core.EmailRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/emails/" + emailId
	request, err := http.NewRequest("GET", endpoint, nil)

	// Make the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.EmailRoot{}, fmt.Errorf("DoRequest Error during GetEmail : %v\n", err)
	}

	var jsonBody core.EmailRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.EmailRoot{}, fmt.Errorf("Error in GetEmail unmarshalling : %w ", err)
	}

	return jsonBody, nil

}

/*
POST HTTP Method to create an email object.  On create, the email is automatically
assigned to the person ID listed in the endpoint.  If you include an empty Relationship
field, the api will return a 422 error.

Assignable Attributes
  - address
  - location
  - primary

Endpoint = /people/v2/people/<people ID>/emails
*/
func CreateEmail(client *core.PC_Client, peopleId string, responseData *core.EmailRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/people/" + peopleId + "/emails"

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
		return nil, fmt.Errorf("DoRequest error during CreateEmail : %w\n", err)
	}

	return body, nil
}

/*
Delete HTTP Method to remove an email.

Endpoint = /people/v2/emails/<email ID>
*/
func DeleteEmail(client *core.PC_Client, emailId string) error {
	endpoint := client.Endpoint + "/people/v2/emails/" + emailId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("NewRequest error in DeleteEmail : %v\n", err)
	}

	_, err = client.DoRequest(request)
	if err != nil {
		return err
	}

	return nil

}

/*
PATCH HTTP Method to update the email for a person.

Assignable Attributes
  - email
  - location
  - primary

Endpoint = /people/v2/emails/<email ID>
*/
func UpdateEmail(client *core.PC_Client, emailId string, responseData *core.EmailRoot) ([]byte, error) {
	endpoint := client.Endpoint + "/people/v2/emails/" + emailId

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
		return nil, fmt.Errorf("DoRequest error during UpdateEmail : %v\n", err)
	}

	return body, nil

}
