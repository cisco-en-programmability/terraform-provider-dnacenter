package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTopologyPhysical() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Topology.

- Returns the raw physical topology by specified criteria of nodeType
`,

		ReadContext: dataSourceTopologyPhysicalRead,
		Schema: map[string]*schema.Schema{
			"node_type": &schema.Schema{
				Description: `nodeType query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"links": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_info": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_ipv4_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_ipv4_mask": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_port_speed": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"grey_out": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_ipv4_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_ipv4_mask": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_port_speed": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"tag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"target": &schema.Schema{
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
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"additional_info": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"custom_param": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"label": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"parent_node_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"x": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"y": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"data_path_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"family": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"fixed": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"grey_out": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"label": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"network_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"node_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"order": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"os_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"platform_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"role_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"upper_node": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"user_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"vlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"x": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"y": &schema.Schema{
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

func dataSourceTopologyPhysicalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNodeType, okNodeType := d.GetOk("node_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPhysicalTopology")
		queryParams1 := dnacentersdkgo.GetPhysicalTopologyQueryParams{}

		if okNodeType {
			queryParams1.NodeType = vNodeType.(string)
		}

		response1, restyResp1, err := client.Topology.GetPhysicalTopology(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPhysicalTopology", err,
				"Failure at GetPhysicalTopology, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTopologyGetPhysicalTopologyItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPhysicalTopology response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTopologyGetPhysicalTopologyItem(item *dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["links"] = flattenTopologyGetPhysicalTopologyItemLinks(item.Links)
	respItem["nodes"] = flattenTopologyGetPhysicalTopologyItemNodes(item.Nodes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTopologyGetPhysicalTopologyItemLinks(items *[]dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponseLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["additional_info"] = flattenTopologyGetPhysicalTopologyItemLinksAdditionalInfo(item.AdditionalInfo)
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

func flattenTopologyGetPhysicalTopologyItemLinksAdditionalInfo(item *dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponseLinksAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenTopologyGetPhysicalTopologyItemNodes(items *[]dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponseNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_applied"] = boolPtrToString(item.ACLApplied)
		respItem["additional_info"] = flattenTopologyGetPhysicalTopologyItemNodesAdditionalInfo(item.AdditionalInfo)
		respItem["custom_param"] = flattenTopologyGetPhysicalTopologyItemNodesCustomParam(item.CustomParam)
		respItem["data_path_id"] = item.DataPathID
		respItem["device_type"] = item.DeviceType
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

func flattenTopologyGetPhysicalTopologyItemNodesAdditionalInfo(item *dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponseNodesAdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenTopologyGetPhysicalTopologyItemNodesCustomParam(item *dnacentersdkgo.ResponseTopologyGetPhysicalTopologyResponseNodesCustomParam) []map[string]interface{} {
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
