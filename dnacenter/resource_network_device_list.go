package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceList() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Devices.

- Adds the device with given credential

- Sync the devices provided as input
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
							Description: `Ap Ethernet Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ap_manager_interface_ip": &schema.Schema{
							Description: `Ap Manager Interface Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"associated_wlc_ip": &schema.Schema{
							Description: `Associated Wlc Ip`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"boot_date_time": &schema.Schema{
							Description: `Boot Date Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"collection_interval": &schema.Schema{
							Description: `Collection Interval`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"collection_status": &schema.Schema{
							Description: `Collection Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_support_level": &schema.Schema{
							Description: `Device Support Level`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"error_code": &schema.Schema{
							Description: `Error Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"error_description": &schema.Schema{
							Description: `Error Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"family": &schema.Schema{
							Description: `Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"hostname": &schema.Schema{
							Description: `Hostname`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interface_count": &schema.Schema{
							Description: `Interface Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inventory_status_detail": &schema.Schema{
							Description: `Inventory Status Detail`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_updated": &schema.Schema{
							Description: `Last Updated`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"line_card_count": &schema.Schema{
							Description: `Line Card Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"line_card_id": &schema.Schema{
							Description: `Line Card Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"location": &schema.Schema{
							Description: `Location`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"location_name": &schema.Schema{
							Description: `Location Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"managed_atleast_once": &schema.Schema{
							Description: `Managed Atleast Once`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"management_state": &schema.Schema{
							Description: `Management State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"memory_size": &schema.Schema{
							Description: `Memory Size`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reachability_failure_reason": &schema.Schema{
							Description: `Reachability Failure Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reachability_status": &schema.Schema{
							Description: `Reachability Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"role": &schema.Schema{
							Description: `Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"role_source": &schema.Schema{
							Description: `Role Source`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"series": &schema.Schema{
							Description: `Series`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"snmp_contact": &schema.Schema{
							Description: `Snmp Contact`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"snmp_location": &schema.Schema{
							Description: `Snmp Location`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"software_type": &schema.Schema{
							Description: `Software Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"software_version": &schema.Schema{
							Description: `Software Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tag_count": &schema.Schema{
							Description: `Tag Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tunnel_udp_port": &schema.Schema{
							Description: `Tunnel Udp Port`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"up_time": &schema.Schema{
							Description: `Up Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"uptime_seconds": &schema.Schema{
							Description: `Uptime Seconds`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"waas_device_mode": &schema.Schema{
							Description: `Waas Device Mode`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_transport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"compute_device": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extended_discovery_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_secure": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"http_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"meraki_org_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_retry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_mgmt_ipaddress_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exist_mgmt_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"new_mgmt_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
	request1 := expandRequestNetworkDeviceListAddDevice2(ctx, "parameters.0", d)
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

	if err != nil || response1 != nil {
		resourceMap := make(map[string]string)
		resourceMap["serial_number"] = vvSerialNumber
		resourceMap["ip_address"] = vvIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceNetworkDeviceListRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Devices.AddDevice2(request1)
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
			diags = append(diags, diagError(
				"Failure when executing AddDevice2", err))
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

		response1, restyResp1, err := client.Devices.GetDeviceList(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDeviceList", err,
			// 	"Failure at GetDeviceList, unexpected response", ""))
			// return diags
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
		// diags = append(diags, diagErrorWithAlt(
		// 	"Failure when executing GetDeviceList", err,
		// 	"Failure at GetDeviceList, unexpected response", ""))
		// return diags
		d.SetId("")
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vSerialNumber)
		request1 := expandRequestNetworkDeviceListSyncDevices2(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil && item != nil && request1.ID == "" {
			request1.ID = item.ID
		}
		response1, restyResp1, err := client.Devices.SyncDevices2(request1)
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
				diags = append(diags, diagError(
					"Failure when executing SyncDevices2", err))
				return diags
			}
		}
	}

	return resourceNetworkDeviceListRead(ctx, d, m)
}

func resourceNetworkDeviceListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete NetworkDeviceList on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNetworkDeviceListAddDevice2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddDevice2 {
	request := dnacentersdkgo.RequestDevicesAddDevice2{}
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
		request.UpdateMgmtIPaddressList = expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressListArray(ctx, key+".update_mgmt_ipaddress_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList {
	request := []dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList {
	request := dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList{}
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

func expandRequestNetworkDeviceListSyncDevices2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesSyncDevices2 {
	request := dnacentersdkgo.RequestDevicesSyncDevices2{}
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
		request.UpdateMgmtIPaddressList = expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressListArray(ctx, key+".update_mgmt_ipaddress_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList {
	request := []dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList {
	request := dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList{}
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
