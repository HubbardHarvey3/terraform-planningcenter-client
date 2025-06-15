package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestCreateEmail(t *testing.T) {
	var emailResponse = `{
		"data": {
			"type": "Email",
			"attributes": {
				"address": "john.doe@gmail.com",
				"location": "Home",
				"primary": true
			}
		} 
	}`
	var mockServer *httptest.Server = setupMockServer(emailResponse, http.StatusOK)
	var dataEmail core.EmailRoot

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	err := json.Unmarshal([]byte(emailResponse), &dataEmail)
	if err != nil {
		t.Error(err)
	}

	emailBytes, err := CreateEmail(client, "123456789", &dataEmail)
	if err != nil {
		t.Error(err)
	}

	var email core.EmailRoot
	err = json.Unmarshal(emailBytes, &email)
	if err != nil {
		t.Error(err)
	}

	if email.Data.Attributes.Address != "john.doe@gmail.com" {
		t.Errorf("Address is not john.doe@gmail.com, but is showing as : %v", email.Data.Attributes.Address)
	}

}

func TestGetEmail(t *testing.T) {
	var emailResponse = `{
		"data": {
			"type": "Email",
			"id": "12345678",
			"attributes": {
			"address": "random@gmail.com",
			"blocked": false,
			"created_at": "2023-11-25T15:33:03Z",
			"location": "Work",
			"primary": true,
			"updated_at": "2023-11-25T15:33:03Z"
			},
			"relationships": {
			"person": {
				"data": {
				"type": "Person",
				"id": "123456789"
				}
			}
			},
			"links": {
			"person": "https://api.planningcenteronline.com/people/v2/people/138378248",
			"self": "https://api.planningcenteronline.com/people/v2/emails/89020410"
			}
		},
		"included": [],
		"meta": {
			"parent": {
			"id": "458241",
			"type": "Organization"
			}
		}
	}`
	var email core.EmailRoot

	var mockServer = setupMockServer(emailResponse, http.StatusOK)
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	email, err := GetEmail(client, "12345678")
	if err != nil {
		t.Errorf("GetPerson failed with an error ::: %v\n", err)
	}

	if email.Data.Attributes.Address != "random@gmail.com" {
		t.Errorf("Address is not random@gmail.com, but is showing as : %v", email.Data.Attributes.Address)
	}
}

func TestUpdateEmail(t *testing.T) {
	var emailResponse = `{
		"data": {
			"type": "Email",
			"id": "12345678",
			"attributes": {
			"address": "john.doe.updated@gmail.com",
			"blocked": false,
			"created_at": "2023-11-25T15:33:03Z",
			"location": "Work",
			"primary": true,
			"updated_at": "2023-11-25T15:33:03Z"
			},
			"relationships": {
			"person": {
				"data": {
				"type": "Person",
				"id": "123456789"
				}
			}
			},
			"links": {
			"person": "https://api.planningcenteronline.com/people/v2/people/138378248",
			"self": "https://api.planningcenteronline.com/people/v2/emails/89020410"
			}
		},
		"included": [],
		"meta": {
			"parent": {
			"id": "458241",
			"type": "Organization"
			}
		}
	}`
	var email core.EmailRoot

	var mockServer = setupMockServer(emailResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response, err := UpdateEmail(client, "12345678", &email)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(response, &email)
	if err != nil {
		t.Error(err)
	}

	if email.Data.Attributes.Address != "john.doe.updated@gmail.com" {
		t.Errorf("email is not 'john.doe.updated@gmail.com', but is showing as : %v", email.Data.Attributes.Address)
	}

}

func TestDeleteEmail(t *testing.T) {

	var mockServer = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeleteEmail(client, "12345678")

	if response != nil {
		t.Errorf("DeletePerson returned something, but should be nil")
	}

}
