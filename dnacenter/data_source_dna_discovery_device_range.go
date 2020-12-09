package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryDevicesRange() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryDevicesRangeRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"start_index": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerGeqThan(0),
			},
			"records_to_return": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(0, 500),
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"anchor_wlc_for_ap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_model_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"avg_update_frequency": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cli_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"duplicate_device_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"http_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ingress_queue_config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"netconf_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_updates": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ping_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_range": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"qos_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_contact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"up_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wlc_ap_device_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryDevicesRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)
	startIndex := d.Get("start_index").(int)
	recordsToReturn := d.Get("records_to_return").(int)
	queryParams := dnac.GetDiscoveredDevicesByRangeQueryParams{}
	if v, ok := d.GetOk("task_id"); ok {
		queryParams.TaskID = v.(string)
	}

	response, _, err := client.Discovery.GetDiscoveredDevicesByRange(id, startIndex, recordsToReturn, &queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenDiscoveryDevicesByRangeReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
