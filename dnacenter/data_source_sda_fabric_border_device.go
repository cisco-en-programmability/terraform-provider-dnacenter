package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Gets border device detail from SDA Fabric
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
		log.Printf("[DEBUG] Selected method 1: GetsBorderDeviceDetailFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetsBorderDeviceDetailFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetsBorderDeviceDetailFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetsBorderDeviceDetailFromSdaFabric", err,
				"Failure at GetsBorderDeviceDetailFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetsBorderDeviceDetailFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsBorderDeviceDetailFromSdaFabric response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItem(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["payload"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayload(item.Payload)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayload(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayload) []map[string]interface{} {
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
	respItem["target_id_list"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadTargetIDList(item.TargetIDList)
	respItem["type"] = item.Type
	respItem["cfs_change_info"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadCfsChangeInfo(item.CfsChangeInfo)
	respItem["custom_provisions"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadCustomProvisions(item.CustomProvisions)
	respItem["configs"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadConfigs(item.Configs)
	respItem["managed_sites"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadManagedSites(item.ManagedSites)
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["roles"] = item.Roles
	respItem["save_wan_connectivity_details_only"] = boolPtrToString(item.SaveWanConnectivityDetailsOnly)
	respItem["site_id"] = item.SiteID
	respItem["akc_settings_cfs"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadAkcSettingsCfs(item.AkcSettingsCfs)
	respItem["device_interface_info"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceInterfaceInfo(item.DeviceInterfaceInfo)
	respItem["device_settings"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettings(item.DeviceSettings)
	respItem["network_wide_settings"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettings(item.NetworkWideSettings)
	respItem["other_device"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadOtherDevice(item.OtherDevice)
	respItem["transit_networks"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadTransitNetworks(item.TransitNetworks)
	respItem["virtual_network"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadVirtualNetwork(item.VirtualNetwork)
	respItem["wlan"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadWLAN(item.WLAN)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadTargetIDList(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTargetIDList) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadCfsChangeInfo(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadCustomProvisions(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCustomProvisions) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadConfigs(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadConfigs) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadManagedSites(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadManagedSites) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadAkcSettingsCfs(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceInterfaceInfo(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettings(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettings) []map[string]interface{} {
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
	respItem["connected_to"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsConnectedTo(item.ConnectedTo)
	respItem["cpu"] = item.CPU
	respItem["dhcp_enabled"] = boolPtrToString(item.DhcpEnabled)
	respItem["external_connectivity_ip_pool"] = item.ExternalConnectivityIPPool
	respItem["external_domain_routing_protocol"] = item.ExternalDomainRoutingProtocol
	respItem["internal_domain_protocol_number"] = item.InternalDomainProtocolNumber
	respItem["memory"] = item.Memory
	respItem["node_type"] = item.NodeType
	respItem["storage"] = item.Storage
	respItem["ext_connectivity_settings"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettings(item.ExtConnectivitySettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsConnectedTo(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettings(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings) []map[string]interface{} {
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
		respItem["l2_handoff"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL2Handoff(item.L2Handoff)
		respItem["l3_handoff"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3Handoff(item.L3Handoff)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL2Handoff(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3Handoff(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff) []map[string]interface{} {
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
		respItem["virtual_network"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork(item.VirtualNetwork)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id_ref"] = item.IDRef

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettings(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettings) []map[string]interface{} {
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
	respItem["aaa"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsAAA(item.AAA)
	respItem["cmx"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsCmx(item.Cmx)
	respItem["dhcp"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcp(item.Dhcp)
	respItem["dns"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNS(item.DNS)
	respItem["ldap"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsLdap(item.Ldap)
	respItem["native_vlan"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNativeVLAN(item.NativeVLAN)
	respItem["netflow"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNetflow(item.Netflow)
	respItem["ntp"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNtp(item.Ntp)
	respItem["snmp"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSNMP(item.SNMP)
	respItem["syslogs"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSyslogs(item.Syslogs)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsAAA(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsAAA) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsCmx(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsCmx) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcp(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcp) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ip_address"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcpIPAddress(item.IPAddress)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDhcpIPAddress(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcpIPAddress) []map[string]interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNS(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNS) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["domain_name"] = item.DomainName
		respItem["ip"] = flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNSIP(item.IP)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsDNSIP(item *dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNSIP) []map[string]interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsLdap(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsLdap) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNativeVLAN(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNativeVLAN) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNetflow(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNetflow) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsNtp(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNtp) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSNMP(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSNMP) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadNetworkWideSettingsSyslogs(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSyslogs) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadOtherDevice(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadOtherDevice) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadTransitNetworks(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTransitNetworks) []map[string]interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadVirtualNetwork(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork) []interface{} {
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

func flattenSdaGetsBorderDeviceDetailFromSdaFabricItemPayloadWLAN(items *[]dnacentersdkgo.ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadWLAN) []interface{} {
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
