package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSiteCountRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSiteCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	SiteCountQueryParams := dnac.GetSiteCountQueryParams{}
	if v, ok := d.GetOk("site_id"); ok {
		SiteCountQueryParams.SiteID = v.(string)
	}

	response, _, err := client.Sites.GetSiteCount(&SiteCountQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("response", response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
