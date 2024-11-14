package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexibleReportSchedules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get all flexible report schedules
`,

		ReadContext: dataSourceFlexibleReportSchedulesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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

						"schedule": &schema.Schema{
							Description: `Schedule information
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFlexibleReportSchedulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllFlexibleReportSchedules")

		response1, restyResp1, err := client.Reports.GetAllFlexibleReportSchedules()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllFlexibleReportSchedules", err,
				"Failure at GetAllFlexibleReportSchedules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenReportsGetAllFlexibleReportSchedulesItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllFlexibleReportSchedules response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetAllFlexibleReportSchedulesItems(items *dnacentersdkgo.ResponseReportsGetAllFlexibleReportSchedules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["report_id"] = item.ReportID
		respItem["schedule"] = flattenReportsGetAllFlexibleReportSchedulesItemsSchedule(item.Schedule)
		respItem["report_name"] = item.ReportName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetAllFlexibleReportSchedulesItemsSchedule(item *dnacentersdkgo.ResponseItemReportsGetAllFlexibleReportSchedulesSchedule) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
