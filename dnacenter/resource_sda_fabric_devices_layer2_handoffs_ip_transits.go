package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricDevicesLayer2HandoffsIPTransits() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds layer 3 handoffs with ip transit in fabric devices based on user input.

- Updates layer 3 handoffs with ip transit of fabric devices based on user input.

- Deletes layer 3 handoffs with ip transit of a fabric device based on user input.

- Deletes a layer 3 handoff with ip transit of a fabric device by id.
`,

		CreateContext: resourceSdaFabricDevicesLayer2HandoffsIPTransitsCreate,
		ReadContext:   resourceSdaFabricDevicesLayer2HandoffsIPTransitsRead,
		UpdateContext: resourceSdaFabricDevicesLayer2HandoffsIPTransitsUpdate,
		DeleteContext: resourceSdaFabricDevicesLayer2HandoffsIPTransitsDelete,
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
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddFabricDevicesLayer3HandoffsWithIpTransit`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_connectivity_ip_pool_name": &schema.Schema{
										Description: `External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this device is assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the fabric device layer 3 handoff ip transit. (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name of the layer 3 handoff ip transit.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"local_ip_address": &schema.Schema{
										Description: `Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"local_ipv6_address": &schema.Schema{
										Description: `Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device ID of the fabric device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"remote_ip_address": &schema.Schema{
										Description: `Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"remote_ipv6_address": &schema.Schema{
										Description: `Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tcp_mss_adjustment": &schema.Schema{
										Description: `TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"transit_network_id": &schema.Schema{
										Description: `ID of the transit network of the layer 3 handoff ip transit.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"virtual_network_name": &schema.Schema{
										Description: `Name of the virtual network associated with this fabric site.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vlan_id": &schema.Schema{
										Description: `VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network.  Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricDevicesLayer2HandoffsIPTransitsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransit(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	vInterfaceName := resourceItem["interface_name"]
	vvInterfaceName := interfaceToString(vInterfaceName)
	queryParamImport := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams{}
	queryParamImport.FabricID = vvFabricID
	item2, err := searchSdaGetFabricDevicesLayer3HandoffsWithIPTransit(m, queryParamImport, vvID, vvInterfaceName)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["fabric_id"] = vvFabricID
		resourceMap["id"] = item2.ID
		resourceMap["name"] = item2.InterfaceName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricDevicesLayer2HandoffsIPTransitsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddFabricDevicesLayer3HandoffsWithIPTransit(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabricDevicesLayer3HandoffsWithIPTransit", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevicesLayer3HandoffsWithIPTransit", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevicesLayer3HandoffsWithIPTransit", err))
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
				"Failure when executing AddFabricDevicesLayer3HandoffsWithIPTransit", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams{}
	queryParamValidate.FabricID = vvFabricID
	item3, err := searchSdaGetFabricDevicesLayer3HandoffsWithIPTransit(m, queryParamValidate, vvID, vvInterfaceName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddFabricDevicesLayer3HandoffsWithIPTransit", err,
			"Failure at AddFabricDevicesLayer3HandoffsWithIPTransit, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["fabric_id"] = vvFabricID
	resourceMap["id"] = item3.ID
	resourceMap["name"] = item3.InterfaceName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricDevicesLayer2HandoffsIPTransitsRead(ctx, d, m)
}

func resourceSdaFabricDevicesLayer2HandoffsIPTransitsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]
	vName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevicesLayer3HandoffsWithIPTransit")
		queryParams1 := dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams{}

		queryParams1.FabricID = vFabricID

		item1, err := searchSdaGetFabricDevicesLayer3HandoffsWithIPTransit(m, queryParams1, vvID, vName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}

		items := []dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse{
			*item1,
		}

		// Review flatten function used
		vItem1 := flattenSdaGetFabricDevicesLayer3HandoffsWithIPTransitItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevicesLayer3HandoffsWithIPTransit search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricDevicesLayer2HandoffsIPTransitsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransit(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateFabricDevicesLayer3HandoffsWithIPTransit(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFabricDevicesLayer3HandoffsWithIPTransit", err, restyResp1.String(),
					"Failure at UpdateFabricDevicesLayer3HandoffsWithIPTransit, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFabricDevicesLayer3HandoffsWithIPTransit", err,
				"Failure at UpdateFabricDevicesLayer3HandoffsWithIPTransit, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateFabricDevicesLayer3HandoffsWithIPTransit", err))
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
					"Failure when executing UpdateFabricDevicesLayer3HandoffsWithIPTransit", err1))
				return diags
			}
		}

	}

	return resourceSdaFabricDevicesLayer2HandoffsIPTransitsRead(ctx, d, m)
}

func resourceSdaFabricDevicesLayer2HandoffsIPTransitsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Sda.DeleteFabricDeviceLayer3HandoffWithIPTransitByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFabricDeviceLayer3HandoffWithIPTransitByID", err, restyResp1.String(),
				"Failure at DeleteFabricDeviceLayer3HandoffWithIPTransitByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFabricDeviceLayer3HandoffWithIPTransitByID", err,
			"Failure at DeleteFabricDeviceLayer3HandoffWithIPTransitByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricDeviceLayer3HandoffWithIPTransitByID", err))
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
				"Failure when executing DeleteFabricDeviceLayer3HandoffWithIPTransitByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransit(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransit {
	request := dnacentersdkgo.RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransit{}
	if v := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransitItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransitItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit {
	request := []dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransitItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsAddFabricDevicesLayer3HandoffsWithIPTransitItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit {
	request := dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_network_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_network_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_network_id")))) {
		request.TransitNetworkID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) {
		request.ExternalConnectivityIPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tcp_mss_adjustment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) {
		request.TCPMssAdjustment = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_ip_address")))) {
		request.LocalIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_ip_address")))) {
		request.RemoteIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_ipv6_address")))) {
		request.LocalIPv6Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_ipv6_address")))) {
		request.RemoteIPv6Address = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransit(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit {
	request := dnacentersdkgo.RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit{}
	if v := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransitItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransitItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit {
	request := []dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransitItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsIPTransitsUpdateFabricDevicesLayer3HandoffsWithIPTransitItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit {
	request := dnacentersdkgo.RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transit_network_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transit_network_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transit_network_id")))) {
		request.TransitNetworkID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) {
		request.ExternalConnectivityIPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tcp_mss_adjustment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) {
		request.TCPMssAdjustment = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_ip_address")))) {
		request.LocalIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_ip_address")))) {
		request.RemoteIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_ipv6_address")))) {
		request.LocalIPv6Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_ipv6_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_ipv6_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_ipv6_address")))) {
		request.RemoteIPv6Address = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetFabricDevicesLayer3HandoffsWithIPTransit(m interface{}, queryParams dnacentersdkgo.GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams, vID string, vInterfaceName string) (*dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse
	// var ite *dnacentersdkgo.ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransit

	queryParams.Offset = 1
	nResponse, _, err := client.Sda.GetFabricDevicesLayer3HandoffsWithIPTransit(&queryParams)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.ID || vInterfaceName == item.InterfaceName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Sda.GetFabricDevicesLayer3HandoffsWithIPTransit(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
