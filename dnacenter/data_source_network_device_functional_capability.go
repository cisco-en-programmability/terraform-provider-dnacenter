package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceFunctionalCapability() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the functional-capability for given devices

- Returns functional capability with given Id
`,

		ReadContext: dataSourceNetworkDeviceFunctionalCapabilityRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"function_name": &schema.Schema{
				Description: `functionName query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": &schema.Schema{
				Description: `id path parameter. Functional Capability UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"function_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_info": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"property_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"string_value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"function_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"function_op_state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"functional_capability": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_info": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"function_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_info": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"property_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"string_value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"function_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"function_op_state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceFunctionalCapabilityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vFunctionName, okFunctionName := d.GetOk("function_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okDeviceID, okFunctionName}
	log.Printf("[DEBUG] Selecting method. Method 1 %q", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %q", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetFunctionalCapabilityForDevices")
		queryParams1 := dnacentersdkgo.GetFunctionalCapabilityForDevicesQueryParams{}

		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okFunctionName {
			queryParams1.FunctionName = interfaceToSliceString(vFunctionName)
		}

		response1, restyResp1, err := client.Devices.GetFunctionalCapabilityForDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetFunctionalCapabilityForDevices", err,
				"Failure at GetFunctionalCapabilityForDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetFunctionalCapabilityForDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFunctionalCapabilityForDevices response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method 2: GetFunctionalCapabilityByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Devices.GetFunctionalCapabilityByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetFunctionalCapabilityByID", err,
				"Failure at GetFunctionalCapabilityByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetFunctionalCapabilityByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFunctionalCapabilityByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetFunctionalCapabilityForDevicesItems(items *[]dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesItemsAttributeInfo(item.AttributeInfo)
		respItem["device_id"] = item.DeviceID
		respItem["functional_capability"] = flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapability(item.FunctionalCapability)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesItemsAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapability(items *[]dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityAttributeInfo(item.AttributeInfo)
		respItem["function_details"] = flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityFunctionDetails(item.FunctionDetails)
		respItem["function_name"] = item.FunctionName
		respItem["function_op_state"] = item.FunctionOpState
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityFunctionDetails(items *[]dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityFunctionDetailsAttributeInfo(item.AttributeInfo)
		respItem["id"] = item.ID
		respItem["property_name"] = item.PropertyName
		respItem["string_value"] = item.StringValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesItemsFunctionalCapabilityFunctionDetailsAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityByIDItem(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityByIDItemAttributeInfo(item.AttributeInfo)
	respItem["function_details"] = flattenDevicesGetFunctionalCapabilityByIDItemFunctionDetails(item.FunctionDetails)
	respItem["function_name"] = item.FunctionName
	respItem["function_op_state"] = item.FunctionOpState
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetFunctionalCapabilityByIDItemAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityByIDItemFunctionDetails(items *[]dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityByIDItemFunctionDetailsAttributeInfo(item.AttributeInfo)
		respItem["id"] = item.ID
		respItem["property_name"] = item.PropertyName
		respItem["string_value"] = item.StringValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityByIDItemFunctionDetailsAttributeInfo(item *dnacentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
