package dnacenter

import (
	"context"
	"log"
	"reflect"
	"time"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricBorderDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Adds border device in SDA Fabric

- Deletes border device from SDA Fabric
`,

		CreateContext: resourceSdaFabricBorderDeviceCreate,
		ReadContext:   resourceSdaFabricBorderDeviceRead,
		UpdateContext: resourceSdaFabricBorderDeviceUpdate,
		DeleteContext: resourceSdaFabricBorderDeviceDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"border_session_type": &schema.Schema{
							Description: `Border Session Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"connected_to_internet": &schema.Schema{
							Description: `Connected to Internet
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_autonomou_system_number": &schema.Schema{
							Description: `External Autonomous System Number  will be used to automate IP routing between Border Node and remote peer (e.g.,1-65535)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_connectivity_ip_pool_name": &schema.Schema{
							Description: `IP pool to use to automate IP routing between the border node and remote peer.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_connectivity_settings": &schema.Schema{
							Description: `External Connectivity Settings information of L3 Handoff
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"external_domain_routing_protocol_name": &schema.Schema{
							Description: `External Domain Routing Protocol  Name. (Example: BGP)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"internal_autonomou_system_number": &schema.Schema{
							Description: `Internal Autonomouns System Number used by border node to communicate with remote peer (e.g.,1-65535)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"l3_handoff": &schema.Schema{
							Description: `L3 Handoff information
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy for device location(site should be fabric site)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network": &schema.Schema{
							Description: `Virtual Network information of L3 Hand off
`,
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name assigned to site
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id (e.g.,2-4096 except for reserved VLANs (1002-1005, 2046, 4095))
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricBorderDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vDeviceManagementIPAddress := resourceItem["device_management_ip_address"]
	vvDeviceManagementIPAddress := interfaceToString(vDeviceManagementIPAddress)

	queryParams1 := dnacentersdkgo.GetsBorderDeviceDetailFromSdaFabricQueryParams{}

	queryParams1.DeviceManagementIPAddress = vvDeviceManagementIPAddress

	getResponse2, restyResp1, err := client.Sda.GetsBorderDeviceDetailFromSdaFabric(&queryParams1)

	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricBorderDeviceRead(ctx, d, m)
	}

	response1, restyResp1, err := client.Sda.AddsBorderDeviceInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddsBorderDeviceInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddsBorderDeviceInSdaFabric", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddsBorderDeviceInSdaFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricBorderDeviceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetsBorderDeviceDetailFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetsBorderDeviceDetailFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		response1, restyResp1, err := client.Sda.GetsBorderDeviceDetailFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetsBorderDeviceDetailFromSdaFabric", err,
			// 	"Failure at GetsBorderDeviceDetailFromSdaFabric, unexpected response", ""))
			// return diags
			d.SetId("")
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
		return diags

	}
	return diags
}

func resourceSdaFabricBorderDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricBorderDeviceRead(ctx, d, m)
}

func resourceSdaFabricBorderDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	queryParams1 := dnacentersdkgo.GetsBorderDeviceDetailFromSdaFabricQueryParams{}
	queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

	item, restyResp1, err := client.Sda.GetsBorderDeviceDetailFromSdaFabric(&queryParams1)

	if err != nil || item == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetsBorderDeviceDetailFromSdaFabric", err,
		"Failure at GetsBorderDeviceDetailFromSdaFabric, unexpected response", ""))*/
		d.SetId("")
		return diags
	}

	queryParams2 := dnacentersdkgo.DeletesBorderDeviceFromSdaFabricQueryParams{}
	queryParams2.DeviceManagementIPAddress = vDeviceManagementIPAddress
	response1, restyResp1, err := client.Sda.DeletesBorderDeviceFromSdaFabric(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesBorderDeviceFromSdaFabric", err, restyResp1.String(),
				"Failure at DeletesBorderDeviceFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesBorderDeviceFromSdaFabric", err,
			"Failure at DeletesBorderDeviceFromSdaFabric, unexpected response", ""))
		return diags
	}

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeletesBorderDeviceFromSdaFabric", err))
			return diags
		}
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_domain_routing_protocol_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_domain_routing_protocol_name")))) {
		request.ExternalDomainRoutingProtocolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_ip_pool_name")))) {
		request.ExternalConnectivityIPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".internal_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".internal_autonomou_system_number")))) {
		request.InternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".border_session_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".border_session_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".border_session_type")))) {
		request.BorderSessionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_to_internet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_to_internet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_to_internet")))) {
		request.ConnectedToInternet = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_connectivity_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_connectivity_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_connectivity_settings")))) {
		request.ExternalConnectivitySettings = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricExternalConnectivitySettings(ctx, key+".external_connectivity_settings.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_autonomou_system_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_autonomou_system_number")))) {
		request.ExternalAutonomouSystemNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".l3_handoff")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".l3_handoff")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".l3_handoff")))) {
		request.L3Handoff = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricL3Handoff(ctx, key+".l3_handoff.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network")))) {
		request.VirtualNetwork = expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricVirtualNetwork(ctx, key+".virtual_network.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricExternalConnectivitySettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricL3Handoff(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricL3Handoff {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricL3Handoff
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSdaFabricBorderDeviceAddsBorderDeviceInSdaFabricVirtualNetwork(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricVirtualNetwork {
	var request dnacentersdkgo.RequestSdaAddsBorderDeviceInSdaFabricVirtualNetwork
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
