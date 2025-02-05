package dnacenter

import (
	"context"

	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceFlexibleReportExecute() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Reports.

- This data source action is used for executing the report
`,

		CreateContext: resourceFlexibleReportExecuteCreate,
		ReadContext:   resourceFlexibleReportExecuteRead,
		DeleteContext: resourceFlexibleReportExecuteDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
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
							Description: `Errors associated to the report execution
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
						"request_status": &schema.Schema{
							Description: `Report  request status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Description: `Report execution start time (Represent the specified number of milliseconds since the epoch time)
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"warnings": &schema.Schema{
							Description: `Warnings associated to the report execution
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"report_id": &schema.Schema{
							Description: `reportId path parameter. Id of the Report
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceFlexibleReportExecuteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vReportID := resourceItem["report_id"]

	vvReportID := vReportID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Reports.ExecutingTheFlexibleReport(vvReportID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ExecutingTheFlexibleReport", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ExecutingTheFlexibleReport", err,
				"Failure at ExecutingTheFlexibleReport execution", bapiError))
			return diags
		}
	}
	vItem1 := flattenReportsExecutingTheFlexibleReportItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ExecutingTheFlexibleReport response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceFlexibleReportExecuteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFlexibleReportExecuteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenReportsExecutingTheFlexibleReportItem(item *dnacentersdkgo.ResponseReportsExecutingTheFlexibleReport) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["request_status"] = item.RequestStatus
	respItem["errors"] = item.Errors
	respItem["warnings"] = flattenReportsExecutingTheFlexibleReportItemWarnings(item.Warnings)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsExecutingTheFlexibleReportItemWarnings(items *[]dnacentersdkgo.ResponseReportsExecutingTheFlexibleReportWarnings) []interface{} {
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
