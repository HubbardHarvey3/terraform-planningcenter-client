package people

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/HubbardHarvey3/terraform-planningcenter-client/core"
)

var appIdOrganization = os.Getenv("PC_APP_ID")
var secretTokenOrganization = os.Getenv("PC_SECRET_TOKEN")

func TestGetOrganization(t *testing.T) {
	if appIdOrganization == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenOrganization == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdOrganization, secretTokenOrganization)

	org, err := GetOrganization(client)
	if err != nil {
		t.Errorf("GetPeople failed with an error ::: %v\n", err)
	}

	if org.Data.Attributes.Name != "CBC" {
		t.Errorf("Expected org.Data.Attributes.Name to be CBC, instead got %v\n", org.Data.Attributes.Name)
	}

}

func TestGetOrganizationPeople(t *testing.T) {
	if appIdOrganization == "" {
		t.Errorf("Need Env Vars PC_APP_ID Set")
	}
	if secretTokenOrganization == "" {
		t.Errorf("Need Env Vars PC_SECRET_TOKEN Set")
	}

	client := core.NewPCClient(appIdOrganization, secretTokenOrganization)

	org, err := GetOrganizationPeople(client)
	if err != nil {
		t.Errorf("GetPeople failed with an error ::: %v\n", err)
	}

	orgJSON, _ := json.Marshal(org)

	fmt.Println(string(orgJSON))

}
