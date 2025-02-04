package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsTrend() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get field notices results trend over time. The default sort is by scan time descending.
`,

		ReadContext: dataSourceFieldNoticesResultsTrendRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Default value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"scan_time": &schema.Schema{
				Description: `scanTime query parameter. Return field notices trend with scanTime greater than this scanTime
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hardware_field_notices_count": &schema.Schema{
							Description: `Number of field notices of type HARDWARE
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"scan_time": &schema.Schema{
							Description: `End time for the scan
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"software_field_notices_count": &schema.Schema{
							Description: `Number of field notices of type SOFTWARE
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFieldNoticesResultsTrendRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vScanTime, okScanTime := d.GetOk("scan_time")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFieldNoticesResultsTrendOverTime")
		queryParams1 := dnacentersdkgo.GetFieldNoticesResultsTrendOverTimeQueryParams{}

		if okScanTime {
			queryParams1.ScanTime = vScanTime.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Compliance.GetFieldNoticesResultsTrendOverTime(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFieldNoticesResultsTrendOverTime", err,
				"Failure at GetFieldNoticesResultsTrendOverTime, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetFieldNoticesResultsTrendOverTimeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFieldNoticesResultsTrendOverTime response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetFieldNoticesResultsTrendOverTimeItems(items *[]dnacentersdkgo.ResponseComplianceGetFieldNoticesResultsTrendOverTimeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["scan_time"] = item.ScanTime
		respItem["software_field_notices_count"] = item.SoftwareFieldNoticesCount
		respItem["hardware_field_notices_count"] = item.HardwareFieldNoticesCount
		respItems = append(respItems, respItem)
	}
	return respItems
}
