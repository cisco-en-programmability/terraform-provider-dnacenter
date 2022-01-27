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
func dataSourceCommandRunnerRunCommand() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Command Runner.

- Submit request for read-only CLIs
`,

		ReadContext: dataSourceCommandRunnerRunCommandRead,
		Schema: map[string]*schema.Schema{
			"commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuids": &schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataSourceCommandRunnerRunCommandRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration")
		request1 := expandRequestCommandRunnerRunCommandRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(ctx, "", d)

		response1, restyResp1, err := client.CommandRunner.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration", err,
				"Failure at RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestCommandRunnerRunCommandRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration {
	request := dnacentersdkgo.RequestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timeout")))) {
		request.Timeout = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationItem(item *dnacentersdkgo.ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse) []map[string]interface{} {
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
