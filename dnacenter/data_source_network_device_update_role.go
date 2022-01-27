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
func dataSourceNetworkDeviceUpdateRole() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.

- Updates the role of the device as access, core, distribution, border router
`,

		ReadContext: dataSourceNetworkDeviceUpdateRoleRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
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
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceNetworkDeviceUpdateRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateDeviceRole")
		request1 := expandRequestNetworkDeviceUpdateRoleUpdateDeviceRole(ctx, "", d)

		response1, restyResp1, err := client.Devices.UpdateDeviceRole(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceRole", err,
				"Failure at UpdateDeviceRole, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesUpdateDeviceRoleItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateDeviceRole response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNetworkDeviceUpdateRoleUpdateDeviceRole(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateDeviceRole {
	request := dnacentersdkgo.RequestDevicesUpdateDeviceRole{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role_source")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role_source")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role_source")))) {
		request.RoleSource = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDevicesUpdateDeviceRoleItem(item *dnacentersdkgo.ResponseDevicesUpdateDeviceRoleResponse) []map[string]interface{} {
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
