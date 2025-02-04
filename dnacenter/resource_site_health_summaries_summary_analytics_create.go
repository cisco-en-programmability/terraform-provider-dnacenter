package dnacenter

import (
	"context"

	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSiteHealthSummariesSummaryAnalyticsCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sites.

- Query an aggregated summary of all site health This data source action provides the latest health data from a given
*endTime* If data is not ready for the provided endTime, the request will fail, and the error message will indicate the
recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime,
since we are not a real time system. When *endTime* is not provided, the API returns the latest data. This data source
action also provides issue data. The *startTime* query param can be used to specify the beginning point of time range to
retrieve the active issue counts in. When this param is not provided, the default *startTime* will be 24 hours before
endTime.

 Aggregated response data will NOT have unique identifier data populated.

 List of unique identifier data: [*id*, *siteHierarchy*,
*siteHierarchyId*, *siteType*, *latitude*, *longitude*] Please refer to the 'API Support Documentation' section to
understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml
`,

		CreateContext: resourceSiteHealthSummariesSummaryAnalyticsCreateCreate,
		ReadContext:   resourceSiteHealthSummariesSummaryAnalyticsCreateRead,
		DeleteContext: resourceSiteHealthSummariesSummaryAnalyticsCreateDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_device_count": &schema.Schema{
							Description: `Access Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"access_device_good_health_count": &schema.Schema{
							Description: `Access Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"access_device_good_health_percentage": &schema.Schema{
							Description: `Access Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"ap_device_count": &schema.Schema{
							Description: `Ap Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"ap_device_good_health_count": &schema.Schema{
							Description: `Ap Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"ap_device_good_health_percentage": &schema.Schema{
							Description: `Ap Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"client_count": &schema.Schema{
							Description: `Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"client_data_usage": &schema.Schema{
							Description: `Client Data Usage`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"client_good_health_count": &schema.Schema{
							Description: `Client Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"client_good_health_percentage": &schema.Schema{
							Description: `Client Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"core_device_count": &schema.Schema{
							Description: `Core Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"core_device_good_health_count": &schema.Schema{
							Description: `Core Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"core_device_good_health_percentage": &schema.Schema{
							Description: `Core Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"distribution_device_count": &schema.Schema{
							Description: `Distribution Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"distribution_device_good_health_count": &schema.Schema{
							Description: `Distribution Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"distribution_device_good_health_percentage": &schema.Schema{
							Description: `Distribution Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"issue_count": &schema.Schema{
							Description: `Issue Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"latitude": &schema.Schema{
							Description: `Latitude`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"longitude": &schema.Schema{
							Description: `Longitude`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"network_device_count": &schema.Schema{
							Description: `Network Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"network_device_good_health_count": &schema.Schema{
							Description: `Network Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"network_device_good_health_percentage": &schema.Schema{
							Description: `Network Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"p1_issue_count": &schema.Schema{
							Description: `P1 Issue Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"p2_issue_count": &schema.Schema{
							Description: `P2 Issue Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"p3_issue_count": &schema.Schema{
							Description: `P3 Issue Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"p4_issue_count": &schema.Schema{
							Description: `P4 Issue Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"router_device_count": &schema.Schema{
							Description: `Router Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"router_device_good_health_count": &schema.Schema{
							Description: `Router Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"router_device_good_health_percentage": &schema.Schema{
							Description: `Router Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_type": &schema.Schema{
							Description: `Site Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"switch_device_count": &schema.Schema{
							Description: `Switch Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"switch_device_good_health_count": &schema.Schema{
							Description: `Switch Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"switch_device_good_health_percentage": &schema.Schema{
							Description: `Switch Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wired_client_count": &schema.Schema{
							Description: `Wired Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wired_client_good_health_count": &schema.Schema{
							Description: `Wired Client Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wired_client_good_health_percentage": &schema.Schema{
							Description: `Wired Client Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_client_count": &schema.Schema{
							Description: `Wireless Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_client_good_health_count": &schema.Schema{
							Description: `Wireless Client Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_client_good_health_percentage": &schema.Schema{
							Description: `Wireless Client Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_device_count": &schema.Schema{
							Description: `Wireless Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_device_good_health_count": &schema.Schema{
							Description: `Wireless Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wireless_device_good_health_percentage": &schema.Schema{
							Description: `Wireless Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wlc_device_count": &schema.Schema{
							Description: `Wlc Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wlc_device_good_health_count": &schema.Schema{
							Description: `Wlc Device Good Health Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"wlc_device_good_health_percentage": &schema.Schema{
							Description: `Wlc Device Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (***) character search support. E.g. **/San*, */San, /San**
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"site_hierarchy_id": &schema.Schema{
							Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. **uuid*, *uuid, uuid**
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
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
							ForceNew: true,
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"views": &schema.Schema{
							Description: `Views`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceSiteHealthSummariesSummaryAnalyticsCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSiteHealthSummariesSummaryAnalyticsCreateQueryAnAggregatedSummaryOfSiteHealthData(ctx, "parameters.0", d)
	queryParams1 := dnacentersdkgo.QueryAnAggregatedSummaryOfSiteHealthDataQueryParams{}

	// has_unknown_response: None

	response1, restyResp1, err := client.Sites.QueryAnAggregatedSummaryOfSiteHealthData(request1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing QueryAnAggregatedSummaryOfSiteHealthData", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenSitesQueryAnAggregatedSummaryOfSiteHealthDataItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting QueryAnAggregatedSummaryOfSiteHealthData response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSiteHealthSummariesSummaryAnalyticsCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSiteHealthSummariesSummaryAnalyticsCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSiteHealthSummariesSummaryAnalyticsCreateQueryAnAggregatedSummaryOfSiteHealthData(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesQueryAnAggregatedSummaryOfSiteHealthData {
	request := dnacentersdkgo.RequestSitesQueryAnAggregatedSummaryOfSiteHealthData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	return &request
}

func flattenSitesQueryAnAggregatedSummaryOfSiteHealthDataItem(item *dnacentersdkgo.ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["site_hierarchy_id"] = item.SiteHierarchyID
	respItem["site_type"] = item.SiteType
	respItem["latitude"] = item.Latitude
	respItem["longitude"] = item.Longitude
	respItem["network_device_good_health_percentage"] = item.NetworkDeviceGoodHealthPercentage
	respItem["network_device_good_health_count"] = item.NetworkDeviceGoodHealthCount
	respItem["client_good_health_count"] = item.ClientGoodHealthCount
	respItem["client_good_health_percentage"] = item.ClientGoodHealthPercentage
	respItem["wired_client_good_health_percentage"] = item.WiredClientGoodHealthPercentage
	respItem["wireless_client_good_health_percentage"] = item.WirelessClientGoodHealthPercentage
	respItem["client_count"] = item.ClientCount
	respItem["wired_client_count"] = item.WiredClientCount
	respItem["wireless_client_count"] = item.WirelessClientCount
	respItem["wired_client_good_health_count"] = item.WiredClientGoodHealthCount
	respItem["wireless_client_good_health_count"] = item.WirelessClientGoodHealthCount
	respItem["network_device_count"] = item.NetworkDeviceCount
	respItem["access_device_count"] = item.AccessDeviceCount
	respItem["access_device_good_health_count"] = item.AccessDeviceGoodHealthCount
	respItem["core_device_count"] = item.CoreDeviceCount
	respItem["core_device_good_health_count"] = item.CoreDeviceGoodHealthCount
	respItem["distribution_device_count"] = item.DistributionDeviceCount
	respItem["distribution_device_good_health_count"] = item.DistributionDeviceGoodHealthCount
	respItem["router_device_count"] = item.RouterDeviceCount
	respItem["router_device_good_health_count"] = item.RouterDeviceGoodHealthCount
	respItem["wireless_device_count"] = item.WirelessDeviceCount
	respItem["wireless_device_good_health_count"] = item.WirelessDeviceGoodHealthCount
	respItem["ap_device_count"] = item.ApDeviceCount
	respItem["ap_device_good_health_count"] = item.ApDeviceGoodHealthCount
	respItem["wlc_device_count"] = item.WlcDeviceCount
	respItem["wlc_device_good_health_count"] = item.WlcDeviceGoodHealthCount
	respItem["switch_device_count"] = item.SwitchDeviceCount
	respItem["switch_device_good_health_count"] = item.SwitchDeviceGoodHealthCount
	respItem["access_device_good_health_percentage"] = item.AccessDeviceGoodHealthPercentage
	respItem["core_device_good_health_percentage"] = item.CoreDeviceGoodHealthPercentage
	respItem["distribution_device_good_health_percentage"] = item.DistributionDeviceGoodHealthPercentage
	respItem["router_device_good_health_percentage"] = item.RouterDeviceGoodHealthPercentage
	respItem["ap_device_good_health_percentage"] = item.ApDeviceGoodHealthPercentage
	respItem["wlc_device_good_health_percentage"] = item.WlcDeviceGoodHealthPercentage
	respItem["switch_device_good_health_percentage"] = item.SwitchDeviceGoodHealthPercentage
	respItem["wireless_device_good_health_percentage"] = item.WirelessDeviceGoodHealthPercentage
	respItem["client_data_usage"] = item.ClientDataUsage
	respItem["p1_issue_count"] = item.P1IssueCount
	respItem["p2_issue_count"] = item.P2IssueCount
	respItem["p3_issue_count"] = item.P3IssueCount
	respItem["p4_issue_count"] = item.P4IssueCount
	respItem["issue_count"] = item.IssueCount
	return []map[string]interface{}{
		respItem,
	}
}
