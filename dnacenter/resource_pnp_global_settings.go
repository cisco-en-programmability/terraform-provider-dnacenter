package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

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

						"accept_eula": &schema.Schema{
							Description: `Accept Eula`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"default_profile": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cert": &schema.Schema{
										Description: `Cert`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"fqdn_addresses": &schema.Schema{
										Description: `Fqdn Addresses`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ip_addresses": &schema.Schema{
										Description: `Ip Addresses`,
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"proxy": &schema.Schema{
										Description: `Proxy`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"sava_mapping_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cco_user": &schema.Schema{
										Description: `Cco User`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"expiry": &schema.Schema{
										Description: `Expiry`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"profile": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address_fqdn": &schema.Schema{
													Description: `Address Fqdn`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"address_ip_v4": &schema.Schema{
													Description: `Address Ip V4`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"cert": &schema.Schema{
													Description: `Cert`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"make_default": &schema.Schema{
													Description: `Make Default`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"port": &schema.Schema{
													Description: `Port`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"profile_id": &schema.Schema{
													Description: `Profile Id`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"proxy": &schema.Schema{
													Description: `Proxy`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"smart_account_id": &schema.Schema{
										Description: `Smart Account Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"virtual_account_id": &schema.Schema{
										Description: `Virtual Account Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
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

		response1, restyResp1, _ := client.DeviceOnboardingPnp.GetPnpGlobalSettings()
		/*		if err != nil {
				diags = append(diags, diagError(
					"Failure when setting GetPnpGlobalSettings response",
					err))
				return diags
			}*/
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
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
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing PnpGlobalSettingsDelete", err, "Delete method is not supported",
		"Failure at PnpGlobalSettingsDelete, unexpected response", ""))

	return diags
}
func expandRequestPnpGlobalSettingsUpdatePnpGlobalSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettings {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUpdatePnpGlobalSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".accept_eula")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".accept_eula")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".accept_eula")))) {
		request.AcceptEula = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_profile")))) {
		request.DefaultProfile = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsDefaultProfile(ctx, key+".default_profile.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sava_mapping_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sava_mapping_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sava_mapping_list")))) {
		request.SavaMappingList = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListArray(ctx, key+".sava_mapping_list", d)
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
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy")))) {
		request.Proxy = interfaceToString(v)
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
	for item_no := range objs {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cco_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cco_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cco_user")))) {
		request.CcoUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiry")))) {
		request.Expiry = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = expandRequestPnpGlobalSettingsUpdatePnpGlobalSettingsSavaMappingListProfile(ctx, key+".profile.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
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
		request.MakeDefault = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy")))) {
		request.Proxy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
