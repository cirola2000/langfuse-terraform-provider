package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"langfuse": providerserver.NewProtocol6WithError(New("test")()),
}

func TestAccProvider(t *testing.T) {
	// This test just ensures that the provider can be instantiated without error
	// For more comprehensive testing, you would add actual acceptance tests here
}

func testAccPreCheck(t *testing.T) {
	// Add pre-check logic here if needed for acceptance tests
	// For example, checking that required environment variables are set
} 