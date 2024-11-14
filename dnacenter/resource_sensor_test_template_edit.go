package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSensorTestTemplateEdit() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Sensors.

- Intent API to deploy, schedule, or edit and existing SENSOR test template
`,

		CreateContext: resourceSensorTestTemplateEditCreate,
		ReadContext:   resourceSensorTestTemplateEditRead,
		DeleteContext: resourceSensorTestTemplateEditDelete,
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
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
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
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
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
						"id": &schema.Schema{
							Description: `The sensor test template unique identifier, generated at test creation time
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"action_in_progress": &schema.Schema{
							Description: `Indication of inprogress action
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `The WIFI bands
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number of APs to test
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `connection type of test: WIRED, WIRELESS, BOTH
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption mode
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"frequency": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"unit": &schema.Schema{
										Description: `Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Description: `Value of the unit
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"last_modified_time": &schema.Schema{
							Description: `Last modify time
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `Location string
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensors": &schema.Schema{
										Description: `Use all sensors in the site for test
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"custom_management_vlan": &schema.Schema{
										Description: `Custom Management VLAN
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_type": &schema.Schema{
										Description: `Site type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address_list": &schema.Schema{
										Description: `MAC addresses
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"management_vlan": &schema.Schema{
										Description: `Management VLAN
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site name hierarhy
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"model_version": &schema.Schema{
							Description: `Test template object model version (must be 2)
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor test template name, which is the same as in 'templateName'
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"num_associated_sensor": &schema.Schema{
							Description: `Number of associated sensor
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"num_neighbor_apthreshold": &schema.Schema{
							Description: `Number of neighboring AP threshold
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"profiles": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_vlan_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"location_id": &schema.Schema{
													Description: `Site UUID
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Description: `Array of VLANs
`,
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
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
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"vlan": &schema.Schema{
										Description: `VLAN
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"radio_as_sensor_removed": &schema.Schema{
							Description: `Radio as sensor removed
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"rssi_threshold": &schema.Schema{
							Description: `RSSI threshold
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"run_now": &schema.Schema{
							Description: `Run now (YES, NO)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"schedule_in_days": &schema.Schema{
							Description: `Bit-wise value of scheduled test days
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sensors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensor_addition": &schema.Schema{
										Description: `Is all sensor addition
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"assigned": &schema.Schema{
										Description: `Is assigned
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"config_updated": &schema.Schema{
										Description: `Configuration updated: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"host_name": &schema.Schema{
										Description: `Host name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"i_perf_info": &schema.Schema{
										Description: `A string-stringList iPerf information
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `Sensor ID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `IP address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"marked_for_uninstall": &schema.Schema{
										Description: `Is marked for uninstall
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"name": &schema.Schema{
										Description: `Sensor name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"run_now": &schema.Schema{
										Description: `Run now: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"sensor_type": &schema.Schema{
										Description: `Sensor type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"service_policy": &schema.Schema{
										Description: `Service policy
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Sensor device status: UP, DOWN, REBOOT
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_mac": &schema.Schema{
										Description: `Switch MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_serial_number": &schema.Schema{
										Description: `Switch serial number
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_uuid": &schema.Schema{
										Description: `Switch device UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"target_a_ps": &schema.Schema{
										Description: `Array of target APs
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"test_mac_addresses": &schema.Schema{
										Description: `A string-string test MAC address
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wired_application_message": &schema.Schema{
										Description: `Wired application message
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wired_application_status": &schema.Schema{
										Description: `Wired application status
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"xor_sensor": &schema.Schema{
										Description: `Is XOR sensor
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"show_wlc_upgrade_banner": &schema.Schema{
							Description: `Show WLC upgrade banner
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site hierarchy
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"bands": &schema.Schema{
										Description: `WIFI bands: 2.4GHz or 5GHz
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Identification number
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer 3 WEB Auth email address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer 3 WEB Auth password
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer 3 WEB Auth security
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer 3 WEB Auth user name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"num_aps": &schema.Schema{
										Description: `Number of APs in the test
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"num_sensors": &schema.Schema{
										Description: `Number of Sensors in the test
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `The SSID profile name string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_password": &schema.Schema{
										Description: `Proxy server password
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_port": &schema.Schema{
										Description: `Proxy server port
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_server": &schema.Schema{
										Description: `Proxy server for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_user_name": &schema.Schema{
										Description: `Proxy server user name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ssid": &schema.Schema{
										Description: `The SSID string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `WLAN status: ENABLED or DISABLED
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `true: the SSID is third party
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"valid_from": &schema.Schema{
										Description: `Valid From UTC timestamp
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"valid_to": &schema.Schema{
										Description: `Valid To UTC timestamp
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"wlan_id": &schema.Schema{
										Description: `WLAN ID
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wlc": &schema.Schema{
										Description: `WLC IP addres
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start time
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status of the test (RUNNING, NOTRUNNING)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Description: `The test template name that is to be edited
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"test_schedule_mode": &schema.Schema{
							Description: `Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Description: `The sensor test template version (must be 2)
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"wlans": &schema.Schema{
							Description: `WLANs list
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceSensorTestTemplateEditCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSensorTestTemplateEditEditSensorTestTemplate(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sensors.EditSensorTestTemplate(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing EditSensorTestTemplate", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenSensorsEditSensorTestTemplateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting EditSensorTestTemplate response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSensorTestTemplateEditRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSensorTestTemplateEditDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_name")))) {
		request.TemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_version")))) {
		request.ModelVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_modified_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_modified_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_modified_time")))) {
		request.LastModifiedTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_associated_sensor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_associated_sensor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_associated_sensor")))) {
		request.NumAssociatedSensor = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_hierarchy")))) {
		request.SiteHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection")))) {
		request.Connection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action_in_progress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action_in_progress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action_in_progress")))) {
		request.ActionInProgress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".frequency")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".frequency")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".frequency")))) {
		request.Frequency = expandRequestSensorTestTemplateEditEditSensorTestTemplateFrequency(ctx, key+".frequency.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_neighbor_apthreshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_neighbor_apthreshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_neighbor_apthreshold")))) {
		request.NumNeighborApThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".schedule_in_days")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".schedule_in_days")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".schedule_in_days")))) {
		request.ScheduleInDays = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlans")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlans")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlans")))) {
		request.WLANs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssids")))) {
		request.SSIDs = expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsArray(ctx, key+".ssids", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profiles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profiles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profiles")))) {
		request.Profiles = expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesArray(ctx, key+".profiles", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".test_schedule_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".test_schedule_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".test_schedule_mode")))) {
		request.TestScheduleMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".show_wlc_upgrade_banner")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".show_wlc_upgrade_banner")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".show_wlc_upgrade_banner")))) {
		request.ShowWlcUpgradeBanner = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_as_sensor_removed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_as_sensor_removed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_as_sensor_removed")))) {
		request.RadioAsSensorRemoved = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".encryption_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encryption_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encryption_mode")))) {
		request.EncryptionMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_now")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_now")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_now")))) {
		request.RunNow = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_info_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_info_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_info_list")))) {
		request.LocationInfoList = expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoListArray(ctx, key+".location_info_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensors")))) {
		request.Sensors = expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsArray(ctx, key+".sensors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_coverage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_coverage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_coverage")))) {
		request.ApCoverage = expandRequestSensorTestTemplateEditEditSensorTestTemplateApCoverageArray(ctx, key+".ap_coverage", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateFrequency(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateFrequency {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateFrequency{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".unit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".unit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".unit")))) {
		request.Unit = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDs {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDs{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDs {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_aps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_aps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_aps")))) {
		request.NumAps = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_sensors")))) {
		request.NumSensors = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authsecurity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authsecurity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authsecurity")))) {
		request.Layer3WebAuthsecurity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authuser_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authuser_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authuser_name")))) {
		request.Layer3WebAuthuserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authpassword")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authpassword")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authpassword")))) {
		request.Layer3WebAuthpassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_auth_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_auth_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_auth_email_address")))) {
		request.Layer3WebAuthEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsThirdParty(ctx, key+".third_party.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_id")))) {
		request.WLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlc")))) {
		request.Wlc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_from")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_from")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_from")))) {
		request.ValidFrom = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_to")))) {
		request.ValidTo = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_type")))) {
		request.PasswordType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_method")))) {
		request.EapMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scep")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scep")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scep")))) {
		request.Scep = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protocol")))) {
		request.AuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certfilename")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certfilename")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certfilename")))) {
		request.Certfilename = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certxferprotocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certxferprotocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certxferprotocol")))) {
		request.Certxferprotocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certstatus")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certstatus")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certstatus")))) {
		request.Certstatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certpassphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certpassphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certpassphrase")))) {
		request.Certpassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certdownloadurl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certdownloadurl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certdownloadurl")))) {
		request.Certdownloadurl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_virtual_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) {
		request.ExtWebAuthVirtualIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth")))) {
		request.ExtWebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".white_list")))) {
		request.WhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) {
		request.ExtWebAuthPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_access_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) {
		request.ExtWebAuthAccessURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_html_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) {
		request.ExtWebAuthHTMLTag = expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsArray(ctx, key+".tests", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsThirdParty(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsThirdParty {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsThirdParty{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected")))) {
		request.Selected = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTests {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTests{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTests {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsConfigArray(ctx, key+".config", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTestsConfig {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTestsConfig{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSSIDsTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTestsConfig {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSSIDsTestsConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server")))) {
		request.Server = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direction")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direction")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direction")))) {
		request.Direction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_port")))) {
		request.StartPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_port")))) {
		request.EndPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".udp_bandwidth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".udp_bandwidth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".udp_bandwidth")))) {
		request.UDPBandwidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".probe_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".probe_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".probe_type")))) {
		request.ProbeType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_packets")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_packets")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_packets")))) {
		request.NumPackets = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".path_to_download")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".path_to_download")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".path_to_download")))) {
		request.PathToDownload = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transfer_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transfer_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transfer_type")))) {
		request.TransferType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server")))) {
		request.NdtServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_port")))) {
		request.NdtServerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_path")))) {
		request.NdtServerPath = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".uplink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".uplink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".uplink_test")))) {
		request.UplinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".downlink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".downlink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".downlink_test")))) {
		request.DownlinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_prompt")))) {
		request.UserNamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exit_command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exit_command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exit_command")))) {
		request.ExitCommand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".final_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".final_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".final_prompt")))) {
		request.FinalPrompt = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfiles {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfiles{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateProfiles(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfiles(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfiles {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfiles{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_type")))) {
		request.PasswordType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_method")))) {
		request.EapMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scep")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scep")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scep")))) {
		request.Scep = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protocol")))) {
		request.AuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certfilename")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certfilename")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certfilename")))) {
		request.Certfilename = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certxferprotocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certxferprotocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certxferprotocol")))) {
		request.Certxferprotocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certstatus")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certstatus")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certstatus")))) {
		request.Certstatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certpassphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certpassphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certpassphrase")))) {
		request.Certpassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certdownloadurl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certdownloadurl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certdownloadurl")))) {
		request.Certdownloadurl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_virtual_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) {
		request.ExtWebAuthVirtualIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth")))) {
		request.ExtWebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".white_list")))) {
		request.WhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) {
		request.ExtWebAuthPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_access_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) {
		request.ExtWebAuthAccessURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_html_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) {
		request.ExtWebAuthHTMLTag = expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsArray(ctx, key+".tests", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan")))) {
		request.VLAN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_vlan_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_vlan_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_vlan_list")))) {
		request.LocationVLANList = expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesLocationVLANListArray(ctx, key+".location_vlan_list", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTests {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTests{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTests {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsConfigArray(ctx, key+".config", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTestsConfig {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTestsConfig{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTestsConfig {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesTestsConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server")))) {
		request.Server = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direction")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direction")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direction")))) {
		request.Direction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_port")))) {
		request.StartPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_port")))) {
		request.EndPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".udp_bandwidth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".udp_bandwidth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".udp_bandwidth")))) {
		request.UDPBandwidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".probe_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".probe_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".probe_type")))) {
		request.ProbeType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_packets")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_packets")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_packets")))) {
		request.NumPackets = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".path_to_download")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".path_to_download")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".path_to_download")))) {
		request.PathToDownload = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transfer_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transfer_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transfer_type")))) {
		request.TransferType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server")))) {
		request.NdtServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_port")))) {
		request.NdtServerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_path")))) {
		request.NdtServerPath = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".uplink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".uplink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".uplink_test")))) {
		request.UplinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".downlink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".downlink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".downlink_test")))) {
		request.DownlinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_prompt")))) {
		request.UserNamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exit_command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exit_command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exit_command")))) {
		request.ExitCommand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".final_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".final_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".final_prompt")))) {
		request.FinalPrompt = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesLocationVLANListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesLocationVLANList {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesLocationVLANList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesLocationVLANList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateProfilesLocationVLANList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesLocationVLANList {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateProfilesLocationVLANList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlans")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlans")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlans")))) {
		request.VLANs = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateLocationInfoList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateLocationInfoList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_type")))) {
		request.LocationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sensors")))) {
		request.AllSensors = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_hierarchy")))) {
		request.SiteHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address_list")))) {
		request.MacAddressList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_vlan")))) {
		request.ManagementVLAN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_management_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_management_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_management_vlan")))) {
		request.CustomManagementVLAN = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensors {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensors{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateSensors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSensors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensors {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_mac")))) {
		request.SwitchMac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_uuid")))) {
		request.SwitchUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_serial_number")))) {
		request.SwitchSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".marked_for_uninstall")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".marked_for_uninstall")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".marked_for_uninstall")))) {
		request.MarkedForUninstall = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wired_application_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wired_application_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wired_application_status")))) {
		request.WiredApplicationStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wired_application_message")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wired_application_message")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wired_application_message")))) {
		request.WiredApplicationMessage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned")))) {
		request.Assigned = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".xor_sensor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".xor_sensor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".xor_sensor")))) {
		request.XorSensor = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_a_ps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_a_ps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_a_ps")))) {
		request.TargetAPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_now")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_now")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_now")))) {
		request.RunNow = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sensor_addition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sensor_addition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sensor_addition")))) {
		request.AllSensorAddition = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_updated")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_updated")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_updated")))) {
		request.ConfigUpdated = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensor_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensor_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensor_type")))) {
		request.SensorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".test_mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".test_mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".test_mac_addresses")))) {
		request.TestMacAddresses = expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsTestMacAddresses(ctx, key+".test_mac_addresses.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_policy")))) {
		request.ServicePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".i_perf_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".i_perf_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".i_perf_info")))) {
		request.IPerfInfo = expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsIPerfInfo(ctx, key+".i_perf_info.0", d)
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsTestMacAddresses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses {
	var request dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateSensorsIPerfInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensorsIPerfInfo {
	var request dnacentersdkgo.RequestSensorsEditSensorTestTemplateSensorsIPerfInfo
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateApCoverageArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsEditSensorTestTemplateApCoverage {
	request := []dnacentersdkgo.RequestSensorsEditSensorTestTemplateApCoverage{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestSensorTestTemplateEditEditSensorTestTemplateApCoverage(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestTemplateEditEditSensorTestTemplateApCoverage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsEditSensorTestTemplateApCoverage {
	request := dnacentersdkgo.RequestSensorsEditSensorTestTemplateApCoverage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_aps_to_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) {
		request.NumberOfApsToTest = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToIntPtr(v)
	}
	return &request
}

func flattenSensorsEditSensorTestTemplateItem(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponse) []map[string]interface{} {
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
	respItem["frequency"] = flattenSensorsEditSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = item.WLANs
	respItem["ssids"] = flattenSensorsEditSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["profiles"] = flattenSensorsEditSensorTestTemplateItemProfiles(item.Profiles)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsEditSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["sensors"] = flattenSensorsEditSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsEditSensorTestTemplateItemApCoverage(item.ApCoverage)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsEditSensorTestTemplateItemFrequency(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseFrequency) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSSIDs(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDs) []map[string]interface{} {
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
		respItem["third_party"] = flattenSensorsEditSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
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
		respItem["ext_web_auth_html_tag"] = flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsEditSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemSSIDsThirdParty(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsEditSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSSIDsTests(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsEditSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemSSIDsTestsConfig(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemProfiles(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseProfiles) []map[string]interface{} {
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
		respItem["ext_web_auth_html_tag"] = flattenSensorsEditSensorTestTemplateItemProfilesExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsEditSensorTestTemplateItemProfilesTests(item.Tests)
		respItem["profile_name"] = item.ProfileName
		respItem["device_type"] = item.DeviceType
		respItem["vlan"] = item.VLAN
		respItem["location_vlan_list"] = flattenSensorsEditSensorTestTemplateItemProfilesLocationVLANList(item.LocationVLANList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemProfilesExtWebAuthHTMLTag(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseProfilesExtWebAuthHTMLTag) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemProfilesTests(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseProfilesTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsEditSensorTestTemplateItemProfilesTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemProfilesTestsConfig(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseProfilesTestsConfig) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemProfilesLocationVLANList(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseProfilesLocationVLANList) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemLocationInfoList(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseLocationInfoList) []map[string]interface{} {
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

func flattenSensorsEditSensorTestTemplateItemSensors(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSensors) []map[string]interface{} {
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
		respItem["test_mac_addresses"] = flattenSensorsEditSensorTestTemplateItemSensorsTestMacAddresses(item.TestMacAddresses)
		respItem["id"] = item.ID
		respItem["service_policy"] = item.ServicePolicy
		respItem["i_perf_info"] = flattenSensorsEditSensorTestTemplateItemSensorsIPerfInfo(item.IPerfInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsEditSensorTestTemplateItemSensorsTestMacAddresses(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSensorsTestMacAddresses) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemSensorsIPerfInfo(item *dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseSensorsIPerfInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsEditSensorTestTemplateItemApCoverage(items *[]dnacentersdkgo.ResponseSensorsEditSensorTestTemplateResponseApCoverage) []map[string]interface{} {
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
