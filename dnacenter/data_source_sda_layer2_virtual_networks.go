package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaLayer2VirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of layer 2 virtual networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaLayer2VirtualNetworksRead,
		Schema: map[string]*schema.Schema{
			"associated_layer3_virtual_network_name": &schema.Schema{
				Description: `associatedLayer3VirtualNetworkName query parameter. Name of the associated layer 3 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the layer 2 virtual network is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. ID of the layer 2 virtual network.
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
			"traffic_type": &schema.Schema{
				Description: `trafficType query parameter. The traffic type of the layer 2 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Description: `vlanId query parameter. The vlan ID of the layer 2 virtual network.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Description: `vlanName query parameter. The vlan name of the layer 2 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"associated_layer3_virtual_network_name": &schema.Schema{
							Description: `Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this layer 2 virtual network is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the layer 2 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_fabric_enabled_wireless": &schema.Schema{
							Description: `Set to true to enable wireless.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_multiple_ip_to_mac_addresses": &schema.Schema{
							Description: `Set to true to enable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine). This field will only be present on layer 2 virtual networks associated with a layer 3 virtual network.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"traffic_type": &schema.Schema{
							Description: `The type of traffic that is served.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `ID of the VLAN of the layer 2 virtual network.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"vlan_name": &schema.Schema{
							Description: `Name of the VLAN of the layer 2 virtual network.
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

func dataSourceSdaLayer2VirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vVLANName, okVLANName := d.GetOk("vlan_name")
	vVLANID, okVLANID := d.GetOk("vlan_id")
	vTrafficType, okTrafficType := d.GetOk("traffic_type")
	vAssociatedLayer3VirtualNetworkName, okAssociatedLayer3VirtualNetworkName := d.GetOk("associated_layer3_virtual_network_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer2VirtualNetworks")
		queryParams1 := dnacentersdkgo.GetLayer2VirtualNetworksQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okVLANName {
			queryParams1.VLANName = vVLANName.(string)
		}
		if okVLANID {
			queryParams1.VLANID = vVLANID.(float64)
		}
		if okTrafficType {
			queryParams1.TrafficType = vTrafficType.(string)
		}
		if okAssociatedLayer3VirtualNetworkName {
			queryParams1.AssociatedLayer3VirtualNetworkName = vAssociatedLayer3VirtualNetworkName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetLayer2VirtualNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLayer2VirtualNetworks", err,
				"Failure at GetLayer2VirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetLayer2VirtualNetworksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer2VirtualNetworks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetLayer2VirtualNetworksItems(items *[]dnacentersdkgo.ResponseSdaGetLayer2VirtualNetworksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["vlan_name"] = item.VLANName
		respItem["vlan_id"] = item.VLANID
		respItem["traffic_type"] = item.TrafficType
		respItem["is_fabric_enabled_wireless"] = boolPtrToString(item.IsFabricEnabledWireless)
		respItem["is_multiple_ip_to_mac_addresses"] = boolPtrToString(item.IsMultipleIPToMacAddresses)
		respItem["associated_layer3_virtual_network_name"] = item.AssociatedLayer3VirtualNetworkName
		respItems = append(respItems, respItem)
	}
	return respItems
}
