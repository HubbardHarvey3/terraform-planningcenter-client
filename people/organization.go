package people

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func GetOrganization(client *core.PC_Client) (core.OrganizationRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "people/v2"
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.OrganizationRoot{}, err
	}

	//Convert from json to the struct
	var jsonBody core.OrganizationRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.OrganizationRoot{}, fmt.Errorf("Error unmarshalling during GetOrganization ::: %v\n", err)
	}

	return jsonBody, nil
}
