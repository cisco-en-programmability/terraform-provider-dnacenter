package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func discoveryRangeHTTPCredentialParam() *schema.Resource {
	return &schema.Resource{
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
	}
}

func dataSourceDiscoveryRange() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryRangeRead,
		Schema: map[string]*schema.Schema{
			"start_index": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerGeqThan(0),
			},
			"records_to_return": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerGeqThan(0),
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cdp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"discovery_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_password_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"global_credential_id_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_read_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,

							Elem: discoveryRangeHTTPCredentialParam(),
						},
						"http_write_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,

							Elem: discoveryRangeHTTPCredentialParam(),
						},
						"ip_address_list": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"lldp_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"no_add_new_device": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"parent_discovery_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"preferred_mgmt_ip_method": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol_order": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"re_discovery": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"retry": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Computed: true,
							Type:     schema.TypeString,
						},
						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_ro_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_rw_community_desc": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"update_mgmt_ip": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"user_name_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"device_ids": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_devices": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"discovery_condition": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_auto_cdp": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	startIndex := d.Get("start_index").(int)
	recordsToReturn := d.Get("records_to_return").(int)

	response, _, err := client.Discovery.GetDiscoveriesByRange(startIndex, recordsToReturn)
	if err != nil {
		return diag.FromErr(err)
	}

	items := flattenDiscoveryByRangeReadItems(response)
	if err := d.Set("items", items); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
