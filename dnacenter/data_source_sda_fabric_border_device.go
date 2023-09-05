package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get border device detail from SDA Fabric
`,

		ReadContext: dataSourceSdaFabricBorderDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"payload": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"akc_settings_cfs": &schema.Schema{
										Description: `Akc Settings Cfs`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"auth_entity_class": &schema.Schema{
										Description: `Auth Entity Class`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"auth_entity_id": &schema.Schema{
										Description: `Auth Entity Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"cfs_change_info": &schema.Schema{
										Description: `Cfs Change Info`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"configs": &schema.Schema{
										Description: `Configs`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"create_time": &schema.Schema{
										Description: `Create Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"custom_provisions": &schema.Schema{
										Description: `Custom Provisions`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"deploy_pending": &schema.Schema{
										Description: `Deploy Pending`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"deployed": &schema.Schema{
										Description: `Deployed`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_interface_info": &schema.Schema{
										Description: `Device Interface Info`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"device_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connected_to": &schema.Schema{
													Description: `Connected To`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"cpu": &schema.Schema{
													Description: `Cpu`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"deploy_pending": &schema.Schema{
													Description: `Deploy Pending`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"dhcp_enabled": &schema.Schema{
													Description: `Dhcp Enabled`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_name": &schema.Schema{
													Description: `Display Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"ext_connectivity_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"deploy_pending": &schema.Schema{
																Description: `Deploy Pending`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"display_name": &schema.Schema{
																Description: `Display Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"external_domain_protocol_number": &schema.Schema{
																Description: `External Domain Protocol Number`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"id": &schema.Schema{
																Description: `Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"instance_id": &schema.Schema{
																Description: `Instance Id`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"instance_tenant_id": &schema.Schema{
																Description: `Instance Tenant Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"instance_version": &schema.Schema{
																Description: `Instance Version`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"interface_uuid": &schema.Schema{
																Description: `Interface Uuid`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"l2_handoff": &schema.Schema{
																Description: `L2 Handoff`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"l3_handoff": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"deploy_pending": &schema.Schema{
																			Description: `Deploy Pending`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"display_name": &schema.Schema{
																			Description: `Display Name`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"instance_id": &schema.Schema{
																			Description: `Instance Id`,
																			Type:        schema.TypeInt,
																			Computed:    true,
																		},

																		"instance_tenant_id": &schema.Schema{
																			Description: `Instance Tenant Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"instance_version": &schema.Schema{
																			Description: `Instance Version`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},

																		"local_ip_address": &schema.Schema{
																			Description: `Local Ip Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"remote_ip_address": &schema.Schema{
																			Description: `Remote Ip Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"virtual_network": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"id_ref": &schema.Schema{
																						Description: `Id Ref`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},

																		"vlan_id": &schema.Schema{
																			Description: `Vlan Id`,
																			Type:        schema.TypeInt,
																			Computed:    true,
																		},
																	},
																},
															},

															"policy_propagation_enabled": &schema.Schema{
																Description: `Policy Propagation Enabled`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"policy_sgt_tag": &schema.Schema{
																Description: `Policy Sgt Tag`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
														},
													},
												},

												"external_connectivity_ip_pool": &schema.Schema{
													Description: `External Connectivity Ip Pool`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"external_domain_routing_protocol": &schema.Schema{
													Description: `External Domain Routing Protocol`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"instance_tenant_id": &schema.Schema{
													Description: `Instance Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance Version`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"internal_domain_protocol_number": &schema.Schema{
													Description: `Internal Domain Protocol Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"memory": &schema.Schema{
													Description: `Memory`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"node_type": &schema.Schema{
													Description: `Node Type`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"storage": &schema.Schema{
													Description: `Storage`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance Version`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"is_seeded": &schema.Schema{
										Description: `Is Seeded`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_stale": &schema.Schema{
										Description: `Is Stale`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Last Update Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"managed_sites": &schema.Schema{
										Description: `Managed Sites`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"namespace": &schema.Schema{
										Description: `Namespace`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"network_device_id": &schema.Schema{
										Description: `Network Device Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"network_wide_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa": &schema.Schema{
													Description: `Aaa`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"cmx": &schema.Schema{
													Description: `Cmx`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"deploy_pending": &schema.Schema{
													Description: `Deploy Pending`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"dhcp": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ip_address": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"address": &schema.Schema{
																			Description: `Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"address_type": &schema.Schema{
																			Description: `Address Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"padded_address": &schema.Schema{
																			Description: `Padded Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},

												"display_name": &schema.Schema{
													Description: `Display Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"dns": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"domain_name": &schema.Schema{
																Description: `Domain Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"id": &schema.Schema{
																Description: `Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ip": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"address": &schema.Schema{
																			Description: `Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"address_type": &schema.Schema{
																			Description: `Address Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"padded_address": &schema.Schema{
																			Description: `Padded Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"instance_tenant_id": &schema.Schema{
													Description: `Instance Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_version": &schema.Schema{
													Description: `Instance Version`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"ldap": &schema.Schema{
													Description: `Ldap`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"native_vlan": &schema.Schema{
													Description: `Native Vlan`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"netflow": &schema.Schema{
													Description: `Netflow`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"ntp": &schema.Schema{
													Description: `Ntp`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"snmp": &schema.Schema{
													Description: `Snmp`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"syslogs": &schema.Schema{
													Description: `Syslogs`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"other_device": &schema.Schema{
										Description: `Other Device`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"provisioning_state": &schema.Schema{
										Description: `Provisioning State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"resource_version": &schema.Schema{
										Description: `Resource Version`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"roles": &schema.Schema{
										Description: `Roles`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"save_wan_connectivity_details_only": &schema.Schema{
										Description: `Save Wan Connectivity Details Only`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"target_id_list": &schema.Schema{
										Description: `Target Id List`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"transit_networks": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id Ref`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"virtual_network": &schema.Schema{
										Description: `Virtual Network`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"wlan": &schema.Schema{
										Description: `Wlan`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaFabricBorderDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetBorderDeviceDetailFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetBorderDeviceDetailFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetBorderDeviceDetailFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBorderDeviceDetailFromSdaFabric", err,
				"Failure at GetBorderDeviceDetailFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetBorderDeviceDetailFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetBorderDeviceDetailFromSdaFabric response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["payload"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayload(item.Payload)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayload(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayload) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["auth_entity_id"] = item.AuthEntityID
	respItem["display_name"] = item.DisplayName
	respItem["auth_entity_class"] = item.AuthEntityClass
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["deploy_pending"] = item.DeployPending
	respItem["instance_version"] = item.InstanceVersion
	respItem["create_time"] = item.CreateTime
	respItem["deployed"] = boolPtrToString(item.Deployed)
	respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
	respItem["is_stale"] = boolPtrToString(item.IsStale)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["namespace"] = item.Namespace
	respItem["provisioning_state"] = item.ProvisioningState
	respItem["resource_version"] = item.ResourceVersion
	respItem["target_id_list"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadTargetIDList(item.TargetIDList)
	respItem["type"] = item.Type
	respItem["cfs_change_info"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadCfsChangeInfo(item.CfsChangeInfo)
	respItem["custom_provisions"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadCustomProvisions(item.CustomProvisions)
	respItem["configs"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadConfigs(item.Configs)
	respItem["managed_sites"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadManagedSites(item.ManagedSites)
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["roles"] = item.Roles
	respItem["save_wan_connectivity_details_only"] = boolPtrToString(item.SaveWanConnectivityDetailsOnly)
	respItem["site_id"] = item.SiteID
	respItem["akc_settings_cfs"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadAkcSettingsCfs(item.AkcSettingsCfs)
	respItem["device_interface_info"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceInterfaceInfo(item.DeviceInterfaceInfo)
	respItem["device_settings"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettings(item.DeviceSettings)
	respItem["network_wide_settings"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettings(item.NetworkWideSettings)
	respItem["other_device"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadOtherDevice(item.OtherDevice)
	respItem["transit_networks"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadTransitNetworks(item.TransitNetworks)
	respItem["virtual_network"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadVirtualNetwork(item.VirtualNetwork)
	respItem["wlan"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadWLAN(item.WLAN)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadTargetIDList(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTargetIDList) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadCfsChangeInfo(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadCustomProvisions(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCustomProvisions) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadConfigs(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadConfigs) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadManagedSites(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadManagedSites) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadAkcSettingsCfs(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceInterfaceInfo(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettings(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["deploy_pending"] = item.DeployPending
	respItem["instance_version"] = item.InstanceVersion
	respItem["connected_to"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsConnectedTo(item.ConnectedTo)
	respItem["cpu"] = item.CPU
	respItem["dhcp_enabled"] = boolPtrToString(item.DhcpEnabled)
	respItem["external_connectivity_ip_pool"] = item.ExternalConnectivityIPPool
	respItem["external_domain_routing_protocol"] = item.ExternalDomainRoutingProtocol
	respItem["internal_domain_protocol_number"] = item.InternalDomainProtocolNumber
	respItem["memory"] = item.Memory
	respItem["node_type"] = item.NodeType
	respItem["storage"] = item.Storage
	respItem["ext_connectivity_settings"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettings(item.ExtConnectivitySettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsConnectedTo(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettings(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_version"] = item.InstanceVersion
		respItem["external_domain_protocol_number"] = item.ExternalDomainProtocolNumber
		respItem["interface_uuid"] = item.InterfaceUUID
		respItem["policy_propagation_enabled"] = boolPtrToString(item.PolicyPropagationEnabled)
		respItem["policy_sgt_tag"] = item.PolicySgtTag
		respItem["l2_handoff"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL2Handoff(item.L2Handoff)
		respItem["l3_handoff"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3Handoff(item.L3Handoff)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL2Handoff(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3Handoff(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_version"] = item.InstanceVersion
		respItem["local_ip_address"] = item.LocalIPAddress
		respItem["remote_ip_address"] = item.RemoteIPAddress
		respItem["vlan_id"] = item.VLANID
		respItem["virtual_network"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork(item.VirtualNetwork)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id_ref"] = item.IDRef

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettings(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["deploy_pending"] = item.DeployPending
	respItem["instance_version"] = item.InstanceVersion
	respItem["aaa"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsAAA(item.AAA)
	respItem["cmx"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsCmx(item.Cmx)
	respItem["dhcp"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcp(item.Dhcp)
	respItem["dns"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNS(item.DNS)
	respItem["ldap"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsLdap(item.Ldap)
	respItem["native_vlan"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNativeVLAN(item.NativeVLAN)
	respItem["netflow"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNetflow(item.Netflow)
	respItem["ntp"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNtp(item.Ntp)
	respItem["snmp"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSNMP(item.SNMP)
	respItem["syslogs"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSyslogs(item.Syslogs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsAAA(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsAAA) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsCmx(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsCmx) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcp(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcp) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ip_address"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcpIPAddress(item.IPAddress)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcpIPAddress(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcpIPAddress) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["padded_address"] = item.PaddedAddress
	respItem["address_type"] = item.AddressType
	respItem["address"] = item.Address

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNS(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNS) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["domain_name"] = item.DomainName
		respItem["ip"] = flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNSIP(item.IP)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNSIP(item *dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNSIP) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["padded_address"] = item.PaddedAddress
	respItem["address_type"] = item.AddressType
	respItem["address"] = item.Address

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsLdap(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsLdap) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNativeVLAN(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNativeVLAN) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNetflow(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNetflow) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNtp(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNtp) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSNMP(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSNMP) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSyslogs(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSyslogs) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadOtherDevice(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadOtherDevice) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadTransitNetworks(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTransitNetworks) []map[string]interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadVirtualNetwork(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork) []interface{} {
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

func flattenSdaGetBorderDeviceDetailFromSdaFabricItemPayloadWLAN(items *[]dnacentersdkgo.ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadWLAN) []interface{} {
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
