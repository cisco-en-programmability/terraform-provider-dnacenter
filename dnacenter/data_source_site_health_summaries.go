package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteHealthSummaries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Get a paginated list of site health summaries. Use the available query parameters to identify a subset of sites you
want health summaries for. This data source provides the latest health data from a given *endTime* If data is not ready
for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to
retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time
system. When *endTime* is not provided, the API returns the latest data. This data source also provides issue data. The
*startTime* query param can be used to specify the beginning point of time range to retrieve the active issue counts in.
When this param is not provided, the default *startTime* will be 24 hours before endTime. Valid values for *sortBy*
param in this API are limited to the attributes provided in the *site* view. Default sortBy is 'siteHierarchy' in order
'asc' (ascending). For detailed information about the usage of the API, please refer to the Open API specification
document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
siteHealthSummaries-1.0.3-resolved.yaml

- Get a health summary for a specific site by providing the unique site id in the url path. This data source provides
the latest health data from a given *endTime* If data is not ready for the provided endTime, the request will fail, and
the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur
if the provided endTime=currentTime, since we are not a real time system. When *endTime* is not provided, the API
returns the latest data. This data source also provides issue data. The *startTime* query param can be used to specify
the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default
*startTime* will be 24 hours before endTime. For detailed information about the usage of the API, please refer to the
Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml
`,

		ReadContext: dataSourceSiteHealthSummariesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. Supported Attributes:
[id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]
If length of attribute list is too long, please use 'view' param instead.
Examples:
attribute=siteHierarchy (single attribute requested)
attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
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
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (***) character search support. E.g. **/San*, */San, /San**
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
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
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *startTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites.
### Response data proviced by each view:
1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude]
2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage]
3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage]
4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount]
When this query parameter is not added the default summaries are:
**[site,client,network,issue]**
Examples:
view=client (single view requested)
view=client&view=network&view=issue (multiple views requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Optional: true,
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

			"items": &schema.Schema{
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
		},
	}
}

func dataSourceSiteHealthSummariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteType, okSiteType := d.GetOk("site_type")
	vID, okID := d.GetOk("id")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")

	method1 := []bool{okStartTime, okEndTime, okLimit, okOffset, okSortBy, okOrder, okSiteHierarchy, okSiteHierarchyID, okSiteType, okID, okView, okAttribute, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okStartTime, okEndTime, okView, okAttribute, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadListOfSiteHealthSummaries")

		headerParams1 := dnacentersdkgo.ReadListOfSiteHealthSummariesHeaderParams{}
		queryParams1 := dnacentersdkgo.ReadListOfSiteHealthSummariesQueryParams{}

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
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Sites.ReadListOfSiteHealthSummaries(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadListOfSiteHealthSummaries", err,
				"Failure at ReadListOfSiteHealthSummaries, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesReadListOfSiteHealthSummariesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadListOfSiteHealthSummaries response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: ReadSiteHealthSummaryDataBySiteID")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.ReadSiteHealthSummaryDataBySiteIDHeaderParams{}
		queryParams2 := dnacentersdkgo.ReadSiteHealthSummaryDataBySiteIDQueryParams{}

		if okStartTime {
			queryParams2.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams2.EndTime = vEndTime.(float64)
		}
		if okView {
			queryParams2.View = vView.(string)
		}
		if okAttribute {
			queryParams2.Attribute = vAttribute.(string)
		}
		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Sites.ReadSiteHealthSummaryDataBySiteID(vvID, &headerParams2, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadSiteHealthSummaryDataBySiteID", err,
				"Failure at ReadSiteHealthSummaryDataBySiteID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSitesReadSiteHealthSummaryDataBySiteIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadSiteHealthSummaryDataBySiteID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesReadListOfSiteHealthSummariesItems(items *[]dnacentersdkgo.ResponseSitesReadListOfSiteHealthSummariesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
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
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesReadSiteHealthSummaryDataBySiteIDItem(item *dnacentersdkgo.ResponseSitesReadSiteHealthSummaryDataBySiteIDResponse) []map[string]interface{} {
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
