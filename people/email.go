package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func GetEmail(client *core.PC_Client, emailId string) (core.EmailRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/emails/" + emailId
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

func CreateEmail(client *core.PC_Client, peopleId string, responseData *core.EmailRootNoRelationship) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId + "/emails"

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
		return nil, fmt.Errorf("DoRequest error during CreateEmail : %w\n", err)
	}

	return body, nil
}

func DeleteEmail(client *core.PC_Client, emailId string) error {
	endpoint := client.Endpoint + "people/v2/emails/" + emailId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("NewRequest error in DeleteEmail : %v\n", err)
	}

	// Make the request
	body, err := client.DoRequest(request)
	if err != nil {
		return fmt.Errorf("DoRequest error during DeleteEmail: %v\n", err)
	}

	fmt.Println(string(body))
	return nil

}

func UpdateEmail(client *core.PC_Client, emailId string, responseData *core.EmailRootNoRelationship) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/emails/" + emailId

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	// Create a request with the JSON data
	request, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
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
