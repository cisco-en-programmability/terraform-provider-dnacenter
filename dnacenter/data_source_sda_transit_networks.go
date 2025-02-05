package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaTransitNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of transit networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaTransitNetworksRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. ID of the transit network.
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
			"name": &schema.Schema{
				Description: `name query parameter. Name of the transit network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID of the transit network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autonomous_system_number": &schema.Schema{
										Description: `Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295].
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"routing_protocol_name": &schema.Schema{
										Description: `Routing Protocol Name of the IP transit network.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `Name of the transit network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sda_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_plane_network_device_ids": &schema.Schema{
										Description: `List of network device IDs that are used as control plane nodes.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"is_multicast_over_transit_enabled": &schema.Schema{
										Description: `This indicates that multicast is enabled over SD-Access Transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"type": &schema.Schema{
							Description: `Type of the transit network.
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

func dataSourceSdaTransitNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransitNetworks")
		queryParams1 := dnacentersdkgo.GetTransitNetworksQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetTransitNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTransitNetworks", err,
				"Failure at GetTransitNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetTransitNetworksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransitNetworks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetTransitNetworksItems(items *[]dnacentersdkgo.ResponseSdaGetTransitNetworksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["ip_transit_settings"] = flattenSdaGetTransitNetworksItemsIPTransitSettings(item.IPTransitSettings)
		respItem["sda_transit_settings"] = flattenSdaGetTransitNetworksItemsSdaTransitSettings(item.SdaTransitSettings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetTransitNetworksItemsIPTransitSettings(item *dnacentersdkgo.ResponseSdaGetTransitNetworksResponseIPTransitSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["routing_protocol_name"] = item.RoutingProtocolName
	respItem["autonomous_system_number"] = item.AutonomousSystemNumber

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetTransitNetworksItemsSdaTransitSettings(item *dnacentersdkgo.ResponseSdaGetTransitNetworksResponseSdaTransitSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["is_multicast_over_transit_enabled"] = boolPtrToString(item.IsMulticastOverTransitEnabled)
	respItem["control_plane_network_device_ids"] = item.ControlPlaneNetworkDeviceIDs

	return []map[string]interface{}{
		respItem,
	}

}
