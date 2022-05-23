package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Returns Task count
`,

		ReadContext: dataSourceTaskCountRead,
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
			"start_time": &schema.Schema{
				Description: `startTime query parameter. This is the epoch start time from which tasks need to be fetched
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

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTaskCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetTaskCount")
		queryParams1 := dnacentersdkgo.GetTaskCountQueryParams{}

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

		response1, restyResp1, err := client.Task.GetTaskCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskCount", err,
				"Failure at GetTaskCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTaskGetTaskCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskGetTaskCountItem(item *dnacentersdkgo.ResponseTaskGetTaskCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
