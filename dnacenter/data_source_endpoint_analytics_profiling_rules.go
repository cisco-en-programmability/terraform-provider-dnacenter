package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEndpointAnalyticsProfilingRules() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AIEndpointAnalytics.

- This data source fetches the list of profiling rules. It can be used to show profiling rules in client applications,
or export those from an environment. 'POST /profiling-rules/bulk' API can be used to import such exported rules into
another environment. If this API is used to export rules to be imported into another Cisco DNA Center system, then
ensure that 'includeDeleted' parameter is 'true', so that deleted rules get synchronized correctly. Use query parameters
to filter the data, as required. If no filter is provided, then it will include only rules of type 'Custom Rule' in the
response. By default, the response is limited to 500 records. Use 'limit' parameter to fetch higher number of records,
if required. 'GET /profiling-rules/count' API can be used to find out the total number of rules in the system.

- Fetches details of the profiling rule for the given 'ruleId'.
`,

		ReadContext: dataSourceEndpointAnalyticsProfilingRulesRead,
		Schema: map[string]*schema.Schema{
			"include_deleted": &schema.Schema{
				Description: `includeDeleted query parameter. Flag to indicate whether deleted rules should be part of the records fetched.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to be fetched. If not provided, 500 records will be fetched by default. To fetch all the records in the system, provide a large value for this parameter.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Record offset to start data fetch at. Offset starts at zero.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order to be used for sorting.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": &schema.Schema{
				Description: `ruleId path parameter. Unique rule identifier
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_type": &schema.Schema{
				Description: `ruleType query parameter. Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested.
`,
				Type:     schema.TypeString,
				Optional: true,
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

			"items": &schema.Schema{
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
		},
	}
}

func dataSourceEndpointAnalyticsProfilingRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vRuleType, okRuleType := d.GetOk("rule_type")
	vIncludeDeleted, okIncludeDeleted := d.GetOk("include_deleted")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vRuleID, okRuleID := d.GetOk("rule_id")

	method1 := []bool{okRuleType, okIncludeDeleted, okLimit, okOffset, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okRuleID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetListOfProfilingRules")
		queryParams1 := dnacentersdkgo.GetListOfProfilingRulesQueryParams{}

		if okRuleType {
			queryParams1.RuleType = vRuleType.(string)
		}
		if okIncludeDeleted {
			queryParams1.IncludeDeleted = vIncludeDeleted.(bool)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.AIEndpointAnalytics.GetListOfProfilingRules(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetListOfProfilingRules", err,
				"Failure at GetListOfProfilingRules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenAIEndpointAnalyticsGetListOfProfilingRulesItems(response1.ProfilingRules)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfProfilingRules response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDetailsOfASingleProfilingRule")
		vvRuleID := vRuleID.(string)

		response2, restyResp2, err := client.AIEndpointAnalytics.GetDetailsOfASingleProfilingRule(vvRuleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDetailsOfASingleProfilingRule", err,
				"Failure at GetDetailsOfASingleProfilingRule, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDetailsOfASingleProfilingRule response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
