package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the config for all devices

- Returns the device config by specified device ID
`,

		ReadContext: dataSourceNetworkDeviceConfigRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"cdp_neighbors": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"health_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"intf_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_intf_brief": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address_table": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"running_config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okNetworkDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceConfigForAllDevices")

		response1, restyResp1, err := client.Devices.GetDeviceConfigForAllDevices()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceConfigForAllDevices", err,
				"Failure at GetDeviceConfigForAllDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceConfigForAllDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceConfigForAllDevices response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetDeviceConfigByID")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response2, restyResp2, err := client.Devices.GetDeviceConfigByID(vvNetworkDeviceID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceConfigByID", err,
				"Failure at GetDeviceConfigByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetDeviceConfigByIDItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceConfigByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceConfigForAllDevicesItems(items *[]dnacentersdkgo.ResponseDevicesGetDeviceConfigForAllDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetDeviceConfigForAllDevicesItemsAttributeInfo(item.AttributeInfo)
		respItem["cdp_neighbors"] = item.CdpNeighbors
		respItem["health_monitor"] = item.HealthMonitor
		respItem["id"] = item.ID
		respItem["intf_description"] = item.IntfDescription
		respItem["inventory"] = item.Inventory
		respItem["ip_intf_brief"] = item.IPIntfBrief
		respItem["mac_address_table"] = item.MacAddressTable
		respItem["running_config"] = item.RunningConfig
		respItem["snmp"] = item.SNMP
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceConfigForAllDevicesItemsAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceConfigByIDItem(item *dnacentersdkgo.ResponseDevicesGetDeviceConfigByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
