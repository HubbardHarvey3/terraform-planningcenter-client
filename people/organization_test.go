package people

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestGetOrganization(t *testing.T) {
	var orgResponse string = `{
		"data": {
		  "type": "Organization",
		  "id": "1",
		  "attributes": {
			"name": "CBC",
			"country_code": "US",
			"date_format": 1,
			"time_zone": "EST",
			"contact_website": "cbc.example.com",
			"created_at": "2000-01-01T12:00:00Z",
			"avatar_url": "string"
		  },
		  "relationships": {}
		}
	}`
	var mockServer *httptest.Server = setupMockServer(orgResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	org, err := GetOrganization(client)
	if err != nil {
		t.Errorf("GetPerson failed with an error ::: %v\n", err)
	}

	if org.Data.Attributes.Name != "CBC" {
		t.Errorf("Expected org.Data.Attributes.Name to be CBC, instead got %v\n", org.Data.Attributes.Name)
	}

}

func TestGetOrganizationAddress(t *testing.T) {
	var orgResponse string = `{
		"data": [{
		  "type": "Address",
		  "id": "1",
		  "attributes": {
			"city": "Hometown",
			"state": "FL",
			"zip": "333333",
			"country_code": "US",
			"location": "Here",
			"primary": true,
			"street_line_1": "1111 NW 41st BLVD",
			"street_line_2": "",
			"created_at": "2000-01-01T12:00:00Z",
			"updated_at": "2000-01-01T12:00:00Z",
			"country_name": "United States"
		  },
		  "relationships": {
			"person": {
			  "data": {
				"type": "Person",
				"id": "1"
			  }
			}
		  }
		}]
	}`

	var mockServer *httptest.Server = setupMockServer(orgResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	org, err := GetOrganizationAddress(client)
	if err != nil {
		t.Errorf("GetPerson failed with an error ::: %v\n", err)
	}

	if org.Data[0].Attributes.City != "Hometown" {
		t.Errorf("Expected org.Data[0].Attributes.City to be Hometown, instead got %v\n", org.Data[0].Attributes.City)
	}
}
