package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteHealthSummariesTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Retrieves the time series information of health and issue data for sites specified by query parameters, or all sites.
The data will be grouped based on the specified trend time interval. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-2.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteHealthSummariesTrendAnalyticsRead,
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
				Description: `id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
`,
				Type:     schema.TypeString,
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
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (***) character search support. E.g. */San*, */San*, */San*
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. *uuid*, *uuid*, *uuid*
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_type": &schema.Schema{
				Description: `siteType query parameter. The type of the site. A site can be an area, building, or floor.
Default when not provided will be *[floor,building,area]*
Examples:
*?siteType=area* (single siteType requested)
*?siteType=area&siteType=building&siteType=floor* (multiple siteTypes requested)
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

func dataSourceSiteHealthSummariesTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteType, okSiteType := d.GetOk("site_type")
	vID, okID := d.GetOk("id")
	vTrendInterval, okTrendInterval := d.GetOk("trend_interval")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vTimeSortOrder, okTimeSortOrder := d.GetOk("time_sort_order")
	vAttribute, okAttribute := d.GetOk("attribute")
	vTaskID, okTaskID := d.GetOk("task_id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork")

		headerParams1 := dnacentersdkgo.ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams{}
		queryParams1 := dnacentersdkgo.ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okSiteHierarchy {
			queryParams1.SiteHierarchy = vSiteHierarchy.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okSiteType {
			queryParams1.SiteType = vSiteType.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
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

		response1, restyResp1, err := client.Sites.ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork", err,
				"Failure at ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkItems(items *[]dnacentersdkgo.ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkItemsAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkItemsAttributes(items *[]dnacentersdkgo.ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponseAttributes) []map[string]interface{} {
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
