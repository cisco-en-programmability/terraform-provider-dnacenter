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

func resourceQosDeviceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing qos device interface infos associate with network device id

- Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as
WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces

- Delete all qos device interface infos associate with network device id
`,

		CreateContext: resourceQosDeviceInterfaceCreate,
		ReadContext:   resourceQosDeviceInterfaceRead,
		UpdateContext: resourceQosDeviceInterfaceUpdate,
		DeleteContext: resourceQosDeviceInterfaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"excluded_interfaces": &schema.Schema{
							Description: `Excluded interfaces ids
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `Id of Qos device info
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device id
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"qos_device_interface_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dmvpn_remote_sites_bw": &schema.Schema{
										Description: `Dmvpn remote sites bandwidth
`,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interface_id": &schema.Schema{
										Description: `Interface id
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"label": &schema.Schema{
										Description: `SP Profile name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"role": &schema.Schema{
										Description: `Interface role
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"upload_bw": &schema.Schema{
										Description: `Upload bandwidth
`,
										Type:     schema.TypeInt,
										Optional: true,
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

func resourceQosDeviceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.ApplicationPolicy.CreateQosDeviceInterfaceInfo(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateQosDeviceInterfaceInfo", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateQosDeviceInterfaceInfo", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceQosDeviceInterfaceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetworkDeviceID := resourceMap["network_device_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetQosDeviceInterfaceInfo")
		queryParams1 := dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetQosDeviceInterfaceInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetQosDeviceInterfaceInfo", err,
				"Failure at GetQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsApplicationPolicyGetQosDeviceInterfaceInfo(m, response1, nil)
		item1, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetQosDeviceInterfaceInfo response", err,
				"Failure when searching item from GetQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenApplicationPolicyGetQosDeviceInterfaceInfoByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetQosDeviceInterfaceInfo search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceQosDeviceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetworkDeviceID := resourceMap["network_device_id"]

	queryParams1 := dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams
	queryParams1.NetworkDeviceID = vNetworkDeviceID
	item, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetQosDeviceInterfaceInfo", err,
			"Failure at GetQosDeviceInterfaceInfo, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx, "parameters", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ApplicationPolicy.UpdateQosDeviceInterfaceInfo(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateQosDeviceInterfaceInfo", err, restyResp1.String(),
					"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateQosDeviceInterfaceInfo", err,
				"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}
	}

	return resourceQosDeviceInterfaceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vNetworkDeviceID := resourceMap["network_device_id"]

	queryParams1 := dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams
	queryParams1.NetworkDeviceID = vNetworkDeviceID
	item, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetQosDeviceInterfaceInfo", err,
			"Failure at GetQosDeviceInterfaceInfo, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.ApplicationPolicy.GetQosDeviceInterfaceInfo(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsApplicationPolicyGetQosDeviceInterfaceInfo(m, getResp1, nil)
		item1, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	response1, restyResp1, err := client.ApplicationPolicy.DeleteQosDeviceInterfaceInfo(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteQosDeviceInterfaceInfo", err, restyResp1.String(),
				"Failure at DeleteQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteQosDeviceInterfaceInfo", err,
			"Failure at DeleteQosDeviceInterfaceInfo, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo{}
	if v := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo{}
	if v := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetQosDeviceInterfaceInfo(m interface{}, queryParams dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams) (*dnacentersdkgo.ResponseItemApplicationPolicyGetQosDeviceInterfaceInfo, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemApplicationPolicyGetQosDeviceInterfaceInfo
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfo
	ite, _, err = client.ApplicationPolicy.GetQosDeviceInterfaceInfo(&queryParams)
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
			var getItem *dnacentersdkgo.ResponseItemApplicationPolicyGetQosDeviceInterfaceInfo
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
