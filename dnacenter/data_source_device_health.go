package dnacenter

import (
	"context"

	"log"

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
				Description: `deviceRole query parameter. The device role (One of CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, AP)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. UTC epoch time in miliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"health": &schema.Schema{
				Description: `health query parameter. The device overall health (One of POOR, FAIR, GOOD)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Max number of device entries in the response (default to 50.  Max at 1000)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset of the first device in the returned data
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Assurance site UUID value
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
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio0": &schema.Schema{
										Description: `Radio0`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"cpu_health": &schema.Schema{
							Description: `Cpu Health`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"cpu_ulitilization": &schema.Schema{
							Description: `Cpu Ulitilization`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"device_family": &schema.Schema{
							Description: `Device Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inter_device_link_avail_health": &schema.Schema{
							Description: `Inter Device Link Avail Health`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"interface_link_err_health": &schema.Schema{
							Description: `Interface Link Err Health`,
							Type:        schema.TypeInt,
							Computed:    true,
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
								},
							},
						},

						"ip_address": &schema.Schema{
							Description: `Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"issue_count": &schema.Schema{
							Description: `Issue Count`,
							Type:        schema.TypeFloat,
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

						"memory_utilization": &schema.Schema{
							Description: `Memory Utilization`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"memory_utilization_health": &schema.Schema{
							Description: `Memory Utilization Health`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"model": &schema.Schema{
							Description: `Model`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"noise_health": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ghz50": &schema.Schema{
										Description: `Ghz50`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio1": &schema.Schema{
										Description: `Radio1`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"os_version": &schema.Schema{
							Description: `Os Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"overall_health": &schema.Schema{
							Description: `Overall Health`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"reachability_health": &schema.Schema{
							Description: `Reachability Health`,
							Type:        schema.TypeString,
							Computed:    true,
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
								},
							},
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
		log.Printf("[DEBUG] Selected method 1: Devices")
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
				"Failure when executing Devices", err,
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
		respItem["name"] = item.Name
		respItem["model"] = item.Model
		respItem["os_version"] = item.OsVersion
		respItem["ip_address"] = item.IPAddress
		respItem["overall_health"] = item.OverallHealth
		respItem["issue_count"] = item.IssueCount
		respItem["location"] = item.Location
		respItem["device_family"] = item.DeviceFamily
		respItem["device_type"] = item.DeviceType
		respItem["mac_address"] = item.MacAddress
		respItem["interface_link_err_health"] = item.InterfaceLinkErrHealth
		respItem["cpu_ulitilization"] = item.CPUUlitilization
		respItem["cpu_health"] = item.CPUHealth
		respItem["memory_utilization_health"] = item.MemoryUtilizationHealth
		respItem["memory_utilization"] = item.MemoryUtilization
		respItem["inter_device_link_avail_health"] = item.InterDeviceLinkAvailHealth
		respItem["reachability_health"] = item.ReachabilityHealth
		respItem["client_count"] = flattenDevicesDevicesItemsClientCount(item.ClientCount)
		respItem["interference_health"] = flattenDevicesDevicesItemsInterferenceHealth(item.InterferenceHealth)
		respItem["noise_health"] = flattenDevicesDevicesItemsNoiseHealth(item.NoiseHealth)
		respItem["air_quality_health"] = flattenDevicesDevicesItemsAirQualityHealth(item.AirQualityHealth)
		respItem["utilization_health"] = flattenDevicesDevicesItemsUtilizationHealth(item.UtilizationHealth)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesDevicesItemsClientCount(item *dnacentersdkgo.ResponseDevicesDevicesResponseClientCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
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
	respItem["radio1"] = item.Radio1
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
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesDevicesItemsUtilizationHealth(item *dnacentersdkgo.ResponseDevicesDevicesResponseUtilizationHealth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["radio0"] = item.Radio0
	respItem["radio1"] = item.Radio1
	respItem["ghz24"] = item.Ghz24
	respItem["ghz50"] = item.Ghz50

	return []map[string]interface{}{
		respItem,
	}

}
