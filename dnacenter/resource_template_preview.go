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
func resourceTemplatePreview() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Configuration Templates.

- API to preview a template.
`,

		CreateContext: resourceTemplatePreviewCreate,
		ReadContext:   resourceTemplatePreviewRead,
		DeleteContext: resourceTemplatePreviewDelete,
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
							Type:     schema.TypeString, //TEST,
							Computed: true,
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
						"device_id": &schema.Schema{
							Description: `UUID of device to get template preview
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"params": &schema.Schema{
							Description: `Params to render preview
`,
							Type:     schema.TypeString, //TEST,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"resource_params": &schema.Schema{
							Description: `Resource params to render preview
`,
							Type:     schema.TypeString, //TEST,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template to get template preview
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceTemplatePreviewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestTemplatePreviewPreviewTemplate(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ConfigurationTemplates.PreviewTemplate(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing PreviewTemplate", err))
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
func resourceTemplatePreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceTemplatePreviewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
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
		request.ResourceParams = expandRequestTemplatePreviewPreviewTemplateResourceParams(ctx, key+".resource_params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestTemplatePreviewPreviewTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateParams
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestTemplatePreviewPreviewTemplateResourceParams(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams {
	var request dnacentersdkgo.RequestConfigurationTemplatesPreviewTemplateResourceParams
	request = d.Get(fixKeyAccess(key))
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
