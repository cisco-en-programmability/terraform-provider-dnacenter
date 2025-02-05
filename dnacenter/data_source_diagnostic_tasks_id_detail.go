package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiagnosticTasksIDDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source retrieves the details of the diagnostic task identified by the specified *id*.
`,

		ReadContext: dataSourceDiagnosticTasksIDDetailRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The *id* of the diagnostic task to be retrieved
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Description: `Any data associated with this task; the value may vary significantly across different tasks
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Description: `An error code if in case this task has failed in its execution
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_reason": &schema.Schema{
							Description: `A textual description indicating the reason why a task has failed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"progress": &schema.Schema{
							Description: `A textual representation which indicates the progress of this task; the value may vary significantly across different tasks
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiagnosticTasksIDDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesDiagnosticTaskDetailsByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.HealthAndPerformance.RetrievesDiagnosticTaskDetailsByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesDiagnosticTaskDetailsByID", err,
				"Failure at RetrievesDiagnosticTaskDetailsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceRetrievesDiagnosticTaskDetailsByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesDiagnosticTaskDetailsByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceRetrievesDiagnosticTaskDetailsByIDItem(item *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesDiagnosticTaskDetailsByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["data"] = item.Data
	respItem["progress"] = item.Progress
	respItem["error_code"] = item.ErrorCode
	respItem["failure_reason"] = item.FailureReason
	return []map[string]interface{}{
		respItem,
	}
}
