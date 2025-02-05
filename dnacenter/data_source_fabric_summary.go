package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Read Fabric summary for overall deployment. Get an aggregated summary of all fabric entities in a deployment including
the entity health.
This data source provides the latest health data until the given *endTime*. If data is not ready for the provided
endTime, the request will fail with error code *400 Bad Request*, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When *endTime* is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
fabricSummary-1.0.1-oas3-resolved.yaml
`,

		ReadContext: dataSourceFabricSummaryRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
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

						"protocol_summaries": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fabric_device_count": &schema.Schema{
										Description: `Fabric Device Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"fabric_site_count": &schema.Schema{
										Description: `Fabric Site Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"fabric_site_fair_health_count": &schema.Schema{
										Description: `Fabric Site Fair Health Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"fabric_site_good_health_count": &schema.Schema{
										Description: `Fabric Site Good Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"fabric_site_good_health_percentage": &schema.Schema{
										Description: `Fabric Site Good Health Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"fabric_site_no_health_count": &schema.Schema{
										Description: `Fabric Site No Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"fabric_site_poor_health_count": &schema.Schema{
										Description: `Fabric Site Poor Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"ip_transit_network_count": &schema.Schema{
										Description: `Ip Transit Network Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l2_vn_count": &schema.Schema{
										Description: `L2 Vn Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l2_vn_fair_health_count": &schema.Schema{
										Description: `L2 Vn Fair Health Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"l2_vn_good_health_count": &schema.Schema{
										Description: `L2 Vn Good Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l2_vn_good_health_percentage": &schema.Schema{
										Description: `L2 Vn Good Health Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l2_vn_no_health_count": &schema.Schema{
										Description: `L2 Vn No Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l2_vn_poor_health_count": &schema.Schema{
										Description: `L2 Vn Poor Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l3_vn_count": &schema.Schema{
										Description: `L3 Vn Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l3_vn_fair_health_count": &schema.Schema{
										Description: `L3 Vn Fair Health Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"l3_vn_good_health_count": &schema.Schema{
										Description: `L3 Vn Good Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l3_vn_good_health_percentage": &schema.Schema{
										Description: `L3 Vn Good Health Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l3_vn_no_health_count": &schema.Schema{
										Description: `L3 Vn No Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"l3_vn_poor_health_count": &schema.Schema{
										Description: `L3 Vn Poor Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"network_segment_protocol": &schema.Schema{
										Description: `Network Segment Protocol`,
										Type:        schema.TypeString,
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

									"transit_network_count": &schema.Schema{
										Description: `Transit Network Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"transit_network_fair_health_count": &schema.Schema{
										Description: `Transit Network Fair Health Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"transit_network_good_health_count": &schema.Schema{
										Description: `Transit Network Good Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"transit_network_good_health_percentage": &schema.Schema{
										Description: `Transit Network Good Health Percentage`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"transit_network_no_health_count": &schema.Schema{
										Description: `Transit Network No Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"transit_network_poor_health_count": &schema.Schema{
										Description: `Transit Network Poor Health Count`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceFabricSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadFabricEntitySummary")

		headerParams1 := dnacentersdkgo.ReadFabricEntitySummaryHeaderParams{}
		queryParams1 := dnacentersdkgo.ReadFabricEntitySummaryQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sda.ReadFabricEntitySummary(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadFabricEntitySummary", err,
				"Failure at ReadFabricEntitySummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaReadFabricEntitySummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadFabricEntitySummary response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaReadFabricEntitySummaryItem(item *dnacentersdkgo.ResponseSdaReadFabricEntitySummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["protocol_summaries"] = flattenSdaReadFabricEntitySummaryItemProtocolSummaries(item.ProtocolSummaries)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaReadFabricEntitySummaryItemProtocolSummaries(items *[]dnacentersdkgo.ResponseSdaReadFabricEntitySummaryResponseProtocolSummaries) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["fabric_site_good_health_count"] = item.FabricSiteGoodHealthCount
		respItem["fabric_site_count"] = item.FabricSiteCount
		respItem["fabric_site_good_health_percentage"] = item.FabricSiteGoodHealthPercentage
		respItem["fabric_site_no_health_count"] = item.FabricSiteNoHealthCount
		respItem["fabric_site_poor_health_count"] = item.FabricSitePoorHealthCount
		respItem["fabric_site_fair_health_count"] = item.FabricSiteFairHealthCount
		respItem["l3_vn_good_health_count"] = item.L3VnGoodHealthCount
		respItem["l3_vn_count"] = item.L3VnCount
		respItem["l3_vn_good_health_percentage"] = item.L3VnGoodHealthPercentage
		respItem["l3_vn_no_health_count"] = item.L3VnNoHealthCount
		respItem["l3_vn_fair_health_count"] = item.L3VnFairHealthCount
		respItem["l3_vn_poor_health_count"] = item.L3VnPoorHealthCount
		respItem["l2_vn_good_health_count"] = item.L2VnGoodHealthCount
		respItem["l2_vn_count"] = item.L2VnCount
		respItem["l2_vn_good_health_percentage"] = item.L2VnGoodHealthPercentage
		respItem["l2_vn_no_health_count"] = item.L2VnNoHealthCount
		respItem["l2_vn_poor_health_count"] = item.L2VnPoorHealthCount
		respItem["l2_vn_fair_health_count"] = item.L2VnFairHealthCount
		respItem["transit_network_good_health_count"] = item.TransitNetworkGoodHealthCount
		respItem["transit_network_count"] = item.TransitNetworkCount
		respItem["transit_network_good_health_percentage"] = item.TransitNetworkGoodHealthPercentage
		respItem["transit_network_no_health_count"] = item.TransitNetworkNoHealthCount
		respItem["transit_network_poor_health_count"] = item.TransitNetworkPoorHealthCount
		respItem["transit_network_fair_health_count"] = item.TransitNetworkFairHealthCount
		respItem["ip_transit_network_count"] = item.IPTransitNetworkCount
		respItem["fabric_device_count"] = item.FabricDeviceCount
		respItem["p1_issue_count"] = item.P1IssueCount
		respItem["p2_issue_count"] = item.P2IssueCount
		respItem["p3_issue_count"] = item.P3IssueCount
		respItem["network_segment_protocol"] = item.NetworkSegmentProtocol
		respItems = append(respItems, respItem)
	}
	return respItems
}
