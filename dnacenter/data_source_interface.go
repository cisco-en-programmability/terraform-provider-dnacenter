package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get list of all properties & operations valid for an interface.
`,

		ReadContext: dataSourceInterfaceRead,
		Schema: map[string]*schema.Schema{
			"interface_uuid": &schema.Schema{
				Description: `interfaceUuid path parameter. Interface ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_uuid": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"operations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"properties": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"applicable": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},

																		"failure_reason": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},

																		"name": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"required": &schema.Schema{
																Description: `Required`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"type": &schema.Schema{
																Description: `Type`,
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
											},
										},
									},

									"properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"properties": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"applicable": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},

																		"failure_reason": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},

																		"name": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"type": &schema.Schema{
																						Description: `Type`,
																						Type:        schema.TypeString,
																						Computed:    true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"required": &schema.Schema{
																Description: `Required`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"type": &schema.Schema{
																Description: `Type`,
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
											},
										},
									},
								},
							},
						},

						"required": &schema.Schema{
							Description: `Required`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInterfaceUUID := d.Get("interface_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LegitOperationsForInterface")
		vvInterfaceUUID := vInterfaceUUID.(string)

		response1, restyResp1, err := client.Devices.LegitOperationsForInterface(vvInterfaceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LegitOperationsForInterface", err,
				"Failure at LegitOperationsForInterface, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesLegitOperationsForInterfaceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LegitOperationsForInterface response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesLegitOperationsForInterfaceItem(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["properties"] = flattenDevicesLegitOperationsForInterfaceItemProperties(item.Properties)
	respItem["required"] = item.Required
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesLegitOperationsForInterfaceItemProperties(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interface_uuid"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesInterfaceUUID(item.InterfaceUUID)
	respItem["properties"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesProperties(item.Properties)
	respItem["operations"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperations(item.Operations)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesInterfaceUUID(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesInterfaceUUID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesProperties(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["items"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItems(item.Items)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItems(items *[]dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesPropertiesItems) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["properties"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsProperties(item.Properties)
		respItem["required"] = item.Required
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsProperties(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesPropertiesItemsProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesName(item.Name)
	respItem["applicable"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesApplicable(item.Applicable)
	respItem["failure_reason"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesFailureReason(item.FailureReason)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesName(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesPropertiesItemsPropertiesName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesApplicable(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesPropertiesItemsPropertiesApplicable) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesPropertiesItemsPropertiesFailureReason(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesPropertiesItemsPropertiesFailureReason) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperations(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperations) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["items"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItems(item.Items)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItems(items *[]dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperationsItems) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["properties"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsProperties(item.Properties)
		respItem["required"] = item.Required
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsProperties(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperationsItemsProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesName(item.Name)
	respItem["applicable"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesApplicable(item.Applicable)
	respItem["failure_reason"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesFailureReason(item.FailureReason)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesName(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperationsItemsPropertiesName) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesApplicable(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperationsItemsPropertiesApplicable) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesOperationsItemsPropertiesFailureReason(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesOperationsItemsPropertiesFailureReason) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
