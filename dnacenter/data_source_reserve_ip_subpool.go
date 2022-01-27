package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReserveIPSubpool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get the ip subpool info.
`,

		ReadContext: dataSourceReserveIPSubpoolRead,
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
			"site_id": &schema.Schema{
				Description: `siteId query parameter. site id to get the reserve ip associated with the site
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"group_name": &schema.Schema{
							Description: `Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"group_owner": &schema.Schema{
							Description: `Group Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_pools": &schema.Schema{
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
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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

									"group_uuid": &schema.Schema{
										Description: `Group Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
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
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Last Update Time`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"overlapping": &schema.Schema{
										Description: `Overlapping`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceReserveIPSubpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetReserveIPSubpool")
		queryParams1 := dnacentersdkgo.GetReserveIPSubpoolQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetReserveIPSubpool(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetReserveIPSubpool", err,
				"Failure at GetReserveIPSubpool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetReserveIPSubpoolItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetReserveIPSubpool response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetReserveIPSubpoolItems(items *[]dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["group_name"] = item.GroupName
		respItem["ip_pools"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPools(item.IPPools)
		respItem["site_id"] = item.SiteID
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["type"] = item.Type
		respItem["group_owner"] = item.GroupOwner
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenNetworkSettingsGetReserveIPSubpoolItem(item *dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}

	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["group_name"] = item.GroupName
	respItem["ip_pools"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPools(item.IPPools)
	respItem["site_id"] = item.SiteID
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["type"] = item.Type
	respItem["group_owner"] = item.GroupOwner
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetReserveIPSubpoolItemsIPPools(items *[]dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPools) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["dhcp_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDhcpServerIPs(item.DhcpServerIPs)
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
		respItem["client_options"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsClientOptions(item.ClientOptions)
		respItem["group_uuid"] = item.GroupUUID
		respItem["dns_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDNSServerIPs(item.DNSServerIPs)
		respItem["context"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsContext(item.Context)
		respItem["ipv6"] = boolPtrToString(item.IPv6)
		respItem["id"] = item.ID
		respItem["ip_pool_cidr"] = item.IPPoolCidr
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenNetworkSettingsGetReserveIPSubpoolItemIPPools(item *dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPools) []map[string]interface{} {
	if item == nil {
		return nil
	}

	respItem := make(map[string]interface{})
	respItem["ip_pool_name"] = item.IPPoolName
	respItem["dhcp_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDhcpServerIPs(item.DhcpServerIPs)
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
	respItem["client_options"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsClientOptions(item.ClientOptions)
	respItem["group_uuid"] = item.GroupUUID
	respItem["dns_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDNSServerIPs(item.DNSServerIPs)
	respItem["context"] = flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsContext(item.Context)
	respItem["ipv6"] = boolPtrToString(item.IPv6)
	respItem["id"] = item.ID
	respItem["ip_pool_cidr"] = item.IPPoolCidr

	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDhcpServerIPs(items *[]dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDhcpServerIPs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsClientOptions(item *dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsClientOptions) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsDNSServerIPs(items *[]dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDNSServerIPs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolItemsIPPoolsContext(items *[]dnacentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsContext) []map[string]interface{} {
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
