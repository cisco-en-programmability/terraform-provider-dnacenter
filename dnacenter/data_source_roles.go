package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRoles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get all roles for the Cisco DNA Center system
`,

		ReadContext: dataSourceRolesRead,
		Schema: map[string]*schema.Schema{
			"invoke_source": &schema.Schema{
				Description: `invokeSource header parameter. The source that invoke this API
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"roles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"meta": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"created": &schema.Schema{
													Description: `The timestamp that the resource type was created
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"created_by": &schema.Schema{
													Description: `The user that creates the resource type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"last_modified": &schema.Schema{
													Description: `The latestest timestamp that the resource type was updated
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"name": &schema.Schema{
										Description: `Role name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"resource_types": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"operations": &schema.Schema{
													Description: `Operations`,
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

									"role_id": &schema.Schema{
										Description: `Role Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Role type, possible values are: "DEFAULT", "SYSTEM", "CUSTOM"
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

func dataSourceRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInvokeSource := d.Get("invoke_source")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRolesApI")

		headerParams1 := dnacentersdkgo.GetRolesApIHeaderParams{}

		headerParams1.InvokeSource = vInvokeSource.(string)

		response1, restyResp1, err := client.UserandRoles.GetRolesApI(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRolesApI", err,
				"Failure at GetRolesApI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetRolesApIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRolesApI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetRolesApIItem(item *dnacentersdkgo.ResponseUserandRolesGetRolesAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["roles"] = flattenUserandRolesGetRolesApIItemRoles(item.Roles)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetRolesApIItemRoles(items *[]dnacentersdkgo.ResponseUserandRolesGetRolesAPIResponseRoles) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["resource_types"] = flattenUserandRolesGetRolesApIItemRolesResourceTypes(item.ResourceTypes)
		respItem["meta"] = flattenUserandRolesGetRolesApIItemRolesMeta(item.Meta)
		respItem["role_id"] = item.RoleID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenUserandRolesGetRolesApIItemRolesResourceTypes(items *[]dnacentersdkgo.ResponseUserandRolesGetRolesAPIResponseRolesResourceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["operations"] = item.Operations
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenUserandRolesGetRolesApIItemRolesMeta(item *dnacentersdkgo.ResponseUserandRolesGetRolesAPIResponseRolesMeta) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["created_by"] = item.CreatedBy
	respItem["created"] = item.Created
	respItem["last_modified"] = item.LastModified

	return []map[string]interface{}{
		respItem,
	}

}
