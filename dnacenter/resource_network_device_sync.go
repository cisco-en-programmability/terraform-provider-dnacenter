package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.
- Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If
forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result
can be seen in the child task of each device
`,

		CreateContext: resourceNetworkDeviceSyncCreate,
		ReadContext:   resourceNetworkDeviceSyncRead,
		DeleteContext: resourceNetworkDeviceSyncDelete,

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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_sync": &schema.Schema{
							Description: `forceSync query parameter.`,
							Type:        schema.TypeBool,
							ForceNew:    true,
							Optional:    true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestDevicesSyncDevices`,
							Type:        schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							ForceNew: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceSyncCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vForceSync, okForceSync := resourceItem["force_sync"]

	request1 := expandRequestNetworkDeviceSyncSyncDevices(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	queryParams1 := dnacentersdkgo.SyncDevicesQueryParams{}
	if okForceSync {
		queryParams1.ForceSync = vForceSync.(bool)
	}
	response1, restyResp1, err := client.Devices.SyncDevices(request1, &queryParams1)
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
				"Failure when executing SyncDevices", err1))
			return diags
		}
	}
	d.SetId(getUnixTimeString())
	return diags
}

func resourceNetworkDeviceSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDeviceSyncUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDeviceSyncRead(ctx, d, m)
}

func resourceNetworkDeviceSyncDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
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
	for item_no := range objs {
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
