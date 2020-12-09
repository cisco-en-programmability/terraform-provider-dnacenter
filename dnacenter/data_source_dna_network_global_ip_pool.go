package dnacenter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkGlobalIPPool() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkGlobalIPPoolRead,
		Schema: map[string]*schema.Schema{

			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"configure_external_dhcp": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"context_key": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"context_value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dhcp_server_ips": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dns_server_ips": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gateways": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_pool_cidr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_pool_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"overlapping": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"owner": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"total_ip_address_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_ip_address_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_percentage": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkGlobalIPPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetGlobalPoolQueryParams{}
	if v, ok := d.GetOk("offset"); ok {
		queryParams.Offset = fmt.Sprintf("%d", v.(int))
	}
	if v, ok := d.GetOk("limit"); ok {
		queryParams.Limit = fmt.Sprintf("%d", v.(int))
	}

	response, _, err := client.NetworkSettings.GetGlobalPool(&queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenNetworkGlobalIPPoolReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
