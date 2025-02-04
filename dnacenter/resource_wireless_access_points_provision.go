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
func resourceWirelessAccessPointsProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action is used to provision access points
`,

		CreateContext: resourceWirelessAccessPointsProvisionCreate,
		ReadContext:   resourceWirelessAccessPointsProvisionRead,
		DeleteContext: resourceWirelessAccessPointsProvisionDelete,
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
							Description: `Task URL
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
						"ap_zone_name": &schema.Schema{
							Description: `AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"network_devices": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_id": &schema.Schema{
										Description: `Network device ID of access points
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mesh_role": &schema.Schema{
										Description: `Mesh Role (Applicable only when AP is in Bridge Mode)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name. RF Profile is not allowed for custom AP Zones.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `Site ID
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

func resourceWirelessAccessPointsProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessAccessPointsProvisionApProvision(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.ApProvision(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApProvision", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing APProvision", err))
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
				"Failure when executing APProvision", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessApProvisionItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApProvision response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessAccessPointsProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessAccessPointsProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessAccessPointsProvisionApProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessApProvision {
	request := dnacentersdkgo.RequestWirelessApProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_devices")))) {
		request.NetworkDevices = expandRequestWirelessAccessPointsProvisionApProvisionNetworkDevicesArray(ctx, key+".network_devices", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile_name")))) {
		request.RfProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_zone_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_zone_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_zone_name")))) {
		request.ApZoneName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccessPointsProvisionApProvisionNetworkDevicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessApProvisionNetworkDevices {
	request := []dnacentersdkgo.RequestWirelessApProvisionNetworkDevices{}
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
		i := expandRequestWirelessAccessPointsProvisionApProvisionNetworkDevices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessAccessPointsProvisionApProvisionNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessApProvisionNetworkDevices {
	request := dnacentersdkgo.RequestWirelessApProvisionNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mesh_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mesh_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mesh_role")))) {
		request.MeshRole = interfaceToString(v)
	}
	return &request
}

func flattenWirelessApProvisionItem(item *dnacentersdkgo.ResponseWirelessApProvisionResponse) []map[string]interface{} {
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
