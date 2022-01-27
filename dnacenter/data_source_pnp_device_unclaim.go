package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourcePnpDeviceUnclaim() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Un-Claims one of more devices with specified workflow
`,

		ReadContext: dataSourcePnpDeviceUnclaimRead,
		Schema: map[string]*schema.Schema{
			"device_id_list": &schema.Schema{
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

						"json_array_response": &schema.Schema{
							Description: `Json Array Response`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Description: `Json Response`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_code": &schema.Schema{
							Description: `Status Code`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpDeviceUnclaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UnClaimDevice")
		request1 := expandRequestPnpDeviceUnclaimUnClaimDevice(ctx, "", d)

		response1, restyResp1, err := client.DeviceOnboardingPnp.UnClaimDevice(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UnClaimDevice", err,
				"Failure at UnClaimDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpUnClaimDeviceItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UnClaimDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPnpDeviceUnclaimUnClaimDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpUnClaimDevice {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpUnClaimDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id_list")))) {
		request.DeviceIDList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDeviceOnboardingPnpUnClaimDeviceItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpUnClaimDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["json_array_response"] = flattenDeviceOnboardingPnpUnClaimDeviceItemJSONArrayResponse(item.JSONArrayResponse)
	respItem["json_response"] = flattenDeviceOnboardingPnpUnClaimDeviceItemJSONResponse(item.JSONResponse)
	respItem["message"] = item.Message
	respItem["status_code"] = item.StatusCode
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpUnClaimDeviceItemJSONArrayResponse(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpUnClaimDeviceJSONArrayResponse) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenDeviceOnboardingPnpUnClaimDeviceItemJSONResponse(item *dnacentersdkgo.ResponseDeviceOnboardingPnpUnClaimDeviceJSONResponse) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
