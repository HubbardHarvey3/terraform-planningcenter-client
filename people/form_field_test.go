package people

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestGetFormField(t *testing.T) {
	var formFieldResponse string = `{
		"data": {
			"type": "FormField",
			"id": "6003884",
			"attributes": {
				"created_at": "2024-07-01T11:10:11Z",
				"description": null,
				"field_type": "string",
				"label": "Phone number",
				"required": true,
				"sequence": 1,
				"settings": {},
				"updated_at": "2024-07-01T11:10:11Z"
			}
		}
	}`
	var form core.FormFieldsRoot
	var mockServer *httptest.Server = setupMockServer(formFieldResponse, http.StatusOK)

	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	form, err := GetFormField(client, "873385", "6003884")
	if err != nil {
		t.Errorf("GetFormField failed with an error ::: %v\n", err)
	}

	if form.Data.Attributes.Label != "Phone number" {
		t.Errorf("Field ID 6830109 is not 'Phone number', but is showing as : %v", form.Data.Attributes.Label)
	}
}
