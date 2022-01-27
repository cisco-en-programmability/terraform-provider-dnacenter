package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceNetworkDeviceSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.

- Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If
forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result
can be seen in the child task of each device
`,

		ReadContext: dataSourceNetworkDeviceSyncRead,
		Schema: map[string]*schema.Schema{
			"force_sync": &schema.Schema{
				Description: `forceSync query parameter.`,
				Type:        schema.TypeBool,
				Optional:    true,
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
			"payload": &schema.Schema{
				Description: `Array of RequestDevicesSyncDevices`,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func dataSourceNetworkDeviceSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vForceSync, okForceSync := d.GetOk("force_sync")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SyncDevices")
		request1 := expandRequestNetworkDeviceSyncSyncDevices(ctx, "", d)
		queryParams1 := dnacentersdkgo.SyncDevicesQueryParams{}

		if okForceSync {
			queryParams1.ForceSync = vForceSync.(bool)
		}

		response1, restyResp1, err := client.Devices.SyncDevices(request1, &queryParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SyncDevices", err,
				"Failure at SyncDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesSyncDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SyncDevices response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNetworkDeviceSyncSyncDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesSyncDevices {
	request := dnacentersdkgo.RequestDevicesSyncDevices{}
	if v := expandRequestNetworkDeviceSyncSyncDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceSyncSyncDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDevicesSyncDevices {
	request := []dnacentersdkgo.RequestItemDevicesSyncDevices{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestNetworkDeviceSyncSyncDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestNetworkDeviceSyncSyncDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDevicesSyncDevices {
	var request dnacentersdkgo.RequestItemDevicesSyncDevices
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenDevicesSyncDevicesItem(item *dnacentersdkgo.ResponseDevicesSyncDevicesResponse) []map[string]interface{} {
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
