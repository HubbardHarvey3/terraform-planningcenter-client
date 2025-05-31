package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var campusResponse = `{
	"data": {
		"type": "Campus",
		"id": "94987",
		"attributes": {
			"avatar_url": null,
			"church_center_enabled": true,
			"city": "Mac",
			"contact_email_address": null,
			"country": "US",
			"created_at": "2024-06-22T17:15:33Z",
			"date_format": null,
			"description": null,
			"geolocation_set_manually": false,
			"latitude": "30.3714823",
			"longitude": "-82.3234843",
			"name": "CBC Campus 1",
			"phone_number": null,
			"state": "FL",
			"street": "555111 Campus Dr",
			"time_zone": "America/New_York",
			"twenty_four_hour_time": null,
			"updated_at": "2024-06-22T17:15:33Z",
			"website": null,
			"zip": "32087"
		}
	}
}`
var campusId = "94987"

func TestCreateCampus(t *testing.T) {
	var dataCampus core.CampusRoot

	var mockServer *httptest.Server = setupMockServer(campusResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	err := json.Unmarshal([]byte(campusResponse), &dataCampus)
	if err != nil {
		t.Error(err)
	}

	campusBytes, err := CreateCampus(client, &dataCampus)
	if err != nil {
		t.Errorf("CreateCampus failed: %v", err)
	}

	var campus core.CampusRoot
	json.Unmarshal(campusBytes, &campus)

	if campus.Data.Attributes.Name != "CBC Campus 1" {
		t.Errorf("Name is not 'CBC Campus 1', but is showing as : %v", campus.Data.Attributes.Name)
	}

}

func TestGetCampus(t *testing.T) {
	var campus core.CampusRoot

	var mockServer *httptest.Server = setupMockServer(campusResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	campus, err := GetCampus(client, "94987")
	if err != nil {
		t.Errorf("GetCampus failed with an error ::: %v\n", err)
	}

	if campus.Data.Attributes.City != "Mac" {
		t.Errorf("Campus is not 'Mac', but is showing as : %v", campus.Data.Attributes.City)
	}
}

func TestUpdateCampus(t *testing.T) {
	var campus core.CampusRoot

	var mockServer *httptest.Server = setupMockServer(campusResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	campus, err := GetCampus(client, campusId)
	if err != nil {
		t.Errorf("GetCampus failed with an error ::: %v\n", err)
	}

	var updatedCampus core.CampusRoot
	updatedCampus.Data.Attributes = campus.Data.Attributes

	response, err := UpdateCampus(client, campusId, &updatedCampus)

	json.Unmarshal(response, &updatedCampus)

	if updatedCampus.Data.Attributes.City != "Mac" {
		t.Errorf("Campus is not 'Mac', but is showing as : %v", updatedCampus.Data.Attributes.City)
	}

}

func TestDeleteCampus(t *testing.T) {
	var mockServer *httptest.Server = setupMockServer(campusResponse, http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeleteCampus(client, campusId)

	if response != nil {
		t.Errorf("DeleteCampus returned something, but should have returned nil")
	}

}
