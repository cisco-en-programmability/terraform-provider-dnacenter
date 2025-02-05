package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClients() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Retrieves the list of clients, while also offering basic filtering and sorting capabilities. For detailed information
about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml

- Retrieves specific client information matching the MAC address. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceClientsRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Refer to ClientAttribute schema for list of attributes supported Examples: *attribute=band* (single attribute requested) *attribute=band&attribute=ssid&attribute=overallScore* (multiple attribute requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"band": &schema.Schema{
				Description: `band query parameter. WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz GHz Examples:
*band=5GHZ* (single band requested)
*band=2.4GHZ&band=6GHZ* (multiple band requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"connected_network_device_name": &schema.Schema{
				Description: `connectedNetworkDeviceName query parameter. Name of the neighbor network device that client is connected to. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search. Ex: **ap-25** or *ap-25** or **ap-25*
Examples:
*connectedNetworkDeviceName=ap-25* (single connectedNetworkDeviceName requested)
*connectedNetworkDeviceName=ap-25&ap-34* (multiple connectedNetworkDeviceName requested)
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
			"id": &schema.Schema{
				Description: `id path parameter. id is the client mac address. It can be specified is any notational conventions  01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_address": &schema.Schema{
				Description: `ipv4Address query parameter. IPv4 Address of the network entity either network device or client This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples:
*ipv4Address=1.1.1.1* (single ipv4Address requested)
*ipv4Address=1.1.1.1&ipv4Address=2.2.2.2* (multiple ipv4Address requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": &schema.Schema{
				Description: `ipv6Address query parameter. IPv6 Address of the network entity either network device or client This field supports wildcard (***) character-based search. Ex: **2001:db8** or *2001:db8** or **2001:db8*
Examples:
*ipv6Address=2001:db8:0:0:0:0:2:1* (single ipv6Address requested)
*ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348* (multiple ipv6Address requested)
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
			"os_type": &schema.Schema{
				Description: `osType query parameter. Client device operating system type. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **iOS** or *iOS** or **iOS* Examples:
*osType=iOS* (single osType requested)
*osType=iOS&osType=Android* (multiple osType requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_version": &schema.Schema{
				Description: `osVersion query parameter. Client device operating system version This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **14.3** or *14.3** or **14.3* Examples:
*osVersion=14.3* (single osVersion requested)
*osVersion=14.3&osVersion=10.1* (multiple osVersion requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **BuildingName** or *BuildingName** or **BuildingName* Examples: *siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested) *siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2* (multiple siteHierarchy requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (***) character-based search.  Ex: **buildingUuid** or *buildingUuid** or **buildingUuid* Examples: *siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid* (single siteHierarchyId requested) *siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2* (multiple siteHierarchyId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: *siteId=floorUuid* (single siteId requested) *siteId=floorUuid1&siteId=floorUuid2* (multiple siteId requested)
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
			"ssid": &schema.Schema{
				Description: `ssid query parameter. SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID Wireless Local Area Network Identifier. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search.  Ex: **Alpha** or *Alpha** or **Alpha*
Examples:
*ssid=Alpha* (single ssid requested)
*ssid=Alpha&ssid=Guest* (multiple ssid requested)
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
				Description: `type query parameter. The client device type whether client is connected to network through Wired or Wireless medium.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. Client related Views Refer to ClientView schema for list of views supported Examples:
*view=Wireless* (single view requested)
*view=WirelessHealth&view=WirelessTraffic* (multiple view requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlc_name": &schema.Schema{
				Description: `wlcName query parameter. Wireless Controller name that reports the wireless client. This field supports wildcard (***) character-based search. If the value contains the (***) character, please use the /query API for regex search. Ex: **wlc-25** or *wlc-25** or **wlc-25*
Examples:
*wlcName=wlc-25* (single wlcName requested)
*wlcName=wlc-25&wlc-34* (multiple wlcName requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_network_device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connected_network_device_id": &schema.Schema{
										Description: `Connected Network Device Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_mac": &schema.Schema{
										Description: `Connected Network Device Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_management_ip": &schema.Schema{
										Description: `Connected Network Device Management Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_name": &schema.Schema{
										Description: `Connected Network Device Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_type": &schema.Schema{
										Description: `Connected Network Device Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"duplex_mode": &schema.Schema{
										Description: `Duplex Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"interface_name": &schema.Schema{
										Description: `Interface Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"interface_speed": &schema.Schema{
										Description: `Interface Speed`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"connection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_ethernet_mac": &schema.Schema{
										Description: `Ap Ethernet Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mode": &schema.Schema{
										Description: `Ap Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"band": &schema.Schema{
										Description: `Band`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bridge_vmmode": &schema.Schema{
										Description: `Bridge V M Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel": &schema.Schema{
										Description: `Channel`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel_width": &schema.Schema{
										Description: `Channel Width`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"data_rate": &schema.Schema{
										Description: `Data Rate`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"is_ios_analytics_capable": &schema.Schema{
										Description: `Is Ios Analytics Capable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"l2_vn": &schema.Schema{
										Description: `L2 Vn`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"l3_vn": &schema.Schema{
										Description: `L3 Vn`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"link_speed": &schema.Schema{
										Description: `Link Speed`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol_capability": &schema.Schema{
										Description: `Protocol Capability`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio_id": &schema.Schema{
										Description: `Radio Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"security_group_tag": &schema.Schema{
										Description: `Security Group Tag`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"session_duration": &schema.Schema{
										Description: `Session Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_duid": &schema.Schema{
										Description: `Upn Duid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_id": &schema.Schema{
										Description: `Upn Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_name": &schema.Schema{
										Description: `Upn Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_owner": &schema.Schema{
										Description: `Upn Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vn_id": &schema.Schema{
										Description: `Vn Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wlc_id": &schema.Schema{
										Description: `Wlc Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wlc_name": &schema.Schema{
										Description: `Wlc Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"connection_status": &schema.Schema{
							Description: `Connection Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"form_factor": &schema.Schema{
							Description: `Form Factor`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connected_score": &schema.Schema{
										Description: `Connected Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"is_link_error_included": &schema.Schema{
										Description: `Is Link Error Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_rssi_included": &schema.Schema{
										Description: `Is Rssi Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_snr_included": &schema.Schema{
										Description: `Is Snr Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_error_percentage_threshold": &schema.Schema{
										Description: `Link Error Percentage Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"onboarding_score": &schema.Schema{
										Description: `Onboarding Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_score": &schema.Schema{
										Description: `Overall Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rssi_threshold": &schema.Schema{
										Description: `Rssi Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"snr_threshold": &schema.Schema{
										Description: `Snr Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
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

						"ipv6_addresses": &schema.Schema{
							Description: `Ipv6 Addresses`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"is_private_mac_address": &schema.Schema{
							Description: `Is Private Mac Address`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"latency": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"background": &schema.Schema{
										Description: `Background`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"best_effort": &schema.Schema{
										Description: `Best Effort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"video": &schema.Schema{
										Description: `Video`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"voice": &schema.Schema{
										Description: `Voice`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onboarding": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_failure_reason": &schema.Schema{
										Description: `Aaa Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"aaa_server_ip": &schema.Schema{
										Description: `Aaa Server Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"assoc_done_time": &schema.Schema{
										Description: `Assoc Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"assoc_failure_reason": &schema.Schema{
										Description: `Assoc Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_done_time": &schema.Schema{
										Description: `Auth Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_assoc_duration": &schema.Schema{
										Description: `Avg Assoc Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_auth_duration": &schema.Schema{
										Description: `Avg Auth Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_dhcp_duration": &schema.Schema{
										Description: `Avg Dhcp Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_run_duration": &schema.Schema{
										Description: `Avg Run Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dhcp_done_time": &schema.Schema{
										Description: `Dhcp Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dhcp_failure_reason": &schema.Schema{
										Description: `Dhcp Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dhcp_server_ip": &schema.Schema{
										Description: `Dhcp Server Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"failed_roaming_count": &schema.Schema{
										Description: `Failed Roaming Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"latest_failure_reason": &schema.Schema{
										Description: `Latest Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"max_assoc_duration": &schema.Schema{
										Description: `Max Assoc Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_auth_duration": &schema.Schema{
										Description: `Max Auth Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_dhcp_duration": &schema.Schema{
										Description: `Max Dhcp Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_roaming_duration": &schema.Schema{
										Description: `Max Roaming Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_run_duration": &schema.Schema{
										Description: `Max Run Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"onboarding_time": &schema.Schema{
										Description: `Onboarding Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"other_failure_reason": &schema.Schema{
										Description: `Other Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"roaming_time": &schema.Schema{
										Description: `Roaming Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"successful_roaming_count": &schema.Schema{
										Description: `Successful Roaming Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"total_roaming_attempts": &schema.Schema{
										Description: `Total Roaming Attempts`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"os_type": &schema.Schema{
							Description: `Os Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"os_version": &schema.Schema{
							Description: `Os Version`,
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

						"tracked": &schema.Schema{
							Description: `Tracked`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"traffic": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_request_count": &schema.Schema{
										Description: `Dns Request Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dns_response_count": &schema.Schema{
										Description: `Dns Response Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_bytes": &schema.Schema{
										Description: `Rx Bytes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_link_error_percentage": &schema.Schema{
										Description: `Rx Link Error Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"rx_packets": &schema.Schema{
										Description: `Rx Packets`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_rate": &schema.Schema{
										Description: `Rx Rate`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"rx_retries": &schema.Schema{
										Description: `Rx Retries`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_retry_percentage": &schema.Schema{
										Description: `Rx Retry Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_drop_percentage": &schema.Schema{
										Description: `Tx Drop Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_drops": &schema.Schema{
										Description: `Tx Drops`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_link_error_percentage": &schema.Schema{
										Description: `Tx Link Error Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"tx_packets": &schema.Schema{
										Description: `Tx Packets`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_rate": &schema.Schema{
										Description: `Tx Rate`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"usage": &schema.Schema{
										Description: `Usage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"user_id": &schema.Schema{
							Description: `User Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"username": &schema.Schema{
							Description: `Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vendor": &schema.Schema{
							Description: `Vendor`,
							Type:        schema.TypeString,
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

						"connected_network_device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connected_network_device_id": &schema.Schema{
										Description: `Connected Network Device Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_mac": &schema.Schema{
										Description: `Connected Network Device Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_management_ip": &schema.Schema{
										Description: `Connected Network Device Management Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_name": &schema.Schema{
										Description: `Connected Network Device Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_network_device_type": &schema.Schema{
										Description: `Connected Network Device Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"duplex_mode": &schema.Schema{
										Description: `Duplex Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"interface_name": &schema.Schema{
										Description: `Interface Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"interface_speed": &schema.Schema{
										Description: `Interface Speed`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"connection": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_ethernet_mac": &schema.Schema{
										Description: `Ap Ethernet Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mode": &schema.Schema{
										Description: `Ap Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"band": &schema.Schema{
										Description: `Band`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bridge_vmmode": &schema.Schema{
										Description: `Bridge V M Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel": &schema.Schema{
										Description: `Channel`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel_width": &schema.Schema{
										Description: `Channel Width`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"data_rate": &schema.Schema{
										Description: `Data Rate`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"is_ios_analytics_capable": &schema.Schema{
										Description: `Is Ios Analytics Capable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"l2_vn": &schema.Schema{
										Description: `L2 Vn`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"l3_vn": &schema.Schema{
										Description: `L3 Vn`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"link_speed": &schema.Schema{
										Description: `Link Speed`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol_capability": &schema.Schema{
										Description: `Protocol Capability`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio_id": &schema.Schema{
										Description: `Radio Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"security_group_tag": &schema.Schema{
										Description: `Security Group Tag`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"session_duration": &schema.Schema{
										Description: `Session Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_duid": &schema.Schema{
										Description: `Upn Duid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_id": &schema.Schema{
										Description: `Upn Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_name": &schema.Schema{
										Description: `Upn Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"upn_owner": &schema.Schema{
										Description: `Upn Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vn_id": &schema.Schema{
										Description: `Vn Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wlc_id": &schema.Schema{
										Description: `Wlc Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wlc_name": &schema.Schema{
										Description: `Wlc Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"connection_status": &schema.Schema{
							Description: `Connection Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"form_factor": &schema.Schema{
							Description: `Form Factor`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connected_score": &schema.Schema{
										Description: `Connected Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"is_link_error_included": &schema.Schema{
										Description: `Is Link Error Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_rssi_included": &schema.Schema{
										Description: `Is Rssi Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_snr_included": &schema.Schema{
										Description: `Is Snr Included`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_error_percentage_threshold": &schema.Schema{
										Description: `Link Error Percentage Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"onboarding_score": &schema.Schema{
										Description: `Onboarding Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overall_score": &schema.Schema{
										Description: `Overall Score`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rssi_threshold": &schema.Schema{
										Description: `Rssi Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"snr_threshold": &schema.Schema{
										Description: `Snr Threshold`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
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

						"ipv6_addresses": &schema.Schema{
							Description: `Ipv6 Addresses`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"is_private_mac_address": &schema.Schema{
							Description: `Is Private Mac Address`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"latency": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"background": &schema.Schema{
										Description: `Background`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"best_effort": &schema.Schema{
										Description: `Best Effort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"video": &schema.Schema{
										Description: `Video`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"voice": &schema.Schema{
										Description: `Voice`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onboarding": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_failure_reason": &schema.Schema{
										Description: `Aaa Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"aaa_server_ip": &schema.Schema{
										Description: `Aaa Server Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"assoc_done_time": &schema.Schema{
										Description: `Assoc Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"assoc_failure_reason": &schema.Schema{
										Description: `Assoc Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_done_time": &schema.Schema{
										Description: `Auth Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_assoc_duration": &schema.Schema{
										Description: `Avg Assoc Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_auth_duration": &schema.Schema{
										Description: `Avg Auth Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_dhcp_duration": &schema.Schema{
										Description: `Avg Dhcp Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"avg_run_duration": &schema.Schema{
										Description: `Avg Run Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dhcp_done_time": &schema.Schema{
										Description: `Dhcp Done Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dhcp_failure_reason": &schema.Schema{
										Description: `Dhcp Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dhcp_server_ip": &schema.Schema{
										Description: `Dhcp Server Ip`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"failed_roaming_count": &schema.Schema{
										Description: `Failed Roaming Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"latest_failure_reason": &schema.Schema{
										Description: `Latest Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"max_assoc_duration": &schema.Schema{
										Description: `Max Assoc Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_auth_duration": &schema.Schema{
										Description: `Max Auth Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_dhcp_duration": &schema.Schema{
										Description: `Max Dhcp Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_roaming_duration": &schema.Schema{
										Description: `Max Roaming Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"max_run_duration": &schema.Schema{
										Description: `Max Run Duration`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"onboarding_time": &schema.Schema{
										Description: `Onboarding Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"other_failure_reason": &schema.Schema{
										Description: `Other Failure Reason`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"roaming_time": &schema.Schema{
										Description: `Roaming Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"successful_roaming_count": &schema.Schema{
										Description: `Successful Roaming Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"total_roaming_attempts": &schema.Schema{
										Description: `Total Roaming Attempts`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"os_type": &schema.Schema{
							Description: `Os Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"os_version": &schema.Schema{
							Description: `Os Version`,
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

						"tracked": &schema.Schema{
							Description: `Tracked`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"traffic": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_request_count": &schema.Schema{
										Description: `Dns Request Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"dns_response_count": &schema.Schema{
										Description: `Dns Response Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_bytes": &schema.Schema{
										Description: `Rx Bytes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_link_error_percentage": &schema.Schema{
										Description: `Rx Link Error Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"rx_packets": &schema.Schema{
										Description: `Rx Packets`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_rate": &schema.Schema{
										Description: `Rx Rate`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"rx_retries": &schema.Schema{
										Description: `Rx Retries`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"rx_retry_percentage": &schema.Schema{
										Description: `Rx Retry Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_drop_percentage": &schema.Schema{
										Description: `Tx Drop Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_drops": &schema.Schema{
										Description: `Tx Drops`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_link_error_percentage": &schema.Schema{
										Description: `Tx Link Error Percentage`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"tx_packets": &schema.Schema{
										Description: `Tx Packets`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"tx_rate": &schema.Schema{
										Description: `Tx Rate`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"usage": &schema.Schema{
										Description: `Usage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"user_id": &schema.Schema{
							Description: `User Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"username": &schema.Schema{
							Description: `Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vendor": &schema.Schema{
							Description: `Vendor`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClientsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vType, okType := d.GetOk("type")
	vOsType, okOsType := d.GetOk("os_type")
	vOsVersion, okOsVersion := d.GetOk("os_version")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vIPv4Address, okIPv4Address := d.GetOk("ipv4_address")
	vIPv6Address, okIPv6Address := d.GetOk("ipv6_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vWlcName, okWlcName := d.GetOk("wlc_name")
	vConnectedNetworkDeviceName, okConnectedNetworkDeviceName := d.GetOk("connected_network_device_name")
	vSSID, okSSID := d.GetOk("ssid")
	vBand, okBand := d.GetOk("band")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")
	vID, okID := d.GetOk("id")

	method1 := []bool{okStartTime, okEndTime, okLimit, okOffset, okSortBy, okOrder, okType, okOsType, okOsVersion, okSiteHierarchy, okSiteHierarchyID, okSiteID, okIPv4Address, okIPv6Address, okMacAddress, okWlcName, okConnectedNetworkDeviceName, okSSID, okBand, okView, okAttribute, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okStartTime, okEndTime, okView, okAttribute, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities")

		headerParams1 := dnacentersdkgo.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams{}

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
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOsType {
			queryParams1.OsType = vOsType.(string)
		}
		if okOsVersion {
			queryParams1.OsVersion = vOsVersion.(string)
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
		if okIPv4Address {
			queryParams1.IPv4Address = vIPv4Address.(string)
		}
		if okIPv6Address {
			queryParams1.IPv6Address = vIPv6Address.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okWlcName {
			queryParams1.WlcName = vWlcName.(string)
		}
		if okConnectedNetworkDeviceName {
			queryParams1.ConnectedNetworkDeviceName = vConnectedNetworkDeviceName.(string)
		}
		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okBand {
			queryParams1.Band = vBand.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Clients.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities", err,
				"Failure at RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrievesSpecificClientInformationMatchingTheMacaddress")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.RetrievesSpecificClientInformationMatchingTheMacaddressHeaderParams{}
		queryParams2 := dnacentersdkgo.RetrievesSpecificClientInformationMatchingTheMacaddressQueryParams{}

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
		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Clients.RetrievesSpecificClientInformationMatchingTheMacaddress(vvID, &headerParams2, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesSpecificClientInformationMatchingTheMacaddress", err,
				"Failure at RetrievesSpecificClientInformationMatchingTheMacaddress, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesSpecificClientInformationMatchingTheMacaddress response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItems(items *[]dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["mac_address"] = item.MacAddress
		respItem["type"] = item.Type
		respItem["name"] = item.Name
		respItem["user_id"] = item.UserID
		respItem["username"] = item.Username
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_addresses"] = item.IPv6Addresses
		respItem["vendor"] = item.Vendor
		respItem["os_type"] = item.OsType
		respItem["os_version"] = item.OsVersion
		respItem["form_factor"] = item.FormFactor
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["site_id"] = item.SiteID
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItem["connection_status"] = item.ConnectionStatus
		respItem["tracked"] = item.Tracked
		respItem["is_private_mac_address"] = boolPtrToString(item.IsPrivateMacAddress)
		respItem["health"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsHealth(item.Health)
		respItem["traffic"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsTraffic(item.Traffic)
		respItem["connected_network_device"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsConnectedNetworkDevice(item.ConnectedNetworkDevice)
		respItem["connection"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsConnection(item.Connection)
		respItem["onboarding"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsOnboarding(item.Onboarding)
		respItem["latency"] = flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsLatency(item.Latency)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsHealth(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["overall_score"] = item.OverallScore
	respItem["onboarding_score"] = item.OnboardingScore
	respItem["connected_score"] = item.ConnectedScore
	respItem["link_error_percentage_threshold"] = item.LinkErrorPercentageThreshold
	respItem["is_link_error_included"] = boolPtrToString(item.IsLinkErrorIncluded)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["snr_threshold"] = item.SnrThreshold
	respItem["is_rssi_included"] = boolPtrToString(item.IsRssiIncluded)
	respItem["is_snr_included"] = boolPtrToString(item.IsSnrIncluded)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsTraffic(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseTraffic) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tx_bytes"] = item.TxBytes
	respItem["rx_bytes"] = item.RxBytes
	respItem["usage"] = item.Usage
	respItem["rx_packets"] = item.RxPackets
	respItem["tx_packets"] = item.TxPackets
	respItem["rx_rate"] = item.RxRate
	respItem["tx_rate"] = item.TxRate
	respItem["rx_link_error_percentage"] = item.RxLinkErrorPercentage
	respItem["tx_link_error_percentage"] = item.TxLinkErrorPercentage
	respItem["rx_retries"] = item.RxRetries
	respItem["rx_retry_percentage"] = item.RxRetryPercentage
	respItem["tx_drops"] = item.TxDrops
	respItem["tx_drop_percentage"] = item.TxDropPercentage
	respItem["dns_request_count"] = item.DNSRequestCount
	respItem["dns_response_count"] = item.DNSResponseCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsConnectedNetworkDevice(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnectedNetworkDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connected_network_device_id"] = item.ConnectedNetworkDeviceID
	respItem["connected_network_device_name"] = item.ConnectedNetworkDeviceName
	respItem["connected_network_device_management_ip"] = item.ConnectedNetworkDeviceManagementIP
	respItem["connected_network_device_mac"] = item.ConnectedNetworkDeviceMac
	respItem["connected_network_device_type"] = item.ConnectedNetworkDeviceType
	respItem["interface_name"] = item.InterfaceName
	respItem["interface_speed"] = item.InterfaceSpeed
	respItem["duplex_mode"] = item.DuplexMode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsConnection(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["vlan_id"] = item.VLANID
	respItem["session_duration"] = item.SessionDuration
	respItem["vn_id"] = item.VnID
	respItem["l2_vn"] = item.L2Vn
	respItem["l3_vn"] = item.L3Vn
	respItem["security_group_tag"] = item.SecurityGroupTag
	respItem["link_speed"] = item.LinkSpeed
	respItem["bridge_vmmode"] = item.BridgeVMMode
	respItem["band"] = item.Band
	respItem["ssid"] = item.SSID
	respItem["auth_type"] = item.AuthType
	respItem["wlc_name"] = item.WlcName
	respItem["wlc_id"] = item.WlcID
	respItem["ap_mac"] = item.ApMac
	respItem["ap_ethernet_mac"] = item.ApEthernetMac
	respItem["ap_mode"] = item.ApMode
	respItem["radio_id"] = item.RadioID
	respItem["channel"] = item.Channel
	respItem["channel_width"] = item.ChannelWidth
	respItem["protocol"] = item.Protocol
	respItem["protocol_capability"] = item.ProtocolCapability
	respItem["upn_id"] = item.UpnID
	respItem["upn_name"] = item.UpnName
	respItem["upn_owner"] = item.UpnOwner
	respItem["upn_duid"] = item.UpnDuid
	respItem["rssi"] = item.Rssi
	respItem["snr"] = item.Snr
	respItem["data_rate"] = item.DataRate
	respItem["is_ios_analytics_capable"] = boolPtrToString(item.IsIosAnalyticsCapable)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsOnboarding(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["avg_run_duration"] = item.AvgRunDuration
	respItem["max_run_duration"] = item.MaxRunDuration
	respItem["avg_assoc_duration"] = item.AvgAssocDuration
	respItem["max_assoc_duration"] = item.MaxAssocDuration
	respItem["avg_auth_duration"] = item.AvgAuthDuration
	respItem["max_auth_duration"] = item.MaxAuthDuration
	respItem["avg_dhcp_duration"] = item.AvgDhcpDuration
	respItem["max_dhcp_duration"] = item.MaxDhcpDuration
	respItem["max_roaming_duration"] = item.MaxRoamingDuration
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["dhcp_server_ip"] = item.DhcpServerIP
	respItem["onboarding_time"] = item.OnboardingTime
	respItem["auth_done_time"] = item.AuthDoneTime
	respItem["assoc_done_time"] = item.AssocDoneTime
	respItem["dhcp_done_time"] = item.DhcpDoneTime
	respItem["roaming_time"] = item.RoamingTime
	respItem["failed_roaming_count"] = item.FailedRoamingCount
	respItem["successful_roaming_count"] = item.SuccessfulRoamingCount
	respItem["total_roaming_attempts"] = item.TotalRoamingAttempts
	respItem["assoc_failure_reason"] = item.AssocFailureReason
	respItem["aaa_failure_reason"] = item.AAAFailureReason
	respItem["dhcp_failure_reason"] = item.DhcpFailureReason
	respItem["other_failure_reason"] = item.OtherFailureReason
	respItem["latest_failure_reason"] = item.LatestFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesItemsLatency(item *dnacentersdkgo.ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseLatency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["video"] = item.Video
	respItem["voice"] = item.Voice
	respItem["best_effort"] = item.BestEffort
	respItem["background"] = item.Background

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItem(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["mac_address"] = item.MacAddress
	respItem["type"] = item.Type
	respItem["name"] = item.Name
	respItem["user_id"] = item.UserID
	respItem["username"] = item.Username
	respItem["ipv4_address"] = item.IPv4Address
	respItem["ipv6_addresses"] = item.IPv6Addresses
	respItem["vendor"] = item.Vendor
	respItem["os_type"] = item.OsType
	respItem["os_version"] = item.OsVersion
	respItem["form_factor"] = item.FormFactor
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["site_hierarchy_id"] = item.SiteHierarchyID
	respItem["site_id"] = item.SiteID
	respItem["last_updated_time"] = item.LastUpdatedTime
	respItem["connection_status"] = item.ConnectionStatus
	respItem["tracked"] = item.Tracked
	respItem["is_private_mac_address"] = boolPtrToString(item.IsPrivateMacAddress)
	respItem["health"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemHealth(item.Health)
	respItem["traffic"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemTraffic(item.Traffic)
	respItem["connected_network_device"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemConnectedNetworkDevice(item.ConnectedNetworkDevice)
	respItem["connection"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemConnection(item.Connection)
	respItem["onboarding"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemOnboarding(item.Onboarding)
	respItem["latency"] = flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemLatency(item.Latency)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemHealth(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["overall_score"] = item.OverallScore
	respItem["onboarding_score"] = item.OnboardingScore
	respItem["connected_score"] = item.ConnectedScore
	respItem["link_error_percentage_threshold"] = item.LinkErrorPercentageThreshold
	respItem["is_link_error_included"] = boolPtrToString(item.IsLinkErrorIncluded)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["snr_threshold"] = item.SnrThreshold
	respItem["is_rssi_included"] = boolPtrToString(item.IsRssiIncluded)
	respItem["is_snr_included"] = boolPtrToString(item.IsSnrIncluded)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemTraffic(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseTraffic) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tx_bytes"] = item.TxBytes
	respItem["rx_bytes"] = item.RxBytes
	respItem["usage"] = item.Usage
	respItem["rx_packets"] = item.RxPackets
	respItem["tx_packets"] = item.TxPackets
	respItem["rx_rate"] = item.RxRate
	respItem["tx_rate"] = item.TxRate
	respItem["rx_link_error_percentage"] = item.RxLinkErrorPercentage
	respItem["tx_link_error_percentage"] = item.TxLinkErrorPercentage
	respItem["rx_retries"] = item.RxRetries
	respItem["rx_retry_percentage"] = item.RxRetryPercentage
	respItem["tx_drops"] = item.TxDrops
	respItem["tx_drop_percentage"] = item.TxDropPercentage
	respItem["dns_request_count"] = item.DNSRequestCount
	respItem["dns_response_count"] = item.DNSResponseCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemConnectedNetworkDevice(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnectedNetworkDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connected_network_device_id"] = item.ConnectedNetworkDeviceID
	respItem["connected_network_device_name"] = item.ConnectedNetworkDeviceName
	respItem["connected_network_device_management_ip"] = item.ConnectedNetworkDeviceManagementIP
	respItem["connected_network_device_mac"] = item.ConnectedNetworkDeviceMac
	respItem["connected_network_device_type"] = item.ConnectedNetworkDeviceType
	respItem["interface_name"] = item.InterfaceName
	respItem["interface_speed"] = item.InterfaceSpeed
	respItem["duplex_mode"] = item.DuplexMode

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemConnection(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["vlan_id"] = item.VLANID
	respItem["session_duration"] = item.SessionDuration
	respItem["vn_id"] = item.VnID
	respItem["l2_vn"] = item.L2Vn
	respItem["l3_vn"] = item.L3Vn
	respItem["security_group_tag"] = item.SecurityGroupTag
	respItem["link_speed"] = item.LinkSpeed
	respItem["bridge_vmmode"] = item.BridgeVMMode
	respItem["band"] = item.Band
	respItem["ssid"] = item.SSID
	respItem["auth_type"] = item.AuthType
	respItem["wlc_name"] = item.WlcName
	respItem["wlc_id"] = item.WlcID
	respItem["ap_mac"] = item.ApMac
	respItem["ap_ethernet_mac"] = item.ApEthernetMac
	respItem["ap_mode"] = item.ApMode
	respItem["radio_id"] = item.RadioID
	respItem["channel"] = item.Channel
	respItem["channel_width"] = item.ChannelWidth
	respItem["protocol"] = item.Protocol
	respItem["protocol_capability"] = item.ProtocolCapability
	respItem["upn_id"] = item.UpnID
	respItem["upn_name"] = item.UpnName
	respItem["upn_owner"] = item.UpnOwner
	respItem["upn_duid"] = item.UpnDuid
	respItem["rssi"] = item.Rssi
	respItem["snr"] = item.Snr
	respItem["data_rate"] = item.DataRate
	respItem["is_ios_analytics_capable"] = boolPtrToString(item.IsIosAnalyticsCapable)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemOnboarding(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["avg_run_duration"] = item.AvgRunDuration
	respItem["max_run_duration"] = item.MaxRunDuration
	respItem["avg_assoc_duration"] = item.AvgAssocDuration
	respItem["max_assoc_duration"] = item.MaxAssocDuration
	respItem["avg_auth_duration"] = item.AvgAuthDuration
	respItem["max_auth_duration"] = item.MaxAuthDuration
	respItem["avg_dhcp_duration"] = item.AvgDhcpDuration
	respItem["max_dhcp_duration"] = item.MaxDhcpDuration
	respItem["max_roaming_duration"] = item.MaxRoamingDuration
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["dhcp_server_ip"] = item.DhcpServerIP
	respItem["onboarding_time"] = item.OnboardingTime
	respItem["auth_done_time"] = item.AuthDoneTime
	respItem["assoc_done_time"] = item.AssocDoneTime
	respItem["dhcp_done_time"] = item.DhcpDoneTime
	respItem["roaming_time"] = item.RoamingTime
	respItem["failed_roaming_count"] = item.FailedRoamingCount
	respItem["successful_roaming_count"] = item.SuccessfulRoamingCount
	respItem["total_roaming_attempts"] = item.TotalRoamingAttempts
	respItem["assoc_failure_reason"] = item.AssocFailureReason
	respItem["aaa_failure_reason"] = item.AAAFailureReason
	respItem["dhcp_failure_reason"] = item.DhcpFailureReason
	respItem["other_failure_reason"] = item.OtherFailureReason
	respItem["latest_failure_reason"] = item.LatestFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsRetrievesSpecificClientInformationMatchingTheMacaddressItemLatency(item *dnacentersdkgo.ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseLatency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["video"] = item.Video
	respItem["voice"] = item.Voice
	respItem["best_effort"] = item.BestEffort
	respItem["background"] = item.Background

	return []map[string]interface{}{
		respItem,
	}

}
