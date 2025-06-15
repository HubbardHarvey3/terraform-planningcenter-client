package core

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const URL = "https://api.planningcenteronline.com/people/v2/people"

func setupMockServer(responseBody string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			if responseBody != "" {
				w.Write([]byte(responseBody))
			}
		}))
}

func TestPCClient_DoRequest(t *testing.T) {
	var responsePeople string = `{
		"data":{
		"type": "Person",
			"id": "138378248",
			"attributes": {
				"accounting_administrator": false,
				"anniversary": null,
				"birthdate": "1888-11-11",
				"can_create_forms": true,
				"can_email_lists": true,
				"child": false,
				"created_at": "2025-11-25T15:33:03Z",
				"directory_status": "no_access",
				"first_name": "Joe",
				"gender": "Male",
				"given_name": null,
				"grade": null,
				"graduation_year": null,
				"inactivated_at": null,
				"last_name": "H.",
				"login_identifier": "joe.blow@gmail.com",
				"medical_notes": null,
				"membership": "Member",
				"middle_name": "CURLING",
				"name": "Joe Blow",
				"passed_background_check": false,
				"people_permissions": "Manager",
				"remote_id": null,
				"school_type": null,
				"site_administrator": true,
				"status": "active",
				"updated_at": "2025-06-08T03:30:53Z"
			}
		}
	}`

	var mockServer *httptest.Server = setupMockServer(responsePeople, http.StatusOK)

	// Initialize your PC_Client with the mock server URL
	client := NewPCClient("MockAppID", "MockSecret", mockServer.URL)

	req, err := http.NewRequest(http.MethodGet, mockServer.URL, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	responseBody, err := client.DoRequest(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	// Verify the response
	if strings.Contains(string(responseBody), "404") {
		t.Errorf("Status Code is 404, not 200")
	}
}
