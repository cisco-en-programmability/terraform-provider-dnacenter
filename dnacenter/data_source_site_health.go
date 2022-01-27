package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteHealth() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Returns Overall Health information for all sites
`,

		ReadContext: dataSourceSiteHealthRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The max number of sites in the returned data set.  Default is 25, and max at 50
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset value, starting from 1, of the first returned site entry.  Default is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_type": &schema.Schema{
				Description: `siteType query parameter. Type of the site to return.  AREA or BUILDING.  Default to AREA
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the Site Hierarchy data is required
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_good_count": &schema.Schema{
							Description: `Access Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"access_total_count": &schema.Schema{
							Description: `Access Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"application_bytes_total_count": &schema.Schema{
							Description: `Application Bytes Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"application_good_count": &schema.Schema{
							Description: `Application Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"application_health": &schema.Schema{
							Description: `Application Health`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"application_health_stats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_total_count": &schema.Schema{
										Description: `App Total Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"business_irrelevant_app_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fair": &schema.Schema{
													Description: `Fair`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"good": &schema.Schema{
													Description: `Good`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"poor": &schema.Schema{
													Description: `Poor`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},

									"business_relevant_app_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fair": &schema.Schema{
													Description: `Fair`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"good": &schema.Schema{
													Description: `Good`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"poor": &schema.Schema{
													Description: `Poor`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},

									"default_health_app_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fair": &schema.Schema{
													Description: `Fair`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"good": &schema.Schema{
													Description: `Good`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"poor": &schema.Schema{
													Description: `Poor`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"application_total_count": &schema.Schema{
							Description: `Application Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_health_wired": &schema.Schema{
							Description: `Client Health Wired`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_health_wireless": &schema.Schema{
							Description: `Client Health Wireless`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"core_good_count": &schema.Schema{
							Description: `Core Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"core_total_count": &schema.Schema{
							Description: `Core Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"distribution_good_count": &schema.Schema{
							Description: `Distribution Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"distribution_total_count": &schema.Schema{
							Description: `Distribution Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dnac_info": &schema.Schema{
							Description: `Dnac Info`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"healthy_clients_percentage": &schema.Schema{
							Description: `Healthy Clients Percentage`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"healthy_network_device_percentage": &schema.Schema{
							Description: `Healthy Network Device Percentage`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"latitude": &schema.Schema{
							Description: `Latitude`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"longitude": &schema.Schema{
							Description: `Longitude`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"network_health_access": &schema.Schema{
							Description: `Network Health Access`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_average": &schema.Schema{
							Description: `Network Health Average`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_core": &schema.Schema{
							Description: `Network Health Core`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_distribution": &schema.Schema{
							Description: `Network Health Distribution`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_others": &schema.Schema{
							Description: `Network Health Others`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_router": &schema.Schema{
							Description: `Network Health Router`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_health_wireless": &schema.Schema{
							Description: `Network Health Wireless`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"number_of_clients": &schema.Schema{
							Description: `Number Of Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"number_of_network_device": &schema.Schema{
							Description: `Number Of Network Device`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"number_of_wired_clients": &schema.Schema{
							Description: `Number Of Wired Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"number_of_wireless_clients": &schema.Schema{
							Description: `Number Of Wireless Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"overall_good_devices": &schema.Schema{
							Description: `Overall Good Devices`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_site_id": &schema.Schema{
							Description: `Parent Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_site_name": &schema.Schema{
							Description: `Parent Site Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"router_good_count": &schema.Schema{
							Description: `Router Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"router_total_count": &schema.Schema{
							Description: `Router Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_type": &schema.Schema{
							Description: `Site Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"total_number_of_active_wireless_clients": &schema.Schema{
							Description: `Total Number Of Active Wireless Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"total_number_of_connected_wired_clients": &schema.Schema{
							Description: `Total Number Of Connected Wired Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wired_good_clients": &schema.Schema{
							Description: `Wired Good Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wireless_device_good_count": &schema.Schema{
							Description: `Wireless Device Good Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wireless_device_total_count": &schema.Schema{
							Description: `Wireless Device Total Count`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wireless_good_clients": &schema.Schema{
							Description: `Wireless Good Clients`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteHealthRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTimestamp, okTimestamp := d.GetOk("timestamp")
	vSiteType, okSiteType := d.GetOk("site_type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSiteHealth")
		queryParams1 := dnacentersdkgo.GetSiteHealthQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(string)
		}
		if okSiteType {
			queryParams1.SiteType = vSiteType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sites.GetSiteHealth(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSiteHealth", err,
				"Failure at GetSiteHealth, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetSiteHealthItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteHealth response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetSiteHealthItems(items *[]dnacentersdkgo.ResponseSitesGetSiteHealthResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_name"] = item.SiteName
		respItem["site_id"] = item.SiteID
		respItem["parent_site_id"] = item.ParentSiteID
		respItem["parent_site_name"] = item.ParentSiteName
		respItem["site_type"] = item.SiteType
		respItem["latitude"] = item.Latitude
		respItem["longitude"] = item.Longitude
		respItem["healthy_network_device_percentage"] = flattenSitesGetSiteHealthItemsHealthyNetworkDevicePercentage(item.HealthyNetworkDevicePercentage)
		respItem["healthy_clients_percentage"] = flattenSitesGetSiteHealthItemsHealthyClientsPercentage(item.HealthyClientsPercentage)
		respItem["client_health_wired"] = flattenSitesGetSiteHealthItemsClientHealthWired(item.ClientHealthWired)
		respItem["client_health_wireless"] = flattenSitesGetSiteHealthItemsClientHealthWireless(item.ClientHealthWireless)
		respItem["number_of_clients"] = flattenSitesGetSiteHealthItemsNumberOfClients(item.NumberOfClients)
		respItem["number_of_network_device"] = flattenSitesGetSiteHealthItemsNumberOfNetworkDevice(item.NumberOfNetworkDevice)
		respItem["network_health_average"] = flattenSitesGetSiteHealthItemsNetworkHealthAverage(item.NetworkHealthAverage)
		respItem["network_health_access"] = flattenSitesGetSiteHealthItemsNetworkHealthAccess(item.NetworkHealthAccess)
		respItem["network_health_core"] = flattenSitesGetSiteHealthItemsNetworkHealthCore(item.NetworkHealthCore)
		respItem["network_health_distribution"] = flattenSitesGetSiteHealthItemsNetworkHealthDistribution(item.NetworkHealthDistribution)
		respItem["network_health_router"] = flattenSitesGetSiteHealthItemsNetworkHealthRouter(item.NetworkHealthRouter)
		respItem["network_health_wireless"] = flattenSitesGetSiteHealthItemsNetworkHealthWireless(item.NetworkHealthWireless)
		respItem["network_health_others"] = flattenSitesGetSiteHealthItemsNetworkHealthOthers(item.NetworkHealthOthers)
		respItem["number_of_wired_clients"] = flattenSitesGetSiteHealthItemsNumberOfWiredClients(item.NumberOfWiredClients)
		respItem["number_of_wireless_clients"] = flattenSitesGetSiteHealthItemsNumberOfWirelessClients(item.NumberOfWirelessClients)
		respItem["total_number_of_connected_wired_clients"] = flattenSitesGetSiteHealthItemsTotalNumberOfConnectedWiredClients(item.TotalNumberOfConnectedWiredClients)
		respItem["total_number_of_active_wireless_clients"] = flattenSitesGetSiteHealthItemsTotalNumberOfActiveWirelessClients(item.TotalNumberOfActiveWirelessClients)
		respItem["wired_good_clients"] = flattenSitesGetSiteHealthItemsWiredGoodClients(item.WiredGoodClients)
		respItem["wireless_good_clients"] = flattenSitesGetSiteHealthItemsWirelessGoodClients(item.WirelessGoodClients)
		respItem["overall_good_devices"] = flattenSitesGetSiteHealthItemsOverallGoodDevices(item.OverallGoodDevices)
		respItem["access_good_count"] = flattenSitesGetSiteHealthItemsAccessGoodCount(item.AccessGoodCount)
		respItem["access_total_count"] = flattenSitesGetSiteHealthItemsAccessTotalCount(item.AccessTotalCount)
		respItem["core_good_count"] = flattenSitesGetSiteHealthItemsCoreGoodCount(item.CoreGoodCount)
		respItem["core_total_count"] = flattenSitesGetSiteHealthItemsCoreTotalCount(item.CoreTotalCount)
		respItem["distribution_good_count"] = flattenSitesGetSiteHealthItemsDistributionGoodCount(item.DistributionGoodCount)
		respItem["distribution_total_count"] = flattenSitesGetSiteHealthItemsDistributionTotalCount(item.DistributionTotalCount)
		respItem["router_good_count"] = flattenSitesGetSiteHealthItemsRouterGoodCount(item.RouterGoodCount)
		respItem["router_total_count"] = flattenSitesGetSiteHealthItemsRouterTotalCount(item.RouterTotalCount)
		respItem["wireless_device_good_count"] = flattenSitesGetSiteHealthItemsWirelessDeviceGoodCount(item.WirelessDeviceGoodCount)
		respItem["wireless_device_total_count"] = flattenSitesGetSiteHealthItemsWirelessDeviceTotalCount(item.WirelessDeviceTotalCount)
		respItem["application_health"] = flattenSitesGetSiteHealthItemsApplicationHealth(item.ApplicationHealth)
		respItem["application_good_count"] = flattenSitesGetSiteHealthItemsApplicationGoodCount(item.ApplicationGoodCount)
		respItem["application_total_count"] = flattenSitesGetSiteHealthItemsApplicationTotalCount(item.ApplicationTotalCount)
		respItem["application_bytes_total_count"] = flattenSitesGetSiteHealthItemsApplicationBytesTotalCount(item.ApplicationBytesTotalCount)
		respItem["dnac_info"] = flattenSitesGetSiteHealthItemsDnacInfo(item.DnacInfo)
		respItem["application_health_stats"] = flattenSitesGetSiteHealthItemsApplicationHealthStats(item.ApplicationHealthStats)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteHealthItemsHealthyNetworkDevicePercentage(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseHealthyNetworkDevicePercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsHealthyClientsPercentage(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseHealthyClientsPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsClientHealthWired(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseClientHealthWired) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsClientHealthWireless(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseClientHealthWireless) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNumberOfClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNumberOfClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNumberOfNetworkDevice(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNumberOfNetworkDevice) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthAverage(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthAverage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthAccess(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthAccess) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthCore(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthCore) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthDistribution(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthDistribution) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthRouter(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthRouter) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthWireless(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthWireless) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNetworkHealthOthers(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNetworkHealthOthers) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNumberOfWiredClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNumberOfWiredClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsNumberOfWirelessClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseNumberOfWirelessClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsTotalNumberOfConnectedWiredClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseTotalNumberOfConnectedWiredClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsTotalNumberOfActiveWirelessClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseTotalNumberOfActiveWirelessClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsWiredGoodClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseWiredGoodClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsWirelessGoodClients(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseWirelessGoodClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsOverallGoodDevices(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseOverallGoodDevices) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsAccessGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseAccessGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsAccessTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseAccessTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsCoreGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseCoreGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsCoreTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseCoreTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsDistributionGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseDistributionGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsDistributionTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseDistributionTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsRouterGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseRouterGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsRouterTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseRouterTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsWirelessDeviceGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseWirelessDeviceGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsWirelessDeviceTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseWirelessDeviceTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsApplicationHealth(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealth) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsApplicationGoodCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationGoodCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsApplicationTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsApplicationBytesTotalCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationBytesTotalCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsDnacInfo(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseDnacInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSitesGetSiteHealthItemsApplicationHealthStats(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealthStats) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["app_total_count"] = item.AppTotalCount
	respItem["business_relevant_app_count"] = flattenSitesGetSiteHealthItemsApplicationHealthStatsBusinessRelevantAppCount(item.BusinessRelevantAppCount)
	respItem["business_irrelevant_app_count"] = flattenSitesGetSiteHealthItemsApplicationHealthStatsBusinessIrrelevantAppCount(item.BusinessIrrelevantAppCount)
	respItem["default_health_app_count"] = flattenSitesGetSiteHealthItemsApplicationHealthStatsDefaultHealthAppCount(item.DefaultHealthAppCount)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteHealthItemsApplicationHealthStatsBusinessRelevantAppCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessRelevantAppCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["poor"] = item.Poor
	respItem["fair"] = item.Fair
	respItem["good"] = item.Good

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteHealthItemsApplicationHealthStatsBusinessIrrelevantAppCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessIrrelevantAppCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["poor"] = item.Poor
	respItem["fair"] = item.Fair
	respItem["good"] = item.Good

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetSiteHealthItemsApplicationHealthStatsDefaultHealthAppCount(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealthStatsDefaultHealthAppCount) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["poor"] = item.Poor
	respItem["fair"] = item.Fair
	respItem["good"] = item.Good

	return []map[string]interface{}{
		respItem,
	}

}
