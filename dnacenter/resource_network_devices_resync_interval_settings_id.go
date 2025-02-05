package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevicesResyncIntervalSettingsID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Devices.

- Update the resync interval (in minutes) for the given network device id.
To disable periodic resync, set interval as *0*.
To use global settings, set interval as *null*.
`,

		CreateContext: resourceNetworkDevicesResyncIntervalSettingsIDCreate,
		ReadContext:   resourceNetworkDevicesResyncIntervalSettingsIDRead,
		UpdateContext: resourceNetworkDevicesResyncIntervalSettingsIDUpdate,
		DeleteContext: resourceNetworkDevicesResyncIntervalSettingsIDDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"interval": &schema.Schema{
							Description: `Resync interval of the device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. The id of the network device.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"interval": &schema.Schema{
							Description: `Resync interval in minutes. To disable periodic resync, set interval as *0*. To use global settings, set interval as *null*.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDevicesResyncIntervalSettingsIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDevicesResyncIntervalSettingsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetResyncIntervalForTheNetworkDevice")
		vvID := vID

		response1, restyResp1, err := client.Devices.GetResyncIntervalForTheNetworkDevice(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetResyncIntervalForTheNetworkDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetResyncIntervalForTheNetworkDevice response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceNetworkDevicesResyncIntervalSettingsIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestNetworkDevicesResyncIntervalSettingsIDUpdateResyncIntervalForTheNetworkDevice(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Devices.UpdateResyncIntervalForTheNetworkDevice(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateResyncIntervalForTheNetworkDevice", err, restyResp1.String(),
					"Failure at UpdateResyncIntervalForTheNetworkDevice, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateResyncIntervalForTheNetworkDevice", err,
				"Failure at UpdateResyncIntervalForTheNetworkDevice, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateResyncIntervalForTheNetworkDevice", err))
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateResyncIntervalForTheNetworkDevice", err1))
				return diags
			}
		}

	}

	return resourceNetworkDevicesResyncIntervalSettingsIDRead(ctx, d, m)
}

func resourceNetworkDevicesResyncIntervalSettingsIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing NetworkDevicesResyncIntervalSettingsID", err, "Delete method is not supported",
		"Failure at NetworkDevicesResyncIntervalSettingsIDDelete, unexpected response", ""))
	return diags
}
func expandRequestNetworkDevicesResyncIntervalSettingsIDUpdateResyncIntervalForTheNetworkDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesUpdateResyncIntervalForTheNetworkDevice {
	request := dnacentersdkgo.RequestDevicesUpdateResyncIntervalForTheNetworkDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interval")))) {
		request.Interval = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
