package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateRead,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_series": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_conflicting_templates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"items": &schema.Schema{
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

func dataSourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	getsTheTemplatesAvailableQueryParams := dnac.GetsTheTemplatesAvailableQueryParams{}
	if v, ok := d.GetOk("project_id"); ok {
		getsTheTemplatesAvailableQueryParams.ProjectID = v.(string)
	}
	if v, ok := d.GetOk("software_type"); ok {
		getsTheTemplatesAvailableQueryParams.SoftwareType = v.(string)
	}
	if v, ok := d.GetOk("software_version"); ok {
		getsTheTemplatesAvailableQueryParams.SoftwareVersion = v.(string)
	}
	if v, ok := d.GetOk("product_family"); ok {
		getsTheTemplatesAvailableQueryParams.ProductFamily = v.(string)
	}
	if v, ok := d.GetOk("product_series"); ok {
		getsTheTemplatesAvailableQueryParams.ProductSeries = v.(string)
	}
	if v, ok := d.GetOk("product_type"); ok {
		getsTheTemplatesAvailableQueryParams.ProductType = v.(string)
	}
	if v, ok := d.GetOk("product_type"); ok {
		getsTheTemplatesAvailableQueryParams.FilterConflictingTemplates = v.(bool)
	}

	response, _, err := client.ConfigurationTemplates.GetsTheTemplatesAvailable(&getsTheTemplatesAvailableQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	templates := flattenTemplatesAvailableReadItems(response)
	if err := d.Set("items", templates); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
