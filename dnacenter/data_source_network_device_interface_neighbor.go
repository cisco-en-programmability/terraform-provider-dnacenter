package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceInterfaceNeighbor() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get connected device detail for given deviceUuid and interfaceUuid
`,

		ReadContext: dataSourceNetworkDeviceInterfaceNeighborRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. instanceuuid of Device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"interface_uuid": &schema.Schema{
				Description: `interfaceUuid path parameter. instanceuuid of interface
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"capabilities": &schema.Schema{
							Description: `Info about capabilities of the connected device
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"neighbor_device": &schema.Schema{
							Description: `Info about the devices connected to the interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"neighbor_port": &schema.Schema{
							Description: `Info about the connected interface
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

func dataSourceNetworkDeviceInterfaceNeighborRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")
	vInterfaceUUID := d.Get("interface_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConnectedDeviceDetail")
		vvDeviceUUID := vDeviceUUID.(string)
		vvInterfaceUUID := vInterfaceUUID.(string)

		response1, restyResp1, err := client.Devices.GetConnectedDeviceDetail(vvDeviceUUID, vvInterfaceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConnectedDeviceDetail", err,
				"Failure at GetConnectedDeviceDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetConnectedDeviceDetailItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectedDeviceDetail response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetConnectedDeviceDetailItem(item *dnacentersdkgo.ResponseDevicesGetConnectedDeviceDetailResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["neighbor_device"] = item.NeighborDevice
	respItem["neighbor_port"] = item.NeighborPort
	respItem["capabilities"] = item.Capabilities
	return []map[string]interface{}{
		respItem,
	}
}
