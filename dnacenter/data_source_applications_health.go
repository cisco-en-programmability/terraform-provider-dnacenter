package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationsHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Applications.

- Intent API to get a list of applications for a specific site, a device, or a client device's MAC address. For a
combination of a specific application with site and/or device the API gets list of issues/devices/endpoints.
`,

		ReadContext: dataSourceApplicationsHealthRead,
		Schema: map[string]*schema.Schema{
			"application_health": &schema.Schema{
				Description: `applicationHealth query parameter. Application health category (POOR, FAIR, or GOOD.  Optionally use with siteId only)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_name": &schema.Schema{
				Description: `applicationName query parameter. The name of the application to get information on
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. Assurance device UUID value (Cannot be submitted together with siteId and clientMac)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Ending epoch time in milliseconds of time window
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The max number of application entries in returned data [1, 1000] (optionally used with siteId only)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Client device's MAC address (Cannot be submitted together with siteId and deviceId)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset of the first application in the returned data (optionally used with siteId only)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Assurance site UUID value (Cannot be submitted together with deviceId and clientMac)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Starting epoch time in milliseconds of time window
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application": &schema.Schema{
							Description: `Issue reltaed application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"application_server_latency": &schema.Schema{
							Description: `Latency of application server
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"average_throughput": &schema.Schema{
							Description: `Average throughput of application
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"business_relevance": &schema.Schema{
							Description: `Application's business relevance
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_ip": &schema.Schema{
							Description: `Endpoint client ip
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_mac_address": &schema.Schema{
							Description: `Endpoint mac address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_name": &schema.Schema{
							Description: `Endpoint client name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_network_latency": &schema.Schema{
							Description: `Latency of client network
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_type": &schema.Schema{
							Description: `Type of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exporter_family": &schema.Schema{
							Description: `Devices family of exporter device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exporter_ip_address": &schema.Schema{
							Description: `Ip address of exporter device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exporter_name": &schema.Schema{
							Description: `Name of exporter device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exporter_uui_d": &schema.Schema{
							Description: `UUID of exporter device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"health": &schema.Schema{
							Description: `Health score
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"issue_id": &schema.Schema{
							Description: `Id number of issue
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_name": &schema.Schema{
							Description: `Name of issue
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"jitter": &schema.Schema{
							Description: `Jitter for application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `Site location
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Application name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_latency": &schema.Schema{
							Description: `Network latency for application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"occurrences": &schema.Schema{
							Description: `Issue occurrences
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"operating_system": &schema.Schema{
							Description: `Endpoint's operating system
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"packet_loss_percent": &schema.Schema{
							Description: `Packet loss percentage for application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"priority": &schema.Schema{
							Description: `Issue priority
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"root_cause": &schema.Schema{
							Description: `Issue's root cause
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"server_network_latency": &schema.Schema{
							Description: `Latency of server network
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"severity": &schema.Schema{
							Description: `Issue severity
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"summary": &schema.Schema{
							Description: `Issue summary
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"timestamp": &schema.Schema{
							Description: `Issue's timestamp
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"traffic_class": &schema.Schema{
							Description: `Application's traffic class
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"usage_bytes": &schema.Schema{
							Description: `Usage amount in bytes
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

func dataSourceApplicationsHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vApplicationHealth, okApplicationHealth := d.GetOk("application_health")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vApplicationName, okApplicationName := d.GetOk("application_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Applications")
		queryParams1 := dnacentersdkgo.ApplicationsQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okApplicationHealth {
			queryParams1.ApplicationHealth = vApplicationHealth.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okApplicationName {
			queryParams1.ApplicationName = vApplicationName.(string)
		}

		response1, restyResp1, err := client.Applications.Applications(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Applications", err,
				"Failure at Applications, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationsApplicationsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Applications response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationsApplicationsItems(items *[]dnacentersdkgo.ResponseApplicationsApplicationsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["health"] = item.Health
		respItem["business_relevance"] = item.BusinessRelevance
		respItem["traffic_class"] = item.TrafficClass
		respItem["usage_bytes"] = item.UsageBytes
		respItem["average_throughput"] = item.AverageThroughput
		respItem["packet_loss_percent"] = flattenApplicationsApplicationsItemsPacketLossPercent(item.PacketLossPercent)
		respItem["network_latency"] = flattenApplicationsApplicationsItemsNetworkLatency(item.NetworkLatency)
		respItem["jitter"] = flattenApplicationsApplicationsItemsJitter(item.Jitter)
		respItem["application_server_latency"] = flattenApplicationsApplicationsItemsApplicationServerLatency(item.ApplicationServerLatency)
		respItem["client_network_latency"] = flattenApplicationsApplicationsItemsClientNetworkLatency(item.ClientNetworkLatency)
		respItem["server_network_latency"] = flattenApplicationsApplicationsItemsServerNetworkLatency(item.ServerNetworkLatency)
		respItem["exporter_ip_address"] = item.ExporterIPAddress
		respItem["exporter_name"] = item.ExporterName
		respItem["exporter_uui_d"] = item.ExporterUUID
		respItem["exporter_family"] = item.ExporterFamily
		respItem["client_name"] = item.ClientName
		respItem["client_ip"] = item.ClientIP
		respItem["location"] = item.Location
		respItem["operating_system"] = item.OperatingSystem
		respItem["device_type"] = item.DeviceType
		respItem["client_mac_address"] = item.ClientMacAddress
		respItem["issue_id"] = item.IssueID
		respItem["issue_name"] = item.IssueName
		respItem["application"] = item.Application
		respItem["severity"] = item.Severity
		respItem["summary"] = item.Summary
		respItem["root_cause"] = item.RootCause
		respItem["timestamp"] = item.Timestamp
		respItem["occurrences"] = item.Occurrences
		respItem["priority"] = item.Priority
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationsApplicationsItemsPacketLossPercent(item *dnacentersdkgo.ResponseApplicationsApplicationsResponsePacketLossPercent) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenApplicationsApplicationsItemsNetworkLatency(item *dnacentersdkgo.ResponseApplicationsApplicationsResponseNetworkLatency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenApplicationsApplicationsItemsJitter(item *dnacentersdkgo.ResponseApplicationsApplicationsResponseJitter) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenApplicationsApplicationsItemsApplicationServerLatency(item *dnacentersdkgo.ResponseApplicationsApplicationsResponseApplicationServerLatency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenApplicationsApplicationsItemsClientNetworkLatency(item *dnacentersdkgo.ResponseApplicationsApplicationsResponseClientNetworkLatency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenApplicationsApplicationsItemsServerNetworkLatency(item *dnacentersdkgo.ResponseApplicationsApplicationsResponseServerNetworkLatency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
