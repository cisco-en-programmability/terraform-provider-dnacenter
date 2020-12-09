package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpDataGlobalSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aaa_credentials": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"accept_eula": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_profile": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ip_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"proxy": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"sava_mapping_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_sync_period": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cco_user": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_sync": &schema.Schema{
							Type:     schema.TypeInt,
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
										Type:     schema.TypeInt,
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
							Type:     schema.TypeInt,
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
			"task_time_outs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"general_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"image_download_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPGlobalSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPGlobalSettingsRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     pnpDataGlobalSettings(),
			},
		},
	}
}

func dataSourcePnPGlobalSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	response, _, err := client.DeviceOnboardingPnP.GetPnPGlobalSettings()
	if err != nil {
		return diag.FromErr(err)
	}

	sItem := flattenPnPGlobalSettingsReadItems(response)
	if err := d.Set("item", sItem); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
