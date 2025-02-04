package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMapsSupportedAccessPoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.
`,

		ReadContext: dataSourceMapsSupportedAccessPointsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"antenna_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"band": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"names": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"ap_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMapsSupportedAccessPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: MapsSupportedAccessPoints")

		response1, restyResp1, err := client.Sites.MapsSupportedAccessPoints()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 MapsSupportedAccessPoints", err,
				"Failure at MapsSupportedAccessPoints, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesMapsSupportedAccessPointsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting MapsSupportedAccessPoints response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesMapsSupportedAccessPointsItems(items *dnacentersdkgo.ResponseSitesMapsSupportedAccessPoints) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["antenna_patterns"] = flattenSitesMapsSupportedAccessPointsItemsAntennaPatterns(item.AntennaPatterns)
		respItem["ap_type"] = item.ApType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesMapsSupportedAccessPointsItemsAntennaPatterns(items *[]dnacentersdkgo.ResponseItemSitesMapsSupportedAccessPointsAntennaPatterns) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["band"] = item.Band
		respItem["names"] = item.Names
		respItems = append(respItems, respItem)
	}
	return respItems
}
