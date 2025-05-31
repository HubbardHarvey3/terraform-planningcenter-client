package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestCreatePerson(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "person",
			"id": "12345",
			"attributes": {
				"accounting_administrator": false,
				"anniversary": null,
				"birthdate": "1990-01-01",
				"first_name": "UnitTest",
				"gender": "male",
				"given_name": null,
				"grade": null,
				"graduation_year": null,
				"inactivated_at": null,
				"last_name": "Doe",
				"medical_notes": null,
				"membership": "member",
				"middle_name": null,
				"nickname": null,
				"site_administrator": false,
				"status": "active"
			}
		}
	}`

	mockServer := setupMockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	var data core.PeopleRoot

	//Convert json into PeopleRoot
	json.Unmarshal([]byte(mockResponse), &data)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	person, err := CreatePeople(client, &data)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var response core.PeopleRoot
	json.Unmarshal(person, &response)

	if response.Data.Attributes.FirstName != "UnitTest" {
		t.Errorf("Expected person.Data.Attributes.FirstName to be UnitTest, instead got %v\n", response.Data.Attributes.FirstName)
	}

}

func TestGetPerson(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "person",
			"id": "12345",
			"attributes": {
				"accounting_administrator": false,
				"anniversary": null,
				"birthdate": "1990-01-01",
				"first_name": "UnitTest",
				"gender": "male",
				"given_name": null,
				"grade": null,
				"graduation_year": null,
				"inactivated_at": null,
				"last_name": "Doe",
				"medical_notes": null,
				"membership": "member",
				"middle_name": null,
				"nickname": null,
				"site_administrator": false,
				"status": "active"
			}
		}
	}`

	var mockServer *httptest.Server = setupMockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	person, err := GetPerson(client, "12345")
	if err != nil {
		t.Errorf("GetPerson failed with an error ::: %v\n", err)
	}

	if person.Data.Attributes.FirstName != "UnitTest" {
		t.Errorf("Expected person.Data.ATtributes.FirstName to be UnitTest, instead got %v\n", person.Data.Attributes.FirstName)
	}

	if person.Data.Attributes.Birthdate != "1990-01-01" {
		t.Errorf("Expected person.Data.Attributes.Birthdate to be 1990-01-01, instead got %v\n", person.Data.Attributes.Birthdate)
	}
}

func TestUpdatePerson(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "person",
			"id": "12345",
			"attributes": {
				"accounting_administrator": false,
				"anniversary": null,
				"birthdate": "1990-01-01",
				"first_name": "UpdateName",
				"gender": "male",
				"given_name": null,
				"grade": null,
				"graduation_year": null,
				"inactivated_at": null,
				"last_name": "Doe",
				"medical_notes": null,
				"membership": "member",
				"middle_name": null,
				"nickname": null,
				"site_administrator": false,
				"status": "active"
			}
		}
	}`
	var person core.PeopleRoot

	var mockServer *httptest.Server = setupMockServer(mockResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response, err := UpdatePerson(client, "12345678", &person)
	if err != nil {
		t.Errorf("GetPerson failed with an error ::: %v\n", err)
	}

	json.Unmarshal(response, &person)

	if person.Data.Attributes.FirstName != "UpdateName" {
		t.Errorf("First Name is not 'UpdateName', but is showing as : %v", person.Data.Attributes.FirstName)
	}

}

func TestDeletePerson(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	var response = DeletePerson(client, "12345")

	if response != nil {
		t.Errorf("DeletePerson returned something, but should be nil")
	}

}
