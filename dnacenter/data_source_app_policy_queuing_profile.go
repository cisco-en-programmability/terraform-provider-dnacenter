package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppPolicyQueuingProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get all or by name, existing application policy queuing profiles
`,

		ReadContext: dataSourceAppPolicyQueuingProfileRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name query parameter. queuing profile name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"clause": &schema.Schema{
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
										Description: `Id
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

									"interface_speed_bandwidth_clauses": &schema.Schema{
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
													Description: `Id
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

												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"display_name": &schema.Schema{
																Description: `Display name
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"id": &schema.Schema{
																Description: `Id
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

															"traffic_class": &schema.Schema{
																Description: `Traffic Class
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

									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tc_dscp_settings": &schema.Schema{
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

												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id
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

												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
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

						"contract_classifier": &schema.Schema{
							Description: `Contract classifier
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

						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Free test description
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

						"gen_id": &schema.Schema{
							Description: `Gen id
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of Queueing profile
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
							Description: `Queueing profile name
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

						"provisioning_state": &schema.Schema{
							Description: `Provisioning State
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

func dataSourceAppPolicyQueuingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetApplicationPolicyQueuingProfile")
		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationPolicyQueuingProfile", err,
				"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationPolicyQueuingProfileItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyQueuingProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItems(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse) []map[string]interface{} {
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
		respItem["description"] = item.Description
		respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
		respItem["is_stale"] = boolPtrToString(item.IsStale)
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["provisioning_state"] = item.ProvisioningState
		respItem["qualifier"] = item.Qualifier
		respItem["resource_version"] = item.ResourceVersion
		respItem["target_id_list"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsTargetIDList(item.TargetIDList)
		respItem["type"] = item.Type
		respItem["cfs_change_info"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCfsChangeInfo(item.CfsChangeInfo)
		respItem["custom_provisions"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCustomProvisions(item.CustomProvisions)
		respItem["gen_id"] = item.GenID
		respItem["internal"] = boolPtrToString(item.Internal)
		respItem["is_deleted"] = boolPtrToString(item.IsDeleted)
		respItem["ise_reserved"] = boolPtrToString(item.IseReserved)
		respItem["pushed"] = boolPtrToString(item.Pushed)
		respItem["clause"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClause(item.Clause)
		respItem["contract_classifier"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsContractClassifier(item.ContractClassifier)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItem(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse) []map[string]interface{} {
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
	respItem["create_time"] = item.CreateTime
	respItem["deployed"] = boolPtrToString(item.Deployed)
	respItem["description"] = item.Description
	respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
	respItem["is_stale"] = boolPtrToString(item.IsStale)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["namespace"] = item.Namespace
	respItem["provisioning_state"] = item.ProvisioningState
	respItem["qualifier"] = item.Qualifier
	respItem["resource_version"] = item.ResourceVersion
	respItem["target_id_list"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsTargetIDList(item.TargetIDList)
	respItem["type"] = item.Type
	respItem["cfs_change_info"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCfsChangeInfo(item.CfsChangeInfo)
	respItem["custom_provisions"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCustomProvisions(item.CustomProvisions)
	respItem["gen_id"] = item.GenID
	respItem["internal"] = boolPtrToString(item.Internal)
	respItem["is_deleted"] = boolPtrToString(item.IsDeleted)
	respItem["ise_reserved"] = boolPtrToString(item.IseReserved)
	respItem["pushed"] = boolPtrToString(item.Pushed)
	respItem["clause"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClause(item.Clause)
	respItem["contract_classifier"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsContractClassifier(item.ContractClassifier)

	return []map[string]interface{}{
		respItem,
	}
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsTargetIDList(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseTargetIDList) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCfsChangeInfo(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCfsChangeInfo) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsCustomProvisions(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCustomProvisions) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClause(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClause) []map[string]interface{} {
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
		respItem["is_common_between_all_interface_speeds"] = boolPtrToString(item.IsCommonBetweenAllInterfaceSpeeds)
		respItem["interface_speed_bandwidth_clauses"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(item.InterfaceSpeedBandwidthClauses)
		respItem["tc_dscp_settings"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(item.TcDscpSettings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClauses(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClauses) []map[string]interface{} {
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
		respItem["interface_speed"] = item.InterfaceSpeed
		respItem["tc_bandwidth_settings"] = flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(item.TcBandwidthSettings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings) []map[string]interface{} {
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
		respItem["bandwidth_percentage"] = item.BandwidthPercentage
		respItem["traffic_class"] = item.TrafficClass
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsClauseTcDscpSettings(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseTcDscpSettings) []map[string]interface{} {
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
		respItem["dscp"] = item.Dscp
		respItem["traffic_class"] = item.TrafficClass
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyQueuingProfileItemsContractClassifier(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseContractClassifier) []interface{} {
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
