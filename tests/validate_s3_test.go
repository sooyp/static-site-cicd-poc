package test

import (
	"testing"
	"fmt"

	terratest_aws "github.com/gruntwork-io/terratest/modules/aws"
	terratest_terraform "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestCheckS3(t *testing.T) {
	// hardcoded for now
	env := "dev"
	region := "eu-west-1"
	bucketName := "my-static-website-" + env

	tfOptions := &terratest_terraform.Options{
		TerraformDir: "../terraform",

		Vars: map[string]interface{}{
			"environment": env,
		},

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": region,
		},

		NoColor: true,
	}
	// Safety check to skip test for production
	if env == "prod" || env == "production" {
		t.Skip("Skipping Terratest against production environment for safety.")
	}
	// Apply the Terraform code
	terratest_terraform.InitAndApply(t, tfOptions)

	// S3 check bit
	doesExist := terratest_aws.DoesS3BucketExist(t, region, bucketName)

	// Not sure if this is enough, might need more validation later
	assert.True(t, doesExist, fmt.Sprintf("Bucket %s should exist but doesn't", bucketName))
}
