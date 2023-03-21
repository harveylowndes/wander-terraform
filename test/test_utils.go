package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func Setup(t *testing.T) *terraform.Options {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/project_infrastructure/backend",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})

	terraform.WorkspaceSelectOrNew(t, terraformOptions, "test")
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "bucket_url")
	assert.Equal(t, "gs://wander-project-remote-tfstate-bucket", output)
	return terraformOptions
}