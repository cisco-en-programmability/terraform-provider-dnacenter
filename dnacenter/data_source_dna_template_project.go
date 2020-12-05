package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplateProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateProjectRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_deletable": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"templates": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"composite": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"language": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_params_order": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"last_update_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"latest_version_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"project_associated": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"document_database": &schema.Schema{
										Type:     schema.TypeBool,
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

func dataSourceTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	getProjectsQueryParams := dnac.GetProjectsQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		getProjectsQueryParams.Name = v.(string)
	}

	response, _, err := client.ConfigurationTemplates.GetProjects(&getProjectsQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	projects := flattenTemplateProjectsReadItems(response)
	if err := d.Set("items", projects); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
