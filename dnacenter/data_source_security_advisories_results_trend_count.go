package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsTrendCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of security advisories results trend over time.
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsTrendCountRead,
		Schema: map[string]*schema.Schema{
			"scan_time": &schema.Schema{
				Description: `scanTime query parameter. Return advisories trend with scanTime greater than this scanTime
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityAdvisoriesResultsTrendCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vScanTime, okScanTime := d.GetOk("scan_time")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfSecurityAdvisoriesResultsTrendOverTime")
		queryParams1 := dnacentersdkgo.GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams{}

		if okScanTime {
			queryParams1.ScanTime = vScanTime.(float64)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfSecurityAdvisoriesResultsTrendOverTime(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfSecurityAdvisoriesResultsTrendOverTime", err,
				"Failure at GetCountOfSecurityAdvisoriesResultsTrendOverTime, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfSecurityAdvisoriesResultsTrendOverTime response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeItem(item *dnacentersdkgo.ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
