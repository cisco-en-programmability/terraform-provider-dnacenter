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

// dataSourceAction
func dataSourceSensorTestTemplateEdit() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.

- Intent API to deploy, schedule, or edit and existing SENSOR test template
`,

		ReadContext: dataSourceSensorTestTemplateEditRead,
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
										Type:        schema.TypeList,
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
			"location_info_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"all_sensors": &schema.Schema{
							Description: `All Sensors`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"location_id": &schema.Schema{
							Description: `Location Id`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"location_type": &schema.Schema{
							Description: `Location Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"frequency": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"unit": &schema.Schema{
										Description: `Unit`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeInt,
										Optional:    true,
									},
								},
							},
						},
						"schedule_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"day": &schema.Schema{
										Description: `Day`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"time_range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"frequency": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"unit": &schema.Schema{
																Description: `Unit`,
																Type:        schema.TypeString,
																Optional:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeInt,
																Optional:    true,
															},
														},
													},
												},
												"from": &schema.Schema{
													Description: `From`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"to": &schema.Schema{
													Description: `To`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"test_schedule_mode": &schema.Schema{
							Description: `Test Schedule Mode`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"template_name": &schema.Schema{
				Description: `Template Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSensorTestTemplateEditRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: EditSensorTestTemplate")
		request1 := expandRequestSensorTestTemplateEditEditSensorTestTemplate(ctx, "", d)

		response1, restyResp1, err := client.Sensors.EditSensorTestTemplate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing EditSensorTestTemplate", err,
				"Failure at EditSensorTestTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsEditSensorTestTemplateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting EditSensorTestTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_name")))) {
		request.TemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_info_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_info_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_info_list")))) {
		request.LocationInfoList = expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoListArray(ctx, key+".location_info_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schedule")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schedule")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schedule")))) {
		request.Schedule = expandRequestSensorTestTemplateEditEditSensorTestTemplateSchedule(ctx, key+".schedule.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList{}
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
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_type")))) {
		request.LocationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_hierarchy")))) {
		request.SiteHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sensors")))) {
		request.AllSensors = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSchedule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSchedule {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSchedule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".test_schedule_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".test_schedule_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".test_schedule_mode")))) {
		request.TestScheduleMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".frequency")))) {
		request.Frequency = expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleFrequency(ctx, key+".frequency.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schedule_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schedule_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schedule_range")))) {
		request.ScheduleRange = expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeArray(ctx, key+".schedule_range", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleFrequency(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleFrequency {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleFrequency{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unit")))) {
		request.Unit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRange {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRange{}
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
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRange {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".day")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".day")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".day")))) {
		request.Day = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_range")))) {
		request.TimeRange = expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRangeArray(ctx, key+".time_range", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange{}
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
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from")))) {
		request.From = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to")))) {
		request.To = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".frequency")))) {
		request.Frequency = expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency(ctx, key+".frequency.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unit")))) {
		request.Unit = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSensorsEditSensorTestTemplateItem(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponse) []map[string]interface{} {
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
	respItem["location"] = flattenSensorsEditSensorTestTemplateItemLocation(item.Location)
	respItem["site_hierarchy"] = flattenSensorsEditSensorTestTemplateItemSiteHierarchy(item.SiteHierarchy)
	respItem["status"] = item.Status
	respItem["connection"] = item.Connection
	respItem["frequency"] = flattenSensorsEditSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = flattenSensorsEditSensorTestTemplateItemWLANs(item.WLANs)
	respItem["ssids"] = flattenSensorsEditSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsEditSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["schedule"] = flattenSensorsEditSensorTestTemplateItemSchedule(item.Schedule)
	respItem["tests"] = flattenSensorsEditSensorTestTemplateItemTests(item.Tests)
	respItem["sensors"] = flattenSensorsEditSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsEditSensorTestTemplateItemApCoverage(item.ApCoverage)
	respItem["test_duration_estimate"] = item.TestDurationEstimate
	respItem["test_template"] = boolPtrToString(item.TestTemplate)
	respItem["legacy_test_suite"] = boolPtrToString(item.LegacyTestSuite)
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsEditSensorTestTemplateItemLocation(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSiteHierarchy(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemFrequency(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemWLANs(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseWLANs) []interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSSIDs(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = flattenSensorsEditSensorTestTemplateItemSSIDsBands(item.Bands)
		respItem["ssid"] = item.SSID
		respItem["profile_name"] = item.ProfileName
		respItem["auth_type"] = item.AuthType
		respItem["auth_type_rcvd"] = flattenSensorsEditSensorTestTemplateItemSSIDsAuthTypeRcvd(item.AuthTypeRcvd)
		respItem["psk"] = item.Psk
		respItem["username"] = flattenSensorsEditSensorTestTemplateItemSSIDsUsername(item.Username)
		respItem["password"] = flattenSensorsEditSensorTestTemplateItemSSIDsPassword(item.Password)
		respItem["eap_method"] = flattenSensorsEditSensorTestTemplateItemSSIDsEapMethod(item.EapMethod)
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = flattenSensorsEditSensorTestTemplateItemSSIDsAuthProtocol(item.AuthProtocol)
		respItem["certfilename"] = flattenSensorsEditSensorTestTemplateItemSSIDsCertfilename(item.Certfilename)
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = flattenSensorsEditSensorTestTemplateItemSSIDsCertpassphrase(item.Certpassphrase)
		respItem["certdownloadurl"] = flattenSensorsEditSensorTestTemplateItemSSIDsCertdownloadurl(item.Certdownloadurl)
		respItem["num_aps"] = item.NumAps
		respItem["num_sensors"] = item.NumSensors
		respItem["layer3web_authsecurity"] = flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item.Layer3WebAuthsecurity)
		respItem["layer3web_authuser_name"] = flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item.Layer3WebAuthuserName)
		respItem["layer3web_authpassword"] = flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item.Layer3WebAuthpassword)
		respItem["ext_web_auth_virtual_ip"] = flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item.ExtWebAuthVirtualIP)
		respItem["layer3web_auth_email_address"] = flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item.Layer3WebAuthEmailAddress)
		respItem["qos_policy"] = item.QosPolicy
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthPortal(item.ExtWebAuthPortal)
		respItem["ext_web_auth_access_url"] = flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item.ExtWebAuthAccessURL)
		respItem["ext_web_auth_html_tag"] = flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["third_party"] = flattenSensorsEditSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
		respItem["id"] = item.ID
		respItem["wlan_id"] = item.WLANID
		respItem["wlc"] = flattenSensorsEditSensorTestTemplateItemSSIDsWlc(item.Wlc)
		respItem["valid_from"] = item.ValidFrom
		respItem["valid_to"] = item.ValidTo
		respItem["status"] = item.Status
		respItem["tests"] = flattenSensorsEditSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemSSIDsBands(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsBands) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsAuthTypeRcvd(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthTypeRcvd) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsUsername(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsUsername) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsPassword(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsPassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsEapMethod(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsEapMethod) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsAuthProtocol(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthProtocol) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsCertfilename(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsCertfilename) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsCertpassphrase(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsCertpassphrase) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsCertdownloadurl(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsCertdownloadurl) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthuserName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthpassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthPortal(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthPortal) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthAccessURL) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSSIDsThirdParty(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsEditSensorTestTemplateItemSSIDsWlc(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsWlc) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSSIDsTests(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsEditSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemSSIDsTestsConfig(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig) []interface{} {
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

func flattenSensorsEditSensorTestTemplateItemLocationInfoList(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseLocationInfoList) []map[string]interface{} {
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
		respItem["mac_address_list"] = flattenSensorsEditSensorTestTemplateItemLocationInfoListMacAddressList(item.MacAddressList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemLocationInfoListMacAddressList(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseLocationInfoListMacAddressList) []interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSchedule(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSchedule) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["schedule_range"] = flattenSensorsEditSensorTestTemplateItemScheduleScheduleRange(item.ScheduleRange)
	respItem["start_time"] = item.StartTime
	respItem["frequency"] = flattenSensorsEditSensorTestTemplateItemScheduleFrequency(item.Frequency)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsEditSensorTestTemplateItemScheduleScheduleRange(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["time_range"] = flattenSensorsEditSensorTestTemplateItemScheduleScheduleRangeTimeRange(item.TimeRange)
		respItem["day"] = item.Day
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemScheduleScheduleRangeTimeRange(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["from"] = item.From
		respItem["to"] = item.To
		respItem["frequency"] = flattenSensorsEditSensorTestTemplateItemScheduleScheduleRangeTimeRangeFrequency(item.Frequency)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemScheduleScheduleRangeTimeRangeFrequency(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemScheduleFrequency(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseScheduleFrequency) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemTests(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseTests) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSensors(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSensors) []interface{} {
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

func flattenSensorsEditSensorTestTemplateItemApCoverage(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseApCoverage) []map[string]interface{} {
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
