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
func dataSourcePnpDeviceConfigPreview() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Triggers a preview for site-based Day 0 Configuration
`,

		ReadContext: dataSourcePnpDeviceConfigPreviewRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"complete": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"expired_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rf_profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sensor_profile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePnpDeviceConfigPreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: PreviewConfig")
		request1 := expandRequestPnpDeviceConfigPreviewPreviewConfig(ctx, "", d)

		response1, restyResp1, err := client.DeviceOnboardingPnp.PreviewConfig(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PreviewConfig", err,
				"Failure at PreviewConfig, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpPreviewConfigItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting PreviewConfig response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestPnpDeviceConfigPreviewPreviewConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceOnboardingPnpPreviewConfig {
	request := dnacentersdkgo.RequestDeviceOnboardingPnpPreviewConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDeviceOnboardingPnpPreviewConfigItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpPreviewConfigResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["complete"] = boolPtrToString(item.Complete)
	respItem["config"] = item.Config
	respItem["error"] = boolPtrToString(item.Error)
	respItem["error_message"] = item.ErrorMessage
	respItem["expired_time"] = item.ExpiredTime
	respItem["rf_profile"] = item.RfProfile
	respItem["sensor_profile"] = item.SensorProfile
	respItem["site_id"] = item.SiteID
	respItem["start_time"] = item.StartTime
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
