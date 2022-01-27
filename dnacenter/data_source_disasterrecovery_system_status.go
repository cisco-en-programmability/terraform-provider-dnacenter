package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDisasterrecoverySystemStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Disaster Recovery.

- Detailed and Summarized status of DR components (Active, Standby and Witness system's health).
`,

		ReadContext: dataSourceDisasterrecoverySystemStatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ipconfig": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface": &schema.Schema{
										Description: `Enterprise or Management interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip": &schema.Schema{
										Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.  
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vip": &schema.Schema{
										Description: `Is this interface an Virtual IP address or not. This is true for Global DR VIP
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"side_a": &schema.Schema{
										Description: `A Side of the IPSec Tunnel
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"side_b": &schema.Schema{
										Description: `Other Side of the IPSec Tunnel 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Status of this IPSec Tunnel
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"main": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is true for cluster level.
`,
													// Type:        schema.TypeBool,
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

												"hostname": &schema.Schema{
													Description: `Hostname of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipaddresses": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `Node IP address
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface a Virtual IP address or not. This is false for node level.
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Main Site. 
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"recovery": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is true for cluster level.
`,
													// Type:        schema.TypeBool,
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

												"hostname": &schema.Schema{
													Description: `Hostname of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipconfig": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `Node IP Address
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface a Virtual IP Address or not. This is false for node level.
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Recovery site
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"site": &schema.Schema{
							Description: `Site of the disaster recovery system. 
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"state": &schema.Schema{
							Description: `State of the Disaster Recovery System.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"witness": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `In case of witness, this is only an IP. 
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is false for witness.
`,
													// Type:        schema.TypeBool,
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

												"hostname": &schema.Schema{
													Description: `Hostname of the witness node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipconfig": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `In case of witness, this is only an IP
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface an Virtual IP address or not. This is false for Witness 
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Witness Site
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
	}
}

func dataSourceDisasterrecoverySystemStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DisasterRecoveryStatus")

		response1, restyResp1, err := client.DisasterRecovery.DisasterRecoveryStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DisasterRecoveryStatus", err,
				"Failure at DisasterRecoveryStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDisasterRecoveryDisasterRecoveryStatusItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DisasterRecoveryStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDisasterRecoveryDisasterRecoveryStatusItem(item *dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemIPconfig(item.IPconfig)
	respItem["site"] = item.Site
	respItem["main"] = flattenDisasterRecoveryDisasterRecoveryStatusItemMain(item.Main)
	respItem["recovery"] = flattenDisasterRecoveryDisasterRecoveryStatusItemRecovery(item.Recovery)
	respItem["witness"] = flattenDisasterRecoveryDisasterRecoveryStatusItemWitness(item.Witness)
	respItem["state"] = item.State
	respItem["ipsec_tunnel"] = flattenDisasterRecoveryDisasterRecoveryStatusItemIPsecTunnel(item.IPsecTunnel)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemMain(item *dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusMain) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemMainIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusItemMainNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusItemMainIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusMainIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemMainNodes(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusMainNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipaddresses"] = flattenDisasterRecoveryDisasterRecoveryStatusItemMainNodesIPaddresses(item.IPaddresses)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemMainNodesIPaddresses(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusMainNodesIPaddresses) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemRecovery(item *dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusRecovery) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryNodes(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryNodesIPconfig(item.IPconfig)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemRecoveryNodesIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodesIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemWitness(item *dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusWitness) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusWitnessIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessNodes(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessNodesIPconfig(item.IPconfig)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemWitnessNodesIPconfig(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodesIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusItemIPsecTunnel(items *[]dnacentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusIPsecTunnel) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["side_a"] = item.SideA
		respItem["side_b"] = item.SideB
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
