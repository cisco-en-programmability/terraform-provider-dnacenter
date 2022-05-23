package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePlatformReleaseSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Platform Configuration.

- Provides information such as API version, mandatory core packages for installation or upgrade, optional packages,
Cisco DNA Center name and version, supported direct updates, and tenant ID.
`,

		ReadContext: dataSourcePlatformReleaseSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"core_packages": &schema.Schema{
							Description: `The set of packages that are mandatory to be installed/upgraded with the release
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"installed_version": &schema.Schema{
							Description: `The installed Cisco DNAC version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the release (example "dnac")
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"packages": &schema.Schema{
							Description: `The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"supported_direct_updates": &schema.Schema{
							Description: `The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"system_version": &schema.Schema{
							Description: `The MAGLEV-SYSTEM version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant ID (for multi tenant Cisco DNA Center)
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

func dataSourcePlatformReleaseSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CiscoDnaCenterReleaseSummary")

		response1, restyResp1, err := client.PlatformConfiguration.CiscoDnaCenterReleaseSummary()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CiscoDnaCenterReleaseSummary", err,
				"Failure at CiscoDnaCenterReleaseSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPlatformConfigurationCiscoDnaCenterReleaseSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CiscoDnaCenterReleaseSummary response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPlatformConfigurationCiscoDnaCenterReleaseSummaryItem(item *dnacentersdkgo.ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["core_packages"] = item.CorePackages
	respItem["packages"] = item.Packages
	respItem["name"] = item.Name
	respItem["installed_version"] = item.InstalledVersion
	respItem["system_version"] = item.SystemVersion
	respItem["supported_direct_updates"] = flattenPlatformConfigurationCiscoDnaCenterReleaseSummaryItemSupportedDirectUpdates(item.SupportedDirectUpdates)
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPlatformConfigurationCiscoDnaCenterReleaseSummaryItemSupportedDirectUpdates(items *[]dnacentersdkgo.ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponseSupportedDirectUpdates) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
