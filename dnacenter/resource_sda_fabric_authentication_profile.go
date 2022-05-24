package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Add default authentication template in SDA Fabric

- Update default authentication profile in SDA Fabric

- Delete default authentication profile in SDA Fabric
`,

		CreateContext: resourceSdaFabricAuthenticationProfileCreate,
		ReadContext:   resourceSdaFabricAuthenticationProfileRead,
		UpdateContext: resourceSdaFabricAuthenticationProfileUpdate,
		DeleteContext: resourceSdaFabricAuthenticationProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddDefaultAuthenticationTemplateInSDAFabric`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_order": &schema.Schema{
							Description: `Authentication Order
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `Dot1x To MabFallback Timeout( Allowed range is [3-120])
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"number_of_hosts": &schema.Schema{
							Description: `Number Of Hosts
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"wake_on_lan": &schema.Schema{
							Description: `Wake On Lan
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricAuthenticationProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabric(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Sda.AddDefaultAuthenticationTemplateInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDefaultAuthenticationTemplateInSdaFabric", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSdaFabricAuthenticationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDefaultAuthenticationProfileFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams{}

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		if okAuthenticateTemplateName {
			queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
		}

		response1, restyResp1, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDefaultAuthenticationProfileFromSdaFabric", err,
				"Failure at GetDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetDefaultAuthenticationProfileFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDefaultAuthenticationProfileFromSdaFabric response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaFabricAuthenticationProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
	item, err := searchSdaGetDefaultAuthenticationProfileFromSDAFabric(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDefaultAuthenticationProfileFromSDAFabric", err,
			"Failure at GetDefaultAuthenticationProfileFromSDAFabric, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabric(ctx, "parameters", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.UpdateDefaultAuthenticationProfileInSdaFabric(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDefaultAuthenticationProfileInSdaFabric", err, restyResp1.String(),
					"Failure at UpdateDefaultAuthenticationProfileInSdaFabric, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDefaultAuthenticationProfileInSdaFabric", err,
				"Failure at UpdateDefaultAuthenticationProfileInSdaFabric, unexpected response", ""))
			return diags
		}
	}

	return resourceSdaFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSdaFabricAuthenticationProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	vAuthenticateTemplateName := resourceMap["authenticate_template_name"]

	queryParams1 := dnacentersdkgo.GetDefaultAuthenticationProfileFromSdaFabricQueryParams
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	queryParams1.AuthenticateTemplateName = vAuthenticateTemplateName
	item, err := searchSdaGetDefaultAuthenticationProfileFromSDAFabric(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDefaultAuthenticationProfileFromSDAFabric", err,
			"Failure at GetDefaultAuthenticationProfileFromSDAFabric, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Sda.GetDefaultAuthenticationProfileFromSdaFabric(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsSdaGetDefaultAuthenticationProfileFromSdaFabric(m, getResp1, nil)
		item1, err := searchSdaGetDefaultAuthenticationProfileFromSdaFabric(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.Sda.DeleteDefaultAuthenticationProfileFromSdaFabric()
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDefaultAuthenticationProfileFromSdaFabric", err, restyResp1.String(),
				"Failure at DeleteDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDefaultAuthenticationProfileFromSdaFabric", err,
			"Failure at DeleteDefaultAuthenticationProfileFromSdaFabric, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddDefaultAuthenticationTemplateInSdaFabric{}
	if v := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := []dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileAddDefaultAuthenticationTemplateInSdaFabricItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric {
	request := dnacentersdkgo.RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := dnacentersdkgo.RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
	if v := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := []dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricAuthenticationProfileUpdateDefaultAuthenticationProfileInSdaFabricItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric {
	request := dnacentersdkgo.RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_order")))) {
		request.AuthenticationOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot1x_to_mab_fallback_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) {
		request.Dot1XToMabFallbackTimeout = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wake_on_lan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wake_on_lan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wake_on_lan")))) {
		request.WakeOnLan = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_hosts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_hosts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_hosts")))) {
		request.NumberOfHosts = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
