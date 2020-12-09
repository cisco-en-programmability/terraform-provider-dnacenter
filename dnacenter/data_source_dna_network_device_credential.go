package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceCredential() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNetworkDeviceCredentialRead,
		Schema: map[string]*schema.Schema{

			"site_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cli": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// ...
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
									"enable_password": &schema.Schema{
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
									"password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"http_read": &schema.Schema{
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
									"password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"secure": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"http_write": &schema.Schema{
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
									"password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"secure": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"snmp_v2_read": &schema.Schema{
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
									"read_community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"snmp_v2_write": &schema.Schema{
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
									"write_community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
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
									"privacy_password": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"privacy_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"snmp_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
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

func dataSourceNetworkDeviceCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParams := dnac.GetDeviceCredentialDetailsQueryParams{}
	if v, ok := d.GetOk("site_id"); ok {
		queryParams.SiteID = v.(string)
	}

	response, _, err := client.NetworkSettings.GetDeviceCredentialDetails(&queryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenNetworkDeviceCredentialReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
