package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTransitNetworkHealthSummariesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get health summary for a specific transit Network by providing the unique transit networks id in the url path.
This data source provides the latest health data until the given *endTime*. If data is not ready for the provided
endTime, the request will fail with error code *400 Bad Request*, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When *endTime* is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
transitNetworkHealthSummaries-1.0.1-resolved.yaml
`,

		ReadContext: dataSourceTransitNetworkHealthSummariesIDRead,
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
				Description: `id path parameter. The unique transit network id, Ex “1551156a-bc97-3c63-aeda-8a6d3765b5b9"
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites.
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

						"bgp_tcp_fair_health_device_count": &schema.Schema{
							Description: `Bgp Tcp Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_tcp_good_health_device_count": &schema.Schema{
							Description: `Bgp Tcp Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_tcp_health_percentage": &schema.Schema{
							Description: `Bgp Tcp Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_tcp_poor_health_device_count": &schema.Schema{
							Description: `Bgp Tcp Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_tcp_total_device_count": &schema.Schema{
							Description: `Bgp Tcp Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"control_plane_count": &schema.Schema{
							Description: `Control Plane Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"fabric_sites_count": &schema.Schema{
							Description: `Fabric Sites Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"fair_health_device_count": &schema.Schema{
							Description: `Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"good_health_device_count": &schema.Schema{
							Description: `Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"good_health_percentage": &schema.Schema{
							Description: `Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"internet_avail_transit_fair_health_device_count": &schema.Schema{
							Description: `Internet Avail Transit Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_transit_good_health_device_count": &schema.Schema{
							Description: `Internet Avail Transit Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_transit_health_percentage": &schema.Schema{
							Description: `Internet Avail Transit Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_transit_poor_health_device_count": &schema.Schema{
							Description: `Internet Avail Transit Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_transit_total_device_count": &schema.Schema{
							Description: `Internet Avail Transit Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_transit_fair_health_device_count": &schema.Schema{
							Description: `Lisp Transit Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_transit_good_health_device_count": &schema.Schema{
							Description: `Lisp Transit Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_transit_health_percentage": &schema.Schema{
							Description: `Lisp Transit Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_transit_poor_health_device_count": &schema.Schema{
							Description: `Lisp Transit Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"lisp_transit_total_device_count": &schema.Schema{
							Description: `Lisp Transit Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"poor_health_device_count": &schema.Schema{
							Description: `Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_transit_fair_health_device_count": &schema.Schema{
							Description: `Pubsub Transit Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_transit_good_health_device_count": &schema.Schema{
							Description: `Pubsub Transit Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_transit_health_percentage": &schema.Schema{
							Description: `Pubsub Transit Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_transit_poor_health_device_count": &schema.Schema{
							Description: `Pubsub Transit Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_transit_total_device_count": &schema.Schema{
							Description: `Pubsub Transit Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_device_count": &schema.Schema{
							Description: `Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_control_plane_fair_health_device_count": &schema.Schema{
							Description: `Transit Control Plane Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_control_plane_good_health_device_count": &schema.Schema{
							Description: `Transit Control Plane Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_control_plane_health_percentage": &schema.Schema{
							Description: `Transit Control Plane Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_control_plane_poor_health_device_count": &schema.Schema{
							Description: `Transit Control Plane Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_control_plane_total_device_count": &schema.Schema{
							Description: `Transit Control Plane Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_services_fair_health_device_count": &schema.Schema{
							Description: `Transit Services Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_services_good_health_device_count": &schema.Schema{
							Description: `Transit Services Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_services_health_percentage": &schema.Schema{
							Description: `Transit Services Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_services_poor_health_device_count": &schema.Schema{
							Description: `Transit Services Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"transit_services_total_device_count": &schema.Schema{
							Description: `Transit Services Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"transit_type": &schema.Schema{
							Description: `Transit Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTransitNetworkHealthSummariesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vEndTime, okEndTime := d.GetOk("end_time")
	vStartTime, okStartTime := d.GetOk("start_time")
	vAttribute, okAttribute := d.GetOk("attribute")
	vView, okView := d.GetOk("view")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadTransitNetworkWithItsHealthSummaryFromID")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.ReadTransitNetworkWithItsHealthSummaryFromIDHeaderParams{}
		queryParams1 := dnacentersdkgo.ReadTransitNetworkWithItsHealthSummaryFromIDQueryParams{}

		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sda.ReadTransitNetworkWithItsHealthSummaryFromID(vvID, &headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadTransitNetworkWithItsHealthSummaryFromID", err,
				"Failure at ReadTransitNetworkWithItsHealthSummaryFromID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaReadTransitNetworkWithItsHealthSummaryFromIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadTransitNetworkWithItsHealthSummaryFromID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaReadTransitNetworkWithItsHealthSummaryFromIDItem(item *dnacentersdkgo.ResponseSdaReadTransitNetworkWithItsHealthSummaryFromIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["control_plane_count"] = item.ControlPlaneCount
	respItem["transit_type"] = item.TransitType
	respItem["fabric_sites_count"] = item.FabricSitesCount
	respItem["good_health_percentage"] = item.GoodHealthPercentage
	respItem["good_health_device_count"] = item.GoodHealthDeviceCount
	respItem["total_device_count"] = item.TotalDeviceCount
	respItem["poor_health_device_count"] = item.PoorHealthDeviceCount
	respItem["fair_health_device_count"] = item.FairHealthDeviceCount
	respItem["transit_control_plane_health_percentage"] = item.TransitControlPlaneHealthPercentage
	respItem["transit_control_plane_total_device_count"] = item.TransitControlPlaneTotalDeviceCount
	respItem["transit_control_plane_good_health_device_count"] = item.TransitControlPlaneGoodHealthDeviceCount
	respItem["transit_control_plane_poor_health_device_count"] = item.TransitControlPlanePoorHealthDeviceCount
	respItem["transit_control_plane_fair_health_device_count"] = item.TransitControlPlaneFairHealthDeviceCount
	respItem["transit_services_health_percentage"] = item.TransitServicesHealthPercentage
	respItem["transit_services_total_device_count"] = item.TransitServicesTotalDeviceCount
	respItem["transit_services_good_health_device_count"] = item.TransitServicesGoodHealthDeviceCount
	respItem["transit_services_poor_health_device_count"] = item.TransitServicesPoorHealthDeviceCount
	respItem["transit_services_fair_health_device_count"] = item.TransitServicesFairHealthDeviceCount
	respItem["pubsub_transit_health_percentage"] = item.PubsubTransitHealthPercentage
	respItem["pubsub_transit_total_device_count"] = item.PubsubTransitTotalDeviceCount
	respItem["pubsub_transit_good_health_device_count"] = item.PubsubTransitGoodHealthDeviceCount
	respItem["pubsub_transit_poor_health_device_count"] = item.PubsubTransitPoorHealthDeviceCount
	respItem["pubsub_transit_fair_health_device_count"] = item.PubsubTransitFairHealthDeviceCount
	respItem["lisp_transit_health_percentage"] = item.LispTransitHealthPercentage
	respItem["lisp_transit_total_device_count"] = item.LispTransitTotalDeviceCount
	respItem["lisp_transit_good_health_device_count"] = item.LispTransitGoodHealthDeviceCount
	respItem["lisp_transit_poor_health_device_count"] = item.LispTransitPoorHealthDeviceCount
	respItem["lisp_transit_fair_health_device_count"] = item.LispTransitFairHealthDeviceCount
	respItem["internet_avail_transit_health_percentage"] = item.InternetAvailTransitHealthPercentage
	respItem["internet_avail_transit_total_device_count"] = item.InternetAvailTransitTotalDeviceCount
	respItem["internet_avail_transit_good_health_device_count"] = item.InternetAvailTransitGoodHealthDeviceCount
	respItem["internet_avail_transit_poor_health_device_count"] = item.InternetAvailTransitPoorHealthDeviceCount
	respItem["internet_avail_transit_fair_health_device_count"] = item.InternetAvailTransitFairHealthDeviceCount
	respItem["bgp_tcp_health_percentage"] = item.BgpTCPHealthPercentage
	respItem["bgp_tcp_total_device_count"] = item.BgpTCPTotalDeviceCount
	respItem["bgp_tcp_good_health_device_count"] = item.BgpTCPGoodHealthDeviceCount
	respItem["bgp_tcp_poor_health_device_count"] = item.BgpTCPPoorHealthDeviceCount
	respItem["bgp_tcp_fair_health_device_count"] = item.BgpTCPFairHealthDeviceCount
	return []map[string]interface{}{
		respItem,
	}
}
