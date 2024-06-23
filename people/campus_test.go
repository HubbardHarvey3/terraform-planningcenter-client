package people

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonCampus = `{
	"data": {
		"type": "person",
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
			"last_name": "DoeCampus",
			"medical_notes": null,
			"membership": "member",
			"middle_name": null,
			"nickname": null,
			"site_administrator": false,
			"status": "active"
		}
	}
}`

var responseCampus = `{
	"data": {
		"type": "Campus",
		"attributes": {
			"latitude": "1.42",
			"longitude": "1.42",
			"description": "New Campus",
			"street": "221B Baker Street",
			"city": "Baker",
			"state": "Florida",
			"zip": "32088",
			"country": "US",
			"phone_number": "555-555-5555",
			"website": "https://www.unittest.org",
			"twenty_four_hour_time": null,
			"date_format": null,
			"church_center_enabled": true,
			"contact_email_address": "campusfaq@cbc.org",
			"geolocation_set_manually": false,
			"time_zone": "UTC",
			"name": "Test Campus"
		}
	}
}`

var personIdCampus string
var campusId string
var appIdCampus = os.Getenv("PC_APP_ID")
var secretTokenCampus = os.Getenv("PC_SECRET_TOKEN")

func TestCreateCampus(t *testing.T) {
	var dataCampus core.CampusRoot

	if appIdCampus == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenCampus == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	err := json.Unmarshal([]byte(responseCampus), &dataCampus)
	if err != nil {
		t.Error(err)
	}

	//Add Org ID to dataCampus
	dataCampus.Data.Relationships = nil

	campusBytes, err := CreateCampus(client, &dataCampus)
	if err != nil {
		t.Errorf("CreateCampus failed: %v", err)
	}

	var campus core.CampusRoot
	json.Unmarshal(campusBytes, &campus)
	campusId = campus.Data.ID

	if campus.Data.Attributes.Name != "Test Campus" {
		t.Errorf("Name is not 'Test Campus', but is showing as : %v", campus.Data.Attributes.Name)
	}

}

func TestGetCampus(t *testing.T) {
	var campus core.CampusRoot

	if appIdCampus == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenCampus == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdCampus, secretTokenCampus)

	campus, err := GetCampus(client, campusId)
	if err != nil {
		t.Errorf("GetCampus failed with an error ::: %v\n", err)
	}

	if campus.Data.Attributes.City != "Baker" {
		t.Errorf("Campus is not 'Baker', but is showing as : %v", campus.Data.Attributes.City)
	}
}

func TestUpdateCampus(t *testing.T) {
	var campus core.CampusRoot

	if appIdCampus == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenCampus == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdCampus, secretTokenCampus)

	campus, err := GetCampus(client, campusId)
	if err != nil {
		t.Errorf("GetCampus failed with an error ::: %v\n", err)
	}

	campus.Data.Attributes.City = "Updated"
	// Alter to without Relationships .... TODO Make this better
	var updatedCampus core.CampusRoot
	updatedCampus.Data.Attributes = campus.Data.Attributes

	response, err := UpdateCampus(client, campusId, &updatedCampus)

	json.Unmarshal(response, &updatedCampus)

	if updatedCampus.Data.Attributes.City != "Updated" {
		t.Errorf("Campus is not 'Updated', but is showing as : %v", updatedCampus.Data.Attributes.City)
	}

}

func TestDeleteCampus(t *testing.T) {
	if appIdCampus == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenCampus == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdCampus, secretTokenCampus)

	//err := DeleteCampus(client, "95005")
	err := DeleteCampus(client, campusId)
	if err != nil {
		t.Errorf("Error during DeleteCampus : %v\n", err)
	}

	_, err = GetCampus(client, campusId)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetCampus should be throwing a 404 after the person was deleted.  Error was %v", err)
	}

}
