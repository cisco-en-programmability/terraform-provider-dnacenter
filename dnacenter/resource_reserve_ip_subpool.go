package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReserveIPSubpool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to delete the reserved ip subpool

- API to reserve an ip subpool from the global pool

- API to update ip subpool from the global pool
`,

		CreateContext: resourceReserveIPSubpoolCreate,
		ReadContext:   resourceReserveIPSubpoolRead,
		UpdateContext: resourceReserveIPSubpoolUpdate,
		DeleteContext: resourceReserveIPSubpoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. Id of reserve ip subpool to be deleted.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"ipv4_dhcp_servers": &schema.Schema{
							Description: `IPv4 input for dhcp server ip example: 1.1.1.1
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_dns_servers": &schema.Schema{
							Description: `IPv4 input for dns server ip example: 4.4.4.4
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_gate_way": &schema.Schema{
							Description: `Gateway ip address details, example: 175.175.0.1
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_global_pool": &schema.Schema{
							Description: `IP v4 Global pool address with cidr, example: 175.175.0.0/16
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_prefix": &schema.Schema{
							Description: `IPv4 prefix value is true, the ip4 prefix length input field is enabled , if it is false ipv4 total Host input is enable
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"ipv4_prefix_length": &schema.Schema{
							Description: `The ipv4 prefix length is required when ipv4prefix value is true.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv4_subnet": &schema.Schema{
							Description: `IPv4 Subnet address, example: 175.175.0.0
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_total_host": &schema.Schema{
							Description: `IPv4 total host is required when ipv4prefix value is false.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv6_address_space": &schema.Schema{
							Description: `If the value is false only ipv4 input are required, otherwise both ipv6 and ipv4 are required
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"ipv6_dhcp_servers": &schema.Schema{
							Description: `IPv6 format dhcp server as input example : 2001:db8::1234
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv6_dns_servers": &schema.Schema{
							Description: `IPv6 format dns server input example: 2001:db8::1234
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv6_gate_way": &schema.Schema{
							Description: `Gateway ip address details, example: 2001:db8:85a3:0:100::1
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_global_pool": &schema.Schema{
							Description: `IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_prefix": &schema.Schema{
							Description: `Ipv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false ipv6 total Host input is enable
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"ipv6_prefix_length": &schema.Schema{
							Description: `IPv6 prefix length is required when the ipv6prefix value is true
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipv6_subnet": &schema.Schema{
							Description: `IPv6 Subnet address, example :2001:db8:85a3:0:100::
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_total_host": &schema.Schema{
							Description: `IPv6 total host is required when ipv6prefix value is false.
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Name of the reserve ip sub pool
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id of site to update sub pool.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"slaac_support": &schema.Schema{
							Description: `Slaac Support`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"type": &schema.Schema{
							Description: `Type of the reserve ip sub pool
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceReserveIPSubpoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestReserveIPSubpoolReserveIPSubpool(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSiteID, okSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.NetworkSettings.ReserveIPSubpool(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ReserveIPSubpool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ReserveIPSubpool", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceReserveIPSubpoolRead(ctx, d, m)
}

func resourceReserveIPSubpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetReserveIPSubpool")
		queryParams1 := dnacentersdkgo.GetReserveIPSubpoolQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID
		}
		if okOffset {
			queryParams1.Offset = vOffset
		}
		if okLimit {
			queryParams1.Limit = vLimit
		}

		response1, restyResp1, err := client.NetworkSettings.GetReserveIPSubpool(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetReserveIPSubpool", err,
				"Failure at GetReserveIPSubpool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNetworkSettingsGetReserveIPSubpool(m, response1, nil)
		item1, err := searchNetworkSettingsGetReserveIPSubpool(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetReserveIPSubpool response", err,
				"Failure when searching item from GetReserveIPSubpool, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenNetworkSettingsGetReserveIPSubpoolByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetReserveIPSubpool search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceReserveIPSubpoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	queryParams1 := dnacentersdkgo.GetReserveIPSubpoolQueryParams
	queryParams1.SiteID = vSiteID
	queryParams1.Offset = vOffset
	queryParams1.Limit = vLimit
	item, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetReserveIPSubpool", err,
			"Failure at GetReserveIPSubpool, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestReserveIPSubpoolUpdateReserveIPSubpool(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateReserveIPSubpool(vvSiteID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateReserveIPSubpool", err, restyResp1.String(),
					"Failure at UpdateReserveIPSubpool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateReserveIPSubpool", err,
				"Failure at UpdateReserveIPSubpool, unexpected response", ""))
			return diags
		}
	}

	return resourceReserveIPSubpoolRead(ctx, d, m)
}

func resourceReserveIPSubpoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	queryParams1 := dnacentersdkgo.GetReserveIPSubpoolQueryParams
	queryParams1.SiteID = vSiteID
	queryParams1.Offset = vOffset
	queryParams1.Limit = vLimit
	item, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetReserveIPSubpool", err,
			"Failure at GetReserveIPSubpool, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.NetworkSettings.GetReserveIPSubpool(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkSettingsGetReserveIPSubpool(m, getResp1, nil)
		item1, err := searchNetworkSettingsGetReserveIPSubpool(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	response1, restyResp1, err := client.NetworkSettings.ReleaseReserveIPSubpool(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ReleaseReserveIPSubpool", err, restyResp1.String(),
				"Failure at ReleaseReserveIPSubpool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ReleaseReserveIPSubpool", err,
			"Failure at ReleaseReserveIPSubpool, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestReserveIPSubpoolReserveIPSubpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsReserveIPSubpool {
	request := dnacentersdkgo.RequestNetworkSettingsReserveIPSubpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_space")))) {
		request.IPv6AddressSpace = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_global_pool")))) {
		request.IPv4GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_prefix")))) {
		request.IPv4Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_prefix_length")))) {
		request.IPv4PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_subnet")))) {
		request.IPv4Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_gate_way")))) {
		request.IPv4GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) {
		request.IPv4DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) {
		request.IPv4DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_global_pool")))) {
		request.IPv6GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix")))) {
		request.IPv6Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) {
		request.IPv6PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_gate_way")))) {
		request.IPv6GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) {
		request.IPv6DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) {
		request.IPv6DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_total_host")))) {
		request.IPv4TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_total_host")))) {
		request.IPv6TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestReserveIPSubpoolUpdateReserveIPSubpool(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateReserveIPSubpool {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateReserveIPSubpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_space")))) {
		request.IPv6AddressSpace = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) {
		request.IPv4DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) {
		request.IPv4DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_global_pool")))) {
		request.IPv6GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix")))) {
		request.IPv6Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) {
		request.IPv6PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_gate_way")))) {
		request.IPv6GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) {
		request.IPv6DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) {
		request.IPv6DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_total_host")))) {
		request.IPv6TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_gate_way")))) {
		request.IPv4GateWay = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsGetReserveIPSubpool(m interface{}, queryParams dnacentersdkgo.GetReserveIPSubpoolQueryParams) (*dnacentersdkgo.ResponseItemNetworkSettingsGetReserveIPSubpool, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemNetworkSettingsGetReserveIPSubpool
	var ite *dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpool
	ite, _, err = client.NetworkSettings.GetReserveIPSubpool(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemNetworkSettingsGetReserveIPSubpool
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
