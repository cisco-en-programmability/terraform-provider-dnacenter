package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDevicesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		CreateContext: resourceNetworkDevicesQueryCreate,
		ReadContext:   resourceNetworkDevicesQueryRead,
		DeleteContext: resourceNetworkDevicesQueryDelete,
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
						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
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

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
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
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ap_details": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_state": &schema.Schema{
													Description: `Admin State`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_group": &schema.Schema{
													Description: `Ap Group`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_operational_state": &schema.Schema{
													Description: `Ap Operational State`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_type": &schema.Schema{
													Description: `Ap Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_time": &schema.Schema{
													Description: `Connected Time`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"connected_wlc_name": &schema.Schema{
													Description: `Connected Wlc Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"ethernet_mac": &schema.Schema{
													Description: `Ethernet Mac`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"flex_group": &schema.Schema{
													Description: `Flex Group`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"home_ap_enabled": &schema.Schema{
													Description: `Home Ap Enabled`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"icap_capability": &schema.Schema{
													Description: `Icap Capability`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"led_flash_enabled": &schema.Schema{
													Description: `Led Flash Enabled`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"led_flash_seconds": &schema.Schema{
													Description: `Led Flash Seconds`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"operational_mode": &schema.Schema{
													Description: `Operational Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"policy_tag_name": &schema.Schema{
													Description: `Policy Tag Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_calendar_profile": &schema.Schema{
													Description: `Power Calendar Profile`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_mode": &schema.Schema{
													Description: `Power Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_profile": &schema.Schema{
													Description: `Power Profile`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_save_mode": &schema.Schema{
													Description: `Power Save Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_save_mode_capable": &schema.Schema{
													Description: `Power Save Mode Capable`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"power_type": &schema.Schema{
													Description: `Power Type`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"radios": &schema.Schema{
													Type:     schema.TypeList,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"air_quality": &schema.Schema{
																Description: `Air Quality`,
																Type:        schema.TypeFloat,
																ForceNew:    true,
																Computed:    true,
															},
															"band": &schema.Schema{
																Description: `Band`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
															"client_count": &schema.Schema{
																Description: `Client Count`,
																Type:        schema.TypeInt,
																ForceNew:    true,
																Computed:    true,
															},
															"id": &schema.Schema{
																Description: `Id`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
															"interference": &schema.Schema{
																Description: `Interference`,
																Type:        schema.TypeFloat,
																ForceNew:    true,
																Computed:    true,
															},
															"noise": &schema.Schema{
																Description: `Noise`,
																Type:        schema.TypeInt,
																ForceNew:    true,
																Computed:    true,
															},
															"traffic_util": &schema.Schema{
																Description: `Traffic Util`,
																Type:        schema.TypeInt,
																ForceNew:    true,
																Computed:    true,
															},
															"utilization": &schema.Schema{
																Description: `Utilization`,
																Type:        schema.TypeFloat,
																ForceNew:    true,
																Computed:    true,
															},
														},
													},
												},
												"regulatory_domain": &schema.Schema{
													Description: `Regulatory Domain`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"reset_reason": &schema.Schema{
													Description: `Reset Reason`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"rf_tag_name": &schema.Schema{
													Description: `Rf Tag Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"site_tag_name": &schema.Schema{
													Description: `Site Tag Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"sub_mode": &schema.Schema{
													Description: `Sub Mode`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"client_count": &schema.Schema{
										Description: `Client Count`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"collection_status": &schema.Schema{
										Description: `Collection Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"communication_state": &schema.Schema{
										Description: `Communication State`,
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
									"device_group_hierarchy_id": &schema.Schema{
										Description: `Device Group Hierarchy Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_role": &schema.Schema{
										Description: `Device Role`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_series": &schema.Schema{
										Description: `Device Series`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"fabric_details": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fabric_role": &schema.Schema{
													Description: `Fabric Role`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"fabric_site_name": &schema.Schema{
													Description: `Fabric Site Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"transit_fabrics": &schema.Schema{
													Description: `Transit Fabrics`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"feature_flag_list": &schema.Schema{
										Description: `Feature Flag List`,
										Type:        schema.TypeList,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ha_last_reset_reason": &schema.Schema{
										Description: `Ha Last Reset Reason`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ha_status": &schema.Schema{
										Description: `Ha Status`,
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
									"ipv4_address": &schema.Schema{
										Description: `Ipv4 Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv6_address": &schema.Schema{
										Description: `Ipv6 Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"last_boot_time": &schema.Schema{
										Description: `Last Boot Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"mac_address": &schema.Schema{
										Description: `Mac Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"maintenance_mode_enabled": &schema.Schema{
										Description: `Maintenance Mode Enabled`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"management_ip_address": &schema.Schema{
										Description: `Management Ip Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"metrics_details": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"air_quality_score": &schema.Schema{
													Description: `Air Quality Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"ap_count": &schema.Schema{
													Description: `Ap Count`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"avg_temperature": &schema.Schema{
													Description: `Avg Temperature`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"cpu_score": &schema.Schema{
													Description: `Cpu Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"cpu_utilization": &schema.Schema{
													Description: `Cpu Utilization`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"discard_interfaces": &schema.Schema{
													Description: `Discard Interfaces`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"discard_score": &schema.Schema{
													Description: `Discard Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"error_interfaces": &schema.Schema{
													Description: `Error Interfaces`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"error_score": &schema.Schema{
													Description: `Error Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"free_memory_buffer": &schema.Schema{
													Description: `Free Memory Buffer`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"free_memory_buffer_score": &schema.Schema{
													Description: `Free Memory Buffer Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"free_timer": &schema.Schema{
													Description: `Free Timer`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"free_timer_score": &schema.Schema{
													Description: `Free Timer Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"high_link_utilization_interfaces": &schema.Schema{
													Description: `High Link Utilization Interfaces`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"inter_device_connected_down_interfaces": &schema.Schema{
													Description: `Inter Device Connected Down Interfaces`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"inter_device_link_score": &schema.Schema{
													Description: `Inter Device Link Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"interference_score": &schema.Schema{
													Description: `Interference Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"link_utilization_score": &schema.Schema{
													Description: `Link Utilization Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"max_temperature": &schema.Schema{
													Description: `Max Temperature`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"memory_score": &schema.Schema{
													Description: `Memory Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"memory_utilization": &schema.Schema{
													Description: `Memory Utilization`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
												"noise_score": &schema.Schema{
													Description: `Noise Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"overall_fabric_score": &schema.Schema{
													Description: `Overall Fabric Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"overall_health_score": &schema.Schema{
													Description: `Overall Health Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"packet_pool": &schema.Schema{
													Description: `Packet Pool`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"packet_pool_score": &schema.Schema{
													Description: `Packet Pool Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"utilization_score": &schema.Schema{
													Description: `Utilization Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"wqe_pool": &schema.Schema{
													Description: `Wqe Pool`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"wqe_pool_score": &schema.Schema{
													Description: `Wqe Pool Score`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"os_type": &schema.Schema{
										Description: `Os Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"platform_id": &schema.Schema{
										Description: `Platform Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"port_count": &schema.Schema{
										Description: `Port Count`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"product_vendor": &schema.Schema{
										Description: `Product Vendor`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"redundancy_mode": &schema.Schema{
										Description: `Redundancy Mode`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"redundancy_peer_state": &schema.Schema{
										Description: `Redundancy Peer State`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"redundancy_peer_state_derived": &schema.Schema{
										Description: `Redundancy Peer State Derived`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"redundancy_state": &schema.Schema{
										Description: `Redundancy State`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"redundancy_state_derived": &schema.Schema{
										Description: `Redundancy State Derived`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ring_status": &schema.Schema{
										Description: `Ring Status`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"serial_number": &schema.Schema{
										Description: `Serial Number`,
										Type:        schema.TypeString,
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
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"software_version": &schema.Schema{
										Description: `Software Version`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"stack_type": &schema.Schema{
										Description: `Stack Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"tag_names": &schema.Schema{
										Description: `Tag Names`,
										Type:        schema.TypeList,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"up_time": &schema.Schema{
										Description: `Up Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"wired_client_count": &schema.Schema{
										Description: `Wired Client Count`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"wireless_client_count": &schema.Schema{
										Description: `Wireless Client Count`,
										Type:        schema.TypeInt,
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

									"count": &schema.Schema{
										Description: `Count`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
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
										Description: `Sort By`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
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

func resourceNetworkDevicesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters{}
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
		i := expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters{}
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

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes{}
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
		i := expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesQueryGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".count")))) {
		request.Count = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItems(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["platform_id"] = item.PlatformID
		respItem["device_family"] = item.DeviceFamily
		respItem["serial_number"] = item.SerialNumber
		respItem["mac_address"] = item.MacAddress
		respItem["device_series"] = item.DeviceSeries
		respItem["software_version"] = item.SoftwareVersion
		respItem["product_vendor"] = item.ProductVendor
		respItem["device_role"] = item.DeviceRole
		respItem["device_type"] = item.DeviceType
		respItem["communication_state"] = item.CommunicationState
		respItem["collection_status"] = item.CollectionStatus
		respItem["ha_status"] = item.HaStatus
		respItem["last_boot_time"] = item.LastBootTime
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_id"] = item.SiteID
		respItem["device_group_hierarchy_id"] = item.DeviceGroupHierarchyID
		respItem["tag_names"] = item.TagNames
		respItem["stack_type"] = item.StackType
		respItem["os_type"] = item.OsType
		respItem["ring_status"] = boolPtrToString(item.RingStatus)
		respItem["maintenance_mode_enabled"] = boolPtrToString(item.MaintenanceModeEnabled)
		respItem["up_time"] = item.UpTime
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_address"] = item.IPv6Address
		respItem["redundancy_mode"] = item.RedundancyMode
		respItem["feature_flag_list"] = item.FeatureFlagList
		respItem["ha_last_reset_reason"] = item.HaLastResetReason
		respItem["redundancy_peer_state_derived"] = item.RedundancyPeerStateDerived
		respItem["redundancy_peer_state"] = item.RedundancyPeerState
		respItem["redundancy_state_derived"] = item.RedundancyStateDerived
		respItem["redundancy_state"] = item.RedundancyState
		respItem["wired_client_count"] = item.WiredClientCount
		respItem["wireless_client_count"] = item.WirelessClientCount
		respItem["port_count"] = item.PortCount
		respItem["client_count"] = item.ClientCount
		respItem["ap_details"] = flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsApDetails(item.ApDetails)
		respItem["metrics_details"] = flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsMetricsDetails(item.MetricsDetails)
		respItem["fabric_details"] = flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsFabricDetails(item.FabricDetails)
		respItem["aggregate_attributes"] = flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsApDetails(item *dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connected_wlc_name"] = item.ConnectedWlcName
	respItem["policy_tag_name"] = item.PolicyTagName
	respItem["ap_operational_state"] = item.ApOperationalState
	respItem["power_save_mode"] = item.PowerSaveMode
	respItem["operational_mode"] = item.OperationalMode
	respItem["reset_reason"] = item.ResetReason
	respItem["protocol"] = item.Protocol
	respItem["power_mode"] = item.PowerMode
	respItem["connected_time"] = item.ConnectedTime
	respItem["led_flash_enabled"] = boolPtrToString(item.LedFlashEnabled)
	respItem["led_flash_seconds"] = item.LedFlashSeconds
	respItem["sub_mode"] = item.SubMode
	respItem["home_ap_enabled"] = boolPtrToString(item.HomeApEnabled)
	respItem["power_type"] = item.PowerType
	respItem["ap_type"] = item.ApType
	respItem["admin_state"] = item.AdminState
	respItem["icap_capability"] = item.IcapCapability
	respItem["regulatory_domain"] = item.RegulatoryDomain
	respItem["ethernet_mac"] = item.EthernetMac
	respItem["rf_tag_name"] = item.RfTagName
	respItem["site_tag_name"] = item.SiteTagName
	respItem["power_save_mode_capable"] = item.PowerSaveModeCapable
	respItem["power_profile"] = item.PowerProfile
	respItem["flex_group"] = item.FlexGroup
	respItem["power_calendar_profile"] = item.PowerCalendarProfile
	respItem["ap_group"] = item.ApGroup
	respItem["radios"] = flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsApDetailsRadios(item.Radios)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsApDetailsRadios(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetailsRadios) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["band"] = item.Band
		respItem["noise"] = item.Noise
		respItem["air_quality"] = item.AirQuality
		respItem["interference"] = item.Interference
		respItem["traffic_util"] = item.TrafficUtil
		respItem["utilization"] = item.Utilization
		respItem["client_count"] = item.ClientCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsMetricsDetails(item *dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseMetricsDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["overall_health_score"] = item.OverallHealthScore
	respItem["overall_fabric_score"] = item.OverallFabricScore
	respItem["cpu_utilization"] = item.CPUUtilization
	respItem["cpu_score"] = item.CPUScore
	respItem["memory_utilization"] = item.MemoryUtilization
	respItem["memory_score"] = item.MemoryScore
	respItem["avg_temperature"] = item.AvgTemperature
	respItem["max_temperature"] = item.MaxTemperature
	respItem["discard_score"] = item.DiscardScore
	respItem["discard_interfaces"] = item.DiscardInterfaces
	respItem["error_score"] = item.ErrorScore
	respItem["error_interfaces"] = item.ErrorInterfaces
	respItem["inter_device_link_score"] = item.InterDeviceLinkScore
	respItem["inter_device_connected_down_interfaces"] = item.InterDeviceConnectedDownInterfaces
	respItem["link_utilization_score"] = item.LinkUtilizationScore
	respItem["high_link_utilization_interfaces"] = item.HighLinkUtilizationInterfaces
	respItem["free_timer_score"] = item.FreeTimerScore
	respItem["free_timer"] = item.FreeTimer
	respItem["packet_pool_score"] = item.PacketPoolScore
	respItem["packet_pool"] = item.PacketPool
	respItem["free_memory_buffer_score"] = item.FreeMemoryBufferScore
	respItem["free_memory_buffer"] = item.FreeMemoryBuffer
	respItem["wqe_pool_score"] = item.WqePoolScore
	respItem["wqe_pool"] = item.WqePool
	respItem["ap_count"] = item.ApCount
	respItem["noise_score"] = item.NoiseScore
	respItem["utilization_score"] = item.UtilizationScore
	respItem["interference_score"] = item.InterferenceScore
	respItem["air_quality_score"] = item.AirQualityScore

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsFabricDetails(item *dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["fabric_role"] = item.FabricRole
	respItem["fabric_site_name"] = item.FabricSiteName
	respItem["transit_fabrics"] = item.TransitFabrics

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["function"] = item.Function
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
