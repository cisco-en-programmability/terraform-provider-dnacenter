package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpGlobalSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns global PnP settings of the user
`,

		ReadContext: dataSourcePnpGlobalSettingsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"aaa_credentials": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"accept_eula": &schema.Schema{
							Description: `Accept Eula`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cert": &schema.Schema{
										Description: `Cert`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"fqdn_addresses": &schema.Schema{
										Description: `Fqdn Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ip_addresses": &schema.Schema{
										Description: `Ip Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"proxy": &schema.Schema{
										Description: `Proxy`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sava_mapping_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auto_sync_period": &schema.Schema{
										Description: `Auto Sync Period`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"cco_user": &schema.Schema{
										Description: `Cco User`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"expiry": &schema.Schema{
										Description: `Expiry`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"last_sync": &schema.Schema{
										Description: `Last Sync`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"profile": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address_fqdn": &schema.Schema{
													Description: `Address Fqdn`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"address_ip_v4": &schema.Schema{
													Description: `Address Ip V4`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"cert": &schema.Schema{
													Description: `Cert`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"make_default": &schema.Schema{
													Description: `Make Default`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"port": &schema.Schema{
													Description: `Port`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"profile_id": &schema.Schema{
													Description: `Profile Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"proxy": &schema.Schema{
													Description: `Proxy`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"smart_account_id": &schema.Schema{
										Description: `Smart Account Id`,
										Type:        schema.TypeString,
										Computed:    true,
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
																Description: `Device Sn List`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"sync_type": &schema.Schema{
																Description: `Sync Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"sync_msg": &schema.Schema{
													Description: `Sync Msg`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"sync_result_str": &schema.Schema{
										Description: `Sync Result Str`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sync_start_time": &schema.Schema{
										Description: `Sync Start Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"sync_status": &schema.Schema{
										Description: `Sync Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"token": &schema.Schema{
										Description: `Token`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"virtual_account_id": &schema.Schema{
										Description: `Virtual Account Id`,
										Type:        schema.TypeString,
										Computed:    true,
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
										Description: `Config Time Out`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"general_time_out": &schema.Schema{
										Description: `General Time Out`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"image_download_time_out": &schema.Schema{
										Description: `Image Download Time Out`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpGlobalSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPnpGlobalSettings")

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetPnpGlobalSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPnpGlobalSettings", err,
				"Failure at GetPnpGlobalSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpGetPnpGlobalSettingsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPnpGlobalSettings response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sava_mapping_list"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingList(item.SavaMappingList)
	respItem["task_time_outs"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemTaskTimeOuts(item.TaskTimeOuts)
	respItem["tenant_id"] = item.TenantID
	respItem["aaa_credentials"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemAAACredentials(item.AAACredentials)
	respItem["default_profile"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemDefaultProfile(item.DefaultProfile)
	respItem["accept_eula"] = boolPtrToString(item.AcceptEula)
	respItem["id"] = item.ID
	respItem["type_id"] = item.TypeID
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sync_status"] = item.SyncStatus
		respItem["sync_start_time"] = item.SyncStartTime
		respItem["sync_result"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListSyncResult(item.SyncResult)
		respItem["last_sync"] = item.LastSync
		respItem["tenant_id"] = item.TenantID
		respItem["profile"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListProfile(item.Profile)
		respItem["token"] = item.Token
		respItem["expiry"] = item.Expiry
		respItem["cco_user"] = item.CcoUser
		respItem["smart_account_id"] = item.SmartAccountID
		respItem["virtual_account_id"] = item.VirtualAccountID
		respItem["auto_sync_period"] = item.AutoSyncPeriod
		respItem["sync_result_str"] = item.SyncResultStr
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListSyncResult(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sync_list"] = flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListSyncResultSyncList(item.SyncList)
	respItem["sync_msg"] = item.SyncMsg

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListSyncResultSyncList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResultSyncList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sync_type"] = item.SyncType
		respItem["device_sn_list"] = item.DeviceSnList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemSavaMappingListProfile(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["address_ip_v4"] = item.AddressIPV4
	respItem["address_fqdn"] = item.AddressFqdn
	respItem["profile_id"] = item.ProfileID
	respItem["proxy"] = boolPtrToString(item.Proxy)
	respItem["make_default"] = boolPtrToString(item.MakeDefault)
	respItem["cert"] = item.Cert
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemTaskTimeOuts(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsTaskTimeOuts) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["image_download_time_out"] = item.ImageDownloadTimeOut
	respItem["config_time_out"] = item.ConfigTimeOut
	respItem["general_time_out"] = item.GeneralTimeOut

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemAAACredentials(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsAAACredentials) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["password"] = item.Password
	respItem["username"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetPnpGlobalSettingsItemDefaultProfile(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetPnpGlobalSettingsDefaultProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["fqdn_addresses"] = item.FqdnAddresses
	respItem["proxy"] = boolPtrToString(item.Proxy)
	respItem["cert"] = item.Cert
	respItem["ip_addresses"] = item.IPAddresses
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}
