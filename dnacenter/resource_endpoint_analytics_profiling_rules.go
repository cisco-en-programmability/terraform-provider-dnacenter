package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpointAnalyticsProfilingRules() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on AIEndpointAnalytics.

- Creates profiling rule from the request body.

- Updates the profiling rule for the given 'ruleId'.

- Deletes the profiling rule for the given 'ruleId'.
`,

		CreateContext: resourceEndpointAnalyticsProfilingRulesCreate,
		ReadContext:   resourceEndpointAnalyticsProfilingRulesRead,
		UpdateContext: resourceEndpointAnalyticsProfilingRulesUpdate,
		DeleteContext: resourceEndpointAnalyticsProfilingRulesDelete,
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

						"cluster_id": &schema.Schema{
							Description: `Unique identifier for ML cluster. Only applicable for 'ML Rule'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"condition_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"attribute_dictionary": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"operator": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"condition_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_deleted": &schema.Schema{
							Description: `Flag to indicate whether the rule was deleted.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_by": &schema.Schema{
							Description: `User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified_on": &schema.Schema{
							Description: `Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"plugin_id": &schema.Schema{
							Description: `Plugin for the rule. Only applicable for 'Cisco Default' rules.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rejected": &schema.Schema{
							Description: `Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"result": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_type": &schema.Schema{
										Description: `List of device types determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_manufacturer": &schema.Schema{
										Description: `List of hardware manufacturers determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_model": &schema.Schema{
										Description: `List of hardware models determined by the current rule.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operating_system": &schema.Schema{
										Description: `List of operating systems determined by the current rule.
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
						"rule_id": &schema.Schema{
							Description: `Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_name": &schema.Schema{
							Description: `Human readable name for the rule.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_priority": &schema.Schema{
							Description: `Priority for the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rule_type": &schema.Schema{
							Description: `Type of the rule.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_version": &schema.Schema{
							Description: `Version of the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"source_priority": &schema.Schema{
							Description: `Source priority for the rule.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_attributes": &schema.Schema{
							Description: `List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
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
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cluster_id": &schema.Schema{
							Description: `Unique identifier for ML cluster. Only applicable for 'ML Rule'.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"condition_groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"condition": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"attribute_dictionary": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"operator": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"condition_group": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operator": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"is_deleted": &schema.Schema{
							Description: `Flag to indicate whether the rule was deleted.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"last_modified_by": &schema.Schema{
							Description: `User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified_on": &schema.Schema{
							Description: `Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"plugin_id": &schema.Schema{
							Description: `Plugin for the rule. Only applicable for 'Cisco Default' rules.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rejected": &schema.Schema{
							Description: `Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"result": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_type": &schema.Schema{
										Description: `List of device types determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_manufacturer": &schema.Schema{
										Description: `List of hardware manufacturers determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"hardware_model": &schema.Schema{
										Description: `List of hardware models determined by the current rule.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"operating_system": &schema.Schema{
										Description: `List of operating systems determined by the current rule.
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
						"rule_id": &schema.Schema{
							Description: `Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule_name": &schema.Schema{
							Description: `Human readable name for the rule.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule_priority": &schema.Schema{
							Description: `Priority for the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rule_type": &schema.Schema{
							Description: `Type of the rule.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule_version": &schema.Schema{
							Description: `Version of the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_priority": &schema.Schema{
							Description: `Source priority for the rule.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"used_attributes": &schema.Schema{
							Description: `List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
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
	}
}

func resourceEndpointAnalyticsProfilingRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRule(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vRuleID, okRuleID := resourceItem["rule_id"]
	vvRuleID := interfaceToString(vRuleID)
	vName, okRuleID := resourceItem["rule_name"]
	vvName := interfaceToString(vName)
	if okRuleID && vvRuleID != "" {
		getResponse2, _, err := client.AIEndpointAnalytics.GetDetailsOfASingleProfilingRule(vvRuleID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["rule_id"] = vvRuleID
			d.SetId(joinResourceID(resourceMap))
			return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.GetListOfProfilingRulesQueryParams{}

		response2, err := searchPolicyGetListOfProfilingRules(m, queryParamImport, vvRuleID, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["rule_id"] = response2.RuleID
			d.SetId(joinResourceID(resourceMap))
			return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.AIEndpointAnalytics.CreateAProfilingRule(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAProfilingRule", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAProfilingRule", err))
		return diags
	}
	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.GetListOfProfilingRulesQueryParams{}
	item3, err := searchPolicyGetListOfProfilingRules(m, queryParamValidate, vvRuleID, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateAProfilingRule", err,
			"Failure at CreateAProfilingRule, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["rule_id"] = item3.RuleID
	d.SetId(joinResourceID(resourceMap))
	return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
}

func resourceEndpointAnalyticsProfilingRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDetailsOfASingleProfilingRule")
		vvRuleID := vID

		// has_unknown_response: None

		response1, restyResp1, err := client.AIEndpointAnalytics.GetDetailsOfASingleProfilingRule(vvRuleID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		// Review flatten function used
		vItem1 := flattenAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfProfilingRules search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEndpointAnalyticsProfilingRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRule(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		restyResp1, err := client.AIEndpointAnalytics.UpdateAnExistingProfilingRule(vvID, request1)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAnExistingProfilingRule", err, restyResp1.String(),
					"Failure at UpdateAnExistingProfilingRule, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAnExistingProfilingRule", err,
				"Failure at UpdateAnExistingProfilingRule, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceEndpointAnalyticsProfilingRulesRead(ctx, d, m)
}

func resourceEndpointAnalyticsProfilingRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvRuleID := resourceMap["id"]

	restyResp1, err := client.AIEndpointAnalytics.DeleteAnExistingProfilingRule(vvRuleID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteAnExistingProfilingRule", err, restyResp1.String(),
				"Failure at DeleteAnExistingProfilingRule, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteAnExistingProfilingRule", err,
			"Failure at DeleteAnExistingProfilingRule, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRule {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_id")))) {
		request.RuleID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_name")))) {
		request.RuleName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_type")))) {
		request.RuleType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_version")))) {
		request.RuleVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_priority")))) {
		request.RulePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_priority")))) {
		request.SourcePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_deleted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_deleted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_deleted")))) {
		request.IsDeleted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_by")))) {
		request.LastModifiedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_on")))) {
		request.LastModifiedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plugin_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plugin_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plugin_id")))) {
		request.PluginID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cluster_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cluster_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cluster_id")))) {
		request.ClusterID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rejected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rejected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rejected")))) {
		request.Rejected = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".result")))) {
		request.Result = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleResult(ctx, key+".result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_groups")))) {
		request.ConditionGroups = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroups(ctx, key+".condition_groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".used_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".used_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".used_attributes")))) {
		request.UsedAttributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleResult(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleResult {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operating_system")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operating_system")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operating_system")))) {
		request.OperatingSystem = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroups {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_group")))) {
		request.ConditionGroup = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesCreateAProfilingRuleConditionGroupsCondition(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroupsCondition {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroupsCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute")))) {
		request.Attribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_dictionary")))) {
		request.AttributeDictionary = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRule(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRule {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRule{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_id")))) {
		request.RuleID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_name")))) {
		request.RuleName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_type")))) {
		request.RuleType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_version")))) {
		request.RuleVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rule_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rule_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rule_priority")))) {
		request.RulePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_priority")))) {
		request.SourcePriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_deleted")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_deleted")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_deleted")))) {
		request.IsDeleted = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_by")))) {
		request.LastModifiedBy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_on")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_on")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_on")))) {
		request.LastModifiedOn = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".plugin_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".plugin_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".plugin_id")))) {
		request.PluginID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cluster_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cluster_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cluster_id")))) {
		request.ClusterID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rejected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rejected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rejected")))) {
		request.Rejected = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".result")))) {
		request.Result = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleResult(ctx, key+".result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_groups")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_groups")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_groups")))) {
		request.ConditionGroups = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroups(ctx, key+".condition_groups.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".used_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".used_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".used_attributes")))) {
		request.UsedAttributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleResult(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleResult {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_manufacturer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_manufacturer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_manufacturer")))) {
		request.HardwareManufacturer = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_model")))) {
		request.HardwareModel = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operating_system")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operating_system")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operating_system")))) {
		request.OperatingSystem = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroups(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroups {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroups{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition")))) {
		request.Condition = expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsCondition(ctx, key+".condition.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".condition_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".condition_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".condition_group")))) {
		request.ConditionGroup = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEndpointAnalyticsProfilingRulesUpdateAnExistingProfilingRuleConditionGroupsCondition(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroupsCondition {
	request := dnacentersdkgo.RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroupsCondition{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute")))) {
		request.Attribute = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attribute_dictionary")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attribute_dictionary")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attribute_dictionary")))) {
		request.AttributeDictionary = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchPolicyGetListOfProfilingRules(m interface{}, queryParams dnacentersdkgo.GetListOfProfilingRulesQueryParams, vID string, vName string) (*dnacentersdkgo.ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRules, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRules
	if vID != "" || vName != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.AIEndpointAnalytics.GetListOfProfilingRules(&queryParams)
		maxPageSize := len(*nResponse.ProfilingRules)
		for len(*nResponse.ProfilingRules) > 0 {
			for _, item := range *nResponse.ProfilingRules {
				if vID == item.RuleID || vName == item.RuleName {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.AIEndpointAnalytics.GetListOfProfilingRules(&queryParams)
			if nResponse == nil || nResponse.ProfilingRules == nil {
				break
			}
		}
		return nil, err
	}
	return foundItem, err
}
