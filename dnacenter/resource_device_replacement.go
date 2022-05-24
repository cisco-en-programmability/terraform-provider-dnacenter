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
			"parameters": &schema.Schema{
				Description: `Array of RequestDeviceReplacementMarkDeviceForReplacement`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_platform": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"faulty_device_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"neighbour_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_readiness_task_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_device_platform": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_device_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"workflow_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceDeviceReplacementCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceReplacementMarkDeviceForReplacement(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

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
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceDeviceReplacementRead(ctx, d, m)
}

func resourceDeviceReplacementRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFaultyDeviceName := resourceMap["faulty_device_name"]
	vFaultyDevicePlatform := resourceMap["faulty_device_platform"]
	vReplacementDevicePlatform := resourceMap["replacement_device_platform"]
	vFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]
	vReplacementStatus := resourceMap["replacement_status"]
	vFamily := resourceMap["family"]
	vSortBy := resourceMap["sort_by"]
	vSortOrder := resourceMap["sort_order"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnListOfReplacementDevicesWithReplacementDetails")
		queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams{}

		if okFaultyDeviceName {
			queryParams1.FaultyDeviceName = vFaultyDeviceName
		}
		if okFaultyDevicePlatform {
			queryParams1.FaultyDevicePlatform = vFaultyDevicePlatform
		}
		if okReplacementDevicePlatform {
			queryParams1.ReplacementDevicePlatform = vReplacementDevicePlatform
		}
		if okFaultyDeviceSerialNumber {
			queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber
		}
		if okReplacementDeviceSerialNumber {
			queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber
		}
		if okReplacementStatus {
			queryParams1.ReplacementStatus = interfaceToSliceString(vReplacementStatus)
		}
		if okFamily {
			queryParams1.Family = interfaceToSliceString(vFamily)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder
		}
		if okOffset {
			queryParams1.Offset = *stringToIntPtr(vOffset)
		}
		if okLimit {
			queryParams1.Limit = *stringToIntPtr(vLimit)
		}

		response1, restyResp1, err := client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReturnListOfReplacementDevicesWithReplacementDetails", err,
				"Failure at ReturnListOfReplacementDevicesWithReplacementDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, response1, nil)
		item1, err := searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from ReturnListOfReplacementDevicesWithReplacementDetails response", err,
				"Failure when searching item from ReturnListOfReplacementDevicesWithReplacementDetails, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsByIDItem(item1)
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
	vFaultyDeviceName := resourceMap["faulty_device_name"]
	vFaultyDevicePlatform := resourceMap["faulty_device_platform"]
	vReplacementDevicePlatform := resourceMap["replacement_device_platform"]
	vFaultyDeviceSerialNumber := resourceMap["faulty_device_serial_number"]
	vReplacementDeviceSerialNumber := resourceMap["replacement_device_serial_number"]
	vReplacementStatus := resourceMap["replacement_status"]
	vFamily := resourceMap["family"]
	vSortBy := resourceMap["sort_by"]
	vSortOrder := resourceMap["sort_order"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	queryParams1 := dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams
	queryParams1.FaultyDeviceName = vFaultyDeviceName
	queryParams1.FaultyDevicePlatform = vFaultyDevicePlatform
	queryParams1.ReplacementDevicePlatform = vReplacementDevicePlatform
	queryParams1.FaultyDeviceSerialNumber = vFaultyDeviceSerialNumber
	queryParams1.ReplacementDeviceSerialNumber = vReplacementDeviceSerialNumber
	queryParams1.ReplacementStatus = interfaceToSliceString(vReplacementStatus)
	queryParams1.Family = interfaceToSliceString(vFamily)
	queryParams1.SortBy = vSortBy
	queryParams1.SortOrder = vSortOrder
	queryParams1.Offset = *stringToIntPtr(vOffset)
	queryParams1.Limit = *stringToIntPtr(vLimit)
	item, err := searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ReturnListOfReplacementDevicesWithReplacementDetails", err,
			"Failure at ReturnListOfReplacementDevicesWithReplacementDetails, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestDeviceReplacementUnmarkDeviceForReplacement(ctx, "parameters", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
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

func searchDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails(m interface{}, queryParams dnacentersdkgo.ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams) (*dnacentersdkgo.ResponseItemDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails
	var ite *dnacentersdkgo.ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails
	ite, _, err = client.DeviceReplacement.ReturnListOfReplacementDevicesWithReplacementDetails(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
