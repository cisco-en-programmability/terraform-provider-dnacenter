package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiagnosticValidationWorkflowsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- Retrieves the count of workflows that have been successfully submitted and are currently available.
`,

		ReadContext: dataSourceDiagnosticValidationWorkflowsCountRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Workflows started before the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"run_status": &schema.Schema{
				Description: `runStatus query parameter. Execution status of the workflow. If the workflow is successfully submitted, runStatus is *PENDING*. If the workflow execution has started, runStatus is *IN_PROGRESS*. If the workflow executed is completed with all validations executed, runStatus is *COMPLETED*. If the workflow execution fails while running validations, runStatus is *FAILED*.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Workflows started after the given time (as milliseconds since UNIX epoch).
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiagnosticValidationWorkflowsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vRunStatus, okRunStatus := d.GetOk("run_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfValidationWorkflows")
		queryParams1 := dnacentersdkgo.RetrievesTheCountOfValidationWorkflowsQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okRunStatus {
			queryParams1.RunStatus = vRunStatus.(string)
		}

		response1, restyResp1, err := client.HealthAndPerformance.RetrievesTheCountOfValidationWorkflows(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfValidationWorkflows", err,
				"Failure at RetrievesTheCountOfValidationWorkflows, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfValidationWorkflows response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsItem(item *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
