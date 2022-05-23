package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpSmartAccountDomains() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the list of Smart Account domains
`,

		ReadContext: dataSourcePnpSmartAccountDomainsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourcePnpSmartAccountDomainsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSmartAccountList")

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetSmartAccountList()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSmartAccountList", err,
				"Failure at GetSmartAccountList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetSmartAccountListItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSmartAccountList response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetSmartAccountListItems(items *dnacentersdkgo.ResponseDeviceOnboardingPnpGetSmartAccountList) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, respItem)
	}
	return respItems
}
