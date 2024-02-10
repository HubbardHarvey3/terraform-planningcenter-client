package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEmail(client *PC_Client, appID, secretToken, emailID string) EmailRoot {
	//Fetch the data
	endpoint := HostURL + "people/v2/emails/" + emailID
	request, err := http.NewRequest("GET", endpoint, nil)

	// Make the request
	body, err := client.doRequest(request, secretToken, appID)
	if err != nil {
		fmt.Println("Failure during doRequest: ")
		fmt.Print(err)
	}

	var jsonBody EmailRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		fmt.Print(err)
	}

	return jsonBody

}

func CreateEmail(client *PC_Client, appID, secretToken, peopleID string, responseData *EmailRootNoRelationship) []byte {
	endpoint := HostURL + "people/v2/people/" + peopleID + "/emails"

	// Convert struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	// Create a request with the JSON data
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the content type to application/json
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	body, err := client.doRequest(request, secretToken, appID)
	if err != nil {
		fmt.Println("Failure during doRequest: ")
		fmt.Print(err)
	}

	return body
}

func DeleteEmail(client *PC_Client, appID, secretToken, emailID string) {
	endpoint := HostURL + "people/v2/emails/" + emailID

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Make the request
	body, err := client.doRequest(request, secretToken, appID)
	if err != nil {
		fmt.Println("Failure during doRequest: ")
		fmt.Print(err)
	}

	fmt.Println(string(body))

}

func UpdateEmail(client *PC_Client, appID, secretToken, emailID string, responseData *EmailRoot) []byte {
	endpoint := HostURL + "people/v2/emails/" + emailID

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
	body, err := client.doRequest(request, secretToken, appID)
	if err != nil {
		fmt.Println("Failure during doRequest: ")
		fmt.Print(err)
	}

	return body

}
