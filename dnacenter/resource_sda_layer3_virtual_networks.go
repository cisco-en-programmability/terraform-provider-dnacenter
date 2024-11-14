package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaLayer3VirtualNetworks() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds layer 3 virtual networks based on user input.

- Deletes layer 3 virtual networks based on user input.

- Updates layer 3 virtual networks based on user input.
`,

		CreateContext: resourceSdaLayer3VirtualNetworksCreate,
		ReadContext:   resourceSdaLayer3VirtualNetworksRead,
		UpdateContext: resourceSdaLayer3VirtualNetworksUpdate,
		DeleteContext: resourceSdaLayer3VirtualNetworksDelete,
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

						"anchored_site_id": &schema.Schema{
							Description: `Fabric ID of the fabric site this layer 3 virtual network is anchored at.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_ids": &schema.Schema{
							Description: `IDs of the fabrics this layer 3 virtual network is assigned to.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `ID of the layer 3 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Name of the layer 3 virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddLayer3VirtualNetworks`,
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
									"anchored_site_id": &schema.Schema{
										Description: `Fabric ID of the fabric site this layer 3 virtual network is to be anchored at.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_ids": &schema.Schema{
										Description: `IDs of the fabrics this layer 3 virtual network is to be assigned to.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `ID of the layer 3 virtual network (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"virtual_network_name": &schema.Schema{
										Description: `Name of the layer 3 virtual network.
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

func resourceSdaLayer3VirtualNetworksCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworks(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vName := resourceItem["virtual_network_name"]
	vvName := interfaceToString(vName)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParamImport := dnacentersdkgo.GetLayer3VirtualNetworksQueryParams{}
	queryParamImport.VirtualNetworkName = vvName
	item2, err := searchSdaGetLayer3VirtualNetworks(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["virtual_network_name"] = item2.VirtualNetworkName
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaLayer3VirtualNetworksRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddLayer3VirtualNetworks(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddLayer3VirtualNetworks", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddLayer3VirtualNetworks", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddLayer3VirtualNetworks", err))
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
				"Failure when executing AddLayer3VirtualNetworks", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetLayer3VirtualNetworksQueryParams{}
	queryParamValidate.VirtualNetworkName = vvName
	item3, err := searchSdaGetLayer3VirtualNetworks(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddLayer3VirtualNetworks", err,
			"Failure at AddLayer3VirtualNetworks, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["virtual_network_name"] = item3.VirtualNetworkName
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaLayer3VirtualNetworksRead(ctx, d, m)
}

func resourceSdaLayer3VirtualNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["virtual_network_name"]
	vvID := resourceMap["id"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer3VirtualNetworks")
		queryParams1 := dnacentersdkgo.GetLayer3VirtualNetworksQueryParams{}
		queryParams1.VirtualNetworkName = vName
		item1, err := searchSdaGetLayer3VirtualNetworks(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used

		items := []dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworksResponse{
			*item1,
		}

		vItem1 := flattenSdaGetLayer3VirtualNetworksItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer3VirtualNetworks search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaLayer3VirtualNetworksUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworks(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateLayer3VirtualNetworks(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateLayer3VirtualNetworks", err, restyResp1.String(),
					"Failure at UpdateLayer3VirtualNetworks, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateLayer3VirtualNetworks", err,
				"Failure at UpdateLayer3VirtualNetworks, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateLayer3VirtualNetworks", err))
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
					"Failure when executing UpdateLayer3VirtualNetworks", err1))
				return diags
			}
		}

	}

	return resourceSdaLayer3VirtualNetworksRead(ctx, d, m)
}

func resourceSdaLayer3VirtualNetworksDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := dnacentersdkgo.DeleteLayer3VirtualNetworksQueryParams{}

	vvVirtualNetworkName := resourceMap["virtual_network_name"]
	queryParamDelete.VirtualNetworkName = vvVirtualNetworkName

	response1, restyResp1, err := client.Sda.DeleteLayer3VirtualNetworks(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteLayer3VirtualNetworks", err, restyResp1.String(),
				"Failure at DeleteLayer3VirtualNetworks, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteLayer3VirtualNetworks", err,
			"Failure at DeleteLayer3VirtualNetworks, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteLayer3VirtualNetworks", err))
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
				"Failure when executing DeleteLayer3VirtualNetworks", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddLayer3VirtualNetworks {
	request := dnacentersdkgo.RequestSdaAddLayer3VirtualNetworks{}
	if v := expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddLayer3VirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaAddLayer3VirtualNetworks{}
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
		i := expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer3VirtualNetworksAddLayer3VirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddLayer3VirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaAddLayer3VirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_ids")))) {
		request.FabricIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anchored_site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anchored_site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anchored_site_id")))) {
		request.AnchoredSiteID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworks(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateLayer3VirtualNetworks {
	request := dnacentersdkgo.RequestSdaUpdateLayer3VirtualNetworks{}
	if v := expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworksItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworksItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateLayer3VirtualNetworks {
	request := []dnacentersdkgo.RequestItemSdaUpdateLayer3VirtualNetworks{}
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
		i := expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworksItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaLayer3VirtualNetworksUpdateLayer3VirtualNetworksItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateLayer3VirtualNetworks {
	request := dnacentersdkgo.RequestItemSdaUpdateLayer3VirtualNetworks{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_ids")))) {
		request.FabricIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anchored_site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anchored_site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anchored_site_id")))) {
		request.AnchoredSiteID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetLayer3VirtualNetworks(m interface{}, queryParams dnacentersdkgo.GetLayer3VirtualNetworksQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworksResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworksResponse
	var ite *dnacentersdkgo.ResponseSdaGetLayer3VirtualNetworks
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetLayer3VirtualNetworks(nil)
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
			nResponse, _, err = client.Sda.GetLayer3VirtualNetworks(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.VirtualNetworkName != "" {
		ite, _, err = client.Sda.GetLayer3VirtualNetworks(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.VirtualNetworkName == queryParams.VirtualNetworkName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
