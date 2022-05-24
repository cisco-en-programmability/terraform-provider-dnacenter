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

func resourceServiceProvider() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to create service provider profile(QOS).

- API to update SP profile
`,

		CreateContext: resourceServiceProviderCreate,
		ReadContext:   resourceServiceProviderRead,
		UpdateContext: resourceServiceProviderUpdate,
		DeleteContext: resourceServiceProviderDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"qos": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"model": &schema.Schema{
													Description: `Model`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"old_profile_name": &schema.Schema{
													Description: `Old Profile Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"profile_name": &schema.Schema{
													Description: `Profile Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"wan_provider": &schema.Schema{
													Description: `Wan Provider`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
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

func resourceServiceProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestServiceProviderCreateSpProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.NetworkSettings.CreateSpProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSpProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSpProfile", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceServiceProviderRead(ctx, d, m)
}

func resourceServiceProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetServiceProviderDetails")

		response1, restyResp1, err := client.NetworkSettings.GetServiceProviderDetails()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetServiceProviderDetails", err,
				"Failure at GetServiceProviderDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenNetworkSettingsGetServiceProviderDetailsItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetServiceProviderDetails search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceServiceProviderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestServiceProviderUpdateSpProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateSpProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSpProfile", err, restyResp1.String(),
					"Failure at UpdateSpProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSpProfile", err,
				"Failure at UpdateSpProfile, unexpected response", ""))
			return diags
		}
	}

	return resourceServiceProviderRead(ctx, d, m)
}

func resourceServiceProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete ServiceProvider on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestServiceProviderCreateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfile{}
	request.Settings = expandRequestServiceProviderCreateSpProfileSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderCreateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestServiceProviderCreateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfile{}
	request.Settings = expandRequestServiceProviderUpdateSpProfileSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderUpdateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := []dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestServiceProviderUpdateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderUpdateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateSpProfileSettingsQos{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model")))) {
		request.Model = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wan_provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wan_provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wan_provider")))) {
		request.WanProvider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".old_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".old_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".old_profile_name")))) {
		request.OldProfileName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchNetworkSettingsGetServiceProviderDetails(m interface{}, queryParams dnacentersdkgo.GetServiceProviderDetailsQueryParams) (*dnacentersdkgo.ResponseItemNetworkSettingsGetServiceProviderDetails, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemNetworkSettingsGetServiceProviderDetails
	var ite *dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetails
	ite, _, err = client.NetworkSettings.GetServiceProviderDetails(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemNetworkSettingsGetServiceProviderDetails
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
