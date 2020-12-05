package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func templateParamComputed() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instruction_text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"not_param": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"param_array": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parameter_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"provider": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_value": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_value": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"selection": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"selection_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"selection_values": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplateDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateDetailsRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"latest_version": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"author": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"composite": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"failure_policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_template_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rollback_template_content": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_variant": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_content": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"composite": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"product_family": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_series": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     templateParamComputed(),
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem:     templateParamComputed(),
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplateDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	getTemplateDetailsQueryParams := dnac.GetTemplateDetailsQueryParams{}
	templateID := d.Get("template_id").(string)

	if v, ok := d.GetOk("latest_version"); ok {
		getTemplateDetailsQueryParams.LatestVersion = v.(bool)
	}

	response, _, err := client.ConfigurationTemplates.GetTemplateDetails(templateID, &getTemplateDetailsQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	templates := flattenTemplateReadItem(response)
	if err := d.Set("item", templates); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
