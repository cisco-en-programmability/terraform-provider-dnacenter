package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceQosPolicySetting() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- API to retrieve the application QoS policy setting.
`,

		ReadContext: dataSourceQosPolicySettingRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deploy_by_default_on_wired_devices": &schema.Schema{
							Description: `Flag to indicate whether QoS policy should be deployed automatically on wired network device when it is provisioned. This would be only applicable for cases where the network device is assigned to a site where a QoS policy has been configured.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceQosPolicySettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheApplicationQoSPolicySetting")

		response1, restyResp1, err := client.ApplicationPolicy.RetrievesTheApplicationQoSPolicySetting()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheApplicationQoSPolicySetting", err,
				"Failure at RetrievesTheApplicationQoSPolicySetting, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyRetrievesTheApplicationQoSPolicySettingItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheApplicationQoSPolicySetting response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyRetrievesTheApplicationQoSPolicySettingItem(item *dnacentersdkgo.ResponseApplicationPolicyRetrievesTheApplicationQoSPolicySettingResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deploy_by_default_on_wired_devices"] = boolPtrToString(item.DeployByDefaultOnWiredDevices)
	return []map[string]interface{}{
		respItem,
	}
}
