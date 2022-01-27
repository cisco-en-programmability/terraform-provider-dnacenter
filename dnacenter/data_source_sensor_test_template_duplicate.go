package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSensorTestTemplateDuplicate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.

- Intent API to duplicate an existing SENSOR test template
`,

		ReadContext: dataSourceSensorTestTemplateDuplicateRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `Bands`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number Of Aps To Test`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `Rssi Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `Connection`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"frequency": &schema.Schema{
							Description: `Frequency`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_modified_time": &schema.Schema{
							Description: `Last Modified Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"legacy_test_suite": &schema.Schema{
							Description: `Legacy Test Suite`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `Location`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensors": &schema.Schema{
										Description: `All Sensors`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Location Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"location_type": &schema.Schema{
										Description: `Location Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"mac_address_list": &schema.Schema{
										Description: `Mac Address List`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"model_version": &schema.Schema{
							Description: `Model Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"num_associated_sensor": &schema.Schema{
							Description: `Num Associated Sensor`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"num_neighbor_apthreshold": &schema.Schema{
							Description: `Num Neighbor A P Threshold`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"radio_as_sensor_removed": &schema.Schema{
							Description: `Radio As Sensor Removed`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rssi_threshold": &schema.Schema{
							Description: `Rssi Threshold`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"run_now": &schema.Schema{
							Description: `Run Now`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"frequency": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"unit": &schema.Schema{
													Description: `Unit`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},
									"schedule_range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"day": &schema.Schema{
													Description: `Day`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"time_range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"frequency": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"unit": &schema.Schema{
																			Description: `Unit`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"value": &schema.Schema{
																			Description: `Value`,
																			Type:        schema.TypeInt,
																			Computed:    true,
																		},
																	},
																},
															},
															"from": &schema.Schema{
																Description: `From`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"to": &schema.Schema{
																Description: `To`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"test_schedule_mode": &schema.Schema{
										Description: `Test Schedule Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"schedule_in_days": &schema.Schema{
							Description: `Schedule In Days`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"sensors": &schema.Schema{
							Description: `Sensors`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"show_wlc_upgrade_banner": &schema.Schema{
							Description: `Show Wlc Upgrade Banner`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"auth_type_rcvd": &schema.Schema{
										Description: `Auth Type Rcvd`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"bands": &schema.Schema{
										Description: `Bands`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certdownloadurl`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"certfilename": &schema.Schema{
										Description: `Certfilename`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certpassphrase`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"certstatus": &schema.Schema{
										Description: `Certstatus`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certxferprotocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"eap_method": &schema.Schema{
										Description: `Eap Method`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Ext Web Auth`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `Ext Web Auth Access Url`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Description: `Ext Web Auth Html Tag`,
										Type:        schema.TypeString,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `Ext Web Auth Portal`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `Ext Web Auth Virtual Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer3web Auth Email Address`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer3web Authpassword`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer3web Authsecurity`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer3web Authuser Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"num_aps": &schema.Schema{
										Description: `Num Aps`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"num_sensors": &schema.Schema{
										Description: `Num Sensors`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"psk": &schema.Schema{
										Description: `Psk`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"qos_policy": &schema.Schema{
										Description: `Qos Policy`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"scep": &schema.Schema{
										Description: `Scep`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"status": &schema.Schema{
										Description: `Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Description: `Config`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `Selected`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"valid_from": &schema.Schema{
										Description: `Valid From`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"valid_to": &schema.Schema{
										Description: `Valid To`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"white_list": &schema.Schema{
										Description: `White List`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Description: `Wlan Id`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"wlc": &schema.Schema{
										Description: `Wlc`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"test_duration_estimate": &schema.Schema{
							Description: `Test Duration Estimate`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"test_schedule_mode": &schema.Schema{
							Description: `Test Schedule Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"test_template": &schema.Schema{
							Description: `Test Template`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tests": &schema.Schema{
							Description: `Tests`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"wlans": &schema.Schema{
							Description: `Wlans`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"new_template_name": &schema.Schema{
				Description: `New Template Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"template_name": &schema.Schema{
				Description: `Template Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSensorTestTemplateDuplicateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DuplicateSensorTestTemplate")
		request1 := expandRequestSensorTestTemplateDuplicateDuplicateSensorTestTemplate(ctx, "", d)

		response1, restyResp1, err := client.Sensors.DuplicateSensorTestTemplate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DuplicateSensorTestTemplate", err,
				"Failure at DuplicateSensorTestTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsDuplicateSensorTestTemplateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DuplicateSensorTestTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSensorTestTemplateDuplicateDuplicateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsDuplicateSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsDuplicateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_name")))) {
		request.TemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_template_name")))) {
		request.NewTemplateName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSensorsDuplicateSensorTestTemplateItem(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["model_version"] = item.ModelVersion
	respItem["start_time"] = item.StartTime
	respItem["last_modified_time"] = item.LastModifiedTime
	respItem["num_associated_sensor"] = item.NumAssociatedSensor
	respItem["location"] = flattenSensorsDuplicateSensorTestTemplateItemLocation(item.Location)
	respItem["site_hierarchy"] = flattenSensorsDuplicateSensorTestTemplateItemSiteHierarchy(item.SiteHierarchy)
	respItem["status"] = item.Status
	respItem["connection"] = item.Connection
	respItem["frequency"] = flattenSensorsDuplicateSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = flattenSensorsDuplicateSensorTestTemplateItemWLANs(item.WLANs)
	respItem["ssids"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsDuplicateSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["schedule"] = flattenSensorsDuplicateSensorTestTemplateItemSchedule(item.Schedule)
	respItem["tests"] = flattenSensorsDuplicateSensorTestTemplateItemTests(item.Tests)
	respItem["sensors"] = flattenSensorsDuplicateSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsDuplicateSensorTestTemplateItemApCoverage(item.ApCoverage)
	respItem["test_duration_estimate"] = item.TestDurationEstimate
	respItem["test_template"] = boolPtrToString(item.TestTemplate)
	respItem["legacy_test_suite"] = boolPtrToString(item.LegacyTestSuite)
	respItem["tenant_id"] = flattenSensorsDuplicateSensorTestTemplateItemTenantID(item.TenantID)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsDuplicateSensorTestTemplateItemLocation(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSiteHierarchy(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemFrequency(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemWLANs(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseWLANs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDs(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsBands(item.Bands)
		respItem["ssid"] = item.SSID
		respItem["profile_name"] = item.ProfileName
		respItem["auth_type"] = item.AuthType
		respItem["auth_type_rcvd"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsAuthTypeRcvd(item.AuthTypeRcvd)
		respItem["psk"] = item.Psk
		respItem["username"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsUsername(item.Username)
		respItem["password"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsPassword(item.Password)
		respItem["eap_method"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsEapMethod(item.EapMethod)
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsAuthProtocol(item.AuthProtocol)
		respItem["certfilename"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertfilename(item.Certfilename)
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertpassphrase(item.Certpassphrase)
		respItem["certdownloadurl"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertdownloadurl(item.Certdownloadurl)
		respItem["num_aps"] = item.NumAps
		respItem["num_sensors"] = item.NumSensors
		respItem["layer3web_authsecurity"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item.Layer3WebAuthsecurity)
		respItem["layer3web_authuser_name"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item.Layer3WebAuthuserName)
		respItem["layer3web_authpassword"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item.Layer3WebAuthpassword)
		respItem["ext_web_auth_virtual_ip"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item.ExtWebAuthVirtualIP)
		respItem["layer3web_auth_email_address"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item.Layer3WebAuthEmailAddress)
		respItem["qos_policy"] = item.QosPolicy
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthPortal(item.ExtWebAuthPortal)
		respItem["ext_web_auth_access_url"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item.ExtWebAuthAccessURL)
		respItem["ext_web_auth_html_tag"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["third_party"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
		respItem["id"] = item.ID
		respItem["wlan_id"] = item.WLANID
		respItem["wlc"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsWlc(item.Wlc)
		respItem["valid_from"] = item.ValidFrom
		respItem["valid_to"] = item.ValidTo
		respItem["status"] = item.Status
		respItem["tests"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsBands(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsBands) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsAuthTypeRcvd(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthTypeRcvd) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsUsername(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsUsername) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsPassword(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsPassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsEapMethod(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsEapMethod) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsAuthProtocol(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthProtocol) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertfilename(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertfilename) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertpassphrase(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertpassphrase) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsCertdownloadurl(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertdownloadurl) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthPortal(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthPortal) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsThirdParty(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsWlc(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsWlc) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsTests(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsTestsConfig(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemLocationInfoList(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["location_id"] = item.LocationID
		respItem["location_type"] = item.LocationType
		respItem["all_sensors"] = boolPtrToString(item.AllSensors)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["mac_address_list"] = flattenSensorsDuplicateSensorTestTemplateItemLocationInfoListMacAddressList(item.MacAddressList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemLocationInfoListMacAddressList(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoListMacAddressList) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSchedule(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["schedule_range"] = flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRange(item.ScheduleRange)
	respItem["start_time"] = item.StartTime
	respItem["frequency"] = flattenSensorsDuplicateSensorTestTemplateItemScheduleFrequency(item.Frequency)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRange(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["time_range"] = flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRangeTimeRange(item.TimeRange)
		respItem["day"] = item.Day
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRangeTimeRange(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["from"] = item.From
		respItem["to"] = item.To
		respItem["frequency"] = flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRangeTimeRangeFrequency(item.Frequency)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemScheduleScheduleRangeTimeRangeFrequency(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["value"] = item.Value
	respItem["unit"] = item.Unit

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemScheduleFrequency(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseScheduleFrequency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["value"] = item.Value
	respItem["unit"] = item.Unit

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemTests(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseTests) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSensors(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSensors) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemApCoverage(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["number_of_aps_to_test"] = item.NumberOfApsToTest
		respItem["rssi_threshold"] = item.RssiThreshold
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemTenantID(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseTenantID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
