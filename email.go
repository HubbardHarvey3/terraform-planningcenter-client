package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEmail(client *PC_Client, appId, secretToken, emailId string) (EmailRoot, error) {
	//Fetch the data
	endpoint := HostURL + "people/v2/emails/" + emailId
	request, err := http.NewRequest("GET", endpoint, nil)

	// Make the request
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return EmailRoot{}, fmt.Errorf("doRequest Error during GetEmail : %v\n", err)
	}

	var jsonBody EmailRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Print(err)
	}

	return jsonBody, nil

}

func CreateEmail(client *PC_Client, appId, secretToken, peopleId string, responseData *EmailRootNoRelationship) ([]byte, error) {
	endpoint := HostURL + "people/v2/people/" + peopleId + "/emails"

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Errorf("Error marshalling JSON: %w", err)
	}

	// Create a request with the JSON data
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Errorf("Error creating request: %w", err)
	}

	// Set the content type to application/json
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return nil, fmt.Errorf("doRequest error during CreateEmail : %w\n", err)
	}

	return body, nil
}

func DeleteEmail(client *PC_Client, appId, secretToken, emailId string) error {
	endpoint := HostURL + "people/v2/emails/" + emailId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("NewRequest error in DeleteEmail : %v\n", err)
	}

	// Make the request
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return fmt.Errorf("doRequest error during DeleteEmail: %v\n", err)
	}

	fmt.Println(string(body))
	return nil

}

func UpdateEmail(client *PC_Client, appId, secretToken, emailId string, responseData *EmailRoot) ([]byte, error) {
	endpoint := HostURL + "people/v2/emails/" + emailId

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
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return nil, fmt.Errorf("doRequest error during UpdateEmail : %v\n", err)
	}

	return body, nil

}
