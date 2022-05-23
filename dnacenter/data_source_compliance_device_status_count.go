package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceStatusCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return Compliance Status Count
`,

		ReadContext: dataSourceComplianceDeviceStatusCountRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
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
							Description: `Response`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceComplianceDeviceStatusCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetComplianceStatusCount")
		queryParams1 := dnacentersdkgo.GetComplianceStatusCountQueryParams{}

		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceStatusCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetComplianceStatusCount", err,
				"Failure at GetComplianceStatusCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetComplianceStatusCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceStatusCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceStatusCountItem(item *dnacentersdkgo.ResponseComplianceGetComplianceStatusCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
