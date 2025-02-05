package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteKpiSummaries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Returns site analytics for all child sites of given parent site. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteKpiSummariesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. List of attributes related to site analytics. If these are provided, then only those attributes will be part of response along with the default attributes. Examples: *attribute=coverageAverage* (single attribute requested) *attribute=coverageFailureMetrics&attribute=coverageTotalCount* (multiple attributes requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"band": &schema.Schema{
				Description: `band query parameter. WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz GHz Examples: *band=5* (single band requested) *band=2.4&band=6* (multiple band requested)
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
			"failure_category": &schema.Schema{
				Description: `failureCategory query parameter. Category of failure when a client fails to meet the threshold. Examples: *failureCategory=AUTH* (single failure category requested) *failureCategory=AUTH&failureCategory=DHCP* (multiple failure categories requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"failure_reason": &schema.Schema{
				Description: `failureReason query parameter. Reason for failure when a client fails to meet the threshold. Examples: *failureReason=MOBILITY_FAILURE* (single ssid requested) *failureReason=REASON_IPLEARN_CONNECT_TIMEOUT&failureReason=ST_EAP_TIMEOUT*   (multiple ssid requested)
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
This field supports wildcard asterisk (***) character search support. E.g. */San*, */San, /San*
Examples:
*?siteHierarchy=Global/AreaName/BuildingName/FloorName* (single siteHierarchy requested)
*?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2* (multiple siteHierarchies requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_hierarchy_id": &schema.Schema{
				Description: `siteHierarchyId query parameter. The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. *globalUuid/areaUuid/buildingUuid/floorUuid*)
This field supports wildcard asterisk (***) character search support. E.g. *uuid*, *uuid, uuid*
Examples:
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid *(single siteHierarchyId requested)
*?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2* (multiple siteHierarchyIds requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The UUID of the site. (Ex. *flooruuid*)
Examples:
*?siteId=id1* (single id requested)
*?siteId=id1&siteId=id2&siteId=id3* (multiple ids requested)
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
				Description: `sortBy query parameter. Field name on which sorting needs to be done
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID Wireless Local Area Network Identifier. Examples: *ssid=Alpha* (single ssid requested) *ssid=Alpha&ssid=Guest* (multiple ssid requested)
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
			"view": &schema.Schema{
				Description: `view query parameter.
The name of the View. Each view represents a specific data set. Please refer to the
SiteAnalyticsView
 Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned.
View Name
Included Attributes
coverage
coverageAverage, coverageSuccessPercentage, coverageSuccessCount, coverageTotalCount, coverageFailureCount, coverageClientCount, coverageImpactedEntities, coverageFailureImpactedEntities, coverageFailureMetrics
onboardingAttempts
onboardingAttemptsSuccessPercentage, onboardingAttemptsSuccessCount, onboardingAttemptsTotalCount, onboardingAttemptsFailureCount, onboardingAttemptsClientCount, onboardingAttemptsImpactedEntities, onboardingAttemptsFailureImpactedEntities, onboardingAttemptsFailureMetrics
onboardingDuration
onboardingDurationAverage, onboardingDurationSuccessPercentage, onboardingDurationSuccessCount, onboardingDurationTotalCount, onboardingDurationFailureCount, onboardingDurationClientCount, onboardingDurationImpactedEntities, onboardingDurationFailureImpactedEntities, onboardingDurationFailureMetrics
roamingAttempts
roamingAttemptsSuccessPercentage, roamingAttemptsSuccessCount, roamingAttemptsTotalCount, roamingAttemptsFailureCount, roamingAttemptsClientCount, roamingAttemptsImpactedEntities, roamingAttemptsFailureImpactedEntities, roamingAttemptsFailureMetrics
roamingDuration
roamingDurationAverage, roamingDurationSuccessPercentage, roamingDurationSuccessCount, roamingDurationTotalCount, roamingDurationFailureCount, roamingDurationClientCount, roamingDurationImpactedEntities, roamingDurationFailureImpactedEntities, roamingDurationFailureMetrics
connectionSpeed
connectionSpeedAverage, connectionSpeedSuccessPercentage, connectionSpeedSuccessCount, connectionSpeedTotalCount, connectionSpeedFailureCount, connectionSpeedClientCount, connectionSpeedImpactedEntities, connectionSpeedFailureImpactedEntities, connectionSpeedFailureMetrics
Examples:
view=connectionSpeed
 (single view requested)
view=roamingDuration&view=roamingAttempts
 (multiple views requested)

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

						"ap_count": &schema.Schema{
							Description: `Ap Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_average": &schema.Schema{
							Description: `Connection Speed Average`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_client_count": &schema.Schema{
							Description: `Connection Speed Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_failure_count": &schema.Schema{
							Description: `Connection Speed Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"connection_speed_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"connection_speed_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"connection_speed_success_count": &schema.Schema{
							Description: `Connection Speed Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_success_percentage": &schema.Schema{
							Description: `Connection Speed Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connection_speed_total_count": &schema.Schema{
							Description: `Connection Speed Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_average": &schema.Schema{
							Description: `Coverage Average`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_client_count": &schema.Schema{
							Description: `Coverage Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_failure_count": &schema.Schema{
							Description: `Coverage Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"coverage_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"coverage_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"coverage_success_count": &schema.Schema{
							Description: `Coverage Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_success_percentage": &schema.Schema{
							Description: `Coverage Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"coverage_total_count": &schema.Schema{
							Description: `Coverage Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"onboarding_attempts_client_count": &schema.Schema{
							Description: `Onboarding Attempts Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_attempts_failure_count": &schema.Schema{
							Description: `Onboarding Attempts Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_attempts_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_attempts_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_attempts_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_attempts_success_count": &schema.Schema{
							Description: `Onboarding Attempts Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_attempts_success_percentage": &schema.Schema{
							Description: `Onboarding Attempts Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_attempts_total_count": &schema.Schema{
							Description: `Onboarding Attempts Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_average": &schema.Schema{
							Description: `Onboarding Duration Average`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_client_count": &schema.Schema{
							Description: `Onboarding Duration Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_failure_count": &schema.Schema{
							Description: `Onboarding Duration Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_duration_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_duration_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"onboarding_duration_success_count": &schema.Schema{
							Description: `Onboarding Duration Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_success_percentage": &schema.Schema{
							Description: `Onboarding Duration Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"onboarding_duration_total_count": &schema.Schema{
							Description: `Onboarding Duration Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_attempts_client_count": &schema.Schema{
							Description: `Roaming Attempts Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_attempts_failure_count": &schema.Schema{
							Description: `Roaming Attempts Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_attempts_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_attempts_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_attempts_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_attempts_success_count": &schema.Schema{
							Description: `Roaming Attempts Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_attempts_success_percentage": &schema.Schema{
							Description: `Roaming Attempts Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_attempts_total_count": &schema.Schema{
							Description: `Roaming Attempts Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_average": &schema.Schema{
							Description: `Roaming Duration Average`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_client_count": &schema.Schema{
							Description: `Roaming Duration Client Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_failure_count": &schema.Schema{
							Description: `Roaming Duration Failure Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_failure_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_duration_failure_metrics": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"failure_ap_count": &schema.Schema{
										Description: `Failure Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_client_count": &schema.Schema{
										Description: `Failure Client Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"failure_percentage": &schema.Schema{
										Description: `Failure Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_duration_impacted_entities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_count": &schema.Schema{
										Description: `Ap Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"building_count": &schema.Schema{
										Description: `Building Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"floor_count": &schema.Schema{
										Description: `Floor Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"sites_count": &schema.Schema{
										Description: `Sites Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"roaming_duration_success_count": &schema.Schema{
							Description: `Roaming Duration Success Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_success_percentage": &schema.Schema{
							Description: `Roaming Duration Success Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"roaming_duration_total_count": &schema.Schema{
							Description: `Roaming Duration Total Count`,
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

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_type": &schema.Schema{
							Description: `Site Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteKpiSummariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID, okTaskID := d.GetOk("task_id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteHierarchy, okSiteHierarchy := d.GetOk("site_hierarchy")
	vSiteHierarchyID, okSiteHierarchyID := d.GetOk("site_hierarchy_id")
	vSiteID, okSiteID := d.GetOk("site_id")
	vSiteType, okSiteType := d.GetOk("site_type")
	vSSID, okSSID := d.GetOk("ssid")
	vBand, okBand := d.GetOk("band")
	vFailureCategory, okFailureCategory := d.GetOk("failure_category")
	vFailureReason, okFailureReason := d.GetOk("failure_reason")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters")

		headerParams1 := dnacentersdkgo.GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams{}
		queryParams1 := dnacentersdkgo.GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
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
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okSiteType {
			queryParams1.SiteType = vSiteType.(string)
		}
		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okBand {
			queryParams1.Band = vBand.(string)
		}
		if okFailureCategory {
			queryParams1.FailureCategory = vFailureCategory.(string)
		}
		if okFailureReason {
			queryParams1.FailureReason = vFailureReason.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
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
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sites.GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters", err,
				"Failure at GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItems(items *[]dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["site_id"] = item.SiteID
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_type"] = item.SiteType
		respItem["ap_count"] = item.ApCount
		respItem["coverage_average"] = item.CoverageAverage
		respItem["coverage_success_percentage"] = item.CoverageSuccessPercentage
		respItem["coverage_success_count"] = item.CoverageSuccessCount
		respItem["coverage_total_count"] = item.CoverageTotalCount
		respItem["coverage_failure_count"] = item.CoverageFailureCount
		respItem["coverage_client_count"] = item.CoverageClientCount
		respItem["coverage_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageImpactedEntities(item.CoverageImpactedEntities)
		respItem["coverage_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageFailureImpactedEntities(item.CoverageFailureImpactedEntities)
		respItem["coverage_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageFailureMetrics(item.CoverageFailureMetrics)
		respItem["onboarding_attempts_success_percentage"] = item.OnboardingAttemptsSuccessPercentage
		respItem["onboarding_attempts_success_count"] = item.OnboardingAttemptsSuccessCount
		respItem["onboarding_attempts_total_count"] = item.OnboardingAttemptsTotalCount
		respItem["onboarding_attempts_failure_count"] = item.OnboardingAttemptsFailureCount
		respItem["onboarding_attempts_client_count"] = item.OnboardingAttemptsClientCount
		respItem["onboarding_attempts_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsImpactedEntities(item.OnboardingAttemptsImpactedEntities)
		respItem["onboarding_attempts_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsFailureImpactedEntities(item.OnboardingAttemptsFailureImpactedEntities)
		respItem["onboarding_attempts_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsFailureMetrics(item.OnboardingAttemptsFailureMetrics)
		respItem["onboarding_duration_average"] = item.OnboardingDurationAverage
		respItem["onboarding_duration_success_percentage"] = item.OnboardingDurationSuccessPercentage
		respItem["onboarding_duration_success_count"] = item.OnboardingDurationSuccessCount
		respItem["onboarding_duration_total_count"] = item.OnboardingDurationTotalCount
		respItem["onboarding_duration_failure_count"] = item.OnboardingDurationFailureCount
		respItem["onboarding_duration_client_count"] = item.OnboardingDurationClientCount
		respItem["onboarding_duration_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationImpactedEntities(item.OnboardingDurationImpactedEntities)
		respItem["onboarding_duration_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationFailureImpactedEntities(item.OnboardingDurationFailureImpactedEntities)
		respItem["onboarding_duration_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationFailureMetrics(item.OnboardingDurationFailureMetrics)
		respItem["roaming_attempts_success_percentage"] = item.RoamingAttemptsSuccessPercentage
		respItem["roaming_attempts_success_count"] = item.RoamingAttemptsSuccessCount
		respItem["roaming_attempts_total_count"] = item.RoamingAttemptsTotalCount
		respItem["roaming_attempts_failure_count"] = item.RoamingAttemptsFailureCount
		respItem["roaming_attempts_client_count"] = item.RoamingAttemptsClientCount
		respItem["roaming_attempts_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsImpactedEntities(item.RoamingAttemptsImpactedEntities)
		respItem["roaming_attempts_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsFailureImpactedEntities(item.RoamingAttemptsFailureImpactedEntities)
		respItem["roaming_attempts_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsFailureMetrics(item.RoamingAttemptsFailureMetrics)
		respItem["roaming_duration_average"] = item.RoamingDurationAverage
		respItem["roaming_duration_success_percentage"] = item.RoamingDurationSuccessPercentage
		respItem["roaming_duration_success_count"] = item.RoamingDurationSuccessCount
		respItem["roaming_duration_total_count"] = item.RoamingDurationTotalCount
		respItem["roaming_duration_failure_count"] = item.RoamingDurationFailureCount
		respItem["roaming_duration_client_count"] = item.RoamingDurationClientCount
		respItem["roaming_duration_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationImpactedEntities(item.RoamingDurationImpactedEntities)
		respItem["roaming_duration_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationFailureImpactedEntities(item.RoamingDurationFailureImpactedEntities)
		respItem["roaming_duration_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationFailureMetrics(item.RoamingDurationFailureMetrics)
		respItem["connection_speed_average"] = item.ConnectionSpeedAverage
		respItem["connection_speed_success_percentage"] = item.ConnectionSpeedSuccessPercentage
		respItem["connection_speed_success_count"] = item.ConnectionSpeedSuccessCount
		respItem["connection_speed_total_count"] = item.ConnectionSpeedTotalCount
		respItem["connection_speed_failure_count"] = item.ConnectionSpeedFailureCount
		respItem["connection_speed_client_count"] = item.ConnectionSpeedClientCount
		respItem["connection_speed_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedImpactedEntities(item.ConnectionSpeedImpactedEntities)
		respItem["connection_speed_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedFailureImpactedEntities(item.ConnectionSpeedFailureImpactedEntities)
		respItem["connection_speed_failure_metrics"] = flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedFailureMetrics(item.ConnectionSpeedFailureMetrics)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsCoverageFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingAttemptsFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsOnboardingDurationFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingAttemptsFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsRoamingDurationFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureImpactedEntities) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["building_count"] = item.BuildingCount
	respItem["floor_count"] = item.FloorCount
	respItem["sites_count"] = item.SitesCount
	respItem["ap_count"] = item.ApCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersItemsConnectionSpeedFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureMetrics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["failure_ap_count"] = item.FailureApCount
	respItem["failure_client_count"] = item.FailureClientCount
	respItem["failure_percentage"] = item.FailurePercentage

	return []map[string]interface{}{
		respItem,
	}

}
