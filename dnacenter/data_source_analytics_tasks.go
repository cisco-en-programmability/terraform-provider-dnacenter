package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnalyticsTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AI Endpoint Analytics.

- Fetches the details of backend task. Task is typically created by making call to some other API that takes longer time
to execute.
`,

		ReadContext: dataSourceAnalyticsTasksRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId path parameter. Unique identifier for the task.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_info": &schema.Schema{
							Description: `Additional information about the task.
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"created_by": &schema.Schema{
							Description: `Name of the user that created the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"created_on": &schema.Schema{
							Description: `Task creation timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"errors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"code": &schema.Schema{
										Description: `Error code.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"details": &schema.Schema{
										Description: `Optional details about the error.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"index": &schema.Schema{
										Description: `Index of the input records which had error during processing. In case the input is not an array, or the error is not record specific, this will be -1.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"message": &schema.Schema{
										Description: `Error message.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Unique identifier for the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_on": &schema.Schema{
							Description: `Last update timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the task.
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

func dataSourceAnalyticsTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTaskDetails")
		vvTaskID := vTaskID.(string)

		response1, restyResp1, err := client.AIEndpointAnalytics.GetTaskDetails(vvTaskID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTaskDetails", err,
				"Failure at GetTaskDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAIEndpointAnalyticsGetTaskDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAIEndpointAnalyticsGetTaskDetailsItem(item *dnacentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["status"] = item.Status
	respItem["errors"] = flattenAIEndpointAnalyticsGetTaskDetailsItemErrors(item.Errors)
	respItem["additional_info"] = flattenAIEndpointAnalyticsGetTaskDetailsItemAdditionalInfo(item.AdditionalInfo)
	respItem["created_by"] = item.CreatedBy
	respItem["created_on"] = item.CreatedOn
	respItem["last_updated_on"] = item.LastUpdatedOn
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAIEndpointAnalyticsGetTaskDetailsItemErrors(items *[]dnacentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetailsErrors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["index"] = item.Index
		respItem["code"] = item.Code
		respItem["message"] = item.Message
		respItem["details"] = item.Details
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAIEndpointAnalyticsGetTaskDetailsItemAdditionalInfo(item *dnacentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetailsAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
