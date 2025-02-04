package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTopologyLayer3() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Topology.

- Returns the Layer 3 network topology by routing protocol
`,

		ReadContext: dataSourceTopologyLayer3Read,
		Schema: map[string]*schema.Schema{
			"topology_type": &schema.Schema{
				Description: `topologyType path parameter. Type of topology(OSPF,ISIS,etc)
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `[Deprecated]
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"links": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_info": &schema.Schema{
										Description: `Additional information about the link
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"end_port_id": &schema.Schema{
										Description: `Device port ID corresponding to the end device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_ipv4_address": &schema.Schema{
										Description: `Interface port IPv4 address corresponding to the end device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_ipv4_mask": &schema.Schema{
										Description: `Interface port IPv4 mask corresponding to the end device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_name": &schema.Schema{
										Description: `Interface port name corresponding to the end device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_speed": &schema.Schema{
										Description: `Interface port speed corresponding to end device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"grey_out": &schema.Schema{
										Description: `Indicates if the link is greyed out
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the link
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_status": &schema.Schema{
										Description: `Indicates whether link is up or down
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"source": &schema.Schema{
										Description: `Device ID corresponding to the source device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_id": &schema.Schema{
										Description: `Device port ID corresponding to start device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_ipv4_address": &schema.Schema{
										Description: `Interface port IPv4 address corresponding to start device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_ipv4_mask": &schema.Schema{
										Description: `Interface port IPv4 mask corresponding to start device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_name": &schema.Schema{
										Description: `Interface port name corresponding to start device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_speed": &schema.Schema{
										Description: `Interface port speed corresponding to start device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tag": &schema.Schema{
										Description: `[Deprecated]
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"target": &schema.Schema{
										Description: `Device ID corresponding to the target device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"acl_applied": &schema.Schema{
										Description: `Indicates if the Access Control List (ACL) is applied on the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"additional_info": &schema.Schema{
										Description: `Additional information about the node
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"connected_device_id": &schema.Schema{
										Description: `ID of the connected device when the nodeType is HOST
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"custom_param": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `[Deprecated] Please refer to nodes.id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"label": &schema.Schema{
													Description: `Label of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"parent_node_id": &schema.Schema{
													Description: `Id of the parent node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"x": &schema.Schema{
													Description: `X coordinate for this node in the topology view
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"y": &schema.Schema{
													Description: `Y coordinate for this node in the topology view
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"data_path_id": &schema.Schema{
										Description: `ID of the path between devices
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_series": &schema.Schema{
										Description: `The series of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_type": &schema.Schema{
										Description: `Type of the device.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"family": &schema.Schema{
										Description: `The product family of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fixed": &schema.Schema{
										Description: `Boolean value indicating whether the position is fixed or will use auto layout
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"grey_out": &schema.Schema{
										Description: `Boolean value indicating whether the node is active for the topology view.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Unique identifier for the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip": &schema.Schema{
										Description: `IP address of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"label": &schema.Schema{
										Description: `Label of the node, typically the hostname of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"network_type": &schema.Schema{
										Description: `Type of the network
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"node_type": &schema.Schema{
										Description: `Type of the node can be 'device' or 'HOST'
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"order": &schema.Schema{
										Description: `Device order by link number
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"os_type": &schema.Schema{
										Description: `OS type of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"platform_id": &schema.Schema{
										Description: `Platform description of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Description: `Role of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role_source": &schema.Schema{
										Description: `Indicates whether the role is assigned manually or automatically
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_version": &schema.Schema{
										Description: `Device OS version
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tags": &schema.Schema{
										Description: `[Deprecated]
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"upper_node": &schema.Schema{
										Description: `ID of the start node
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"user_id": &schema.Schema{
										Description: `ID of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vlan_id": &schema.Schema{
										Description: `VLAN ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"x": &schema.Schema{
										Description: `[Deprecated] Please refer to customParam.x
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"y": &schema.Schema{
										Description: `[Deprecated] Please refer to customerParam.y
`,
										Type:     schema.TypeInt,
										Computed: true,
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

func dataSourceTopologyLayer3Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTopologyType := d.Get("topology_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetL3TopologyDetails")
		vvTopologyType := vTopologyType.(string)

		response1, restyResp1, err := client.Topology.GetL3TopologyDetails(vvTopologyType)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetL3TopologyDetails", err,
				"Failure at GetL3TopologyDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTopologyGetL3TopologyDetailsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetL3TopologyDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTopologyGetL3TopologyDetailsItem(item *dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["links"] = flattenTopologyGetL3TopologyDetailsItemLinks(item.Links)
	respItem["nodes"] = flattenTopologyGetL3TopologyDetailsItemNodes(item.Nodes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTopologyGetL3TopologyDetailsItemLinks(items *[]dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponseLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["additional_info"] = flattenTopologyGetL3TopologyDetailsItemLinksAdditionalInfo(item.AdditionalInfo)
		respItem["end_port_id"] = item.EndPortID
		respItem["end_port_ipv4_address"] = item.EndPortIPv4Address
		respItem["end_port_ipv4_mask"] = item.EndPortIPv4Mask
		respItem["end_port_name"] = item.EndPortName
		respItem["end_port_speed"] = item.EndPortSpeed
		respItem["grey_out"] = boolPtrToString(item.GreyOut)
		respItem["id"] = item.ID
		respItem["link_status"] = item.LinkStatus
		respItem["source"] = item.Source
		respItem["start_port_id"] = item.StartPortID
		respItem["start_port_ipv4_address"] = item.StartPortIPv4Address
		respItem["start_port_ipv4_mask"] = item.StartPortIPv4Mask
		respItem["start_port_name"] = item.StartPortName
		respItem["start_port_speed"] = item.StartPortSpeed
		respItem["tag"] = item.Tag
		respItem["target"] = item.Target
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTopologyGetL3TopologyDetailsItemLinksAdditionalInfo(item *dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponseLinksAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenTopologyGetL3TopologyDetailsItemNodes(items *[]dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponseNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_applied"] = boolPtrToString(item.ACLApplied)
		respItem["additional_info"] = flattenTopologyGetL3TopologyDetailsItemNodesAdditionalInfo(item.AdditionalInfo)
		respItem["custom_param"] = flattenTopologyGetL3TopologyDetailsItemNodesCustomParam(item.CustomParam)
		respItem["connected_device_id"] = item.ConnectedDeviceID
		respItem["data_path_id"] = item.DataPathID
		respItem["device_type"] = item.DeviceType
		respItem["device_series"] = item.DeviceSeries
		respItem["family"] = item.Family
		respItem["fixed"] = boolPtrToString(item.Fixed)
		respItem["grey_out"] = boolPtrToString(item.GreyOut)
		respItem["id"] = item.ID
		respItem["ip"] = item.IP
		respItem["label"] = item.Label
		respItem["network_type"] = item.NetworkType
		respItem["node_type"] = item.NodeType
		respItem["order"] = item.Order
		respItem["os_type"] = item.OsType
		respItem["platform_id"] = item.PlatformID
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["software_version"] = item.SoftwareVersion
		respItem["tags"] = item.Tags
		respItem["upper_node"] = item.UpperNode
		respItem["user_id"] = item.UserID
		respItem["vlan_id"] = item.VLANID
		respItem["x"] = item.X
		respItem["y"] = item.Y
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTopologyGetL3TopologyDetailsItemNodesAdditionalInfo(item *dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponseNodesAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenTopologyGetL3TopologyDetailsItemNodesCustomParam(item *dnacentersdkgo.ResponseTopologyGetL3TopologyDetailsResponseNodesCustomParam) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["label"] = item.Label
	respItem["parent_node_id"] = item.ParentNodeID
	respItem["x"] = item.X
	respItem["y"] = item.Y

	return []map[string]interface{}{
		respItem,
	}

}
