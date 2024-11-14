package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSensor() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Sensors.

- Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID

- Intent API to delete an existing SENSOR test template
`,

		CreateContext: resourceSensorCreate,
		ReadContext:   resourceSensorRead,
		UpdateContext: resourceSensorUpdate,
		DeleteContext: resourceSensorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"backhaul_type": &schema.Schema{
							Description: `Backhall type: WIRED, WIRELESS
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ethernet_mac_address": &schema.Schema{
							Description: `Sensor device's ethernet MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Description: `IP Address
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_seen": &schema.Schema{
							Description: `Last seen timestamp
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"led": &schema.Schema{
							Description: `Is LED Enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `Site name in hierarchy form
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"radio_mac_address": &schema.Schema{
							Description: `Sensor device's radio MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Description: `Serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssh": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_password": &schema.Schema{
										Description: `Enable password
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssh_password": &schema.Schema{
										Description: `SSH password
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssh_state": &schema.Schema{
										Description: `SSH state
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssh_user_name": &schema.Schema{
										Description: `SSH user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Description: `Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Description: `Sensor version
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `The WIFI bands
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number of APs to test
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeInt,
										Optional: true,
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
							Computed: true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption mode
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
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
										Computed:     true,
									},
									"custom_management_vlan": &schema.Schema{
										Description: `Custom Management VLAN
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"location_type": &schema.Schema{
										Description: `Site type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_address_list": &schema.Schema{
										Description: `MAC addresses
`,
										Type:     schema.TypeList,
										Optional: true,
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
										Computed: true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site name hierarhy
`,
										Type:     schema.TypeString,
										Optional: true,
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
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor test template name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profiles": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
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
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"location_vlan_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"location_id": &schema.Schema{
													Description: `Site UUID
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Description: `Array of VLANs
`,
													Type:     schema.TypeList,
													Optional: true,
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
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
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
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
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
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
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
										Computed: true,
									},
									"vlan": &schema.Schema{
										Description: `VLAN
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"run_now": &schema.Schema{
							Description: `Run now (YES, NO)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sensors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
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
										Computed:     true,
									},
									"assigned": &schema.Schema{
										Description: `Is assigned
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"config_updated": &schema.Schema{
										Description: `Configuration updated: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"host_name": &schema.Schema{
										Description: `Host name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"i_perf_info": &schema.Schema{
										Description: `A string-stringList iPerf information
`,
										Type:     schema.TypeList,
										Optional: true,
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
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `IP address
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"marked_for_uninstall": &schema.Schema{
										Description: `Is marked for uninstall
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"name": &schema.Schema{
										Description: `Sensor name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"run_now": &schema.Schema{
										Description: `Run now: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sensor_type": &schema.Schema{
										Description: `Sensor type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"service_policy": &schema.Schema{
										Description: `Service policy
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Sensor device status: UP, DOWN, REBOOT
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_mac": &schema.Schema{
										Description: `Switch MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_serial_number": &schema.Schema{
										Description: `Switch serial number
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"switch_uuid": &schema.Schema{
										Description: `Switch device UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"target_a_ps": &schema.Schema{
										Description: `Array of target APs
`,
										Type:     schema.TypeList,
										Optional: true,
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
										Computed: true,
									},
									"wired_application_message": &schema.Schema{
										Description: `Wired application message
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"wired_application_status": &schema.Schema{
										Description: `Wired application status
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"xor_sensor": &schema.Schema{
										Description: `Is XOR sensor
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bands": &schema.Schema{
										Description: `WIFI bands: 2.4GHz or 5GHz
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
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
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer 3 WEB Auth email address
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer 3 WEB Auth password
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer 3 WEB Auth security
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer 3 WEB Auth user name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `The SSID profile name string
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"proxy_password": &schema.Schema{
										Description: `Proxy server password
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"proxy_port": &schema.Schema{
										Description: `Proxy server port
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"proxy_server": &schema.Schema{
										Description: `Proxy server for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"proxy_user_name": &schema.Schema{
										Description: `Proxy server user name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"ssid": &schema.Schema{
										Description: `The SSID string
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
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
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
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
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
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
													Computed: true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
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
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"wlan_id": &schema.Schema{
										Description: `WLAN ID
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"wlc": &schema.Schema{
										Description: `WLC IP addres
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `The sensor test template version (must be 2)
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSensorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSensorCreateSensorTestTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	item, err := searchSensorsSensors(m, vvName)
	if err != nil || item == nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceSensorRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sensors.CreateSensorTestTemplate(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSensorTestTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSensorTestTemplate", err))
		return diags
	}
	//Falta verificar por medio de EXECUTION ID, no pude porque el esquema de respuesta esta mal.

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceSensorRead(ctx, d, m)
}

func resourceSensorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Sensors")
		//queryParams1 := dnacentersdkgo.SensorsQueryParams{}

		response1, err := searchSensorsSensors(m, vName)

		if err != nil || response1 == nil {
			/*if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}*/
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC
		items := []dnacentersdkgo.ResponseSensorsSensorsResponse{
			*response1,
		}
		vItem1 := flattenSensorsSensorsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Sensors search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSensorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SensorUpdate", err, "Update method is not supported",
		"Failure at SensorUpdate, unexpected response", ""))

	return diags
}

func resourceSensorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	item, err := searchSensorsSensors(m, vName)
	if err != nil || item == nil {
		d.SetId("")
		return diags
	}

	//selectedMethod := 1
	//var vvID string
	//var vvName string
	// REVIEW: Add getAllItems and search function to get missing params

	queryParams := dnacentersdkgo.DeleteSensorTestQueryParams{
		TemplateName: vName,
	}
	response1, restyResp1, err := client.Sensors.DeleteSensorTest(&queryParams)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteSensorTest", err, restyResp1.String(),
				"Failure at DeleteSensorTest, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteSensorTest", err,
			"Failure at DeleteSensorTest, unexpected response", ""))
		return diags
	}

	//Falta verificar por medio de EXECUTION ID, no pude porque el esquema de respuesta esta mal.

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSensorCreateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplate {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_version")))) {
		request.ModelVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection")))) {
		request.Connection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssids")))) {
		request.SSIDs = expandRequestSensorCreateSensorTestTemplateSSIDsArray(ctx, key+".ssids", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profiles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profiles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profiles")))) {
		request.Profiles = expandRequestSensorCreateSensorTestTemplateProfilesArray(ctx, key+".profiles", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".encryption_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encryption_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encryption_mode")))) {
		request.EncryptionMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_now")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_now")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_now")))) {
		request.RunNow = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_info_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_info_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_info_list")))) {
		request.LocationInfoList = expandRequestSensorCreateSensorTestTemplateLocationInfoListArray(ctx, key+".location_info_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensors")))) {
		request.Sensors = expandRequestSensorCreateSensorTestTemplateSensorsArray(ctx, key+".sensors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_coverage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_coverage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_coverage")))) {
		request.ApCoverage = expandRequestSensorCreateSensorTestTemplateApCoverageArray(ctx, key+".ap_coverage", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDs(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
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
		request.ThirdParty = expandRequestSensorCreateSensorTestTemplateSSIDsThirdParty(ctx, key+".third_party.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_id")))) {
		request.WLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlc")))) {
		request.Wlc = interfaceToString(v)
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
		request.ExtWebAuthHTMLTag = expandRequestSensorCreateSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorCreateSensorTestTemplateSSIDsTestsArray(ctx, key+".tests", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsThirdParty(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected")))) {
		request.Selected = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDsTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfigArray(ctx, key+".config", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
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
		i := expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSSIDsTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles{}
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
		i := expandRequestSensorCreateSensorTestTemplateProfiles(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfiles(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles{}
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
		request.ExtWebAuthHTMLTag = expandRequestSensorCreateSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorCreateSensorTestTemplateProfilesTestsArray(ctx, key+".tests", d)
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
		request.LocationVLANList = expandRequestSensorCreateSensorTestTemplateProfilesLocationVLANListArray(ctx, key+".location_vlan_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag{}
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
		i := expandRequestSensorCreateSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests{}
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
		i := expandRequestSensorCreateSensorTestTemplateProfilesTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesTests(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorCreateSensorTestTemplateProfilesTestsConfigArray(ctx, key+".config", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig{}
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
		i := expandRequestSensorCreateSensorTestTemplateProfilesTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig{}
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesLocationVLANListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList{}
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
		i := expandRequestSensorCreateSensorTestTemplateProfilesLocationVLANList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateProfilesLocationVLANList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlans")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlans")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlans")))) {
		request.VLANs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateLocationInfoListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList{}
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
		i := expandRequestSensorCreateSensorTestTemplateLocationInfoList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateLocationInfoList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList{}
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSensorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensors {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensors{}
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
		i := expandRequestSensorCreateSensorTestTemplateSensors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSensors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensors {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensors{}
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
		request.TestMacAddresses = expandRequestSensorCreateSensorTestTemplateSensorsTestMacAddresses(ctx, key+".test_mac_addresses.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_policy")))) {
		request.ServicePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".i_perf_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".i_perf_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".i_perf_info")))) {
		request.IPerfInfo = expandRequestSensorCreateSensorTestTemplateSensorsIPerfInfo(ctx, key+".i_perf_info.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSensorsTestMacAddresses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses {
	var request dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateSensorsIPerfInfo(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo {
	var request dnacentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateApCoverageArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := []dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
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
		i := expandRequestSensorCreateSensorTestTemplateApCoverage(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSensorCreateSensorTestTemplateApCoverage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := dnacentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_aps_to_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) {
		request.NumberOfApsToTest = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSensorsSensors(m interface{}, vName string) (*dnacentersdkgo.ResponseSensorsSensorsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSensorsSensorsResponse
	var ite *dnacentersdkgo.ResponseSensorsSensors
	ite, _, err = client.Sensors.Sensors(nil)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == vName {
			var getItem *dnacentersdkgo.ResponseSensorsSensorsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
