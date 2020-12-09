package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteHealth() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSiteHealthRead,
		Schema: map[string]*schema.Schema{
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_bytes_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_health": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"client_health_wired": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"client_health_wireless": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"core_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"core_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"distribution_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"distribution_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dnac_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"healthy_clients_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"healthy_network_device_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"latitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"longitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"network_health_access": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_average": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_core": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_distribution": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_others": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_router": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_health_wireless": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_network_device": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_wired_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_wireless_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"overall_good_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"parent_site_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_site_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"router_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"router_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_number_of_active_wireless_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"total_number_of_connected_wired_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wired_good_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wireless_device_good_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wireless_device_total_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wireless_good_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_health_stats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_total_count": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_irrelevant_app_fair": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_irrelevant_app_good": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_irrelevant_app_poor": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_relevant_app_fair": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_relevant_app_good": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"business_relevant_app_poor": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"default_health_app_fair": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"default_health_app_good": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"default_health_app_poor": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	SiteHealthQueryParams := dnac.GetSiteHealthQueryParams{}
	if v, ok := d.GetOk("timestamp"); ok {
		SiteHealthQueryParams.Timestamp = v.(string)
	}

	response, _, err := client.Sites.GetSiteHealth(&SiteHealthQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	siteHealthItems := flattenSiteHealthReadItem(response)
	if err := d.Set("items", siteHealthItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
