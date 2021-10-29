package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

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
		if v, ok := si["device_sn_list"]; ok && v != nil {
			syncItem.DeviceSnList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := si["sync_type"]; ok && v != nil {
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
		if v, ok := sr["sync_list"]; ok && v != nil {
			if w := constructSyncVirtualAccountDevicesRequestSyncResultSyncList(v.([]interface{})); w != nil {
				result.SyncList = *w
			}
		}
		if v, ok := sr["sync_msg"]; ok && v != nil {
			result.SyncMsg = v.(string)
		}
	}
	return &result
}

func constructSyncVirtualAccountDevicesRequestProfile(response []interface{}) *dnac.SyncVirtualAccountDevicesRequestProfile {
	var result dnac.SyncVirtualAccountDevicesRequestProfile
	if len(response) > 0 {
		drp := response[0].(map[string]interface{})

		if v, ok := drp["address_fqdn"]; ok && v != nil {
			result.AddressFqdn = v.(string)
		}
		if v, ok := drp["address_ip_v4"]; ok && v != nil {
			result.AddressIPV4 = v.(string)
		}
		if v, ok := drp["cert"]; ok && v != nil {
			result.Cert = v.(string)
		}
		if v, ok := drp["make_default"]; ok && v != nil {
			result.MakeDefault = v.(bool)
		}
		if v, ok := drp["name"]; ok && v != nil {
			result.Name = v.(string)
		}
		if v, ok := drp["port"]; ok && v != nil {
			result.Port = v.(int)
		}
		if v, ok := drp["profile_id"]; ok && v != nil {
			result.ProfileID = v.(string)
		}
		if v, ok := drp["proxy"]; ok && v != nil {
			result.Proxy = v.(bool)
		}
	}
	return &result
}

func constructSyncVirtualAccountDevicesRequest(response []interface{}) *dnac.SyncVirtualAccountDevicesRequest {
	var result dnac.SyncVirtualAccountDevicesRequest
	if len(response) > 0 {
		dr := response[0].(map[string]interface{})
		if v, ok := dr["auto_sync_period"]; ok && v != nil {
			result.AutoSyncPeriod = v.(int)
		}
		if v, ok := dr["cco_user"]; ok && v != nil {
			result.CcoUser = v.(string)
		}
		if v, ok := dr["expiry"]; ok && v != nil {
			result.Expiry = v.(int)
		}
		if v, ok := dr["last_sync"]; ok && v != nil {
			result.LastSync = v.(int)
		}
		if v, ok := dr["profile"]; ok && v != nil {
			if w := constructSyncVirtualAccountDevicesRequestProfile(v.([]interface{})); w != nil {
				result.Profile = *w
			}
		}
		if v, ok := dr["smart_account_id"]; ok && v != nil {
			result.SmartAccountID = v.(string)
		}
		if v, ok := dr["sync_result"]; ok && v != nil {
			if w := constructSyncVirtualAccountDevicesRequestSyncResult(v.([]interface{})); w != nil {
				result.SyncResult = *w
			}
		}
		if v, ok := dr["sync_result_str"]; ok && v != nil {
			result.SyncResultStr = v.(string)
		}
		if v, ok := dr["sync_start_time"]; ok && v != nil {
			result.SyncStartTime = v.(int)
		}
		if v, ok := dr["sync_status"]; ok && v != nil {
			result.SyncStatus = v.(string)
		}
		if v, ok := dr["tenant_id"]; ok && v != nil {
			result.TenantID = v.(string)
		}
		if v, ok := dr["token"]; ok && v != nil {
			result.Token = v.(string)
		}
		if v, ok := dr["virtual_account_id"]; ok && v != nil {
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
