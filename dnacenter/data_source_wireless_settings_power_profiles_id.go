package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsPowerProfilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get a Power Profile by Power Profile ID that captured in wireless settings design
`,

		ReadContext: dataSourceWirelessSettingsPowerProfilesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Power Profile ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `The description of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique Identifier of the power profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `The Name of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_id": &schema.Schema{
										Description: `Interface Id for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"interface_type": &schema.Schema{
										Description: `Interface Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_type": &schema.Schema{
										Description: `Parameter Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_value": &schema.Schema{
										Description: `Parameter Value for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sequence": &schema.Schema{
										Description: `Sequential Ordered List of rules for Power Profile.
`,
										Type:     schema.TypeInt,
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

func dataSourceWirelessSettingsPowerProfilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPowerProfileByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Wireless.GetPowerProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPowerProfileByID", err,
				"Failure at GetPowerProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetPowerProfileByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPowerProfileByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetPowerProfileByIDItem(item *dnacentersdkgo.ResponseWirelessGetPowerProfileByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["profile_name"] = item.ProfileName
	respItem["description"] = item.Description
	respItem["rules"] = flattenWirelessGetPowerProfileByIDItemRules(item.Rules)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetPowerProfileByIDItemRules(items *[]dnacentersdkgo.ResponseWirelessGetPowerProfileByIDResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sequence"] = item.Sequence
		respItem["interface_type"] = item.InterfaceType
		respItem["interface_id"] = item.InterfaceID
		respItem["parameter_type"] = item.ParameterType
		respItem["parameter_value"] = item.ParameterValue
		respItems = append(respItems, respItem)
	}
	return respItems
}
