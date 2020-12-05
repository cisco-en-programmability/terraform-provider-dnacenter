package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagQuery() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTagQueryRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_space": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringIsValueFunc("name"),
			},
			"order": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"asc", "des"}),
			},
			"system_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_tag": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": &schema.Schema{
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
						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
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

func dataSourceTagQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	tagQueryParams := dnac.GetTagQueryParams{}

	if v, ok := d.GetOk("name"); ok {
		tagQueryParams.Name = v.(string)
	}
	if v, ok := d.GetOk("name_space"); ok {
		tagQueryParams.AdditionalInfameSpace = v.(string)
	}
	if v, ok := d.GetOk("attributes"); ok {
		tagQueryParams.AdditionalInfttributes = v.(string)
	}
	if v, ok := d.GetOk("level"); ok {
		tagQueryParams.Level = v.(string)
	}
	if v, ok := d.GetOk("offset"); ok {
		tagQueryParams.Offset = v.(string)
	}
	if v, ok := d.GetOk("limit"); ok {
		tagQueryParams.Limit = v.(string)
	}
	if v, ok := d.GetOk("size"); ok {
		tagQueryParams.Size = v.(string)
	}
	if v, ok := d.GetOk("field"); ok {
		tagQueryParams.Field = v.(string)
	}
	if v, ok := d.GetOk("sort_by"); ok {
		tagQueryParams.SortBy = v.(string)
	}
	if v, ok := d.GetOk("order"); ok {
		tagQueryParams.Order = v.(string)
	}
	if v, ok := d.GetOk("system_tag"); ok {
		tagQueryParams.SystemTag = v.(string)
	}

	// Prepare Request
	response, _, err := client.Tag.GetTag(&tagQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	tagsRead := flattenTagQueryReadItems(response)
	if err := d.Set("items", tagsRead); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
