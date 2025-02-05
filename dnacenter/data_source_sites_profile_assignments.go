package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesProfileAssignments() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieves the list of profiles that the given site has been assigned.  These profiles may either be directly assigned
to this site, or were assigned to a parent site and have been inherited.
These assigments can be modified via the */dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments*
resources.
`,

		ReadContext: dataSourceSitesProfileAssignmentsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. The *id* of the site, retrievable from */dna/intent/api/v1/sites*
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
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSitesProfileAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned")
		vvSiteID := vSiteID.(string)
		queryParams1 := dnacentersdkgo.RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.SiteDesign.RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned", err,
				"Failure at RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedItems(items *[]dnacentersdkgo.ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
