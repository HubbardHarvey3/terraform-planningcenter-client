package people

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestGetForm(t *testing.T) {
	var responseForms string = `{
		"data": {
			"type": "Form",
			"id": "771246",
			"attributes": {
				"active": true,
				"archived": false,
				"archived_at": null,
				"created_at": "2024-07-01T11:09:20Z",
				"deleted_at": null,
				"description": "<div>A description for the test form</div>",
				"login_required": false,
				"name": "Test Form",
				"public_url": "https://cbc-458241.churchcenter.com/people/forms/771246",
				"recently_viewed": false,
				"send_submission_notification_to_submitter": false,
				"submission_count": 0,
				"updated_at": "2024-07-01T11:09:37Z"
			}
		}
	}`

	var form core.FormRoot
	var mockServer *httptest.Server = setupMockServer(responseForms, http.StatusOK)

	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	form, err := GetForm(client, "771246")
	if err != nil {
		t.Errorf("GetForm failed with an error ::: %v\n", err)
	}

	if form.Data.Attributes.Name != "Test Form" {
		t.Errorf("Form Name is not 'Test Form', but is showing as : %v", form.Data.Attributes.Name)
	}

}
