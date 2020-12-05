package dnacenter

import (
	"context"
	dnac "dnacenter-go-sdk/sdk"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceConfigPreview() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceConfigPreviewRead,
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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"complete": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"config": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"error": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"error_message": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"expired_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rf_profile": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensor_profile": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"task_id": &schema.Schema{
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
	}
}

func dataSourcePnPDeviceConfigPreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	previewConfigRequest := dnac.PreviewConfigRequest{}

	if v, ok := d.GetOk("device_id"); ok {
		previewConfigRequest.DeviceID = v.(string)
	}
	if v, ok := d.GetOk("site_id"); ok {
		previewConfigRequest.SiteID = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		previewConfigRequest.Type = v.(string)
	}

	response, _, err := client.DeviceOnboardingPnP.PreviewConfig(&previewConfigRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenPnPDeviceConfigPreviewReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
