package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceApplicationVisibilityNetworkDevicesEnableAppTelemetry() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Application Policy.

- This data source action can be used to enable application telemetry feature on multiple network devices. Request
payload should include the list of network devices where application telemetry has to be enabled. For wireless
controllers, it also needs the WLAN modes / SSID details to be included for enablement.
Please note that this operation can be performed even if the feature is already enabled on the network device. It would
push the updated configurations to the network device.
This operation pushes configuration to the network devices, and is only permitted if the provisioning settings do not
mandate a config preview for application telemetry enablement. In cases where such settings are active, attempting to
use this endpoint will result in *422 Unprocessable Content* error.
`,

		CreateContext: resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryCreate,
		ReadContext:   resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryRead,
		DeleteContext: resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryDelete,
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

						"task_id": &schema.Schema{
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"network_devices": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Network device identifier
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"include_guest_ssids": &schema.Schema{
										Description: `Flag to indicate whether guest SSIDs should be included for application telemetry enablement. Applicable only for wireless devices. Default value is false.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"include_wlan_modes": &schema.Schema{
										Description: `Types of WLAN modes which needs to be included for enablement. Applicable and mandatory only for wireless devices. Available values: LOCAL or NON_LOCAL.
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevices(ctx, "parameters.0", d)

	response1, restyResp1, err := client.ApplicationPolicy.EnableApplicationTelemetryFeatureOnMultipleNetworkDevices(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing EnableApplicationTelemetryFeatureOnMultipleNetworkDevices", err))
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
				"Failure when executing EnableApplicationTelemetryFeatureOnMultipleNetworkDevices", err1))
			return diags
		}
	}

	vItem1 := flattenApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting EnableApplicationTelemetryFeatureOnMultipleNetworkDevices response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceApplicationVisibilityNetworkDevicesEnableAppTelemetryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevices {
	request := dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_devices")))) {
		request.NetworkDevices = expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevicesArray(ctx, key+".network_devices", d)
	}
	return &request
}

func expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices {
	request := []dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestApplicationVisibilityNetworkDevicesEnableAppTelemetryEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices {
	request := dnacentersdkgo.RequestApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_wlan_modes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_wlan_modes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_wlan_modes")))) {
		request.IncludeWLANModes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".include_guest_ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".include_guest_ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".include_guest_ssids")))) {
		request.IncludeGuestSSIDs = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesItem(item *dnacentersdkgo.ResponseApplicationPolicyEnableApplicationTelemetryFeatureOnMultipleNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
