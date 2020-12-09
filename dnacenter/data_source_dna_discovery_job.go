package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryJobRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cli_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_time": &schema.Schema{
							Type:     schema.TypeInt,
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
						"inventory_collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"netconf_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ping_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_status": &schema.Schema{
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
		},
	}
}

func dataSourceDiscoveryJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetListOfDiscoveriesByDiscoveryIDQueryParams{}

	id := d.Get("id").(string)
	if v, ok := d.GetOk("offset"); ok {
		queryParams.Offset = v.(int)
	}
	if v, ok := d.GetOk("limit"); ok {
		queryParams.Limit = v.(int)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		queryParams.IPAddress = v.(string)
	}

	response, _, err := client.Discovery.GetListOfDiscoveriesByDiscoveryID(id, &queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenDiscoveryJobReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
