package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLanAutomationV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on LAN Automation.

- Invoke V2 LAN Automation Start API, which supports optional auto-stop processing feature based on the provided timeout
or a specific device list, or both. The stop processing will be executed automatically when either of the cases is
satisfied, without specifically calling the stop API. The V2 API behaves similarly to  if no timeout or device list is
provided, and the user needs to call the stop API for LAN Automation stop processing. With the V2 API, the user can also
specify the level up to which the devices can be LAN automated.
`,

		CreateContext: resourceLanAutomationV2Create,
		ReadContext:   resourceLanAutomationV2Read,
		DeleteContext: resourceLanAutomationV2Delete,
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

						"task_id": &schema.Schema{
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `url to check the status of task
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
							Description: `Array of RequestLanAutomationLANAutomationStartV2`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"discovered_device_site_name_hierarchy": &schema.Schema{
										Description: `Discovered device site name.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"discovery_devices": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_host_name": &schema.Schema{
													Description: `Hostname of the device
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"device_management_ipaddress": &schema.Schema{
													Description: `Management IP Address of the device
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"device_serial_number": &schema.Schema{
													Description: `Serial number of the device
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"device_site_name_hierarchy": &schema.Schema{
													Description: `Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"discovery_level": &schema.Schema{
										Description: `Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"discovery_timeout": &schema.Schema{
										Description: `Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080]. If both timeout and discovery devices list are provided, the stop processing will be attempted whichever happens earlier. Users can always use the LAN Automation delete API to force stop processing.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"host_name_file_id": &schema.Schema{
										Description: `Use /dna/intent/api/v1/file/namespace/nw_orch API to get the file ID for the already uploaded file in the nw_orch namespace.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"host_name_prefix": &schema.Schema{
										Description: `Host name prefix assigned to the discovered device.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ip_pools": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_pool_name": &schema.Schema{
													Description: `Name of the IP pool.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"ip_pool_role": &schema.Schema{
													Description: `Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
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
										Computed: true,
									},
									"multicast_enabled": &schema.Schema{
										Description: `Enable underlay native multicast.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"peer_device_managment_ipaddress": &schema.Schema{
										Description: `Peer seed management IP address.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"primary_device_interface_names": &schema.Schema{
										Description: `The list of interfaces on primary seed via which the discovered devices are connected.
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
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
										Computed: true,
									},
									"redistribute_isis_to_bgp": &schema.Schema{
										Description: `Advertise LAN Automation summary route into BGP.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
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

func resourceLanAutomationV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestLanAutomationV2LanAutomationStartV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.LanAutomation.LanAutomationStartV2(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing LanAutomationStartV2", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing LANAutomationStartV2", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing LANAutomationStartV2", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenLanAutomationLanAutomationStartV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LanAutomationStartV2 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceLanAutomationV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLanAutomationV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestLanAutomationV2LanAutomationStartV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLanAutomationLanAutomationStartV2 {
	request := dnacentersdkgo.RequestLanAutomationLanAutomationStartV2{}
	if v := expandRequestLanAutomationV2LanAutomationStartV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2 {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2{}
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
		i := expandRequestLanAutomationV2LanAutomationStartV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2Item(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2 {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2{}
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
		request.IPPools = expandRequestLanAutomationV2LanAutomationStartV2ItemIPPoolsArray(ctx, key+".ip_pools", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_enabled")))) {
		request.MulticastEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name_prefix")))) {
		request.HostNamePrefix = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name_file_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name_file_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name_file_id")))) {
		request.HostNameFileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".redistribute_isis_to_bgp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".redistribute_isis_to_bgp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".redistribute_isis_to_bgp")))) {
		request.RedistributeIsisToBgp = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".isis_domain_pwd")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".isis_domain_pwd")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".isis_domain_pwd")))) {
		request.IsisDomainPwd = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_level")))) {
		request.DiscoveryLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_timeout")))) {
		request.DiscoveryTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_devices")))) {
		request.DiscoveryDevices = expandRequestLanAutomationV2LanAutomationStartV2ItemDiscoveryDevicesArray(ctx, key+".discovery_devices", d)
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2ItemIPPoolsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2IPPools {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2IPPools{}
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
		i := expandRequestLanAutomationV2LanAutomationStartV2ItemIPPools(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2ItemIPPools(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2IPPools {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2IPPools{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_role")))) {
		request.IPPoolRole = interfaceToString(v)
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2ItemDiscoveryDevicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices{}
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
		i := expandRequestLanAutomationV2LanAutomationStartV2ItemDiscoveryDevices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationV2LanAutomationStartV2ItemDiscoveryDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_serial_number")))) {
		request.DeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_host_name")))) {
		request.DeviceHostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_site_name_hierarchy")))) {
		request.DeviceSiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ipaddress")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	return &request
}

func flattenLanAutomationLanAutomationStartV2Item(item *dnacentersdkgo.ResponseLanAutomationLanAutomationStartV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
