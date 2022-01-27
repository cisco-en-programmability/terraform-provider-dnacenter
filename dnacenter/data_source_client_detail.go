package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Returns detailed Client information retrieved by Mac Address for any given point of time.
`,

		ReadContext: dataSourceClientDetailRead,
		Schema: map[string]*schema.Schema{
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. MAC Address of the client
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp": &schema.Schema{
				Description: `timestamp query parameter. Epoch time(in milliseconds) when the Client health data is required
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"band": &schema.Schema{
										Description: `Band`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel": &schema.Schema{
										Description: `Channel`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel_width": &schema.Schema{
										Description: `Channel Width`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_type": &schema.Schema{
										Description: `Host Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"nw_device_mac": &schema.Schema{
										Description: `Nw Device Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"nw_device_name": &schema.Schema{
										Description: `Nw Device Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"spatial_stream": &schema.Schema{
										Description: `Spatial Stream`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"uapsd": &schema.Schema{
										Description: `Uapsd`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wmm": &schema.Schema{
										Description: `Wmm`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"detail": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_group": &schema.Schema{
										Description: `Ap Group`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"avg_rssi": &schema.Schema{
										Description: `Avg Rssi`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"avg_snr": &schema.Schema{
										Description: `Avg Snr`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel": &schema.Schema{
										Description: `Channel`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"client_connection": &schema.Schema{
										Description: `Client Connection`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"client_type": &schema.Schema{
										Description: `Client Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_device": &schema.Schema{
										Description: `Connected Device`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"connection_status": &schema.Schema{
										Description: `Connection Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"data_rate": &schema.Schema{
										Description: `Data Rate`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dns_failure": &schema.Schema{
										Description: `Dns Failure`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dns_success": &schema.Schema{
										Description: `Dns Success`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"frequency": &schema.Schema{
										Description: `Frequency`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"health_score": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"health_type": &schema.Schema{
													Description: `Health Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reason": &schema.Schema{
													Description: `Reason`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"score": &schema.Schema{
													Description: `Score`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},

									"host_ip_v4": &schema.Schema{
										Description: `Host Ip V4`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_ip_v6": &schema.Schema{
										Description: `Host Ip V6`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"host_mac": &schema.Schema{
										Description: `Host Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_os": &schema.Schema{
										Description: `Host Os`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_type": &schema.Schema{
										Description: `Host Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_version": &schema.Schema{
										Description: `Host Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ios_capable": &schema.Schema{
										Description: `Ios Capable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"issue_count": &schema.Schema{
										Description: `Issue Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"last_updated": &schema.Schema{
										Description: `Last Updated`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"location": &schema.Schema{
										Description: `Location`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"onboarding": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa_rootcause_list": &schema.Schema{
													Description: `Aaa Rootcause List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"aaa_server_ip": &schema.Schema{
													Description: `Aaa Server Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"assoc_done_time": &schema.Schema{
													Description: `Assoc Done Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"assoc_rootcause_list": &schema.Schema{
													Description: `Assoc Rootcause List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"auth_done_time": &schema.Schema{
													Description: `Auth Done Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"average_assoc_duration": &schema.Schema{
													Description: `Average Assoc Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"average_auth_duration": &schema.Schema{
													Description: `Average Auth Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"average_dhcp_duration": &schema.Schema{
													Description: `Average Dhcp Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"average_run_duration": &schema.Schema{
													Description: `Average Run Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"dhcp_done_time": &schema.Schema{
													Description: `Dhcp Done Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"dhcp_rootcause_list": &schema.Schema{
													Description: `Dhcp Rootcause List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"dhcp_server_ip": &schema.Schema{
													Description: `Dhcp Server Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"max_assoc_duration": &schema.Schema{
													Description: `Max Assoc Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"max_auth_duration": &schema.Schema{
													Description: `Max Auth Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"max_dhcp_duration": &schema.Schema{
													Description: `Max Dhcp Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"max_run_duration": &schema.Schema{
													Description: `Max Run Duration`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"other_rootcause_list": &schema.Schema{
													Description: `Other Rootcause List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"onboarding_time": &schema.Schema{
										Description: `Onboarding Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"rx_bytes": &schema.Schema{
										Description: `Rx Bytes`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sub_type": &schema.Schema{
										Description: `Sub Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vnid": &schema.Schema{
										Description: `Vnid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"topology": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"links": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"link_status": &schema.Schema{
													Description: `Link Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"port_utilization": &schema.Schema{
													Description: `Port Utilization`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"source": &schema.Schema{
													Description: `Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"target": &schema.Schema{
													Description: `Target`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"nodes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"clients": &schema.Schema{
													Description: `Clients`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"connected_device": &schema.Schema{
													Description: `Connected Device`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"count": &schema.Schema{
													Description: `Count`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"device_type": &schema.Schema{
													Description: `Device Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"fabric_group": &schema.Schema{
													Description: `Fabric Group`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"family": &schema.Schema{
													Description: `Family`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"health_score": &schema.Schema{
													Description: `Health Score`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"ip": &schema.Schema{
													Description: `Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"level": &schema.Schema{
													Description: `Level`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"node_type": &schema.Schema{
													Description: `Node Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"platform_id": &schema.Schema{
													Description: `Platform Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"radio_frequency": &schema.Schema{
													Description: `Radio Frequency`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"role": &schema.Schema{
													Description: `Role`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"software_version": &schema.Schema{
													Description: `Software Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"user_id": &schema.Schema{
													Description: `User Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
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

func dataSourceClientDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTimestamp, okTimestamp := d.GetOk("timestamp")
	vMacAddress := d.Get("mac_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetClientDetail")
		queryParams1 := dnacentersdkgo.GetClientDetailQueryParams{}

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(string)
		}
		queryParams1.MacAddress = vMacAddress.(string)

		response1, restyResp1, err := client.Clients.GetClientDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetClientDetail", err,
				"Failure at GetClientDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenClientsGetClientDetailItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetClientDetail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsGetClientDetailItem(item *dnacentersdkgo.ResponseClientsGetClientDetail) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["detail"] = flattenClientsGetClientDetailItemDetail(item.Detail)
	respItem["connection_info"] = flattenClientsGetClientDetailItemConnectionInfo(item.ConnectionInfo)
	respItem["topology"] = flattenClientsGetClientDetailItemTopology(item.Topology)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenClientsGetClientDetailItemDetail(item *dnacentersdkgo.ResponseClientsGetClientDetailDetail) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["connection_status"] = item.ConnectionStatus
	respItem["host_type"] = item.HostType
	respItem["user_id"] = flattenClientsGetClientDetailItemDetailUserID(item.UserID)
	respItem["host_name"] = item.HostName
	respItem["host_os"] = flattenClientsGetClientDetailItemDetailHostOs(item.HostOs)
	respItem["host_version"] = flattenClientsGetClientDetailItemDetailHostVersion(item.HostVersion)
	respItem["sub_type"] = item.SubType
	respItem["last_updated"] = item.LastUpdated
	respItem["health_score"] = flattenClientsGetClientDetailItemDetailHealthScore(item.HealthScore)
	respItem["host_mac"] = item.HostMac
	respItem["host_ip_v4"] = item.HostIPV4
	respItem["host_ip_v6"] = item.HostIPV6
	respItem["auth_type"] = item.AuthType
	respItem["vlan_id"] = item.VLANID
	respItem["vnid"] = item.Vnid
	respItem["ssid"] = item.SSID
	respItem["frequency"] = item.Frequency
	respItem["channel"] = item.Channel
	respItem["ap_group"] = flattenClientsGetClientDetailItemDetailApGroup(item.ApGroup)
	respItem["location"] = flattenClientsGetClientDetailItemDetailLocation(item.Location)
	respItem["client_connection"] = item.ClientConnection
	respItem["connected_device"] = flattenClientsGetClientDetailItemDetailConnectedDevice(item.ConnectedDevice)
	respItem["issue_count"] = item.IssueCount
	respItem["rssi"] = item.Rssi
	respItem["avg_rssi"] = flattenClientsGetClientDetailItemDetailAvgRssi(item.AvgRssi)
	respItem["snr"] = item.Snr
	respItem["avg_snr"] = flattenClientsGetClientDetailItemDetailAvgSnr(item.AvgSnr)
	respItem["data_rate"] = item.DataRate
	respItem["tx_bytes"] = item.TxBytes
	respItem["rx_bytes"] = item.RxBytes
	respItem["dns_success"] = flattenClientsGetClientDetailItemDetailDNSSuccess(item.DNSSuccess)
	respItem["dns_failure"] = flattenClientsGetClientDetailItemDetailDNSFailure(item.DNSFailure)
	respItem["onboarding"] = flattenClientsGetClientDetailItemDetailOnboarding(item.Onboarding)
	respItem["client_type"] = item.ClientType
	respItem["onboarding_time"] = flattenClientsGetClientDetailItemDetailOnboardingTime(item.OnboardingTime)
	respItem["port"] = flattenClientsGetClientDetailItemDetailPort(item.Port)
	respItem["ios_capable"] = boolPtrToString(item.IosCapable)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientDetailItemDetailUserID(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailHostOs(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailHostOs) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailHostVersion(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailHostVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailHealthScore(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailHealthScore) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["health_type"] = item.HealthType
		respItem["reason"] = item.Reason
		respItem["score"] = item.Score
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientDetailItemDetailApGroup(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailApGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailLocation(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailConnectedDevice(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailConnectedDevice) []interface{} {
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

func flattenClientsGetClientDetailItemDetailAvgRssi(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailAvgRssi) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailAvgSnr(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailAvgSnr) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailDNSSuccess(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailDNSSuccess) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailDNSFailure(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailDNSFailure) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboarding(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["average_run_duration"] = flattenClientsGetClientDetailItemDetailOnboardingAverageRunDuration(item.AverageRunDuration)
	respItem["max_run_duration"] = flattenClientsGetClientDetailItemDetailOnboardingMaxRunDuration(item.MaxRunDuration)
	respItem["average_assoc_duration"] = flattenClientsGetClientDetailItemDetailOnboardingAverageAssocDuration(item.AverageAssocDuration)
	respItem["max_assoc_duration"] = flattenClientsGetClientDetailItemDetailOnboardingMaxAssocDuration(item.MaxAssocDuration)
	respItem["average_auth_duration"] = flattenClientsGetClientDetailItemDetailOnboardingAverageAuthDuration(item.AverageAuthDuration)
	respItem["max_auth_duration"] = flattenClientsGetClientDetailItemDetailOnboardingMaxAuthDuration(item.MaxAuthDuration)
	respItem["average_dhcp_duration"] = flattenClientsGetClientDetailItemDetailOnboardingAverageDhcpDuration(item.AverageDhcpDuration)
	respItem["max_dhcp_duration"] = flattenClientsGetClientDetailItemDetailOnboardingMaxDhcpDuration(item.MaxDhcpDuration)
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["dhcp_server_ip"] = flattenClientsGetClientDetailItemDetailOnboardingDhcpServerIP(item.DhcpServerIP)
	respItem["auth_done_time"] = flattenClientsGetClientDetailItemDetailOnboardingAuthDoneTime(item.AuthDoneTime)
	respItem["assoc_done_time"] = flattenClientsGetClientDetailItemDetailOnboardingAssocDoneTime(item.AssocDoneTime)
	respItem["dhcp_done_time"] = flattenClientsGetClientDetailItemDetailOnboardingDhcpDoneTime(item.DhcpDoneTime)
	respItem["assoc_rootcause_list"] = flattenClientsGetClientDetailItemDetailOnboardingAssocRootcauseList(item.AssocRootcauseList)
	respItem["aaa_rootcause_list"] = flattenClientsGetClientDetailItemDetailOnboardingAAARootcauseList(item.AAARootcauseList)
	respItem["dhcp_rootcause_list"] = flattenClientsGetClientDetailItemDetailOnboardingDhcpRootcauseList(item.DhcpRootcauseList)
	respItem["other_rootcause_list"] = flattenClientsGetClientDetailItemDetailOnboardingOtherRootcauseList(item.OtherRootcauseList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientDetailItemDetailOnboardingAverageRunDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAverageRunDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingMaxRunDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingMaxRunDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAverageAssocDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAverageAssocDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingMaxAssocDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingMaxAssocDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAverageAuthDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAverageAuthDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingMaxAuthDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingMaxAuthDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAverageDhcpDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAverageDhcpDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingMaxDhcpDuration(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingMaxDhcpDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingDhcpServerIP(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingDhcpServerIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAuthDoneTime(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAuthDoneTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAssocDoneTime(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAssocDoneTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingDhcpDoneTime(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingDhcpDoneTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailOnboardingAssocRootcauseList(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAssocRootcauseList) []interface{} {
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

func flattenClientsGetClientDetailItemDetailOnboardingAAARootcauseList(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingAAARootcauseList) []interface{} {
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

func flattenClientsGetClientDetailItemDetailOnboardingDhcpRootcauseList(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingDhcpRootcauseList) []interface{} {
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

func flattenClientsGetClientDetailItemDetailOnboardingOtherRootcauseList(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingOtherRootcauseList) []interface{} {
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

func flattenClientsGetClientDetailItemDetailOnboardingTime(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboardingTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemDetailPort(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemConnectionInfo(item *dnacentersdkgo.ResponseClientsGetClientDetailConnectionInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_type"] = item.HostType
	respItem["nw_device_name"] = item.NwDeviceName
	respItem["nw_device_mac"] = item.NwDeviceMac
	respItem["protocol"] = item.Protocol
	respItem["band"] = item.Band
	respItem["spatial_stream"] = item.SpatialStream
	respItem["channel"] = item.Channel
	respItem["channel_width"] = item.ChannelWidth
	respItem["wmm"] = item.Wmm
	respItem["uapsd"] = item.Uapsd
	respItem["timestamp"] = item.Timestamp

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientDetailItemTopology(item *dnacentersdkgo.ResponseClientsGetClientDetailTopology) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["nodes"] = flattenClientsGetClientDetailItemTopologyNodes(item.Nodes)
	respItem["links"] = flattenClientsGetClientDetailItemTopologyLinks(item.Links)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientDetailItemTopologyNodes(items *[]dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["role"] = item.Role
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["device_type"] = item.DeviceType
		respItem["platform_id"] = flattenClientsGetClientDetailItemTopologyNodesPlatformID(item.PlatformID)
		respItem["family"] = flattenClientsGetClientDetailItemTopologyNodesFamily(item.Family)
		respItem["ip"] = item.IP
		respItem["software_version"] = flattenClientsGetClientDetailItemTopologyNodesSoftwareVersion(item.SoftwareVersion)
		respItem["user_id"] = flattenClientsGetClientDetailItemTopologyNodesUserID(item.UserID)
		respItem["node_type"] = item.NodeType
		respItem["radio_frequency"] = flattenClientsGetClientDetailItemTopologyNodesRadioFrequency(item.RadioFrequency)
		respItem["clients"] = flattenClientsGetClientDetailItemTopologyNodesClients(item.Clients)
		respItem["count"] = flattenClientsGetClientDetailItemTopologyNodesCount(item.Count)
		respItem["health_score"] = item.HealthScore
		respItem["level"] = item.Level
		respItem["fabric_group"] = flattenClientsGetClientDetailItemTopologyNodesFabricGroup(item.FabricGroup)
		respItem["connected_device"] = flattenClientsGetClientDetailItemTopologyNodesConnectedDevice(item.ConnectedDevice)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientDetailItemTopologyNodesPlatformID(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesPlatformID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesFamily(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesFamily) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesSoftwareVersion(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesSoftwareVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesUserID(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesRadioFrequency(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesRadioFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesClients(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesCount(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesFabricGroup(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesFabricGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyNodesConnectedDevice(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyNodesConnectedDevice) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyLinks(items *[]dnacentersdkgo.ResponseClientsGetClientDetailTopologyLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["source"] = item.Source
		respItem["link_status"] = item.LinkStatus
		respItem["label"] = item.Label
		respItem["target"] = item.Target
		respItem["id"] = flattenClientsGetClientDetailItemTopologyLinksID(item.ID)
		respItem["port_utilization"] = flattenClientsGetClientDetailItemTopologyLinksPortUtilization(item.PortUtilization)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientDetailItemTopologyLinksID(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyLinksID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientDetailItemTopologyLinksPortUtilization(item *dnacentersdkgo.ResponseClientsGetClientDetailTopologyLinksPortUtilization) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
