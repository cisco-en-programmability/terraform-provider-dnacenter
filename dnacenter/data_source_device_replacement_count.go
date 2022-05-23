package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceReplacementCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Replacement.

- Get replacement devices count
`,

		ReadContext: dataSourceDeviceReplacementCountRead,
		Schema: map[string]*schema.Schema{
			"replacement_status": &schema.Schema{
				Description: `replacementStatus query parameter. Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
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

func dataSourceDeviceReplacementCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vReplacementStatus, okReplacementStatus := d.GetOk("replacement_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ReturnReplacementDevicesCount")
		queryParams1 := dnacentersdkgo.ReturnReplacementDevicesCountQueryParams{}

		if okReplacementStatus {
			queryParams1.ReplacementStatus = interfaceToSliceString(vReplacementStatus)
		}

		response1, restyResp1, err := client.DeviceReplacement.ReturnReplacementDevicesCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReturnReplacementDevicesCount", err,
				"Failure at ReturnReplacementDevicesCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceReplacementReturnReplacementDevicesCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnReplacementDevicesCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceReplacementReturnReplacementDevicesCountItem(item *dnacentersdkgo.ResponseDeviceReplacementReturnReplacementDevicesCount) []map[string]interface{} {
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
