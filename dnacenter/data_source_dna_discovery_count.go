package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Discovery ID",
				Required:    true,
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDiscoveryCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetDevicesDiscoveredByIDQueryParams{}
	id := d.Get("id").(string)
	if v, ok := d.GetOk("task_id"); ok {
		queryParams.TaskID = v.(string)
	}

	// Prepare Request
	response, _, err := client.Discovery.GetDevicesDiscoveredByID(id, &queryParams)
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
