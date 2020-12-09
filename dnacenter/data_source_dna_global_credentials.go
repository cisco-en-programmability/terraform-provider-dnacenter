package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryGlobalCredentials() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryGlobalCredentialsRead,
		Schema: map[string]*schema.Schema{
			"credential_sub_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"CLI", "SNMPV2_READ_COMMUNITY", "SNMPV2_WRITE_COMMUNITY", "SNMPV3", "HTTP_WRITE", "HTTP_READ", "NETCONF"}),
			},
			"sort_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"asc", "des"}),
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryGlobalCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetGlobalCredentialsQueryParams{}
	queryParams.CredentialSubType = d.Get("credential_sub_type").(string)
	if v, ok := d.GetOk("sort_by"); ok {
		queryParams.SortBy = v.(string)
	}
	if v, ok := d.GetOk("order"); ok {
		queryParams.Order = v.(string)
	}

	response, _, err := client.Discovery.GetGlobalCredentials(&queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenDiscoveryReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
