package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteHealthSummariesIDTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Retrieves the time series information of health and issue data for a site specified by the path parameter. The data
will be grouped based on the specified trend time interval. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-2.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteHealthSummariesIDTrendAnalyticsRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. Supported Analytics Attributes:
[networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]
attribute=networkDeviceCount (single attribute requested)
attribute=networkDeviceCount&attribute=clientCount (multiple attributes requested)
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
			"id": &schema.Schema{
				Description: `id path parameter. unique site uuid
`,
				Type:     schema.TypeString,
				Required: true,
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
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"task_id": &schema.Schema{
				Description: `taskId query parameter. used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_sort_order": &schema.Schema{
				Description: `timeSortOrder query parameter. The sort order of a time sorted API response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"trend_interval": &schema.Schema{
				Description: `trendInterval query parameter. The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
`,
				Type:     schema.TypeString,
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

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
								},
							},
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteHealthSummariesIDTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vTrendInterval, okTrendInterval := d.GetOk("trend_interval")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vTimeSortOrder, okTimeSortOrder := d.GetOk("time_sort_order")
	vAttribute, okAttribute := d.GetOk("attribute")
	vTaskID, okTaskID := d.GetOk("task_id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadTrendAnalyticsDataForASpecificSiteInYourNetwork")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams{}
		queryParams1 := dnacentersdkgo.ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okTrendInterval {
			queryParams1.TrendInterval = vTrendInterval.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okTimeSortOrder {
			queryParams1.TimeSortOrder = vTimeSortOrder.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sites.ReadTrendAnalyticsDataForASpecificSiteInYourNetwork(vvID, &headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadTrendAnalyticsDataForASpecificSiteInYourNetwork", err,
				"Failure at ReadTrendAnalyticsDataForASpecificSiteInYourNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadTrendAnalyticsDataForASpecificSiteInYourNetwork response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkItems(items *[]dnacentersdkgo.ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkItemsAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkItemsAttributes(items *[]dnacentersdkgo.ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponseAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
