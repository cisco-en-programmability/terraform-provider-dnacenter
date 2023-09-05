package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBusinessSdaVirtualNetworkSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

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
							Description: `Virtual Network summary retrieved successfully from SDA Fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_count": &schema.Schema{
							Description: `Virtual Networks Count
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"virtual_network_summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"site_name_hierarchy": &schema.Schema{
										Description: `Site Name Hierarchy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"virtual_network_name": &schema.Schema{
										Description: `Virtual Network Name
`,
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

func dataSourceBusinessSdaVirtualNetworkSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteNameHierarchy := d.Get("site_name_hierarchy")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetVirtualNetworkSummary")
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

		vItem1 := flattenSdaGetVirtualNetworkSummaryItem(response1)
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

func flattenSdaGetVirtualNetworkSummaryItem(item *dnacentersdkgo.ResponseSdaGetVirtualNetworkSummary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["virtual_network_count"] = item.VirtualNetworkCount
	respItem["virtual_network_summary"] = flattenSdaGetVirtualNetworkSummaryItemVirtualNetworkSummary(item.VirtualNetworkSummary)
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetVirtualNetworkSummaryItemVirtualNetworkSummary(items *[]dnacentersdkgo.ResponseSdaGetVirtualNetworkSummaryVirtualNetworkSummary) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItems = append(respItems, respItem)
	}
	return respItems
}
