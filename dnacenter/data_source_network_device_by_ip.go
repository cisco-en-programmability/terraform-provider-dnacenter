package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceByIP() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the network device by specified IP address
`,

		ReadContext: dataSourceNetworkDeviceByIPRead,
		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Description: `ipAddress path parameter. Device IP address
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceNetworkDeviceByIPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vIPAddress := d.Get("ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDeviceByIP")
		vvIPAddress := vIPAddress.(string)

		response1, restyResp1, err := client.Devices.GetNetworkDeviceByIP(vvIPAddress)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDeviceByIP", err,
				"Failure at GetNetworkDeviceByIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetNetworkDeviceByIPItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDeviceByIP response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
