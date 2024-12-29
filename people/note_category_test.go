package people

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonNoteCategory = `{
	"data": {
		"type": "person",
		"attributes": {
			"accounting_administrator": false,
			"anniversary": null,
			"birthdate": "1990-01-01",
			"first_name": "UnitTestNote",
			"gender": "male",
			"given_name": null,
			"grade": null,
			"graduation_year": null,
			"inactivated_at": null,
			"last_name": "NoteDoe",
			"medical_notes": null,
			"membership": "member",
			"middle_name": null,
			"nickname": null,
			"site_administrator": false,
			"status": "active"
		}
	}
}`

var responseNoteCategory = `{
	"data": {
		"type": "NoteCategory",
		"attributes": {
			"name": "ClientTestNoteCategory"
		}
	} 
}`

var noteCategoryId string
var appIdNoteCategory = os.Getenv("PC_APP_ID")
var secretTokenNoteCategory = os.Getenv("PC_SECRET_TOKEN")

func TestCreateNoteCategory(t *testing.T) {
	var dataPerson core.PeopleRoot
	var dataNote core.NoteCategoryRoot
	dataNote.Data.Relationships = nil

	if appIdNoteCategory == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenNoteCategory == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into core.PeopleRoot
	err := json.Unmarshal([]byte(responsePersonNoteCategory), &dataPerson)
	if err != nil {
		t.Error(err)
	}

	client := core.NewPCClient(appIdNoteCategory, secretTokenNoteCategory)

	person, err := CreatePeople(client, &dataPerson)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var responsePerson core.PeopleRoot
	json.Unmarshal(person, &responsePerson)

	personId = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responseNoteCategory), &dataNote)
	if err != nil {
		t.Error(err)
	}

	noteBytes, err := CreateNoteCategory(client, &dataNote)
	if err != nil {
		t.Error(err)
	}

	var note core.NoteCategoryRoot
	json.Unmarshal(noteBytes, &note)
	noteCategoryId = note.Data.ID

	if note.Data.Attributes.Name != "ClientTestNoteCategory" {
		t.Errorf("NoteCategory Name is not 'ClientTestNoteCategory', but is showing as : %v", note.Data.Attributes.Name)
	}

}

func TestGetNoteCategory(t *testing.T) {
	var note core.NoteCategoryRoot

	if appIdNoteCategory == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenNoteCategory == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdNoteCategory, secretTokenNoteCategory)

	note, err := GetNoteCategory(client, noteCategoryId)
	if err != nil {
		t.Errorf("GetNote failed with an error ::: %v\n", err)
	}

	if note.Data.Attributes.Name != "ClientTestNoteCategory" {
		t.Errorf("Note is not 'ClientTestNoteCategory', but is showing as : %v", note.Data.Attributes.Name)
	}

}

func TestDeleteNoteCategory(t *testing.T) {

	if appIdNoteCategory == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenNoteCategory == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdNoteCategory, secretTokenNoteCategory)

	err := DeleteNoteCategory(client, noteCategoryId)
	if err != nil {
		t.Errorf("Error during DeleteNote : %v\n", err)
	}

	_, err = GetNote(client, noteCategoryId)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetNote should be throwing a 404 after the person was deleted")
	}

	err = DeletePeople(client, personId)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}
}
