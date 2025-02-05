package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterfaceNetworkDeviceDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns interface by specified device Id and interface name
`,

		ReadContext: dataSourceInterfaceNetworkDeviceDetailRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Interface name
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"addresses": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"address": &schema.Schema{
																Description: `IP address of the interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"ip_mask": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"address": &schema.Schema{
																Description: `IP Mask of the interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"is_inverse_mask": &schema.Schema{
													Description: `Inverse Mask of the IP address is enabled or not
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"type": &schema.Schema{
										Description: `Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"admin_status": &schema.Schema{
							Description: `Admin status as ('UP'/'DOWN')
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"class_name": &schema.Schema{
							Description: `Classifies the port as switch port ,loopback interface etc.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description for the Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"duplex": &schema.Schema{
							Description: `Interface duplex as AutoNegotiate or FullDuplex
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"if_index": &schema.Schema{
							Description: `Interface index
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id of the Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_type": &schema.Schema{
							Description: `Interface type as Physical or Virtual
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_address": &schema.Schema{
							Description: `IPV4 Address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ipv4_mask": &schema.Schema{
							Description: `IPV4 Mask of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"isis_support": &schema.Schema{
							Description: `Flag for ISIS enabled / disabled
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_incoming_packet_time": &schema.Schema{
							Description: `Time, in milliseconds since UNIX epoch, when the last packet was received on this interface
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"last_outgoing_packet_time": &schema.Schema{
							Description: `Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Description: `Time when the device interface info last got updated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_compute_element": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_compute_element_url": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_network_element": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_network_element_url": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_id": &schema.Schema{
							Description: `ID of physical interface mapped with the virtual interface of WLC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mapped_physical_interface_name": &schema.Schema{
							Description: `Physical interface name mapped with the virtual interface of WLC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"media_type": &schema.Schema{
							Description: `Media Type of the interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mtu": &schema.Schema{
							Description: `MTU Information of Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name for the interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"native_vlan_id": &schema.Schema{
							Description: `Vlan to receive untagged frames on trunk port
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"networkdevice_id": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ospf_support": &schema.Schema{
							Description: `Flag for OSPF enabled / disabled
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pid": &schema.Schema{
							Description: `Platform ID of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_mode": &schema.Schema{
							Description: `Port mode as access, trunk, routed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_name": &schema.Schema{
							Description: `Interface name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_type": &schema.Schema{
							Description: `Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"poweroverethernet": &schema.Schema{
							Description: `This is internal attribute.  Not to be used.  Deprecated
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_no": &schema.Schema{
							Description: `Serial number of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Description: `Series of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"speed": &schema.Schema{
							Description: `Speed of the interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Interface status as Down / Up
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan ID of interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"voice_vlan": &schema.Schema{
							Description: `Vlan information of the interface
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

func dataSourceInterfaceNetworkDeviceDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")
	vName := d.Get("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInterfaceDetailsByDeviceIDAndInterfaceName")
		vvDeviceID := vDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams{}

		queryParams1.Name = vName.(string)

		response1, restyResp1, err := client.Devices.GetInterfaceDetailsByDeviceIDAndInterfaceName(vvDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetInterfaceDetailsByDeviceIDAndInterfaceName", err,
				"Failure at GetInterfaceDetailsByDeviceIDAndInterfaceName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInterfaceDetailsByDeviceIDAndInterfaceName response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItem(item *dnacentersdkgo.ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["addresses"] = flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddresses(item.Addresses)
	respItem["admin_status"] = item.AdminStatus
	respItem["class_name"] = item.ClassName
	respItem["description"] = item.Description
	respItem["name"] = item.Name
	respItem["device_id"] = item.DeviceID
	respItem["duplex"] = item.Duplex
	respItem["id"] = item.ID
	respItem["if_index"] = item.IfIndex
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["interface_type"] = item.InterfaceType
	respItem["ipv4_address"] = item.IPv4Address
	respItem["ipv4_mask"] = item.IPv4Mask
	respItem["isis_support"] = item.IsisSupport
	respItem["last_outgoing_packet_time"] = item.LastOutgoingPacketTime
	respItem["last_incoming_packet_time"] = item.LastIncomingPacketTime
	respItem["last_updated"] = item.LastUpdated
	respItem["mac_address"] = item.MacAddress
	respItem["mapped_physical_interface_id"] = item.MappedPhysicalInterfaceID
	respItem["mapped_physical_interface_name"] = item.MappedPhysicalInterfaceName
	respItem["media_type"] = item.MediaType
	respItem["mtu"] = item.Mtu
	respItem["native_vlan_id"] = item.NativeVLANID
	respItem["ospf_support"] = item.OspfSupport
	respItem["pid"] = item.Pid
	respItem["port_mode"] = item.PortMode
	respItem["port_name"] = item.PortName
	respItem["port_type"] = item.PortType
	respItem["serial_no"] = item.SerialNo
	respItem["series"] = item.Series
	respItem["speed"] = item.Speed
	respItem["status"] = item.Status
	respItem["vlan_id"] = item.VLANID
	respItem["voice_vlan"] = item.VoiceVLAN
	respItem["poweroverethernet"] = item.Poweroverethernet
	respItem["networkdevice_id"] = item.NetworkdeviceID
	respItem["managed_compute_element"] = item.ManagedComputeElement
	respItem["managed_network_element"] = item.ManagedNetworkElement
	respItem["managed_network_element_url"] = item.ManagedNetworkElementURL
	respItem["managed_compute_element_url"] = item.ManagedComputeElementURL
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddresses(items *[]dnacentersdkgo.ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddresses) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["address"] = flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddress(item.Address)
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddress(item *dnacentersdkgo.ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddress) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ip_address"] = flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddressIPAddress(item.IPAddress)
	respItem["ip_mask"] = flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddressIPMask(item.IPMask)
	respItem["is_inverse_mask"] = boolPtrToString(item.IsInverseMask)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddressIPAddress(item *dnacentersdkgo.ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPAddress) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["address"] = item.Address

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameItemAddressesAddressIPMask(item *dnacentersdkgo.ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPMask) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["address"] = item.Address

	return []map[string]interface{}{
		respItem,
	}

}
