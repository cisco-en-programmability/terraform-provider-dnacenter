package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSite() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSiteRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
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
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"latitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"longitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"height": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"rf_model": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"width": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	SiteQueryParams := dnac.GetSiteQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		SiteQueryParams.Name = v.(string)
	}
	if v, ok := d.GetOk("site_id"); ok {
		SiteQueryParams.SiteID = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		SiteQueryParams.Type = v.(string)
	}
	if v, ok := d.GetOk("offset"); ok {
		SiteQueryParams.Offset = v.(string)
	}
	if v, ok := d.GetOk("limit"); ok {
		SiteQueryParams.Limit = v.(string)
	}

	response, _, err := client.Sites.GetSite(&SiteQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	siteItems := flattenSiteReadItem(response)
	if err := d.Set("items", siteItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
