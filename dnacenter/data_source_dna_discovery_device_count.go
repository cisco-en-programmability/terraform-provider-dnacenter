package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryDeviceCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"response": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDiscoveryDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Prepare Request
	response, _, err := client.Discovery.GetCountOfAllDiscoveryJobs()
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source
	if err := d.Set("response", response.Response); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
