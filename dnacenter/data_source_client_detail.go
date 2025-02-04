package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

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
				Type:     schema.TypeFloat,
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
										Description: `The band at which the host is connected. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"channel": &schema.Schema{
										Description: `The channel used by the host. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"channel_width": &schema.Schema{
										Description: `The channel width used by the host. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_type": &schema.Schema{
										Description: `Host Type - WIRELESS or WIRED
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nw_device_mac": &schema.Schema{
										Description: `Device MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"nw_device_name": &schema.Schema{
										Description: `Name of the network device it is connected to. In case of wireless, it would be an AccessPoint
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Connection Protocol used. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_stream": &schema.Schema{
										Description: `The spatial stream of host. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"timestamp": &schema.Schema{
										Description: `Epoch/Unix time in milliseconds
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"uapsd": &schema.Schema{
										Description: `The UAPSD of the host. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wmm": &schema.Schema{
										Description: `The wmm of the host. This information is present for wireless hosts only
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"detail": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_server_eap_latency": &schema.Schema{
										Description: `The AAA server EAP latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"aaa_server_failed_transaction": &schema.Schema{
										Description: `The number of failed AAA server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"aaa_server_ip": &schema.Schema{
										Description: `The AAA server IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"aaa_server_latency": &schema.Schema{
										Description: `The AAA server latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"aaa_server_mab_latency": &schema.Schema{
										Description: `The AAA server MAB latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"aaa_server_success_transaction": &schema.Schema{
										Description: `The number of successful AAA server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"aaa_server_transaction": &schema.Schema{
										Description: `The number of AAA server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ap_group": &schema.Schema{
										Description: `AP group to which the client belongs
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_type": &schema.Schema{
										Description: `Authentication mechanism of the client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"avg_rssi": &schema.Schema{
										Description: `Average RSSI value for the client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"avg_snr": &schema.Schema{
										Description: `Average signal to noise ratio for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bridge_vmmode": &schema.Schema{
										Description: `Bridge VM mode
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"channel": &schema.Schema{
										Description: `Channel to which the client is connected
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"client_connection": &schema.Schema{
										Description: `AP/Switch to which the client is connected
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"client_type": &schema.Schema{
										Description: `OLD or NEW
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"connected_device": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"band": &schema.Schema{
													Description: `Band of the AP
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Unique identifier of the device
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip_address": &schema.Schema{
													Description: `Management IP address of the connected device.  (deprecated soon in favor of 'mgmtIp')
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"mac": &schema.Schema{
													Description: `MAC address of the access point
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"mgmt_ip": &schema.Schema{
													Description: `The IP address of the connected device
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"mode": &schema.Schema{
													Description: `The mode of the access point
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name of the device
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `Type of device (AP or SWITCH)
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"connected_upn": &schema.Schema{
										Description: `connected UPN ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"connected_upn_id": &schema.Schema{
										Description: `Connected UPN ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"connected_upn_owner": &schema.Schema{
										Description: `Connected UPN owner
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"connection_status": &schema.Schema{
										Description: `The client is connected, connecting, or disconnected
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"country_code": &schema.Schema{
										Description: `The country code of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"data_rate": &schema.Schema{
										Description: `MCS data rates for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_form": &schema.Schema{
										Description: `The device form of the host (e.g. Phone/Tablet)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_vendor": &schema.Schema{
										Description: `The device vendor string
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_decline_ip": &schema.Schema{
										Description: `DHCP decline IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_nak_ip": &schema.Schema{
										Description: `DHCP NAK IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_server_dolatency": &schema.Schema{
										Description: `The DHCP server DO latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"dhcp_server_failed_transaction": &schema.Schema{
										Description: `The number of failed DHCP server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"dhcp_server_ip": &schema.Schema{
										Description: `The DHCP server IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dhcp_server_latency": &schema.Schema{
										Description: `The DHCP server latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"dhcp_server_ralatency": &schema.Schema{
										Description: `The DHCP RA latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"dhcp_server_success_transaction": &schema.Schema{
										Description: `The number of successful DHCP server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"dhcp_server_transaction": &schema.Schema{
										Description: `The number of DHCP server transactions
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"dns_request": &schema.Schema{
										Description: `DNS request attempts for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dns_response": &schema.Schema{
										Description: `DNS response attempts for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot11_protocol": &schema.Schema{
										Description: `Description of dot11 protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot11_protocol_capability": &schema.Schema{
										Description: `description of dot11 protocol capability
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"duid": &schema.Schema{
										Description: `Device UID for MAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"firmware_version": &schema.Schema{
										Description: `The firmware version of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"frequency": &schema.Schema{
										Description: `Frequency band to which the client is connected
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"health_score": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"health_type": &schema.Schema{
													Description: `Type of device health
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"reason": &schema.Schema{
													Description: `Reason for the health score value
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"score": &schema.Schema{
													Description: `health score of client device in the range of 1 to 10. Value 0 for a client represents an IDLE client
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"host_ip_v4": &schema.Schema{
										Description: `IPv4 address of the interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_ip_v6": &schema.Schema{
										Description: `List of IPv6 addresses
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"host_mac": &schema.Schema{
										Description: `MAC address of the interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_name": &schema.Schema{
										Description: `The hostname of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_os": &schema.Schema{
										Description: `The OS of host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_type": &schema.Schema{
										Description: `WIRED or WIRELESS
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"host_version": &schema.Schema{
										Description: `The version of OS of host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"hw_model": &schema.Schema{
										Description: `Hardware model
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Unique identifier representing a specific host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"identifier": &schema.Schema{
										Description: `The host's unique identifier, which is populated by and in order of userId, hostName, hostIpV4, hostIpV6, or hostMac
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"intel_capable": &schema.Schema{
										Description: `Whether support Intel device analytics
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ios_capable": &schema.Schema{
										Description: `IOS Capable`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_guest_upn_endpoint": &schema.Schema{
										Description: `Whether it is guest UPN endpoint
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"issue_count": &schema.Schema{
										Description: `Issue count for a device
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"l2_virtual_network": &schema.Schema{
										Description: `Comma separated Level 2 virtual network names
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"l3_virtual_network": &schema.Schema{
										Description: `Comma separated Level 3 virtual network names
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_updated": &schema.Schema{
										Description: `Epoch/Unix time in milliseconds
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"latency_be": &schema.Schema{
										Description: `Best-effort latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"latency_bg": &schema.Schema{
										Description: `Background latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"latency_video": &schema.Schema{
										Description: `Video latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"latency_voice": &schema.Schema{
										Description: `Voice latency
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"link_speed": &schema.Schema{
										Description: `The speed of wired client
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"link_threshold": &schema.Schema{
										Description: `Link error threshold of wired client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"location": &schema.Schema{
										Description: `Site location of client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_roaming_duration": &schema.Schema{
										Description: `Max roaming duration for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"model_name": &schema.Schema{
										Description: `System model
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"onboarding": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa_rootcause_list": &schema.Schema{
													Description: `Root cause list of AAA failure category
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"aaa_server_ip": &schema.Schema{
													Description: `AAA server IP for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"assoc_done_time": &schema.Schema{
													Description: `Epoch/Unix time in milliseconds
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"assoc_rootcause_list": &schema.Schema{
													Description: `Root cause list of ASSOC failure category
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"auth_done_time": &schema.Schema{
													Description: `Epoch/Unix time in milliseconds
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"average_assoc_duration": &schema.Schema{
													Description: `Average association duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"average_auth_duration": &schema.Schema{
													Description: `Average auth duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"average_dhcp_duration": &schema.Schema{
													Description: `Average DHCP duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"average_run_duration": &schema.Schema{
													Description: `Average run Duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dhcp_done_time": &schema.Schema{
													Description: `Epoch/Unix time in milliseconds
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dhcp_rootcause_list": &schema.Schema{
													Description: `Root cause list of DHCP failure category
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"dhcp_server_ip": &schema.Schema{
													Description: `DHCP server IP for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"latest_root_cause_list": &schema.Schema{
													Description: `Root cause list of latest root cause category
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"max_assoc_duration": &schema.Schema{
													Description: `Max association duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_auth_duration": &schema.Schema{
													Description: `Max auth duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_dhcp_duration": &schema.Schema{
													Description: `Max DHCP duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_run_duration": &schema.Schema{
													Description: `Max run duration for a client
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"other_rootcause_list": &schema.Schema{
													Description: `Root cause list of other failure category
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"onboarding_time": &schema.Schema{
										Description: `Epoch/Unix time in milliseconds
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"port": &schema.Schema{
										Description: `switch port for client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"port_description": &schema.Schema{
										Description: `Port desctiption of wired client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_type": &schema.Schema{
										Description: `AC/DC voltage
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"private_mac": &schema.Schema{
										Description: `Private Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"remote_end_duplex_mode": &schema.Schema{
										Description: `The remote end duplex mode of wired client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rssi": &schema.Schema{
										Description: `Min RSSI value for the client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rssi_is_include": &schema.Schema{
										Description: `RSSI include/exclude flag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_bytes": &schema.Schema{
										Description: `Total received bytes for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_link_error": &schema.Schema{
										Description: `The error of rx link
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"rx_rate": &schema.Schema{
										Description: `The rate of rx
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"rx_retry_pct": &schema.Schema{
										Description: `The retry count as percentage wrt to total rx packets
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sales_code": &schema.Schema{
										Description: `The Sales Code of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"session_duration": &schema.Schema{
										Description: `Time duration the session took from run time to delete time
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sgt": &schema.Schema{
										Description: `Security group tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"slot_id": &schema.Schema{
										Description: `Slot ID of AP which client is connected
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"snr": &schema.Schema{
										Description: `Min signal to noise ratio for the client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"snr_is_include": &schema.Schema{
										Description: `SNR include/exclude flag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"snr_threshold": &schema.Schema{
										Description: `SNR threshold
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid": &schema.Schema{
										Description: `WLAN SSID to which the client is connected
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sub_type": &schema.Schema{
										Description: `The device type of host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tracked": &schema.Schema{
										Description: `Tracking status of this host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"trust_details": &schema.Schema{
										Description: `Trust details explaining reason for corresponding Trust score
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"trust_score": &schema.Schema{
										Description: `Trust score of Client received from EndPoint Analytics
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tx_bytes": &schema.Schema{
										Description: `total transmitted bytes for a client
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tx_link_error": &schema.Schema{
										Description: `The error of tx link
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"tx_rate": &schema.Schema{
										Description: `The rate of tx
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"upn_id": &schema.Schema{
										Description: `Registered UPN ID of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"upn_name": &schema.Schema{
										Description: `Registered UPN name of the host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"upn_owner": &schema.Schema{
										Description: `Owner of registered UPN name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"usage": &schema.Schema{
										Description: `Usage of txBytes and rxBytes of client
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"user_id": &schema.Schema{
										Description: `The user ID of this host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_time": &schema.Schema{
										Description: `The metric modification time of the new version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"vlan_id": &schema.Schema{
										Description: `VLAN ID for the host
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"vnid": &schema.Schema{
										Description: `VNID of the host
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"wlc_name": &schema.Schema{
										Description: `The name of the connected wireless controller
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlc_uuid": &schema.Schema{
										Description: `UUID of the WLC the client connected to
`,
										Type:     schema.TypeString,
										Computed: true,
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

												"ap_radio_admin_status": &schema.Schema{
													Description: `The admin status of the radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ap_radio_oper_status": &schema.Schema{
													Description: `The oper status of the radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Identifier of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_details": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Description: `The admin status of network device interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"client_mac_address": &schema.Schema{
																Description: `The MAC address of the client device
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"connected_device_int_name": &schema.Schema{
																Description: `The interface name of the network device
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"duplex": &schema.Schema{
																Description: `The duplex info of the network device interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"port_mode": &schema.Schema{
																Description: `The port mode info of network device interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"label": &schema.Schema{
													Description: `The details of the edge
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"link_status": &schema.Schema{
													Description: `Link status of the link
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"port_utilization": &schema.Schema{
													Description: `Port utilization
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"source": &schema.Schema{
													Description: `Edge line starting node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_admin_status": &schema.Schema{
													Description: `The admin status of the source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_duplex_info": &schema.Schema{
													Description: `The duplex info of the source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_interface_name": &schema.Schema{
													Description: `The interface name of the source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_link_status": &schema.Schema{
													Description: `The status of the link
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_port_mode": &schema.Schema{
													Description: `The port mode of the source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_port_vla_n_info": &schema.Schema{
													Description: `List of VLANs configured on the source port
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target": &schema.Schema{
													Description: `End node of the edge line
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_admin_status": &schema.Schema{
													Description: `The admin status of the target
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_duplex_info": &schema.Schema{
													Description: `The duplex info of the target
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_interface_name": &schema.Schema{
													Description: `The interface name of the target
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_link_status": &schema.Schema{
													Description: `The status of the link
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_port_mode": &schema.Schema{
													Description: `The port mode of the target
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"target_port_vla_n_info": &schema.Schema{
													Description: `List of VLANs configured on the target port
`,
													Type:     schema.TypeString,
													Computed: true,
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
													Description: `Number of clients
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"connected_device": &schema.Schema{
													Description: `Connected Device`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"count": &schema.Schema{
													Description: `Count`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"description": &schema.Schema{
													Description: `Description of the topology node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"device_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"fabric_group": &schema.Schema{
													Description: `Fabric Group
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"fabric_role": &schema.Schema{
													Description: `Fabric Role
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"health_score": &schema.Schema{
													Description: `Device health score
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `User ID, hostname, IP address, or MAC address
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `Device IP
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipv6": &schema.Schema{
													Description: `Device IPv6
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"level": &schema.Schema{
													Description: `Level in the topology
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Device name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"node_type": &schema.Schema{
													Description: `Node type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"platform_id": &schema.Schema{
													Description: `Device platform ID
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"radio_frequency": &schema.Schema{
													Description: `Radio frequency
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"role": &schema.Schema{
													Description: `Device role
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"software_version": &schema.Schema{
													Description: `Device software version
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"stack_type": &schema.Schema{
													Description: `Stack Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"user_id": &schema.Schema{
													Description: `User ID
`,
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
				},
			},
		},
	}
}

func dataSourceClientDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vMacAddress := d.Get("mac_address")
	vTimestamp, okTimestamp := d.GetOk("timestamp")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetClientDetail")
		queryParams1 := dnacentersdkgo.GetClientDetailQueryParams{}

		queryParams1.MacAddress = vMacAddress.(string)

		if okTimestamp {
			queryParams1.Timestamp = vTimestamp.(float64)
		}

		response1, restyResp1, err := client.Clients.GetClientDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetClientDetail", err,
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
	respItem["tracked"] = item.Tracked
	respItem["host_type"] = item.HostType
	respItem["user_id"] = item.UserID
	respItem["duid"] = item.Duid
	respItem["identifier"] = item.IDentifier
	respItem["host_name"] = item.HostName
	respItem["host_os"] = item.HostOs
	respItem["host_version"] = item.HostVersion
	respItem["sub_type"] = item.SubType
	respItem["firmware_version"] = item.FirmwareVersion
	respItem["device_vendor"] = item.DeviceVendor
	respItem["device_form"] = item.DeviceForm
	respItem["sales_code"] = item.SalesCode
	respItem["country_code"] = item.CountryCode
	respItem["last_updated"] = item.LastUpdated
	respItem["health_score"] = flattenClientsGetClientDetailItemDetailHealthScore(item.HealthScore)
	respItem["host_mac"] = item.HostMac
	respItem["host_ip_v4"] = item.HostIPV4
	respItem["host_ip_v6"] = item.HostIPV6
	respItem["auth_type"] = item.AuthType
	respItem["vlan_id"] = item.VLANID
	respItem["l3_virtual_network"] = item.L3VirtualNetwork
	respItem["l2_virtual_network"] = item.L2VirtualNetwork
	respItem["vnid"] = item.Vnid
	respItem["upn_id"] = item.UpnID
	respItem["upn_name"] = item.UpnName
	respItem["ssid"] = item.SSID
	respItem["frequency"] = item.Frequency
	respItem["channel"] = item.Channel
	respItem["ap_group"] = item.ApGroup
	respItem["sgt"] = item.Sgt
	respItem["location"] = item.Location
	respItem["client_connection"] = item.ClientConnection
	respItem["connected_device"] = flattenClientsGetClientDetailItemDetailConnectedDevice(item.ConnectedDevice)
	respItem["issue_count"] = item.IssueCount
	respItem["rssi"] = item.Rssi
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["rssi_is_include"] = item.RssiIsInclude
	respItem["avg_rssi"] = item.AvgRssi
	respItem["snr"] = item.Snr
	respItem["snr_threshold"] = item.SnrThreshold
	respItem["snr_is_include"] = item.SnrIsInclude
	respItem["avg_snr"] = item.AvgSnr
	respItem["data_rate"] = item.DataRate
	respItem["tx_bytes"] = item.TxBytes
	respItem["rx_bytes"] = item.RxBytes
	respItem["dns_response"] = item.DNSResponse
	respItem["dns_request"] = item.DNSRequest
	respItem["onboarding"] = flattenClientsGetClientDetailItemDetailOnboarding(item.Onboarding)
	respItem["client_type"] = item.ClientType
	respItem["onboarding_time"] = item.OnboardingTime
	respItem["port"] = item.Port
	respItem["ios_capable"] = boolPtrToString(item.IosCapable)
	respItem["usage"] = item.Usage
	respItem["link_speed"] = item.LinkSpeed
	respItem["link_threshold"] = item.LinkThreshold
	respItem["remote_end_duplex_mode"] = item.RemoteEndDuplexMode
	respItem["tx_link_error"] = item.TxLinkError
	respItem["rx_link_error"] = item.RxLinkError
	respItem["tx_rate"] = item.TxRate
	respItem["rx_rate"] = item.RxRate
	respItem["rx_retry_pct"] = item.RxRetryPct
	respItem["version_time"] = item.VersionTime
	respItem["dot11_protocol"] = item.Dot11Protocol
	respItem["slot_id"] = item.SlotID
	respItem["dot11_protocol_capability"] = item.Dot11ProtocolCapability
	respItem["private_mac"] = boolPtrToString(item.PrivateMac)
	respItem["dhcp_server_ip"] = item.DhcpServerIP
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["aaa_server_transaction"] = item.AAAServerTransaction
	respItem["aaa_server_failed_transaction"] = item.AAAServerFailedTransaction
	respItem["aaa_server_success_transaction"] = item.AAAServerSuccessTransaction
	respItem["aaa_server_latency"] = item.AAAServerLatency
	respItem["aaa_server_mab_latency"] = item.AAAServerMABLatency
	respItem["aaa_server_eap_latency"] = item.AAAServerEApLatency
	respItem["dhcp_server_transaction"] = item.DhcpServerTransaction
	respItem["dhcp_server_failed_transaction"] = item.DhcpServerFailedTransaction
	respItem["dhcp_server_success_transaction"] = item.DhcpServerSuccessTransaction
	respItem["dhcp_server_latency"] = item.DhcpServerLatency
	respItem["dhcp_server_dolatency"] = item.DhcpServerDOLatency
	respItem["dhcp_server_ralatency"] = item.DhcpServerRALatency
	respItem["max_roaming_duration"] = item.MaxRoamingDuration
	respItem["upn_owner"] = item.UpnOwner
	respItem["connected_upn"] = item.ConnectedUpn
	respItem["connected_upn_owner"] = item.ConnectedUpnOwner
	respItem["connected_upn_id"] = item.ConnectedUpnID
	respItem["is_guest_upn_endpoint"] = boolPtrToString(item.IsGuestUPNEndpoint)
	respItem["wlc_name"] = item.WlcName
	respItem["wlc_uuid"] = item.WlcUUID
	respItem["session_duration"] = item.SessionDuration
	respItem["intel_capable"] = boolPtrToString(item.IntelCapable)
	respItem["hw_model"] = item.HwModel
	respItem["power_type"] = item.PowerType
	respItem["model_name"] = item.ModelName
	respItem["bridge_vmmode"] = item.BridgeVMMode
	respItem["dhcp_nak_ip"] = item.DhcpNakIP
	respItem["dhcp_decline_ip"] = item.DhcpDeclineIP
	respItem["port_description"] = item.PortDescription
	respItem["latency_voice"] = item.LatencyVoice
	respItem["latency_video"] = item.LatencyVideo
	respItem["latency_bg"] = item.LatencyBg
	respItem["latency_be"] = item.LatencyBe
	respItem["trust_score"] = item.TrustScore
	respItem["trust_details"] = item.TrustDetails

	return []map[string]interface{}{
		respItem,
	}

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

func flattenClientsGetClientDetailItemDetailConnectedDevice(items *[]dnacentersdkgo.ResponseClientsGetClientDetailDetailConnectedDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["name"] = item.Name
		respItem["mac"] = item.Mac
		respItem["id"] = item.ID
		respItem["ip_address"] = item.IPaddress
		respItem["mgmt_ip"] = item.MgmtIP
		respItem["band"] = item.Band
		respItem["mode"] = item.Mode
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientDetailItemDetailOnboarding(item *dnacentersdkgo.ResponseClientsGetClientDetailDetailOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["average_run_duration"] = item.AverageRunDuration
	respItem["max_run_duration"] = item.MaxRunDuration
	respItem["average_assoc_duration"] = item.AverageAssocDuration
	respItem["max_assoc_duration"] = item.MaxAssocDuration
	respItem["average_auth_duration"] = item.AverageAuthDuration
	respItem["max_auth_duration"] = item.MaxAuthDuration
	respItem["average_dhcp_duration"] = item.AverageDhcpDuration
	respItem["max_dhcp_duration"] = item.MaxDhcpDuration
	respItem["aaa_server_ip"] = item.AAAServerIP
	respItem["dhcp_server_ip"] = item.DhcpServerIP
	respItem["auth_done_time"] = item.AuthDoneTime
	respItem["assoc_done_time"] = item.AssocDoneTime
	respItem["dhcp_done_time"] = item.DhcpDoneTime
	respItem["assoc_rootcause_list"] = item.AssocRootcauseList
	respItem["aaa_rootcause_list"] = item.AAARootcauseList
	respItem["dhcp_rootcause_list"] = item.DhcpRootcauseList
	respItem["other_rootcause_list"] = item.OtherRootcauseList
	respItem["latest_root_cause_list"] = item.LatestRootCauseList

	return []map[string]interface{}{
		respItem,
	}

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
		respItem["platform_id"] = item.PlatformID
		respItem["family"] = item.Family
		respItem["ip"] = item.IP
		respItem["ipv6"] = item.IPv6
		respItem["software_version"] = item.SoftwareVersion
		respItem["user_id"] = item.UserID
		respItem["node_type"] = item.NodeType
		respItem["radio_frequency"] = item.RadioFrequency
		respItem["clients"] = item.Clients
		respItem["count"] = item.Count
		respItem["health_score"] = item.HealthScore
		respItem["level"] = item.Level
		respItem["fabric_group"] = item.FabricGroup
		respItem["fabric_role"] = item.FabricRole
		respItem["connected_device"] = item.ConnectedDevice
		respItem["stack_type"] = item.StackType
		respItems = append(respItems, respItem)
	}
	return respItems
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
		respItem["source_link_status"] = item.SourceLinkStatus
		respItem["target_link_status"] = item.TargetLinkStatus
		respItem["label"] = item.Label
		respItem["target"] = item.Target
		respItem["id"] = item.ID
		respItem["port_utilization"] = item.PortUtilization
		respItem["source_interface_name"] = item.SourceInterfaceName
		respItem["target_interface_name"] = item.TargetInterfaceName
		respItem["source_duplex_info"] = item.SourceDuplexInfo
		respItem["target_duplex_info"] = item.TargetDuplexInfo
		respItem["source_port_mode"] = item.SourcePortMode
		respItem["target_port_mode"] = item.TargetPortMode
		respItem["source_admin_status"] = item.SourceAdminStatus
		respItem["target_admin_status"] = item.TargetAdminStatus
		respItem["ap_radio_admin_status"] = item.ApRadioAdminStatus
		respItem["ap_radio_oper_status"] = item.ApRadioOperStatus
		respItem["source_port_vla_n_info"] = item.SourcePortVLANInfo
		respItem["target_port_vla_n_info"] = item.TargetPortVLANInfo
		respItem["interface_details"] = flattenClientsGetClientDetailItemTopologyLinksInterfaceDetails(item.InterfaceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientDetailItemTopologyLinksInterfaceDetails(items *[]dnacentersdkgo.ResponseClientsGetClientDetailTopologyLinksInterfaceDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["client_mac_address"] = item.ClientMacAddress
		respItem["connected_device_int_name"] = item.ConnectedDeviceIntName
		respItem["duplex"] = item.Duplex
		respItem["port_mode"] = item.PortMode
		respItem["admin_status"] = item.AdminStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}
