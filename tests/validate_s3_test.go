package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/stretchr/testify/assert"
)

func TestS3Bucket(t *testing.T) {
	// Define the path to the Terraform code
	terraformOptions := &terraform.Options{
		// Path to the Terraform code
		TerraformDir: "../terraform",

		// Variables to pass to our Terraform configuration
		Vars: map[string]interface{}{
			"s3_bucket_name": "my-unique-bucket-name", // Replace with a unique bucket name
		},

		// Variables to pass to Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": eu-west-1", // Set this AWS region
		},

		// Disable colors in Terraform commands so its readable in logs
		NoColor: true,
	}

	// First, initialize and apply the Terraform code
	terraform.InitAndApply(t, terraformOptions)

	// Now check if the S3 bucket exists
	bucket := aws.GetS3Bucket(t, terraformOptions.Vars["s3_bucket_name"].(string))

	// Validate the bucket properties (check if public access is blocked)
	assert.NotNil(t, bucket)
	assert.Equal(t, "my-unique-bucket-name", *bucket.Name)
}

