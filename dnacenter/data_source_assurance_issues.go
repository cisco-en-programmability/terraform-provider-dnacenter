package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceIssues() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Returns all details of each issue along with suggested actions for given set of filters specified in query parameters.
If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted
to 24-hours ago from end time. All string type query parameters support wildcard search (using *). For example:
siteHierarchy=Global/San Jose/* returns issues under all sites whole siteHierarchy starts with "Global/San Jose/".
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
IssuesList-1.0.0-resolved.yaml

- Returns all the details and suggested actions of an issue for the given issue id. https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceIssuesRead,
		Schema: map[string]*schema.Schema{
			"accept_language": &schema.Schema{
				Description: `Accept-Language header parameter. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ai_driven": &schema.Schema{
				Description: `aiDriven query parameter. Flag whether the issue is AI driven issue
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"attribute": &schema.Schema{
				Description: `attribute query parameter. List of attributes related to the issue. If these are provided, then only those attributes will be part of response along with the default attributes. Please refer to the *IssuesResponseAttribute* Model for supported attributes. Examples: *attribute=deviceType* (single attribute requested) *attribute=deviceType&attribute=updatedBy* (multiple attributes requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Description: `category query parameter. Categories of the issue. Supports single category and multiple categories. Examples: category=availability (single status requested) category=availability&category=onboarding (multiple categories requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. Device Type of the device to which this issue belongs to. Supports single device type and multiple device types. Examples: deviceType=wireless controller (single device type requested) deviceType=wireless controller&deviceType=core (multiple device types requested)
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
			"entity_id": &schema.Schema{
				Description: `entityId query parameter. Id of the entity for which this issue belongs to. For example, it
    could be mac address of AP or UUID of Sensor
  example: 68:ca:e4:79:3f:20 4de02167-901b-43cf-8822-cffd3caa286f
Examples: entityId=68:ca:e4:79:3f:20 (single entity id requested) entityId=68:ca:e4:79:3f:20&entityId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple entity ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_type": &schema.Schema{
				Description: `entityType query parameter. Entity type of the issue. Supports single entity type and multiple entity types. Examples: entityType=networkDevice (single entity type requested) entityType=network device&entityType=client (multiple entity types requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_driven": &schema.Schema{
				Description: `fabricDriven query parameter. Flag whether the issue is related to a Fabric site, a virtual network or a transit.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fabric_site_driven": &schema.Schema{
				Description: `fabricSiteDriven query parameter. Flag whether the issue is Fabric site driven issue
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fabric_site_id": &schema.Schema{
				Description: `fabricSiteId query parameter. The UUID of the fabric site. (Ex. "flooruuid") Examples: fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26,864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_transit_driven": &schema.Schema{
				Description: `fabricTransitDriven query parameter. Flag whether the issue is Fabric Transit driven issue
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fabric_transit_site_id": &schema.Schema{
				Description: `fabricTransitSiteId query parameter. The UUID of the fabric transit site. (Ex. "flooruuid") Examples: fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26&fabricTransitSiteId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_vn_driven": &schema.Schema{
				Description: `fabricVnDriven query parameter. Flag whether the issue is Fabric Virtual Network driven issue
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fabric_vn_name": &schema.Schema{
				Description: `fabricVnName query parameter. The name of the fabric virtual network Examples: fabricVnName=name1 (single fabric virtual network name requested) fabricVnName=name1&fabricVnName=name2&fabricVnName=name3 (multiple fabric virtual network names requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. The issue Uuid
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_global": &schema.Schema{
				Description: `isGlobal query parameter. Global issues are those issues which impacts across many devices, sites. They are also displayed on Issue Dashboard in Catalyst Center UI. Non-Global issues are displayed only on Client 360 or Device 360 pages. If this flag is 'true', only global issues are returned. If it if 'false', all issues are returned.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"issue_id": &schema.Schema{
				Description: `issueId query parameter. UUID of the issue Examples: issueId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single issue id requested) issueId=e52aecfe-b142-4287-a587-11a16ba6dd26&issueId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple issue ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of issues to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. The macAddress of the network device or client This field supports wildcard (***) character-based search.  Ex: **AB:AB:AB** or *AB:AB:AB** or **AB:AB:AB* Examples:
*macAddress=AB:AB:AB:CD:CD:CD* (single macAddress requested)
*macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE* (multiple macAddress requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The name of the issue Examples: name=ap_down (single issue name requested) name=ap_down&name=wlc_monitor (multiple issue names requested) Issue names can be retrieved using the API /data/api/v1/assuranceIssueConfigurations
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. The list of Network Device Uuids. (Ex. *6bef213c-19ca-4170-8375-b694e251101c*)
Examples:
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c* (single networkDeviceId requested)
*networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0* (multiple networkDeviceIds with & separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_ip_address": &schema.Schema{
				Description: `networkDeviceIpAddress query parameter. The list of Network Device management IP Address. (Ex. *121.1.1.10*)
This field supports wildcard (***) character-based search.  Ex: **1.1** or *1.1** or **1.1*
Examples:
*networkDeviceIpAddress=121.1.1.10*
*networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10* (multiple networkDevice IP Address with & separator)
`,
				Type:     schema.TypeString,
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
			"priority": &schema.Schema{
				Description: `priority query parameter. Priority of the issue. Supports single priority and multiple priorities Examples: priority=P1 (single priority requested) priority=P1&priority=P2&priority=P3 (multiple priorities requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. Severity of the issue. Supports single severity and multiple severities. Examples: severity=high (single severity requested) severity=high&severity=medium (multiple severities requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy": &schema.Schema{
				Description: `siteHierarchy query parameter. The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. *Global/AreaName/BuildingName/FloorName*)
This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San*
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (*) character search support. E.g. **uuid*, *uuid, uuid*
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid*
Examples:
*?siteId=id1* (single id requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple ids requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_name": &schema.Schema{
				Description: `siteName query parameter. The name of the site. (Ex. *FloorName*)
This field supports wildcard asterisk (*) character search support. E.g. *San*, *San, San*
Examples:
*?siteName=building1* (single siteName requested)
*?siteName=building1&siteName=building2&siteName=building3* (multiple siteNames requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
If *startTime* is not provided, API will default to current time.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Status of the issue. Supports single status and multiple statuses. Examples: status=active (single status requested) status=active&status=resolved (multiple statuses requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_by": &schema.Schema{
				Description: `updatedBy query parameter. The user who last updated this issue. Examples: updatedBy=admin (single updatedBy requested) updatedBy=admin&updatedBy=john (multiple updatedBy requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The name of the View. Each view represents a specific data set. Please refer to the *IssuesView* Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned.
| View Name | Included Attributes | | --| --| | *update* | updatedTime, updatedBy | | *site* | siteName, siteHierarchy, siteId, siteHierarchyId | Examples: *view=update* (single view requested) *view=update&view=site* (multiple views requested)       
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"entity_id": &schema.Schema{
							Description: `Entity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"entity_type": &schema.Schema{
							Description: `Entity Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"first_occurred_time": &schema.Schema{
							Description: `First Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"is_global": &schema.Schema{
							Description: `Is Global`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_id": &schema.Schema{
							Description: `Issue Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"most_recent_occurred_time": &schema.Schema{
							Description: `Most Recent Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"notes": &schema.Schema{
							Description: `Notes`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"suggested_actions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Description: `Message`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"steps": &schema.Schema{
										Description: `Steps`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"summary": &schema.Schema{
							Description: `Summary`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"updated_by": &schema.Schema{
							Description: `Updated By`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"updated_time": &schema.Schema{
							Description: `Updated Time`,
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

						"additional_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"entity_id": &schema.Schema{
							Description: `Entity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"entity_type": &schema.Schema{
							Description: `Entity Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"first_occurred_time": &schema.Schema{
							Description: `First Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"is_global": &schema.Schema{
							Description: `Is Global`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_id": &schema.Schema{
							Description: `Issue Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"most_recent_occurred_time": &schema.Schema{
							Description: `Most Recent Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"notes": &schema.Schema{
							Description: `Notes`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"suggested_actions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Description: `Message`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"steps": &schema.Schema{
										Description: `Steps`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"summary": &schema.Schema{
							Description: `Summary`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"updated_by": &schema.Schema{
							Description: `Updated By`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"updated_time": &schema.Schema{
							Description: `Updated Time`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAssuranceIssuesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vIsGlobal, okIsGlobal := d.GetOk("is_global")
	vPriority, okPriority := d.GetOk("priority")
	vSeverity, okSeverity := d.GetOk("severity")
	vStatus, okStatus := d.GetOk("status")
	vEntityType, okEntityType := d.GetOk("entity_type")
	vCategory, okCategory := d.GetOk("category")
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vName, okName := d.GetOk("name")
	vIssueID, okIssueID := d.GetOk("issue_id")
	vEntityID, okEntityID := d.GetOk("entity_id")
	vUpdatedBy, okUpdatedBy := d.GetOk("updated_by")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteName, okSiteName := d.GetOk("site_name")
	vSiteID, okSiteID := d.GetOk("site_id")
	vFabricSiteID, okFabricSiteID := d.GetOk("fabric_site_id")
	vFabricVnName, okFabricVnName := d.GetOk("fabric_vn_name")
	vFabricTransitSiteID, okFabricTransitSiteID := d.GetOk("fabric_transit_site_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vNetworkDeviceIPAddress, okNetworkDeviceIPAddress := d.GetOk("network_device_ip_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vAiDriven, okAiDriven := d.GetOk("ai_driven")
	vFabricDriven, okFabricDriven := d.GetOk("fabric_driven")
	vFabricSiteDriven, okFabricSiteDriven := d.GetOk("fabric_site_driven")
	vFabricVnDriven, okFabricVnDriven := d.GetOk("fabric_vn_driven")
	vFabricTransitDriven, okFabricTransitDriven := d.GetOk("fabric_transit_driven")
	vAcceptLanguage, okAcceptLanguage := d.GetOk("accept_language")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")
	vID, okID := d.GetOk("id")

	method1 := []bool{okStartTime, okEndTime, okLimit, okOffset, okSortBy, okOrder, okIsGlobal, okPriority, okSeverity, okStatus, okEntityType, okCategory, okDeviceType, okName, okIssueID, okEntityID, okUpdatedBy, okSiteHierarchy, okSiteHierarchyID, okSiteName, okSiteID, okFabricSiteID, okFabricVnName, okFabricTransitSiteID, okNetworkDeviceID, okNetworkDeviceIPAddress, okMacAddress, okView, okAttribute, okAiDriven, okFabricDriven, okFabricSiteDriven, okFabricVnDriven, okFabricTransitDriven, okAcceptLanguage, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okView, okAttribute, okAcceptLanguage, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork")

		headerParams1 := dnacentersdkgo.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams{}
		queryParams1 := dnacentersdkgo.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams{}

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
		if okIsGlobal {
			queryParams1.IsGlobal = vIsGlobal.(bool)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okEntityType {
			queryParams1.EntityType = vEntityType.(string)
		}
		if okCategory {
			queryParams1.Category = vCategory.(string)
		}
		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okIssueID {
			queryParams1.IssueID = vIssueID.(string)
		}
		if okEntityID {
			queryParams1.EntityID = vEntityID.(string)
		}
		if okUpdatedBy {
			queryParams1.UpdatedBy = vUpdatedBy.(string)
		}
		if okSiteHierarchy {
			queryParams1.SiteHierarchy = vSiteHierarchy.(string)
		}
		if okSiteHierarchyID {
			queryParams1.SiteHierarchyID = vSiteHierarchyID.(string)
		}
		if okSiteName {
			queryParams1.SiteName = vSiteName.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okFabricSiteID {
			queryParams1.FabricSiteID = vFabricSiteID.(string)
		}
		if okFabricVnName {
			queryParams1.FabricVnName = vFabricVnName.(string)
		}
		if okFabricTransitSiteID {
			queryParams1.FabricTransitSiteID = vFabricTransitSiteID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okNetworkDeviceIPAddress {
			queryParams1.NetworkDeviceIPAddress = vNetworkDeviceIPAddress.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okAiDriven {
			queryParams1.AiDriven = vAiDriven.(bool)
		}
		if okFabricDriven {
			queryParams1.FabricDriven = vFabricDriven.(bool)
		}
		if okFabricSiteDriven {
			queryParams1.FabricSiteDriven = vFabricSiteDriven.(bool)
		}
		if okFabricVnDriven {
			queryParams1.FabricVnDriven = vFabricVnDriven.(bool)
		}
		if okFabricTransitDriven {
			queryParams1.FabricTransitDriven = vFabricTransitDriven.(bool)
		}
		if okAcceptLanguage {
			headerParams1.AcceptLanguage = vAcceptLanguage.(string)
		}
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Issues.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork", err,
				"Failure at GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDHeaderParams{}
		queryParams2 := dnacentersdkgo.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDQueryParams{}

		if okView {
			queryParams2.View = vView.(string)
		}
		if okAttribute {
			queryParams2.Attribute = vAttribute.(string)
		}
		if okAcceptLanguage {
			headerParams2.AcceptLanguage = vAcceptLanguage.(string)
		}
		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Issues.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID(vvID, &headerParams2, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID", err,
				"Failure at GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItems(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["summary"] = item.Summary
		respItem["priority"] = item.Priority
		respItem["severity"] = item.Severity
		respItem["device_type"] = item.DeviceType
		respItem["category"] = item.Category
		respItem["entity_type"] = item.EntityType
		respItem["entity_id"] = item.EntityID
		respItem["first_occurred_time"] = item.FirstOccurredTime
		respItem["most_recent_occurred_time"] = item.MostRecentOccurredTime
		respItem["status"] = item.Status
		respItem["is_global"] = boolPtrToString(item.IsGlobal)
		respItem["updated_by"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsUpdatedBy(item.UpdatedBy)
		respItem["updated_time"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsUpdatedTime(item.UpdatedTime)
		respItem["notes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsNotes(item.Notes)
		respItem["site_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteID(item.SiteID)
		respItem["site_hierarchy_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteHierarchyID(item.SiteHierarchyID)
		respItem["site_name"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteName(item.SiteName)
		respItem["site_hierarchy"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteHierarchy(item.SiteHierarchy)
		respItem["suggested_actions"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSuggestedActions(item.SuggestedActions)
		respItem["additional_attributes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsAdditionalAttributes(item.AdditionalAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsUpdatedBy(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedBy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsUpdatedTime(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsNotes(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseNotes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteID(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteHierarchyID(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchyID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteName(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSiteHierarchy(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSuggestedActions(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsSuggestedActionsSteps(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActionsSteps) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkItemsAdditionalAttributes(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseAdditionalAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItem(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["issue_id"] = item.IssueID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["summary"] = item.Summary
	respItem["priority"] = item.Priority
	respItem["severity"] = item.Severity
	respItem["device_type"] = item.DeviceType
	respItem["category"] = item.Category
	respItem["entity_type"] = item.EntityType
	respItem["entity_id"] = item.EntityID
	respItem["first_occurred_time"] = item.FirstOccurredTime
	respItem["most_recent_occurred_time"] = item.MostRecentOccurredTime
	respItem["status"] = item.Status
	respItem["is_global"] = boolPtrToString(item.IsGlobal)
	respItem["updated_by"] = item.UpdatedBy
	respItem["updated_time"] = item.UpdatedTime
	respItem["notes"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemNotes(item.Notes)
	respItem["site_id"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteID(item.SiteID)
	respItem["site_hierarchy_id"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteHierarchyID(item.SiteHierarchyID)
	respItem["site_name"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteName(item.SiteName)
	respItem["site_hierarchy"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteHierarchy(item.SiteHierarchy)
	respItem["suggested_actions"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSuggestedActions(item.SuggestedActions)
	respItem["additional_attributes"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemAdditionalAttributes(item.AdditionalAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemNotes(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseNotes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteID(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteHierarchyID(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchyID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteName(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSiteHierarchy(item *dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSuggestedActions(items *[]dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemSuggestedActionsSteps(items *[]dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActionsSteps) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDItemAdditionalAttributes(items *[]dnacentersdkgo.ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseAdditionalAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
