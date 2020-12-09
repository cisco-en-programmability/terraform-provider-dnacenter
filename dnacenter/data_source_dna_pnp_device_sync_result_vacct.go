package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceSyncResultVacct() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceSyncResultVacctRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_sync_period": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"cco_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"last_sync": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_fqdn": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"address_ip_v4": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"cert": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"make_default": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"profile_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_sn_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sync_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"sync_msg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sync_result_str": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sync_start_time": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"sync_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnPDeviceSyncResultVacctRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	name := d.Get("name").(string)
	domain := d.Get("domain").(string)

	response, _, err := client.DeviceOnboardingPnP.GetSyncResultForVirtualAccount(domain, name)
	if err != nil {
		return diag.FromErr(err)
	}

	sItem := flattenPnPDeviceSyncResultVacctItem(response)
	if err := d.Set("item", sItem); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
