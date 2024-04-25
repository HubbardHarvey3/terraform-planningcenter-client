package people

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func GetAddress(client *core.PC_Client, appId, secretToken, addressId string) (core.AddressRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2/addresses/" + addressId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	//Send the request
	body, err := client.DoRequest(request, secretToken, appId)
	if err != nil {
		return core.AddressRoot{}, err
	}

	//Convert from json to the struct
	var jsonBody core.AddressRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.AddressRoot{}, fmt.Errorf("Error unmarshalling during GetAddress ::: %v\n", err)
	}

	return jsonBody, nil

}

func CreateAddress(client *core.PC_Client, appId, secretToken, peopleId string, responseData *core.AddressRootNoRelationship) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId + "/addresses"

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
	body, err := client.DoRequest(request, secretToken, appId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func DeleteAddress(client *core.PC_Client, appId, secretToken, addressId string) error {
	endpoint := client.Endpoint + "people/v2/addresses/" + addressId

	// Create a request with the JSON data
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Error creating request ::: %v", err)
	}

	_, err = client.DoRequest(request, secretToken, appId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAddress(client *core.PC_Client, appId, secretToken, peopleId string, responseData *core.PeopleRoot) ([]byte, error) {
	endpoint := client.Endpoint + "people/v2/people/" + peopleId

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
	body, err := client.DoRequest(request, secretToken, peopleId)
	if err != nil {
		return nil, err
	}

	return body, nil

}
