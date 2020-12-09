package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoverySNMPProperty() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoverySNMPPropertyRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"int_value": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"system_property_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoverySNMPPropertyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	response, _, err := client.Discovery.GetSNMPProperties()
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenSNMPPropertiesRangeReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
