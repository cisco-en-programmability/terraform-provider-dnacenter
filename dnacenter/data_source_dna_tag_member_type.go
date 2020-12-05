package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagMemberType() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagMemberTypeRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceTagMemberTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Prepare Request
	response, _, err := client.Tag.GetTagResourceTypes()
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source
	if err := d.Set("items", response.Response); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
