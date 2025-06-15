package people

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var noteResponse = `{
	"data": {
		"type": "Note",
		"id": "1",
		"attributes": {
			"note": "Test note from the api",
			"created_at": "2000-01-01T12:00:00Z",
			"updated_at": "2000-01-01T12:00:00Z",
			"display_date": "2000-01-01T12:00:00Z",
			"note_category_id": 11,
			"organization_id": "11111",
			"person_id": "12345678"
		}
	}
}`

func TestCreateNote(t *testing.T) {
	var dataNote core.NoteRoot

	var mockServer *httptest.Server = setupMockServer(noteResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	err := json.Unmarshal([]byte(noteResponse), &dataNote)
	if err != nil {
		t.Error(err)
	}

	noteBytes, err := CreateNote(client, "12345678", &dataNote)
	if err != nil {
		t.Error(err)
	}

	var note core.NoteRoot
	json.Unmarshal(noteBytes, &note)

	if note.Data.Attributes.Note != "Test note from the api" {
		t.Errorf("Address is not 'Test note from the api', but is showing as : %v", note.Data.Attributes.Note)
	}

}

func TestGetNote(t *testing.T) {
	var note core.NoteRoot

	var mockServer *httptest.Server = setupMockServer(noteResponse, http.StatusOK)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)
	// Initialize your PC_Client with the mock server URL

	note, err := GetNote(client, "1")
	if err != nil {
		t.Errorf("GetNote failed with an error ::: %v\n", err)
	}

	if note.Data.Attributes.Note != "Test note from the api" {
		t.Errorf("Note is not 'Test note from the api', but is showing as : %v", note.Data.Attributes.Note)
	}

}

func TestDeleteNote(t *testing.T) {

	var mockServer *httptest.Server = setupMockServer("", http.StatusNoContent)

	client := core.NewPCClient(mockAppId, mockSecret, mockServer.URL)

	response := DeleteNote(client, "1")
	if response != nil {
		t.Errorf("DeleteNote did not return 'nil'")
	}

}
