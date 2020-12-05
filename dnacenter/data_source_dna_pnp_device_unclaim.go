package dnacenter

import (
	"context"
	dnac "dnacenter-go-sdk/sdk"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceUnclaim() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceUnclaimRead,
		Schema: map[string]*schema.Schema{
			"device_id_list": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"json_array_response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_code": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnPDeviceUnclaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deviceIDListInterface := d.Get("device_id_list").([]interface{})
	deviceIDList := convertSliceInterfaceToSliceString(deviceIDListInterface)

	unClaimDeviceRequest := dnac.UnClaimDeviceRequest{DeviceIDList: deviceIDList}

	response, _, err := client.DeviceOnboardingPnP.UnClaimDevice(&unClaimDeviceRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenPnPDeviceUnClaimReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
