package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricSiteHealthSummariesIDTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get health time series for a specific Fabric Site by providing the unique Fabric site id in the url path. The data
will be grouped based on the specified trend time interval. If startTime and endTime are not provided, the API defaults
to the last 24 hours.
By default: the number of records returned will be 500. the records will be sorted in time ascending (*asc*) order
ex: id:93a25378-7740-4e20-8d90-0060ad9a1be0
This data source provides the latest health data until the given *endTime*. If data is not ready for the provided
endTime, the request will fail with error code *400 Bad Request*, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When *endTime* is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
fabricSiteHealthSummaries-1.0.1-resolved.yaml
`,

		ReadContext: dataSourceFabricSiteHealthSummariesIDTrendAnalyticsRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The interested fields in the request. For valid attributes, verify the documentation.
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
				Description: `id path parameter. unique fabric site id
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
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
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
			"trend_interval": &schema.Schema{
				Description: `trendInterval query parameter. The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
`,
				Type:     schema.TypeString,
				Required: true,
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
										Type:        schema.TypeInt,
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

func dataSourceFabricSiteHealthSummariesIDTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vTrendInterval := d.Get("trend_interval")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vOrder, okOrder := d.GetOk("order")
	vAttribute, okAttribute := d.GetOk("attribute")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams{}
		queryParams1 := dnacentersdkgo.TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		queryParams1.TrendInterval = vTrendInterval.(string)

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sda.TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange(vvID, &headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange", err,
				"Failure at TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeItems(items *[]dnacentersdkgo.ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeItemsAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeItemsAttributes(items *[]dnacentersdkgo.ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponseAttributes) []map[string]interface{} {
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
