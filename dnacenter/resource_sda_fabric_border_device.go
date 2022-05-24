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

func resourceSdaFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Add border device in SDA Fabric

- Delete border device from SDA Fabric
`,

		CreateContext: resourceSdaFabricBorderDeviceCreate,
		ReadContext:   resourceSdaFabricBorderDeviceRead,
		UpdateContext: resourceSdaFabricBorderDeviceUpdate,
		DeleteContext: resourceSdaFabricBorderDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddBorderDeviceInSDAFabric`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"border_session_type": &schema.Schema{
							Description: `Border Session Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"border_with_external_connectivity": &schema.Schema{
							Description: `Border With External Connectivity (Note: True for transit and False for non-transit border)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"connected_to_internet": &schema.Schema{
							Description: `Connected to Internet
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the provisioned Device
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_role": &schema.Schema{
							Description: `Supported Device Roles in SD-Access fabric. Allowed roles are "Border_Node","Control_Plane_Node","Edge_Node". E.g. ["Border_Node"] or ["Border_Node", "Control_Plane_Node"] or ["Border_Node", "Control_Plane_Node","Edge_Node"]
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"external_connectivity_ip_pool_name": &schema.Schema{
							Description: `External Connectivity IpPool Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_connectivity_settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"external_autonomou_system_number": &schema.Schema{
										Description: `External Autonomous System Number peer (e.g.,1-65535)
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_description": &schema.Schema{
										Description: `Interface Description
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface Name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"l2_handoff": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"virtual_network_name": &schema.Schema{
													Description: `Virtual Network Name, that is associated to Fabric Site
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"vlan_name": &schema.Schema{
													Description: `Vlan Name of L2 Handoff
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"l3_handoff": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"virtual_network": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"virtual_network_name": &schema.Schema{
																Description: `Virtual Network Name, that is associated to Fabric Site
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"vlan_id": &schema.Schema{
																Description: `Vlan Id (e.g.,2-4096 except for reserved VLANs (1002-1005, 2046, 4095))
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"external_domain_routing_protocol_name": &schema.Schema{
							Description: `External Domain Routing Protocol Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"internal_autonomou_system_number": &schema.Schema{
							Description: `Internal Autonomouns System Number (e.g.,1-65535)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"sda_transit_network_name": &schema.Schema{
							Description: `SD-Access Transit Network Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy of provisioned Device(site should be part of Fabric Site)
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

func resourceSdaFabricBorderDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabric(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddBorderDeviceInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddBorderDeviceInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddBorderDeviceInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricBorderDeviceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetBorderDeviceDetailFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetBorderDeviceDetailFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		response1, restyResp1, err := client.Sda.GetBorderDeviceDetailFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBorderDeviceDetailFromSdaFabric", err,
				"Failure at GetBorderDeviceDetailFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetBorderDeviceDetailFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetBorderDeviceDetailFromSdaFabric response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaFabricBorderDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricBorderDeviceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	queryParams1 := dnacentersdkgo.GetBorderDeviceDetailFromSdaFabricQueryParams
	queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress
	item, err := searchSdaGetBorderDeviceDetailFromSDAFabric(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetBorderDeviceDetailFromSDAFabric", err,
			"Failure at GetBorderDeviceDetailFromSDAFabric, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Sda.GetBorderDeviceDetailFromSdaFabric(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSdaGetBorderDeviceDetailFromSdaFabric(m, getResp1, nil)
		item1, err := searchSdaGetBorderDeviceDetailFromSdaFabric(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Sda.DeleteBorderDeviceFromSdaFabric()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteBorderDeviceFromSdaFabric", err, restyResp1.String(),
				"Failure at DeleteBorderDeviceFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteBorderDeviceFromSdaFabric", err,
			"Failure at DeleteBorderDeviceFromSdaFabric, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddBorderDeviceInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddBorderDeviceInSdaFabric{}
	if v := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabric {
	request := []dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabric{}
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
		i := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabric {
	request := dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_role")))) {
		request.DeviceRole = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_domain_routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) {
		request.ExternalDomainRoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) {
		request.ExternalConnectivityIPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) {
		request.InternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_session_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_session_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_session_type")))) {
		request.BorderSessionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_to_internet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_to_internet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_to_internet")))) {
		request.ConnectedToInternet = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sda_transit_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sda_transit_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sda_transit_network_name")))) {
		request.SdaTransitNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_with_external_connectivity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_with_external_connectivity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_with_external_connectivity")))) {
		request.BorderWithExternalConnectivity = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_settings")))) {
		request.ExternalConnectivitySettings = expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsArray(ctx, key+".external_connectivity_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings {
	request := []dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings{}
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
		i := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings {
	request := dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_description")))) {
		request.InterfaceDescription = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) {
		request.ExternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l3_handoff")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l3_handoff")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l3_handoff")))) {
		request.L3Handoff = expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffArray(ctx, key+".l3_handoff", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l2_handoff")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l2_handoff")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l2_handoff")))) {
		request.L2Handoff = expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL2HandoffArray(ctx, key+".l2_handoff", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff {
	request := []dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff{}
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
		i := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3Handoff(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3Handoff(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff {
	request := dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network")))) {
		request.VirtualNetwork = expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffVirtualNetworkArray(ctx, key+".virtual_network", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffVirtualNetworkArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork {
	request := []dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork{}
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
		i := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffVirtualNetwork(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL3HandoffVirtualNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork {
	request := dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL2HandoffArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff {
	request := []dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff{}
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
		i := expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL2Handoff(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricBorderDeviceAddBorderDeviceInSdaFabricItemExternalConnectivitySettingsL2Handoff(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff {
	request := dnacentersdkgo.RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
