package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricDevicesLayer2HandoffsIPTransits() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.
`,

		ReadContext: dataSourceSdaFabricDevicesLayer2HandoffsIPTransitsRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric this device belongs to.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device ID of the fabric device.
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"external_connectivity_ip_pool_name": &schema.Schema{
							Description: `External connectivity ip pool is used by Catalyst Center to allocate IP address for the connection between the border node and peer.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the fabric device layer 3 handoff ip transit.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface name of the layer 3 handoff ip transit.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"local_ip_address": &schema.Schema{
							Description: `Local ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"local_ipv6_address": &schema.Schema{
							Description: `Local ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"remote_ip_address": &schema.Schema{
							Description: `Remote ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"remote_ipv6_address": &schema.Schema{
							Description: `Remote ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tcp_mss_adjustment": &schema.Schema{
							Description: `TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"transit_network_id": &schema.Schema{
							Description: `ID of the transit network of the layer 3 handoff ip transit.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_network_name": &schema.Schema{
							Description: `Name of the virtual network associated with this fabric site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
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

func dataSourceSdaFabricDevicesLayer2HandoffsIPTransitsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevicesLayer3HandoffsWithIPTransit")
		queryParams1 := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams{}

		queryParams1.FabricID = vFabricID.(string)

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetFabricDevicesLayer3HandoffsWithIPTransit(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFabricDevicesLayer3HandoffsWithIPTransit", err,
				"Failure at GetFabricDevicesLayer3HandoffsWithIPTransit, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetFabricDevicesLayer3HandoffsWithIPTransitItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevicesLayer3HandoffsWithIPTransit response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetFabricDevicesLayer3HandoffsWithIPTransitItems(items *[]dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["fabric_id"] = item.FabricID
		respItem["transit_network_id"] = item.TransitNetworkID
		respItem["interface_name"] = item.InterfaceName
		respItem["external_connectivity_ip_pool_name"] = item.ExternalConnectivityIPPoolName
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItem["vlan_id"] = item.VLANID
		respItem["tcp_mss_adjustment"] = item.TCPMssAdjustment
		respItem["local_ip_address"] = item.LocalIPAddress
		respItem["remote_ip_address"] = item.RemoteIPAddress
		respItem["local_ipv6_address"] = item.LocalIPv6Address
		respItem["remote_ipv6_address"] = item.RemoteIPv6Address
		respItems = append(respItems, respItem)
	}
	return respItems
}
