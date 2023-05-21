package todotest

import (
	"DEMO__Terraform_provider_test/provider/client"
	todo "DEMO__Terraform_provider_test/provider/testing/data_source"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		Schema: map[string]*schema.Schema{
			"some_string_config": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"todotest_item": todo.Test(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := "https://jsonplaceholder.typicode.com"
	return client.Create(address), nil
}
