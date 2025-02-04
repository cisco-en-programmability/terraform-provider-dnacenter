package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteKpiSummariesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Returns site analytics for the given site. For detailed information about the usage of the API, please refer to the
Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteKpiSummariesIDRead,
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
			"id": &schema.Schema{
				Description: `id path parameter. The Site UUID
`,
				Type:     schema.TypeString,
				Required: true,
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

			"item": &schema.Schema{
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

func dataSourceSiteKpiSummariesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vTaskID, okTaskID := d.GetOk("task_id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSSID, okSSID := d.GetOk("ssid")
	vBand, okBand := d.GetOk("band")
	vFailureCategory, okFailureCategory := d.GetOk("failure_category")
	vFailureReason, okFailureReason := d.GetOk("failure_reason")
	vView, okView := d.GetOk("view")
	vAttribute, okAttribute := d.GetOk("attribute")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAnalyticsForOneSite")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.GetSiteAnalyticsForOneSiteHeaderParams{}
		queryParams1 := dnacentersdkgo.GetSiteAnalyticsForOneSiteQueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
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
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sites.GetSiteAnalyticsForOneSite(vvID, &headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteAnalyticsForOneSite", err,
				"Failure at GetSiteAnalyticsForOneSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetSiteAnalyticsForOneSiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAnalyticsForOneSite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteAnalyticsForOneSiteItem(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["coverage_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemCoverageImpactedEntities(item.CoverageImpactedEntities)
	respItem["coverage_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemCoverageFailureImpactedEntities(item.CoverageFailureImpactedEntities)
	respItem["coverage_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemCoverageFailureMetrics(item.CoverageFailureMetrics)
	respItem["onboarding_attempts_success_percentage"] = item.OnboardingAttemptsSuccessPercentage
	respItem["onboarding_attempts_success_count"] = item.OnboardingAttemptsSuccessCount
	respItem["onboarding_attempts_total_count"] = item.OnboardingAttemptsTotalCount
	respItem["onboarding_attempts_failure_count"] = item.OnboardingAttemptsFailureCount
	respItem["onboarding_attempts_client_count"] = item.OnboardingAttemptsClientCount
	respItem["onboarding_attempts_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsImpactedEntities(item.OnboardingAttemptsImpactedEntities)
	respItem["onboarding_attempts_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsFailureImpactedEntities(item.OnboardingAttemptsFailureImpactedEntities)
	respItem["onboarding_attempts_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsFailureMetrics(item.OnboardingAttemptsFailureMetrics)
	respItem["onboarding_duration_average"] = item.OnboardingDurationAverage
	respItem["onboarding_duration_success_percentage"] = item.OnboardingDurationSuccessPercentage
	respItem["onboarding_duration_success_count"] = item.OnboardingDurationSuccessCount
	respItem["onboarding_duration_total_count"] = item.OnboardingDurationTotalCount
	respItem["onboarding_duration_failure_count"] = item.OnboardingDurationFailureCount
	respItem["onboarding_duration_client_count"] = item.OnboardingDurationClientCount
	respItem["onboarding_duration_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationImpactedEntities(item.OnboardingDurationImpactedEntities)
	respItem["onboarding_duration_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationFailureImpactedEntities(item.OnboardingDurationFailureImpactedEntities)
	respItem["onboarding_duration_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationFailureMetrics(item.OnboardingDurationFailureMetrics)
	respItem["roaming_attempts_success_percentage"] = item.RoamingAttemptsSuccessPercentage
	respItem["roaming_attempts_success_count"] = item.RoamingAttemptsSuccessCount
	respItem["roaming_attempts_total_count"] = item.RoamingAttemptsTotalCount
	respItem["roaming_attempts_failure_count"] = item.RoamingAttemptsFailureCount
	respItem["roaming_attempts_client_count"] = item.RoamingAttemptsClientCount
	respItem["roaming_attempts_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsImpactedEntities(item.RoamingAttemptsImpactedEntities)
	respItem["roaming_attempts_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsFailureImpactedEntities(item.RoamingAttemptsFailureImpactedEntities)
	respItem["roaming_attempts_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsFailureMetrics(item.RoamingAttemptsFailureMetrics)
	respItem["roaming_duration_average"] = item.RoamingDurationAverage
	respItem["roaming_duration_success_percentage"] = item.RoamingDurationSuccessPercentage
	respItem["roaming_duration_success_count"] = item.RoamingDurationSuccessCount
	respItem["roaming_duration_total_count"] = item.RoamingDurationTotalCount
	respItem["roaming_duration_failure_count"] = item.RoamingDurationFailureCount
	respItem["roaming_duration_client_count"] = item.RoamingDurationClientCount
	respItem["roaming_duration_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationImpactedEntities(item.RoamingDurationImpactedEntities)
	respItem["roaming_duration_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationFailureImpactedEntities(item.RoamingDurationFailureImpactedEntities)
	respItem["roaming_duration_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationFailureMetrics(item.RoamingDurationFailureMetrics)
	respItem["connection_speed_average"] = item.ConnectionSpeedAverage
	respItem["connection_speed_success_percentage"] = item.ConnectionSpeedSuccessPercentage
	respItem["connection_speed_success_count"] = item.ConnectionSpeedSuccessCount
	respItem["connection_speed_total_count"] = item.ConnectionSpeedTotalCount
	respItem["connection_speed_failure_count"] = item.ConnectionSpeedFailureCount
	respItem["connection_speed_client_count"] = item.ConnectionSpeedClientCount
	respItem["connection_speed_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedImpactedEntities(item.ConnectionSpeedImpactedEntities)
	respItem["connection_speed_failure_impacted_entities"] = flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedFailureImpactedEntities(item.ConnectionSpeedFailureImpactedEntities)
	respItem["connection_speed_failure_metrics"] = flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedFailureMetrics(item.ConnectionSpeedFailureMetrics)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSitesGetSiteAnalyticsForOneSiteItemCoverageImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemCoverageFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemCoverageFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureMetrics) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingAttemptsFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureMetrics) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemOnboardingDurationFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureMetrics) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingAttemptsFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureMetrics) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemRoamingDurationFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureMetrics) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedFailureImpactedEntities(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureImpactedEntities) []map[string]interface{} {
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

func flattenSitesGetSiteAnalyticsForOneSiteItemConnectionSpeedFailureMetrics(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureMetrics) []map[string]interface{} {
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
