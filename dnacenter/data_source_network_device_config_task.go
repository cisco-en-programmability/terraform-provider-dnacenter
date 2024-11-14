package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfigTask() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Archive.

- Returns a config task result details by specified id
`,

		ReadContext: dataSourceNetworkDeviceConfigTaskRead,
		Schema: map[string]*schema.Schema{
			"parent_task_id": &schema.Schema{
				Description: `parentTaskId query parameter. task Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"completion_time": &schema.Schema{
							Description: `Timestamp when the task was completed.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"detail_message": &schema.Schema{
							Description: `Details of the task, if available.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `UUID of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_ip_address": &schema.Schema{
							Description: `IP address of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_code": &schema.Schema{
							Description: `Error code if the task failed.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_message": &schema.Schema{
							Description: `Failure message, if the task failed.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"host_name": &schema.Schema{
							Description: `Host name of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_task_id": &schema.Schema{
							Description: `UUID of the parent task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Timestamp when the task started.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"task_id": &schema.Schema{
							Description: `UUID of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"task_status": &schema.Schema{
							Description: `Status of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"task_type": &schema.Schema{
							Description: `Task type can be 0,1,2 etc(ARCHIVE_RUNNING(0),ARCHIVE_STARTUP(1),ARCHIVE_VLAN(2),DEPLOY_RUNNING(3),DEPLOY_STARTUP(4),DEPLOY_VLAN(5),COPY_RUNNING_TO_STARTUP(6)) 
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

func dataSourceNetworkDeviceConfigTaskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vParentTaskID := d.Get("parent_task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConfigTaskDetails")
		queryParams1 := dnacentersdkgo.GetConfigTaskDetailsQueryParams{}

		queryParams1.ParentTaskID = vParentTaskID.(string)

		response1, restyResp1, err := client.Compliance.GetConfigTaskDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConfigTaskDetails", err,
				"Failure at GetConfigTaskDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetConfigTaskDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConfigTaskDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetConfigTaskDetailsItems(items *[]dnacentersdkgo.ResponseComplianceGetConfigTaskDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["start_time"] = item.StartTime
		respItem["error_code"] = item.ErrorCode
		respItem["device_id"] = item.DeviceID
		respItem["task_id"] = item.TaskID
		respItem["task_status"] = item.TaskStatus
		respItem["parent_task_id"] = item.ParentTaskID
		respItem["device_ip_address"] = item.DeviceIPAddress
		respItem["detail_message"] = item.DetailMessage
		respItem["failure_message"] = item.FailureMessage
		respItem["task_type"] = item.TaskType
		respItem["completion_time"] = item.CompletionTime
		respItem["host_name"] = item.HostName
		respItems = append(respItems, respItem)
	}
	return respItems
}
