package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGlobalPool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get global pool.
`,

		ReadContext: dataSourceGlobalPoolRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. No of Global Pools to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_options": &schema.Schema{
							Description: `Client Options`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"configure_external_dhcp": &schema.Schema{
							Description: `Configure External Dhcp`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"context_key": &schema.Schema{
										Description: `Context Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"context_value": &schema.Schema{
										Description: `Context Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"owner": &schema.Schema{
										Description: `Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"create_time": &schema.Schema{
							Description: `Create Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"dhcp_server_ips": &schema.Schema{
							Description: `Dhcp Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"dns_server_ips": &schema.Schema{
							Description: `Dns Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"gateways": &schema.Schema{
							Description: `Gateways`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_cidr": &schema.Schema{
							Description: `Ip Pool Cidr`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"overlapping": &schema.Schema{
							Description: `Overlapping`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"owner": &schema.Schema{
							Description: `Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_uuid": &schema.Schema{
							Description: `Parent Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"shared": &schema.Schema{
							Description: `Shared`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"total_ip_address_count": &schema.Schema{
							Description: `Total Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"used_ip_address_count": &schema.Schema{
							Description: `Used Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"used_percentage": &schema.Schema{
							Description: `Used Percentage`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalPool")
		queryParams1 := dnacentersdkgo.GetGlobalPoolQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetGlobalPool(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetGlobalPool", err,
				"Failure at GetGlobalPool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetGlobalPoolItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalPool response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetGlobalPoolItems(items *[]dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["dhcp_server_ips"] = item.DhcpServerIPs
		respItem["gateways"] = item.Gateways
		respItem["create_time"] = item.CreateTime
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["total_ip_address_count"] = item.TotalIPAddressCount
		respItem["used_ip_address_count"] = item.UsedIPAddressCount
		respItem["parent_uuid"] = item.ParentUUID
		respItem["owner"] = item.Owner
		respItem["shared"] = boolPtrToString(item.Shared)
		respItem["overlapping"] = boolPtrToString(item.Overlapping)
		respItem["configure_external_dhcp"] = boolPtrToString(item.ConfigureExternalDhcp)
		respItem["used_percentage"] = item.UsedPercentage
		respItem["client_options"] = flattenNetworkSettingsGetGlobalPoolItemsClientOptions(item.ClientOptions)
		respItem["dns_server_ips"] = item.DNSServerIPs
		respItem["context"] = flattenNetworkSettingsGetGlobalPoolItemsContext(item.Context)
		respItem["ipv6"] = boolPtrToString(item.IPv6)
		respItem["id"] = item.ID
		respItem["ip_pool_cidr"] = item.IPPoolCidr
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetGlobalPoolItem(item *dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse) map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ip_pool_name"] = item.IPPoolName
	respItem["dhcp_server_ips"] = item.DhcpServerIPs
	respItem["gateways"] = item.Gateways
	respItem["create_time"] = item.CreateTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["total_ip_address_count"] = item.TotalIPAddressCount
	respItem["used_ip_address_count"] = item.UsedIPAddressCount
	respItem["parent_uuid"] = item.ParentUUID
	respItem["owner"] = item.Owner
	respItem["shared"] = boolPtrToString(item.Shared)
	respItem["overlapping"] = boolPtrToString(item.Overlapping)
	respItem["configure_external_dhcp"] = boolPtrToString(item.ConfigureExternalDhcp)
	respItem["used_percentage"] = item.UsedPercentage
	respItem["client_options"] = flattenNetworkSettingsGetGlobalPoolItemsClientOptions(item.ClientOptions)
	respItem["dns_server_ips"] = item.DNSServerIPs
	respItem["context"] = flattenNetworkSettingsGetGlobalPoolItemsContext(item.Context)
	respItem["ipv6"] = boolPtrToString(item.IPv6)
	respItem["id"] = item.ID
	respItem["ip_pool_cidr"] = item.IPPoolCidr

	return respItem
}

func flattenNetworkSettingsGetGlobalPoolItemsClientOptions(item *dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponseClientOptions) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenNetworkSettingsGetGlobalPoolItemsContext(items *[]dnacentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponseContext) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["owner"] = item.Owner
		respItem["context_key"] = item.ContextKey
		respItem["context_value"] = item.ContextValue
		respItems = append(respItems, respItem)
	}
	return respItems
}
