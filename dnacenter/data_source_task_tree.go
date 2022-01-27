package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskTree() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Returns a task with its children tasks by based on their id
`,

		ReadContext: dataSourceTaskTreeRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId path parameter. UUID of the Task
`,
				Type:     schema.TypeString,
				Required: true,
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
							Type:     schema.TypeString,
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
							Type:     schema.TypeString,
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

func dataSourceTaskTreeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTaskTree")
		vvTaskID := vTaskID.(string)

		response1, restyResp1, err := client.Task.GetTaskTree(vvTaskID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskTree", err,
				"Failure at GetTaskTree, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTaskGetTaskTreeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskTree response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskGetTaskTreeItems(items *[]dnacentersdkgo.ResponseTaskGetTaskTreeResponse) []map[string]interface{} {
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
		respItem["operation_id_list"] = flattenTaskGetTaskTreeItemsOperationIDList(item.OperationIDList)
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

func flattenTaskGetTaskTreeItemsOperationIDList(item *dnacentersdkgo.ResponseTaskGetTaskTreeResponseOperationIDList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
