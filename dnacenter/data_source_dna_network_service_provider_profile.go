package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkServiceProviderProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkServiceProviderProfileRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_group_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inherited_group_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sla_profile_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sp_profile_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"wan_provider": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkServiceProviderProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	response, _, err := client.NetworkSettings.GetServiceProviderDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenNetworkServiceProviderProfileReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
