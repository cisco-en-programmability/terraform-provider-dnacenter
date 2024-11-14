package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexibleReportExecutions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get Execution Id by Report Id
`,

		ReadContext: dataSourceFlexibleReportExecutionsRead,
		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Description: `reportId path parameter. Id of the report
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_count": &schema.Schema{
							Description: `Total number of report executions
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"executions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Report execution end time (Represent the specified number of milliseconds since the epoch time)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"errors": &schema.Schema{
										Description: `Errors associated with the report execution
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"execution_id": &schema.Schema{
										Description: `Report ExecutionId (Unique UUID)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"process_status": &schema.Schema{
										Description: `Report execution status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"request_status": &schema.Schema{
										Description: `Report request status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Report execution start time (Represent the specified number of milliseconds since the epoch time)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"warnings": &schema.Schema{
										Description: `Warnings associated with the report execution
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"report_id": &schema.Schema{
							Description: `Report Id (Unique UUID)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_name": &schema.Schema{
							Description: `Name of the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_was_executed": &schema.Schema{
							Description: `Report execution status flag (true if execution is started, false if the execution is not started)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFlexibleReportExecutionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExecutionIDByReportID")
		vvReportID := vReportID.(string)

		response1, restyResp1, err := client.Reports.GetExecutionIDByReportID(vvReportID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetExecutionIDByReportID", err,
				"Failure at GetExecutionIDByReportID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenReportsGetExecutionIDByReportIDItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExecutionIDByReportID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetExecutionIDByReportIDItem(item *dnacentersdkgo.ResponseReportsGetExecutionIDByReportID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["report_id"] = item.ReportID
	respItem["report_name"] = item.ReportName
	respItem["executions"] = flattenReportsGetExecutionIDByReportIDItemExecutions(item.Executions)
	respItem["execution_count"] = item.ExecutionCount
	respItem["report_was_executed"] = boolPtrToString(item.ReportWasExecuted)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetExecutionIDByReportIDItemExecutions(items *[]dnacentersdkgo.ResponseReportsGetExecutionIDByReportIDExecutions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["execution_id"] = item.ExecutionID
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["process_status"] = item.ProcessStatus
		respItem["request_status"] = item.RequestStatus
		respItem["errors"] = item.Errors
		respItem["warnings"] = flattenReportsGetExecutionIDByReportIDItemExecutionsWarnings(item.Warnings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetExecutionIDByReportIDItemExecutionsWarnings(items *[]dnacentersdkgo.ResponseReportsGetExecutionIDByReportIDExecutionsWarnings) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
