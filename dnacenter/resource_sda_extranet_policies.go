package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaExtranetPolicies() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Deletes extranet policies based on user input.

- Updates an extranet policy based on user input.

- Adds an extranet policy based on user input.

- Deletes an extranet policy based on id.
`,

		CreateContext: resourceSdaExtranetPoliciesCreate,
		ReadContext:   resourceSdaExtranetPoliciesRead,
		UpdateContext: resourceSdaExtranetPoliciesUpdate,
		DeleteContext: resourceSdaExtranetPoliciesDelete,
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

						"extranet_policy_name": &schema.Schema{
							Description: `Name of the extranet policy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_ids": &schema.Schema{
							Description: `IDs of the fabric sites associated with this extranet policy.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `ID of the extranet policy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"provider_virtual_network_name": &schema.Schema{
							Description: `Name of the provider virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscriber_virtual_network_names": &schema.Schema{
							Description: `Name of the subscriber virtual network names.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddExtranetPolicy`,
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

									"extranet_policy_name": &schema.Schema{
										Description: `Name of the extranet policy to be created.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_ids": &schema.Schema{
										Description: `IDs of the fabric sites to be associated with this extranet policy.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `ID of the existing extranet policy (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"provider_virtual_network_name": &schema.Schema{
										Description: `Name of the existing provider virtual network.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"subscriber_virtual_network_names": &schema.Schema{
										Description: `Name of the subscriber virtual networks.
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
				},
			},
		},
	}
}

func resourceSdaExtranetPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaExtranetPoliciesAddExtranetPolicy(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vExtranetPolicyName := resourceItem["extranet_policy_name"]
	vvExtranetPolicyName := interfaceToString(vExtranetPolicyName)
	queryParamImport := dnacentersdkgo.GetExtranetPoliciesQueryParams{}
	queryParamImport.ExtranetPolicyName = vvExtranetPolicyName
	item2, err := searchSdaGetExtranetPolicies(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["name"] = item2.ExtranetPolicyName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaExtranetPoliciesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddExtranetPolicy(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddExtranetPolicy", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddExtranetPolicy", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddExtranetPolicy", err))
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
				"Failure when executing AddExtranetPolicy", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetExtranetPoliciesQueryParams{}
	queryParamValidate.ExtranetPolicyName = vvExtranetPolicyName
	item3, err := searchSdaGetExtranetPolicies(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddExtranetPolicy", err,
			"Failure at AddExtranetPolicy, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["name"] = item3.ExtranetPolicyName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaExtranetPoliciesRead(ctx, d, m)
}

func resourceSdaExtranetPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	vvName := resourceMap["name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExtranetPolicies")
		queryParams1 := dnacentersdkgo.GetExtranetPoliciesQueryParams{}
		queryParams1.ExtranetPolicyName = vvName
		item1, err := searchSdaGetExtranetPolicies(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseSdaGetExtranetPoliciesResponse{
			*item1,
		}

		vItem1 := flattenSdaGetExtranetPoliciesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExtranetPolicies search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaExtranetPoliciesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaExtranetPoliciesUpdateExtranetPolicy(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateExtranetPolicy(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateExtranetPolicy", err, restyResp1.String(),
					"Failure at UpdateExtranetPolicy, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateExtranetPolicy", err,
				"Failure at UpdateExtranetPolicy, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateExtranetPolicy", err))
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
					"Failure when executing UpdateExtranetPolicy", err1))
				return diags
			}
		}

	}

	return resourceSdaExtranetPoliciesRead(ctx, d, m)
}

func resourceSdaExtranetPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteExtranetPolicyByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteExtranetPolicyByID", err, restyResp1.String(),
				"Failure at DeleteExtranetPolicyByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteExtranetPolicyByID", err,
			"Failure at DeleteExtranetPolicyByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteExtranetPolicyByID", err))
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
				"Failure when executing DeleteExtranetPolicyByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaExtranetPoliciesAddExtranetPolicy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddExtranetPolicy {
	request := dnacentersdkgo.RequestSdaAddExtranetPolicy{}
	if v := expandRequestSdaExtranetPoliciesAddExtranetPolicyItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaExtranetPoliciesAddExtranetPolicyItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddExtranetPolicy {
	request := []dnacentersdkgo.RequestItemSdaAddExtranetPolicy{}
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
		i := expandRequestSdaExtranetPoliciesAddExtranetPolicyItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaExtranetPoliciesAddExtranetPolicyItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddExtranetPolicy {
	request := dnacentersdkgo.RequestItemSdaAddExtranetPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extranet_policy_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extranet_policy_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extranet_policy_name")))) {
		request.ExtranetPolicyName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_ids")))) {
		request.FabricIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider_virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider_virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider_virtual_network_name")))) {
		request.ProviderVirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscriber_virtual_network_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscriber_virtual_network_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscriber_virtual_network_names")))) {
		request.SubscriberVirtualNetworkNames = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaExtranetPoliciesUpdateExtranetPolicy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateExtranetPolicy {
	request := dnacentersdkgo.RequestSdaUpdateExtranetPolicy{}
	if v := expandRequestSdaExtranetPoliciesUpdateExtranetPolicyItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaExtranetPoliciesUpdateExtranetPolicyItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateExtranetPolicy {
	request := []dnacentersdkgo.RequestItemSdaUpdateExtranetPolicy{}
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
		i := expandRequestSdaExtranetPoliciesUpdateExtranetPolicyItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaExtranetPoliciesUpdateExtranetPolicyItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateExtranetPolicy {
	request := dnacentersdkgo.RequestItemSdaUpdateExtranetPolicy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extranet_policy_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extranet_policy_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extranet_policy_name")))) {
		request.ExtranetPolicyName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_ids")))) {
		request.FabricIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider_virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider_virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider_virtual_network_name")))) {
		request.ProviderVirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscriber_virtual_network_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscriber_virtual_network_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscriber_virtual_network_names")))) {
		request.SubscriberVirtualNetworkNames = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetExtranetPolicies(m interface{}, queryParams dnacentersdkgo.GetExtranetPoliciesQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetExtranetPoliciesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetExtranetPoliciesResponse

	ite, _, err := client.Sda.GetExtranetPolicies(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err
	}
	itemsCopy := *ite.Response
	if itemsCopy == nil {
		return foundItem, err
	}
	for _, item := range itemsCopy {
		if item.ExtranetPolicyName == queryParams.ExtranetPolicyName {
			foundItem = &item
			return foundItem, err
		}
	}
	return foundItem, err

}
