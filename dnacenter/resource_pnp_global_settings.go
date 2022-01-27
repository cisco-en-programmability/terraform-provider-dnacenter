package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePnpGlobalSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Device Onboarding (PnP).

- Updates the user's list of global PnP settings
`,

		CreateContext: resourcePnpGlobalSettingsCreate,
		ReadContext:   resourcePnpGlobalSettingsRead,
		UpdateContext: resourcePnpGlobalSettingsUpdate,
		DeleteContext: resourcePnpGlobalSettingsDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"aaa_credentials": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"password": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"accept_eula": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"default_profile": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"fqdn_addresses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_addresses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"proxy": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
								},
							},
						},
						"sava_mapping_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auto_sync_period": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"cco_user": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"expiry": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"last_sync": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"profile": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
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
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"port": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"profile_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"proxy": &schema.Schema{
													// Type:     schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
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
										Type:     schema.TypeInt,
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
						"task_time_outs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_time_out": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"general_time_out": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"image_download_time_out": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpGlobalSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourcePnpGlobalSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceID := d.Id()
	//resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPnpGlobalSettings")

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetPnpGlobalSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetPnpGlobalSettings", err,
			// 	"Failure at GetPnpGlobalSettings, unexpected response", ""))
			// return diags
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourcePnpGlobalSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		request1 := expandRequestPnpGlobalSettingsUpdatePnpGlobalSettings(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.DeviceOnboardingPnp.UpdatePnpGlobalSettings(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePnpGlobalSettings", err, restyResp1.String(),
					"Failure at UpdatePnpGlobalSettings, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePnpGlobalSettings", err,
				"Failure at UpdatePnpGlobalSettings, unexpected response", ""))
			return diags
		}
	}

	return resourcePnpGlobalSettingsRead(ctx, d, m)
}

func resourcePnpGlobalSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete PnpGlobalSettings on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettings {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_credentials")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_credentials")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_credentials")))) {
		request.AAACredentials = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsAAACredentials(ctx, key+".aaa_credentials.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accept_eula")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accept_eula")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accept_eula")))) {
		request.AcceptEula = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_profile")))) {
		request.DefaultProfile = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsDefaultProfile(ctx, key+".default_profile.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sava_mapping_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sava_mapping_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sava_mapping_list")))) {
		request.SavaMappingList = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListArray(ctx, key+".sava_mapping_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".task_time_outs")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".task_time_outs")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".task_time_outs")))) {
		request.TaskTimeOuts = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsTaskTimeOuts(ctx, key+".task_time_outs.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsAAACredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsDefaultProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert")))) {
		request.Cert = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fqdn_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fqdn_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fqdn_addresses")))) {
		request.FqdnAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_addresses")))) {
		request.IPAddresses = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy")))) {
		request.Proxy = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_sync_period")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_sync_period")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_sync_period")))) {
		request.AutoSyncPeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cco_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cco_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cco_user")))) {
		request.CcoUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiry")))) {
		request.Expiry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync")))) {
		request.LastSync = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListProfile(ctx, key+".profile.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_result")))) {
		request.SyncResult = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResult(ctx, key+".sync_result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_result_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_result_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_result_str")))) {
		request.SyncResultStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_start_time")))) {
		request.SyncStartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_status")))) {
		request.SyncStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".token")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".token")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".token")))) {
		request.Token = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_fqdn")))) {
		request.AddressFqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_ip_v4")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_ip_v4")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_ip_v4")))) {
		request.AddressIPV4 = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert")))) {
		request.Cert = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".make_default")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".make_default")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".make_default")))) {
		request.MakeDefault = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy")))) {
		request.Proxy = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResult(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_list")))) {
		request.SyncList = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncListArray(ctx, key+".sync_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_msg")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_msg")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_msg")))) {
		request.SyncMsg = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList {
	request := []dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sn_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sn_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sn_list")))) {
		request.DeviceSnList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_type")))) {
		request.SyncType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsTaskTimeOuts(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_time_out")))) {
		request.ConfigTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".general_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".general_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".general_time_out")))) {
		request.GeneralTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_download_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_download_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_download_time_out")))) {
		request.ImageDownloadTimeOut = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
