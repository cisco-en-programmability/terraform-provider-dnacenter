package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLanAutomationUpdateV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on LAN Automation.

- Invoke this API to stop LAN Automation and update device parameters such as Loopback0 IP address and/or hostname
discovered in the current session.
`,

		CreateContext: resourceLanAutomationUpdateV2Create,
		ReadContext:   resourceLanAutomationUpdateV2Read,
		DeleteContext: resourceLanAutomationUpdateV2Delete,
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
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `url to check the status of task
`,
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
						"id": &schema.Schema{
							Description: `id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestLanAutomationLANAutomationStopAndUpdateDevicesV2`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_management_ipaddress": &schema.Schema{
										Description: `Device Management IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"new_host_name": &schema.Schema{
										Description: `New hostname to be assigned to the device
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"new_loopback0_ipaddress": &schema.Schema{
										Description: `New Loopback0 IP Address from LAN pool of Device Discovery Site.
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
				},
			},
		},
	}
}

func resourceLanAutomationUpdateV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.LanAutomation.LanAutomationStopAndUpdateDevicesV2(vvID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing LanAutomationStopAndUpdateDevicesV2", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing LANAutomationStopAndUpdateDevicesV2", err))
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
				"Failure when executing LANAutomationStopAndUpdateDevicesV2", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenLanAutomationLanAutomationStopAndUpdateDevicesV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LanAutomationStopAndUpdateDevicesV2 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceLanAutomationUpdateV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLanAutomationUpdateV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLanAutomationLanAutomationStopAndUpdateDevicesV2 {
	request := dnacentersdkgo.RequestLanAutomationLanAutomationStopAndUpdateDevicesV2{}
	if v := expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 {
	request := []dnacentersdkgo.RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2{}
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
		i := expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationUpdateV2LanAutomationStopAndUpdateDevicesV2Item(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 {
	request := dnacentersdkgo.RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ipaddress")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_loopback0_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_loopback0_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_loopback0_ipaddress")))) {
		request.NewLoopback0IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_host_name")))) {
		request.NewHostName = interfaceToString(v)
	}
	return &request
}

func flattenLanAutomationLanAutomationStopAndUpdateDevicesV2Item(item *dnacentersdkgo.ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2Response) []map[string]interface{} {
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
