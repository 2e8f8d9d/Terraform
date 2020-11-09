package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

func TestTerraformAws(t *testing.T) {
	var user User

	user.FirstName = "Caige"
	user.LastName = "Kelly"

	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the IP of the instance
	publicIp := terraform.Output(t, terraformOptions, "public_ip")

	// Make an HTTP request to the instance
	url := fmt.Sprintf("http://%s:3000/users", publicIp)
	http_helper.HttpGetWithRetry(t, url, nil, 200, "null", 10, 2*time.Second)
}
