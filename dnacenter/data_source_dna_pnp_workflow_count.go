package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPWorkflowCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPWorkflowCountRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPWorkflowCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParam := dnac.GetWorkflowCountQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		queryParam.Name = convertSliceInterfaceToSliceString(v.([]interface{}))
	}

	// Prepare Request
	response, _, err := client.DeviceOnboardingPnP.GetWorkflowCount(&queryParam)
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
