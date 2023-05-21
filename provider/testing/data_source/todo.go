package todo

import (
	"DEMO__Terraform_provider_test/provider/client"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// initNameAllocationSchema is where we define the schema of the Terraform data source
func Test() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description",
				ForceNew:    true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"completed": {
				Type:     schema.TypeBool,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
		},
	}
}

type Item struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func dataSourceRead(d *schema.ResourceData, m interface{}) error {

	api := m.(*client.Client)

	todoId := d.Get("id").(string)
	endpoint := fmt.Sprintf("todos/%s", todoId)
	body, err := api.Get(endpoint)
	if err != nil {
		return err
	}

	item := Item{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return err
	}

	d.SetId(todoId)
	d.Set("title", item.Title)
	d.Set("user_id", item.UserId)
	d.Set("completed", item.Completed)

	return nil
}
