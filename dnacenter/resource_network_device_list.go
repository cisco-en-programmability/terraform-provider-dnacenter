package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceList() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Devices.

- Adds the device with given credential

- Update the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger
an inventory sync.
`,

		CreateContext: resourceNetworkDeviceListCreate,
		ReadContext:   resourceNetworkDeviceListRead,
		UpdateContext: resourceNetworkDeviceListUpdate,
		DeleteContext: resourceNetworkDeviceListDelete,
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

						"ap_ethernet_mac_address": &schema.Schema{
							Description: `AccessPoint Ethernet MacAddress of AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ap_manager_interface_ip": &schema.Schema{
							Description: `IP address of WLC on AP manager interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_wlc_ip": &schema.Schema{
							Description: `Associated Wlc Ip address of the AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_date_time": &schema.Schema{
							Description: `Device boot time
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_interval": &schema.Schema{
							Description: `Re sync Interval of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_status": &schema.Schema{
							Description: `Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `System description
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_support_level": &schema.Schema{
							Description: `Support level of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_resolved_management_address": &schema.Schema{
							Description: `Specifies the resolved ip address of dns name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Description: `Inventory status error code
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_description": &schema.Schema{
							Description: `Inventory status description
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": &schema.Schema{
							Description: `Family of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Instance Uuid of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_count": &schema.Schema{
							Description: `Number of interfaces on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_status_detail": &schema.Schema{
							Description: `Status detail of inventory sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_device_resync_start_time": &schema.Schema{
							Description: `Start time for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_managed_resync_reasons": &schema.Schema{
							Description: `Reasons for last successful sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Time in epoch when the network device info last got updated
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"last_updated": &schema.Schema{
							Description: `Time when the network device info last got updated
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_count": &schema.Schema{
							Description: `Number of linecards on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_id": &schema.Schema{
							Description: `IDs of linecards of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `[Deprecated] Location ID that is associated with the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_name": &schema.Schema{
							Description: `[Deprecated] Name of the associated location
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Description: `MAC address of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_atleast_once": &schema.Schema{
							Description: `Indicates if device went into Managed state atleast once
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_ip_address": &schema.Schema{
							Description: `IP address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_state": &schema.Schema{
							Description: `Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": &schema.Schema{
							Description: `Processor memory size
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"pending_sync_requests_count": &schema.Schema{
							Description: `Count of pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": &schema.Schema{
							Description: `Platform ID of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_failure_reason": &schema.Schema{
							Description: `Failure reason for unreachable devices
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_status": &schema.Schema{
							Description: `Device reachability status as Reachable / Unreachable
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"reasons_for_device_resync": &schema.Schema{
							Description: `Reason for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"reasons_for_pending_sync_requests": &schema.Schema{
							Description: `Reasons for pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": &schema.Schema{
							Description: `Role of device as access, distribution, border router, core
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_source": &schema.Schema{
							Description: `Role source as manual / auto
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Description: `Serial number of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"series": &schema.Schema{
							Description: `Device series
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_contact": &schema.Schema{
							Description: `SNMP contact on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_location": &schema.Schema{
							Description: `SNMP location on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Description: `Software type on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Description: `Software version on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"sync_requested_by_app": &schema.Schema{
							Description: `Applications which requested for the resync of network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_count": &schema.Schema{
							Description: `Number of tags associated with the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_udp_port": &schema.Schema{
							Description: `Mobility protocol port is stored as tunneludpport for WLC
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"up_time": &schema.Schema{
							Description: `Time that shows for how long the device has been up
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"uptime_seconds": &schema.Schema{
							Description: `Uptime in Seconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"waas_device_mode": &schema.Schema{
							Description: `WAAS device mode
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

						"cli_transport": &schema.Schema{
							Description: `CLI transport. Supported values: telnet, ssh. Required if type is NETWORK_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"compute_device": &schema.Schema{
							Description: `Compute Device or not. Options are true / false.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_password": &schema.Schema{
							Description: `CLI enable password of the device. Required if device is configured to use enable password.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extended_discovery_info": &schema.Schema{
							Description: `This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_password": &schema.Schema{
							Description: `HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_port": &schema.Schema{
							Description: `HTTP port of the device. Required if type is COMPUTE_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_secure": &schema.Schema{
							Description: `Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP. Default is true.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"http_user_name": &schema.Schema{
							Description: `HTTP Username of the device. Required if type is COMPUTE_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Description: `IP Address of the device. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"meraki_org_id": &schema.Schema{
							Description: `Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"netconf_port": &schema.Schema{
							Description: `Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password": &schema.Schema{
							Description: `CLI Password of the device. Required if type is NETWORK_DEVICE.
`,
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"serial_number": &schema.Schema{
							Description: `Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Description: `SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Description: `SNMPv3 auth protocol. Supported values: sha, md5. Required if snmpMode is authNoPriv or authPriv.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_mode": &schema.Schema{
							Description: `SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Description: `SNMPv3 priv passphrase. Required if snmpMode is authPriv.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Description: `SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_ro_community": &schema.Schema{
							Description: `SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_rw_community": &schema.Schema{
							Description: `SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_retry": &schema.Schema{
							Description: `SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"snmp_timeout": &schema.Schema{
							Description: `SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"snmp_user_name": &schema.Schema{
							Description: `SNMPV3 user name of the device. Required if snmpVersion is v3.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_version": &schema.Schema{
							Description: `SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type of device being added. Default is NETWORK_DEVICE.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_mgmt_ipaddress_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exist_mgmt_ip_address": &schema.Schema{
										Description: `existMgmtIpAddress IP Address of the device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"new_mgmt_ip_address": &schema.Schema{
										Description: `New IP Address to be Updated.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"user_name": &schema.Schema{
							Description: `CLI user name of the device. Required if type is NETWORK_DEVICE.
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

func resourceNetworkDeviceListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkDeviceListAddDeviceKnowYourNetwork(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vSerialNumber := resourceItem["serial_number"]
	vvSerialNumber := interfaceToString(vSerialNumber)
	var vIPAddress []string
	var vvIPAddress string
	if v, ok := d.GetOk("parameters.0.ip_address"); ok {
		objs := v.([]interface{})
		if len(objs) == 0 {
			return nil
		}
		for item_no := range objs {
			if i, ok := d.GetOk(fmt.Sprintf("parameters.0.ip_address.%d", item_no)); ok {
				vIPAddress = append(vIPAddress, interfaceToString(i))
			}
		}
	}
	vvIPAddress = strings.Join(vIPAddress, ",")

	queryParams1 := dnacentersdkgo.GetDeviceListQueryParams{}
	if vSerialNumber != "" {
		queryParams1.SerialNumber = []string{vvSerialNumber}
	}
	queryParams1.ManagementIPAddress = vIPAddress

	response1, _, err := client.Devices.GetDeviceList(&queryParams1)

	if err != nil || len(*response1.Response) > 0 {
		log.Printf("Prueba2 %v", response1)
		resourceMap := make(map[string]string)
		resourceMap["serial_number"] = vvSerialNumber
		resourceMap["ip_address"] = vvIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceNetworkDeviceListRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Devices.AddDeviceKnowYourNetwork(request1)
	log.Printf("ADDDEVICE ERROR %v", restyResp1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDevice2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDevice2", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddDevice2", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AddDevice2", err1))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["serial_number"] = vvSerialNumber
	resourceMap["ip_address"] = vvIPAddress
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkDeviceListRead(ctx, d, m)
}

func resourceNetworkDeviceListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSerialNumber := resourceMap["serial_number"]
	vvIPAddress := resourceMap["ip_address"]
	vIPAddress := strings.Split(vvIPAddress, ",")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceList")
		queryParams1 := dnacentersdkgo.GetDeviceListQueryParams{}
		if vSerialNumber != "" {
			queryParams1.SerialNumber = []string{vSerialNumber}
		}
		queryParams1.ManagementIPAddress = vIPAddress

		response1, restyResp1, _ := client.Devices.GetDeviceList(&queryParams1)

		/*		if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetDeviceList", err,
					"Failure at GetDeviceList, unexpected response", ""))
				return diags
			}*/
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceListItems(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceList search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceNetworkDeviceListUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSerialNumber := resourceMap["serial_number"]
	vvIPAddress := resourceMap["ip_address"]
	vIPAddress := strings.Split(vvIPAddress, ",")

	queryParams1 := dnacentersdkgo.GetDeviceListQueryParams{}
	if vSerialNumber != "" {
		queryParams1.SerialNumber = []string{vSerialNumber}
	}
	queryParams1.ManagementIPAddress = vIPAddress
	item, err := searchDevicesGetDeviceList(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDeviceList", err,
			"Failure at GetDeviceList, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vSerialNumber)
		request1 := expandRequestNetworkDeviceListUpdateDeviceDetails(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Devices.UpdateDeviceDetails(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SyncDevices2", err, restyResp1.String(),
					"Failure at SyncDevices2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SyncDevices2", err,
				"Failure at SyncDevices2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing AddDevice2", err))
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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing SyncDevices2", err1))
				return diags
			}
		}
	}

	return resourceNetworkDeviceListRead(ctx, d, m)
}

func resourceNetworkDeviceListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDeviceListDelete", err, "Delete method is not supported",
		"Failure at NetworkDeviceListDelete, unexpected response", ""))

	return diags
}
func expandRequestNetworkDeviceListAddDeviceKnowYourNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddDeviceKnowYourNetwork {
	request := dnacentersdkgo.RequestDevicesAddDeviceKnowYourNetwork{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_transport")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_transport")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_transport")))) {
		request.CliTransport = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".compute_device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".compute_device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".compute_device")))) {
		request.ComputeDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extended_discovery_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extended_discovery_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extended_discovery_info")))) {
		request.ExtendedDiscoveryInfo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_password")))) {
		request.HTTPPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_port")))) {
		request.HTTPPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_secure")))) {
		request.HTTPSecure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_user_name")))) {
		request.HTTPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".meraki_org_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".meraki_org_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".meraki_org_id")))) {
		request.MerakiOrgID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_retry")))) {
		request.SNMPRetry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_timeout")))) {
		request.SNMPTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListUpdateDeviceDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateDeviceDetails {
	request := dnacentersdkgo.RequestDevicesUpdateDeviceDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_transport")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_transport")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_transport")))) {
		request.CliTransport = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".compute_device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".compute_device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".compute_device")))) {
		request.ComputeDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extended_discovery_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extended_discovery_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extended_discovery_info")))) {
		request.ExtendedDiscoveryInfo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_password")))) {
		request.HTTPPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_port")))) {
		request.HTTPPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_secure")))) {
		request.HTTPSecure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_user_name")))) {
		request.HTTPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".meraki_org_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".meraki_org_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".meraki_org_id")))) {
		request.MerakiOrgID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_retry")))) {
		request.SNMPRetry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_timeout")))) {
		request.SNMPTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_mgmt_ipaddress_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) {
		request.UpdateMgmtIPaddressList = expandRequestNetworkDeviceListUpdateDeviceDetailsUpdateMgmtIPaddressListArray(ctx, key+".update_mgmt_ipaddress_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListUpdateDeviceDetailsUpdateMgmtIPaddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList {
	request := []dnacentersdkgo.RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList{}
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
		i := expandRequestNetworkDeviceListUpdateDeviceDetailsUpdateMgmtIPaddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListUpdateDeviceDetailsUpdateMgmtIPaddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList {
	request := dnacentersdkgo.RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exist_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) {
		request.ExistMgmtIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) {
		request.NewMgmtIPAddress = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDevicesGetDeviceList(m interface{}, queryParams dnacentersdkgo.GetDeviceListQueryParams) (*dnacentersdkgo.ResponseDevicesGetDeviceListResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDevicesGetDeviceListResponse
	var ite *dnacentersdkgo.ResponseDevicesGetDeviceList
	ite, _, err = client.Devices.GetDeviceList(&queryParams)
	if err != nil {
		return nil, err
	}
	if ite == nil {
		return nil, err
	}

	if ite.Response == nil {
		return nil, err
	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if strings.Contains(strings.Join(queryParams.SerialNumber, ","), item.SerialNumber) {
			var getItem *dnacentersdkgo.ResponseDevicesGetDeviceListResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		} else if strings.Contains(strings.Join(queryParams.ManagementIPAddress, ","), item.ManagementIPAddress) {
			var getItem *dnacentersdkgo.ResponseDevicesGetDeviceListResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
