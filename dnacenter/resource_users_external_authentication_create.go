package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceUsersExternalAuthenticationCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on User and Roles.

- Enable or disable external authentication on Cisco DNA Center System.
Please find the Administrator Guide for your particular release from the list linked below and follow the steps required
to enable external authentication before trying to do so from this API.
https://www.cisco.com/c/en/us/support/cloud-systems-management/dna-center/products-maintenance-guides-list.html
`,

		CreateContext: resourceUsersExternalAuthenticationCreateCreate,
		ReadContext:   resourceUsersExternalAuthenticationCreateRead,
		DeleteContext: resourceUsersExternalAuthenticationCreateDelete,
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

						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": &schema.Schema{
							Description: `Enable/disable External Authentication.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceUsersExternalAuthenticationCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestUsersExternalAuthenticationCreateManageExternalAuthenticationSettingAPI(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.UserandRoles.ManageExternalAuthenticationSettingAPI(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ManageExternalAuthenticationSettingAPI", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenUserandRolesManageExternalAuthenticationSettingAPIItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ManageExternalAuthenticationSettingAPI response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceUsersExternalAuthenticationCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceUsersExternalAuthenticationCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestUsersExternalAuthenticationCreateManageExternalAuthenticationSettingAPI(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestUserandRolesManageExternalAuthenticationSettingAPI {
	request := dnacentersdkgo.RequestUserandRolesManageExternalAuthenticationSettingAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable")))) {
		request.Enable = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenUserandRolesManageExternalAuthenticationSettingAPIItem(item *dnacentersdkgo.ResponseUserandRolesManageExternalAuthenticationSettingAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
