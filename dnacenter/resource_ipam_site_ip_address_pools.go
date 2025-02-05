package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPamSiteIPAddressPools() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Network Settings.

- Reserves (creates) an IP address subpool, which reserves address space from a global pool (or global pools) for a
particular site (and it's child sites). A subpool must be either an IPv4 or dual-stack pool, with *ipV4AddressSpace* and
optionally *ipV6AddressSpace* properties specified.
`,

		CreateContext: resourceIPamSiteIPAddressPoolsCreate,
		ReadContext:   resourceIPamSiteIPAddressPoolsRead,
		UpdateContext: resourceIPamSiteIPAddressPoolsUpdate,
		DeleteContext: resourceIPamSiteIPAddressPoolsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ip_v4_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"assigned_addresses": &schema.Schema{
										Description: `The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_assigned_addresses": &schema.Schema{
										Description: `The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dhcp_servers": &schema.Schema{
										Description: `The DHCP server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dns_servers": &schema.Schema{
										Description: `The DNS server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"gateway_ip_address": &schema.Schema{
										Description: `The gateway IP address for this subnet.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"global_pool_id": &schema.Schema{
										Description: `The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"slaac_support": &schema.Schema{
										Description: `If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"subnet": &schema.Schema{
										Description: `The IP address component of the CIDR notation for this subnet.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"total_addresses": &schema.Schema{
										Description: `The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unassignable_addresses": &schema.Schema{
										Description: `The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ip_v6_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"assigned_addresses": &schema.Schema{
										Description: `The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_assigned_addresses": &schema.Schema{
										Description: `The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dhcp_servers": &schema.Schema{
										Description: `The DHCP server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dns_servers": &schema.Schema{
										Description: `The DNS server(s) for this subnet.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"gateway_ip_address": &schema.Schema{
										Description: `The gateway IP address for this subnet.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"global_pool_id": &schema.Schema{
										Description: `The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"slaac_support": &schema.Schema{
										Description: `If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"subnet": &schema.Schema{
										Description: `The IP address component of the CIDR notation for this subnet.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"total_addresses": &schema.Schema{
										Description: `The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"unassignable_addresses": &schema.Schema{
										Description: `The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pool_type": &schema.Schema{
							Description: `Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `The *id* of the site that this subpool belongs to. This must be the *id* of a non-Global site.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name": &schema.Schema{
							Description: `The name of the site that this subpool belongs to.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIPamSiteIPAddressPoolsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpools(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vSiteID := resourceItem["site_id"]
	vvSiteID := vSiteID.(string)

	queryParamImport := dnacentersdkgo.RetrievesIPAddressSubpoolsQueryParams{}
	queryParamImport.SiteID = vvSiteID
	item2, err := searchNetworkSettingsRetrievesIPAddressSubpools(m, queryParamImport, vvSiteID)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceItem["site_id"] = item2.SiteID
		d.SetId(joinResourceID(resourceMap))
		return resourceIPamSiteIPAddressPoolsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.NetworkSettings.ReservecreateIPAddressSubpools(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ReservecreateIPAddressSubpools", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ReservecreateIPAddressSubpools", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ReservecreateIPAddressSubpools", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ReservecreateIPAddressSubpools", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.RetrievesIPAddressSubpoolsQueryParams{}
	queryParamValidate.SiteID = vvSiteID
	item3, err := searchNetworkSettingsRetrievesIPAddressSubpools(m, queryParamValidate, vvSiteID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ReservecreateIPAddressSubpools", err,
			"Failure at ReservecreateIPAddressSubpools, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["site_id"] = item3.SiteID
	d.SetId(joinResourceID(resourceMap))
	return resourceIPamSiteIPAddressPoolsRead(ctx, d, m)
}

func resourceIPamSiteIPAddressPoolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesIPAddressSubpools")
		queryParams1 := dnacentersdkgo.RetrievesIPAddressSubpoolsQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrievesIPAddressSubpools(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchNetworkSettingsRetrievesIPAddressSubpools(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesIPAddressSubpools search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItem(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ipV4AddressSpace"] = flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItemIPV4AddressSpace(item.IPV4AddressSpace)
	respItem["ipV6AddressSpace"] = flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItemIPV6AddressSpace(item.IPV6AddressSpace)
	respItem["name"] = item.Name
	respItem["poolType"] = item.PoolType
	respItem["siteId"] = item.SiteID
	respItem["siteName"] = item.SiteName
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItemIPV4AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponseIPV4AddressSpace) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["subnet"] = item.Subnet
	respItem["prefixLength"] = item.PrefixLength
	respItem["gatewayIpAddress"] = item.GatewayIPAddress
	respItem["dhcpServers"] = item.DhcpServers
	respItem["dnsServers"] = item.DNSServers
	respItem["totalAddresses"] = item.TotalAddresses
	respItem["unassignableAddresses"] = item.UnassignableAddresses
	respItem["assignedAddresses"] = item.AssignedAddresses
	respItem["defaultAssignedAddresses"] = item.DefaultAssignedAddresses
	respItem["slaacSupport"] = item.SLAacSupport
	respItem["globalPoolId"] = item.GlobalPoolID

	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrievesIPAddressSubpoolsByIDItemIPV6AddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponseIPV6AddressSpace) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["subnet"] = item.Subnet
	respItem["prefixLength"] = item.PrefixLength
	respItem["gatewayIpAddress"] = item.GatewayIPAddress
	respItem["dhcpServers"] = item.DhcpServers
	respItem["dnsServers"] = item.DNSServers
	respItem["totalAddresses"] = item.TotalAddresses
	respItem["unassignableAddresses"] = item.UnassignableAddresses
	respItem["assignedAddresses"] = item.AssignedAddresses
	respItem["defaultAssignedAddresses"] = item.DefaultAssignedAddresses
	respItem["slaacSupport"] = item.SLAacSupport
	respItem["globalPoolId"] = item.GlobalPoolID
	return []map[string]interface{}{
		respItem,
	}
}

func resourceIPamSiteIPAddressPoolsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceIPamSiteIPAddressPoolsRead(ctx, d, m)
}

func resourceIPamSiteIPAddressPoolsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete IPamSiteIPAddressPools on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpools(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpools {
	request := dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpools{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_v4_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_v4_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_v4_address_space")))) {
		request.IPV4AddressSpace = expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpoolsIPV4AddressSpace(ctx, key+".ip_v4_address_space.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_v6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_v6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_v6_address_space")))) {
		request.IPV6AddressSpace = expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpoolsIPV6AddressSpace(ctx, key+".ip_v6_address_space.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pool_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pool_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pool_type")))) {
		request.PoolType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name")))) {
		request.SiteName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpoolsIPV4AddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpoolsIPV4AddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpoolsIPV4AddressSpace{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnet")))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_length")))) {
		request.PrefixLength = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway_ip_address")))) {
		request.GatewayIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_servers")))) {
		request.DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_servers")))) {
		request.DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_addresses")))) {
		request.TotalAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unassignable_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unassignable_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unassignable_addresses")))) {
		request.UnassignableAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned_addresses")))) {
		request.AssignedAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_assigned_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_assigned_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_assigned_addresses")))) {
		request.DefaultAssignedAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_pool_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_pool_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_pool_id")))) {
		request.GlobalPoolID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPamSiteIPAddressPoolsReservecreateIPAddressSubpoolsIPV6AddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpoolsIPV6AddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsReservecreateIPAddressSubpoolsIPV6AddressSpace{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnet")))) {
		request.Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".prefix_length")))) {
		request.PrefixLength = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway_ip_address")))) {
		request.GatewayIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_servers")))) {
		request.DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_servers")))) {
		request.DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_addresses")))) {
		request.TotalAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unassignable_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unassignable_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unassignable_addresses")))) {
		request.UnassignableAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned_addresses")))) {
		request.AssignedAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_assigned_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_assigned_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_assigned_addresses")))) {
		request.DefaultAssignedAddresses = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".global_pool_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".global_pool_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".global_pool_id")))) {
		request.GlobalPoolID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsRetrievesIPAddressSubpools(m interface{}, queryParams dnacentersdkgo.RetrievesIPAddressSubpoolsQueryParams, vID string) (*dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpoolsResponse
	var ite *dnacentersdkgo.ResponseNetworkSettingsRetrievesIPAddressSubpools
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.NetworkSettings.RetrievesIPAddressSubpools(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.NetworkSettings.RetrievesIPAddressSubpools(&queryParams)
		}
		return nil, err
	} else if queryParams.SiteID != "" {
		ite, _, err = client.NetworkSettings.RetrievesIPAddressSubpools(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.SiteID {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
