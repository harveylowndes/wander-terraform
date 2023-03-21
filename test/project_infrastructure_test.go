package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestProjectInfrastructureIAM(t *testing.T) {
	defer terraform.Destroy(t, Setup(t))

	// Start
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/project_infrastructure/iam",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptions)

	terraform.WorkspaceSelectOrNew(t, terraformOptions, "test")
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "cluster_service_account_id")
	assert.Equal(t, fmt.Sprintf("projects/%s/serviceAccounts/%s", "vernal-union-381217", terraform.Output(t, terraformOptions, "cluster_service_account_email")), output)
}