package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestCreateListCategory(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "ListCategory",
			"id": "123456789",
			"attributes": {
				"name": "Mock List"
			}
		}
	}`

	mockServer := setupMockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	var data core.ListCategoryRoot

	//Convert json into ListCategoryRoot
	json.Unmarshal([]byte(mockResponse), &data)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	listCategory, err := CreateListCategory(client, &data)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var response core.ListCategoryRoot
	json.Unmarshal(listCategory, &response)

	if response.Data.Attributes.Name != "Mock List" {
		t.Errorf("Expected listCategory.Data.Attributes.Name to be Mock List, instead got %v\n", response.Data.Attributes.Name)
	}

}

func TestGetListCategory(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "ListCategory",
			"id": "123456789",
			"attributes": {
				"name": "Mock List"
			}
		}
	}`

	var mockServer *httptest.Server = setupMockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	listCategory, err := GetListCategory(client, "123456789")
	if err != nil {
		t.Errorf("GetListCategory failed with an error ::: %v\n", err)
	}

	if listCategory.Data.Attributes.Name != "Mock List" {
		t.Errorf("Expected listCategory.Data.Attributes.Name to be 'Mock List', instead got %v\n", listCategory.Data.Attributes.Name)
	}

}

func TestUpdateListCategory(t *testing.T) {
	mockResponse := `{
		"data": {
			"type": "ListCategory",
			"id": "123456789",
			"attributes": {
				"name": "Updated List Name"
			}
		}
	}`
	var listCategory core.ListCategoryRoot

	var mockServer *httptest.Server = setupMockServer(mockResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response, err := UpdateListCategory(client, "123456789", &listCategory)
	if err != nil {
		t.Errorf("GetlistCategory failed with an error ::: %v\n", err)
	}

	json.Unmarshal(response, &listCategory)

	if listCategory.Data.Attributes.Name != "Updated List Name" {
		t.Errorf("First Name is not 'Updated List Name', but is showing as : %v", listCategory.Data.Attributes.Name)
	}

}

func TestDeletelistCategory(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	var response = DeleteListCategory(client, "123456789")

	if response != nil {
		t.Errorf("DeletelistCategory returned something, but should be nil")
	}

}
