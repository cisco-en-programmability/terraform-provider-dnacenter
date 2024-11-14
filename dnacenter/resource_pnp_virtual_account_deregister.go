package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpVirtualAccountDeregister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Device Onboarding (PnP).

- Deregisters the specified smart account & virtual account info and the associated device information from the PnP
System & database. The devices associated with the deregistered virtual account are removed from the PnP database as
well. The response payload contains the deregistered smart & virtual account information
`,

		CreateContext: resourcePnpVirtualAccountDeregisterCreate,
		ReadContext:   resourcePnpVirtualAccountDeregisterRead,
		DeleteContext: resourcePnpVirtualAccountDeregisterDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auto_sync_period": &schema.Schema{
							Description: `Auto Sync Period`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"cco_user": &schema.Schema{
							Description: `Cco User`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"expiry": &schema.Schema{
							Description: `Expiry`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"last_sync": &schema.Schema{
							Description: `Last Sync`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address_fqdn": &schema.Schema{
										Description: `Address Fqdn`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"address_ip_v4": &schema.Schema{
										Description: `Address Ip V4`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"address_ip_v6": &schema.Schema{
										Description: `Address Ip V6`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"cert": &schema.Schema{
										Description: `Cert`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"make_default": &schema.Schema{
										Description: `Make Default`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"profile_id": &schema.Schema{
										Description: `Profile Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"proxy": &schema.Schema{
										Description: `Proxy`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Description: `Smart Account Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sync_start_time": &schema.Schema{
							Description: `Sync Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"sync_status": &schema.Schema{
							Description: `Sync Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"virtual_account_id": &schema.Schema{
							Description: `Virtual Account Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": &schema.Schema{
							Description: `domain query parameter. Smart Account Domain
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"name": &schema.Schema{
							Description: `name query parameter. Virtual Account Name
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourcePnpVirtualAccountDeregisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDomain := resourceItem["domain"]

	vName := resourceItem["name"]

	queryParams1 := dnacentersdkgo.DeregisterVirtualAccountQueryParams{}

	queryParams1.Domain = vDomain.(string)

	queryParams1.Name = vName.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.DeregisterVirtualAccount(&queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing DeregisterVirtualAccount", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpDeregisterVirtualAccountItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeregisterVirtualAccount response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourcePnpVirtualAccountDeregisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpVirtualAccountDeregisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenDeviceOnboardingPnpDeregisterVirtualAccountItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpDeregisterVirtualAccount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["last_sync"] = item.LastSync
	respItem["cco_user"] = item.CcoUser
	respItem["expiry"] = item.Expiry
	respItem["auto_sync_period"] = item.AutoSyncPeriod
	respItem["profile"] = flattenDeviceOnboardingPnpDeregisterVirtualAccountItemProfile(item.Profile)
	respItem["sync_status"] = item.SyncStatus
	respItem["sync_start_time"] = item.SyncStartTime
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpDeregisterVirtualAccountItemProfile(item *dnacentersdkgo.ResponseDeviceOnboardingPnpDeregisterVirtualAccountProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["profile_id"] = item.ProfileID
	respItem["make_default"] = boolPtrToString(item.MakeDefault)
	respItem["address_ip_v4"] = item.AddressIPV4
	respItem["address_ip_v6"] = item.AddressIPV6
	respItem["address_fqdn"] = item.AddressFqdn
	respItem["port"] = item.Port
	respItem["cert"] = item.Cert
	respItem["proxy"] = boolPtrToString(item.Proxy)

	return []map[string]interface{}{
		respItem,
	}

}
