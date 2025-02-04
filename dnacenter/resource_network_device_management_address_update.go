package dnacenter

import (
	"context"

	"errors"

	"time"

	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDeviceManagementAddressUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.

- This is a simple PUT API to edit the management IP Address of the device.
`,

		CreateContext: resourceNetworkDeviceManagementAddressUpdateCreate,
		ReadContext:   resourceNetworkDeviceManagementAddressUpdateRead,
		DeleteContext: resourceNetworkDeviceManagementAddressUpdateDelete,
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
						"deviceid": &schema.Schema{
							Description: `deviceid path parameter. The UUID of the device whose management IP address is to be updated.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"new_ip": &schema.Schema{
							Description: `New IP Address of the device to be Updated
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceManagementAddressUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDeviceid := resourceItem["deviceid"]

	vvDeviceid := vDeviceid.(string)
	request1 := expandRequestNetworkDeviceManagementAddressUpdateUpdateDeviceManagementAddress(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.UpdateDeviceManagementAddress(vvDeviceid, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateDeviceManagementAddress", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UpdateDeviceManagementAddress", err))
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
				"Failure when executing UpdateDeviceManagementAddress", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenDevicesUpdateDeviceManagementAddressItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateDeviceManagementAddress response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceNetworkDeviceManagementAddressUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDeviceManagementAddressUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDeviceManagementAddressUpdateUpdateDeviceManagementAddress(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateDeviceManagementAddress {
	request := dnacentersdkgo.RequestDevicesUpdateDeviceManagementAddress{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_ip")))) {
		request.NewIP = interfaceToString(v)
	}
	return &request
}

func flattenDevicesUpdateDeviceManagementAddressItem(item *dnacentersdkgo.ResponseDevicesUpdateDeviceManagementAddressResponse) []map[string]interface{} {
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
