package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplateVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateVersionRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"composite": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"versions_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"author": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_comment": &schema.Schema{
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

func dataSourceTemplateVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	templateID := d.Get("template_id").(string)
	response, _, err := client.ConfigurationTemplates.GetTemplateVersions(templateID)
	if err != nil {
		return diag.FromErr(err)
	}

	templates := flattenTemplateVersionsReadItem(response)
	if err := d.Set("item", templates); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
