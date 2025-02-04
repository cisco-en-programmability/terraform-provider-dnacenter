package dnacenter

import (
	"context"
	"errors"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on User and Roles.

- Add a new user for Cisco DNA Center System.

- Update a user for Cisco DNA Center System.
`,

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"email": &schema.Schema{
							Description: `Email`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"first_name": &schema.Schema{
							Description: `First Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"last_name": &schema.Schema{
							Description: `Last Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Computed:    true,
						},
						"role_list": &schema.Schema{
							Description: `Role id list
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_id": &schema.Schema{
							Description: `User Id`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"username": &schema.Schema{
							Description: `Username`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestUserAddUserAPI(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vInvokeSource := resourceItem["invoke_source"]
	vvInvokeSource := interfaceToString(vInvokeSource)
	queryParamImport := dnacentersdkgo.GetUsersAPIQueryParams{}
	queryParamImport.InvokeSource = vvInvokeSource
	item2, _, err := client.UserandRoles.GetUsersAPI(&queryParamImport)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["invoke_source"] = vvInvokeSource
		d.SetId(joinResourceID(resourceMap))
		return resourceUserRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.UserandRoles.AddUserAPI(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddUserAPI", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddUserAPI", err))
		return diags
	}
	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.GetUsersAPIQueryParams{}
	queryParamValidate.InvokeSource = vvInvokeSource
	item3, _, err := client.UserandRoles.GetUsersAPI(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddUserAPI", err,
			"Failure at AddUserAPI, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["invoke_source"] = vvInvokeSource

	d.SetId(joinResourceID(resourceMap))
	return resourceUserRead(ctx, d, m)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vInvokeSource := resourceMap["invoke_source"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUsersAPI")
		queryParams1 := dnacentersdkgo.GetUsersAPIQueryParams{}

		queryParams1.InvokeSource = vInvokeSource

		response1, restyResp1, err := client.UserandRoles.GetUsersAPI(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
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

		return diags

	}
	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestUserUpdateUserAPI(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.UserandRoles.UpdateUserAPI(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateUserAPI", err, restyResp1.String(),
					"Failure at UpdateUserAPI, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateUserAPI", err,
				"Failure at UpdateUserAPI, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing User", err, "Delete method is not supported",
		"Failure at UserDelete, unexpected response", ""))
	return diags
}
func expandRequestUserAddUserAPI(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestUserandRolesAddUserAPI {
	request := dnacentersdkgo.RequestUserandRolesAddUserAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role_list")))) {
		request.RoleList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestUserUpdateUserAPI(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestUserandRolesUpdateUserAPI {
	request := dnacentersdkgo.RequestUserandRolesUpdateUserAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".first_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".first_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".first_name")))) {
		request.FirstName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_name")))) {
		request.LastName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email")))) {
		request.Email = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_id")))) {
		request.UserID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role_list")))) {
		request.RoleList = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
