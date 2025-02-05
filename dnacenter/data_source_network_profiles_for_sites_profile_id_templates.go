package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkProfilesForSitesProfileIDTemplates() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves a list of CLI templates attached to a network profile based on the network profile ID.
`,

		ReadContext: dataSourceNetworkProfilesForSitesProfileIDTemplatesRead,
		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Description: `profileId path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The id of the template attached to the site profile - */intent/api/v1/templates*
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The name of the template attached to the site profile - */intent/api/v1/templates*
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

func dataSourceNetworkProfilesForSitesProfileIDTemplatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileID := d.Get("profile_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveCliTemplatesAttachedToANetworkProfile")
		vvProfileID := vProfileID.(string)

		response1, restyResp1, err := client.NetworkSettings.RetrieveCliTemplatesAttachedToANetworkProfile(vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveCliTemplatesAttachedToANetworkProfile", err,
				"Failure at RetrieveCliTemplatesAttachedToANetworkProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveCliTemplatesAttachedToANetworkProfile response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileItems(items *[]dnacentersdkgo.ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
