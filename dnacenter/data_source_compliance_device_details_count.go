package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceDetailsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return  Compliance Count Detail
`,

		ReadContext: dataSourceComplianceDeviceDetailsCountRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. complianceType can have any value among 'NETWORK_PROFILE', 'IMAGE', 'APPLICATION_VISIBILITY', 'FABRIC', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
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

func dataSourceComplianceDeviceDetailsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceType, okComplianceType := d.GetOk("compliance_type")
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetComplianceDetailCount")
		queryParams1 := dnacentersdkgo.GetComplianceDetailCountQueryParams{}

		if okComplianceType {
			queryParams1.ComplianceType = vComplianceType.(string)
		}
		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceDetailCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetComplianceDetailCount", err,
				"Failure at GetComplianceDetailCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetComplianceDetailCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceDetailCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceDetailCountItem(item *dnacentersdkgo.ResponseComplianceGetComplianceDetailCount) []map[string]interface{} {
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
