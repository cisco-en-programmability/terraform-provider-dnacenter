package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesDeviceControllabilitySettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some
device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs
to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when
assigning a device to a site. If changes are made to any settings that are under the scope of this process, these
changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device
Controllability is disabled. The following device settings will be enabled as part of Device Controllability when
devices are discovered. SNMP Credentials. NETCONF Credentials. Subsequent to discovery, devices will be added to
Inventory. The following device settings will be enabled when devices are added to inventory. Cisco TrustSec (CTS)
Credentials. The following device settings will be enabled when devices are assigned to a site. Some of these settings
can be defined at a site level under Design > Network Settings > Telemetry & Wireless. Wired Endpoint Data Collection
Enablement. Controller Certificates. SNMP Trap Server Definitions. Syslog Server Definitions. Application Visibility.
Application QoS Policy. Wireless Service Assurance (WSA). Wireless Telemetry. DTLS Ciphersuite. AP Impersonation. If
Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings on
devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related
configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device. SWIM
certificate issue. IOS WLC NA certificate issue. PKCS12 certificate issue. IOS telemetry configuration issu
`,

		ReadContext: dataSourceNetworkDevicesDeviceControllabilitySettingsRead,
		Schema: map[string]*schema.Schema{

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
		},
	}
}

func dataSourceNetworkDevicesDeviceControllabilitySettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceControllabilitySettings", err,
				"Failure at GetDeviceControllabilitySettings, unexpected response", ""))
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

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetDeviceControllabilitySettingsItem(item *dnacentersdkgo.ResponseSiteDesignGetDeviceControllabilitySettingsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["autocorrect_telemetry_config"] = boolPtrToString(item.AutocorrectTelemetryConfig)
	respItem["device_controllability"] = boolPtrToString(item.DeviceControllability)
	return []map[string]interface{}{
		respItem,
	}
}
