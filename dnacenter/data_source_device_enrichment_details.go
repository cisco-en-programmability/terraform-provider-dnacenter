package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Enriches a given network device context (device id or device Mac Address or device management IP address) with details
about the device and neighbor topology
`,

		ReadContext: dataSourceDeviceEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
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

			"items": &schema.Schema{
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
																Type:        schema.TypeString,
																Computed:    true,
															},

															"connected_device": &schema.Schema{
																Description: `Connected Device`,
																Type:        schema.TypeString,
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
																Type:        schema.TypeInt,
																Computed:    true,
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
		},
	}
}

func dataSourceDeviceEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceEnrichmentDetails")

		headerParams1 := dnacentersdkgo.GetDeviceEnrichmentDetailsHeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		response1, restyResp1, err := client.Devices.GetDeviceEnrichmentDetails(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceEnrichmentDetails", err,
				"Failure at GetDeviceEnrichmentDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceEnrichmentDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceEnrichmentDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceEnrichmentDetailsItems(items *dnacentersdkgo.ResponseDevicesGetDeviceEnrichmentDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_details"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetails(item.DeviceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetails(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["family"] = item.Family
	respItem["type"] = item.Type
	respItem["location"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsLocation(item.Location)
	respItem["error_code"] = item.ErrorCode
	respItem["mac_address"] = item.MacAddress
	respItem["role"] = item.Role
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = item.BootDateTime
	respItem["collection_status"] = item.CollectionStatus
	respItem["interface_count"] = item.InterfaceCount
	respItem["line_card_count"] = item.LineCardCount
	respItem["line_card_id"] = item.LineCardID
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["tunnel_udp_port"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsTunnelUDPPort(item.TunnelUDPPort)
	respItem["waas_device_mode"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsWaasDeviceMode(item.WaasDeviceMode)
	respItem["series"] = item.Series
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["collection_interval"] = item.CollectionInterval
	respItem["serial_number"] = item.SerialNumber
	respItem["software_version"] = item.SoftwareVersion
	respItem["role_source"] = item.RoleSource
	respItem["hostname"] = item.Hostname
	respItem["up_time"] = item.UpTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["error_description"] = item.ErrorDescription
	respItem["location_name"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsLocationName(item.LocationName)
	respItem["tag_count"] = item.TagCount
	respItem["last_updated"] = item.LastUpdated
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["id"] = item.ID
	respItem["neighbor_topology"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopology(item.NeighborTopology)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsLocation(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsTunnelUDPPort(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsWaasDeviceMode(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsLocationName(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopology(items *[]dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nodes"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodes(item.Nodes)
		respItem["links"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinks(item.Links)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodes(items *[]dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes) []map[string]interface{} {
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
		respItem["device_type"] = item.DeviceType
		respItem["platform_id"] = item.PlatformID
		respItem["family"] = item.Family
		respItem["ip"] = item.IP
		respItem["software_version"] = item.SoftwareVersion
		respItem["user_id"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesUserID(item.UserID)
		respItem["node_type"] = item.NodeType
		respItem["radio_frequency"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesRadioFrequency(item.RadioFrequency)
		respItem["clients"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesClients(item.Clients)
		respItem["count"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesCount(item.Count)
		respItem["health_score"] = item.HealthScore
		respItem["level"] = item.Level
		respItem["fabric_group"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesFabricGroup(item.FabricGroup)
		respItem["connected_device"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesConnectedDevice(item.ConnectedDevice)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesUserID(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesRadioFrequency(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesClients(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesCount(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesFabricGroup(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyNodesConnectedDevice(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinks(items *[]dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["source"] = item.Source
		respItem["link_status"] = item.LinkStatus
		respItem["label"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksLabel(item.Label)
		respItem["target"] = item.Target
		respItem["id"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksID(item.ID)
		respItem["port_utilization"] = flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksPortUtilization(item.PortUtilization)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksLabel(items *[]dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel) []interface{} {
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

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksID(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsItemsDeviceDetailsNeighborTopologyLinksPortUtilization(item *dnacentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
