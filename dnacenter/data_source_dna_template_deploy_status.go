package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplateDeployStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateDeployStatusRead,
		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"duration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"duration": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
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

func dataSourceTemplateDeployStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deploymentID := d.Get("deployment_id").(string)
	response, _, err := client.ConfigurationTemplates.GetTemplateDeploymentStatus(deploymentID)
	if err != nil {
		return diag.FromErr(err)
	}

	templates := flattenTemplateDeployStatusReadItem(response)
	if err := d.Set("item", templates); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
