package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricsFabricIDSwitchWirelessSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Get the SDA Wireless details from the switches on the fabric site that have wireless capability enabled. A maximum of
two switches can have a wireless role in a fabric site.
`,

		ReadContext: dataSourceFabricsFabricIDSwitchWirelessSettingRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_wireless": &schema.Schema{
							Description: `Enable Wireless`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Network Device ID of the Wireless Capable Switch
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rolling_ap_upgrade": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_reboot_percentage": &schema.Schema{
										Description: `AP Reboot Percentage. Permissible values - 5, 15, 25
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"enable_rolling_ap_upgrade": &schema.Schema{
										Description: `Enable Rolling Ap Upgrade`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceFabricsFabricIDSwitchWirelessSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSdaWirelessDetailsFromSwitches")
		vvFabricID := vFabricID.(string)

		response1, restyResp1, err := client.FabricWireless.GetSdaWirelessDetailsFromSwitches(vvFabricID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSdaWirelessDetailsFromSwitches", err,
				"Failure at GetSdaWirelessDetailsFromSwitches, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaWirelessDetailsFromSwitches response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesItems(items *[]dnacentersdkgo.ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["enable_wireless"] = boolPtrToString(item.EnableWireless)
		respItem["rolling_ap_upgrade"] = flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesItemsRollingApUpgrade(item.RollingApUpgrade)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenFabricWirelessGetSdaWirelessDetailsFromSwitchesItemsRollingApUpgrade(item *dnacentersdkgo.ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponseRollingApUpgrade) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_rolling_ap_upgrade"] = boolPtrToString(item.EnableRollingApUpgrade)
	respItem["ap_reboot_percentage"] = item.ApRebootPercentage

	return []map[string]interface{}{
		respItem,
	}

}
