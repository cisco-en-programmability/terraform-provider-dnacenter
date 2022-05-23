package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesIDsPerDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Security Advisories.

- Retrieves list of advisory IDs for a device
`,

		ReadContext: dataSourceSecurityAdvisoriesIDsPerDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device instance UUID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advisory_ids": &schema.Schema{
							Description: `Advisory Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityAdvisoriesIDsPerDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAdvisoryIDsPerDevice")
		vvDeviceID := vDeviceID.(string)

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoryIDsPerDevice(vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAdvisoryIDsPerDevice", err,
				"Failure at GetAdvisoryIDsPerDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSecurityAdvisoriesGetAdvisoryIDsPerDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoryIDsPerDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoryIDsPerDeviceItems(items *[]dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoryIDsPerDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_id"] = item.DeviceID
		respItem["advisory_ids"] = item.AdvisoryIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}
