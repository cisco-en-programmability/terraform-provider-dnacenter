package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpVirtualAccountSyncResult() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the summary of devices synced from the given smart account & virtual account with PnP
`,

		ReadContext: dataSourcePnpVirtualAccountSyncResultRead,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Description: `domain path parameter. Smart Account Domain
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Description: `name path parameter. Virtual Account Name
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auto_sync_period": &schema.Schema{
							Description: `Auto Sync Period`,
							Type:        schema.TypeFloat,
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
										Type:        schema.TypeFloat,
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

						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_sn_list": &schema.Schema{
													Description: `Device Sn List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"sync_type": &schema.Schema{
													Description: `Sync Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"sync_msg": &schema.Schema{
										Description: `Sync Msg`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"sync_result_str": &schema.Schema{
							Description: `Sync Result Str`,
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

						"token": &schema.Schema{
							Description: `Token`,
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
		},
	}
}

func dataSourcePnpVirtualAccountSyncResultRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDomain := d.Get("domain")
	vName := d.Get("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSyncResultForVirtualAccount")
		vvDomain := vDomain.(string)
		vvName := vName.(string)

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetSyncResultForVirtualAccount(vvDomain, vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSyncResultForVirtualAccount", err,
				"Failure at GetSyncResultForVirtualAccount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSyncResultForVirtualAccount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItem(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["auto_sync_period"] = item.AutoSyncPeriod
	respItem["sync_result_str"] = item.SyncResultStr
	respItem["profile"] = flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemProfile(item.Profile)
	respItem["cco_user"] = item.CcoUser
	respItem["sync_result"] = flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemSyncResult(item.SyncResult)
	respItem["token"] = item.Token
	respItem["sync_start_time"] = item.SyncStartTime
	respItem["last_sync"] = item.LastSync
	respItem["tenant_id"] = item.TenantID
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["expiry"] = item.Expiry
	respItem["sync_status"] = item.SyncStatus
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemProfile(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["proxy"] = boolPtrToString(item.Proxy)
	respItem["make_default"] = boolPtrToString(item.MakeDefault)
	respItem["port"] = item.Port
	respItem["profile_id"] = item.ProfileID
	respItem["name"] = item.Name
	respItem["address_ip_v4"] = item.AddressIPV4
	respItem["cert"] = item.Cert
	respItem["address_fqdn"] = item.AddressFqdn

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemSyncResult(item *dnacentersdkgo.ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sync_list"] = flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemSyncResultSyncList(item.SyncList)
	respItem["sync_msg"] = item.SyncMsg

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetSyncResultForVirtualAccountItemSyncResultSyncList(items *[]dnacentersdkgo.ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResultSyncList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sync_type"] = item.SyncType
		respItem["device_sn_list"] = item.DeviceSnList
		respItems = append(respItems, respItem)
	}
	return respItems
}
