package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaLayer3VirtualNetworksCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of layer 3 virtual networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaLayer3VirtualNetworksCountRead,
		Schema: map[string]*schema.Schema{
			"anchored_site_id": &schema.Schema{
				Description: `anchoredSiteId query parameter. Fabric ID of the fabric site the layer 3 virtual network is anchored at.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the layer 3 virtual network is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Number of layer 3 virtual networks.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaLayer3VirtualNetworksCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vAnchoredSiteID, okAnchoredSiteID := d.GetOk("anchored_site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer3VirtualNetworksCount")
		queryParams1 := dnacentersdkgo.GetLayer3VirtualNetworksCountQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okAnchoredSiteID {
			queryParams1.AnchoredSiteID = vAnchoredSiteID.(string)
		}

		response1, restyResp1, err := client.Sda.GetLayer3VirtualNetworksCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLayer3VirtualNetworksCount", err,
				"Failure at GetLayer3VirtualNetworksCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetLayer3VirtualNetworksCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer3VirtualNetworksCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetLayer3VirtualNetworksCountItem(item *dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworksCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
