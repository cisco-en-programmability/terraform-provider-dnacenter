package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the
user, the devices that the user is connected to and the assurance issues that the user is impacted by
`,

		ReadContext: dataSourceClientEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_value": &schema.Schema{
				Description: `entity_value header parameter. Contains the actual value for the entity type that has been defined
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"issue_category": &schema.Schema{
				Description: `issueCategory header parameter. The category of the DNA event based on which the underlying issues need to be fetched
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ap_manager_interface_ip": &schema.Schema{
													Description: `Ap Manager Interface Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"associated_wlc_ip": &schema.Schema{
													Description: `Associated Wlc Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"boot_date_time": &schema.Schema{
													Description: `Boot Date Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"cisco360view": &schema.Schema{
													Description: `Cisco360view`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"collection_interval": &schema.Schema{
													Description: `Collection Interval`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"collection_status": &schema.Schema{
													Description: `Collection Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"error_code": &schema.Schema{
													Description: `Error Code`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"error_description": &schema.Schema{
													Description: `Error Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"family": &schema.Schema{
													Description: `Family`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"hostname": &schema.Schema{
													Description: `Hostname`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_uuid": &schema.Schema{
													Description: `Instance Uuid`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"interface_count": &schema.Schema{
													Description: `Interface Count`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"inventory_status_detail": &schema.Schema{
													Description: `Inventory Status Detail`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"last_update_time": &schema.Schema{
													Description: `Last Update Time`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"last_updated": &schema.Schema{
													Description: `Last Updated`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"line_card_count": &schema.Schema{
													Description: `Line Card Count`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"line_card_id": &schema.Schema{
													Description: `Line Card Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"location": &schema.Schema{
													Description: `Location`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"location_name": &schema.Schema{
													Description: `Location Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"mac_address": &schema.Schema{
													Description: `Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"management_ip_address": &schema.Schema{
													Description: `Management Ip Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"memory_size": &schema.Schema{
													Description: `Memory Size`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"neighbor_topology": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"links": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"label": &schema.Schema{
																			Description: `Label`,
																			Type:        schema.TypeList,
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"link_status": &schema.Schema{
																			Description: `Link Status`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"port_utilization": &schema.Schema{
																			Description: `Port Utilization`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"source": &schema.Schema{
																			Description: `Source`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"target": &schema.Schema{
																			Description: `Target`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},

															"nodes": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"clients": &schema.Schema{
																			Description: `Clients`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},

																		"count": &schema.Schema{
																			Description: `Count`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"description": &schema.Schema{
																			Description: `Description`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"device_type": &schema.Schema{
																			Description: `Device Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"fabric_group": &schema.Schema{
																			Description: `Fabric Group`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"family": &schema.Schema{
																			Description: `Family`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"health_score": &schema.Schema{
																			Description: `Health Score`,
																			Type:        schema.TypeList,
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"ip": &schema.Schema{
																			Description: `Ip`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"level": &schema.Schema{
																			Description: `Level`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},

																		"name": &schema.Schema{
																			Description: `Name`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"node_type": &schema.Schema{
																			Description: `Node Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"platform_id": &schema.Schema{
																			Description: `Platform Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"radio_frequency": &schema.Schema{
																			Description: `Radio Frequency`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"role": &schema.Schema{
																			Description: `Role`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"software_version": &schema.Schema{
																			Description: `Software Version`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"user_id": &schema.Schema{
																			Description: `User Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},

												"platform_id": &schema.Schema{
													Description: `Platform Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reachability_failure_reason": &schema.Schema{
													Description: `Reachability Failure Reason`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reachability_status": &schema.Schema{
													Description: `Reachability Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"role": &schema.Schema{
													Description: `Role`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"role_source": &schema.Schema{
													Description: `Role Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"serial_number": &schema.Schema{
													Description: `Serial Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"series": &schema.Schema{
													Description: `Series`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"snmp_contact": &schema.Schema{
													Description: `Snmp Contact`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"snmp_location": &schema.Schema{
													Description: `Snmp Location`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"software_version": &schema.Schema{
													Description: `Software Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"tag_count": &schema.Schema{
													Description: `Tag Count`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"tunnel_udp_port": &schema.Schema{
													Description: `Tunnel Udp Port`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"up_time": &schema.Schema{
													Description: `Up Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"waas_device_mode": &schema.Schema{
													Description: `Waas Device Mode`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"issue_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"issue": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"impacted_hosts": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connected_interface": &schema.Schema{
																Description: `Connected Interface`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"failed_attempts": &schema.Schema{
																Description: `Failed Attempts`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"host_name": &schema.Schema{
																Description: `Host Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"host_os": &schema.Schema{
																Description: `Host Os`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"host_type": &schema.Schema{
																Description: `Host Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"location": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"aps_impacted": &schema.Schema{
																			Description: `Aps Impacted`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"area": &schema.Schema{
																			Description: `Area`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"building": &schema.Schema{
																			Description: `Building`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"floor": &schema.Schema{
																			Description: `Floor`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"site_id": &schema.Schema{
																			Description: `Site Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"site_type": &schema.Schema{
																			Description: `Site Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},

															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ssid": &schema.Schema{
																Description: `Ssid`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"timestamp": &schema.Schema{
																Description: `Timestamp`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"issue_category": &schema.Schema{
													Description: `Issue Category`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_description": &schema.Schema{
													Description: `Issue Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_entity": &schema.Schema{
													Description: `Issue Entity`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_entity_value": &schema.Schema{
													Description: `Issue Entity Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_id": &schema.Schema{
													Description: `Issue Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_name": &schema.Schema{
													Description: `Issue Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_priority": &schema.Schema{
													Description: `Issue Priority`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_severity": &schema.Schema{
													Description: `Issue Severity`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_source": &schema.Schema{
													Description: `Issue Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_summary": &schema.Schema{
													Description: `Issue Summary`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_timestamp": &schema.Schema{
													Description: `Issue Timestamp`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"suggested_actions": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"message": &schema.Schema{
																Description: `Message`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"steps": &schema.Schema{
																Description: `Steps`,
																Type:        schema.TypeList,
																Computed:    true,
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
							},
						},

						"user_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"client_connection": &schema.Schema{
										Description: `Client Connection`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_device": &schema.Schema{
										Description: `Connected Device`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"connection_status": &schema.Schema{
										Description: `Connection Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"data_rate": &schema.Schema{
										Description: `Data Rate`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"health_score": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"health_type": &schema.Schema{
													Description: `Health Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reason": &schema.Schema{
													Description: `Reason`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"score": &schema.Schema{
													Description: `Score`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},

									"host_ip_v4": &schema.Schema{
										Description: `Host Ip V4`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_ip_v6": &schema.Schema{
										Description: `Host Ip V6`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"host_mac": &schema.Schema{
										Description: `Host Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_os": &schema.Schema{
										Description: `Host Os`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_type": &schema.Schema{
										Description: `Host Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_version": &schema.Schema{
										Description: `Host Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_count": &schema.Schema{
										Description: `Issue Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"last_updated": &schema.Schema{
										Description: `Last Updated`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"location": &schema.Schema{
										Description: `Location`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_type": &schema.Schema{
										Description: `Sub Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceClientEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")
	vIssueCategory := d.Get("issue_category")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetClientEnrichmentDetails")

		headerParams1 := dnacentersdkgo.GetClientEnrichmentDetailsHeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		headerParams1.IssueCategory = vIssueCategory.(string)

		response1, restyResp1, err := client.Clients.GetClientEnrichmentDetails(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetClientEnrichmentDetails", err,
				"Failure at GetClientEnrichmentDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenClientsGetClientEnrichmentDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetClientEnrichmentDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsGetClientEnrichmentDetailsItems(items *dnacentersdkgo.ResponseClientsGetClientEnrichmentDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["user_details"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetails(item.UserDetails)
		respItem["connected_device"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDevice(item.ConnectedDevice)
		respItem["issue_details"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetails(item.IssueDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetails(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["connection_status"] = item.ConnectionStatus
	respItem["host_type"] = item.HostType
	respItem["user_id"] = item.UserID
	respItem["host_name"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostName(item.HostName)
	respItem["host_os"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostOs(item.HostOs)
	respItem["host_version"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostVersion(item.HostVersion)
	respItem["sub_type"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSubType(item.SubType)
	respItem["last_updated"] = item.LastUpdated
	respItem["health_score"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHealthScore(item.HealthScore)
	respItem["host_mac"] = item.HostMac
	respItem["host_ip_v4"] = item.HostIPV4
	respItem["host_ip_v6"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostIPV6(item.HostIPV6)
	respItem["auth_type"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsAuthType(item.AuthType)
	respItem["vlan_id"] = item.VLANID
	respItem["ssid"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSSID(item.SSID)
	respItem["location"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsLocation(item.Location)
	respItem["client_connection"] = item.ClientConnection
	respItem["connected_device"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsConnectedDevice(item.ConnectedDevice)
	respItem["issue_count"] = item.IssueCount
	respItem["rssi"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsRssi(item.Rssi)
	respItem["snr"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSnr(item.Snr)
	respItem["data_rate"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsDataRate(item.DataRate)
	respItem["port"] = flattenClientsGetClientEnrichmentDetailsItemsUserDetailsPort(item.Port)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostName(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostOs(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostOs) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostVersion(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSubType(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSubType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHealthScore(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHealthScore) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["health_type"] = item.HealthType
		respItem["reason"] = item.Reason
		respItem["score"] = item.Score
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsHostIPV6(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostIPV6) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsAuthType(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsAuthType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSSID(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSSID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsLocation(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsConnectedDevice(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsConnectedDevice) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsRssi(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsRssi) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsSnr(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSnr) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsDataRate(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsDataRate) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsUserDetailsPort(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsUserDetailsPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDevice(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_details"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetails(item.DeviceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetails(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["family"] = item.Family
	respItem["type"] = item.Type
	respItem["location"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLocation(item.Location)
	respItem["error_code"] = item.ErrorCode
	respItem["mac_address"] = item.MacAddress
	respItem["role"] = item.Role
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsBootDateTime(item.BootDateTime)
	respItem["collection_status"] = item.CollectionStatus
	respItem["interface_count"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsInterfaceCount(item.InterfaceCount)
	respItem["line_card_count"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLineCardCount(item.LineCardCount)
	respItem["line_card_id"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLineCardID(item.LineCardID)
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["tunnel_udp_port"] = item.TunnelUDPPort
	respItem["waas_device_mode"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item.WaasDeviceMode)
	respItem["series"] = item.Series
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["collection_interval"] = item.CollectionInterval
	respItem["serial_number"] = item.SerialNumber
	respItem["software_version"] = item.SoftwareVersion
	respItem["role_source"] = item.RoleSource
	respItem["hostname"] = item.Hostname
	respItem["up_time"] = item.UpTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["error_description"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsErrorDescription(item.ErrorDescription)
	respItem["location_name"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLocationName(item.LocationName)
	respItem["tag_count"] = item.TagCount
	respItem["last_updated"] = item.LastUpdated
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["id"] = item.ID
	respItem["neighbor_topology"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopology(item.NeighborTopology)
	respItem["cisco360view"] = item.Cisco360View

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLocation(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsBootDateTime(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsBootDateTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsInterfaceCount(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsInterfaceCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLineCardCount(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLineCardID(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsErrorDescription(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsLocationName(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopology(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nodes"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodes(item.Nodes)
		respItem["links"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinks(item.Links)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodes(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["role"] = item.Role
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["device_type"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType(item.DeviceType)
		respItem["platform_id"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID(item.PlatformID)
		respItem["family"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily(item.Family)
		respItem["ip"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP(item.IP)
		respItem["software_version"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion(item.SoftwareVersion)
		respItem["user_id"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID(item.UserID)
		respItem["node_type"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType(item.NodeType)
		respItem["radio_frequency"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency(item.RadioFrequency)
		respItem["clients"] = item.Clients
		respItem["count"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount(item.Count)
		respItem["health_score"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore(item.HealthScore)
		respItem["level"] = item.Level
		respItem["fabric_group"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup(item.FabricGroup)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinks(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["source"] = item.Source
		respItem["link_status"] = item.LinkStatus
		respItem["label"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel(item.Label)
		respItem["target"] = item.Target
		respItem["id"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksID(item.ID)
		respItem["port_utilization"] = flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization(item.PortUtilization)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksID(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetails(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["issue"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssue(item.Issue)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssue(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["issue_source"] = item.IssueSource
		respItem["issue_category"] = item.IssueCategory
		respItem["issue_name"] = item.IssueName
		respItem["issue_description"] = item.IssueDescription
		respItem["issue_entity"] = item.IssueEntity
		respItem["issue_entity_value"] = item.IssueEntityValue
		respItem["issue_severity"] = item.IssueSeverity
		respItem["issue_priority"] = item.IssuePriority
		respItem["issue_summary"] = item.IssueSummary
		respItem["issue_timestamp"] = item.IssueTimestamp
		respItem["suggested_actions"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueSuggestedActions(item.SuggestedActions)
		respItem["impacted_hosts"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHosts(item.ImpactedHosts)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueSuggestedActions(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueSuggestedActionsSteps(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHosts(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHosts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["host_type"] = item.HostType
		respItem["host_name"] = item.HostName
		respItem["host_os"] = item.HostOs
		respItem["ssid"] = item.SSID
		respItem["connected_interface"] = item.ConnectedInterface
		respItem["mac_address"] = item.MacAddress
		respItem["failed_attempts"] = item.FailedAttempts
		respItem["location"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocation(item.Location)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocation(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_id"] = item.SiteID
	respItem["site_type"] = item.SiteType
	respItem["area"] = item.Area
	respItem["building"] = item.Building
	respItem["floor"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocationFloor(item.Floor)
	respItem["aps_impacted"] = flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocationApsImpacted(item.ApsImpacted)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocationFloor(item *dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationFloor) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsItemsIssueDetailsIssueImpactedHostsLocationApsImpacted(items *[]dnacentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationApsImpacted) []interface{} {
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
