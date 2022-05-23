package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpVirtualAccounts() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns list of virtual accounts associated with the specified smart account
`,

		ReadContext: dataSourcePnpVirtualAccountsRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: `domain path parameter. Smart Account Domain
`,
				Type:     schema.TypeString,
				Required: true,
			},

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

func dataSourcePnpVirtualAccountsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDomain := d.Get("domain")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVirtualAccountList")
		vvDomain := vDomain.(string)

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetVirtualAccountList(vvDomain)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualAccountList", err,
				"Failure at GetVirtualAccountList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetVirtualAccountListItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualAccountList response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetVirtualAccountListItems(items *dnacentersdkgo.ResponseDeviceOnboardingPnpGetVirtualAccountList) []interface{} {
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
