package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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
				Description: `identifier query parameter. One of keywords : macAddress or uuid or nwDeviceName
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"search_by": &schema.Schema{
				Description: `searchBy query parameter. MAC Address or Device Name value or UUID of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the device data is required
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"halast_reset_reason": &schema.Schema{
							Description: `H A Last Reset Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"haprimary_power_status": &schema.Schema{
							Description: `H A Primary Power Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"hasecondary_power_status": &schema.Schema{
							Description: `H A Secondary Power Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"air_quality": &schema.Schema{
							Description: `Air Quality`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"air_quality_score": &schema.Schema{
							Description: `Air Quality Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"client_count": &schema.Schema{
							Description: `Client Count`,
							Type:        schema.TypeString,
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

						"cpu": &schema.Schema{
							Description: `Cpu`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"cpu_score": &schema.Schema{
							Description: `Cpu Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"device_series": &schema.Schema{
							Description: `Device Series`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"free_mbuf": &schema.Schema{
							Description: `Free Mbuf`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"free_mbuf_score": &schema.Schema{
							Description: `Free Mbuf Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"free_timer": &schema.Schema{
							Description: `Free Timer`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"free_timer_score": &schema.Schema{
							Description: `Free Timer Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"interference": &schema.Schema{
							Description: `Interference`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interference_score": &schema.Schema{
							Description: `Interference Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"location": &schema.Schema{
							Description: `Location`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"management_ip_addr": &schema.Schema{
							Description: `Management Ip Addr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"memory": &schema.Schema{
							Description: `Memory`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"memory_score": &schema.Schema{
							Description: `Memory Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"noise": &schema.Schema{
							Description: `Noise`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"noise_score": &schema.Schema{
							Description: `Noise Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"nw_device_family": &schema.Schema{
							Description: `Nw Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"nw_device_id": &schema.Schema{
							Description: `Nw Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"nw_device_name": &schema.Schema{
							Description: `Nw Device Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"nw_device_role": &schema.Schema{
							Description: `Nw Device Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"nw_device_type": &schema.Schema{
							Description: `Nw Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"os_type": &schema.Schema{
							Description: `Os Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"overall_health": &schema.Schema{
							Description: `Overall Health`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"packet_pool": &schema.Schema{
							Description: `Packet Pool`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"packet_pool_score": &schema.Schema{
							Description: `Packet Pool Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform Id`,
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

						"redundancy_state": &schema.Schema{
							Description: `Redundancy State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"redundancy_unit": &schema.Schema{
							Description: `Redundancy Unit`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"software_version": &schema.Schema{
							Description: `Software Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"utilization": &schema.Schema{
							Description: `Utilization`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"utilization_score": &schema.Schema{
							Description: `Utilization Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wqe": &schema.Schema{
							Description: `Wqe`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wqe_score": &schema.Schema{
							Description: `Wqe Score`,
							Type:        schema.TypeInt,
							Computed:    true,
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
	vSearchBy := d.Get("search_by")
	vIDentifier := d.Get("identifier")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceDetail")
		queryParams1 := dnacentersdkgo.GetDeviceDetailQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(string)
		}
		queryParams1.SearchBy = vSearchBy.(string)

		queryParams1.IDentifier = vIDentifier.(string)

		response1, restyResp1, err := client.Devices.GetDeviceDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceDetail", err,
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
	respItem["halast_reset_reason"] = item.HALastResetReason
	respItem["management_ip_addr"] = item.ManagementIPAddr
	respItem["haprimary_power_status"] = item.HAPrimaryPowerStatus
	respItem["redundancy_mode"] = item.RedundancyMode
	respItem["communication_state"] = item.CommunicationState
	respItem["nw_device_name"] = item.NwDeviceName
	respItem["redundancy_unit"] = item.RedundancyUnit
	respItem["platform_id"] = item.PlatformID
	respItem["redundancy_peer_state"] = item.RedundancyPeerState
	respItem["nw_device_id"] = item.NwDeviceID
	respItem["redundancy_state"] = item.RedundancyState
	respItem["nw_device_role"] = item.NwDeviceRole
	respItem["nw_device_family"] = item.NwDeviceFamily
	respItem["mac_address"] = item.MacAddress
	respItem["collection_status"] = item.CollectionStatus
	respItem["device_series"] = item.DeviceSeries
	respItem["os_type"] = item.OsType
	respItem["client_count"] = item.ClientCount
	respItem["hasecondary_power_status"] = item.HASecondaryPowerStatus
	respItem["software_version"] = item.SoftwareVersion
	respItem["nw_device_type"] = item.NwDeviceType
	respItem["overall_health"] = item.OverallHealth
	respItem["memory_score"] = item.MemoryScore
	respItem["cpu_score"] = item.CPUScore
	respItem["noise_score"] = item.NoiseScore
	respItem["utilization_score"] = item.UtilizationScore
	respItem["air_quality_score"] = item.AirQualityScore
	respItem["interference_score"] = item.InterferenceScore
	respItem["wqe_score"] = item.WqeScore
	respItem["free_mbuf_score"] = item.FreeMbufScore
	respItem["packet_pool_score"] = item.PacketPoolScore
	respItem["free_timer_score"] = item.FreeTimerScore
	respItem["memory"] = item.Memory
	respItem["cpu"] = item.CPU
	respItem["noise"] = item.Noise
	respItem["utilization"] = item.Utilization
	respItem["air_quality"] = item.AirQuality
	respItem["interference"] = item.Interference
	respItem["wqe"] = item.Wqe
	respItem["free_mbuf"] = item.FreeMbuf
	respItem["packet_pool"] = item.PacketPool
	respItem["free_timer"] = item.FreeTimer
	respItem["location"] = item.Location
	respItem["timestamp"] = item.Timestamp
	return []map[string]interface{}{
		respItem,
	}
}
