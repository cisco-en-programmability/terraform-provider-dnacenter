package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersAnchorCapableDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get Anchor capable devices
`,

		ReadContext: dataSourceWirelessControllersAnchorCapableDevicesRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_ip": &schema.Schema{
							Description: `Anchor Controller Ip
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_name": &schema.Schema{
							Description: `Anchor Controller host name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_mgmt_ip": &schema.Schema{
							Description: `Wireless management Ip Address
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessControllersAnchorCapableDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnchorCapableDevices")

		response1, restyResp1, err := client.Wireless.GetAnchorCapableDevices()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAnchorCapableDevices", err,
				"Failure at GetAnchorCapableDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAnchorCapableDevicesItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnchorCapableDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAnchorCapableDevicesItem(item *dnacentersdkgo.ResponseWirelessGetAnchorCapableDevices) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_ip"] = item.DeviceIP
	respItem["device_name"] = item.DeviceName
	respItem["wireless_mgmt_ip"] = item.WirelessMgmtIP
	return []map[string]interface{}{
		respItem,
	}
}
