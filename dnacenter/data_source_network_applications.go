package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkApplications() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Applications.

- Retrieves the list of network applications along with experience and health metrics. If startTime and endTime are not
provided, the API defaults to the last 24 hours. *siteId* is mandatory. *siteId* must be a site UUID of a building. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
NetworkApplications-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceNetworkApplicationsRead,
		Schema: map[string]*schema.Schema{
			"application_name": &schema.Schema{
				Description: `applicationName query parameter. Name of the application for which the experience data is intended.
Examples:
*applicationName=webex* (single applicationName requested)
*applicationName=webex&applicationName=teams* (multiple applicationName requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute": &schema.Schema{
				Description: `attribute query parameter. List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Supported attributes are applicationName, siteId, exporterIpAddress, exporterNetworkDeviceId, healthScore, businessRelevance, usage, throughput, packetLossPercent, networkLatency, applicationServerLatency, clientNetworkLatency, serverNetworkLatency, trafficClass, jitter, ssid Examples: *attribute=healthScore* (single attribute requested) *attribute=healthScore&attribute=ssid&attribute=jitter* (multiple attribute requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"business_relevance": &schema.Schema{
				Description: `businessRelevance query parameter. The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
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
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The site UUID without the top level hierarchy. *siteId* is mandatory. *siteId* must be a site UUID of a building. (Ex."buildingUuid") Examples: *siteId=buildingUuid* (single siteId requested) *siteId=buildingUuid1&siteId=buildingUuid2* (multiple siteId requested)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. In the context of a network application, SSID refers to the name of the wireless network to which the client connects.
Examples:
*ssid=Alpha* (single ssid requested)
*ssid=Alpha&ssid=Guest* (multiple ssid requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_name": &schema.Schema{
							Description: `Application Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"application_server_latency": &schema.Schema{
							Description: `Application Server Latency`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"business_relevance": &schema.Schema{
							Description: `Business Relevance`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_network_latency": &schema.Schema{
							Description: `Client Network Latency`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"exporter_ip_address": &schema.Schema{
							Description: `Exporter Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"exporter_network_device_id": &schema.Schema{
							Description: `Exporter Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"health_score": &schema.Schema{
							Description: `Health Score`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"jitter": &schema.Schema{
							Description: `Jitter`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"network_latency": &schema.Schema{
							Description: `Network Latency`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"packet_loss_percent": &schema.Schema{
							Description: `Packet Loss Percent`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"server_network_latency": &schema.Schema{
							Description: `Server Network Latency`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ssid": &schema.Schema{
							Description: `Ssid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"throughput": &schema.Schema{
							Description: `Throughput`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"traffic_class": &schema.Schema{
							Description: `Traffic Class`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"usage": &schema.Schema{
							Description: `Usage`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSiteID := d.Get("site_id")
	vSSID, okSSID := d.GetOk("ssid")
	vApplicationName, okApplicationName := d.GetOk("application_name")
	vBusinessRelevance, okBusinessRelevance := d.GetOk("business_relevance")
	vAttribute, okAttribute := d.GetOk("attribute")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics")

		headerParams1 := dnacentersdkgo.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams{}

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
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		queryParams1.SiteID = vSiteID.(string)

		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okApplicationName {
			queryParams1.ApplicationName = vApplicationName.(string)
		}
		if okBusinessRelevance {
			queryParams1.BusinessRelevance = vBusinessRelevance.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Applications.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics", err,
				"Failure at RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsItems(items *[]dnacentersdkgo.ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["application_name"] = item.ApplicationName
		respItem["business_relevance"] = item.BusinessRelevance
		respItem["site_id"] = item.SiteID
		respItem["exporter_ip_address"] = item.ExporterIPAddress
		respItem["exporter_network_device_id"] = item.ExporterNetworkDeviceID
		respItem["health_score"] = item.HealthScore
		respItem["usage"] = item.Usage
		respItem["throughput"] = item.Throughput
		respItem["packet_loss_percent"] = item.PacketLossPercent
		respItem["network_latency"] = item.NetworkLatency
		respItem["application_server_latency"] = item.ApplicationServerLatency
		respItem["client_network_latency"] = item.ClientNetworkLatency
		respItem["server_network_latency"] = item.ServerNetworkLatency
		respItem["traffic_class"] = item.TrafficClass
		respItem["jitter"] = item.Jitter
		respItem["ssid"] = item.SSID
		respItems = append(respItems, respItem)
	}
	return respItems
}
