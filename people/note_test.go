package people

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var responsePersonNote = `{
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

var responseNote = `{
	"data": {
		"type": "Note",
		"attributes": {
			"note": "Test note from the api"
		}
	} 
}
`

var noteId string
var appIdNote = os.Getenv("PC_APP_ID")
var secretTokenNote = os.Getenv("PC_SECRET_TOKEN")

func TestCreateNote(t *testing.T) {
	var dataPerson core.PeopleRoot
	var dataNote core.NoteRoot

	if appIdEmail == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenEmail == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	//Convert json into core.PeopleRoot
	err := json.Unmarshal([]byte(responsePersonNote), &dataPerson)
	if err != nil {
		t.Error(err)
	}

	client := core.NewPCClient(appIdNote, secretTokenNote)

	person, err := CreatePeople(client, &dataPerson)
	if err != nil {
		t.Errorf("Error during CreatePeople :: %v\n", err)
	}

	var responsePerson core.PeopleRoot
	json.Unmarshal(person, &responsePerson)

	personId = responsePerson.Data.ID

	err = json.Unmarshal([]byte(responseNote), &dataNote)
	if err != nil {
		t.Error(err)
	}

	noteBytes, err := CreateNote(client, personId, &dataNote)
	if err != nil {
		t.Error(err)
	}

	var note core.NoteRoot
	json.Unmarshal(noteBytes, &note)
	noteId = note.Data.ID

	if note.Data.Attributes.Note != "Test note from the api" {
		t.Errorf("Address is not 'Test note from the api', but is showing as : %v", note.Data.Attributes.Note)
	}

}

func TestDeleteNote(t *testing.T) {

	if appIdNote == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenNote == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdEmail, secretTokenEmail)

	err := DeleteNote(client, noteId)
	if err != nil {
		t.Errorf("Error during DeleteNote : %v\n", err)
	}

	_, err = GetNote(client, noteId)
	if !strings.Contains(err.Error(), "404") {
		t.Errorf("GetNote should be throwing a 404 after the person was deleted")
	}

	err = DeletePeople(client, personId)
	if err != nil {
		t.Errorf("Failed cleaning up testing resource")
	}
}
