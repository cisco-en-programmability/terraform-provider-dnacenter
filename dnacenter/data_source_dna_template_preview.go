package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatePreview() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplatePreviewRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"params": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cli_preview": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_errors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceTemplatePreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	templateID := d.Get("template_id").(string)
	previewTemplateRequest := dnac.PreviewTemplateRequest{TemplateID: templateID}
	if v, ok := d.GetOk("params"); ok {
		previewTemplateRequest.Params = v.(map[string]interface{})
	}

	response, _, err := client.ConfigurationTemplates.PreviewTemplate(&previewTemplateRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templatePreview := flattenPreviewTemplateReadItem(response)
	if err := d.Set("item", templatePreview); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
