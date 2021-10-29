package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aaa_credentials": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"accept_eula": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"default_profile": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ip_addresses": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"proxy": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sava_mapping_list": &schema.Schema{
				Type:     schema.TypeList, //[]UpdatePnPGlobalSettingsRequestSavaMappingList
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_sync_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cco_user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"last_sync": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_fqdn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address_ip_v4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"make_default": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"profile_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"proxy": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_sn_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sync_type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"sync_msg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"sync_result_str": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sync_start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"sync_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"task_time_outs": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"general_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"image_download_time_out": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePnPGlobalSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePnPGlobalSettingsCreate,
		ReadContext:   resourcePnPGlobalSettingsRead,
		UpdateContext: resourcePnPGlobalSettingsUpdate,
		DeleteContext: resourcePnPGlobalSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem:     pnpSettings(),
			},
		},
	}
}

///// start construct for add

func constructUpdatePnPSettingsAAACredentials(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestAAACredentials {
	var result dnac.UpdatePnPGlobalSettingsRequestAAACredentials
	if len(response) > 0 {
		cred := response[0].(map[string]interface{})
		if v, ok := cred["password"]; ok && v != nil {
			result.Password = v.(string)
		}
		if v, ok := cred["username"]; ok && v != nil {
			result.Username = v.(string)
		}
	}
	return &result
}

func constructUpdatePnPSettingsDefaultProfile(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestDefaultProfile {
	var result dnac.UpdatePnPGlobalSettingsRequestDefaultProfile
	if len(response) > 0 {
		profile := response[0].(map[string]interface{})
		if v, ok := profile["cert"]; ok && v != nil {
			result.Cert = v.(string)
		}
		if v, ok := profile["fqdn_addresses"]; ok && v != nil {
			result.FqdnAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := profile["ip_addresses"]; ok && v != nil {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := profile["port"]; ok && v != nil {
			result.Port = v.(int)
		}
		if v, ok := profile["proxy"]; ok && v != nil {
			result.Proxy = v.(bool)
		}
	}
	return &result
}

func constructUpdatePnPSettingsSavaMappingListProfile(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestSavaMappingListProfile {
	var result dnac.UpdatePnPGlobalSettingsRequestSavaMappingListProfile
	if len(response) > 0 {
		p := response[0].(map[string]interface{})
		if v, ok := p["address_fqdn"]; ok && v != nil {
			result.AddressFqdn = v.(string)
		}
		if v, ok := p["address_ip_v4"]; ok && v != nil {
			result.AddressIPV4 = v.(string)
		}
		if v, ok := p["cert"]; ok && v != nil {
			result.Cert = v.(string)
		}
		if v, ok := p["make_default"]; ok && v != nil {
			result.MakeDefault = v.(bool)
		}
		if v, ok := p["name"]; ok && v != nil {
			result.Name = v.(string)
		}
		if v, ok := p["port"]; ok && v != nil {
			result.Port = v.(int)
		}
		if v, ok := p["profile_id"]; ok && v != nil {
			result.ProfileID = v.(string)
		}
		if v, ok := p["proxy"]; ok && v != nil {
			result.Proxy = v.(bool)
		}
	}
	return &result
}

func constructUpdatePnPSettingsSavaMappingListSyncResultSyncList(response []interface{}) *[]dnac.UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList {
	var syncList []dnac.UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList
	for _, sm := range response {
		smi := sm.(map[string]interface{})
		syncItem := dnac.UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList{}
		if v, ok := smi["device_sn_list"]; ok && v != nil {
			syncItem.DeviceSnList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := smi["sync_type"]; ok && v != nil {
			syncItem.SyncType = v.(string)
		}
		syncList = append(syncList, syncItem)
	}
	return &syncList
}

func constructUpdatePnPSettingsSavaMappingListSyncResult(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestSavaMappingListSyncResult {
	var result dnac.UpdatePnPGlobalSettingsRequestSavaMappingListSyncResult
	if len(response) > 0 {
		p := response[0].(map[string]interface{})
		if v, ok := p["sync_list"]; ok && v != nil {
			if w := constructUpdatePnPSettingsSavaMappingListSyncResultSyncList(v.([]interface{})); w != nil {
				result.SyncList = *w
			}
		}
		if v, ok := p["sync_msg"]; ok && v != nil {
			result.SyncMsg = v.(string)
		}
	}
	return &result
}

func constructUpdatePnPSettingsSavaMappingList(response []interface{}) *[]dnac.UpdatePnPGlobalSettingsRequestSavaMappingList {
	var savaMappingList []dnac.UpdatePnPGlobalSettingsRequestSavaMappingList
	for _, sm := range response {
		smi := sm.(map[string]interface{})
		savaMappingItem := dnac.UpdatePnPGlobalSettingsRequestSavaMappingList{}
		if v, ok := smi["auto_sync_period"]; ok && v != nil {
			savaMappingItem.AutoSyncPeriod = v.(int)
		}
		if v, ok := smi["cco_user"]; ok && v != nil {
			savaMappingItem.CcoUser = v.(string)
		}
		if v, ok := smi["expiry"]; ok && v != nil {
			savaMappingItem.Expiry = v.(int)
		}
		if v, ok := smi["last_sync"]; ok && v != nil {
			savaMappingItem.LastSync = v.(int)
		}
		if v, ok := smi["profile"]; ok && v != nil {
			if w := constructUpdatePnPSettingsSavaMappingListProfile(v.([]interface{})); w != nil {
				savaMappingItem.Profile = *w
			}
		}
		if v, ok := smi["smart_account_id"]; ok && v != nil {
			savaMappingItem.SmartAccountID = v.(string)
		}
		if v, ok := smi["sync_result"]; ok && v != nil {
			if w := constructUpdatePnPSettingsSavaMappingListSyncResult(v.([]interface{})); w != nil {
				savaMappingItem.SyncResult = *w
			}
		}
		if v, ok := smi["sync_result_str"]; ok && v != nil {
			savaMappingItem.SyncResultStr = v.(string)
		}
		if v, ok := smi["sync_start_time"]; ok && v != nil {
			savaMappingItem.SyncStartTime = v.(int)
		}
		if v, ok := smi["sync_status"]; ok && v != nil {
			savaMappingItem.SyncStatus = v.(string)
		}
		if v, ok := smi["tenant_id"]; ok && v != nil {
			savaMappingItem.TenantID = v.(string)
		}
		if v, ok := smi["token"]; ok && v != nil {
			savaMappingItem.Token = v.(string)
		}
		if v, ok := smi["virtual_account_id"]; ok && v != nil {
			savaMappingItem.VirtualAccountID = v.(string)
		}

		savaMappingList = append(savaMappingList, savaMappingItem)
	}
	return &savaMappingList
}

func constructUpdatePnPSettingsTaskTimeOuts(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestTaskTimeOuts {
	var result dnac.UpdatePnPGlobalSettingsRequestTaskTimeOuts
	if len(response) > 0 {
		timeOuts := response[0].(map[string]interface{})

		if v, ok := timeOuts["config_time_out"]; ok && v != nil {
			result.ConfigTimeOut = v.(int)
		}
		if v, ok := timeOuts["general_time_out"]; ok && v != nil {
			result.GeneralTimeOut = v.(int)
		}
		if v, ok := timeOuts["image_download_time_out"]; ok && v != nil {
			result.ImageDownloadTimeOut = v.(int)
		}
	}
	return &result
}

func constructUpdatePnPSettings(ws map[string]interface{}) *dnac.UpdatePnPGlobalSettingsRequest {
	var SettingsItem dnac.UpdatePnPGlobalSettingsRequest
	if v, ok := ws["id"]; ok && v != nil {
		SettingsItem.TypeID = v.(string)
	}
	if v, ok := ws["aaa_credentials"]; ok && v != nil {
		if w := constructUpdatePnPSettingsAAACredentials(v.([]interface{})); w != nil {
			SettingsItem.AAACredentials = *w
		}
	}
	if v, ok := ws["accept_eula"]; ok && v != nil {
		SettingsItem.AcceptEula = v.(bool)
	}
	if v, ok := ws["default_profile"]; ok && v != nil {
		if w := constructUpdatePnPSettingsDefaultProfile(v.([]interface{})); w != nil {
			SettingsItem.DefaultProfile = *w
		}
	}
	if v, ok := ws["sava_mapping_list"]; ok && v != nil {
		if w := constructUpdatePnPSettingsSavaMappingList(v.([]interface{})); w != nil {
			SettingsItem.SavaMappingList = *w
		}
	}
	if v, ok := ws["task_time_outs"]; ok && v != nil {
		if w := constructUpdatePnPSettingsTaskTimeOuts(v.([]interface{})); w != nil {
			SettingsItem.TaskTimeOuts = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok && v != nil {
		SettingsItem.TenantID = v.(string)
	}
	if v, ok := ws["version"]; ok && v != nil {
		SettingsItem.Version = v.(int)
	}
	return &SettingsItem
}

///// end construct for add
///// start construct for update

///// end construct for update

func resourcePnPGlobalSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// Update resource id
	d.SetId("global")
	// Update resource on Terraform
	resourcePnPGlobalSettingsRead(ctx, d, m)
	return diags
}

func resourcePnPGlobalSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	response, _, err := client.DeviceOnboardingPnP.GetPnPGlobalSettings()
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if response == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	deviceItem := flattenPnPGlobalSettingsReadItems(response)
	if err := d.Set("item", deviceItem); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePnPGlobalSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		pnpRequest := item.(map[string]interface{})
		request := constructUpdatePnPSettings(pnpRequest)

		_, _, err := client.DeviceOnboardingPnP.UpdatePnPGlobalSettings(request)
		if err != nil {
			return diag.FromErr(err)
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourcePnPGlobalSettingsRead(ctx, d, m)
}

func resourcePnPGlobalSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
