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
func dataSourceDeviceReplacementDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Replacement.

- API to trigger RMA workflow that will replace faulty device with replacement device with same configuration and images
`,

		ReadContext: dataSourceDeviceReplacementDeployRead,
		Schema: map[string]*schema.Schema{
			"faulty_device_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"replacement_device_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceDeviceReplacementDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeployDeviceReplacementWorkflow")
		request1 := expandRequestDeviceReplacementDeployDeployDeviceReplacementWorkflow(ctx, "", d)

		response1, restyResp1, err := client.DeviceReplacement.DeployDeviceReplacementWorkflow(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeployDeviceReplacementWorkflow", err,
				"Failure at DeployDeviceReplacementWorkflow, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceReplacementDeployDeviceReplacementWorkflowItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeployDeviceReplacementWorkflow response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestDeviceReplacementDeployDeployDeviceReplacementWorkflow(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementDeployDeviceReplacementWorkflow {
	request := dnacentersdkgo.RequestDeviceReplacementDeployDeviceReplacementWorkflow{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDeviceReplacementDeployDeviceReplacementWorkflowItem(item *dnacentersdkgo.ResponseDeviceReplacementDeployDeviceReplacementWorkflowResponse) []map[string]interface{} {
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
