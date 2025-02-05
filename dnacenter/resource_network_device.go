package dnacenter

import (
	"context"

	"errors"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and delete operations on Devices.

- This resource allows any network device that is not currently provisioned to be removed from the inventory. Important:
Devices currently provisioned cannot be deleted. To delete a provisioned device, the device must be first deprovisioned.
`,

		CreateContext: resourceNetworkDeviceCreate,
		ReadContext:   resourceNetworkDeviceRead,
		UpdateContext: resourceNetworkDeviceUpdate,
		DeleteContext: resourceNetworkDeviceDelete,
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
						"vendor": &schema.Schema{
							Description: `Vendor details
`,
							Type:     schema.TypeString,
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
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. Device ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	// resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceByID")
		vvID := vID

		response1, restyResp1, err := client.Devices.GetDeviceByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceByID response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceNetworkDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDeviceUpdate", err, "Update method is not supported",
		"Failure at NetworkDeviceUpdate, unexpected response", ""))

	return diags
}

func resourceNetworkDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	// queryParamDelete := dnacentersdkgo.DeleteDeviceByIDQueryParams{}

	vvID := resourceMap["id"]
	// queryParamDelete.CleanConfig = vvCleanConfig

	response1, restyResp1, err := client.Devices.DeleteDeviceByID(vvID, nil)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceByID", err, restyResp1.String(),
				"Failure at DeleteDeviceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceByID", err,
			"Failure at DeleteDeviceByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteDeviceByID", err))
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
				"Failure when executing DeleteDeviceByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
