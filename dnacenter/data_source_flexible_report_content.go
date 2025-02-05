package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexibleReportContent() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- This is used to download the flexible report. The API returns report content. Save the response to a file by
converting the response data as a blob and setting the file format available from content-disposition response header.
`,

		ReadContext: dataSourceFlexibleReportContentRead,
		Schema: map[string]*schema.Schema{
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. Id of execution
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"report_id": &schema.Schema{
				Description: `reportId path parameter. Id of the report
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFlexibleReportContentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")
	vExecutionID := d.Get("execution_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DownloadFlexibleReport")
		vvReportID := vReportID.(string)
		vvExecutionID := vExecutionID.(string)
		response1, err := client.Reports.DownloadFlexibleReport(vvReportID, vvExecutionID)

		if err = d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DownloadFlexibleReport response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
