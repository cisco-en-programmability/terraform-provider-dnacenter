package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Gets the Network Device details based on the provided query parameters.  When there is no start and end time specified
returns the latest device details. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml

- Returns the device data for the given device Uuid in the specified start and end time range. When there is no start
and end time specified returns the latest available data for the given Id. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		ReadContext: dataSourceNetworkDevicesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The List of Network Device model attributes. This is helps to specify the interested fields in the request.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"family": &schema.Schema{
				Description: `family query parameter. The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_score": &schema.Schema{
				Description: `healthScore query parameter. The list of entity health score categories
Examples:
healthScore=good, healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. The device Uuid
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. The macAddress of the network device or client This field supports wildcard (***) character-based search.  Ex: **AB:AB:AB** or *AB:AB:AB** or **AB:AB:AB* Examples:
*macAddress=AB:AB:AB:CD:CD:CD* (single macAddress requested)
*macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE* (multiple macAddress requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_mode": &schema.Schema{
				Description: `maintenanceMode query parameter. The device maintenanceMode status true or false
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"management_ip_address": &schema.Schema{
				Description: `managementIpAddress query parameter. The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10")
This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Description: `role query parameter. The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. The list of network device serial numbers. This field supports wildcard (***) character-based search.  Ex: **MS1SV** or *MS1SV** or **MS1SV* Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San*
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (*) character search support. E.g. **uuid*, *uuid, uuid*
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid*
Examples:
*?siteId=id1* (single id requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter. The list of network device software version This field supports wildcard (***) character-based search. Ex: **17.8** or **17.8* or *17.8** Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *startTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. The list of network device type This field supports wildcard (***) character-based search. Ex: **9407R** or **9407R* or *9407R** Examples: type=SwitchesCisco Catalyst 9407R Switch (single network device types ) type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The List of Network Device model views. Please refer to ***NetworkDeviceView*** for the supported list
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},

						"ap_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"admin_state": &schema.Schema{
										Description: `Admin State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_group": &schema.Schema{
										Description: `Ap Group`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_operational_state": &schema.Schema{
										Description: `Ap Operational State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_type": &schema.Schema{
										Description: `Ap Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_time": &schema.Schema{
										Description: `Connected Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"connected_wlc_name": &schema.Schema{
										Description: `Connected Wlc Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ethernet_mac": &schema.Schema{
										Description: `Ethernet Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"flex_group": &schema.Schema{
										Description: `Flex Group`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"home_ap_enabled": &schema.Schema{
										Description: `Home Ap Enabled`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"icap_capability": &schema.Schema{
										Description: `Icap Capability`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"led_flash_enabled": &schema.Schema{
										Description: `Led Flash Enabled`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"led_flash_seconds": &schema.Schema{
										Description: `Led Flash Seconds`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"operational_mode": &schema.Schema{
										Description: `Operational Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"policy_tag_name": &schema.Schema{
										Description: `Policy Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_calendar_profile": &schema.Schema{
										Description: `Power Calendar Profile`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_mode": &schema.Schema{
										Description: `Power Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_profile": &schema.Schema{
										Description: `Power Profile`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_save_mode": &schema.Schema{
										Description: `Power Save Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_save_mode_capable": &schema.Schema{
										Description: `Power Save Mode Capable`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_type": &schema.Schema{
										Description: `Power Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radios": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"air_quality": &schema.Schema{
													Description: `Air Quality`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"band": &schema.Schema{
													Description: `Band`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"client_count": &schema.Schema{
													Description: `Client Count`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"interference": &schema.Schema{
													Description: `Interference`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"noise": &schema.Schema{
													Description: `Noise`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"traffic_util": &schema.Schema{
													Description: `Traffic Util`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"utilization": &schema.Schema{
													Description: `Utilization`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},

									"regulatory_domain": &schema.Schema{
										Description: `Regulatory Domain`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reset_reason": &schema.Schema{
										Description: `Reset Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rf_tag_name": &schema.Schema{
										Description: `Rf Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"site_tag_name": &schema.Schema{
										Description: `Site Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_mode": &schema.Schema{
										Description: `Sub Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"client_count": &schema.Schema{
							Description: `Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"collection_status": &schema.Schema{
							Description: `Collection Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"communication_state": &schema.Schema{
							Description: `Communication State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_group_hierarchy_id": &schema.Schema{
							Description: `Device Group Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_role": &schema.Schema{
							Description: `Device Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_series": &schema.Schema{
							Description: `Device Series`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fabric_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fabric_role": &schema.Schema{
										Description: `Fabric Role`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"fabric_site_name": &schema.Schema{
										Description: `Fabric Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"transit_fabrics": &schema.Schema{
										Description: `Transit Fabrics`,
										Type:        schema.TypeList,
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
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ha_last_reset_reason": &schema.Schema{
							Description: `Ha Last Reset Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ha_status": &schema.Schema{
							Description: `Ha Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv4_address": &schema.Schema{
							Description: `Ipv4 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6_address": &schema.Schema{
							Description: `Ipv6 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_boot_time": &schema.Schema{
							Description: `Last Boot Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"maintenance_mode_enabled": &schema.Schema{
							Description: `Maintenance Mode Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"metrics_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"air_quality_score": &schema.Schema{
										Description: `Air Quality Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_temperature": &schema.Schema{
										Description: `Avg Temperature`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"cpu_score": &schema.Schema{
										Description: `Cpu Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"cpu_utilization": &schema.Schema{
										Description: `Cpu Utilization`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"discard_interfaces": &schema.Schema{
										Description: `Discard Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"discard_score": &schema.Schema{
										Description: `Discard Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"error_interfaces": &schema.Schema{
										Description: `Error Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"error_score": &schema.Schema{
										Description: `Error Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"free_memory_buffer": &schema.Schema{
										Description: `Free Memory Buffer`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"free_memory_buffer_score": &schema.Schema{
										Description: `Free Memory Buffer Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"free_timer": &schema.Schema{
										Description: `Free Timer`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"free_timer_score": &schema.Schema{
										Description: `Free Timer Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"high_link_utilization_interfaces": &schema.Schema{
										Description: `High Link Utilization Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inter_device_connected_down_interfaces": &schema.Schema{
										Description: `Inter Device Connected Down Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inter_device_link_score": &schema.Schema{
										Description: `Inter Device Link Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"interference_score": &schema.Schema{
										Description: `Interference Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"link_utilization_score": &schema.Schema{
										Description: `Link Utilization Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_temperature": &schema.Schema{
										Description: `Max Temperature`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"memory_score": &schema.Schema{
										Description: `Memory Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"memory_utilization": &schema.Schema{
										Description: `Memory Utilization`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"noise_score": &schema.Schema{
										Description: `Noise Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_fabric_score": &schema.Schema{
										Description: `Overall Fabric Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_health_score": &schema.Schema{
										Description: `Overall Health Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"packet_pool": &schema.Schema{
										Description: `Packet Pool`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"packet_pool_score": &schema.Schema{
										Description: `Packet Pool Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"utilization_score": &schema.Schema{
										Description: `Utilization Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wqe_pool": &schema.Schema{
										Description: `Wqe Pool`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wqe_pool_score": &schema.Schema{
										Description: `Wqe Pool Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"os_type": &schema.Schema{
							Description: `Os Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_count": &schema.Schema{
							Description: `Port Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"product_vendor": &schema.Schema{
							Description: `Product Vendor`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_mode": &schema.Schema{
							Description: `Redundancy Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_peer_state": &schema.Schema{
							Description: `Redundancy Peer State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_peer_state_derived": &schema.Schema{
							Description: `Redundancy Peer State Derived`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_state": &schema.Schema{
							Description: `Redundancy State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_state_derived": &schema.Schema{
							Description: `Redundancy State Derived`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ring_status": &schema.Schema{
							Description: `Ring Status`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"software_version": &schema.Schema{
							Description: `Software Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"stack_type": &schema.Schema{
							Description: `Stack Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tag_names": &schema.Schema{
							Description: `Tag Names`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"up_time": &schema.Schema{
							Description: `Up Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wired_client_count": &schema.Schema{
							Description: `Wired Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wireless_client_count": &schema.Schema{
							Description: `Wireless Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},

						"ap_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"admin_state": &schema.Schema{
										Description: `Admin State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_group": &schema.Schema{
										Description: `Ap Group`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_operational_state": &schema.Schema{
										Description: `Ap Operational State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_type": &schema.Schema{
										Description: `Ap Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_time": &schema.Schema{
										Description: `Connected Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"connected_wlc_name": &schema.Schema{
										Description: `Connected Wlc Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ethernet_mac": &schema.Schema{
										Description: `Ethernet Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"flex_group": &schema.Schema{
										Description: `Flex Group`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"home_ap_enabled": &schema.Schema{
										Description: `Home Ap Enabled`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"icap_capability": &schema.Schema{
										Description: `Icap Capability`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"led_flash_enabled": &schema.Schema{
										Description: `Led Flash Enabled`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"led_flash_seconds": &schema.Schema{
										Description: `Led Flash Seconds`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"operational_mode": &schema.Schema{
										Description: `Operational Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"policy_tag_name": &schema.Schema{
										Description: `Policy Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_calendar_profile": &schema.Schema{
										Description: `Power Calendar Profile`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_mode": &schema.Schema{
										Description: `Power Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_profile": &schema.Schema{
										Description: `Power Profile`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_save_mode": &schema.Schema{
										Description: `Power Save Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_save_mode_capable": &schema.Schema{
										Description: `Power Save Mode Capable`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_type": &schema.Schema{
										Description: `Power Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radios": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"air_quality": &schema.Schema{
													Description: `Air Quality`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"band": &schema.Schema{
													Description: `Band`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"client_count": &schema.Schema{
													Description: `Client Count`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"interference": &schema.Schema{
													Description: `Interference`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"noise": &schema.Schema{
													Description: `Noise`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"traffic_util": &schema.Schema{
													Description: `Traffic Util`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"utilization": &schema.Schema{
													Description: `Utilization`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},

									"regulatory_domain": &schema.Schema{
										Description: `Regulatory Domain`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reset_reason": &schema.Schema{
										Description: `Reset Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rf_tag_name": &schema.Schema{
										Description: `Rf Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"site_tag_name": &schema.Schema{
										Description: `Site Tag Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_mode": &schema.Schema{
										Description: `Sub Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"client_count": &schema.Schema{
							Description: `Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"collection_status": &schema.Schema{
							Description: `Collection Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"communication_state": &schema.Schema{
							Description: `Communication State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_group_hierarchy_id": &schema.Schema{
							Description: `Device Group Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_role": &schema.Schema{
							Description: `Device Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_series": &schema.Schema{
							Description: `Device Series`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fabric_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fabric_role": &schema.Schema{
										Description: `Fabric Role`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"fabric_site_name": &schema.Schema{
										Description: `Fabric Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"transit_fabrics": &schema.Schema{
										Description: `Transit Fabrics`,
										Type:        schema.TypeList,
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
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ha_last_reset_reason": &schema.Schema{
							Description: `Ha Last Reset Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ha_status": &schema.Schema{
							Description: `Ha Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv4_address": &schema.Schema{
							Description: `Ipv4 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6_address": &schema.Schema{
							Description: `Ipv6 Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_boot_time": &schema.Schema{
							Description: `Last Boot Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"maintenance_mode_enabled": &schema.Schema{
							Description: `Maintenance Mode Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"metrics_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"air_quality_score": &schema.Schema{
										Description: `Air Quality Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_temperature": &schema.Schema{
										Description: `Avg Temperature`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"cpu_score": &schema.Schema{
										Description: `Cpu Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"cpu_utilization": &schema.Schema{
										Description: `Cpu Utilization`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"discard_interfaces": &schema.Schema{
										Description: `Discard Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"discard_score": &schema.Schema{
										Description: `Discard Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"error_interfaces": &schema.Schema{
										Description: `Error Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"error_score": &schema.Schema{
										Description: `Error Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"free_memory_buffer": &schema.Schema{
										Description: `Free Memory Buffer`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"free_memory_buffer_score": &schema.Schema{
										Description: `Free Memory Buffer Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"free_timer": &schema.Schema{
										Description: `Free Timer`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"free_timer_score": &schema.Schema{
										Description: `Free Timer Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"high_link_utilization_interfaces": &schema.Schema{
										Description: `High Link Utilization Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inter_device_connected_down_interfaces": &schema.Schema{
										Description: `Inter Device Connected Down Interfaces`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"inter_device_link_score": &schema.Schema{
										Description: `Inter Device Link Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"interference_score": &schema.Schema{
										Description: `Interference Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"link_utilization_score": &schema.Schema{
										Description: `Link Utilization Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_temperature": &schema.Schema{
										Description: `Max Temperature`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"memory_score": &schema.Schema{
										Description: `Memory Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"memory_utilization": &schema.Schema{
										Description: `Memory Utilization`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"noise_score": &schema.Schema{
										Description: `Noise Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_fabric_score": &schema.Schema{
										Description: `Overall Fabric Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_health_score": &schema.Schema{
										Description: `Overall Health Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"packet_pool": &schema.Schema{
										Description: `Packet Pool`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"packet_pool_score": &schema.Schema{
										Description: `Packet Pool Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"utilization_score": &schema.Schema{
										Description: `Utilization Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wqe_pool": &schema.Schema{
										Description: `Wqe Pool`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wqe_pool_score": &schema.Schema{
										Description: `Wqe Pool Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"os_type": &schema.Schema{
							Description: `Os Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_count": &schema.Schema{
							Description: `Port Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"product_vendor": &schema.Schema{
							Description: `Product Vendor`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_mode": &schema.Schema{
							Description: `Redundancy Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_peer_state": &schema.Schema{
							Description: `Redundancy Peer State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_peer_state_derived": &schema.Schema{
							Description: `Redundancy Peer State Derived`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_state": &schema.Schema{
							Description: `Redundancy State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_state_derived": &schema.Schema{
							Description: `Redundancy State Derived`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ring_status": &schema.Schema{
							Description: `Ring Status`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"software_version": &schema.Schema{
							Description: `Software Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"stack_type": &schema.Schema{
							Description: `Stack Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tag_names": &schema.Schema{
							Description: `Tag Names`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"up_time": &schema.Schema{
							Description: `Up Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wired_client_count": &schema.Schema{
							Description: `Wired Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wireless_client_count": &schema.Schema{
							Description: `Wireless Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vID, okID := d.GetOk("id")
	vManagementIPAddress, okManagementIPAddress := d.GetOk("management_ip_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vFamily, okFamily := d.GetOk("family")
	vType, okType := d.GetOk("type")
	vRole, okRole := d.GetOk("role")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vMaintenanceMode, okMaintenanceMode := d.GetOk("maintenance_mode")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vHealthScore, okHealthScore := d.GetOk("health_score")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")

	method1 := []bool{okStartTime, okEndTime, okLimit, okOffset, okSortBy, okOrder, okSiteHierarchy, okSiteHierarchyID, okSiteID, okID, okManagementIPAddress, okMacAddress, okFamily, okType, okRole, okSerialNumber, okMaintenanceMode, okSoftwareVersion, okHealthScore, okView, okAttribute}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okStartTime, okEndTime, okView, okAttribute}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters")
		queryParams1 := dnacentersdkgo.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSiteHierarchy {
			queryParams1.SiteHierarchy = vSiteHierarchy.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okManagementIPAddress {
			queryParams1.ManagementIPAddress = vManagementIPAddress.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okMaintenanceMode {
			queryParams1.MaintenanceMode = vMaintenanceMode.(bool)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = vSoftwareVersion.(string)
		}
		if okHealthScore {
			queryParams1.HealthScore = vHealthScore.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}

		response1, restyResp1, err := client.Devices.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters", err,
				"Failure at GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTheDeviceDataForTheGivenDeviceIDUUID")
		vvID := vID.(string)
		queryParams2 := dnacentersdkgo.GetTheDeviceDataForTheGivenDeviceIDUUIDQueryParams{}

		if okStartTime {
			queryParams2.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams2.EndTime = vEndTime.(float64)
		}
		if okView {
			queryParams2.View = vView.(string)
		}
		if okAttribute {
			queryParams2.Attribute = vAttribute.(string)
		}

		response2, restyResp2, err := client.Devices.GetTheDeviceDataForTheGivenDeviceIDUUID(vvID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheDeviceDataForTheGivenDeviceIDUUID", err,
				"Failure at GetTheDeviceDataForTheGivenDeviceIDUUID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheDeviceDataForTheGivenDeviceIDUUID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItems(items *[]dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponse) []map[string]interface{} {
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
		respItem["ap_details"] = flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsApDetails(item.ApDetails)
		respItem["metrics_details"] = flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsMetricsDetails(item.MetricsDetails)
		respItem["fabric_details"] = flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsFabricDetails(item.FabricDetails)
		respItem["aggregate_attributes"] = flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsApDetails(item *dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetails) []map[string]interface{} {
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
	respItem["radios"] = flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsApDetailsRadios(item.Radios)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsApDetailsRadios(items *[]dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetailsRadios) []map[string]interface{} {
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

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsMetricsDetails(item *dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseMetricsDetails) []map[string]interface{} {
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

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsFabricDetails(item *dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricDetails) []map[string]interface{} {
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

func flattenDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseAggregateAttributes) []map[string]interface{} {
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

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItem(item *dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["ap_details"] = flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemApDetails(item.ApDetails)
	respItem["metrics_details"] = flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemMetricsDetails(item.MetricsDetails)
	respItem["fabric_details"] = flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemFabricDetails(item.FabricDetails)
	respItem["aggregate_attributes"] = flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemAggregateAttributes(item.AggregateAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemApDetails(item *dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetails) []map[string]interface{} {
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
	respItem["radios"] = flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemApDetailsRadios(item.Radios)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemApDetailsRadios(items *[]dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetailsRadios) []map[string]interface{} {
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

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemMetricsDetails(item *dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseMetricsDetails) []map[string]interface{} {
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

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemFabricDetails(item *dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricDetails) []map[string]interface{} {
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

func flattenDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDItemAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseAggregateAttributes) []map[string]interface{} {
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
