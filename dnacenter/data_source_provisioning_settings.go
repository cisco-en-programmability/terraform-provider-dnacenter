package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProvisioningSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on System Settings.

- Returns provisioning settings
`,

		ReadContext: dataSourceProvisioningSettingsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"require_itsm_approval": &schema.Schema{
							Description: `If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"require_preview": &schema.Schema{
							Description: `If require preview is enabled, the device configurations must be reviewed before deploying them
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

func dataSourceProvisioningSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProvisioningSettings")

		response1, restyResp1, err := client.SystemSettings.GetProvisioningSettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetProvisioningSettings", err,
				"Failure at GetProvisioningSettings, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsGetProvisioningSettingsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisioningSettings response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSystemSettingsGetProvisioningSettingsItem(item *dnacentersdkgo.ResponseSystemSettingsGetProvisioningSettingsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["require_itsm_approval"] = boolPtrToString(item.RequireItsmApproval)
	respItem["require_preview"] = boolPtrToString(item.RequirePreview)
	return []map[string]interface{}{
		respItem,
	}
}
