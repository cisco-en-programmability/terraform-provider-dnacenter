package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- API to fetch the details of network device using the *id*. Use the */dna/intent/api/v1/networkDevices/query* API for
advanced filtering. The API supports views to fetch only the required fields. Refer features for more details.
`,

		ReadContext: dataSourceNetworkDevicesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Unique identifier for the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"views": &schema.Schema{
				Description: `views query parameter. The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
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

func dataSourceNetworkDevicesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vViews, okViews := d.GetOk("views")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDetailsOfASingleNetworkDevice")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetDetailsOfASingleNetworkDeviceQueryParams{}

		if okViews {
			queryParams1.Views = vViews.(string)
		}

		response1, restyResp1, err := client.Devices.GetDetailsOfASingleNetworkDevice(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDetailsOfASingleNetworkDevice", err,
				"Failure at GetDetailsOfASingleNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDetailsOfASingleNetworkDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDetailsOfASingleNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDetailsOfASingleNetworkDeviceItem(item *dnacentersdkgo.ResponseDevicesGetDetailsOfASingleNetworkDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["user_defined_fields"] = flattenDevicesGetDetailsOfASingleNetworkDeviceItemUserDefinedFields(item.UserDefinedFields)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetDetailsOfASingleNetworkDeviceItemUserDefinedFields(item *dnacentersdkgo.ResponseDevicesGetDetailsOfASingleNetworkDeviceResponseUserDefinedFields) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
