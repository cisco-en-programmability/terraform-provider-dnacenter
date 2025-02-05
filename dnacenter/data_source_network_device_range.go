package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceRange() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the list of network devices for the given pagination range. The maximum number of records that can be
retrieved is 500
`,

		ReadContext: dataSourceNetworkDeviceRangeRead,
		Schema: map[string]*schema.Schema{
			"records_to_return": &schema.Schema{
				Description: `recordsToReturn path parameter. Number of records to return [1<= recordsToReturn <= 500]
`,
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_index": &schema.Schema{
				Description: `startIndex path parameter. Start index [>=1]
`,
				Type:     schema.TypeInt,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_ethernet_mac_address": &schema.Schema{
							Description: `AccessPoint Ethernet MacAddress of AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_manager_interface_ip": &schema.Schema{
							Description: `IP address of WLC on AP manager interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"associated_wlc_ip": &schema.Schema{
							Description: `Associated Wlc Ip address of the AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"boot_date_time": &schema.Schema{
							Description: `Device boot time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_interval": &schema.Schema{
							Description: `Re sync Interval of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_status": &schema.Schema{
							Description: `Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `System description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_support_level": &schema.Schema{
							Description: `Support level of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dns_resolved_management_address": &schema.Schema{
							Description: `Specifies the resolved ip address of dns name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Description: `Inventory status error code
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_description": &schema.Schema{
							Description: `Inventory status description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"family": &schema.Schema{
							Description: `Family of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Instance Uuid of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_count": &schema.Schema{
							Description: `Number of interfaces on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_status_detail": &schema.Schema{
							Description: `Status detail of inventory sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_device_resync_start_time": &schema.Schema{
							Description: `Start time for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Time in epoch when the network device info last got updated
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Description: `Time when the network device info last got updated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_count": &schema.Schema{
							Description: `Number of linecards on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_id": &schema.Schema{
							Description: `IDs of linecards of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `[Deprecated] Location ID that is associated with the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location_name": &schema.Schema{
							Description: `[Deprecated] Name of the associated location
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_atleast_once": &schema.Schema{
							Description: `Indicates if device went into Managed state atleast once
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `IP address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_state": &schema.Schema{
							Description: `Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"memory_size": &schema.Schema{
							Description: `Processor memory size
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pending_sync_requests_count": &schema.Schema{
							Description: `Count of pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform ID of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_failure_reason": &schema.Schema{
							Description: `Failure reason for unreachable devices
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_status": &schema.Schema{
							Description: `Device reachability status as Reachable / Unreachable
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_device_resync": &schema.Schema{
							Description: `Reason for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_pending_sync_requests": &schema.Schema{
							Description: `Reasons for pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role": &schema.Schema{
							Description: `Role of device as access, distribution, border router, core
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role_source": &schema.Schema{
							Description: `Role source as manual / auto
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial number of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Description: `Device series
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_contact": &schema.Schema{
							Description: `SNMP contact on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_location": &schema.Schema{
							Description: `SNMP location on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_type": &schema.Schema{
							Description: `Software type on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Software version on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag_count": &schema.Schema{
							Description: `Number of tags associated with the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tunnel_udp_port": &schema.Schema{
							Description: `Mobility protocol port is stored as tunneludpport for WLC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"up_time": &schema.Schema{
							Description: `Time that shows for how long the device has been up
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"uptime_seconds": &schema.Schema{
							Description: `Uptime in Seconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"vendor": &schema.Schema{
							Description: `Vendor details
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"waas_device_mode": &schema.Schema{
							Description: `WAAS device mode
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

func dataSourceNetworkDeviceRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartIndex := d.Get("start_index")
	vRecordsToReturn := d.Get("records_to_return")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceByPaginationRange")
		vvStartIndex := vStartIndex.(int)
		vvRecordsToReturn := vRecordsToReturn.(int)

		response1, restyResp1, err := client.Devices.GetNetworkDeviceByPaginationRange(vvStartIndex, vvRecordsToReturn)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDeviceByPaginationRange", err,
				"Failure at GetNetworkDeviceByPaginationRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetNetworkDeviceByPaginationRangeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByPaginationRange response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetNetworkDeviceByPaginationRangeItems(items *[]dnacentersdkgo.ResponseDevicesGetNetworkDeviceByPaginationRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
		respItem["associated_wlc_ip"] = item.AssociatedWlcIP
		respItem["boot_date_time"] = item.BootDateTime
		respItem["collection_interval"] = item.CollectionInterval
		respItem["collection_status"] = item.CollectionStatus
		respItem["error_code"] = item.ErrorCode
		respItem["error_description"] = item.ErrorDescription
		respItem["family"] = item.Family
		respItem["hostname"] = item.Hostname
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["interface_count"] = item.InterfaceCount
		respItem["inventory_status_detail"] = item.InventoryStatusDetail
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["last_updated"] = item.LastUpdated
		respItem["line_card_count"] = item.LineCardCount
		respItem["line_card_id"] = item.LineCardID
		respItem["location"] = item.Location
		respItem["location_name"] = item.LocationName
		respItem["mac_address"] = item.MacAddress
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["memory_size"] = item.MemorySize
		respItem["platform_id"] = item.PlatformID
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["serial_number"] = item.SerialNumber
		respItem["series"] = item.Series
		respItem["snmp_contact"] = item.SNMPContact
		respItem["snmp_location"] = item.SNMPLocation
		respItem["software_type"] = item.SoftwareType
		respItem["software_version"] = item.SoftwareVersion
		respItem["tag_count"] = item.TagCount
		respItem["tunnel_udp_port"] = item.TunnelUDPPort
		respItem["type"] = item.Type
		respItem["up_time"] = item.UpTime
		respItem["waas_device_mode"] = item.WaasDeviceMode
		respItem["dns_resolved_management_address"] = item.DNSResolvedManagementAddress
		respItem["ap_ethernet_mac_address"] = item.ApEthernetMacAddress
		respItem["vendor"] = item.Vendor
		respItem["reasons_for_pending_sync_requests"] = item.ReasonsForPendingSyncRequests
		respItem["pending_sync_requests_count"] = item.PendingSyncRequestsCount
		respItem["reasons_for_device_resync"] = item.ReasonsForDeviceResync
		respItem["last_device_resync_start_time"] = item.LastDeviceResyncStartTime
		respItem["uptime_seconds"] = item.UptimeSeconds
		respItem["managed_atleast_once"] = boolPtrToString(item.ManagedAtleastOnce)
		respItem["device_support_level"] = item.DeviceSupportLevel
		respItem["management_state"] = item.ManagementState
		respItem["description"] = item.Description
		respItems = append(respItems, respItem)
	}
	return respItems
}
