package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var mockTabResponse string = `{
	"data": {
		"type": "Tab",
		"id": "12345",
		"attributes": {
			"name": "Manual Test",
			"sequence": 111,
			"slug": "ManualTestSlug"
		}
	}
}`

func TestCreateTab(t *testing.T) {

	mockServer := setupMockServer(mockTabResponse, http.StatusOK)
	defer mockServer.Close()

	var data core.SingleTabRoot

	err := json.Unmarshal([]byte(mockTabResponse), &data)
	if err != nil {
		t.Errorf("Error during Unmarshalling in TestCreateTab, %v", err)
	}

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	tab, err := CreateTab(client, &data)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var response core.SingleTabRoot
	json.Unmarshal(tab, &response)

	if response.Data.Attributes.Name != "Manual Test" {
		t.Errorf("Expected Tab.Data.Attributes.Name to be 'Manual Test', instead got %v\n", response.Data.Attributes.Name)
	}

}

func TestGetTab(t *testing.T) {
	var mockServer *httptest.Server = setupMockServer(mockTabResponse, http.StatusOK)
	defer mockServer.Close()

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	tab, err := GetTab(client, "12345")
	if err != nil {
		t.Errorf("GetTab failed with an error ::: %v\n", err)
	}

	if tab.Data.Attributes.Name != "Manual Test" {
		t.Errorf("Expected Tab.Data.Attributes.Name to be Manual Test, instead got %v\n", tab.Data.Attributes.Name)
	}

	if tab.Data.Attributes.Name != "Manual Test" {
		t.Errorf("Expected Tab.Data.Attributes.Name to be Manual Test, instead got %v\n", tab.Data.Attributes.Name)
	}
}

func TestUpdateTab(t *testing.T) {
	var tab core.SingleTabRoot

	var mockServer *httptest.Server = setupMockServer(mockTabResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response, err := UpdateTab(client, "12345678", &tab)
	if err != nil {
		t.Errorf("UpdateTab failed with an error ::: %v\n", err)
	}

	json.Unmarshal(response, &tab)

	if tab.Data.Attributes.Name != "Manual Test" {
		t.Errorf("First Name is not 'Manual Test', but is showing as : %v", tab.Data.Attributes.Name)
	}

}

func TestDeleteTab(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	var response = DeleteTab(client, "12345")

	if response != nil {
		t.Errorf("DeleteTab returned something, but should be nil")
	}

}
