package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceInterfacePoe() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns POE interface details for the device, where deviceuuid is mandatory & accepts comma seperated interface names
which is optional and returns information for that particular interfaces where(operStatus = operationalStatus)
`,

		ReadContext: dataSourceNetworkDeviceInterfacePoeRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. uuid of the device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"interface_name_list": &schema.Schema{
				Description: `interfaceNameList query parameter. comma seperated interface names
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Description: `Admin Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"allocated_power": &schema.Schema{
							Description: `Allocated Power`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"max_port_power": &schema.Schema{
							Description: `Max Port Power`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"oper_status": &schema.Schema{
							Description: `Oper Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"port_power_drawn": &schema.Schema{
							Description: `Port Power Drawn`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceInterfacePoeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")
	vInterfaceNameList, okInterfaceNameList := d.GetOk("interface_name_list")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ReturnsPoeInterfaceDetailsForTheDevice")
		vvDeviceUUID := vDeviceUUID.(string)
		queryParams1 := dnacentersdkgo.ReturnsPoeInterfaceDetailsForTheDeviceQueryParams{}

		if okInterfaceNameList {
			queryParams1.InterfaceNameList = vInterfaceNameList.(string)
		}

		response1, restyResp1, err := client.Devices.ReturnsPoeInterfaceDetailsForTheDevice(vvDeviceUUID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReturnsPoeInterfaceDetailsForTheDevice", err,
				"Failure at ReturnsPoeInterfaceDetailsForTheDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesReturnsPoeInterfaceDetailsForTheDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsPoeInterfaceDetailsForTheDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesReturnsPoeInterfaceDetailsForTheDeviceItems(items *[]dnacentersdkgo.ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["admin_status"] = item.AdminStatus
		respItem["oper_status"] = item.OperStatus
		respItem["interface_name"] = item.InterfaceName
		respItem["max_port_power"] = item.MaxPortPower
		respItem["allocated_power"] = item.AllocatedPower
		respItem["port_power_drawn"] = item.PortPowerDrawn
		respItems = append(respItems, respItem)
	}
	return respItems
}
