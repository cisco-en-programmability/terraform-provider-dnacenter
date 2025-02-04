package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaMulticastVirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of multicast configurations for virtual networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaMulticastVirtualNetworksRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric site where multicast is configured.
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
				Description: `virtualNetworkName query parameter. Name of the virtual network associated to the multicast configuration.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the multicast configuration.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Name of the IP Pool.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_ssm_ranges": &schema.Schema{
							Description: `IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"multicast_r_ps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipv4_address": &schema.Schema{
										Description: `IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ipv4_asm_ranges": &schema.Schema{
										Description: `IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ipv6_address": &schema.Schema{
										Description: `IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ipv6_asm_ranges": &schema.Schema{
										Description: `IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"is_default_v4_rp": &schema.Schema{
										Description: `Specifies whether it is a default IPv4 RP.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_default_v6_rp": &schema.Schema{
										Description: `Specifies whether it is a default IPv6 RP.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"network_device_ids": &schema.Schema{
										Description: `IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"rp_device_location": &schema.Schema{
										Description: `Device location of the RP.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"virtual_network_name": &schema.Schema{
							Description: `Name of the virtual network.
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

func dataSourceSdaMulticastVirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vVirtualNetworkName, okVirtualNetworkName := d.GetOk("virtual_network_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMulticastVirtualNetworks")
		queryParams1 := dnacentersdkgo.GetMulticastVirtualNetworksQueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okVirtualNetworkName {
			queryParams1.VirtualNetworkName = vVirtualNetworkName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetMulticastVirtualNetworks(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetMulticastVirtualNetworks", err,
				"Failure at GetMulticastVirtualNetworks, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetMulticastVirtualNetworksItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMulticastVirtualNetworks response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetMulticastVirtualNetworksItems(items *[]dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworksResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["virtual_network_name"] = item.VirtualNetworkName
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["ipv4_ssm_ranges"] = item.IPv4SsmRanges
		respItem["multicast_r_ps"] = flattenSdaGetMulticastVirtualNetworksItemsMulticastRPs(item.MulticastRPs)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetMulticastVirtualNetworksItemsMulticastRPs(items *[]dnacentersdkgo.ResponseSdaGetMulticastVirtualNetworksResponseMulticastRPs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["rp_device_location"] = item.RpDeviceLocation
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_address"] = item.IPv6Address
		respItem["is_default_v4_rp"] = boolPtrToString(item.IsDefaultV4RP)
		respItem["is_default_v6_rp"] = boolPtrToString(item.IsDefaultV6RP)
		respItem["network_device_ids"] = item.NetworkDeviceIDs
		respItem["ipv4_asm_ranges"] = item.IPv4AsmRanges
		respItem["ipv6_asm_ranges"] = item.IPv6AsmRanges
		respItems = append(respItems, respItem)
	}
	return respItems
}
