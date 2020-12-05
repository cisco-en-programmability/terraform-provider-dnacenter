package dnacenter

import (
	"context"
	dnac "dnacenter-go-sdk/sdk"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceClaimSite() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceClaimSiteRead,
		Schema: map[string]*schema.Schema{

			"device_id": &schema.Schema{
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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response": &schema.Schema{
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
		},
	}
}

func dataSourcePnPDeviceClaimSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	claimADeviceToASiteRequest := dnac.ClaimADeviceToASiteRequest{}

	if v, ok := d.GetOk("device_id"); ok {
		claimADeviceToASiteRequest.DeviceID = v.(string)
	}
	if v, ok := d.GetOk("site_id"); ok {
		claimADeviceToASiteRequest.SiteID = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		claimADeviceToASiteRequest.Type = v.(string)
	}

	response, _, err := client.DeviceOnboardingPnP.ClaimADeviceToASite(&claimADeviceToASiteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenPnPDeviceClaimSiteReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
