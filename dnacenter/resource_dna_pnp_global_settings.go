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
		if v, ok := cred["password"]; ok {
			result.Password = v.(string)
		}
		if v, ok := cred["username"]; ok {
			result.Username = v.(string)
		}
	}
	return &result
}

func constructUpdatePnPSettingsDefaultProfile(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestDefaultProfile {
	var result dnac.UpdatePnPGlobalSettingsRequestDefaultProfile
	if len(response) > 0 {
		profile := response[0].(map[string]interface{})
		if v, ok := profile["cert"]; ok {
			result.Cert = v.(string)
		}
		if v, ok := profile["fqdn_addresses"]; ok {
			result.FqdnAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := profile["ip_addresses"]; ok {
			result.IPAddresses = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := profile["port"]; ok {
			result.Port = v.(int)
		}
		if v, ok := profile["proxy"]; ok {
			result.Proxy = v.(bool)
		}
	}
	return &result
}

func constructUpdatePnPSettingsSavaMappingListProfile(response []interface{}) *dnac.UpdatePnPGlobalSettingsRequestSavaMappingListProfile {
	var result dnac.UpdatePnPGlobalSettingsRequestSavaMappingListProfile
	if len(response) > 0 {
		p := response[0].(map[string]interface{})
		if v, ok := p["address_fqdn"]; ok {
			result.AddressFqdn = v.(string)
		}
		if v, ok := p["address_ip_v4"]; ok {
			result.AddressIPV4 = v.(string)
		}
		if v, ok := p["cert"]; ok {
			result.Cert = v.(string)
		}
		if v, ok := p["make_default"]; ok {
			result.MakeDefault = v.(bool)
		}
		if v, ok := p["name"]; ok {
			result.Name = v.(string)
		}
		if v, ok := p["port"]; ok {
			result.Port = v.(int)
		}
		if v, ok := p["profile_id"]; ok {
			result.ProfileID = v.(string)
		}
		if v, ok := p["proxy"]; ok {
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
		if v, ok := smi["device_sn_list"]; ok {
			syncItem.DeviceSnList = convertSliceInterfaceToSliceString(v.([]interface{}))
		}
		if v, ok := smi["sync_type"]; ok {
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
		if v, ok := p["sync_list"]; ok {
			if w := constructUpdatePnPSettingsSavaMappingListSyncResultSyncList(v.([]interface{})); w != nil {
				result.SyncList = *w
			}
		}
		if v, ok := p["sync_msg"]; ok {
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
		if v, ok := smi["auto_sync_period"]; ok {
			savaMappingItem.AutoSyncPeriod = v.(int)
		}
		if v, ok := smi["cco_user"]; ok {
			savaMappingItem.CcoUser = v.(string)
		}
		if v, ok := smi["expiry"]; ok {
			savaMappingItem.Expiry = v.(int)
		}
		if v, ok := smi["last_sync"]; ok {
			savaMappingItem.LastSync = v.(int)
		}
		if v, ok := smi["profile"]; ok {
			if w := constructUpdatePnPSettingsSavaMappingListProfile(v.([]interface{})); w != nil {
				savaMappingItem.Profile = *w
			}
		}
		if v, ok := smi["smart_account_id"]; ok {
			savaMappingItem.SmartAccountID = v.(string)
		}
		if v, ok := smi["sync_result"]; ok {
			if w := constructUpdatePnPSettingsSavaMappingListSyncResult(v.([]interface{})); w != nil {
				savaMappingItem.SyncResult = *w
			}
		}
		if v, ok := smi["sync_result_str"]; ok {
			savaMappingItem.SyncResultStr = v.(string)
		}
		if v, ok := smi["sync_start_time"]; ok {
			savaMappingItem.SyncStartTime = v.(int)
		}
		if v, ok := smi["sync_status"]; ok {
			savaMappingItem.SyncStatus = v.(string)
		}
		if v, ok := smi["tenant_id"]; ok {
			savaMappingItem.TenantID = v.(string)
		}
		if v, ok := smi["token"]; ok {
			savaMappingItem.Token = v.(string)
		}
		if v, ok := smi["virtual_account_id"]; ok {
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

		if v, ok := timeOuts["config_time_out"]; ok {
			result.ConfigTimeOut = v.(int)
		}
		if v, ok := timeOuts["general_time_out"]; ok {
			result.GeneralTimeOut = v.(int)
		}
		if v, ok := timeOuts["image_download_time_out"]; ok {
			result.ImageDownloadTimeOut = v.(int)
		}
	}
	return &result
}

func constructUpdatePnPSettings(ws map[string]interface{}) *dnac.UpdatePnPGlobalSettingsRequest {
	var SettingsItem dnac.UpdatePnPGlobalSettingsRequest
	if v, ok := ws["id"]; ok {
		SettingsItem.TypeID = v.(string)
	}
	if v, ok := ws["aaa_credentials"]; ok {
		if w := constructUpdatePnPSettingsAAACredentials(v.([]interface{})); w != nil {
			SettingsItem.AAACredentials = *w
		}
	}
	if v, ok := ws["accept_eula"]; ok {
		SettingsItem.AcceptEula = v.(bool)
	}
	if v, ok := ws["default_profile"]; ok {
		if w := constructUpdatePnPSettingsDefaultProfile(v.([]interface{})); w != nil {
			SettingsItem.DefaultProfile = *w
		}
	}
	if v, ok := ws["sava_mapping_list"]; ok {
		if w := constructUpdatePnPSettingsSavaMappingList(v.([]interface{})); w != nil {
			SettingsItem.SavaMappingList = *w
		}
	}
	if v, ok := ws["task_time_outs"]; ok {
		if w := constructUpdatePnPSettingsTaskTimeOuts(v.([]interface{})); w != nil {
			SettingsItem.TaskTimeOuts = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok {
		SettingsItem.TenantID = v.(string)
	}
	if v, ok := ws["version"]; ok {
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
