package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceCustomPrompt() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on System Settings.

- Save custom prompt added by user in Cisco DNA Center. API will always override the existing prompts. User should
provide all the custom prompt in case of any update
`,

		CreateContext: resourceNetworkDeviceCustomPromptCreate,
		ReadContext:   resourceNetworkDeviceCustomPromptRead,
		UpdateContext: resourceNetworkDeviceCustomPromptUpdate,
		DeleteContext: resourceNetworkDeviceCustomPromptDelete,
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

						"custom_password_prompt": &schema.Schema{
							Description: `Custom Password`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"custom_username_prompt": &schema.Schema{
							Description: `Custom Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_password_prompt": &schema.Schema{
							Description: `Default Password`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_username_prompt": &schema.Schema{
							Description: `Default Username`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"password_prompt": &schema.Schema{
							Description: `Password Prompt`,
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
						},
						"username_prompt": &schema.Schema{
							Description: `Username Prompt`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceCustomPromptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkDeviceCustomPromptCustomPromptPostAPI(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.SystemSettings.CustomPromptPostAPI(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CustomPromptPostAPI", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CustomPromptPostAPI", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkDeviceCustomPromptRead(ctx, d, m)
}

func resourceNetworkDeviceCustomPromptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CustomPromptSupportGetAPI")

		response1, restyResp1, err := client.SystemSettings.CustomPromptSupportGetAPI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CustomPromptSupportGetAPI", err,
				"Failure at CustomPromptSupportGetAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsCustomPromptSupportGetAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CustomPromptSupportGetAPI response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkDeviceCustomPromptUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDeviceCustomPromptRead(ctx, d, m)
}

func resourceNetworkDeviceCustomPromptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete NetworkDeviceCustomPrompt on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNetworkDeviceCustomPromptCustomPromptPostAPI(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSystemSettingsCustomPromptPostAPI {
	request := dnacentersdkgo.RequestSystemSettingsCustomPromptPostAPI{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username_prompt")))) {
		request.UsernamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
