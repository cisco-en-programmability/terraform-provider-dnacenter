package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteMembership() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSiteMembershipRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"offset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"response": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"response": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
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
						"site": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"response": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parent_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"additional_info": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"group_type_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"group_hierarchy": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"group_name_hierarchy": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"instance_tenant_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"version": &schema.Schema{
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

func dataSourceSiteMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteID := d.Get("site_id").(string)

	queryParams := dnac.GetMembershipQueryParams{}
	if v, ok := d.GetOk("offset"); ok {
		queryParams.Offset = v.(string)
	}
	if v, ok := d.GetOk("limit"); ok {
		queryParams.Limit = v.(string)
	}
	if v, ok := d.GetOk("device_family"); ok {
		queryParams.DeviceFamily = v.(string)
	}
	if v, ok := d.GetOk("serial_number"); ok {
		queryParams.SerialNumber = v.(string)
	}

	response, _, err := client.Sites.GetMembership(siteID, &queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	membershipResponse := flattenSiteMembershipReadItem(response)
	if err := d.Set("response", membershipResponse); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
