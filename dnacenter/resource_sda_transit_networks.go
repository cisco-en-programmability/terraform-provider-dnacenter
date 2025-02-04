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

func resourceSdaTransitNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Updates transit networks based on user input.

- Adds transit networks based on user input.
`,

		CreateContext: resourceSdaTransitNetworksCreate,
		ReadContext:   resourceSdaTransitNetworksRead,
		UpdateContext: resourceSdaTransitNetworksUpdate,
		DeleteContext: resourceSdaTransitNetworksDelete,
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

						"id": &schema.Schema{
							Description: `ID of the transit network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autonomous_system_number": &schema.Schema{
										Description: `Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295].
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"routing_protocol_name": &schema.Schema{
										Description: `Routing Protocol Name of the IP transit network.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Name of the transit network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"sda_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_plane_network_device_ids": &schema.Schema{
										Description: `List of network device IDs that are used as control plane nodes.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_multicast_over_transit_enabled": &schema.Schema{
										Description: `This indicates that multicast is enabled over SD-Access Transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"type": &schema.Schema{
							Description: `Type of the transit network.
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestItsmRetryIntegrationEvents`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `ID of the transit network (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_transit_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"autonomous_system_number": &schema.Schema{
													Description: `Autonomous system number of the IP transit network. Allowed range is [1 to 4294967295].
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"routing_protocol_name": &schema.Schema{
													Description: `Routing protocol name of the IP transit network.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Name of the transit network.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sda_transit_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"control_plane_network_device_ids": &schema.Schema{
													Description: `List of network device IDs that will be used as control plane nodes. Maximum 2 network device IDs can be provided for transit of type SDA_LISP_BGP_TRANSIT and maximum 4 network device IDs can be provided for transit of type SDA_LISP_PUB_SUB_TRANSIT.
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_multicast_over_transit_enabled": &schema.Schema{
													Description: `Set this to true to enable multicast over SD-Access transit.  This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type of the transit network.
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

func resourceSdaTransitNetworksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaTransitNetworksAddTransitNetworks(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParamImport := dnacentersdkgo.GetTransitNetworksQueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchSdaGetTransitNetworks(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaTransitNetworksRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddTransitNetworks(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddTransitNetworks", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddTransitNetworks", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddTransitNetworks", err))
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
				"Failure when executing AddTransitNetworks", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetTransitNetworksQueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchSdaGetTransitNetworks(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddTransitNetworks", err,
			"Failure at AddTransitNetworks, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = item3.Name
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaTransitNetworksRead(ctx, d, m)
}

func resourceSdaTransitNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransitNetworks")
		queryParams1 := dnacentersdkgo.GetTransitNetworksQueryParams{}
		queryParams1.Name = vName

		item1, err := searchSdaGetTransitNetworks(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used

		items := []dnacentersdkgo.ResponseSdaGetTransitNetworksResponse{
			*item1,
		}

		vItem1 := flattenSdaGetTransitNetworksItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransitNetworks search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaTransitNetworksUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaTransitNetworksUpdateTransitNetworks(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateTransitNetworks(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTransitNetworks", err, restyResp1.String(),
					"Failure at UpdateTransitNetworks, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTransitNetworks", err,
				"Failure at UpdateTransitNetworks, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateTransitNetworks", err))
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
					"Failure when executing UpdateTransitNetworks", err1))
				return diags
			}
		}

	}

	return resourceSdaTransitNetworksRead(ctx, d, m)
}

func resourceSdaTransitNetworksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SdaTransitNetworks", err, "Delete method is not supported",
		"Failure at SdaTransitNetworksDelete, unexpected response", ""))
	return diags
}

func expandRequestSdaTransitNetworksAddTransitNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddTransitNetworks {
	request := dnacentersdkgo.RequestSdaAddTransitNetworks{}
	if v := expandRequestSdaTransitNetworksAddTransitNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksAddTransitNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddTransitNetworks {
	request := []dnacentersdkgo.RequestItemSdaAddTransitNetworks{}
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
		i := expandRequestSdaTransitNetworksAddTransitNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksAddTransitNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddTransitNetworks {
	request := dnacentersdkgo.RequestItemSdaAddTransitNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_transit_settings")))) {
		request.IPTransitSettings = expandRequestSdaTransitNetworksAddTransitNetworksItemIPTransitSettings(ctx, key+".ip_transit_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sda_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sda_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sda_transit_settings")))) {
		request.SdaTransitSettings = expandRequestSdaTransitNetworksAddTransitNetworksItemSdaTransitSettings(ctx, key+".sda_transit_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksAddTransitNetworksItemIPTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddTransitNetworksIPTransitSettings {
	request := dnacentersdkgo.RequestItemSdaAddTransitNetworksIPTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".routing_protocol_name")))) {
		request.RoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".autonomous_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".autonomous_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".autonomous_system_number")))) {
		request.AutonomousSystemNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksAddTransitNetworksItemSdaTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddTransitNetworksSdaTransitSettings {
	request := dnacentersdkgo.RequestItemSdaAddTransitNetworksSdaTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_multicast_over_transit_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) {
		request.IsMulticastOverTransitEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".control_plane_network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".control_plane_network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".control_plane_network_device_ids")))) {
		request.ControlPlaneNetworkDeviceIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksUpdateTransitNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateTransitNetworks {
	request := dnacentersdkgo.RequestSdaUpdateTransitNetworks{}
	if v := expandRequestSdaTransitNetworksUpdateTransitNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksUpdateTransitNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateTransitNetworks {
	request := []dnacentersdkgo.RequestItemSdaUpdateTransitNetworks{}
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
		i := expandRequestSdaTransitNetworksUpdateTransitNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksUpdateTransitNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateTransitNetworks {
	request := dnacentersdkgo.RequestItemSdaUpdateTransitNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_transit_settings")))) {
		request.IPTransitSettings = expandRequestSdaTransitNetworksUpdateTransitNetworksItemIPTransitSettings(ctx, key+".ip_transit_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sda_transit_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sda_transit_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sda_transit_settings")))) {
		request.SdaTransitSettings = expandRequestSdaTransitNetworksUpdateTransitNetworksItemSdaTransitSettings(ctx, key+".sda_transit_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksUpdateTransitNetworksItemIPTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateTransitNetworksIPTransitSettings {
	request := dnacentersdkgo.RequestItemSdaUpdateTransitNetworksIPTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".routing_protocol_name")))) {
		request.RoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".autonomous_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".autonomous_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".autonomous_system_number")))) {
		request.AutonomousSystemNumber = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaTransitNetworksUpdateTransitNetworksItemSdaTransitSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateTransitNetworksSdaTransitSettings {
	request := dnacentersdkgo.RequestItemSdaUpdateTransitNetworksSdaTransitSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_multicast_over_transit_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_multicast_over_transit_enabled")))) {
		request.IsMulticastOverTransitEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".control_plane_network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".control_plane_network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".control_plane_network_device_ids")))) {
		request.ControlPlaneNetworkDeviceIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetTransitNetworks(m interface{}, queryParams dnacentersdkgo.GetTransitNetworksQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetTransitNetworksResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetTransitNetworksResponse
	var ite *dnacentersdkgo.ResponseSdaGetTransitNetworks
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetTransitNetworks(nil)
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
			nResponse, _, err = client.Sda.GetTransitNetworks(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.Sda.GetTransitNetworks(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
