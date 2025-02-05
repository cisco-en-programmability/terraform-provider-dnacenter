package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBuildings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Gets a building in the network hierarchy.
`,

		ReadContext: dataSourceBuildingsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Building Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"address": &schema.Schema{
							Description: `Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"country": &schema.Schema{
							Description: `Country name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"latitude": &schema.Schema{
							Description: `Building Latitude. Example: 37.403712
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"longitude": &schema.Schema{
							Description: `Building Longitude. Example: -121.971063
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Building name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Example: building
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

func dataSourceBuildingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsABuildingV2")
		vvID := vID.(string)

		response1, restyResp1, err := client.SiteDesign.GetsABuildingV2(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsABuildingV2", err,
				"Failure at GetsABuildingV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsABuildingV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsABuildingV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetsABuildingV2Item(item *dnacentersdkgo.ResponseSiteDesignGetsABuildingV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_id"] = item.ParentID
	respItem["name"] = item.Name
	respItem["latitude"] = item.Latitude
	respItem["longitude"] = item.Longitude
	respItem["address"] = item.Address
	respItem["country"] = item.Country
	respItem["type"] = item.Type
	return []map[string]interface{}{
		respItem,
	}
}
