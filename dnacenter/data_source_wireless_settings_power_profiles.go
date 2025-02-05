package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsPowerProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get Power Profiles that captured in wireless settings design.
`,

		ReadContext: dataSourceWirelessSettingsPowerProfilesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"profile_name": &schema.Schema{
				Description: `profileName query parameter. Power Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
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
										Description: `The sequence of the power profile rule.
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

func dataSourceWirelessSettingsPowerProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vProfileName, okProfileName := d.GetOk("profile_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPowerProfiles")
		queryParams1 := dnacentersdkgo.GetPowerProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okProfileName {
			queryParams1.ProfileName = vProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetPowerProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPowerProfiles", err,
				"Failure at GetPowerProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetPowerProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPowerProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetPowerProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGetPowerProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["profile_name"] = item.ProfileName
		respItem["description"] = item.Description
		respItem["rules"] = flattenWirelessGetPowerProfilesItemsRules(item.Rules)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetPowerProfilesItemsRules(items *[]dnacentersdkgo.ResponseWirelessGetPowerProfilesResponseRules) []map[string]interface{} {
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
