package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPamSiteIPAddressPoolsID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Network Settings.

- Releases an IP address subpool.
**Release** completely removes the subpool from the site, and from all child sites, and frees the address use from the
global pool(s).  Subpools cannot be released when assigned addresses in use.

- Updates an IP address subpool, which reserves address space from a global pool (or global pools) for a particular
site.
Restrictions on updating an IP address subpool: The *poolType* cannot be changed. The *siteId* cannot be changed. The
*ipV4AddressSpace* may not be removed. The *globalPoolId*, *subnet*, and *prefixLength* cannot be changed once it's
already been set. However you may edit a subpool to add an IP address space if it does not already have one.
`,

		CreateContext: resourceIPamSiteIPAddressPoolsIDCreate,
		ReadContext:   resourceIPamSiteIPAddressPoolsIDRead,
		UpdateContext: resourceIPamSiteIPAddressPoolsIDUpdate,
		DeleteContext: resourceIPamSiteIPAddressPoolsIDDelete,
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

						"id": &schema.Schema{
							Description: `id path parameter. The *id* of the IP address subpool to update.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_v4_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
								},
							},
						},
						"ip_v6_address_space": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
					},
				},
			},
		},
	}
}

func resourceIPamSiteIPAddressPoolsIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceIPamSiteIPAddressPoolsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesAnIPAddressSubpool")
		vvID := vID

		response1, restyResp1, err := client.NetworkSettings.RetrievesAnIPAddressSubpool(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
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

		return diags

	}
	return diags
}

func resourceIPamSiteIPAddressPoolsIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpool(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdatesAnIPAddressSubpool(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAnIPAddressSubpool", err, restyResp1.String(),
					"Failure at UpdatesAnIPAddressSubpool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAnIPAddressSubpool", err,
				"Failure at UpdatesAnIPAddressSubpool, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesAnIPAddressSubpool", err))
			return diags
		}
		taskId := response1.Response.TaskID
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
					"Failure when executing UpdatesAnIPAddressSubpool", err1))
				return diags
			}
		}

	}

	return resourceIPamSiteIPAddressPoolsIDRead(ctx, d, m)
}

func resourceIPamSiteIPAddressPoolsIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.NetworkSettings.ReleaseAnIPAddressSubpool(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ReleaseAnIPAddressSubpool", err, restyResp1.String(),
				"Failure at ReleaseAnIPAddressSubpool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ReleaseAnIPAddressSubpool", err,
			"Failure at ReleaseAnIPAddressSubpool, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ReleaseAnIPAddressSubpool", err))
		return diags
	}
	taskId := response1.Response.TaskID
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
				"Failure when executing ReleaseAnIPAddressSubpool", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpool {
	request := dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_v4_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_v4_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_v4_address_space")))) {
		request.IPV4AddressSpace = expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpoolIPV4AddressSpace(ctx, key+".ip_v4_address_space.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_v6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_v6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_v6_address_space")))) {
		request.IPV6AddressSpace = expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpoolIPV6AddressSpace(ctx, key+".ip_v6_address_space.0", d)
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpoolIPV4AddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpoolIPV4AddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpoolIPV4AddressSpace{}
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

func expandRequestIPamSiteIPAddressPoolsIDUpdatesAnIPAddressSubpoolIPV6AddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpoolIPV6AddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsUpdatesAnIPAddressSubpoolIPV6AddressSpace{}
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
