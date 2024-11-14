package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaAuthenticationProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of authentication profiles that match the provided query parameters.
`,

		ReadContext: dataSourceSdaAuthenticationProfilesRead,
		Schema: map[string]*schema.Schema{
			"authentication_profile_name": &schema.Schema{
				Description: `authenticationProfileName query parameter. Return only the authentication profiles with this specified name. Note that 'No Authentication' is not a valid option for this parameter.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the authentication profile is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authentication_order": &schema.Schema{
							Description: `First authentication method.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"authentication_profile_name": &schema.Schema{
							Description: `The default host authentication template.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `802.1x Timeout.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this authentication profile is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the authentication profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_bpdu_guard_enabled": &schema.Schema{
							Description: `Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication".
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"number_of_hosts": &schema.Schema{
							Description: `Number of Hosts.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wake_on_lan": &schema.Schema{
							Description: `Wake on LAN.
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

func dataSourceSdaAuthenticationProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vAuthenticationProfileName, okAuthenticationProfileName := d.GetOk("authentication_profile_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAuthenticationProfiles")
		queryParams1 := dnacentersdkgo.GetAuthenticationProfilesQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okAuthenticationProfileName {
			queryParams1.AuthenticationProfileName = vAuthenticationProfileName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetAuthenticationProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAuthenticationProfiles", err,
				"Failure at GetAuthenticationProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetAuthenticationProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthenticationProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetAuthenticationProfilesItems(items *[]dnacentersdkgo.ResponseSdaGetAuthenticationProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["authentication_profile_name"] = item.AuthenticationProfileName
		respItem["authentication_order"] = item.AuthenticationOrder
		respItem["dot1x_to_mab_fallback_timeout"] = item.Dot1XToMabFallbackTimeout
		respItem["wake_on_lan"] = boolPtrToString(item.WakeOnLan)
		respItem["number_of_hosts"] = item.NumberOfHosts
		respItem["is_bpdu_guard_enabled"] = boolPtrToString(item.IsBpduGuardEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}
