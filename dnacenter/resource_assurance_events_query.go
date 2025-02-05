package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAssuranceEventsQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API
Support Documentation' section to understand which fields are supported. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml
`,

		CreateContext: resourceAssuranceEventsQueryCreate,
		ReadContext:   resourceAssuranceEventsQueryRead,
		DeleteContext: resourceAssuranceEventsQueryDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"affected_clients": &schema.Schema{
										Description: `Affected Clients`,
										Type:        schema.TypeList,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ap_radio_operation_state": &schema.Schema{
										Description: `Ap Radio Operation State`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ap_role": &schema.Schema{
										Description: `Ap Role`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ap_switch_id": &schema.Schema{
										Description: `Ap Switch Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ap_switch_name": &schema.Schema{
										Description: `Ap Switch Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"assoc_rssi": &schema.Schema{
										Description: `Assoc Rssi`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"assoc_snr": &schema.Schema{
										Description: `Assoc Snr`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"audit_session_id": &schema.Schema{
										Description: `Audit Session Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"auth_server_ip": &schema.Schema{
										Description: `Auth Server Ip`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"candidate_a_ps": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"apid": &schema.Schema{
													Description: `Ap Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_mac": &schema.Schema{
													Description: `Ap Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_name": &schema.Schema{
													Description: `Ap Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"bssid": &schema.Schema{
													Description: `Bssid`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"rssi": &schema.Schema{
													Description: `Rssi`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"child_events": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"details": &schema.Schema{
													Description: `Details`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"failure_category": &schema.Schema{
													Description: `Failure Category`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"reason_code": &schema.Schema{
													Description: `Reason Code`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"reason_description": &schema.Schema{
													Description: `Reason Description`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"result_status": &schema.Schema{
													Description: `Result Status`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"sub_reason_code": &schema.Schema{
													Description: `Sub Reason Code`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"sub_reason_description": &schema.Schema{
													Description: `Sub Reason Description`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"timestamp": &schema.Schema{
													Description: `Timestamp`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"wireless_event_type": &schema.Schema{
													Description: `Wireless Event Type`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"client_mac": &schema.Schema{
										Description: `Client Mac`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"connected_interface_name": &schema.Schema{
										Description: `Connected Interface Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"current_radio_power_level": &schema.Schema{
										Description: `Current Radio Power Level`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"details": &schema.Schema{
										Description: `Details`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_family": &schema.Schema{
										Description: `Device Family`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"dhcp_server_ip": &schema.Schema{
										Description: `Dhcp Server Ip`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"duid": &schema.Schema{
										Description: `Duid`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"event_status": &schema.Schema{
										Description: `Event Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"failure_category": &schema.Schema{
										Description: `Failure Category`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"failure_ip_address": &schema.Schema{
										Description: `Failure Ip Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"frequency": &schema.Schema{
										Description: `Frequency`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"identifier": &schema.Schema{
										Description: `Identifier`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"invalid_ie_a_ps": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"apid": &schema.Schema{
													Description: `Ap Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_mac": &schema.Schema{
													Description: `Ap Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_name": &schema.Schema{
													Description: `Ap Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"bssid": &schema.Schema{
													Description: `Bssid`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"frame_type": &schema.Schema{
													Description: `Frame Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ies": &schema.Schema{
													Description: `Ies`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ipv4": &schema.Schema{
										Description: `Ipv4`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv6": &schema.Schema{
										Description: `Ipv6`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"is_private_mac": &schema.Schema{
										Description: `Is Private Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"last_ap_disconnect_reason": &schema.Schema{
										Description: `Last Ap Disconnect Reason`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"last_ap_reset_type": &schema.Schema{
										Description: `Last Ap Reset Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"management_ip_address": &schema.Schema{
										Description: `Management Ip Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"message_type": &schema.Schema{
										Description: `Message Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"missing_response_a_ps": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"apid": &schema.Schema{
													Description: `Ap Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_mac": &schema.Schema{
													Description: `Ap Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_name": &schema.Schema{
													Description: `Ap Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"bssid": &schema.Schema{
													Description: `Bssid`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"frame_type": &schema.Schema{
													Description: `Frame Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network Device Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"network_device_name": &schema.Schema{
										Description: `Network Device Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"new_radio_channel_list": &schema.Schema{
										Description: `New Radio Channel List`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"new_radio_channel_width": &schema.Schema{
										Description: `New Radio Channel Width`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"old_radio_channel_list": &schema.Schema{
										Description: `Old Radio Channel List`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"old_radio_channel_width": &schema.Schema{
										Description: `Old Radio Channel Width`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"previous_radio_power_level": &schema.Schema{
										Description: `Previous Radio Power Level`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"radio_channel_slot": &schema.Schema{
										Description: `Radio Channel Slot`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"radio_channel_utilization": &schema.Schema{
										Description: `Radio Channel Utilization`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"radio_interference": &schema.Schema{
										Description: `Radio Interference`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"radio_noise": &schema.Schema{
										Description: `Radio Noise`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"reason_description": &schema.Schema{
										Description: `Reason Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"replaced_device_serial_number": &schema.Schema{
										Description: `Replaced Device Serial Number`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"replacing_device_serial_number": &schema.Schema{
										Description: `Replacing Device Serial Number`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"result_status": &schema.Schema{
										Description: `Result Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"roam_type": &schema.Schema{
										Description: `Roam Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy_id": &schema.Schema{
										Description: `Site Hierarchy Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"sub_reason_description": &schema.Schema{
										Description: `Sub Reason Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"switch_number": &schema.Schema{
										Description: `Switch Number`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"udn_id": &schema.Schema{
										Description: `Udn Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"udn_name": &schema.Schema{
										Description: `Udn Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"wireless_client_event_end_time": &schema.Schema{
										Description: `Wireless Client Event End Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"wireless_client_event_start_time": &schema.Schema{
										Description: `Wireless Client Event Start Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"wlc_id": &schema.Schema{
										Description: `Wlc Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"wlc_name": &schema.Schema{
										Description: `Wlc Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"offset": &schema.Schema{
										Description: `Offset`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"order": &schema.Schema{
													Description: `Order`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
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
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"views": &schema.Schema{
							Description: `Views`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceAssuranceEventsQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFilters(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.QueryAssuranceEventsWithFiltersHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.QueryAssuranceEventsWithFilters(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing QueryAssuranceEventsWithFilters", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesQueryAssuranceEventsWithFiltersItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting QueryAssuranceEventsWithFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceAssuranceEventsQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssuranceEventsQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFilters {
	request := dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_family")))) {
		request.DeviceFamily = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersFilters {
	request := []dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersFilters{}
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
		i := expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersFilters {
	request := dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPage {
	request := dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy {
	request := []dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy{}
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
		i := expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceEventsQueryQueryAssuranceEventsWithFiltersPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy {
	request := dnacentersdkgo.RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesQueryAssuranceEventsWithFiltersItems(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsWithFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["old_radio_channel_width"] = item.OldRadioChannelWidth
		respItem["client_mac"] = item.ClientMac
		respItem["switch_number"] = item.SwitchNumber
		respItem["assoc_rssi"] = item.AssocRssi
		respItem["affected_clients"] = item.AffectedClients
		respItem["is_private_mac"] = boolPtrToString(item.IsPrivateMac)
		respItem["frequency"] = item.Frequency
		respItem["ap_role"] = item.ApRole
		respItem["replacing_device_serial_number"] = item.ReplacingDeviceSerialNumber
		respItem["message_type"] = item.MessageType
		respItem["failure_category"] = item.FailureCategory
		respItem["ap_switch_name"] = item.ApSwitchName
		respItem["ap_switch_id"] = item.ApSwitchID
		respItem["radio_channel_utilization"] = item.RadioChannelUtilization
		respItem["mnemonic"] = item.Mnemonic
		respItem["radio_channel_slot"] = item.RadioChannelSlot
		respItem["details"] = item.Details
		respItem["id"] = item.ID
		respItem["last_ap_disconnect_reason"] = item.LastApDisconnectReason
		respItem["network_device_name"] = item.NetworkDeviceName
		respItem["identifier"] = item.IDentifier
		respItem["reason_description"] = item.ReasonDescription
		respItem["vlan_id"] = item.VLANID
		respItem["udn_id"] = item.UdnID
		respItem["audit_session_id"] = item.AuditSessionID
		respItem["ap_mac"] = item.ApMac
		respItem["device_family"] = item.DeviceFamily
		respItem["radio_noise"] = item.RadioNoise
		respItem["wlc_name"] = item.WlcName
		respItem["ap_radio_operation_state"] = item.ApRadioOperationState
		respItem["name"] = item.Name
		respItem["failure_ip_address"] = item.FailureIPAddress
		respItem["new_radio_channel_list"] = item.NewRadioChannelList
		respItem["duid"] = item.Duid
		respItem["roam_type"] = item.RoamType
		respItem["candidate_a_ps"] = flattenDevicesQueryAssuranceEventsWithFiltersItemsCandidateAPs(item.CandidateAPs)
		respItem["replaced_device_serial_number"] = item.ReplacedDeviceSerialNumber
		respItem["old_radio_channel_list"] = item.OldRadioChannelList
		respItem["ssid"] = item.SSID
		respItem["sub_reason_description"] = item.SubReasonDescription
		respItem["wireless_client_event_end_time"] = item.WirelessClientEventEndTime
		respItem["ipv4"] = item.IPv4
		respItem["wlc_id"] = item.WlcID
		respItem["ipv6"] = item.IPv6
		respItem["missing_response_a_ps"] = flattenDevicesQueryAssuranceEventsWithFiltersItemsMissingResponseAPs(item.MissingResponseAPs)
		respItem["timestamp"] = item.Timestamp
		respItem["severity"] = item.Severity
		respItem["current_radio_power_level"] = item.CurrentRadioPowerLevel
		respItem["new_radio_channel_width"] = item.NewRadioChannelWidth
		respItem["assoc_snr"] = item.AssocSnr
		respItem["auth_server_ip"] = item.AuthServerIP
		respItem["child_events"] = flattenDevicesQueryAssuranceEventsWithFiltersItemsChildEvents(item.ChildEvents)
		respItem["connected_interface_name"] = item.ConnectedInterfaceName
		respItem["dhcp_server_ip"] = item.DhcpServerIP
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["previous_radio_power_level"] = item.PreviousRadioPowerLevel
		respItem["result_status"] = item.ResultStatus
		respItem["radio_interference"] = item.RadioInterference
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["event_status"] = item.EventStatus
		respItem["wireless_client_event_start_time"] = item.WirelessClientEventStartTime
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["udn_name"] = item.UdnName
		respItem["facility"] = item.Facility
		respItem["last_ap_reset_type"] = item.LastApResetType
		respItem["invalid_ie_a_ps"] = flattenDevicesQueryAssuranceEventsWithFiltersItemsInvalidIeAPs(item.InvalidIeAPs)
		respItem["username"] = item.Username
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryAssuranceEventsWithFiltersItemsCandidateAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsWithFiltersResponseCandidateAPs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["apid"] = item.APID
		respItem["ap_name"] = item.ApName
		respItem["ap_mac"] = item.ApMac
		respItem["bssid"] = item.Bssid
		respItem["rssi"] = item.Rssi
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryAssuranceEventsWithFiltersItemsMissingResponseAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsWithFiltersResponseMissingResponseAPs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["apid"] = item.APID
		respItem["ap_name"] = item.ApName
		respItem["ap_mac"] = item.ApMac
		respItem["bssid"] = item.Bssid
		respItem["type"] = item.Type
		respItem["frame_type"] = item.FrameType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryAssuranceEventsWithFiltersItemsChildEvents(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsWithFiltersResponseChildEvents) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["timestamp"] = item.Timestamp
		respItem["wireless_event_type"] = item.WirelessEventType
		respItem["details"] = item.Details
		respItem["reason_code"] = item.ReasonCode
		respItem["reason_description"] = item.ReasonDescription
		respItem["sub_reason_code"] = item.SubReasonCode
		respItem["sub_reason_description"] = item.SubReasonDescription
		respItem["result_status"] = item.ResultStatus
		respItem["failure_category"] = item.FailureCategory
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryAssuranceEventsWithFiltersItemsInvalidIeAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsWithFiltersResponseInvalidIeAPs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["apid"] = item.APID
		respItem["ap_name"] = item.ApName
		respItem["ap_mac"] = item.ApMac
		respItem["bssid"] = item.Bssid
		respItem["type"] = item.Type
		respItem["frame_type"] = item.FrameType
		respItem["ies"] = item.Ies
		respItems = append(respItems, respItem)
	}
	return respItems
}
