package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricDevicesLayer2Handoffs() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Deletes layer 2 handoffs of a fabric device based on user input.

- Adds layer 2 handoffs in fabric devices based on user input.

- Deletes a layer 2 handoff of a fabric device based on id.
`,

		CreateContext: resourceSdaFabricDevicesLayer2HandoffsCreate,
		ReadContext:   resourceSdaFabricDevicesLayer2HandoffsRead,
		UpdateContext: resourceSdaFabricDevicesLayer2HandoffsUpdate,
		DeleteContext: resourceSdaFabricDevicesLayer2HandoffsDelete,
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

						"external_vlan_id": &schema.Schema{
							Description: `External VLAN number into which the fabric is extended.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the layer 2 handoff of a fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface name of the layer 2 handoff.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_vlan_id": &schema.Schema{
							Description: `VLAN number associated with this fabric.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device ID of the fabric device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddFabricDevicesLayer2Handoffs`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external_vlan_id": &schema.Schema{
										Description: `External VLAN number into which the fabric must be extended. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this device is assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `id path parameter. ID of the layer 2 handoff of a fabric device.
`,
										Type:     schema.TypeString,
										Required: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name of the layer 2 handoff. E.g., GigabitEthernet1/0/4
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"internal_vlan_id": &schema.Schema{
										Description: `VLAN number associated with this fabric. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device ID of the fabric device.
`,
										Type:     schema.TypeString,
										Optional: true,
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

func resourceSdaFabricDevicesLayer2HandoffsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2Handoffs(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	vInterfaceName := resourceItem["interface_name"]
	vvInterfaceName := interfaceToString(vInterfaceName)
	queryParamImport := dnacentersdkgo.GetFabricDevicesLayer2HandoffsQueryParams{}
	queryParamImport.FabricID = vvFabricID
	item2, err := searchSdaGetFabricDevicesLayer2Handoffs(m, queryParamImport, vvID, vvInterfaceName)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["fabric_id"] = vvFabricID
		resourceMap["id"] = item2.ID
		resourceMap["name"] = item2.InterfaceName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricDevicesLayer2HandoffsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddFabricDevicesLayer2Handoffs(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabricDevicesLayer2Handoffs", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevicesLayer2Handoffs", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddFabricDevicesLayer2Handoffs", err))
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
				"Failure when executing AddFabricDevicesLayer2Handoffs", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetFabricDevicesLayer2HandoffsQueryParams{}
	queryParamValidate.FabricID = vvFabricID
	item3, err := searchSdaGetFabricDevicesLayer2Handoffs(m, queryParamValidate, vvID, vvInterfaceName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddFabricDevicesLayer2Handoffs", err,
			"Failure at AddFabricDevicesLayer2Handoffs, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["fabric_id"] = vvFabricID
	resourceMap["id"] = item2.ID
	resourceMap["name"] = item2.InterfaceName

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricDevicesLayer2HandoffsRead(ctx, d, m)
}

func resourceSdaFabricDevicesLayer2HandoffsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]
	vName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricDevicesLayer2Handoffs")
		queryParams1 := dnacentersdkgo.GetFabricDevicesLayer2HandoffsQueryParams{}

		queryParams1.FabricID = vFabricID

		item1, err := searchSdaGetFabricDevicesLayer2Handoffs(m, queryParams1, vvID, vName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used

		items := []dnacentersdkgo.ResponseSdaGetFabricDevicesLayer2HandoffsResponse{
			*item1,
		}

		vItem1 := flattenSdaGetFabricDevicesLayer2HandoffsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricDevicesLayer2Handoffs search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricDevicesLayer2HandoffsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricDevicesLayer2HandoffsRead(ctx, d, m)
}

func resourceSdaFabricDevicesLayer2HandoffsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteFabricDeviceLayer2HandoffByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFabricDeviceLayer2HandoffByID", err, restyResp1.String(),
				"Failure at DeleteFabricDeviceLayer2HandoffByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFabricDeviceLayer2HandoffByID", err,
			"Failure at DeleteFabricDeviceLayer2HandoffByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricDeviceLayer2HandoffByID", err))
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
				"Failure when executing DeleteFabricDeviceLayer2HandoffByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2Handoffs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddFabricDevicesLayer2Handoffs {
	request := dnacentersdkgo.RequestSdaAddFabricDevicesLayer2Handoffs{}
	if v := expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2HandoffsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2HandoffsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer2Handoffs {
	request := []dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer2Handoffs{}
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
		i := expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2HandoffsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricDevicesLayer2HandoffsAddFabricDevicesLayer2HandoffsItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer2Handoffs {
	request := dnacentersdkgo.RequestItemSdaAddFabricDevicesLayer2Handoffs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_vlan_id")))) {
		request.InternalVLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_vlan_id")))) {
		request.ExternalVLANID = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetFabricDevicesLayer2Handoffs(m interface{}, queryParams dnacentersdkgo.GetFabricDevicesLayer2HandoffsQueryParams, vID string, vInterfaceName string) (*dnacentersdkgo.ResponseSdaGetFabricDevicesLayer2HandoffsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetFabricDevicesLayer2HandoffsResponse

	queryParams.Offset = 1
	nResponse, _, err := client.Sda.GetFabricDevicesLayer2Handoffs(&queryParams)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.ID || vInterfaceName == item.InterfaceName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Sda.GetFabricDevicesLayer2Handoffs(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
