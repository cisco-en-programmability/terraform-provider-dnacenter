package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSensorTestTemplateDuplicate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.

- Intent API to duplicate an existing SENSOR test template
`,

		CreateContext: resourceSensorTestTemplateDuplicateCreate,
		ReadContext:   resourceSensorTestTemplateDuplicateRead,
		DeleteContext: resourceSensorTestTemplateDuplicateDelete,
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

						"id": &schema.Schema{
							Description: `The sensor test template unique identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"action_in_progress": &schema.Schema{
							Description: `Indication of inprogress action
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `The WIFI bands
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number of APs to test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `connection type of test: WIRED, WIRELESS, BOTH
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption mode
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"frequency": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"unit": &schema.Schema{
										Description: `Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Description: `Value of the unit
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"last_modified_time": &schema.Schema{
							Description: `Last modify time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `Location string
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensors": &schema.Schema{
										Description: `Use all sensors in the site for test
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_management_vlan": &schema.Schema{
										Description: `Custom Management VLAN
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_type": &schema.Schema{
										Description: `Site type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address_list": &schema.Schema{
										Description: `MAC addresses
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"management_vlan": &schema.Schema{
										Description: `Management VLAN
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site name hierarhy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"model_version": &schema.Schema{
							Description: `Test template object model version (must be 2)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor test template name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_associated_sensor": &schema.Schema{
							Description: `Number of associated sensor
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"num_neighbor_apthreshold": &schema.Schema{
							Description: `Number of neighboring AP threshold
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"profiles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_vlan_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"location_id": &schema.Schema{
													Description: `Site UUID
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Description: `Array of VLANs
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
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"vlan": &schema.Schema{
										Description: `VLAN
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"radio_as_sensor_removed": &schema.Schema{
							Description: `Radio as sensor removed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rssi_threshold": &schema.Schema{
							Description: `RSSI threshold
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"run_now": &schema.Schema{
							Description: `Run now (YES, NO)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule_in_days": &schema.Schema{
							Description: `Bit-wise value of scheduled test days
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sensors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensor_addition": &schema.Schema{
										Description: `Is all sensor addition
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"assigned": &schema.Schema{
										Description: `Is assigned
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"config_updated": &schema.Schema{
										Description: `Configuration updated: YES, NO
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": &schema.Schema{
										Description: `Host name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"i_perf_info": &schema.Schema{
										Description: `A string-stringList iPerf information
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `Sensor ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `IP address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"marked_for_uninstall": &schema.Schema{
										Description: `Is marked for uninstall
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Sensor name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"run_now": &schema.Schema{
										Description: `Run now: YES, NO
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensor_type": &schema.Schema{
										Description: `Sensor type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_policy": &schema.Schema{
										Description: `Service policy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Sensor device status: UP, DOWN, REBOOT
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_mac": &schema.Schema{
										Description: `Switch MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_serial_number": &schema.Schema{
										Description: `Switch serial number
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_uuid": &schema.Schema{
										Description: `Switch device UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_a_ps": &schema.Schema{
										Description: `Array of target APs
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"test_mac_addresses": &schema.Schema{
										Description: `A string-string test MAC address
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},
									"wired_application_message": &schema.Schema{
										Description: `Wired application message
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wired_application_status": &schema.Schema{
										Description: `Wired application status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"xor_sensor": &schema.Schema{
										Description: `Is XOR sensor
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"show_wlc_upgrade_banner": &schema.Schema{
							Description: `Show WLC upgrade banner
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"bands": &schema.Schema{
										Description: `WIFI bands: 2.4GHz or 5GHz
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Identification number
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer 3 WEB Auth email address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer 3 WEB Auth password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer 3 WEB Auth security
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer 3 WEB Auth user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"num_aps": &schema.Schema{
										Description: `Number of APs in the test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"num_sensors": &schema.Schema{
										Description: `Number of Sensors in the test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `The SSID profile name string
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_password": &schema.Schema{
										Description: `Proxy server password
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_port": &schema.Schema{
										Description: `Proxy server port
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_server": &schema.Schema{
										Description: `Proxy server for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_user_name": &schema.Schema{
										Description: `Proxy server user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Description: `The SSID string
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `WLAN status: ENABLED or DISABLED
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `true: the SSID is third party
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_from": &schema.Schema{
										Description: `Valid From UTC timestamp
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"valid_to": &schema.Schema{
										Description: `Valid To UTC timestamp
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Description: `WLAN ID
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"wlc": &schema.Schema{
										Description: `WLC IP addres
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status of the test (RUNNING, NOTRUNNING)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"test_schedule_mode": &schema.Schema{
							Description: `Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Description: `The sensor test template version (must be 2)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wlans": &schema.Schema{
							Description: `WLANs list
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"new_template_name": &schema.Schema{
							Description: `Destination test template name
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Description: `Source test template name
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSensorTestTemplateDuplicateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSensorTestTemplateDuplicateDuplicateSensorTestTemplate(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sensors.DuplicateSensorTestTemplate(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing DuplicateSensorTestTemplate", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenSensorsDuplicateSensorTestTemplateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DuplicateSensorTestTemplate response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSensorTestTemplateDuplicateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSensorTestTemplateDuplicateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSensorTestTemplateDuplicateDuplicateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsDuplicateSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsDuplicateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_name")))) {
		request.TemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_template_name")))) {
		request.NewTemplateName = interfaceToString(v)
	}
	return &request
}

func flattenSensorsDuplicateSensorTestTemplateItem(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type_id"] = item.TypeID
	respItem["version"] = item.Version
	respItem["model_version"] = item.ModelVersion
	respItem["start_time"] = item.StartTime
	respItem["last_modified_time"] = item.LastModifiedTime
	respItem["num_associated_sensor"] = item.NumAssociatedSensor
	respItem["location"] = item.Location
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["status"] = item.Status
	respItem["connection"] = item.Connection
	respItem["action_in_progress"] = item.ActionInProgress
	respItem["frequency"] = flattenSensorsDuplicateSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = item.WLANs
	respItem["ssids"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["profiles"] = flattenSensorsDuplicateSensorTestTemplateItemProfiles(item.Profiles)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsDuplicateSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["sensors"] = flattenSensorsDuplicateSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsDuplicateSensorTestTemplateItemApCoverage(item.ApCoverage)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsDuplicateSensorTestTemplateItemFrequency(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseFrequency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["value"] = item.Value
	respItem["unit"] = item.Unit

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDs(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["ssid"] = item.SSID
		respItem["profile_name"] = item.ProfileName
		respItem["num_aps"] = item.NumAps
		respItem["num_sensors"] = item.NumSensors
		respItem["layer3web_authsecurity"] = item.Layer3WebAuthsecurity
		respItem["layer3web_authuser_name"] = item.Layer3WebAuthuserName
		respItem["layer3web_authpassword"] = item.Layer3WebAuthpassword
		respItem["layer3web_auth_email_address"] = item.Layer3WebAuthEmailAddress
		respItem["third_party"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
		respItem["id"] = item.ID
		respItem["wlan_id"] = item.WLANID
		respItem["wlc"] = item.Wlc
		respItem["valid_from"] = item.ValidFrom
		respItem["valid_to"] = item.ValidTo
		respItem["status"] = item.Status
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["auth_type"] = item.AuthType
		respItem["psk"] = item.Psk
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["password_type"] = item.PasswordType
		respItem["eap_method"] = item.EapMethod
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = item.AuthProtocol
		respItem["certfilename"] = item.Certfilename
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = item.Certpassphrase
		respItem["certdownloadurl"] = item.Certdownloadurl
		respItem["ext_web_auth_virtual_ip"] = item.ExtWebAuthVirtualIP
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = item.ExtWebAuthPortal
		respItem["ext_web_auth_access_url"] = item.ExtWebAuthAccessURL
		respItem["ext_web_auth_html_tag"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsThirdParty(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["label"] = item.Label
		respItem["tag"] = item.Tag
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsTests(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsDuplicateSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSSIDsTestsConfig(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["domains"] = item.Domains
		respItem["server"] = item.Server
		respItem["user_name"] = item.UserName
		respItem["password"] = item.Password
		respItem["url"] = item.URL
		respItem["port"] = item.Port
		respItem["protocol"] = item.Protocol
		respItem["servers"] = item.Servers
		respItem["direction"] = item.Direction
		respItem["start_port"] = item.StartPort
		respItem["end_port"] = item.EndPort
		respItem["udp_bandwidth"] = item.UDPBandwidth
		respItem["probe_type"] = item.ProbeType
		respItem["num_packets"] = item.NumPackets
		respItem["path_to_download"] = item.PathToDownload
		respItem["transfer_type"] = item.TransferType
		respItem["shared_secret"] = item.SharedSecret
		respItem["ndt_server"] = item.NdtServer
		respItem["ndt_server_port"] = item.NdtServerPort
		respItem["ndt_server_path"] = item.NdtServerPath
		respItem["uplink_test"] = boolPtrToString(item.UplinkTest)
		respItem["downlink_test"] = boolPtrToString(item.DownlinkTest)
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["user_name_prompt"] = item.UserNamePrompt
		respItem["password_prompt"] = item.PasswordPrompt
		respItem["exit_command"] = item.ExitCommand
		respItem["final_prompt"] = item.FinalPrompt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemProfiles(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseProfiles) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["auth_type"] = item.AuthType
		respItem["psk"] = item.Psk
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["password_type"] = item.PasswordType
		respItem["eap_method"] = item.EapMethod
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = item.AuthProtocol
		respItem["certfilename"] = item.Certfilename
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = item.Certpassphrase
		respItem["certdownloadurl"] = item.Certdownloadurl
		respItem["ext_web_auth_virtual_ip"] = item.ExtWebAuthVirtualIP
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = item.ExtWebAuthPortal
		respItem["ext_web_auth_access_url"] = item.ExtWebAuthAccessURL
		respItem["ext_web_auth_html_tag"] = flattenSensorsDuplicateSensorTestTemplateItemProfilesExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsDuplicateSensorTestTemplateItemProfilesTests(item.Tests)
		respItem["profile_name"] = item.ProfileName
		respItem["device_type"] = item.DeviceType
		respItem["vlan"] = item.VLAN
		respItem["location_vlan_list"] = flattenSensorsDuplicateSensorTestTemplateItemProfilesLocationVLANList(item.LocationVLANList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemProfilesExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["label"] = item.Label
		respItem["tag"] = item.Tag
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemProfilesTests(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsDuplicateSensorTestTemplateItemProfilesTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemProfilesTestsConfig(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTestsConfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["domains"] = item.Domains
		respItem["server"] = item.Server
		respItem["user_name"] = item.UserName
		respItem["password"] = item.Password
		respItem["url"] = item.URL
		respItem["port"] = item.Port
		respItem["protocol"] = item.Protocol
		respItem["servers"] = item.Servers
		respItem["direction"] = item.Direction
		respItem["start_port"] = item.StartPort
		respItem["end_port"] = item.EndPort
		respItem["udp_bandwidth"] = item.UDPBandwidth
		respItem["probe_type"] = item.ProbeType
		respItem["num_packets"] = item.NumPackets
		respItem["path_to_download"] = item.PathToDownload
		respItem["transfer_type"] = item.TransferType
		respItem["shared_secret"] = item.SharedSecret
		respItem["ndt_server"] = item.NdtServer
		respItem["ndt_server_port"] = item.NdtServerPort
		respItem["ndt_server_path"] = item.NdtServerPath
		respItem["uplink_test"] = boolPtrToString(item.UplinkTest)
		respItem["downlink_test"] = boolPtrToString(item.DownlinkTest)
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["user_name_prompt"] = item.UserNamePrompt
		respItem["password_prompt"] = item.PasswordPrompt
		respItem["exit_command"] = item.ExitCommand
		respItem["final_prompt"] = item.FinalPrompt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemProfilesLocationVLANList(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseProfilesLocationVLANList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["location_id"] = item.LocationID
		respItem["vlans"] = item.VLANs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemLocationInfoList(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["location_id"] = item.LocationID
		respItem["location_type"] = item.LocationType
		respItem["all_sensors"] = boolPtrToString(item.AllSensors)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["mac_address_list"] = item.MacAddressList
		respItem["management_vlan"] = item.ManagementVLAN
		respItem["custom_management_vlan"] = boolPtrToString(item.CustomManagementVLAN)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSensors(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSensors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["mac_address"] = item.MacAddress
		respItem["switch_mac"] = item.SwitchMac
		respItem["switch_uuid"] = item.SwitchUUID
		respItem["switch_serial_number"] = item.SwitchSerialNumber
		respItem["marked_for_uninstall"] = boolPtrToString(item.MarkedForUninstall)
		respItem["ip_address"] = item.IPAddress
		respItem["host_name"] = item.HostName
		respItem["wired_application_status"] = item.WiredApplicationStatus
		respItem["wired_application_message"] = item.WiredApplicationMessage
		respItem["assigned"] = boolPtrToString(item.Assigned)
		respItem["status"] = item.Status
		respItem["xor_sensor"] = boolPtrToString(item.XorSensor)
		respItem["target_a_ps"] = item.TargetAPs
		respItem["run_now"] = item.RunNow
		respItem["location_id"] = item.LocationID
		respItem["all_sensor_addition"] = boolPtrToString(item.AllSensorAddition)
		respItem["config_updated"] = item.ConfigUpdated
		respItem["sensor_type"] = item.SensorType
		respItem["test_mac_addresses"] = flattenSensorsDuplicateSensorTestTemplateItemSensorsTestMacAddresses(item.TestMacAddresses)
		respItem["id"] = item.ID
		respItem["service_policy"] = item.ServicePolicy
		respItem["i_perf_info"] = flattenSensorsDuplicateSensorTestTemplateItemSensorsIPerfInfo(item.IPerfInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsDuplicateSensorTestTemplateItemSensorsTestMacAddresses(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSensorsTestMacAddresses) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemSensorsIPerfInfo(item *dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseSensorsIPerfInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsDuplicateSensorTestTemplateItemApCoverage(items *[]dnacentersdkgo.ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["number_of_aps_to_test"] = item.NumberOfApsToTest
		respItem["rssi_threshold"] = item.RssiThreshold
		respItems = append(respItems, respItem)
	}
	return respItems
}
