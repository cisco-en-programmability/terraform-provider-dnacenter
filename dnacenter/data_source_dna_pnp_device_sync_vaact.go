package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceSyncVacct() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceSyncVacctRead,
		Schema: map[string]*schema.Schema{
			"request": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_sync_period": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"cco_user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"last_sync": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_fqdn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"address_ip_v4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"make_default": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeFloat,
										Optional: true,
									},
									"profile_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"proxy": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_sn_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sync_type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"sync_msg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sync_result_str": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sync_start_time": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"sync_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

func constructSyncVirtualAccountDevicesRequestSyncResultSyncList(response []interface{}) *[]dnac.SyncVirtualAccountDevicesRequestSyncResultSyncList {
	var result []dnac.SyncVirtualAccountDevicesRequestSyncResultSyncList
	for _, item := range response {
		syncItem := dnac.SyncVirtualAccountDevicesRequestSyncResultSyncList{}
		si := item.(map[string]interface{})
		if v, ok := si["device_sn_list"]; ok {
			syncItem.DeviceSnList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := si["sync_type"]; ok {
			syncItem.SyncType = v.(string)
		}
		result = append(result, syncItem)
	}
	return &result
}

func constructSyncVirtualAccountDevicesRequestSyncResult(response []interface{}) *dnac.SyncVirtualAccountDevicesRequestSyncResult {
	var result dnac.SyncVirtualAccountDevicesRequestSyncResult
	if len(response) > 0 {
		sr := response[0].(map[string]interface{})
		if v, ok := sr["sync_list"]; ok {
			if w := constructSyncVirtualAccountDevicesRequestSyncResultSyncList(v.([]interface{})); v != nil {
				result.SyncList = *w
			}
		}
		if v, ok := sr["sync_msg"]; ok {
			result.SyncMsg = v.(string)
		}
	}
	return &result
}

func constructSyncVirtualAccountDevicesRequestProfile(response []interface{}) *dnac.SyncVirtualAccountDevicesRequestProfile {
	var result dnac.SyncVirtualAccountDevicesRequestProfile
	if len(response) > 0 {
		drp := response[0].(map[string]interface{})

		if v, ok := drp["address_fqdn"]; ok {
			result.AddressFqdn = v.(string)
		}
		if v, ok := drp["address_ip_v4"]; ok {
			result.AddressIPV4 = v.(string)
		}
		if v, ok := drp["cert"]; ok {
			result.Cert = v.(string)
		}
		if v, ok := drp["make_default"]; ok {
			result.MakeDefault = v.(bool)
		}
		if v, ok := drp["name"]; ok {
			result.Name = v.(string)
		}
		if v, ok := drp["port"]; ok {
			result.Port = v.(int)
		}
		if v, ok := drp["profile_id"]; ok {
			result.ProfileID = v.(string)
		}
		if v, ok := drp["proxy"]; ok {
			result.Proxy = v.(bool)
		}
	}
	return &result
}

func constructSyncVirtualAccountDevicesRequest(response []interface{}) *dnac.SyncVirtualAccountDevicesRequest {
	var result dnac.SyncVirtualAccountDevicesRequest
	if len(response) > 0 {
		dr := response[0].(map[string]interface{})
		if v, ok := dr["auto_sync_period"]; ok {
			result.AutoSyncPeriod = v.(int)
		}
		if v, ok := dr["cco_user"]; ok {
			result.CcoUser = v.(string)
		}
		if v, ok := dr["expiry"]; ok {
			result.Expiry = v.(int)
		}
		if v, ok := dr["last_sync"]; ok {
			result.LastSync = v.(int)
		}
		if v, ok := dr["profile"]; ok {
			if w := constructSyncVirtualAccountDevicesRequestProfile(v.([]interface{})); v != nil {
				result.Profile = *w
			}
		}
		if v, ok := dr["smart_account_id"]; ok {
			result.SmartAccountID = v.(string)
		}
		if v, ok := dr["sync_result"]; ok {
			if w := constructSyncVirtualAccountDevicesRequestSyncResult(v.([]interface{})); v != nil {
				result.SyncResult = *w
			}
		}
		if v, ok := dr["sync_result_str"]; ok {
			result.SyncResultStr = v.(string)
		}
		if v, ok := dr["sync_start_time"]; ok {
			result.SyncStartTime = v.(int)
		}
		if v, ok := dr["sync_status"]; ok {
			result.SyncStatus = v.(string)
		}
		if v, ok := dr["tenant_id"]; ok {
			result.TenantID = v.(string)
		}
		if v, ok := dr["token"]; ok {
			result.Token = v.(string)
		}
		if v, ok := dr["virtual_account_id"]; ok {
			result.VirtualAccountID = v.(string)
		}
	}
	return &result
}

func dataSourcePnPDeviceSyncVacctRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	request := dnac.SyncVirtualAccountDevicesRequest{}

	response, _, err := client.DeviceOnboardingPnP.SyncVirtualAccountDevices(&request)
	if err != nil {
		return diag.FromErr(err)
	}

	sItem := flattenPnPDeviceSyncVacctItem(response)
	if err := d.Set("item", sItem); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
