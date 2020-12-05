package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpHistoryWorkItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_taken": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func pnpHistoryKeyValueMap() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPDeviceHistory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceHistoryRead,
		Schema: map[string]*schema.Schema{
			"sort": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"asc", "des"}),
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_flag": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"history_task_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addn_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     pnpHistoryKeyValueMap(),
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_taken": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     pnpHistoryWorkItems(),
									},
								},
							},
						},
						"timestamp": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnPDeviceHistoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParam := dnac.GetDeviceHistoryQueryParams{}

	if v, ok := d.GetOk("sort"); ok {
		queryParam.Sort = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("sort_order"); ok {
		queryParam.SortOrder = v.(string)
	}
	if v, ok := d.GetOk("serial_number"); ok {
		queryParam.SerialNumber = v.(string)
	}

	// Prepare Request
	response, _, err := client.DeviceOnboardingPnP.GetDeviceHistory(&queryParam)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source

	historyResponse := response.Response
	items := flattenPnPDevicesHistoryReadItems(&historyResponse)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
