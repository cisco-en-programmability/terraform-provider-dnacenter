package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppPolicy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get all existing application policies
`,

		ReadContext: dataSourceAppPolicyRead,
		Schema: map[string]*schema.Schema{
			"policy_scope": &schema.Schema{
				Description: `policyScope query parameter. policy scope name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advanced_policy_scope": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope_element": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"group_id": &schema.Schema{
													Description: `Group id
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"id": &schema.Schema{
													Description: `Id of Advanced policy scope element
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"ssid": &schema.Schema{
													Description: `Ssid
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

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Advanced policy scope
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"consumer": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Consumer
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"scalable_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to application Scalable group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"contract_list": &schema.Schema{
							Description: `Contract list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"create_time": &schema.Schema{
							Description: `Create time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"custom_provisions": &schema.Schema{
							Description: `Custom provisions
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"delete_policy_status": &schema.Schema{
							Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exclusive_contract": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"clause": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_removal_behavior": &schema.Schema{
													Description: `Device removal behavior
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"host_tracking_enabled": &schema.Schema{
													Description: `Host tracking enabled
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id of Business relevance or Application policy knobs clause
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"priority": &schema.Schema{
													Description: `Priority
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"relevance_level": &schema.Schema{
													Description: `Relevance level
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `Type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Exclusive contract
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id of Group based policy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Identity source
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `State
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"instance_created_on": &schema.Schema{
							Description: `Instance created on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Description: `Instance updated on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"internal": &schema.Schema{
							Description: `Internal
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_deleted": &schema.Schema{
							Description: `Is deleted
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Is enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_scope_stale": &schema.Schema{
							Description: `Is scope stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_seeded": &schema.Schema{
							Description: `Is seeded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_stale": &schema.Schema{
							Description: `Is stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ise_reserved": &schema.Schema{
							Description: `Is reserved
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_scope": &schema.Schema{
							Description: `Policy name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_status": &schema.Schema{
							Description: `Policy status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"priority": &schema.Schema{
							Description: `Priority
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"producer": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Producer
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"scalable_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to application-set or application Scalable group
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"provisioning_state": &schema.Schema{
							Description: `Provisioning state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pushed": &schema.Schema{
							Description: `Pushed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"qualifier": &schema.Schema{
							Description: `Qualifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"resource_version": &schema.Schema{
							Description: `Resource version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"target_id_list": &schema.Schema{
							Description: `Target id list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAppPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPolicyScope, okPolicyScope := d.GetOk("policy_scope")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicy")
		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueryParams{}

		if okPolicyScope {
			queryParams1.PolicyScope = vPolicyScope.(string)
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

		vItems1 := flattenApplicationPolicyGetApplicationPolicyItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicy response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationPolicyItems(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["create_time"] = item.CreateTime
		respItem["deployed"] = boolPtrToString(item.Deployed)
		respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
		respItem["is_stale"] = boolPtrToString(item.IsStale)
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["provisioning_state"] = item.ProvisioningState
		respItem["qualifier"] = item.Qualifier
		respItem["resource_version"] = item.ResourceVersion
		respItem["target_id_list"] = flattenApplicationPolicyGetApplicationPolicyItemsTargetIDList(item.TargetIDList)
		respItem["type"] = item.Type
		respItem["cfs_change_info"] = flattenApplicationPolicyGetApplicationPolicyItemsCfsChangeInfo(item.CfsChangeInfo)
		respItem["custom_provisions"] = flattenApplicationPolicyGetApplicationPolicyItemsCustomProvisions(item.CustomProvisions)
		respItem["delete_policy_status"] = item.DeletePolicyStatus
		respItem["internal"] = boolPtrToString(item.Internal)
		respItem["is_deleted"] = boolPtrToString(item.IsDeleted)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["is_scope_stale"] = boolPtrToString(item.IsScopeStale)
		respItem["ise_reserved"] = boolPtrToString(item.IseReserved)
		respItem["policy_scope"] = item.PolicyScope
		respItem["policy_status"] = item.PolicyStatus
		respItem["priority"] = item.Priority
		respItem["pushed"] = boolPtrToString(item.Pushed)
		respItem["advanced_policy_scope"] = flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScope(item.AdvancedPolicyScope)
		respItem["contract_list"] = flattenApplicationPolicyGetApplicationPolicyItemsContractList(item.ContractList)
		respItem["exclusive_contract"] = flattenApplicationPolicyGetApplicationPolicyItemsExclusiveContract(item.ExclusiveContract)
		respItem["identity_source"] = flattenApplicationPolicyGetApplicationPolicyItemsIDentitySource(item.IDentitySource)
		respItem["producer"] = flattenApplicationPolicyGetApplicationPolicyItemsProducer(item.Producer)
		respItem["consumer"] = flattenApplicationPolicyGetApplicationPolicyItemsConsumer(item.Consumer)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsTargetIDList(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseTargetIDList) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsCfsChangeInfo(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseCfsChangeInfo) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsCustomProvisions(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseCustomProvisions) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScope(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScope) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["name"] = item.Name
	respItem["advanced_policy_scope_element"] = flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScopeAdvancedPolicyScopeElement(item.AdvancedPolicyScopeElement)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScopeAdvancedPolicyScopeElement(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElement) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["group_id"] = item.GroupID
		respItem["ssid"] = flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScopeAdvancedPolicyScopeElementSSID(item.SSID)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsAdvancedPolicyScopeAdvancedPolicyScopeElementSSID(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsContractList(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseContractList) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsExclusiveContract(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContract) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["clause"] = flattenApplicationPolicyGetApplicationPolicyItemsExclusiveContractClause(item.Clause)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyItemsExclusiveContractClause(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContractClause) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["priority"] = item.Priority
		respItem["type"] = item.Type
		respItem["relevance_level"] = item.RelevanceLevel
		respItem["device_removal_behavior"] = item.DeviceRemovalBehavior
		respItem["host_tracking_enabled"] = boolPtrToString(item.HostTrackingEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsIDentitySource(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseIDentitySource) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["state"] = item.State
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyItemsProducer(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseProducer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["scalable_group"] = flattenApplicationPolicyGetApplicationPolicyItemsProducerScalableGroup(item.ScalableGroup)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyItemsProducerScalableGroup(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseProducerScalableGroup) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id_ref"] = item.IDRef
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyItemsConsumer(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseConsumer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["scalable_group"] = flattenApplicationPolicyGetApplicationPolicyItemsConsumerScalableGroup(item.ScalableGroup)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyItemsConsumerScalableGroup(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyResponseConsumerScalableGroup) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id_ref"] = item.IDRef
		respItems = append(respItems, respItem)
	}
	return respItems
}
