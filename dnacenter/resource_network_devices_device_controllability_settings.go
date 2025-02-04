package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevicesDeviceControllabilitySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Site Design.

- Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some
device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs
to manage devices. Changes are made on network devices  during discovery, when adding a device to Inventory, or when
assigning a device to a site. If changes  are made to any settings that are under the scope of this process, these
changes are applied to the  network devices during the Provision and Update Telemetry Settings operations, even if
Device  Controllability is disabled. The following device settings will be enabled as part of  Device Controllability
when devices are discovered.

  SNMP Credentials.
  NETCONF Credentials.

Subsequent to discovery, devices will be added to Inventory. The following device settings will be  enabled when devices
are added to inventory.

  Cisco TrustSec (CTS) Credentials.

The following device settings will be enabled when devices are assigned to a site. Some of these  settings can be
defined at a site level under Design > Network Settings > Telemetry & Wireless.

  Wired Endpoint Data Collection Enablement.
  Controller Certificates.
  SNMP Trap Server Definitions.
  Syslog Server Definitions.
  Application Visibility.
  Application QoS Policy.
  Wireless Service Assurance (WSA).
  Wireless Telemetry.
  DTLS Ciphersuite.
  AP Impersonation.

If Device Controllability is disabled, Catalyst Center does not configure any of the preceding  credentials or settings
on devices during discovery, at runtime, or during site assignment. However,  the telemetry settings and related
configuration are pushed when the device is provisioned or when the  update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on  the device.

  SWIM certificate issue.
  IOS WLC NA certificate issue.
  PKCS12 certificate issue.
  IOS telemetry configuration issue.

The autocorrect telemetry config feature is supported only when Device Controllability is enabled.
`,

		CreateContext: resourceNetworkDevicesDeviceControllabilitySettingsCreate,
		ReadContext:   resourceNetworkDevicesDeviceControllabilitySettingsRead,
		UpdateContext: resourceNetworkDevicesDeviceControllabilitySettingsUpdate,
		DeleteContext: resourceNetworkDevicesDeviceControllabilitySettingsDelete,
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

						"autocorrect_telemetry_config": &schema.Schema{
							Description: `If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_controllability": &schema.Schema{
							Description: `If it is true, device controllability is enabled. If it is false, device controllability is disabled.
`,
							// Type:        schema.TypeBool,
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

						"autocorrect_telemetry_config": &schema.Schema{
							Description: `If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"device_controllability": &schema.Schema{
							Description: `If it is true, device controllability is enabled. If it is false, device controllability is disabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDevicesDeviceControllabilitySettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceNetworkDevicesDeviceControllabilitySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceControllabilitySettings")

		response1, restyResp1, err := client.SiteDesign.GetDeviceControllabilitySettings()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetDeviceControllabilitySettingsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceControllabilitySettings response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceNetworkDevicesDeviceControllabilitySettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestNetworkDevicesDeviceControllabilitySettingsUpdateDeviceControllabilitySettings(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SiteDesign.UpdateDeviceControllabilitySettings(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceControllabilitySettings", err, restyResp1.String(),
					"Failure at UpdateDeviceControllabilitySettings, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceControllabilitySettings", err,
				"Failure at UpdateDeviceControllabilitySettings, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateDeviceControllabilitySettings", err))
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
					"Failure when executing UpdateDeviceControllabilitySettings", err1))
				return diags
			}
		}

	}

	return resourceNetworkDevicesDeviceControllabilitySettingsRead(ctx, d, m)
}

func resourceNetworkDevicesDeviceControllabilitySettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDevicesDeviceControllabilitySettings", err, "Delete method is not supported",
		"Failure at NetworkDevicesDeviceControllabilitySettingsDelete, unexpected response", ""))
	return diags
}
func expandRequestNetworkDevicesDeviceControllabilitySettingsUpdateDeviceControllabilitySettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateDeviceControllabilitySettings {
	request := dnacentersdkgo.RequestSiteDesignUpdateDeviceControllabilitySettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".autocorrect_telemetry_config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".autocorrect_telemetry_config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".autocorrect_telemetry_config")))) {
		request.AutocorrectTelemetryConfig = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_controllability")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_controllability")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_controllability")))) {
		request.DeviceControllability = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
