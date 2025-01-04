package people

import (
	"os"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var formIdFormField string = "873385"
var appIdFormField = os.Getenv("PC_APP_ID")
var secretTokenFormField = os.Getenv("PC_SECRET_TOKEN")

func TestGetFormField(t *testing.T) {
	var form core.FormFieldsRoot

	if appIdForm == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenForm == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}
	// Initialize your PC_Client with the mock server URL
	client := core.NewPCClient(appIdForm, secretTokenForm)

	form, err := GetFormField(client, formId, "6830109")
	if err != nil {
		t.Errorf("GetFormField failed with an error ::: %v\n", err)
	}

	if form.Data.Attributes.Label != "Phone number" {
		t.Errorf("Field ID 6830109 is not 'Phone number', but is showing as : %v", form.Data.Attributes.Label)
	}
}
