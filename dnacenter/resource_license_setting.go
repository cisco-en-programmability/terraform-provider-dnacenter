package dnacenter

import (
	"context"
	"errors"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Licenses.

- Update license setting Configure default smart account id  and/or virtual account id for auto registration of devices
for smart license flow. Virtual account should be part of default smart account. Default smart account id cannot be set
to 'null'. Auto registration of devices for smart license flow is applicable only for direct or on-prem SSM connection
mode.
`,

		CreateContext: resourceLicenseSettingCreate,
		ReadContext:   resourceLicenseSettingRead,
		UpdateContext: resourceLicenseSettingUpdate,
		DeleteContext: resourceLicenseSettingDelete,
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

						"auto_registration_virtual_account_id": &schema.Schema{
							Description: `Virtual account id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_smart_account_id": &schema.Schema{
							Description: `Default smart account id
`,
							Type:     schema.TypeString,
							Computed: true,
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

						"auto_registration_virtual_account_id": &schema.Schema{
							Description: `Virtual account id
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_smart_account_id": &schema.Schema{
							Description: `Default smart account id
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceLicenseSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// TODO: Add the path params to `item` schema
	//       & return it individually

	return diags
}

func resourceLicenseSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveLicenseSetting")

		response1, restyResp1, err := client.Licenses.RetrieveLicenseSetting()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesRetrieveLicenseSettingItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveLicenseSetting response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceLicenseSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestLicenseSettingUpdateLicenseSetting(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Licenses.UpdateLicenseSetting(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateLicenseSetting", err, restyResp1.String(),
					"Failure at UpdateLicenseSetting, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateLicenseSetting", err,
				"Failure at UpdateLicenseSetting, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceLicenseSettingRead(ctx, d, m)
}

func resourceLicenseSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing LicenseSettingDelete", err, "Delete method is not supported",
		"Failure at LicenseSettingDelete, unexpected response", ""))

	return diags
}
func expandRequestLicenseSettingUpdateLicenseSetting(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLicensesUpdateLicenseSetting {
	request := dnacentersdkgo.RequestLicensesUpdateLicenseSetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_smart_account_id")))) {
		request.DefaultSmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_registration_virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_registration_virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_registration_virtual_account_id")))) {
		request.AutoRegistrationVirtualAccountID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
