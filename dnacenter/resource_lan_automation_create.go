package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLanAutomationCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on LAN Automation.

- Invoke this API to start LAN Automation for the given site.
`,

		CreateContext: resourceLanAutomationCreateCreate,
		ReadContext:   resourceLanAutomationCreateRead,
		DeleteContext: resourceLanAutomationCreateDelete,
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

						"detail": &schema.Schema{
							Description: `Detailed information of the error code.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Description: `Error code value.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Description: `Description of the error code.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestLanAutomationLANAutomation2`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"discovered_device_site_name_hierarchy": &schema.Schema{
										Description: `Discovered device site name.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"host_name_file_id": &schema.Schema{
										Description: `Use /dna/intent/api/v1/file/namespace/nw_orch api to get the file id for the already uploaded file in nw_orch namespace.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"host_name_prefix": &schema.Schema{
										Description: `Host name prefix which shall be assigned to the discovered device.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ip_pools": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_pool_name": &schema.Schema{
													Description: `Name of the IP pool.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"ip_pool_role": &schema.Schema{
													Description: `Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
									"isis_domain_pwd": &schema.Schema{
										Description: `IS-IS domain password in plain text.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mulitcast_enabled": &schema.Schema{
										Description: `To enable underlay native multicast.
`,

										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
									},
									"peer_device_managment_ipaddress": &schema.Schema{
										Description: `Peer seed management IP address.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"primary_device_interface_names": &schema.Schema{
										Description: `The list of interfaces on primary seed via which the discovered devices are connected.
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"primary_device_managment_ipaddress": &schema.Schema{
										Description: `Primary seed management IP address.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"redistribute_isis_to_bgp": &schema.Schema{
										Description: `Advertise LAN Automation summary route into BGP. 
`,

										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
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

func resourceLanAutomationCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestLanAutomationCreateLanAutomation2(ctx, "parameters.0", d)

	response1, restyResp1, err := client.LanAutomation.LanAutomation2(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing LanAutomation2", err,
			"Failure at LanAutomation2, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItem1 := flattenLanAutomationLanAutomation2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LanAutomation2 response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceLanAutomationCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLanAutomationCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestLanAutomationCreateLanAutomation2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLanAutomationLanAutomation2 {
	request := dnacentersdkgo.RequestLanAutomationLanAutomation2{}
	if v := expandRequestLanAutomationCreateLanAutomation2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestLanAutomationCreateLanAutomation2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomation2 {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomation2{}
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
		i := expandRequestLanAutomationCreateLanAutomation2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationCreateLanAutomation2Item(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomation2 {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomation2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovered_device_site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovered_device_site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovered_device_site_name_hierarchy")))) {
		request.DiscoveredDeviceSiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_device_managment_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_device_managment_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_device_managment_ipaddress")))) {
		request.PrimaryDeviceManagmentIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_device_managment_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_device_managment_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_device_managment_ipaddress")))) {
		request.PeerDeviceManagmentIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_device_interface_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_device_interface_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_device_interface_names")))) {
		request.PrimaryDeviceInterfaceNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pools")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pools")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pools")))) {
		request.IPPools = expandRequestLanAutomationCreateLanAutomation2ItemIPPoolsArray(ctx, key+".ip_pools", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mulitcast_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mulitcast_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mulitcast_enabled")))) {
		request.MulitcastEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name_prefix")))) {
		request.HostNamePrefix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name_file_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name_file_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name_file_id")))) {
		request.HostNameFileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".isis_domain_pwd")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".isis_domain_pwd")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".isis_domain_pwd")))) {
		request.IsisDomainPwd = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redistribute_isis_to_bgp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redistribute_isis_to_bgp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redistribute_isis_to_bgp")))) {
		request.RedistributeIsisToBgp = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestLanAutomationCreateLanAutomation2ItemIPPoolsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomation2IPPools {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomation2IPPools{}
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
		i := expandRequestLanAutomationCreateLanAutomation2ItemIPPools(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationCreateLanAutomation2ItemIPPools(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomation2IPPools {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomation2IPPools{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_role")))) {
		request.IPPoolRole = interfaceToString(v)
	}
	return &request
}

func flattenLanAutomationLanAutomation2Item(item *dnacentersdkgo.ResponseLanAutomationLanAutomation2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_code"] = item.ErrorCode
	respItem["message"] = item.Message
	respItem["detail"] = item.Detail
	return []map[string]interface{}{
		respItem,
	}
}
