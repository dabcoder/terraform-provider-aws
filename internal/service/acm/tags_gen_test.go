// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package acm_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	tfacm "github.com/hashicorp/terraform-provider-aws/internal/service/acm"
)

func expectFullTags(resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullTags(tfacm.ServicePackage(context.Background()), resourceAddress, knownValue)
}
