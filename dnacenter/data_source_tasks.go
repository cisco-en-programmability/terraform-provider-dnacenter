package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Returns task(s) based on filter criteria

- Returns the task with the given ID
`,

		ReadContext: dataSourceTasksRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. This is the epoch millisecond end time upto which task records need to be fetched
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. the *id* of the task to retrieve
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": &schema.Schema{
				Description: `parentId query parameter. Fetch tasks that have this parent Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_id": &schema.Schema{
				Description: `rootId query parameter. Fetch tasks that have this root Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. This is the epoch millisecond start time from which tasks need to be fetched
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
`,
				Type:     schema.TypeString,
				Optional: true,
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
							Description: `The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
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
							Description: `The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
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
							Type:     schema.TypeString,
							Computed: true,
						},

						"updated_time": &schema.Schema{
							Description: `A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
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
							Description: `The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
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
							Description: `The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
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
							Type:     schema.TypeString,
							Computed: true,
						},

						"updated_time": &schema.Schema{
							Description: `A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
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

func dataSourceTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vParentID, okParentID := d.GetOk("parent_id")
	vRootID, okRootID := d.GetOk("root_id")
	vStatus, okStatus := d.GetOk("status")
	vID, okID := d.GetOk("id")

	method1 := []bool{okOffset, okLimit, okSortBy, okOrder, okStartTime, okEndTime, okParentID, okRootID, okStatus}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTasks")
		queryParams1 := dnacentersdkgo.GetTasksQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(int)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(int)
		}
		if okParentID {
			queryParams1.ParentID = vParentID.(string)
		}
		if okRootID {
			queryParams1.RootID = vRootID.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}

		response1, restyResp1, err := client.Task.GetTasks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTasks", err,
				"Failure at GetTasks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTaskGetTasksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTasks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTasksByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Task.GetTasksByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTasksByID", err,
				"Failure at GetTasksByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTaskGetTasksByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTasksByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskGetTasksItems(items *[]dnacentersdkgo.ResponseTaskGetTasksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["end_time"] = item.EndTime
		respItem["id"] = item.ID
		respItem["updated_time"] = item.UpdatedTime
		respItem["parent_id"] = item.ParentID
		respItem["result_location"] = item.ResultLocation
		respItem["root_id"] = item.RootID
		respItem["start_time"] = item.StartTime
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTaskGetTasksByIDItem(item *dnacentersdkgo.ResponseTaskGetTasksByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["end_time"] = item.EndTime
	respItem["id"] = item.ID
	respItem["updated_time"] = item.UpdatedTime
	respItem["parent_id"] = item.ParentID
	respItem["result_location"] = item.ResultLocation
	respItem["root_id"] = item.RootID
	respItem["start_time"] = item.StartTime
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}
