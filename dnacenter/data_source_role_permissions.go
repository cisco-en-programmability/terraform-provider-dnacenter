package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRolePermissions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get permissions for a role from Cisco DNA Center System
`,

		ReadContext: dataSourceRolePermissionsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"resource_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_permission": &schema.Schema{
										Description: `Default permission
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
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
		},
	}
}

func dataSourceRolePermissionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPermissionsApI")

		response1, restyResp1, err := client.UserandRoles.GetPermissionsApI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPermissionsApI", err,
				"Failure at GetPermissionsApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetPermissionsApIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPermissionsApI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetPermissionsApIItem(item *dnacentersdkgo.ResponseUserandRolesGetPermissionsAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["resource_types"] = flattenUserandRolesGetPermissionsApIItemResourceTypes(item.ResourceTypes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetPermissionsApIItemResourceTypes(items *[]dnacentersdkgo.ResponseUserandRolesGetPermissionsAPIResponseResourceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["display_name"] = item.DisplayName
		respItem["description"] = item.Description
		respItem["default_permission"] = item.DefaultPermission
		respItems = append(respItems, respItem)
	}
	return respItems
}
