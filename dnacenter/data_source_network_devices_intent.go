package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesIntent() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- API to fetch the list of network devices using basic filters. Use the */dna/intent/api/v1/networkDevices/query* API
for advanced filtering. Refer features for more details.
`,

		ReadContext: dataSourceNetworkDevicesIntentRead,
		Schema: map[string]*schema.Schema{
			"family": &schema.Schema{
				Description: `family query parameter. Product family of the network device. For example, Switches, Routers, etc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Network device Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Min: 1, Max: 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_address": &schema.Schema{
				Description: `managementAddress query parameter. Management address of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_state": &schema.Schema{
				Description: `managementState query parameter. The status of the network device's manageability. Available statuses are MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachability_status": &schema.Schema{
				Description: `reachabilityStatus query parameter. Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Description: `role query parameter. Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Serial number of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by. Available values : id, managementAddress, dnsResolvedManagementIpAddress, hostname, macAddress, type, family, series, platformids, softwareType, softwareVersion, vendor, bootTime, role, roleSource, apEthernetMacAddress, apManagerInterfaceIpAddress, apWlcIpAddress, deviceSupportLevel, reachabilityFailureReason, resyncStartTime, resyncEndTime, resyncReasons, pendingResyncRequestCount, pendingResyncRequestReasons, resyncIntervalSource, resyncIntervalMinutes
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_device": &schema.Schema{
				Description: `stackDevice query parameter. Flag indicating if the device is a stack device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"views": &schema.Schema{
				Description: `views query parameter. The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Refer features for more details. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_ethernet_mac_address": &schema.Schema{
							Description: `Ethernet MAC address of the AP network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_manager_interface_ip_address": &schema.Schema{
							Description: `Management IP address of the AP network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_wlc_ip_address": &schema.Schema{
							Description: `Management IP address of the WLC on which AP is associated to
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"boot_time": &schema.Schema{
							Description: `The time at which the network device was last rebooted or powered on represented as epoch in milliseconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"device_support_level": &schema.Schema{
							Description: `The level of support Catalyst Center provides for the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dns_resolved_management_ip_address": &schema.Schema{
							Description: `DNS-resolved management IP address of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Description: `Error code indicating the reason for the last resync failure
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_description": &schema.Schema{
							Description: `Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"family": &schema.Schema{
							Description: `Product family of the network device. For example, Switches, Routers, etc
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Description: `Hostname of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique identifier of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_successful_resync_reasons": &schema.Schema{
							Description: `List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_address": &schema.Schema{
							Description: `Management address of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_state": &schema.Schema{
							Description: `The status of the network device's manageability. Refer features for more details.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pending_resync_request_count": &schema.Schema{
							Description: `Number of pending resync requests for the device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"pending_resync_request_reasons": &schema.Schema{
							Description: `List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"platform_ids": &schema.Schema{
							Description: `Platform identifier of the network device
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"reachability_failure_reason": &schema.Schema{
							Description: `Reason for reachability failure. This message that provides more information about the reachability failure.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_status": &schema.Schema{
							Description: `Reachability status of the network device. Refer features for more details
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"resync_end_time": &schema.Schema{
							Description: `End time for the last resync represented as epoch in milliseconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"resync_interval_minutes": &schema.Schema{
							Description: `The duration in minutes between the periodic resync attempts for the device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"resync_interval_source": &schema.Schema{
							Description: `Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"resync_reasons": &schema.Schema{
							Description: `List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"resync_requested_by_apps": &schema.Schema{
							Description: `List of applications that requested the last/ongoing resync on the device
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"resync_start_time": &schema.Schema{
							Description: `Start time for the last/ongoing resync represented as epoch in milliseconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"role": &schema.Schema{
							Description: `Role assigned to the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role_source": &schema.Schema{
							Description: `Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_numbers": &schema.Schema{
							Description: `Serial number of the network device. In case of stack device, there will be multiple serial numbers
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"series": &schema.Schema{
							Description: `The model range or series of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_contact": &schema.Schema{
							Description: `SNMP contact of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_location": &schema.Schema{
							Description: `SNMP location of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_type": &schema.Schema{
							Description: `Type of software running on the network device. For example, IOS-XE, etc.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Version of the software running on the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"stack_device": &schema.Schema{
							Description: `Flag indicating if the network device is a stack device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Inventory related status of the network device. Refer features for more details
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"user_defined_fields": &schema.Schema{
							Description: `Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"vendor": &schema.Schema{
							Description: `Vendor of the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesIntentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vManagementAddress, okManagementAddress := d.GetOk("management_address")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vFamily, okFamily := d.GetOk("family")
	vStackDevice, okStackDevice := d.GetOk("stack_device")
	vRole, okRole := d.GetOk("role")
	vStatus, okStatus := d.GetOk("status")
	vReachabilityStatus, okReachabilityStatus := d.GetOk("reachability_status")
	vManagementState, okManagementState := d.GetOk("management_state")
	vViews, okViews := d.GetOk("views")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveNetworkDevices")
		queryParams1 := dnacentersdkgo.RetrieveNetworkDevicesQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okManagementAddress {
			queryParams1.ManagementAddress = vManagementAddress.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okStackDevice {
			queryParams1.StackDevice = vStackDevice.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okReachabilityStatus {
			queryParams1.ReachabilityStatus = vReachabilityStatus.(string)
		}
		if okManagementState {
			queryParams1.ManagementState = vManagementState.(string)
		}
		if okViews {
			queryParams1.Views = vViews.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Devices.RetrieveNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveNetworkDevices", err,
				"Failure at RetrieveNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesRetrieveNetworkDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrieveNetworkDevicesItems(items *[]dnacentersdkgo.ResponseDevicesRetrieveNetworkDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["management_address"] = item.ManagementAddress
		respItem["dns_resolved_management_ip_address"] = item.DNSResolvedManagementIPAddress
		respItem["hostname"] = item.Hostname
		respItem["mac_address"] = item.MacAddress
		respItem["serial_numbers"] = item.SerialNumbers
		respItem["type"] = item.Type
		respItem["family"] = item.Family
		respItem["series"] = item.Series
		respItem["status"] = item.Status
		respItem["platform_ids"] = item.PlatformIDs
		respItem["software_type"] = item.SoftwareType
		respItem["software_version"] = item.SoftwareVersion
		respItem["vendor"] = item.Vendor
		respItem["stack_device"] = boolPtrToString(item.StackDevice)
		respItem["boot_time"] = item.BootTime
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["ap_ethernet_mac_address"] = item.ApEthernetMacAddress
		respItem["ap_manager_interface_ip_address"] = item.ApManagerInterfaceIPAddress
		respItem["ap_wlc_ip_address"] = item.ApWlcIPAddress
		respItem["device_support_level"] = item.DeviceSupportLevel
		respItem["snmp_location"] = item.SNMPLocation
		respItem["snmp_contact"] = item.SNMPContact
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["management_state"] = item.ManagementState
		respItem["last_successful_resync_reasons"] = item.LastSuccessfulResyncReasons
		respItem["resync_start_time"] = item.ResyncStartTime
		respItem["resync_end_time"] = item.ResyncEndTime
		respItem["resync_reasons"] = item.ResyncReasons
		respItem["resync_requested_by_apps"] = item.ResyncRequestedByApps
		respItem["pending_resync_request_count"] = item.PendingResyncRequestCount
		respItem["pending_resync_request_reasons"] = item.PendingResyncRequestReasons
		respItem["resync_interval_source"] = item.ResyncIntervalSource
		respItem["resync_interval_minutes"] = item.ResyncIntervalMinutes
		respItem["error_code"] = item.ErrorCode
		respItem["error_description"] = item.ErrorDescription
		respItem["user_defined_fields"] = flattenDevicesRetrieveNetworkDevicesItemsUserDefinedFields(item.UserDefinedFields)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesRetrieveNetworkDevicesItemsUserDefinedFields(item *dnacentersdkgo.ResponseDevicesRetrieveNetworkDevicesResponseUserDefinedFields) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
