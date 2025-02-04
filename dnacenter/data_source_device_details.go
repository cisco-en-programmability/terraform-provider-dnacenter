package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.
`,

		ReadContext: dataSourceDeviceDetailsRead,
		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Description: `identifier query parameter. One of "macAddress", "nwDeviceName", "uuid" (case insensitive)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"search_by": &schema.Schema{
				Description: `searchBy query parameter. MAC Address, device name, or UUID of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. UTC timestamp of device data in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"halast_reset_reason": &schema.Schema{
							Description: `Last HA reset reason
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"admin_state": &schema.Schema{
							Description: `Device (AP) admin state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"air_quality": &schema.Schema{
							Description: `Device (AP) WIFI air quality
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"air_quality_score": &schema.Schema{
							Description: `Device (AP) air quality health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ap_group": &schema.Schema{
							Description: `Device (AP) AP group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_type": &schema.Schema{
							Description: `Ap Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"avg_temperature": &schema.Schema{
							Description: `Device's average temperature
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"collection_status": &schema.Schema{
							Description: `Device's telemetry data collection status for DNAC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"communication_state": &schema.Schema{
							Description: `Device communication state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"connected_time": &schema.Schema{
							Description: `UTC timestamp
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"connectivity_status": &schema.Schema{
							Description: `Device connectivity status
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"cpu": &schema.Schema{
							Description: `Device CPU utilization
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"cpu_score": &schema.Schema{
							Description: `Device's CPU usage score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"device_group_hierarchy_id": &schema.Schema{
							Description: `Device group site hierarchy UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_series": &schema.Schema{
							Description: `Device series string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ethernet_mac": &schema.Schema{
							Description: `Device (AP) ethernet MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"feature_flag_list": &schema.Schema{
							Description: `List of device feature capabilities
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"flex_group": &schema.Schema{
							Description: `Deivce (A) flexconnect group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"free_mbuf": &schema.Schema{
							Description: `Free memory buffer of the device
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"free_mbuf_score": &schema.Schema{
							Description: `Free memory buffer health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"free_timer": &schema.Schema{
							Description: `Free timer of the device
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"free_timer_score": &schema.Schema{
							Description: `Free Timer Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"ha_status": &schema.Schema{
							Description: `Device's HA status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"home_ap_enabled": &schema.Schema{
							Description: `Home Ap Enabled`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"icap_capability": &schema.Schema{
							Description: `Device (AP) ICAP capability bit values
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interference": &schema.Schema{
							Description: `Device (AP) WIFI signal interference
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interference_score": &schema.Schema{
							Description: `Device (AP) WIFI signal interference health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ip_addr_management_ip_addr": &schema.Schema{
							Description: `Device's management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_boot_time": &schema.Schema{
							Description: `Device's last boot UTC timestamp
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"led_flash_enabled": &schema.Schema{
							Description: `Device (AP) LED flash
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"led_flash_seconds": &schema.Schema{
							Description: `LED flash seconds
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `Device's site hierarchy UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `Device MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"maintenance_mode": &schema.Schema{
							Description: `Whether device is in maintenance mode
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_addr": &schema.Schema{
							Description: `Management IP address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"max_temperature": &schema.Schema{
							Description: `Device's max temperature
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"memory": &schema.Schema{
							Description: `Device memory utilization
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"memory_score": &schema.Schema{
							Description: `Device's memory usage score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"mode": &schema.Schema{
							Description: `Device mode (AP)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"noise": &schema.Schema{
							Description: `Device (AP) WIFI signal noise
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"noise_score": &schema.Schema{
							Description: `Device (AP) WIFI signal noise health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"nw_device_family": &schema.Schema{
							Description: `Device faimly string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"nw_device_id": &schema.Schema{
							Description: `Device's UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"nw_device_name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"nw_device_role": &schema.Schema{
							Description: `Device role
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"nw_device_type": &schema.Schema{
							Description: `Device type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"op_state": &schema.Schema{
							Description: `Operation state of device (AP)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"os_type": &schema.Schema{
							Description: `Device's OS type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"overall_health": &schema.Schema{
							Description: `Device's overall health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"packet_pool": &schema.Schema{
							Description: `Packet pool of the device
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"packet_pool_score": &schema.Schema{
							Description: `Device packet pool health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"platform_id": &schema.Schema{
							Description: `Device's platform ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_tag_name": &schema.Schema{
							Description: `Device (AP) policy tag
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_calendar_profile": &schema.Schema{
							Description: `Device (AP) power calendar profile name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_mode": &schema.Schema{
							Description: `Device's power mode
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_profile": &schema.Schema{
							Description: `Device (AP) power profile name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_save_mode": &schema.Schema{
							Description: `Device power save mode
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_save_mode_capable": &schema.Schema{
							Description: `Device (AP) power save mode capability
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_type": &schema.Schema{
							Description: `Device (AP) power type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol": &schema.Schema{
							Description: `Protocol code
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"redundancy_mode": &schema.Schema{
							Description: `Device redundancy mode
`,
							Type:     schema.TypeString,
							Computed: true,
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
							Description: `Redundancy state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"redundancy_state_derived": &schema.Schema{
							Description: `Derived redundancy state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"regulatory_domain": &schema.Schema{
							Description: `Device (AP) WIFI domain
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reset_reason": &schema.Schema{
							Description: `Device reset reason
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rf_tag_name": &schema.Schema{
							Description: `Device (AP) RF tag name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ring_status": &schema.Schema{
							Description: `Device's ring status
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Device serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_hierarchy_graph_id": &schema.Schema{
							Description: `Site hierarchy UUID in which device is assigned to
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_tag_name": &schema.Schema{
							Description: `Device (AP) site tag name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Device's software version string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"stack_type": &schema.Schema{
							Description: `Device stack type (applicable for stackable devices)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sub_mode": &schema.Schema{
							Description: `Device submode
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag_id_list": &schema.Schema{
							Description: `Tag ID List`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"timestamp": &schema.Schema{
							Description: `UTC timestamp of the device health data
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"up_time": &schema.Schema{
							Description: `Device up time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"utilization": &schema.Schema{
							Description: `Device utilization
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"utilization_score": &schema.Schema{
							Description: `Device utilization health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wqe": &schema.Schema{
							Description: `WQE of the device
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"wqe_score": &schema.Schema{
							Description: `WQE health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTimestamp, okTimestamp := d.GetOk("timestamp")
	vIDentifier := d.Get("identifier")
	vSearchBy := d.Get("search_by")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceDetail")
		queryParams1 := dnacentersdkgo.GetDeviceDetailQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(float64)
		}
		queryParams1.IDentifier = vIDentifier.(string)

		queryParams1.SearchBy = vSearchBy.(string)

		response1, restyResp1, err := client.Devices.GetDeviceDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceDetail", err,
				"Failure at GetDeviceDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceDetailItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceDetail response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceDetailItem(item *dnacentersdkgo.ResponseDevicesGetDeviceDetailResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["noise_score"] = item.NoiseScore
	respItem["policy_tag_name"] = item.PolicyTagName
	respItem["interference_score"] = item.InterferenceScore
	respItem["op_state"] = item.OpState
	respItem["power_save_mode"] = item.PowerSaveMode
	respItem["mode"] = item.Mode
	respItem["reset_reason"] = item.ResetReason
	respItem["nw_device_role"] = item.NwDeviceRole
	respItem["protocol"] = item.Protocol
	respItem["power_mode"] = item.PowerMode
	respItem["connected_time"] = item.ConnectedTime
	respItem["ring_status"] = boolPtrToString(item.RingStatus)
	respItem["led_flash_seconds"] = item.LedFlashSeconds
	respItem["ip_addr_management_ip_addr"] = item.IPAddrManagementIPAddr
	respItem["stack_type"] = item.StackType
	respItem["sub_mode"] = item.SubMode
	respItem["serial_number"] = item.SerialNumber
	respItem["nw_device_name"] = item.NwDeviceName
	respItem["device_group_hierarchy_id"] = item.DeviceGroupHierarchyID
	respItem["cpu"] = item.CPU
	respItem["utilization"] = item.Utilization
	respItem["nw_device_id"] = item.NwDeviceID
	respItem["site_hierarchy_graph_id"] = item.SiteHierarchyGraphID
	respItem["nw_device_family"] = item.NwDeviceFamily
	respItem["mac_address"] = item.MacAddress
	respItem["home_ap_enabled"] = item.HomeApEnabled
	respItem["device_series"] = item.DeviceSeries
	respItem["collection_status"] = item.CollectionStatus
	respItem["utilization_score"] = item.UtilizationScore
	respItem["maintenance_mode"] = boolPtrToString(item.MaintenanceMode)
	respItem["interference"] = item.Interference
	respItem["software_version"] = item.SoftwareVersion
	respItem["tag_id_list"] = flattenDevicesGetDeviceDetailItemTagIDList(item.TagIDList)
	respItem["power_type"] = item.PowerType
	respItem["overall_health"] = item.OverallHealth
	respItem["management_ip_addr"] = item.ManagementIPAddr
	respItem["memory"] = item.Memory
	respItem["communication_state"] = item.CommunicationState
	respItem["ap_type"] = item.ApType
	respItem["admin_state"] = item.AdminState
	respItem["noise"] = item.Noise
	respItem["icap_capability"] = item.IcapCapability
	respItem["regulatory_domain"] = item.RegulatoryDomain
	respItem["ethernet_mac"] = item.EthernetMac
	respItem["nw_device_type"] = item.NwDeviceType
	respItem["air_quality"] = item.AirQuality
	respItem["rf_tag_name"] = item.RfTagName
	respItem["site_tag_name"] = item.SiteTagName
	respItem["platform_id"] = item.PlatformID
	respItem["up_time"] = item.UpTime
	respItem["memory_score"] = item.MemoryScore
	respItem["power_save_mode_capable"] = item.PowerSaveModeCapable
	respItem["power_profile"] = item.PowerProfile
	respItem["air_quality_score"] = item.AirQualityScore
	respItem["location"] = item.Location
	respItem["flex_group"] = item.FlexGroup
	respItem["last_boot_time"] = item.LastBootTime
	respItem["power_calendar_profile"] = item.PowerCalendarProfile
	respItem["connectivity_status"] = item.ConnectivityStatus
	respItem["led_flash_enabled"] = item.LedFlashEnabled
	respItem["cpu_score"] = item.CPUScore
	respItem["avg_temperature"] = item.AvgTemperature
	respItem["max_temperature"] = item.MaxTemperature
	respItem["ha_status"] = item.HaStatus
	respItem["os_type"] = item.OsType
	respItem["timestamp"] = item.Timestamp
	respItem["ap_group"] = item.ApGroup
	respItem["redundancy_mode"] = item.RedundancyMode
	respItem["feature_flag_list"] = item.FeatureFlagList
	respItem["free_mbuf_score"] = item.FreeMbufScore
	respItem["halast_reset_reason"] = item.HALastResetReason
	respItem["wqe_score"] = item.WqeScore
	respItem["redundancy_peer_state_derived"] = item.RedundancyPeerStateDerived
	respItem["free_timer_score"] = item.FreeTimerScore
	respItem["redundancy_peer_state"] = item.RedundancyPeerState
	respItem["redundancy_state_derived"] = item.RedundancyStateDerived
	respItem["redundancy_state"] = item.RedundancyState
	respItem["packet_pool_score"] = item.PacketPoolScore
	respItem["free_timer"] = item.FreeTimer
	respItem["packet_pool"] = item.PacketPool
	respItem["wqe"] = item.Wqe
	respItem["free_mbuf"] = item.FreeMbuf
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetDeviceDetailItemTagIDList(items *[]dnacentersdkgo.ResponseDevicesGetDeviceDetailResponseTagIDList) []interface{} {
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
