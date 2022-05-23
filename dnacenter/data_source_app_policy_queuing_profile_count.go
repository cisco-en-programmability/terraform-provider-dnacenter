package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppPolicyQueuingProfileCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get the number of all existing  application policy queuing profile
`,

		ReadContext: dataSourceAppPolicyQueuingProfileCountRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Total number of Queueing Profile
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAppPolicyQueuingProfileCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicyQueuingProfileCount")

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicyQueuingProfileCount()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationPolicyQueuingProfileCount", err,
				"Failure at GetApplicationPolicyQueuingProfileCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyGetApplicationPolicyQueuingProfileCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyQueuingProfileCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileCountItem(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
