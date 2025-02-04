package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceEvents() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API
Support Documentation' section to understand which fields are supported. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml

- API to fetch the details of an assurance event using event *id*. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceEventsRead,
		Schema: map[string]*schema.Schema{
			"ap_mac": &schema.Schema{
				Description: `apMac query parameter. MAC address of the access point. This parameter is applicable for *Unified AP* and *Wireless Client* events.
This field supports wildcard (***) character-based search. Ex: **50:0F** or *50:0F** or **50:0F*
Examples:
*apMac=50:0F:80:0F:F7:E0* (single apMac requested)
*apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0* (multiple apMac requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes (*id*, *name*, *timestamp*, *details*, *messageType*, *siteHierarchyId*, *siteHierarchy*, *deviceFamily*, *networkDeviceId*, *networkDeviceName*, *managementIpAddress*) would be part of the response.
 Examples:

*attribute=name* (single attribute requested)
*attribute=name&attribute=networkDeviceName* (multiple attribute requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_mac": &schema.Schema{
				Description: `clientMac query parameter. MAC address of the client. This parameter is applicable for *Wired Client* and *Wireless Client* events.
This field supports wildcard (***) character-based search. Ex: **66:2B** or *66:2B** or **66:2B*
Examples:
*clientMac=66:2B:B8:D2:01:56* (single clientMac requested)
*clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89* (multiple clientMac requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_family": &schema.Schema{
				Description: `deviceFamily query parameter. Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing *Routers* along with *Wireless Client* or *Unified AP* is not supported. Examples:
*deviceFamily=Switches and Hubs* (single deviceFamily requested)
*deviceFamily=Switches and Hubs&deviceFamily=Routers* (multiple deviceFamily requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *endTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Unique identifier for the event
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
			"message_type": &schema.Schema{
				Description: `messageType query parameter. Message type for the event.
Examples:
*messageType=Syslog* (single messageType requested)
*messageType=Trap&messageType=Syslog* (multiple messageType requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. The list of Network Device Uuids. (Ex. *6bef213c-19ca-4170-8375-b694e251101c*)
Examples:
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c* (single networkDeviceId requested)
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0* (multiple networkDeviceId with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_name": &schema.Schema{
				Description: `networkDeviceName query parameter. Network device name. This parameter is applicable for network device related families. This field supports wildcard (***) character-based search. Ex: **Branch** or *Branch** or **Branch* Examples:
*networkDeviceName=Branch-3-Gateway* (single networkDeviceName requested)
*networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch* (multiple networkDeviceName requested)
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
			"severity": &schema.Schema{
				Description: `severity query parameter. Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and *Wired Client* events.
| Value | Severity    | | ----| ----------| | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        |
Examples:
*severity=0* (single severity requested)
*severity=0&severity=1* (multiple severity requested)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. **uuid*, *uuid, uuid**
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyId requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
Examples:
*?siteId=id1* (single siteId requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple siteId requested)
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
If *startTime* is not provided, API will default to current time minus 24 hours.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The list of events views. Please refer to *EventViews* for the supported list
 Examples:

*view=network* (single view requested)
*view=network&view=ap* (multiple view requested)
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

						"affected_clients": &schema.Schema{
							Description: `Affected Clients`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ap_mac": &schema.Schema{
							Description: `Ap Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_radio_operation_state": &schema.Schema{
							Description: `Ap Radio Operation State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_role": &schema.Schema{
							Description: `Ap Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_switch_id": &schema.Schema{
							Description: `Ap Switch Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_switch_name": &schema.Schema{
							Description: `Ap Switch Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"assoc_rssi": &schema.Schema{
							Description: `Assoc Rssi`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"assoc_snr": &schema.Schema{
							Description: `Assoc Snr`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"audit_session_id": &schema.Schema{
							Description: `Audit Session Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"auth_server_ip": &schema.Schema{
							Description: `Auth Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"candidate_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"child_events": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"details": &schema.Schema{
										Description: `Details`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"failure_category": &schema.Schema{
										Description: `Failure Category`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reason_code": &schema.Schema{
										Description: `Reason Code`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reason_description": &schema.Schema{
										Description: `Reason Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"result_status": &schema.Schema{
										Description: `Result Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_reason_code": &schema.Schema{
										Description: `Sub Reason Code`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_reason_description": &schema.Schema{
										Description: `Sub Reason Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wireless_event_type": &schema.Schema{
										Description: `Wireless Event Type`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"client_mac": &schema.Schema{
							Description: `Client Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"connected_interface_name": &schema.Schema{
							Description: `Connected Interface Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"current_radio_power_level": &schema.Schema{
							Description: `Current Radio Power Level`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"details": &schema.Schema{
							Description: `Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dhcp_server_ip": &schema.Schema{
							Description: `Dhcp Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duid": &schema.Schema{
							Description: `Duid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"event_status": &schema.Schema{
							Description: `Event Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"facility": &schema.Schema{
							Description: `Facility`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_category": &schema.Schema{
							Description: `Failure Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_ip_address": &schema.Schema{
							Description: `Failure Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"frequency": &schema.Schema{
							Description: `Frequency`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"identifier": &schema.Schema{
							Description: `Identifier`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"invalid_ie_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"frame_type": &schema.Schema{
										Description: `Frame Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ies": &schema.Schema{
										Description: `Ies`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"ipv4": &schema.Schema{
							Description: `Ipv4`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_private_mac": &schema.Schema{
							Description: `Is Private Mac`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_ap_disconnect_reason": &schema.Schema{
							Description: `Last Ap Disconnect Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_ap_reset_type": &schema.Schema{
							Description: `Last Ap Reset Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"message_type": &schema.Schema{
							Description: `Message Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"missing_response_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"frame_type": &schema.Schema{
										Description: `Frame Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"mnemonic": &schema.Schema{
							Description: `Mnemonic`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_name": &schema.Schema{
							Description: `Network Device Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"new_radio_channel_list": &schema.Schema{
							Description: `New Radio Channel List`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"new_radio_channel_width": &schema.Schema{
							Description: `New Radio Channel Width`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"old_radio_channel_list": &schema.Schema{
							Description: `Old Radio Channel List`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"old_radio_channel_width": &schema.Schema{
							Description: `Old Radio Channel Width`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"previous_radio_power_level": &schema.Schema{
							Description: `Previous Radio Power Level`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"radio_channel_slot": &schema.Schema{
							Description: `Radio Channel Slot`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"radio_channel_utilization": &schema.Schema{
							Description: `Radio Channel Utilization`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_interference": &schema.Schema{
							Description: `Radio Interference`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_noise": &schema.Schema{
							Description: `Radio Noise`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reason_description": &schema.Schema{
							Description: `Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"replaced_device_serial_number": &schema.Schema{
							Description: `Replaced Device Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"replacing_device_serial_number": &schema.Schema{
							Description: `Replacing Device Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"result_status": &schema.Schema{
							Description: `Result Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"roam_type": &schema.Schema{
							Description: `Roam Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeInt,
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

						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sub_reason_description": &schema.Schema{
							Description: `Sub Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"switch_number": &schema.Schema{
							Description: `Switch Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"udn_id": &schema.Schema{
							Description: `Udn Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"udn_name": &schema.Schema{
							Description: `Udn Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"username": &schema.Schema{
							Description: `Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wireless_client_event_end_time": &schema.Schema{
							Description: `Wireless Client Event End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wireless_client_event_start_time": &schema.Schema{
							Description: `Wireless Client Event Start Time`,
							Type:        schema.TypeInt,
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"affected_clients": &schema.Schema{
							Description: `Affected Clients`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ap_mac": &schema.Schema{
							Description: `Ap Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_radio_operation_state": &schema.Schema{
							Description: `Ap Radio Operation State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_role": &schema.Schema{
							Description: `Ap Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_switch_id": &schema.Schema{
							Description: `Ap Switch Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_switch_name": &schema.Schema{
							Description: `Ap Switch Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"assoc_rssi": &schema.Schema{
							Description: `Assoc Rssi`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"assoc_snr": &schema.Schema{
							Description: `Assoc Snr`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"audit_session_id": &schema.Schema{
							Description: `Audit Session Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"auth_server_ip": &schema.Schema{
							Description: `Auth Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"candidate_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"child_events": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"details": &schema.Schema{
										Description: `Details`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"failure_category": &schema.Schema{
										Description: `Failure Category`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reason_code": &schema.Schema{
										Description: `Reason Code`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"reason_description": &schema.Schema{
										Description: `Reason Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"result_status": &schema.Schema{
										Description: `Result Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_reason_code": &schema.Schema{
										Description: `Sub Reason Code`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_reason_description": &schema.Schema{
										Description: `Sub Reason Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"wireless_event_type": &schema.Schema{
										Description: `Wireless Event Type`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"client_mac": &schema.Schema{
							Description: `Client Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"connected_interface_name": &schema.Schema{
							Description: `Connected Interface Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"current_radio_power_level": &schema.Schema{
							Description: `Current Radio Power Level`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"details": &schema.Schema{
							Description: `Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dhcp_server_ip": &schema.Schema{
							Description: `Dhcp Server Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duid": &schema.Schema{
							Description: `Duid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"event_status": &schema.Schema{
							Description: `Event Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"facility": &schema.Schema{
							Description: `Facility`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_category": &schema.Schema{
							Description: `Failure Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_ip_address": &schema.Schema{
							Description: `Failure Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"frequency": &schema.Schema{
							Description: `Frequency`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"identifier": &schema.Schema{
							Description: `Identifier`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"invalid_ie_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"frame_type": &schema.Schema{
										Description: `Frame Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ies": &schema.Schema{
										Description: `Ies`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"ipv4": &schema.Schema{
							Description: `Ipv4`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_private_mac": &schema.Schema{
							Description: `Is Private Mac`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_ap_disconnect_reason": &schema.Schema{
							Description: `Last Ap Disconnect Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_ap_reset_type": &schema.Schema{
							Description: `Last Ap Reset Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"message_type": &schema.Schema{
							Description: `Message Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"missing_response_a_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ap_name": &schema.Schema{
										Description: `Ap Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bssid": &schema.Schema{
										Description: `Bssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"frame_type": &schema.Schema{
										Description: `Frame Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"mnemonic": &schema.Schema{
							Description: `Mnemonic`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_device_name": &schema.Schema{
							Description: `Network Device Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"new_radio_channel_list": &schema.Schema{
							Description: `New Radio Channel List`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"new_radio_channel_width": &schema.Schema{
							Description: `New Radio Channel Width`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"old_radio_channel_list": &schema.Schema{
							Description: `Old Radio Channel List`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"old_radio_channel_width": &schema.Schema{
							Description: `Old Radio Channel Width`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"previous_radio_power_level": &schema.Schema{
							Description: `Previous Radio Power Level`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"radio_channel_slot": &schema.Schema{
							Description: `Radio Channel Slot`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"radio_channel_utilization": &schema.Schema{
							Description: `Radio Channel Utilization`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_interference": &schema.Schema{
							Description: `Radio Interference`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_noise": &schema.Schema{
							Description: `Radio Noise`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reason_description": &schema.Schema{
							Description: `Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"replaced_device_serial_number": &schema.Schema{
							Description: `Replaced Device Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"replacing_device_serial_number": &schema.Schema{
							Description: `Replacing Device Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"result_status": &schema.Schema{
							Description: `Result Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"roam_type": &schema.Schema{
							Description: `Roam Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeInt,
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

						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sub_reason_description": &schema.Schema{
							Description: `Sub Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"switch_number": &schema.Schema{
							Description: `Switch Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"udn_id": &schema.Schema{
							Description: `Udn Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"udn_name": &schema.Schema{
							Description: `Udn Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"username": &schema.Schema{
							Description: `Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wireless_client_event_end_time": &schema.Schema{
							Description: `Wireless Client Event End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wireless_client_event_start_time": &schema.Schema{
							Description: `Wireless Client Event Start Time`,
							Type:        schema.TypeInt,
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
		},
	}
}

func dataSourceAssuranceEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceFamily, okDeviceFamily := d.GetOk("device_family")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vMessageType, okMessageType := d.GetOk("message_type")
	vSeverity, okSeverity := d.GetOk("severity")
	vSiteID, okSiteID := d.GetOk("site_id")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vNetworkDeviceName, okNetworkDeviceName := d.GetOk("network_device_name")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vApMac, okApMac := d.GetOk("ap_mac")
	vClientMac, okClientMac := d.GetOk("client_mac")
	vAttribute, okAttribute := d.GetOk("attribute")
	vView, okView := d.GetOk("view")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")
	vID, okID := d.GetOk("id")

	method1 := []bool{okDeviceFamily, okStartTime, okEndTime, okMessageType, okSeverity, okSiteID, okSiteHierarchyID, okNetworkDeviceName, okNetworkDeviceID, okApMac, okClientMac, okAttribute, okView, okOffset, okLimit, okSortBy, okOrder, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okAttribute, okView, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: QueryAssuranceEvents")

		headerParams1 := dnacentersdkgo.QueryAssuranceEventsHeaderParams{}
		queryParams1 := dnacentersdkgo.QueryAssuranceEventsQueryParams{}

		if okDeviceFamily {
			queryParams1.DeviceFamily = vDeviceFamily.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okMessageType {
			queryParams1.MessageType = vMessageType.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(float64)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okNetworkDeviceName {
			queryParams1.NetworkDeviceName = vNetworkDeviceName.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okApMac {
			queryParams1.ApMac = vApMac.(string)
		}
		if okClientMac {
			queryParams1.ClientMac = vClientMac.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Devices.QueryAssuranceEvents(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 QueryAssuranceEvents", err,
				"Failure at QueryAssuranceEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesQueryAssuranceEventsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting QueryAssuranceEvents response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDetailsOfASingleAssuranceEvent")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.GetDetailsOfASingleAssuranceEventHeaderParams{}
		queryParams2 := dnacentersdkgo.GetDetailsOfASingleAssuranceEventQueryParams{}

		if okAttribute {
			queryParams2.Attribute = vAttribute.(string)
		}
		if okView {
			queryParams2.View = vView.(string)
		}
		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Devices.GetDetailsOfASingleAssuranceEvent(vvID, &headerParams2, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDetailsOfASingleAssuranceEvent", err,
				"Failure at GetDetailsOfASingleAssuranceEvent, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetDetailsOfASingleAssuranceEventItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDetailsOfASingleAssuranceEvent response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesQueryAssuranceEventsItems(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsResponse) []map[string]interface{} {
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
		respItem["candidate_a_ps"] = flattenDevicesQueryAssuranceEventsItemsCandidateAPs(item.CandidateAPs)
		respItem["replaced_device_serial_number"] = item.ReplacedDeviceSerialNumber
		respItem["old_radio_channel_list"] = item.OldRadioChannelList
		respItem["ssid"] = item.SSID
		respItem["sub_reason_description"] = item.SubReasonDescription
		respItem["wireless_client_event_end_time"] = item.WirelessClientEventEndTime
		respItem["ipv4"] = item.IPv4
		respItem["wlc_id"] = item.WlcID
		respItem["ipv6"] = item.IPv6
		respItem["missing_response_a_ps"] = flattenDevicesQueryAssuranceEventsItemsMissingResponseAPs(item.MissingResponseAPs)
		respItem["timestamp"] = item.Timestamp
		respItem["severity"] = item.Severity
		respItem["current_radio_power_level"] = item.CurrentRadioPowerLevel
		respItem["new_radio_channel_width"] = item.NewRadioChannelWidth
		respItem["assoc_snr"] = item.AssocSnr
		respItem["auth_server_ip"] = item.AuthServerIP
		respItem["child_events"] = flattenDevicesQueryAssuranceEventsItemsChildEvents(item.ChildEvents)
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
		respItem["invalid_ie_a_ps"] = flattenDevicesQueryAssuranceEventsItemsInvalidIeAPs(item.InvalidIeAPs)
		respItem["username"] = item.Username
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryAssuranceEventsItemsCandidateAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsResponseCandidateAPs) []map[string]interface{} {
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

func flattenDevicesQueryAssuranceEventsItemsMissingResponseAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsResponseMissingResponseAPs) []map[string]interface{} {
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

func flattenDevicesQueryAssuranceEventsItemsChildEvents(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsResponseChildEvents) []map[string]interface{} {
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

func flattenDevicesQueryAssuranceEventsItemsInvalidIeAPs(items *[]dnacentersdkgo.ResponseDevicesQueryAssuranceEventsResponseInvalidIeAPs) []map[string]interface{} {
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

func flattenDevicesGetDetailsOfASingleAssuranceEventItem(item *dnacentersdkgo.ResponseDevicesGetDetailsOfASingleAssuranceEventResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["candidate_a_ps"] = flattenDevicesGetDetailsOfASingleAssuranceEventItemCandidateAPs(item.CandidateAPs)
	respItem["replaced_device_serial_number"] = item.ReplacedDeviceSerialNumber
	respItem["old_radio_channel_list"] = item.OldRadioChannelList
	respItem["ssid"] = item.SSID
	respItem["sub_reason_description"] = item.SubReasonDescription
	respItem["wireless_client_event_end_time"] = item.WirelessClientEventEndTime
	respItem["ipv4"] = item.IPv4
	respItem["wlc_id"] = item.WlcID
	respItem["ipv6"] = item.IPv6
	respItem["missing_response_a_ps"] = flattenDevicesGetDetailsOfASingleAssuranceEventItemMissingResponseAPs(item.MissingResponseAPs)
	respItem["timestamp"] = item.Timestamp
	respItem["severity"] = item.Severity
	respItem["current_radio_power_level"] = item.CurrentRadioPowerLevel
	respItem["new_radio_channel_width"] = item.NewRadioChannelWidth
	respItem["assoc_snr"] = item.AssocSnr
	respItem["auth_server_ip"] = item.AuthServerIP
	respItem["child_events"] = flattenDevicesGetDetailsOfASingleAssuranceEventItemChildEvents(item.ChildEvents)
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
	respItem["invalid_ie_a_ps"] = flattenDevicesGetDetailsOfASingleAssuranceEventItemInvalidIeAPs(item.InvalidIeAPs)
	respItem["username"] = item.Username
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetDetailsOfASingleAssuranceEventItemCandidateAPs(items *[]dnacentersdkgo.ResponseDevicesGetDetailsOfASingleAssuranceEventResponseCandidateAPs) []map[string]interface{} {
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

func flattenDevicesGetDetailsOfASingleAssuranceEventItemMissingResponseAPs(items *[]dnacentersdkgo.ResponseDevicesGetDetailsOfASingleAssuranceEventResponseMissingResponseAPs) []map[string]interface{} {
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

func flattenDevicesGetDetailsOfASingleAssuranceEventItemChildEvents(items *[]dnacentersdkgo.ResponseDevicesGetDetailsOfASingleAssuranceEventResponseChildEvents) []map[string]interface{} {
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

func flattenDevicesGetDetailsOfASingleAssuranceEventItemInvalidIeAPs(items *[]dnacentersdkgo.ResponseDevicesGetDetailsOfASingleAssuranceEventResponseInvalidIeAPs) []map[string]interface{} {
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
