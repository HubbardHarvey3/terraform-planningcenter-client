package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPeople(client *PC_Client, appId, secretToken, peopleId string) (PeopleRoot, error) {
	//Fetch the data
	endpoint := HostURL + "people/v2/people/" + peopleId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	//Send the request
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return PeopleRoot{}, err
	}

	//Convert from json to the struct
	var jsonBody PeopleRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return PeopleRoot{}, fmt.Errorf("Error unmarshalling during GetPeople ::: %v\n", err)
	}

	return jsonBody, nil

}

func CreatePeople(client *PC_Client, appId, secretToken string, responseData *PeopleRoot) ([]byte, error) {
	endpoint := HostURL + "people/v2/people/"

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
	body, err := client.doRequest(request, secretToken, appId)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DeletePeople(client *PC_Client, appId, secretToken, peopleId string) error {
	endpoint := HostURL + "people/v2/people/" + peopleId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Error creating request ::: %v", err)
	}

	_, err = client.doRequest(request, secretToken, appId)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePeople(client *PC_Client, appId, secretToken, peopleId string, responseData *PeopleRoot) ([]byte, error) {
	endpoint := HostURL + "people/v2/people/" + peopleId

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
	body, err := client.doRequest(request, secretToken, peopleId)
	if err != nil {
		return nil, err
	}

	return body, nil

}
