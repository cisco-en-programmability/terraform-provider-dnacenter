package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEoxStatusSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on EoX.

- Retrieves EoX summary for all devices in the network
`,

		ReadContext: dataSourceEoxStatusSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hardware_count": &schema.Schema{
							Description: `Hardware Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"module_count": &schema.Schema{
							Description: `Module Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"software_count": &schema.Schema{
							Description: `Software Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_count": &schema.Schema{
							Description: `Total Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEoxStatusSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEoXSummary")

		response1, restyResp1, err := client.EoX.GetEoXSummary()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEoXSummary", err,
				"Failure at GetEoXSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEoXGetEoXSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEoXSummary response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEoXGetEoXSummaryItem(item *dnacentersdkgo.ResponseEoXGetEoXSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hardware_count"] = item.HardwareCount
	respItem["software_count"] = item.SoftwareCount
	respItem["module_count"] = item.ModuleCount
	respItem["total_count"] = item.TotalCount
	return []map[string]interface{}{
		respItem,
	}
}
