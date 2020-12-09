package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoverySummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoverySummaryRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Discovery ID",
				Required:    true,
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sort_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sort_order": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ping_status": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snmp_status": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cli_status": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"netconf_status": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"http_status": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDiscoverySummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetNetworkDevicesFromDiscoveryQueryParams{}
	id := d.Get("id").(string)
	if v, ok := d.GetOk("task_id"); ok {
		queryParams.TaskID = v.(string)
	}
	if v, ok := d.GetOk("sort_by"); ok {
		queryParams.SortBy = v.(string)
	}
	if v, ok := d.GetOk("sort_order"); ok {
		queryParams.SortOrder = v.(string)
	}
	if v, ok := d.GetOk("ip_address"); ok {
		queryParams.IPAddress = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("ping_status"); ok {
		queryParams.PingStatus = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("snmp_status"); ok {
		queryParams.SNMPStatus = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("cli_status"); ok {
		queryParams.CliStatus = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("netconf_status"); ok {
		queryParams.NetconfStatus = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("http_status"); ok {
		queryParams.HTTPStatus = convertSliceInterfaceToSliceString(v.([]interface{}))
	}

	// Prepare Request
	response, _, err := client.Discovery.GetNetworkDevicesFromDiscovery(id, &queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source
	if err := d.Set("response", response.Response); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
