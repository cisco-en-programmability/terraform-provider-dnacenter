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
func dataSourceTemplatePreview() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Configuration Templates.

- API to preview a template.
`,

		ReadContext: dataSourceTemplatePreviewRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `UUID of device to get template preview
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_preview": &schema.Schema{
							Description: `Generated template preview
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_id": &schema.Schema{
							Description: `UUID of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_errors": &schema.Schema{
							Description: `Validation error in template content if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"params": &schema.Schema{
				Description: `Params to render preview
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_params": &schema.Schema{
				Description: `Resource params to render preview
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"template_id": &schema.Schema{
				Description: `UUID of template to get template preview
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceTemplatePreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: PreviewTemplate")
		request1 := expandRequestTemplatePreviewPreviewTemplate(ctx, "", d)

		response1, restyResp1, err := client.ConfigurationTemplates.PreviewTemplate(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing PreviewTemplate", err,
				"Failure at PreviewTemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesPreviewTemplateItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting PreviewTemplate response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestTemplatePreviewPreviewTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplate {
	request := dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestTemplatePreviewPreviewTemplateParams(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = expandRequestTemplatePreviewPreviewTemplateResourceParamsArray(ctx, key+".resource_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTemplatePreviewPreviewTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTemplatePreviewPreviewTemplateResourceParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams {
	request := []dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams{}
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
		i := expandRequestTemplatePreviewPreviewTemplateResourceParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestTemplatePreviewPreviewTemplateResourceParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenConfigurationTemplatesPreviewTemplateItem(item *dnacentersdkgo.ResponseConfigurationTemplatesPreviewTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli_preview"] = item.CliPreview
	respItem["device_id"] = item.DeviceID
	respItem["template_id"] = item.TemplateID
	respItem["validation_errors"] = flattenConfigurationTemplatesPreviewTemplateItemValidationErrors(item.ValidationErrors)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesPreviewTemplateItemValidationErrors(item *dnacentersdkgo.ResponseConfigurationTemplatesPreviewTemplateValidationErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
