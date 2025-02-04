package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexibleReportSchedule() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get flexible report schedule by report id
`,

		ReadContext: dataSourceFlexibleReportScheduleRead,
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceFlexibleReportScheduleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFlexibleReportScheduleByReportID")
		vvReportID := vReportID.(string)

		response1, restyResp1, err := client.Reports.GetFlexibleReportScheduleByReportID(vvReportID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFlexibleReportScheduleByReportID", err,
				"Failure at GetFlexibleReportScheduleByReportID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := restyResp1.String()
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFlexibleReportScheduleByReportID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
