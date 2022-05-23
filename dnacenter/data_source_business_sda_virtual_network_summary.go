package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBusinessSdaVirtualNetworkSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation.

- Get Virtual Network Summary
`,

		ReadContext: dataSourceBusinessSdaVirtualNetworkSummaryRead,
		Schema: map[string]*schema.Schema{
			"site_name_hierarchy": &schema.Schema{
				Description: `siteNameHierarchy query parameter. Complete fabric siteNameHierarchy Path
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
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"fabric_count": &schema.Schema{
							Description: `Fabric Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBusinessSdaVirtualNetworkSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVirtualNetworkSummary")
		queryParams1 := dnacentersdkgo.GetVirtualNetworkSummaryQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy.(string)

		response1, restyResp1, err := client.Sda.GetVirtualNetworkSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetVirtualNetworkSummary", err,
				"Failure at GetVirtualNetworkSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenGetVirtualNetworkSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkSummary response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenGetVirtualNetworkSummaryItem(item *dnacentersdkgo.ResponseSdaGetVirtualNetworkSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["fabric_count"] = item.FabricCount
	return []map[string]interface{}{
		respItem,
	}
}
