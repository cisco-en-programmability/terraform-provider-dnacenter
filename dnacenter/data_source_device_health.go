package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating
additional value added services.
`,

		ReadContext: dataSourceDeviceHealthRead,
		Schema: map[string]*schema.Schema{
			"device_role": &schema.Schema{
				Description: `deviceRole query parameter. CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, or AP (case insensitive)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. UTC epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"health": &schema.Schema{
				Description: `health query parameter. DNAC health catagory: POOR, FAIR, or GOOD (case insensitive)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Max number of device entries in the response (default to 50. Max at 500)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset of the first device in the returned data (Mutiple of 'limit' + 1)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. DNAC site UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. UTC epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"air_quality_health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz24": &schema.Schema{
										Description: `Ghz24`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"ap_count": &schema.Schema{
							Description: `Number of AP count
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"avg_temperature": &schema.Schema{
							Description: `Average device (switch) temperature
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"band": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"client_count": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz24": &schema.Schema{
										Description: `Ghz24`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"cpu_health": &schema.Schema{
							Description: `Device CPU health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"cpu_ulitilization": &schema.Schema{
							Description: `Device's CPU utilization
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"cpu_utilization": &schema.Schema{
							Description: `Device's CPU utilization
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"device_family": &schema.Schema{
							Description: `Device family
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_type": &schema.Schema{
							Description: `Device type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"free_memory_buffer": &schema.Schema{
							Description: `Device free memory
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"free_memory_buffer_health": &schema.Schema{
							Description: `Device free memory buffer health
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"free_timer": &schema.Schema{
							Description: `Device free timer
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"free_timer_score": &schema.Schema{
							Description: `Device free timer health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"inter_device_link_avail_fabric": &schema.Schema{
							Description: `Device uplink health
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"inter_device_link_avail_health": &schema.Schema{
							Description: `Device connectivity status
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"interface_link_err_health": &schema.Schema{
							Description: `Device (AP) error health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"interference_health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz24": &schema.Schema{
										Description: `Ghz24`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"ip_address": &schema.Schema{
							Description: `Management IP address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_count": &schema.Schema{
							Description: `Number of issues
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `Site location in which this device is assigned to
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"max_temperature": &schema.Schema{
							Description: `Max device (switch) temperature
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"memory_utilization": &schema.Schema{
							Description: `Device memory utilization
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"memory_utilization_health": &schema.Schema{
							Description: `Device memory utilization health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"model": &schema.Schema{
							Description: `Device model string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"noise_health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz24": &schema.Schema{
										Description: `Ghz24`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"os_version": &schema.Schema{
							Description: `Device OS version string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"overall_health": &schema.Schema{
							Description: `Overall health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"packet_pool": &schema.Schema{
							Description: `Device packet pool
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"packet_pool_health": &schema.Schema{
							Description: `Device packet pool
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"reachability_health": &schema.Schema{
							Description: `Device reachability in the network
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"utilization_health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz24": &schema.Schema{
										Description: `Ghz24`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio2": &schema.Schema{
										Description: `Radio2`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio3": &schema.Schema{
										Description: `Radio3`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"uuid": &schema.Schema{
							Description: `Device UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wan_link_utilization": &schema.Schema{
							Description: `WLAN link utilization
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"wqe_pools": &schema.Schema{
							Description: `Device WQE pool
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"wqe_pools_health": &schema.Schema{
							Description: `Device WQE pool health
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

func dataSourceDeviceHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceRole, okDeviceRole := d.GetOk("device_role")
	vSiteID, okSiteID := d.GetOk("site_id")
	vHealth, okHealth := d.GetOk("health")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: Devices")
		queryParams1 := dnacentersdkgo.DevicesQueryParams{}

		if okDeviceRole {
			queryParams1.DeviceRole = vDeviceRole.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okHealth {
			queryParams1.Health = vHealth.(string)
		}
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

		response1, restyResp1, err := client.Devices.Devices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 Devices", err,
				"Failure at Devices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Devices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesDevicesItems(items *[]dnacentersdkgo.ResponseDevicesDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_type"] = item.DeviceType
		respItem["cpu_utilization"] = item.CPUUtilization
		respItem["overall_health"] = item.OverallHealth
		respItem["utilization_health"] = flattenDevicesDevicesItemsUtilizationHealth(item.UtilizationHealth)
		respItem["air_quality_health"] = flattenDevicesDevicesItemsAirQualityHealth(item.AirQualityHealth)
		respItem["ip_address"] = item.IPAddress
		respItem["cpu_health"] = item.CPUHealth
		respItem["device_family"] = item.DeviceFamily
		respItem["issue_count"] = item.IssueCount
		respItem["mac_address"] = item.MacAddress
		respItem["noise_health"] = flattenDevicesDevicesItemsNoiseHealth(item.NoiseHealth)
		respItem["os_version"] = item.OsVersion
		respItem["name"] = item.Name
		respItem["interface_link_err_health"] = item.InterfaceLinkErrHealth
		respItem["memory_utilization"] = item.MemoryUtilization
		respItem["inter_device_link_avail_health"] = item.InterDeviceLinkAvailHealth
		respItem["interference_health"] = flattenDevicesDevicesItemsInterferenceHealth(item.InterferenceHealth)
		respItem["model"] = item.Model
		respItem["location"] = item.Location
		respItem["reachability_health"] = item.ReachabilityHealth
		respItem["band"] = flattenDevicesDevicesItemsBand(item.Band)
		respItem["memory_utilization_health"] = item.MemoryUtilizationHealth
		respItem["client_count"] = flattenDevicesDevicesItemsClientCount(item.ClientCount)
		respItem["avg_temperature"] = item.AvgTemperature
		respItem["max_temperature"] = item.MaxTemperature
		respItem["inter_device_link_avail_fabric"] = item.InterDeviceLinkAvailFabric
		respItem["ap_count"] = item.ApCount
		respItem["free_timer_score"] = item.FreeTimerScore
		respItem["free_timer"] = item.FreeTimer
		respItem["packet_pool_health"] = item.PacketPoolHealth
		respItem["packet_pool"] = item.PacketPool
		respItem["free_memory_buffer_health"] = item.FreeMemoryBufferHealth
		respItem["free_memory_buffer"] = item.FreeMemoryBuffer
		respItem["wqe_pools_health"] = item.WqePoolsHealth
		respItem["wqe_pools"] = item.WqePools
		respItem["wan_link_utilization"] = item.WanLinkUtilization
		respItem["cpu_ulitilization"] = item.CPUUlitilization
		respItem["uuid"] = item.UUID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesDevicesItemsUtilizationHealth(item *dnacentersdkgo.ResponseDevicesDevicesResponseUtilizationHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsAirQualityHealth(item *dnacentersdkgo.ResponseDevicesDevicesResponseAirQualityHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsNoiseHealth(item *dnacentersdkgo.ResponseDevicesDevicesResponseNoiseHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsInterferenceHealth(item *dnacentersdkgo.ResponseDevicesDevicesResponseInterferenceHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsBand(item *dnacentersdkgo.ResponseDevicesDevicesResponseBand) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsClientCount(item *dnacentersdkgo.ResponseDevicesDevicesResponseClientCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["radio2"] = item.Radio2
	respItem["radio3"] = item.Radio3
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}
