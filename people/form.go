package people

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

/*
GET HTTP Method to get a note object.  If no form id is provided,
the request will return all of the forms

Endpoint = /people/v2/forms/<form id>
*/
func GetForm(client *core.PC_Client, formId string) (core.FormRoot, error) {
	//Fetch the data
	endpoint := client.Endpoint + "/people/v2/forms/" + formId
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return core.FormRoot{}, fmt.Errorf("Error creating get request: %w", err)
	}

	//Send the request
	body, err := client.DoRequest(request)
	if err != nil {
		return core.FormRoot{}, fmt.Errorf("Error executing get request: %w", err)
	}

	//Convert from json to the struct
	var jsonBody core.FormRoot
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return core.FormRoot{}, fmt.Errorf("Error unmarshalling during GetNote ::: %v\n", err)
	}

	return jsonBody, nil

}
