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

func resourceIPamGlobalIPAddressPools() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Network Settings.

- Creates a global IP address pool, which is not bound to a particular site. A global pool must be either an IPv4 or
IPv6 pool.
`,

		CreateContext: resourceIPamGlobalIPAddressPoolsCreate,
		ReadContext:   resourceIPamGlobalIPAddressPoolsRead,
		UpdateContext: resourceIPamGlobalIPAddressPoolsUpdate,
		DeleteContext: resourceIPamGlobalIPAddressPoolsDelete,
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

						"address_space": &schema.Schema{
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
									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
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
						"id": &schema.Schema{
							Description: `The UUID for this global IP pool.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"pool_type": &schema.Schema{
							Description: `Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
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

						"address_space": &schema.Schema{
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
									"prefix_length": &schema.Schema{
										Description: `The network mask component, as a decimal, for the CIDR notation of this subnet.
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
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
							Description: `Once created, a global pool type cannot be changed. Tunnel: Assigns IP addresses to site-to-site VPN for IPSec tunneling. Generic: used for all other network types.
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

func resourceIPamGlobalIPAddressPoolsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestIPamGlobalIPAddressPoolsCreateAGlobalIPAddressPool(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	resourceItem := *getResourceItem(d.Get("parameters"))
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)

	queryParamImport := dnacentersdkgo.RetrievesGlobalIPAddressPoolsQueryParams{}
	item2, err := searchNetworkSettingsRetrievesGlobalIPAddressPools(m, queryParamImport, "", vvName)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		d.SetId(joinResourceID(resourceMap))
		return resourceIPamGlobalIPAddressPoolsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.NetworkSettings.CreateAGlobalIPAddressPool(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAGlobalIPAddressPool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAGlobalIPAddressPool", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateAGlobalIPAddressPool", err))
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
				"Failure when executing CreateAGlobalIPAddressPool", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.RetrievesGlobalIPAddressPoolsQueryParams{}
	item3, err := searchNetworkSettingsRetrievesGlobalIPAddressPools(m, queryParamValidate, "", vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateAGlobalIPAddressPool", err,
			"Failure at CreateAGlobalIPAddressPool, unexpected response", ""))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceIPamGlobalIPAddressPoolsRead(ctx, d, m)
}

func resourceIPamGlobalIPAddressPoolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vname := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesGlobalIPAddressPools")
		queryParams1 := dnacentersdkgo.RetrievesGlobalIPAddressPoolsQueryParams{}

		response1, restyResp1, err := client.NetworkSettings.RetrievesGlobalIPAddressPools(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchNetworkSettingsRetrievesGlobalIPAddressPools(m, queryParams1, "", vname)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenNetworkSettingsRetrievesGlobalIPAddressPoolsByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesGlobalIPAddressPools search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceIPamGlobalIPAddressPoolsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceIPamGlobalIPAddressPoolsRead(ctx, d, m)
}

func resourceIPamGlobalIPAddressPoolsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete IPamGlobalIPAddressPools on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestIPamGlobalIPAddressPoolsCreateAGlobalIPAddressPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateAGlobalIPAddressPool {
	request := dnacentersdkgo.RequestNetworkSettingsCreateAGlobalIPAddressPool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_space")))) {
		request.AddressSpace = expandRequestIPamGlobalIPAddressPoolsCreateAGlobalIPAddressPoolAddressSpace(ctx, key+".address_space.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pool_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pool_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pool_type")))) {
		request.PoolType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIPamGlobalIPAddressPoolsCreateAGlobalIPAddressPoolAddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateAGlobalIPAddressPoolAddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsCreateAGlobalIPAddressPoolAddressSpace{}
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsRetrievesGlobalIPAddressPools(m interface{}, queryParams dnacentersdkgo.RetrievesGlobalIPAddressPoolsQueryParams, vID string, name string) (*dnacentersdkgo.ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsResponse
	var ite *dnacentersdkgo.ResponseNetworkSettingsRetrievesGlobalIPAddressPools
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.NetworkSettings.RetrievesGlobalIPAddressPools(nil)
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
			nResponse, _, err = client.NetworkSettings.RetrievesGlobalIPAddressPools(&queryParams)
		}
		return nil, err
	} else if name != "" {
		ite, _, err = client.NetworkSettings.RetrievesGlobalIPAddressPools(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}

func flattenNetworkSettingsRetrievesGlobalIPAddressPoolsByIDItem(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["addressSpace"] = flattenNetworkSettingsRetrievesGlobalIPAddressPoolsByIDItemAddressSpace(item.AddressSpace)
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["poolType"] = item.PoolType
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrievesGlobalIPAddressPoolsByIDItemAddressSpace(item *dnacentersdkgo.ResponseNetworkSettingsRetrievesGlobalIPAddressPoolsResponseAddressSpace) []map[string]interface{} {
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
	return []map[string]interface{}{
		respItem,
	}
}
