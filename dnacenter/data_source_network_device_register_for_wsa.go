package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceRegisterForWsa() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- It fetches devices which are registered to receive WSA notifications. The device serial number and/or MAC address are
required to be provided as query parameters.
`,

		ReadContext: dataSourceNetworkDeviceRegisterForWsaRead,
		Schema: map[string]*schema.Schema{
			"macaddress": &schema.Schema{
				Description: `macaddress query parameter. Mac addres of the device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Serial number of the device
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"model_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceRegisterForWsaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vMacaddress, okMacaddress := d.GetOk("macaddress")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDevicesRegisteredForWsaNotification")
		queryParams1 := dnacentersdkgo.GetDevicesRegisteredForWsaNotificationQueryParams{}

		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okMacaddress {
			queryParams1.Macaddress = vMacaddress.(string)
		}

		response1, restyResp1, err := client.Devices.GetDevicesRegisteredForWsaNotification(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDevicesRegisteredForWsaNotification", err,
				"Failure at GetDevicesRegisteredForWsaNotification, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDevicesRegisteredForWsaNotificationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDevicesRegisteredForWsaNotification response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDevicesRegisteredForWsaNotificationItem(item *dnacentersdkgo.ResponseDevicesGetDevicesRegisteredForWsaNotificationResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mac_address"] = item.MacAddress
	respItem["model_number"] = item.ModelNumber
	respItem["name"] = item.Name
	respItem["serial_number"] = item.SerialNumber
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}
