package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplications() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationsRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_set_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_subtype": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"category_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dscp": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"engine_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"help_string": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ignore_conflict": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"long_description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"popularity": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rank": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"server_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"traffic_class": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"application_network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"lower_port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ports": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"upper_port": &schema.Schema{
										Type:     schema.TypeInt,
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

func dataSourceApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	applicationsQueryParams := dnac.GetApplicationsQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		applicationsQueryParams.Name = v.(string)
	}
	if v, ok := d.GetOk("offset"); ok {
		applicationsQueryParams.Offset = v.(float64)
	}
	if v, ok := d.GetOk("limit"); ok {
		applicationsQueryParams.Limit = v.(float64)
	}

	response, _, err := client.ApplicationPolicy.GetApplications(&applicationsQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	applicationItems := flattenApplicationsReadItems(&response.Response)
	if err := d.Set("items", applicationItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
