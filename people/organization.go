package people

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get an Organization object.

Endpoint = /people/v2/
*/
func GetOrganization(client *core.PC_Client) (core.OrganizationRootNoRelationship, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2"
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.OrganizationRootNoRelationship{}, fmt.Errorf("Error creating request: %v\n", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.OrganizationRootNoRelationship{}, err
	}

	//Convert from json to the struct
	var jsonBody core.OrganizationRootNoRelationship
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.OrganizationRootNoRelationship{}, fmt.Errorf("Error unmarshalling during GetOrganization ::: %v\n", err)
	}

	return jsonBody, nil
}

/*
GET HTTP Method to get all addresses attached to an Organization

Endpoint = /people/v2/addresses
*/
func GetOrganizationAddress(client *core.PC_Client) (core.OrganizationRootAddress, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/addresses"
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.OrganizationRootAddress{}, fmt.Errorf("Error creating request: %v\n", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.OrganizationRootAddress{}, fmt.Errorf("Error executing the request: %v\n", err)
	}

	//Convert from json to the struct
	var jsonBody core.OrganizationRootAddress
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.OrganizationRootAddress{}, fmt.Errorf("Error unmarshalling during GetOrganization ::: %v\n", err)
	}

	return jsonBody, nil
}
