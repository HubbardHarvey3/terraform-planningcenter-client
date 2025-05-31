package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

func TestCreateNoteCategory(t *testing.T) {
	var createNoteResponse string = `{
		"data": {
			"type": "NoteCategory",
			"id": "111222",
			"attributes": {
			"created_at": "2023-12-09T18:06:07Z",
			"locked": true,
			"name": "General",
			"organization_id": 123456,
			"updated_at": "2023-12-09T18:06:07Z"
			},
			"relationships": {
			"organization": {
				"data": {
				"type": "Organization",
				"id": "123456"
				}
			}
			},
			"links": {
				"shares": "https://api.planningcenteronline.com/people/v2/note_categories/221326/shares",
				"subscribers": "https://api.planningcenteronline.com/people/v2/note_categories/221326/subscribers",
				"subscriptions": "https://api.planningcenteronline.com/people/v2/note_categories/221326/subscriptions",
				"self": "https://api.planningcenteronline.com/people/v2/note_categories/221326"
			}
		},
		"included": [],
		"meta": {
			"can_include": [
				"shares",
				"subscribers",
				"subscriptions"
			],
			"parent": {
				"id": "458241",
				"type": "Organization"
			}
		}
	}`
	var mockServer *httptest.Server = setupMockServer(createNoteResponse, http.StatusOK)
	var dataNote core.NoteCategoryRoot

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	noteBytes, err := CreateNoteCategory(client, &dataNote)
	if err != nil {
		t.Error(err)
	}

	var note core.NoteCategoryRoot
	json.Unmarshal(noteBytes, &note)

	if note.Data.Attributes.Name != "General" {
		t.Errorf("Name is not 'General', but is showing as : %v", note.Data.Attributes.Name)
	}

}

func TestGetNoteCategory(t *testing.T) {
	var getNoteResponse string = `{
		"data": {
			"type": "NoteCategory",
			"id": "111222",
			"attributes": {
			"created_at": "2023-12-09T18:06:07Z",
			"locked": true,
			"name": "General",
			"organization_id": 123456,
			"updated_at": "2023-12-09T18:06:07Z"
			},
			"relationships": {
			"organization": {
				"data": {
				"type": "Organization",
				"id": "123456"
				}
			}
			},
			"links": {
				"shares": "https://api.planningcenteronline.com/people/v2/note_categories/221326/shares",
				"subscribers": "https://api.planningcenteronline.com/people/v2/note_categories/221326/subscribers",
				"subscriptions": "https://api.planningcenteronline.com/people/v2/note_categories/221326/subscriptions",
				"self": "https://api.planningcenteronline.com/people/v2/note_categories/221326"
			}
		},
		"included": [],
		"meta": {
			"can_include": [
				"shares",
				"subscribers",
				"subscriptions"
			],
			"parent": {
				"id": "458241",
				"type": "Organization"
			}
		}
	}`
	var note core.NoteCategoryRoot

	var mockServer *httptest.Server = setupMockServer(getNoteResponse, http.StatusOK)

	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	note, err := GetNoteCategory(client, "111222")
	if err != nil {
		t.Errorf("GetNote failed with an error ::: %v\n", err)
	}

	if note.Data.Attributes.Name != "General" {
		t.Errorf("Name is not 'General', but is showing as : %v", note.Data.Attributes.Name)
	}

}

func TestDeleteNoteCategory(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeleteNote(client, "111222")

	if response != nil {
		t.Errorf("DeletePerson returned something, but should be nil")
	}

}
