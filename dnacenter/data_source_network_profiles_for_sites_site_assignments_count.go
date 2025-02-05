package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkProfilesForSitesSiteAssignmentsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieves the count of sites that the given network profile for sites is assigned to.
`,

		ReadContext: dataSourceNetworkProfilesForSitesSiteAssignmentsCountRead,
		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Description: `profileId path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkProfilesForSitesSiteAssignmentsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileID := d.Get("profile_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo")
		vvProfileID := vProfileID.(string)

		response1, restyResp1, err := client.SiteDesign.RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo", err,
				"Failure at RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToItem(item *dnacentersdkgo.ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
