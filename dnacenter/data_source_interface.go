package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

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

						"interface_uuid": &schema.Schema{
							Description: `Interface Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"operations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"applicable": &schema.Schema{
										Description: `Applicable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Failure Reason`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
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

									"applicable": &schema.Schema{
										Description: `Applicable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Failure Reason`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
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
	respItem["interface_uuid"] = item.InterfaceUUID
	respItem["properties"] = flattenDevicesLegitOperationsForInterfaceItemProperties(item.Properties)
	respItem["operations"] = flattenDevicesLegitOperationsForInterfaceItemOperations(item.Operations)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesLegitOperationsForInterfaceItemProperties(items *[]dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponseProperties) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["applicable"] = boolPtrToString(item.Applicable)
		respItem["failure_reason"] = flattenDevicesLegitOperationsForInterfaceItemPropertiesFailureReason(item.FailureReason)
		respItems = append(respItems, respItem)
	}

	return respItems

}

func flattenDevicesLegitOperationsForInterfaceItemPropertiesFailureReason(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponsePropertiesFailureReason) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesLegitOperationsForInterfaceItemOperations(items *[]dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponseOperations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["applicable"] = boolPtrToString(item.Applicable)
		respItem["failure_reason"] = flattenDevicesLegitOperationsForInterfaceItemOperationsFailureReason(item.FailureReason)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesLegitOperationsForInterfaceItemOperationsFailureReason(item *dnacentersdkgo.ResponseDevicesLegitOperationsForInterfaceResponseOperationsFailureReason) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
