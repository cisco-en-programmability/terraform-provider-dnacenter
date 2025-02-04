package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamSiteIPAddressPoolsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves an IP address subpool, which reserves address space from a global pool (or global pools) for a particular
site.
`,

		ReadContext: dataSourceIPamSiteIPAddressPoolsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The *id* of the IP address subpool to retrieve.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceIPamSiteIPAddressPoolsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesAnIPAddressSubpool")
		vvID := vID.(string)

		response1, restyResp1, err := client.NetworkSettings.RetrievesAnIPAddressSubpool(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesAnIPAddressSubpool", err,
				"Failure at RetrievesAnIPAddressSubpool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrievesAnIPAddressSubpoolItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesAnIPAddressSubpool response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrievesAnIPAddressSubpoolItem(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesAnIPAddressSubpoolResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ip_v4_address_space"] = flattenNetworkSettingsRetrievesAnIPAddressSubpoolItemIPV4AddressSpace(item.IPV4AddressSpace)
	respItem["ip_v6_address_space"] = flattenNetworkSettingsRetrievesAnIPAddressSubpoolItemIPV6AddressSpace(item.IPV6AddressSpace)
	respItem["name"] = item.Name
	respItem["pool_type"] = item.PoolType
	respItem["site_id"] = item.SiteID
	respItem["site_name"] = item.SiteName
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrievesAnIPAddressSubpoolItemIPV4AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesAnIPAddressSubpoolResponseIPV4AddressSpace) []map[string]interface{} {
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

func flattenNetworkSettingsRetrievesAnIPAddressSubpoolItemIPV6AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesAnIPAddressSubpoolResponseIPV6AddressSpace) []map[string]interface{} {
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
