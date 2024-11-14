package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get all users for the Cisco DNA Center System.
`,

		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"auth_source": &schema.Schema{
				Description: `authSource query parameter. The source that authenticates the user. The value of this query parameter can be set to "internal" or "external". If not provided, then all users will be returned in the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"invoke_source": &schema.Schema{
				Description: `invokeSource query parameter. The source that invokes this API. The value of this query parameter must be set to "external".
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"users": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_source": &schema.Schema{
										Description: `Authentiction source, internal or external
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"email": &schema.Schema{
										Description: `Email`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"first_name": &schema.Schema{
										Description: `First Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"last_name": &schema.Schema{
										Description: `Last Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"passphrase_update_time": &schema.Schema{
										Description: `Passphrase Update Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"role_list": &schema.Schema{
										Description: `A list of role ids
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"username": &schema.Schema{
										Description: `Username`,
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

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInvokeSource := d.Get("invoke_source")
	vAuthSource, okAuthSource := d.GetOk("auth_source")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUsersAPI")
		queryParams1 := dnacentersdkgo.GetUsersAPIQueryParams{}

		queryParams1.InvokeSource = vInvokeSource.(string)

		if okAuthSource {
			queryParams1.AuthSource = vAuthSource.(string)
		}

		response1, restyResp1, err := client.UserandRoles.GetUsersAPI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetUsersAPI", err,
				"Failure at GetUsersAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetUsersAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUsersAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetUsersAPIItem(item *dnacentersdkgo.ResponseUserandRolesGetUsersAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["users"] = flattenUserandRolesGetUsersAPIItemUsers(item.Users)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetUsersAPIItemUsers(items *[]dnacentersdkgo.ResponseUserandRolesGetUsersAPIResponseUsers) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["first_name"] = item.FirstName
		respItem["last_name"] = item.LastName
		respItem["auth_source"] = item.AuthSource
		respItem["passphrase_update_time"] = item.PassphraseUpdateTime
		respItem["role_list"] = item.RoleList
		respItem["user_id"] = item.UserID
		respItem["email"] = item.Email
		respItem["username"] = item.Username
		respItems = append(respItems, respItem)
	}
	return respItems
}
