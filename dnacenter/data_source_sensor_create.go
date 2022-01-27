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
func dataSourceSensorCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID
`,

		ReadContext: dataSourceSensorCreateRead,
		Schema: map[string]*schema.Schema{
			"ap_coverage": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bands": &schema.Schema{
							Description: `Bands`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"number_of_aps_to_test": &schema.Schema{
							Description: `Number Of Aps To Test`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"rssi_threshold": &schema.Schema{
							Description: `Rssi Threshold`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"r_connection": &schema.Schema{
				Description: `r_connection`,
				Type:        schema.TypeString,
				Optional:    true,
			},
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
						"r_connection": &schema.Schema{
							Description: `r_connection`,
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
							Description: `Location Info List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
							Description: `Schedule`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"schedule_in_days": &schema.Schema{
							Description: `Schedule In Days`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"sensors": &schema.Schema{
							Description: `Sensors`,
							Type:        schema.TypeString,
							Computed:    true,
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
			"model_version": &schema.Schema{
				Description: `Model Version`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `Name`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ssids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_type": &schema.Schema{
							Description: `Auth Type`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"categories": &schema.Schema{
							Description: `Categories`,
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"profile_name": &schema.Schema{
							Description: `Profile Name`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"psk": &schema.Schema{
							Description: `Psk`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"qos_policy": &schema.Schema{
							Description: `Qos Policy`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"tests": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config": &schema.Schema{
										Description: `Config`,
										Type:        schema.TypeList,
										Optional:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"third_party": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"selected": &schema.Schema{
										Description: `Selected`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
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

func dataSourceSensorCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateSensorTestTemplate")
		request1 := expandRequestSensorCreateCreateSensorTestTemplate(ctx, "", d)

		response1, restyResp1, err := client.Sensors.CreateSensorTestTemplate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSensorTestTemplate", err,
				"Failure at CreateSensorTestTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsCreateSensorTestTemplateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSensorTestTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSensorCreateCreateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssids")))) {
		request.SSIDs = expandRequestSensorCreateCreateSensorTestTemplateSSIDsArray(ctx, key+".ssids", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".r_connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".r_connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".r_connection")))) {
		request.Connection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_coverage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_coverage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_coverage")))) {
		request.ApCoverage = expandRequestSensorCreateCreateSensorTestTemplateApCoverageArray(ctx, key+".ap_coverage", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_version")))) {
		request.ModelVersion = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
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
		i := expandRequestSensorCreateCreateSensorTestTemplateSSIDs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = expandRequestSensorCreateCreateSensorTestTemplateSSIDsThirdParty(ctx, key+".third_party.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsArray(ctx, key+".tests", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsThirdParty(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected")))) {
		request.Selected = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
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
		i := expandRequestSensorCreateCreateSensorTestTemplateSSIDsTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsConfigArray(ctx, key+".config", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
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
		i := expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateSSIDsTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	var request dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateApCoverageArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
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
		i := expandRequestSensorCreateCreateSensorTestTemplateApCoverage(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSensorCreateCreateSensorTestTemplateApCoverage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_aps_to_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) {
		request.NumberOfApsToTest = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSensorsCreateSensorTestTemplateItem(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponse) []map[string]interface{} {
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
	respItem["location"] = flattenSensorsCreateSensorTestTemplateItemLocation(item.Location)
	respItem["site_hierarchy"] = flattenSensorsCreateSensorTestTemplateItemSiteHierarchy(item.SiteHierarchy)
	respItem["status"] = item.Status
	respItem["r_connection"] = item.Connection
	respItem["frequency"] = flattenSensorsCreateSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = flattenSensorsCreateSensorTestTemplateItemWLANs(item.WLANs)
	respItem["ssids"] = flattenSensorsCreateSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsCreateSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["schedule"] = flattenSensorsCreateSensorTestTemplateItemSchedule(item.Schedule)
	respItem["tests"] = flattenSensorsCreateSensorTestTemplateItemTests(item.Tests)
	respItem["sensors"] = flattenSensorsCreateSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsCreateSensorTestTemplateItemApCoverage(item.ApCoverage)
	respItem["test_duration_estimate"] = item.TestDurationEstimate
	respItem["test_template"] = boolPtrToString(item.TestTemplate)
	respItem["legacy_test_suite"] = boolPtrToString(item.LegacyTestSuite)
	respItem["tenant_id"] = flattenSensorsCreateSensorTestTemplateItemTenantID(item.TenantID)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsCreateSensorTestTemplateItemLocation(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSiteHierarchy(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemFrequency(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemWLANs(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseWLANs) []interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemSSIDs(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = flattenSensorsCreateSensorTestTemplateItemSSIDsBands(item.Bands)
		respItem["ssid"] = item.SSID
		respItem["profile_name"] = item.ProfileName
		respItem["auth_type"] = item.AuthType
		respItem["auth_type_rcvd"] = flattenSensorsCreateSensorTestTemplateItemSSIDsAuthTypeRcvd(item.AuthTypeRcvd)
		respItem["psk"] = item.Psk
		respItem["username"] = flattenSensorsCreateSensorTestTemplateItemSSIDsUsername(item.Username)
		respItem["password"] = flattenSensorsCreateSensorTestTemplateItemSSIDsPassword(item.Password)
		respItem["eap_method"] = flattenSensorsCreateSensorTestTemplateItemSSIDsEapMethod(item.EapMethod)
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = flattenSensorsCreateSensorTestTemplateItemSSIDsAuthProtocol(item.AuthProtocol)
		respItem["certfilename"] = flattenSensorsCreateSensorTestTemplateItemSSIDsCertfilename(item.Certfilename)
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = flattenSensorsCreateSensorTestTemplateItemSSIDsCertpassphrase(item.Certpassphrase)
		respItem["certdownloadurl"] = flattenSensorsCreateSensorTestTemplateItemSSIDsCertdownloadurl(item.Certdownloadurl)
		respItem["num_aps"] = item.NumAps
		respItem["num_sensors"] = item.NumSensors
		respItem["layer3web_authsecurity"] = flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item.Layer3WebAuthsecurity)
		respItem["layer3web_authuser_name"] = flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item.Layer3WebAuthuserName)
		respItem["layer3web_authpassword"] = flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item.Layer3WebAuthpassword)
		respItem["ext_web_auth_virtual_ip"] = flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item.ExtWebAuthVirtualIP)
		respItem["layer3web_auth_email_address"] = flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item.Layer3WebAuthEmailAddress)
		respItem["qos_policy"] = item.QosPolicy
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthPortal(item.ExtWebAuthPortal)
		respItem["ext_web_auth_access_url"] = flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item.ExtWebAuthAccessURL)
		respItem["ext_web_auth_html_tag"] = flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["third_party"] = flattenSensorsCreateSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
		respItem["id"] = item.ID
		respItem["wlan_id"] = item.WLANID
		respItem["wlc"] = flattenSensorsCreateSensorTestTemplateItemSSIDsWlc(item.Wlc)
		respItem["valid_from"] = item.ValidFrom
		respItem["valid_to"] = item.ValidTo
		respItem["status"] = item.Status
		respItem["tests"] = flattenSensorsCreateSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSSIDsBands(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsBands) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsAuthTypeRcvd(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthTypeRcvd) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsUsername(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsUsername) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsPassword(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsPassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsEapMethod(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsEapMethod) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsAuthProtocol(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthProtocol) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsCertfilename(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertfilename) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsCertpassphrase(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertpassphrase) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsCertdownloadurl(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertdownloadurl) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthsecurity(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthuserName(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthpassword(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthVirtualIP(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsLayer3WebAuthEmailAddress(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthPortal(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthPortal) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthAccessURL(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemSSIDsThirdParty(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsWlc(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsWlc) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsTests(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsCreateSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSSIDsTestsConfig(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig) []interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemLocationInfoList(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList) []interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemSchedule(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSchedule) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemTests(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseTests) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSensors(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSensors) []interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemApCoverage(items *[]dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseApCoverage) []map[string]interface{} {
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

func flattenSensorsCreateSensorTestTemplateItemTenantID(item *dnacentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseTenantID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
