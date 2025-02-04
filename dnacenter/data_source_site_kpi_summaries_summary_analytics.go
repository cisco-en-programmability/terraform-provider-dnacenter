package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteKpiSummariesSummaryAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Get site analytics summary data for the given task id. For detailed information about the usage of the API, please
refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteKpiSummariesSummaryAnalyticsRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId query parameter. used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
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

			"item": &schema.Schema{
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
					},
				},
			},
		},
	}
}

func dataSourceSiteKpiSummariesSummaryAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAnalyticsSummaryDataForTheGivenTaskID")

		headerParams1 := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDHeaderParams{}
		queryParams1 := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams{}

		queryParams1.TaskID = vTaskID.(string)

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sites.GetSiteAnalyticsSummaryDataForTheGivenTaskID(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteAnalyticsSummaryDataForTheGivenTaskID", err,
				"Failure at GetSiteAnalyticsSummaryDataForTheGivenTaskID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAnalyticsSummaryDataForTheGivenTaskID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDItem(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDItemAttributes(item.Attributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDItemAttributes(items *[]dnacentersdkgo.ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponseAttributes) []map[string]interface{} {
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
