package dnacenter

import (
	"context"
	"errors"

	"time"

	"reflect"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpDeviceConfigPreview() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Triggers a preview for site-based Day 0 Configuration
`,

		CreateContext: resourcePnpDeviceConfigPreviewCreate,
		ReadContext:   resourcePnpDeviceConfigPreviewRead,
		DeleteContext: resourcePnpDeviceConfigPreviewDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"site_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpDeviceConfigPreviewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceConfigPreviewPreviewConfig(ctx, "parameters.0", d)

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

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing PreviewConfig", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing PreviewConfig", err1))
			return diags
		}
	}

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
func resourcePnpDeviceConfigPreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceConfigPreviewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
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
