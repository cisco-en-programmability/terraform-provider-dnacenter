package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsersExternalServersAAAAttribute() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get the current value of the custom AAA attribute.
`,

		ReadContext: dataSourceUsersExternalServersAAAAttributeRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_name": &schema.Schema{
										Description: `Value of the custom AAA attribute name
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

func dataSourceUsersExternalServersAAAAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAAAAttributeAPI")

		response1, restyResp1, err := client.UserandRoles.GetAAAAttributeAPI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAAAAttributeAPI", err,
				"Failure at GetAAAAttributeAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetAAAAttributeAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAAAAttributeAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetAAAAttributeAPIItem(item *dnacentersdkgo.ResponseUserandRolesGetAAAAttributeAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_attributes"] = flattenUserandRolesGetAAAAttributeAPIItemAAAAttributes(item.AAAAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetAAAAttributeAPIItemAAAAttributes(items *[]dnacentersdkgo.ResponseUserandRolesGetAAAAttributeAPIResponseAAAAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_name"] = item.AttributeName
		respItems = append(respItems, respItem)
	}
	return respItems
}
