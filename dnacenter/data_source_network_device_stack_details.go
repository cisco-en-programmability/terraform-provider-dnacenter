package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceStackDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieves complete stack details for given device ID
`,

		ReadContext: dataSourceNetworkDeviceStackDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_id": &schema.Schema{
							Description: `Device ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"stack_port_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"is_synch_ok": &schema.Schema{
										Description: `If link partner sends valid protocol message
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_active": &schema.Schema{
										Description: `If stack port is in same state as link partner
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_ok": &schema.Schema{
										Description: `If link is stable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the stack port
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"neighbor_port": &schema.Schema{
										Description: `Neighbor's member number and stack port number
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nr_link_ok_changes": &schema.Schema{
										Description: `Relative stability of the link
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"stack_cable_length_info": &schema.Schema{
										Description: `Cable length
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"stack_port_oper_status_info": &schema.Schema{
										Description: `Port opearation status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"switch_port": &schema.Schema{
										Description: `Member number and stack port number
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"stack_switch_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ent_physical_index": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"hw_priority": &schema.Schema{
										Description: `Hardware priority of the switch
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Description: `Mac address of the switch
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"num_next_reload": &schema.Schema{
										Description: `Stack member number to be used in next reload
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"platform_id": &schema.Schema{
										Description: `Platform Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Description: `Function of the switch
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Serial number
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_image": &schema.Schema{
										Description: `Software image type running on the switch
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"stack_member_number": &schema.Schema{
										Description: `Switch member number
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `Current state of the switch
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"switch_priority": &schema.Schema{
										Description: `Priority of the switch
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"svl_switch_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dad_protocol": &schema.Schema{
										Description: `Stackwise virtual dual active detection config
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dad_recovery_reload_enabled": &schema.Schema{
										Description: `If dad recovery reload enabled
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"domain_number": &schema.Schema{
										Description: `Stackwise virtual switch domain number
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"in_dad_recovery_mode": &schema.Schema{
										Description: `If in dad recovery mode
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sw_virtual_status": &schema.Schema{
										Description: `Stackwise virtual status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"switch_members": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"bandwidth": &schema.Schema{
													Description: `Bandwidth
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"svl_member_end_points": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"svl_member_end_point_ports": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"svl_protocol_status": &schema.Schema{
																			Description: `Stackwise virtual protocol status
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"sw_local_interface": &schema.Schema{
																			Description: `Stackwise virtual local interface
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"sw_remote_interface": &schema.Schema{
																			Description: `Stackwise virtual remote interface
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"svl_number": &schema.Schema{
																Description: `Stackwise virtual link number
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"svl_status": &schema.Schema{
																Description: `Stackwise virtual status
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"svl_member_number": &schema.Schema{
													Description: `Switch member number
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"svl_member_pep_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dad_enabled": &schema.Schema{
																Description: `If dadInterface is configured for dual active detection
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"dad_interface_name": &schema.Schema{
																Description: `Interface for dual active detection
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceStackDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetStackDetailsForDevice")
		vvDeviceID := vDeviceID.(string)

		response1, restyResp1, err := client.Devices.GetStackDetailsForDevice(vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetStackDetailsForDevice", err,
				"Failure at GetStackDetailsForDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetStackDetailsForDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetStackDetailsForDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetStackDetailsForDeviceItem(item *dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_id"] = item.DeviceID
	respItem["stack_port_info"] = flattenDevicesGetStackDetailsForDeviceItemStackPortInfo(item.StackPortInfo)
	respItem["stack_switch_info"] = flattenDevicesGetStackDetailsForDeviceItemStackSwitchInfo(item.StackSwitchInfo)
	respItem["svl_switch_info"] = flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfo(item.SvlSwitchInfo)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetStackDetailsForDeviceItemStackPortInfo(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["is_synch_ok"] = item.IsSynchOk
		respItem["link_active"] = boolPtrToString(item.LinkActive)
		respItem["link_ok"] = boolPtrToString(item.LinkOk)
		respItem["name"] = item.Name
		respItem["neighbor_port"] = item.NeighborPort
		respItem["nr_link_ok_changes"] = item.NrLinkOkChanges
		respItem["stack_cable_length_info"] = item.StackCableLengthInfo
		respItem["stack_port_oper_status_info"] = item.StackPortOperStatusInfo
		respItem["switch_port"] = item.SwitchPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemStackSwitchInfo(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ent_physical_index"] = item.EntPhysicalIndex
		respItem["hw_priority"] = item.HwPriority
		respItem["mac_address"] = item.MacAddress
		respItem["num_next_reload"] = item.NumNextReload
		respItem["platform_id"] = item.PlatformID
		respItem["role"] = item.Role
		respItem["serial_number"] = item.SerialNumber
		respItem["software_image"] = item.SoftwareImage
		respItem["stack_member_number"] = item.StackMemberNumber
		respItem["state"] = item.State
		respItem["switch_priority"] = item.SwitchPriority
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfo(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dad_protocol"] = item.DadProtocol
		respItem["dad_recovery_reload_enabled"] = boolPtrToString(item.DadRecoveryReloadEnabled)
		respItem["domain_number"] = item.DomainNumber
		respItem["in_dad_recovery_mode"] = boolPtrToString(item.InDadRecoveryMode)
		respItem["sw_virtual_status"] = item.SwVirtualStatus
		respItem["switch_members"] = flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembers(item.SwitchMembers)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembers(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bandwidth"] = item.Bandwidth
		respItem["svl_member_end_points"] = flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberEndPoints(item.SvlMemberEndPoints)
		respItem["svl_member_number"] = item.SvlMemberNumber
		respItem["svl_member_pep_settings"] = flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberPepSettings(item.SvlMemberPepSettings)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberEndPoints(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["svl_member_end_point_ports"] = flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts(item.SvlMemberEndPointPorts)
		respItem["svl_number"] = item.SvlNumber
		respItem["svl_status"] = item.SvlStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["svl_protocol_status"] = item.SvlProtocolStatus
		respItem["sw_local_interface"] = item.SwLocalInterface
		respItem["sw_remote_interface"] = item.SwRemoteInterface
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetStackDetailsForDeviceItemSvlSwitchInfoSwitchMembersSvlMemberPepSettings(items *[]dnacentersdkgo.ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dad_enabled"] = boolPtrToString(item.DadEnabled)
		respItem["dad_interface_name"] = item.DadInterfaceName
		respItems = append(respItems, respItem)
	}
	return respItems
}
