package people

import (
	"os"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var formId string = "873385"
var appIdForm = os.Getenv("PC_APP_ID")
var secretTokenForm = os.Getenv("PC_SECRET_TOKEN")

func TestGetForm(t *testing.T) {
	var form core.FormRoot

	if appIdForm == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenForm == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdForm, secretTokenForm)

	form, err := GetForm(client, formId)
	if err != nil {
		t.Errorf("GetNote failed with an error ::: %v\n", err)
	}

	if form.Data.Attributes.Name != "ManualTestForm" {
		t.Errorf("Note is not 'ManualTestForm', but is showing as : %v", form.Data.Attributes.Name)
	}

}
