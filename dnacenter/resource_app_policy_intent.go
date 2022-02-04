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

func resourceAppPolicyIntent() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Application Policy.

- Create/Update/Delete application policy
`,

		CreateContext: resourceAppPolicyIntentCreate,
		ReadContext:   resourceAppPolicyIntentRead,
		UpdateContext: resourceAppPolicyIntentUpdate,
		DeleteContext: resourceAppPolicyIntentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_policy_scope_element": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"group_id": &schema.Schema{
																Description: `Group id
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"ssid": &schema.Schema{
																Description: `Ssid
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Policy name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"consumer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to Queueing profile
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"delete_policy_status": &schema.Schema{
										Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"exclusive_contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"clause": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device_removal_behavior": &schema.Schema{
																Description: `Device eemoval behavior
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"host_tracking_enabled": &schema.Schema{
																Description: `Is host tracking enabled
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"relevance_level": &schema.Schema{
																Description: `Relevance level
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"type": &schema.Schema{
																Description: `Type
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"policy_scope": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"priority": &schema.Schema{
										Description: `Set to 4095 while producer refer to application Scalable group otherwise 100
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"producer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application-set or application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"delete_list": &schema.Schema{
							Description: `Delete list of Group Based Policy ids
`,
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"update_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_policy_scope_element": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"group_id": &schema.Schema{
																Description: `Group id
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `Id of Advance policy scope element
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"ssid": &schema.Schema{
																Description: `Ssid
`,
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `Id of Advance policy scope
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Policy name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"consumer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id of Consumer
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to Queueing profile
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"delete_policy_status": &schema.Schema{
										Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"exclusive_contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"clause": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device_removal_behavior": &schema.Schema{
																Description: `Device removal behavior
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"host_tracking_enabled": &schema.Schema{
																Description: `Host tracking enabled
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"id": &schema.Schema{
																Description: `Id of Business relevance or Application policy knobs clause
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"relevance_level": &schema.Schema{
																Description: `Relevance level
`,
																Type:     schema.TypeString,
																Optional: true,
															},
															"type": &schema.Schema{
																Description: `Type
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `Id of Exclusive contract
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id of Group based policy
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"policy_scope": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"priority": &schema.Schema{
										Description: `Set to 4095 while producer refer to application Scalable group otherwise 100
`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"producer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id of Producer
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application-set or application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
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

func resourceAppPolicyIntentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAppPolicyIntentApplicationPolicyIntent(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.ApplicationPolicy.ApplicationPolicyIntent(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ApplicationPolicyIntent", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntent", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceAppPolicyIntentRead(ctx, d, m)
}

func resourceAppPolicyIntentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vPolicyScope := resourceMap["policy_scope"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicy")
		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueryParams{}

		if okPolicyScope {
			queryParams1.PolicyScope = vPolicyScope
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicy(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationPolicy", err,
				"Failure at GetApplicationPolicy, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenApplicationPolicyGetApplicationPolicyItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicy search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAppPolicyIntentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceAppPolicyIntentRead(ctx, d, m)
}

func resourceAppPolicyIntentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete AppPolicyIntent on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestAppPolicyIntentApplicationPolicyIntent(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntent {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntent{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_list")))) {
		request.CreateList = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListArray(ctx, key+".create_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_list")))) {
		request.UpdateList = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListArray(ctx, key+".update_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_list")))) {
		request.DeleteList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentCreateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_policy_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_policy_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_policy_status")))) {
		request.DeletePolicyStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_scope")))) {
		request.PolicyScope = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope")))) {
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumer(ctx, key+".consumer.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".relevance_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".relevance_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".relevance_level")))) {
		request.RelevanceLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_removal_behavior")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_removal_behavior")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_removal_behavior")))) {
		request.DeviceRemovalBehavior = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_tracking_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_tracking_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_tracking_enabled")))) {
		request.HostTrackingEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListContract(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListContract {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducer {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentCreateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentUpdateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_policy_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_policy_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_policy_status")))) {
		request.DeletePolicyStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_scope")))) {
		request.PolicyScope = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope")))) {
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumer(ctx, key+".consumer.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".relevance_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".relevance_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".relevance_level")))) {
		request.RelevanceLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_removal_behavior")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_removal_behavior")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_removal_behavior")))) {
		request.DeviceRemovalBehavior = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_tracking_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_tracking_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_tracking_enabled")))) {
		request.HostTrackingEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListContract(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListContract {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup {
	request := []dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestAppPolicyIntentApplicationPolicyIntentUpdateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup {
	request := dnacentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchApplicationPolicyGetApplicationPolicy(m interface{}, queryParams dnacentersdkgo.GetApplicationPolicyQueryParams) (*dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicy, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicy
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicy
	ite, _, err = client.ApplicationPolicy.GetApplicationPolicy(&queryParams)
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
			var getItem *dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicy
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
