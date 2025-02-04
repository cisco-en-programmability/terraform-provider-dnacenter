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

func resourceSdaLayer2VirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds layer 2 virtual networks based on user input.

- Deletes layer 2 virtual networks based on user input.

- Updates layer 2 virtual networks based on user input.

- Deletes a layer 2 virtual network based on id.
`,

		CreateContext: resourceSdaLayer2VirtualNetworksCreate,
		ReadContext:   resourceSdaLayer2VirtualNetworksRead,
		UpdateContext: resourceSdaLayer2VirtualNetworksUpdate,
		DeleteContext: resourceSdaLayer2VirtualNetworksDelete,
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

						"associated_layer3_virtual_network_name": &schema.Schema{
							Description: `Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this layer 2 virtual network is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the layer 2 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_fabric_enabled_wireless": &schema.Schema{
							Description: `Set to true to enable wireless.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"traffic_type": &schema.Schema{
							Description: `The type of traffic that is served.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Description: `ID of the VLAN of the layer 2 virtual network.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vlan_name": &schema.Schema{
							Description: `Name of the VLAN of the layer 2 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddLayer2VirtualNetworks`,
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
									"associated_layer3_virtual_network_name": &schema.Schema{
										Description: `Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. The layer 3 virtual network must have already been added to the fabric before association. This field must either be present in all payload elements or none.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this layer 2 virtual network is to be assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the layer 2 virtual network (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_fabric_enabled_wireless": &schema.Schema{
										Description: `Set to true to enable wireless. Default is false.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"traffic_type": &schema.Schema{
										Description: `The type of traffic that is served.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vlan_id": &schema.Schema{
										Description: `ID of the VLAN of the layer 2 virtual network. Allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, and 2046. If deploying on a fabric zone, this vlanId must match the vlanId of the corresponding layer 2 virtual network on the fabric site.
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"vlan_name": &schema.Schema{
										Description: `Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens.
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

func resourceSdaLayer2VirtualNetworksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworks(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vFabricID := resourceItem["fabric_id"]
	vvFabricID := interfaceToString(vFabricID)
	vName := resourceItem["associated_layer3_virtual_network_name"]
	vvName := interfaceToString(vName)

	queryParamImport := dnacentersdkgo.GetLayer2VirtualNetworksQueryParams{}
	queryParamImport.FabricID = vvFabricID
	queryParamImport.AssociatedLayer3VirtualNetworkName = vvName
	item2, err := searchSdaGetLayer2VirtualNetworks(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["fabric_id"] = item2.FabricID
		resourceMap["associated_layer3_virtual_network_name"] = item2.AssociatedLayer3VirtualNetworkName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaLayer2VirtualNetworksRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddLayer2VirtualNetworks(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddLayer2VirtualNetworks", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddLayer2VirtualNetworks", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddLayer2VirtualNetworks", err))
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
				"Failure when executing AddLayer2VirtualNetworks", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetLayer2VirtualNetworksQueryParams{}
	queryParamValidate.ID = vvID
	queryParamValidate.FabricID = vvFabricID
	queryParamValidate.AssociatedLayer3VirtualNetworkName = vvName
	item3, err := searchSdaGetLayer2VirtualNetworks(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddLayer2VirtualNetworks", err,
			"Failure at AddLayer2VirtualNetworks, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["fabric_id"] = item3.FabricID
	resourceMap["associated_layer3_virtual_network_name"] = item3.AssociatedLayer3VirtualNetworkName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaLayer2VirtualNetworksRead(ctx, d, m)
}

func resourceSdaLayer2VirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	vvFabicID := resourceMap["fabric_id"]
	vvName := resourceMap["associated_layer3_virtual_network_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer2VirtualNetworks")
		queryParams1 := dnacentersdkgo.GetLayer2VirtualNetworksQueryParams{}
		queryParams1.FabricID = vvFabicID
		queryParams1.AssociatedLayer3VirtualNetworkName = vvName

		item1, err := searchSdaGetLayer2VirtualNetworks(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseSdaGetLayer2VirtualNetworksResponse{
			*item1,
		}
		vItem1 := flattenSdaGetLayer2VirtualNetworksItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer2VirtualNetworks search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaLayer2VirtualNetworksUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	if d.HasChange("parameters") {
		request1 := expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworks(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateLayer2VirtualNetworks(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateLayer2VirtualNetworks", err, restyResp1.String(),
					"Failure at UpdateLayer2VirtualNetworks, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateLayer2VirtualNetworks", err,
				"Failure at UpdateLayer2VirtualNetworks, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateLayer2VirtualNetworks", err))
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
					"Failure when executing UpdateLayer2VirtualNetworks", err1))
				return diags
			}
		}

	}

	return resourceSdaLayer2VirtualNetworksRead(ctx, d, m)
}

func resourceSdaLayer2VirtualNetworksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Sda.DeleteLayer2VirtualNetworkByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteLayer2VirtualNetworkByID", err, restyResp1.String(),
				"Failure at DeleteLayer2VirtualNetworkByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteLayer2VirtualNetworkByID", err,
			"Failure at DeleteLayer2VirtualNetworkByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteLayer2VirtualNetworkByID", err))
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
				"Failure when executing DeleteLayer2VirtualNetworkByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddLayer2VirtualNetworks {
	request := dnacentersdkgo.RequestSdaAddLayer2VirtualNetworks{}
	if v := expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddLayer2VirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaAddLayer2VirtualNetworks{}
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
		i := expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer2VirtualNetworksAddLayer2VirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddLayer2VirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaAddLayer2VirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_fabric_enabled_wireless")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_fabric_enabled_wireless")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_fabric_enabled_wireless")))) {
		request.IsFabricEnabledWireless = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".associated_layer3_virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".associated_layer3_virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".associated_layer3_virtual_network_name")))) {
		request.AssociatedLayer3VirtualNetworkName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateLayer2VirtualNetworks {
	request := dnacentersdkgo.RequestSdaUpdateLayer2VirtualNetworks{}
	if v := expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateLayer2VirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaUpdateLayer2VirtualNetworks{}
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
		i := expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer2VirtualNetworksUpdateLayer2VirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateLayer2VirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaUpdateLayer2VirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_fabric_enabled_wireless")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_fabric_enabled_wireless")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_fabric_enabled_wireless")))) {
		request.IsFabricEnabledWireless = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".associated_layer3_virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".associated_layer3_virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".associated_layer3_virtual_network_name")))) {
		request.AssociatedLayer3VirtualNetworkName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetLayer2VirtualNetworks(m interface{}, queryParams dnacentersdkgo.GetLayer2VirtualNetworksQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetLayer2VirtualNetworksResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetLayer2VirtualNetworksResponse
	var ite *dnacentersdkgo.ResponseSdaGetLayer2VirtualNetworks
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetLayer2VirtualNetworks(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetLayer2VirtualNetworks(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.AssociatedLayer3VirtualNetworkName != "" {
		ite, _, err = client.Sda.GetLayer2VirtualNetworks(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.AssociatedLayer3VirtualNetworkName == queryParams.AssociatedLayer3VirtualNetworkName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
