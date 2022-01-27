package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceSupervisorCardDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.
`,

		ReadContext: dataSourceNetworkDeviceSupervisorCardDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. instanceuuid of device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"partno": &schema.Schema{
							Description: `Partno`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"serialno": &schema.Schema{
							Description: `Serialno`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"slotno": &schema.Schema{
							Description: `Slotno`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"switchno": &schema.Schema{
							Description: `Switchno`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceSupervisorCardDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSupervisorCardDetail")
		vvDeviceUUID := vDeviceUUID.(string)

		response1, restyResp1, err := client.Devices.GetSupervisorCardDetail(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSupervisorCardDetail", err,
				"Failure at GetSupervisorCardDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetSupervisorCardDetailItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSupervisorCardDetail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetSupervisorCardDetailItems(items *[]dnacentersdkgo.ResponseDevicesGetSupervisorCardDetailResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["serialno"] = item.Serialno
		respItem["partno"] = item.Partno
		respItem["switchno"] = item.Switchno
		respItem["slotno"] = item.Slotno
		respItems = append(respItems, respItem)
	}
	return respItems
}
