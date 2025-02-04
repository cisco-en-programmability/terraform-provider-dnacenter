package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceWirelessLan() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the wireless lan controller info with given device ID
`,

		ReadContext: dataSourceNetworkDeviceWirelessLanRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_enabled_ports": &schema.Schema{
							Description: `Admin Enabled Ports of the Device
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						"ap_group_name": &schema.Schema{
							Description: `Name of the AP Group that Access point assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"eth_mac_address": &schema.Schema{
							Description: `Ethernet MacAddress of the Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"flex_group_name": &schema.Schema{
							Description: `Name of the Flex Group that Access point assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `TenantId of the Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance UUID of the Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"lag_mode_enabled": &schema.Schema{
							Description: `LagMode status of the Device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_enabled": &schema.Schema{
							Description: `Netconf Status of the Device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_license_info": &schema.Schema{
							Description: `License type of Wireless Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_package_installed": &schema.Schema{
							Description: `Status of the Wireless Package on the Device
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

func dataSourceNetworkDeviceWirelessLanRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWirelessLanControllerDetailsByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Devices.GetWirelessLanControllerDetailsByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessLanControllerDetailsByID", err,
				"Failure at GetWirelessLanControllerDetailsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetWirelessLanControllerDetailsByIDItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessLanControllerDetailsByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetWirelessLanControllerDetailsByIDItem(item *dnacentersdkgo.ResponseDevicesGetWirelessLanControllerDetailsByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_enabled_ports"] = item.AdminEnabledPorts
	respItem["ap_group_name"] = item.ApGroupName
	respItem["device_id"] = item.DeviceID
	respItem["eth_mac_address"] = item.EthMacAddress
	respItem["flex_group_name"] = item.FlexGroupName
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["lag_mode_enabled"] = boolPtrToString(item.LagModeEnabled)
	respItem["netconf_enabled"] = boolPtrToString(item.NetconfEnabled)
	respItem["wireless_license_info"] = item.WirelessLicenseInfo
	respItem["wireless_package_installed"] = boolPtrToString(item.WirelessPackageInstalled)
	return []map[string]interface{}{
		respItem,
	}
}
