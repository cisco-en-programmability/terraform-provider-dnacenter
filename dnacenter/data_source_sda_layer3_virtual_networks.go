package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaLayer3VirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of layer 3 virtual networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaLayer3VirtualNetworksRead,
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
			"virtual_network_name": &schema.Schema{
				Description: `virtualNetworkName query parameter. Name of the layer 3 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"anchored_site_id": &schema.Schema{
							Description: `Fabric ID of the fabric site this layer 3 virtual network is anchored at.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_ids": &schema.Schema{
							Description: `IDs of the fabrics this layer 3 virtual network is assigned to.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `ID of the layer 3 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Name of the layer 3 virtual network.
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

func dataSourceSdaLayer3VirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVirtualNetworkName, okVirtualNetworkName := d.GetOk("virtual_network_name")
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vAnchoredSiteID, okAnchoredSiteID := d.GetOk("anchored_site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer3VirtualNetworks")
		queryParams1 := dnacentersdkgo.GetLayer3VirtualNetworksQueryParams{}

		if okVirtualNetworkName {
			queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)
		}
		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okAnchoredSiteID {
			queryParams1.AnchoredSiteID = vAnchoredSiteID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetLayer3VirtualNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLayer3VirtualNetworks", err,
				"Failure at GetLayer3VirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetLayer3VirtualNetworksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer3VirtualNetworks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetLayer3VirtualNetworksItems(items *[]dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItem["fabric_ids"] = item.FabricIDs
		respItem["anchored_site_id"] = item.AnchoredSiteID
		respItems = append(respItems, respItem)
	}
	return respItems
}
