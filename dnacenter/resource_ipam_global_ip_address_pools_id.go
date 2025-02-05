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

func resourceIPamGlobalIPAddressPoolsID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Network Settings.

- Updates a global IP address pool.
Restrictions on updating a global IP address pool: The *poolType* cannot be changed. The *subnet* and *prefixLength*
within *addressSpace* cannot be changed.

- Deletes a global IP address pool.  A global IP address pool can only be deleted if there are no subpools reserving
address space from it.
`,

		CreateContext: resourceIPamGlobalIPAddressPoolsIDCreate,
		ReadContext:   resourceIPamGlobalIPAddressPoolsIDRead,
		UpdateContext: resourceIPamGlobalIPAddressPoolsIDUpdate,
		DeleteContext: resourceIPamGlobalIPAddressPoolsIDDelete,
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
						"id": &schema.Schema{
							Description: `id path parameter. The *id* of the global IP address pool to update.
`,
							Type:     schema.TypeString,
							Required: true,
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

func resourceIPamGlobalIPAddressPoolsIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceIPamGlobalIPAddressPoolsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesAGlobalIPAddressPool")
		vvID := vID

		response1, restyResp1, err := client.NetworkSettings.RetrievesAGlobalIPAddressPool(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrievesAGlobalIPAddressPoolItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesAGlobalIPAddressPool response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceIPamGlobalIPAddressPoolsIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestIPamGlobalIPAddressPoolsIDUpdatesAGlobalIPAddressPool(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdatesAGlobalIPAddressPool(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAGlobalIPAddressPool", err, restyResp1.String(),
					"Failure at UpdatesAGlobalIPAddressPool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAGlobalIPAddressPool", err,
				"Failure at UpdatesAGlobalIPAddressPool, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesAGlobalIPAddressPool", err))
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
					"Failure when executing UpdatesAGlobalIPAddressPool", err1))
				return diags
			}
		}

	}

	return resourceIPamGlobalIPAddressPoolsIDRead(ctx, d, m)
}

func resourceIPamGlobalIPAddressPoolsIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.NetworkSettings.DeleteAGlobalIPAddressPool(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAGlobalIPAddressPool", err, restyResp1.String(),
				"Failure at DeleteAGlobalIPAddressPool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAGlobalIPAddressPool", err,
			"Failure at DeleteAGlobalIPAddressPool, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteAGlobalIPAddressPool", err))
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
				"Failure when executing DeleteAGlobalIPAddressPool", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestIPamGlobalIPAddressPoolsIDUpdatesAGlobalIPAddressPool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdatesAGlobalIPAddressPool {
	request := dnacentersdkgo.RequestNetworkSettingsUpdatesAGlobalIPAddressPool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_space")))) {
		request.AddressSpace = expandRequestIPamGlobalIPAddressPoolsIDUpdatesAGlobalIPAddressPoolAddressSpace(ctx, key+".address_space.0", d)
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

func expandRequestIPamGlobalIPAddressPoolsIDUpdatesAGlobalIPAddressPoolAddressSpace(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdatesAGlobalIPAddressPoolAddressSpace {
	request := dnacentersdkgo.RequestNetworkSettingsUpdatesAGlobalIPAddressPoolAddressSpace{}
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
