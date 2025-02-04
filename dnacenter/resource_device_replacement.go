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

func resourceDeviceReplacement() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Device Replacement.

- UnMarks device for replacement

- Marks device for replacement
`,

		CreateContext: resourceDeviceReplacementCreate,
		ReadContext:   resourceDeviceReplacementRead,
		UpdateContext: resourceDeviceReplacementUpdate,
		DeleteContext: resourceDeviceReplacementDelete,
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

						"creation_time": &schema.Schema{
							Description: `Date and time of marking the device for replacement
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"family": &schema.Schema{
							Description: `Faulty device family
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"faulty_device_id": &schema.Schema{
							Description: `Unique identifier of the faulty device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"faulty_device_name": &schema.Schema{
							Description: `Faulty device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"faulty_device_platform": &schema.Schema{
							Description: `Faulty device platform
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"faulty_device_serial_number": &schema.Schema{
							Description: `Faulty device serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Unique identifier of the device replacement resource
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"neighbour_device_id": &schema.Schema{
							Description: `Unique identifier of the neighbor device to create the DHCP server
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_readiness_task_id": &schema.Schema{
							Description: `Unique identifier of network readiness task
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"readinesscheck_task_id": &schema.Schema{
							Description: `Unique identifier of the readiness check task for the replacement device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"replacement_device_platform": &schema.Schema{
							Description: `Replacement device platform
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"replacement_device_serial_number": &schema.Schema{
							Description: `Replacement device serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"replacement_status": &schema.Schema{
							Description: `Device Replacement status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"replacement_time": &schema.Schema{
							Description: `Date and time of device replacement
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"workflow_failed_step": &schema.Schema{
							Description: `Step in which the device replacement failed
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"workflow_id": &schema.Schema{
							Description: `Unique identifier of the device replacement workflow
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDeviceReplacementMarkDeviceForReplacement`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestDeviceReplacementMarkDeviceForReplacement`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"creation_time": &schema.Schema{
										Description: `Date and time of marking the device for replacement
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"family": &schema.Schema{
										Description: `Faulty device family
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"faulty_device_id": &schema.Schema{
										Description: `Unique identifier of the faulty device
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"faulty_device_name": &schema.Schema{
										Description: `Faulty device name
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"faulty_device_platform": &schema.Schema{
										Description: `Faulty device platform
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"faulty_device_serial_number": &schema.Schema{
										Description: `Faulty device serial number
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Unique identifier of the device replacement resource
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"neighbour_device_id": &schema.Schema{
										Description: `Unique identifier of the neighbor device to create the DHCP server
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_readiness_task_id": &schema.Schema{
										Description: `Unique identifier of network readiness task
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"replacement_device_platform": &schema.Schema{
										Description: `Replacement device platform
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"replacement_device_serial_number": &schema.Schema{
										Description: `Replacement device serial number
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"replacement_status": &schema.Schema{
										Description: `Device replacement status. Use MARKED-FOR-REPLACEMENT to mark the device for replacement.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"replacement_time": &schema.Schema{
										Description: `Date and time of device replacement
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"workflow_id": &schema.Schema{
										Description: `Unique identifier of the device replacement workflow
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								}}}},
				},
			},
		},
	}
}

func resourceDeviceReplacementCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestDeviceReplacementMarkDeviceForReplacement(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vFaultyDeviceID := resourceItem["faulty_device_id"]
	vFaultyDeviceSerialNumber := resourceItem["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber := resourceItem["replacement_device_serial_number"]

	vvFaultyDeviceID := interfaceToString(vFaultyDeviceID)
	vvFaultyDeviceSerialNumber := interfaceToString(vFaultyDeviceSerialNumber)
	vvReplacementDeviceSerialNumber := interfaceToString(vReplacementDeviceSerialNumber)

	log.Printf("[DEBUG] Selected method 1: ReturnListOfReplacementDevicesWithReplacementDetails")
	queryParamImport := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

	queryParamImport.FaultyDeviceSerialNumber = vvFaultyDeviceSerialNumber

	queryParamImport.ReplacementDeviceSerialNumber = vvReplacementDeviceSerialNumber

	item2, err := searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, queryParamImport, vvFaultyDeviceID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["faulty_device_serial_number"] = item2.FaultyDeviceSerialNumber
		resourceMap["replacement_device_serial_number"] = item2.ReplacementDeviceSerialNumber
		resourceMap["faulty_device_id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceDeviceReplacementRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.DeviceReplacement.MarkDeviceForReplacement(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing MarkDeviceForReplacement", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing MarkDeviceForReplacement", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing MarkDeviceForReplacement", err))
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
				"Failure when executing MarkDeviceForReplacement", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}
	queryParamValidate.FaultyDeviceSerialNumber = vvFaultyDeviceSerialNumber
	queryParamValidate.ReplacementDeviceSerialNumber = vvReplacementDeviceSerialNumber
	item3, err := searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, queryParamValidate, vvFaultyDeviceID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing MarkDeviceForReplacement", err,
			"Failure at MarkDeviceForReplacement, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["faulty_device_serial_number"] = item3.FaultyDeviceSerialNumber
	resourceMap["replacement_device_serial_number"] = item3.ReplacementDeviceSerialNumber
	resourceMap["faulty_device_id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceDeviceReplacementRead(ctx, d, m)
}

func resourceDeviceReplacementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFaultyDeviceID, _ := resourceMap["faulty_device_id"]
	vFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnListOfReplacementDevicesWithReplacementDetails")
		queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

		queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber

		queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber

		item1, err := searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, queryParams1, vFaultyDeviceID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		items := []dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse{
			*item1,
		}
		// Review flatten function used
		vItem1 := flattenDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnListOfReplacementDevicesWithReplacementDetails search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceDeviceReplacementUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]
	vFaultyDeviceID := resourceMap["faulty_device_id"]

	log.Printf("[DEBUG] Selected method 1: ReturnListOfReplacementDevicesWithReplacementDetails")
	queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

	if vFaultyDeviceSerialNumber != "" {
		queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber
	}
	if vReplacementDeviceSerialNumber != "" {
		queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber
	}
	if d.HasChange("parameters") {
		request1 := expandRequestDeviceReplacementUnmarkDeviceForReplacement(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vFaultyDeviceID
			request1 = &req
		}
		response1, restyResp1, err := client.DeviceReplacement.UnmarkDeviceForReplacement(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UnmarkDeviceForReplacement", err, restyResp1.String(),
					"Failure at UnmarkDeviceForReplacement, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UnmarkDeviceForReplacement", err,
				"Failure at UnmarkDeviceForReplacement, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UnmarkDeviceForReplacement", err))
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
					"Failure when executing UnmarkDeviceForReplacement", err1))
				return diags
			}
		}

	}

	return resourceDeviceReplacementRead(ctx, d, m)
}

func resourceDeviceReplacementDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete DeviceReplacement on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestDeviceReplacementMarkDeviceForReplacement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacement {
	request := dnacentersdkgo.RequestDeviceReplacementMarkDeviceForReplacement{}
	if v := expandRequestDeviceReplacementMarkDeviceForReplacementItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceReplacementMarkDeviceForReplacement {
	request := []dnacentersdkgo.RequestItemDeviceReplacementMarkDeviceForReplacement{}
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
		i := expandRequestDeviceReplacementMarkDeviceForReplacementItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceReplacementMarkDeviceForReplacementItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceReplacementMarkDeviceForReplacement {
	request := dnacentersdkgo.RequestItemDeviceReplacementMarkDeviceForReplacement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacement {
	request := dnacentersdkgo.RequestDeviceReplacementUnmarkDeviceForReplacement{}
	if v := expandRequestDeviceReplacementUnmarkDeviceForReplacementItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDeviceReplacementUnmarkDeviceForReplacement {
	request := []dnacentersdkgo.RequestItemDeviceReplacementUnmarkDeviceForReplacement{}
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
		i := expandRequestDeviceReplacementUnmarkDeviceForReplacementItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestDeviceReplacementUnmarkDeviceForReplacementItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDeviceReplacementUnmarkDeviceForReplacement {
	request := dnacentersdkgo.RequestItemDeviceReplacementUnmarkDeviceForReplacement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".creation_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".creation_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".creation_time")))) {
		request.CreationTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".family")))) {
		request.Family = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_id")))) {
		request.FaultyDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_name")))) {
		request.FaultyDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_platform")))) {
		request.FaultyDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".faulty_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".faulty_device_serial_number")))) {
		request.FaultyDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".neighbour_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".neighbour_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".neighbour_device_id")))) {
		request.NeighbourDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_readiness_task_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_readiness_task_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_readiness_task_id")))) {
		request.NetworkReadinessTaskID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_platform")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_platform")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_platform")))) {
		request.ReplacementDevicePlatform = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_device_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_device_serial_number")))) {
		request.ReplacementDeviceSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_status")))) {
		request.ReplacementStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".replacement_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".replacement_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".replacement_time")))) {
		request.ReplacementTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m interface{}, queryParams dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams, vID string) (*dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse
	var ite *dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {

			for _, item := range *nResponse.Response {
				if vID == item.FaultyDeviceID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = maxPageSize
			queryParams.Offset += maxPageSize
			nResponse, _, err = client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.FaultyDeviceName != "" {
		ite, _, err = client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.FaultyDeviceName == queryParams.FaultyDeviceName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
