package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceRebootApreboot() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Wireless.

- Users can reboot multiple access points up-to 200 at a time using this API
`,

		CreateContext: resourceDeviceRebootAprebootCreate,
		ReadContext:   resourceDeviceRebootAprebootRead,
		UpdateContext: resourceDeviceRebootAprebootUpdate,
		DeleteContext: resourceDeviceRebootAprebootDelete,
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

						"ap_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"failure_reason": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},
									"reboot_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
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

						"ap_mac_addresses": &schema.Schema{
							Description: `The ethernet MAC address of the access point.
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceDeviceRebootAprebootCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	// resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceRebootAprebootRebootAccessPoints(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	// vApMacAdresses := resourceItem["ap_mac_addresses"]
	// vvApMacAdresses := interfaceToString(vApMacAdresses)

	resp1, restyResp1, err := client.Wireless.RebootAccessPoints(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing RebootAccessPoints", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing RebootAccessPoints", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing RebootAccessPoints", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
				"Failure when executing RebootAccessPoints", err1))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["parent_task_id"] = taskId

	d.SetId(joinResourceID(resourceMap))
	return resourceDeviceRebootAprebootRead(ctx, d, m)
}

func resourceDeviceRebootAprebootRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	client := m.(*dnacentersdkgo.Client)
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vParentTaskID := resourceMap["parent_task_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointRebootTaskResult")
		queryParams1 := dnacentersdkgo.GetAccessPointRebootTaskResultQueryParams{}
		queryParams1.ParentTaskID = vParentTaskID
		response1, _, err := client.Wireless.GetAccessPointRebootTaskResult(&queryParams1)

		if err != nil || response1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenWirelessGetAccessPointRebootTaskResultItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointRebootTaskResult search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceDeviceRebootAprebootUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing DeviceRebootAprebootUpdate", err, "Update method is not supported",
		"Failure at DeviceRebootAprebootUpdate, unexpected response", ""))

	return diags
}

func resourceDeviceRebootAprebootDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing DeviceRebootAprebootDelete", err, "Delete method is not supported",
		"Failure at DeviceRebootAprebootDelete, unexpected response", ""))

	return diags
}
func expandRequestDeviceRebootAprebootRebootAccessPoints(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessRebootAccessPoints {
	request := dnacentersdkgo.RequestWirelessRebootAccessPoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_mac_addresses")))) {
		request.ApMacAddresses = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
