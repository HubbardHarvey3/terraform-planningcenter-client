package people

import (
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

	fmt.Println(org)
	fmt.Println("Org NAME ::: " + org.Data.Attributes.Name)

	//	if person.Data.Attributes.FirstName != "UnitTest" {
	//		t.Errorf("Expected person.Data.ATtributes.FirstName to be UnitTest, instead got %v\n", person.Data.Attributes.FirstName)
	//	}
	//
	//	if person.Data.Attributes.Birthdate != "1990-01-01" {
	//		t.Errorf("Expected person.Data.Attributes.Birthdate to be 1990-01-01, instead got %v\n", person.Data.Attributes.Birthdate)
	//	}

}
