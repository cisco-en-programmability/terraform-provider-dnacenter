package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTask() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Returns task(s) based on filter criteria

- Returns a task by specified id
`,

		ReadContext: dataSourceTaskRead,
		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Description: `data query parameter. Fetch tasks that contains this data
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. This is the epoch end time upto which audit records need to be fetched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_code": &schema.Schema{
				Description: `errorCode query parameter. Fetch tasks that have this error code
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"failure_reason": &schema.Schema{
				Description: `failureReason query parameter. Fetch tasks that contains this failure reason
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_error": &schema.Schema{
				Description: `isError query parameter. Fetch tasks ended as success or failure. Valid values: true, false
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Sort order asc or dsc
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
			"progress": &schema.Schema{
				Description: `progress query parameter. Fetch tasks that contains this progress
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_type": &schema.Schema{
				Description: `serviceType query parameter. Fetch tasks with this service type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort results by this field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. This is the epoch start time from which tasks need to be fetched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_id": &schema.Schema{
				Description: `taskId path parameter. UUID of the Task
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Description: `username query parameter. Fetch tasks with this username
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_status_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"data": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_error": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"operation_id_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"progress": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"root_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"service_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
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

						"additional_status_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"data": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_error": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"operation_id_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"progress": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"root_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"service_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTaskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vData, okData := d.GetOk("data")
	vErrorCode, okErrorCode := d.GetOk("error_code")
	vServiceType, okServiceType := d.GetOk("service_type")
	vUsername, okUsername := d.GetOk("username")
	vProgress, okProgress := d.GetOk("progress")
	vIsError, okIsError := d.GetOk("is_error")
	vFailureReason, okFailureReason := d.GetOk("failure_reason")
	vParentID, okParentID := d.GetOk("parent_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vTaskID, okTaskID := d.GetOk("task_id")

	method1 := []bool{okStartTime, okEndTime, okData, okErrorCode, okServiceType, okUsername, okProgress, okIsError, okFailureReason, okParentID, okOffset, okLimit, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okTaskID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTasks")
		queryParams1 := dnacentersdkgo.GetTasksQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(string)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(string)
		}
		if okData {
			queryParams1.Data = vData.(string)
		}
		if okErrorCode {
			queryParams1.ErrorCode = vErrorCode.(string)
		}
		if okServiceType {
			queryParams1.ServiceType = vServiceType.(string)
		}
		if okUsername {
			queryParams1.Username = vUsername.(string)
		}
		if okProgress {
			queryParams1.Progress = vProgress.(string)
		}
		if okIsError {
			queryParams1.IsError = vIsError.(string)
		}
		if okFailureReason {
			queryParams1.FailureReason = vFailureReason.(string)
		}
		if okParentID {
			queryParams1.ParentID = vParentID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Task.GetTasks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTasks", err,
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
		log.Printf("[DEBUG] Selected method 2: GetTaskByID")
		vvTaskID := vTaskID.(string)

		response2, restyResp2, err := client.Task.GetTaskByID(vvTaskID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTaskGetTaskByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskByID response",
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
		respItem["additional_status_url"] = item.AdditionalStatusURL
		respItem["data"] = item.Data
		respItem["end_time"] = item.EndTime
		respItem["error_code"] = item.ErrorCode
		respItem["error_key"] = item.ErrorKey
		respItem["failure_reason"] = item.FailureReason
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["is_error"] = boolPtrToString(item.IsError)
		respItem["last_update"] = item.LastUpdate
		respItem["operation_id_list"] = flattenTaskGetTasksItemsOperationIDList(item.OperationIDList)
		respItem["parent_id"] = item.ParentID
		respItem["progress"] = item.Progress
		respItem["root_id"] = item.RootID
		respItem["service_type"] = item.ServiceType
		respItem["start_time"] = item.StartTime
		respItem["username"] = item.Username
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTaskGetTasksItemsOperationIDList(item *dnacentersdkgo.ResponseTaskGetTasksResponseOperationIDList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenTaskGetTaskByIDItem(item *dnacentersdkgo.ResponseTaskGetTaskByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["additional_status_url"] = item.AdditionalStatusURL
	respItem["data"] = item.Data
	respItem["end_time"] = item.EndTime
	respItem["error_code"] = item.ErrorCode
	respItem["error_key"] = item.ErrorKey
	respItem["failure_reason"] = item.FailureReason
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["is_error"] = boolPtrToString(item.IsError)
	respItem["last_update"] = item.LastUpdate
	respItem["operation_id_list"] = flattenTaskGetTaskByIDItemOperationIDList(item.OperationIDList)
	respItem["parent_id"] = item.ParentID
	respItem["progress"] = item.Progress
	respItem["root_id"] = item.RootID
	respItem["service_type"] = item.ServiceType
	respItem["start_time"] = item.StartTime
	respItem["username"] = item.Username
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTaskGetTaskByIDItemOperationIDList(item *dnacentersdkgo.ResponseTaskGetTaskByIDResponseOperationIDList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
