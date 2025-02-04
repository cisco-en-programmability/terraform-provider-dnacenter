package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceTasksCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- returns a count of the number of assurance tasks that are not expired For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceTasksCountRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Description: `status query parameter. used to get a subset of tasks by their status
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

func dataSourceAssuranceTasksCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStatus, okStatus := d.GetOk("status")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist")

		headerParams1 := dnacentersdkgo.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams{}

		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Task.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist", err,
				"Failure at RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistItem(item *dnacentersdkgo.ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
