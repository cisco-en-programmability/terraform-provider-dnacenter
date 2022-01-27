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

// dataSourceAction
func dataSourceServiceProviderCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to create service provider profile(QOS).
`,

		ReadContext: dataSourceServiceProviderCreateRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
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
	}
}

func dataSourceServiceProviderCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CreateSpProfile")
		request1 := expandRequestServiceProviderCreateCreateSpProfile(ctx, "", d)

		response1, restyResp1, err := client.NetworkSettings.CreateSpProfile(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateSpProfile", err,
				"Failure at CreateSpProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsCreateSpProfileItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CreateSpProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestServiceProviderCreateCreateSpProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfile {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfile{}
	request.Settings = expandRequestServiceProviderCreateCreateSpProfileSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos")))) {
		request.Qos = expandRequestServiceProviderCreateCreateSpProfileSettingsQosArray(ctx, key+".qos", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettingsQosArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
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
		i := expandRequestServiceProviderCreateCreateSpProfileSettingsQos(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestServiceProviderCreateCreateSpProfileSettingsQos(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateSpProfileSettingsQos {
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

func flattenNetworkSettingsCreateSpProfileItem(item *dnacentersdkgo.ResponseNetworkSettingsCreateSpProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
