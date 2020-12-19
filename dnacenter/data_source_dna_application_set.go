package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationSetRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source": &schema.Schema{
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceApplicationSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetApplicationSetsQueryParams{}
	if v, ok := d.GetOk("name"); ok {
		queryParams.Name = v.(string)
	}
	if v, ok := d.GetOk("offset"); ok {
		queryParams.Offset = v.(float64)
	}
	if v, ok := d.GetOk("limit"); ok {
		queryParams.Limit = v.(float64)
	}

	response, _, err := client.ApplicationPolicy.GetApplicationSets(&queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	appSetItems := flattenApplicationsSetReadItems(&response.Response)
	if err := d.Set("items", appSetItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
