package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkApplicationsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Applications.

- Retrieves the number of network applications by applying basic filtering. If startTime and endTime are not provided,
the API defaults to the last 24 hours. *siteId* is mandatory. *siteId* must be a site UUID of a building. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceNetworkApplicationsCountRead,
		Schema: map[string]*schema.Schema{
			"application_name": &schema.Schema{
				Description: `applicationName query parameter. Name of the application for which the experience data is intended.
Examples:
*applicationName=webex* (single applicationName requested)
*applicationName=webex&applicationName=teams* (multiple applicationName requested)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"business_relevance": &schema.Schema{
				Description: `businessRelevance query parameter. The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
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
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The site UUID without the top level hierarchy. *siteId* is mandatory. *siteId* must be a site UUID of a building. (Ex."buildingUuid") Examples: *siteId=buildingUuid* (single siteId requested) *siteId=buildingUuid1&siteId=buildingUuid2* (multiple siteId requested)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. In the context of a network application, SSID refers to the name of the wireless network to which the client connects.
Examples:
*ssid=Alpha* (single ssid requested)
*ssid=Alpha&ssid=Guest* (multiple ssid requested)
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

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkApplicationsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteID := d.Get("site_id")
	vSSID, okSSID := d.GetOk("ssid")
	vApplicationName, okApplicationName := d.GetOk("application_name")
	vBusinessRelevance, okBusinessRelevance := d.GetOk("business_relevance")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering")

		headerParams1 := dnacentersdkgo.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		queryParams1.SiteID = vSiteID.(string)

		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okApplicationName {
			queryParams1.ApplicationName = vApplicationName.(string)
		}
		if okBusinessRelevance {
			queryParams1.BusinessRelevance = vBusinessRelevance.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Applications.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering", err,
				"Failure at RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringItem(item *dnacentersdkgo.ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
