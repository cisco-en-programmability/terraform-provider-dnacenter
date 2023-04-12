package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessAccesspointConfigurationSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Users can query the access point configuration information per device using the ethernet MAC address
`,

		ReadContext: dataSourceWirelessAccesspointConfigurationSummaryRead,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Description: `key query parameter. The ethernet MAC address of Access point
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_order_index": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"is_being_changed": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ordered_list_oeassoc_name": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"ordered_list_oeindex": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"admin_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_height": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"ap_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_entity_class": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"auth_entity_id": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"change_log_list": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"deploy_pending": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"eth_mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"failover_priority": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_created_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"instance_origin": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"internal_key": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"long_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"lazy_loaded_entities": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"led_brightness_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"led_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mesh_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"creation_order_index": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"is_being_changed": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ordered_list_oeassoc_name": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"ordered_list_oeindex": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"admin_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"antenna_angle": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"antenna_elev_angle": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"antenna_gain": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"antenna_pattern_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_entity_class": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"auth_entity_id": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"change_log_list": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"channel_assignment_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"channel_number": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"channel_width": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"clean_air_si": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"deploy_pending": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"if_type": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"if_type_value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"instance_origin": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"internal_key": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"long_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"url": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"lazy_loaded_entities": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_assignment_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"powerlevel": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_band": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"radio_role_assignment": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"slot_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"secondary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"secondary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tertiary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tertiary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessAccesspointConfigurationSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vKey := d.Get("key")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointConfiguration")
		queryParams1 := dnacentersdkgo.GetAccessPointConfigurationQueryParams{}

		queryParams1.Key = vKey.(string)

		response1, restyResp1, err := client.Wireless.GetAccessPointConfiguration(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAccessPointConfiguration", err,
				"Failure at GetAccessPointConfiguration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAccessPointConfigurationItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointConfiguration response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointConfigurationItem(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfiguration) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["instance_uuid"] = flattenWirelessGetAccessPointConfigurationItemInstanceUUID(item.InstanceUUID)
	respItem["instance_id"] = item.InstanceID
	respItem["auth_entity_id"] = flattenWirelessGetAccessPointConfigurationItemAuthEntityID(item.AuthEntityID)
	respItem["display_name"] = item.DisplayName
	respItem["auth_entity_class"] = flattenWirelessGetAccessPointConfigurationItemAuthEntityClass(item.AuthEntityClass)
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["ordered_list_oeindex"] = item.OrderedListOEIndex
	respItem["ordered_list_oeassoc_name"] = flattenWirelessGetAccessPointConfigurationItemOrderedListOEAssocName(item.OrderedListOEAssocName)
	respItem["creation_order_index"] = item.CreationOrderIndex
	respItem["is_being_changed"] = boolPtrToString(item.IsBeingChanged)
	respItem["deploy_pending"] = item.DeployPending
	respItem["instance_created_on"] = flattenWirelessGetAccessPointConfigurationItemInstanceCreatedOn(item.InstanceCreatedOn)
	respItem["instance_updated_on"] = flattenWirelessGetAccessPointConfigurationItemInstanceUpdatedOn(item.InstanceUpdatedOn)
	respItem["change_log_list"] = flattenWirelessGetAccessPointConfigurationItemChangeLogList(item.ChangeLogList)
	respItem["instance_origin"] = flattenWirelessGetAccessPointConfigurationItemInstanceOrigin(item.InstanceOrigin)
	respItem["lazy_loaded_entities"] = flattenWirelessGetAccessPointConfigurationItemLazyLoadedEntities(item.LazyLoadedEntities)
	respItem["instance_version"] = item.InstanceVersion
	respItem["admin_status"] = item.AdminStatus
	respItem["ap_height"] = item.ApHeight
	respItem["ap_mode"] = item.ApMode
	respItem["ap_name"] = item.ApName
	respItem["eth_mac"] = item.EthMac
	respItem["failover_priority"] = item.FailoverPriority
	respItem["led_brightness_level"] = item.LedBrightnessLevel
	respItem["led_status"] = item.LedStatus
	respItem["location"] = item.Location
	respItem["mac_address"] = item.MacAddress
	respItem["primary_controller_name"] = item.PrimaryControllerName
	respItem["primary_ip_address"] = item.PrimaryIPAddress
	respItem["secondary_controller_name"] = item.SecondaryControllerName
	respItem["secondary_ip_address"] = item.SecondaryIPAddress
	respItem["tertiary_controller_name"] = item.TertiaryControllerName
	respItem["tertiary_ip_address"] = item.TertiaryIPAddress
	respItem["mesh_dtos"] = flattenWirelessGetAccessPointConfigurationItemMeshDTOs(item.MeshDTOs)
	respItem["radio_dtos"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOs(item.RadioDTOs)
	respItem["internal_key"] = flattenWirelessGetAccessPointConfigurationItemInternalKey(item.InternalKey)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetAccessPointConfigurationItemInstanceUUID(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationInstanceUUID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemAuthEntityID(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationAuthEntityID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemAuthEntityClass(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationAuthEntityClass) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemOrderedListOEAssocName(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationOrderedListOEAssocName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemInstanceCreatedOn(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationInstanceCreatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemInstanceUpdatedOn(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationInstanceUpdatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemChangeLogList(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationChangeLogList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemInstanceOrigin(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationInstanceOrigin) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemLazyLoadedEntities(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationLazyLoadedEntities) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemMeshDTOs(items *[]dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationMeshDTOs) []interface{} {
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

func flattenWirelessGetAccessPointConfigurationItemRadioDTOs(items *[]dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceUUID(item.InstanceUUID)
		respItem["instance_id"] = item.InstanceID
		respItem["auth_entity_id"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsAuthEntityID(item.AuthEntityID)
		respItem["display_name"] = item.DisplayName
		respItem["auth_entity_class"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsAuthEntityClass(item.AuthEntityClass)
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["ordered_list_oeindex"] = item.OrderedListOEIndex
		respItem["ordered_list_oeassoc_name"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsOrderedListOEAssocName(item.OrderedListOEAssocName)
		respItem["creation_order_index"] = item.CreationOrderIndex
		respItem["is_being_changed"] = boolPtrToString(item.IsBeingChanged)
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_created_on"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceCreatedOn(item.InstanceCreatedOn)
		respItem["instance_updated_on"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceUpdatedOn(item.InstanceUpdatedOn)
		respItem["change_log_list"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsChangeLogList(item.ChangeLogList)
		respItem["instance_origin"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceOrigin(item.InstanceOrigin)
		respItem["lazy_loaded_entities"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsLazyLoadedEntities(item.LazyLoadedEntities)
		respItem["instance_version"] = item.InstanceVersion
		respItem["admin_status"] = item.AdminStatus
		respItem["antenna_angle"] = item.AntennaAngle
		respItem["antenna_elev_angle"] = item.AntennaElevAngle
		respItem["antenna_gain"] = item.AntennaGain
		respItem["antenna_pattern_name"] = item.AntennaPatternName
		respItem["channel_assignment_mode"] = item.ChannelAssignmentMode
		respItem["channel_number"] = item.ChannelNumber
		respItem["channel_width"] = item.ChannelWidth
		respItem["clean_air_si"] = item.CleanAirSI
		respItem["if_type"] = item.IfType
		respItem["if_type_value"] = item.IfTypeValue
		respItem["mac_address"] = item.MacAddress
		respItem["power_assignment_mode"] = item.PowerAssignmentMode
		respItem["powerlevel"] = item.Powerlevel
		respItem["radio_band"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioBand(item.RadioBand)
		respItem["radio_role_assignment"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioRoleAssignment(item.RadioRoleAssignment)
		respItem["slot_id"] = item.SlotID
		respItem["internal_key"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsInternalKey(item.InternalKey)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceUUID(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUUID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsAuthEntityID(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsAuthEntityClass(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityClass) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsOrderedListOEAssocName(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsOrderedListOEAssocName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceCreatedOn(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceCreatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceUpdatedOn(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUpdatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsChangeLogList(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsChangeLogList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsInstanceOrigin(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceOrigin) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsLazyLoadedEntities(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsLazyLoadedEntities) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioBand(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioBand) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioRoleAssignment(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioRoleAssignment) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsInternalKey(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsInternalKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["long_type"] = item.LongType
	respItem["url"] = item.URL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetAccessPointConfigurationItemInternalKey(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationInternalKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["long_type"] = item.LongType
	respItem["url"] = item.URL

	return []map[string]interface{}{
		respItem,
	}

}
