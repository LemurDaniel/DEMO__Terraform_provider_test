package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	todotest "DEMO__Terraform_provider_test/provider/testing"
	//"DEMO__Terraform_provider_test/azurepim"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return todotest.Provider()
		},
	})
}
