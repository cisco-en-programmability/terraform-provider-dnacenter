package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamSiteIPAddressPools() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves IP address subpools, which reserve address space from a global pool (or global pools).
`,

		ReadContext: dataSourceIPamSiteIPAddressPoolsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The *id* of the site for which to retrieve IP address subpools. Only subpools whose *siteId* exactly matches will be fetched, parent or child site matches will not be included.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
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
							Description: `The UUID for this reserve IP pool (subpool).
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_v4_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"assigned_addresses": &schema.Schema{
										Description: `The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"default_assigned_addresses": &schema.Schema{
										Description: `The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_servers": &schema.Schema{
										Description: `The DHCP server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"dns_servers": &schema.Schema{
										Description: `The DNS server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"gateway_ip_address": &schema.Schema{
										Description: `The gateway IP address for this subnet.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"global_pool_id": &schema.Schema{
										Description: `The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"slaac_support": &schema.Schema{
										Description: `If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"subnet": &schema.Schema{
										Description: `The IP address component of the CIDR notation for this subnet.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"total_addresses": &schema.Schema{
										Description: `The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"unassignable_addresses": &schema.Schema{
										Description: `The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"ip_v6_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"assigned_addresses": &schema.Schema{
										Description: `The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"default_assigned_addresses": &schema.Schema{
										Description: `The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_servers": &schema.Schema{
										Description: `The DHCP server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"dns_servers": &schema.Schema{
										Description: `The DNS server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"gateway_ip_address": &schema.Schema{
										Description: `The gateway IP address for this subnet.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"global_pool_id": &schema.Schema{
										Description: `The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"slaac_support": &schema.Schema{
										Description: `If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"subnet": &schema.Schema{
										Description: `The IP address component of the CIDR notation for this subnet.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"total_addresses": &schema.Schema{
										Description: `The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"unassignable_addresses": &schema.Schema{
										Description: `The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pool_type": &schema.Schema{
							Description: `Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_id": &schema.Schema{
							Description: `The *id* of the site that this subpool belongs to. This must be the *id* of a non-Global site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name": &schema.Schema{
							Description: `The name of the site that this subpool belongs to.
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

func dataSourceIPamSiteIPAddressPoolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesIPAddressSubpools")
		queryParams1 := dnacentersdkgo.RetrievesIPAddressSubpoolsQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrievesIPAddressSubpools(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesIPAddressSubpools", err,
				"Failure at RetrievesIPAddressSubpools, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsRetrievesIPAddressSubpoolsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesIPAddressSubpools response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsItems(items *[]dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ip_v4_address_space"] = flattenNetworkSettingsRetrievesIPAddressSubpoolsItemsIPV4AddressSpace(item.IPV4AddressSpace)
		respItem["ip_v6_address_space"] = flattenNetworkSettingsRetrievesIPAddressSubpoolsItemsIPV6AddressSpace(item.IPV6AddressSpace)
		respItem["name"] = item.Name
		respItem["pool_type"] = item.PoolType
		respItem["site_id"] = item.SiteID
		respItem["site_name"] = item.SiteName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsItemsIPV4AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponseIPV4AddressSpace) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["subnet"] = item.Subnet
	respItem["prefix_length"] = item.PrefixLength
	respItem["gateway_ip_address"] = item.GatewayIPAddress
	respItem["dhcp_servers"] = item.DhcpServers
	respItem["dns_servers"] = item.DNSServers
	respItem["total_addresses"] = item.TotalAddresses
	respItem["unassignable_addresses"] = item.UnassignableAddresses
	respItem["assigned_addresses"] = item.AssignedAddresses
	respItem["default_assigned_addresses"] = item.DefaultAssignedAddresses
	respItem["slaac_support"] = boolPtrToString(item.SLAacSupport)
	respItem["global_pool_id"] = item.GlobalPoolID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsItemsIPV6AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponseIPV6AddressSpace) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["subnet"] = item.Subnet
	respItem["prefix_length"] = item.PrefixLength
	respItem["gateway_ip_address"] = item.GatewayIPAddress
	respItem["dhcp_servers"] = item.DhcpServers
	respItem["dns_servers"] = item.DNSServers
	respItem["total_addresses"] = item.TotalAddresses
	respItem["unassignable_addresses"] = item.UnassignableAddresses
	respItem["assigned_addresses"] = item.AssignedAddresses
	respItem["default_assigned_addresses"] = item.DefaultAssignedAddresses
	respItem["slaac_support"] = boolPtrToString(item.SLAacSupport)
	respItem["global_pool_id"] = item.GlobalPoolID

	return []map[string]interface{}{
		respItem,
	}

}
