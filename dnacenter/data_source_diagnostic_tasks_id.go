package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiagnosticTasksID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- This data source retrieves the diagnostic task identified by the specified *id*.
`,

		ReadContext: dataSourceDiagnosticTasksIDRead,
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

						"end_time": &schema.Schema{
							Description: `An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `The ID of this task
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"result_location": &schema.Schema{
							Description: `A server-relative URL indicating where additional task-specific details may be found
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"root_id": &schema.Schema{
							Description: `The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Summarizes the status of a task
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"updated_time": &schema.Schema{
							Description: `The last modification date and time of this task, expressed in Unix epoch time format to the millisecond precision.
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

func dataSourceDiagnosticTasksIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesDiagnosticTaskByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.HealthAndPerformance.RetrievesDiagnosticTaskByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesDiagnosticTaskByID", err,
				"Failure at RetrievesDiagnosticTaskByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenHealthAndPerformanceRetrievesDiagnosticTaskByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesDiagnosticTaskByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceRetrievesDiagnosticTaskByIDItem(item *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesDiagnosticTaskByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["root_id"] = item.RootID
	respItem["parent_id"] = item.ParentID
	respItem["start_time"] = item.StartTime
	respItem["result_location"] = item.ResultLocation
	respItem["status"] = item.Status
	respItem["updated_time"] = item.UpdatedTime
	respItem["end_time"] = item.EndTime
	return []map[string]interface{}{
		respItem,
	}
}
