package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

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
				Description: `limit query parameter. Max number of data entries in the returned data set [1,50].  Default is 25
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset of the first returned data set entry (Multiple of 'limit' + 1)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_type": &schema.Schema{
				Description: `siteType query parameter. site type: AREA or BUILDING (case insensitive)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the Site Hierarchy data is required
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_good_count": &schema.Schema{
							Description: `Number of GOOD health access devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"access_total_count": &schema.Schema{
							Description: `Number of access devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ap_device_good_count": &schema.Schema{
							Description: `Number of GOOD health AP devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ap_device_total_count": &schema.Schema{
							Description: `Number of AP devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"application_bytes_total_count": &schema.Schema{
							Description: `Total application bytes
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"application_good_count": &schema.Schema{
							Description: `Number of GOOD health applications int the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"application_health": &schema.Schema{
							Description: `Average application health in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"application_health_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bytes_count": &schema.Schema{
										Description: `Byte count of the application
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"health_score": &schema.Schema{
										Description: `Health score of the application
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"traffic_class": &schema.Schema{
										Description: `Traffic class of the application
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"application_health_stats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_total_count": &schema.Schema{
										Description: `Total application count
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"business_irrelevant_app_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fair": &schema.Schema{
													Description: `Fair business irrelevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"good": &schema.Schema{
													Description: `Good business irrelevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"poor": &schema.Schema{
													Description: `Poor business irrelevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
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
													Description: `Fair business relevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"good": &schema.Schema{
													Description: `Good business relevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"poor": &schema.Schema{
													Description: `Poor business relevant application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
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
													Description: `Fair default application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"good": &schema.Schema{
													Description: `Good default application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"poor": &schema.Schema{
													Description: `Poor default application count
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"application_total_count": &schema.Schema{
							Description: `Number of applications int the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_health_wired": &schema.Schema{
							Description: `Health of all wired clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_health_wireless": &schema.Schema{
							Description: `Health of all wireless clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"core_good_count": &schema.Schema{
							Description: `Number of GOOD health core devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"core_total_count": &schema.Schema{
							Description: `Number of core devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"distribution_good_count": &schema.Schema{
							Description: `Number of GOOD health distribution devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"distribution_total_count": &schema.Schema{
							Description: `Number of distribution devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"dnac_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": &schema.Schema{
										Description: `IP address of the DNAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Status of the DNAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"uuid": &schema.Schema{
										Description: `UUID of the DNAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"healthy_clients_percentage": &schema.Schema{
							Description: `Client health of all clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"healthy_network_device_percentage": &schema.Schema{
							Description: `Network health of devices on the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"latitude": &schema.Schema{
							Description: `Site (building) location's latitude
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"longitude": &schema.Schema{
							Description: `Site (building) location's longitude
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"network_health_ap": &schema.Schema{
							Description: `Network health for AP devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_access": &schema.Schema{
							Description: `Network health for access devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_average": &schema.Schema{
							Description: `Average network health in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_core": &schema.Schema{
							Description: `Network health for core devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_distribution": &schema.Schema{
							Description: `Network health for distribution devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_others": &schema.Schema{
							Description: `Network health for other devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_router": &schema.Schema{
							Description: `Network health for router devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_switch": &schema.Schema{
							Description: `Network health for switch devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_wlc": &schema.Schema{
							Description: `Network health for WLC devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_health_wireless": &schema.Schema{
							Description: `Network health for wireless devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"number_of_clients": &schema.Schema{
							Description: `Total number of clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"number_of_network_device": &schema.Schema{
							Description: `Total number of network devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"number_of_wired_clients": &schema.Schema{
							Description: `Number of wired clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"number_of_wireless_clients": &schema.Schema{
							Description: `Number of wireless clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"overall_good_devices": &schema.Schema{
							Description: `Number of GOOD health devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"parent_site_id": &schema.Schema{
							Description: `The parent site's UUID of this site
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_site_name": &schema.Schema{
							Description: `The parent site's name of this site
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"router_good_count": &schema.Schema{
							Description: `Number of GOOD health router in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"router_total_count": &schema.Schema{
							Description: `Number of router devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"site_id": &schema.Schema{
							Description: `Site UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name": &schema.Schema{
							Description: `Name of the site
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_type": &schema.Schema{
							Description: `Site type of this site
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"switch_device_good_count": &schema.Schema{
							Description: `Number of GOOD health switch devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"switch_device_total_count": &schema.Schema{
							Description: `Number of switch devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"total_number_of_active_wireless_clients": &schema.Schema{
							Description: `Number of active wireless clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"total_number_of_connected_wired_clients": &schema.Schema{
							Description: `Number of connected wired clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"usage": &schema.Schema{
							Description: `Total bits used by all clients in a site
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"wired_good_clients": &schema.Schema{
							Description: `Number of GOOD health wired clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wireless_device_good_count": &schema.Schema{
							Description: `Number of GOOD health wireless devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wireless_device_total_count": &schema.Schema{
							Description: `Number of wireless devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wireless_good_clients": &schema.Schema{
							Description: `Number of GOOD health wireless clients in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wlc_device_good_count": &schema.Schema{
							Description: `Number of GOOD health wireless controller devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wlc_device_total_count": &schema.Schema{
							Description: `Number of wireless controller devices in the site
`,
							Type:     schema.TypeInt,
							Computed: true,
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
	vSiteType, okSiteType := d.GetOk("site_type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vTimestamp, okTimestamp := d.GetOk("timestamp")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteHealth")
		queryParams1 := dnacentersdkgo.GetSiteHealthQueryParams{}

		if okSiteType {
			queryParams1.SiteType = vSiteType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(float64)
		}

		response1, restyResp1, err := client.Sites.GetSiteHealth(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteHealth", err,
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
		respItem["healthy_network_device_percentage"] = item.HealthyNetworkDevicePercentage
		respItem["healthy_clients_percentage"] = item.HealthyClientsPercentage
		respItem["client_health_wired"] = item.ClientHealthWired
		respItem["client_health_wireless"] = item.ClientHealthWireless
		respItem["number_of_clients"] = item.NumberOfClients
		respItem["number_of_network_device"] = item.NumberOfNetworkDevice
		respItem["network_health_average"] = item.NetworkHealthAverage
		respItem["network_health_access"] = item.NetworkHealthAccess
		respItem["network_health_core"] = item.NetworkHealthCore
		respItem["network_health_distribution"] = item.NetworkHealthDistribution
		respItem["network_health_router"] = item.NetworkHealthRouter
		respItem["network_health_wireless"] = item.NetworkHealthWireless
		respItem["network_health_ap"] = item.NetworkHealthAP
		respItem["network_health_wlc"] = item.NetworkHealthWLC
		respItem["network_health_switch"] = item.NetworkHealthSwitch
		respItem["network_health_others"] = item.NetworkHealthOthers
		respItem["number_of_wired_clients"] = item.NumberOfWiredClients
		respItem["number_of_wireless_clients"] = item.NumberOfWirelessClients
		respItem["total_number_of_connected_wired_clients"] = item.TotalNumberOfConnectedWiredClients
		respItem["total_number_of_active_wireless_clients"] = item.TotalNumberOfActiveWirelessClients
		respItem["wired_good_clients"] = item.WiredGoodClients
		respItem["wireless_good_clients"] = item.WirelessGoodClients
		respItem["overall_good_devices"] = item.OverallGoodDevices
		respItem["access_good_count"] = item.AccessGoodCount
		respItem["access_total_count"] = item.AccessTotalCount
		respItem["core_good_count"] = item.CoreGoodCount
		respItem["core_total_count"] = item.CoreTotalCount
		respItem["distribution_good_count"] = item.DistributionGoodCount
		respItem["distribution_total_count"] = item.DistributionTotalCount
		respItem["router_good_count"] = item.RouterGoodCount
		respItem["router_total_count"] = item.RouterTotalCount
		respItem["wireless_device_good_count"] = item.WirelessDeviceGoodCount
		respItem["wireless_device_total_count"] = item.WirelessDeviceTotalCount
		respItem["ap_device_good_count"] = item.ApDeviceGoodCount
		respItem["ap_device_total_count"] = item.ApDeviceTotalCount
		respItem["wlc_device_good_count"] = item.WlcDeviceGoodCount
		respItem["wlc_device_total_count"] = item.WlcDeviceTotalCount
		respItem["switch_device_good_count"] = item.SwitchDeviceGoodCount
		respItem["switch_device_total_count"] = item.SwitchDeviceTotalCount
		respItem["application_health"] = item.ApplicationHealth
		respItem["application_health_info"] = flattenSitesGetSiteHealthItemsApplicationHealthInfo(item.ApplicationHealthInfo)
		respItem["application_good_count"] = item.ApplicationGoodCount
		respItem["application_total_count"] = item.ApplicationTotalCount
		respItem["application_bytes_total_count"] = item.ApplicationBytesTotalCount
		respItem["dnac_info"] = flattenSitesGetSiteHealthItemsDnacInfo(item.DnacInfo)
		respItem["usage"] = item.Usage
		respItem["application_health_stats"] = flattenSitesGetSiteHealthItemsApplicationHealthStats(item.ApplicationHealthStats)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteHealthItemsApplicationHealthInfo(items *[]dnacentersdkgo.ResponseSitesGetSiteHealthResponseApplicationHealthInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["traffic_class"] = item.TrafficClass
		respItem["bytes_count"] = item.BytesCount
		respItem["health_score"] = item.HealthScore
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetSiteHealthItemsDnacInfo(item *dnacentersdkgo.ResponseSitesGetSiteHealthResponseDnacInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["uuid"] = item.UUID
	respItem["ip"] = item.IP
	respItem["status"] = item.Status

	return []map[string]interface{}{
		respItem,
	}

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
