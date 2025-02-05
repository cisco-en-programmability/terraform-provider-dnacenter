package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesWirelessSettingsSSIDs() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all SSIDs (Service Set Identifier) at the given site

- This data source allows the user to get an SSID (Service Set Identifier) by ID at the given site
`,

		ReadContext: dataSourceSitesWirelessSettingsSSIDsRead,
		Schema: map[string]*schema.Schema{
			"auth_type": &schema.Schema{
				Description: `authType query parameter. Auth Type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. SSID ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"l3auth_type": &schema.Schema{
				Description: `l3authType query parameter. L3 Auth Type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Description: `ssid query parameter. SSID Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlan_type": &schema.Schema{
				Description: `wlanType query parameter. Wlan Type
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_override": &schema.Schema{
							Description: `Activate the AAA Override feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"acct_servers": &schema.Schema{
							Description: `List of Accounting server IpAddresses
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"acl_name": &schema.Schema{
							Description: `Pre-Auth Access Control List (ACL) Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_server": &schema.Schema{
							Description: `Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_servers": &schema.Schema{
							Description: `List of Authentication/Authorization server IpAddresses
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"auth_type": &schema.Schema{
							Description: `L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"basic_service_set_client_idle_timeout": &schema.Schema{
							Description: `This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"basic_service_set_max_idle_enable": &schema.Schema{
							Description: `Activate the maximum idle feature for the Basic Service Set
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"cckm_tsf_tolerance": &schema.Schema{
							Description: `Cckm TImestamp Tolerance(in milliseconds)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_exclusion_enable": &schema.Schema{
							Description: `Activate the feature that allows for the exclusion of clients
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_exclusion_timeout": &schema.Schema{
							Description: `This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_rate_limit": &schema.Schema{
							Description: `This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"coverage_hole_detection_enable": &schema.Schema{
							Description: `Activate Coverage Hole Detection feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"directed_multicast_service_enable": &schema.Schema{
							Description: `The Directed Multicast Service feature becomes operational when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"egress_qos": &schema.Schema{
							Description: `Egress QOS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"external_auth_ip_address": &schema.Schema{
							Description: `External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fast_transition_over_the_distributed_system_enable": &schema.Schema{
							Description: `Enable Fast Transition over the Distributed System when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ghz24_policy": &schema.Schema{
							Description: `2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ghz6_policy_client_steering": &schema.Schema{
							Description: `True if 6 GHz Policy Client Steering is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `SSID ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ingress_qos": &schema.Schema{
							Description: `Ingress QOS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_id": &schema.Schema{
							Description: `Site UUID from where the SSID is inherited
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_name": &schema.Schema{
							Description: `Site Name from where the SSID is inherited
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_name_hierarchy": &schema.Schema{
							Description: `Inherited Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_uui_d": &schema.Schema{
							Description: `Inherited Site UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_ap_beacon_protection_enabled": &schema.Schema{
							Description: `When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x": &schema.Schema{
							Description: `When set to true, the 802.1X authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x_plus_ft": &schema.Schema{
							Description: `When set to true, the 802.1X-Plus-FT authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x_sha256": &schema.Schema{
							Description: `When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_easy_psk": &schema.Schema{
							Description: `When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_owe": &schema.Schema{
							Description: `When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk": &schema.Schema{
							Description: `When set to true, the Pre-shared Key (PSK) authentication feature is enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk_plus_ft": &schema.Schema{
							Description: `When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk_sha256": &schema.Schema{
							Description: `The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae": &schema.Schema{
							Description: `When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_ext": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_ext_plus_ft": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_plus_ft": &schema.Schema{
							Description: `Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_suite_b1921x": &schema.Schema{
							Description: `When set to true, the SuiteB192-1x authentication key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_suite_b1x": &schema.Schema{
							Description: `When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_broadcast_ssi_d": &schema.Schema{
							Description: `When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_cckm_enabled": &schema.Schema{
							Description: `True if CCKM is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_custom_nas_id_options": &schema.Schema{
							Description: `Set to true if Custom NAS ID Options provided
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Set SSID's admin status as 'Enabled' when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_fast_lane_enabled": &schema.Schema{
							Description: `True if FastLane is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_hex": &schema.Schema{
							Description: `True if passphrase is in Hex format, else False.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_mac_filtering_enabled": &schema.Schema{
							Description: `When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_posturing_enabled": &schema.Schema{
							Description: `Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_random_mac_filter_enabled": &schema.Schema{
							Description: `Deny clients using randomized MAC addresses when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_sensor_pnp": &schema.Schema{
							Description: `True if SSID is a sensor SSID
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"l3_auth_type": &schema.Schema{
							Description: `L3 Authentication Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_frame_protection_clientprotection": &schema.Schema{
							Description: `Management Frame Protection Client
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multi_psk_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"passphrase": &schema.Schema{
										Description: `Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"passphrase_type": &schema.Schema{
										Description: `Passphrase Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"nas_options": &schema.Schema{
							Description: `Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"neighbor_list_enable": &schema.Schema{
							Description: `The Neighbor List feature is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"open_ssid": &schema.Schema{
							Description: `Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"passphrase": &schema.Schema{
							Description: `Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_profile_name": &schema.Schema{
							Description: `Policy Profile Name. If not passed, profileName value will be used to populate this parameter
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `WLAN Profile Name, if not passed autogenerated profile name will be assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protected_management_frame": &schema.Schema{
							Description: `(REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_ccmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_ccmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_gcmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activatedTrue if RSN Cipher Suite GCMP128 is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_gcmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"session_time_out": &schema.Schema{
							Description: `This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"session_time_out_enable": &schema.Schema{
							Description: `Turn on the feature that imposes a time limit on user sessions
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sleeping_client_enable": &schema.Schema{
							Description: `When set to true, this will activate the timeout settings that apply to clients in sleep mode
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sleeping_client_timeout": &schema.Schema{
							Description: `This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ssid": &schema.Schema{
							Description: `Name of the SSID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_radio_type": &schema.Schema{
							Description: `Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"web_passthrough": &schema.Schema{
							Description: `When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlan_band_select_enable": &schema.Schema{
							Description: `Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlan_type": &schema.Schema{
							Description: `Wlan Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_override": &schema.Schema{
							Description: `Activate the AAA Override feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"acct_servers": &schema.Schema{
							Description: `List of Accounting server IpAddresses
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"acl_name": &schema.Schema{
							Description: `Pre-Auth Access Control List (ACL) Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_server": &schema.Schema{
							Description: `Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_servers": &schema.Schema{
							Description: `List of Authentication/Authorization server IpAddresses
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"auth_type": &schema.Schema{
							Description: `L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"basic_service_set_client_idle_timeout": &schema.Schema{
							Description: `This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"basic_service_set_max_idle_enable": &schema.Schema{
							Description: `Activate the maximum idle feature for the Basic Service Set
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"cckm_tsf_tolerance": &schema.Schema{
							Description: `Cckm TImestamp Tolerance(in milliseconds)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_exclusion_enable": &schema.Schema{
							Description: `Activate the feature that allows for the exclusion of clients
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_exclusion_timeout": &schema.Schema{
							Description: `This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"client_rate_limit": &schema.Schema{
							Description: `This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"coverage_hole_detection_enable": &schema.Schema{
							Description: `Activate Coverage Hole Detection feature when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"directed_multicast_service_enable": &schema.Schema{
							Description: `The Directed Multicast Service feature becomes operational when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"egress_qos": &schema.Schema{
							Description: `Egress QOS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"external_auth_ip_address": &schema.Schema{
							Description: `External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fast_transition_over_the_distributed_system_enable": &schema.Schema{
							Description: `Enable Fast Transition over the Distributed System when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ghz24_policy": &schema.Schema{
							Description: `2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ghz6_policy_client_steering": &schema.Schema{
							Description: `True if 6 GHz Policy Client Steering is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `SSID ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ingress_qos": &schema.Schema{
							Description: `Ingress QOS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_id": &schema.Schema{
							Description: `Site UUID from where the SSID is inherited
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_name": &schema.Schema{
							Description: `Site Name from where the SSID is inherited
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_site_name_hierarchy": &schema.Schema{
							Description: `Inherited Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_ap_beacon_protection_enabled": &schema.Schema{
							Description: `When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x": &schema.Schema{
							Description: `When set to true, the 802.1X authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x_plus_ft": &schema.Schema{
							Description: `When set to true, the 802.1X-Plus-FT authentication key is in use
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key8021x_sha256": &schema.Schema{
							Description: `When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_easy_psk": &schema.Schema{
							Description: `When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_owe": &schema.Schema{
							Description: `When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk": &schema.Schema{
							Description: `When set to true, the Pre-shared Key (PSK) authentication feature is enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk_plus_ft": &schema.Schema{
							Description: `When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_psk_sha256": &schema.Schema{
							Description: `The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae": &schema.Schema{
							Description: `When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_ext": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_ext_plus_ft": &schema.Schema{
							Description: `When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_sae_plus_ft": &schema.Schema{
							Description: `Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_suite_b1921x": &schema.Schema{
							Description: `When set to true, the SuiteB192-1x authentication key feature is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_auth_key_suite_b1x": &schema.Schema{
							Description: `When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_broadcast_ssi_d": &schema.Schema{
							Description: `When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_cckm_enabled": &schema.Schema{
							Description: `True if CCKM is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_custom_nas_id_options": &schema.Schema{
							Description: `Set to true if Custom NAS ID Options provided
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Set SSID's admin status as 'Enabled' when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_fast_lane_enabled": &schema.Schema{
							Description: `When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_hex": &schema.Schema{
							Description: `True if passphrase is in Hex format, else False.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_mac_filtering_enabled": &schema.Schema{
							Description: `True if MAC Filtering is enabled, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_posturing_enabled": &schema.Schema{
							Description: `Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_random_mac_filter_enabled": &schema.Schema{
							Description: `Deny clients using randomized MAC addresses when set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_sensor_pnp": &schema.Schema{
							Description: `True if SSID is a sensor SSID
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"l3_auth_type": &schema.Schema{
							Description: `L3 Authentication Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_frame_protection_clientprotection": &schema.Schema{
							Description: `Management Frame Protection Client
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"multi_psk_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"passphrase": &schema.Schema{
										Description: `Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"passphrase_type": &schema.Schema{
										Description: `Passphrase Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"nas_options": &schema.Schema{
							Description: `Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"neighbor_list_enable": &schema.Schema{
							Description: `The Neighbor List feature is enabled when it is set to true
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"open_ssid": &schema.Schema{
							Description: `Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"passphrase": &schema.Schema{
							Description: `Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_profile_name": &schema.Schema{
							Description: `Policy Profile Name. If not passed, profileName value will be used to populate this parameter
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `WLAN Profile Name, if not passed autogenerated profile name will be assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protected_management_frame": &schema.Schema{
							Description: `(REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_ccmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_ccmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_gcmp128": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rsn_cipher_suite_gcmp256": &schema.Schema{
							Description: `When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"session_time_out": &schema.Schema{
							Description: `This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"session_time_out_enable": &schema.Schema{
							Description: `Turn on the feature that imposes a time limit on user sessions
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sleeping_client_enable": &schema.Schema{
							Description: `When set to true, this will activate the timeout settings that apply to clients in sleep mode
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sleeping_client_timeout": &schema.Schema{
							Description: `This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"ssid": &schema.Schema{
							Description: `Name of the SSID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_radio_type": &schema.Schema{
							Description: `Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"web_passthrough": &schema.Schema{
							Description: `When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlan_band_select_enable": &schema.Schema{
							Description: `Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlan_type": &schema.Schema{
							Description: `Wlan Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSitesWirelessSettingsSSIDsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSSID, okSSID := d.GetOk("ssid")
	vWLANType, okWLANType := d.GetOk("wlan_type")
	vAuthType, okAuthType := d.GetOk("auth_type")
	vL3AuthType, okL3AuthType := d.GetOk("l3auth_type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okSiteID, okLimit, okOffset, okSSID, okWLANType, okAuthType, okL3AuthType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okSiteID, okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSSIDBySite")
		vvSiteID := vSiteID.(string)
		queryParams1 := dnacentersdkgo.GetSSIDBySiteQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSSID {
			queryParams1.SSID = vSSID.(string)
		}
		if okWLANType {
			queryParams1.WLANType = vWLANType.(string)
		}
		if okAuthType {
			queryParams1.AuthType = vAuthType.(string)
		}
		if okL3AuthType {
			queryParams1.L3AuthType = vL3AuthType.(string)
		}

		response1, restyResp1, err := client.Wireless.GetSSIDBySite(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDBySite", err,
				"Failure at GetSSIDBySite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetSSIDBySiteItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDBySite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetSSIDByID")
		vvSiteID := vSiteID.(string)
		vvID := vID.(string)

		response2, restyResp2, err := client.Wireless.GetSSIDByID(vvSiteID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDByID", err,
				"Failure at GetSSIDByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetSSIDByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetSSIDBySiteItems(items *[]dnacentersdkgo.ResponseWirelessGetSSIDBySiteResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid"] = item.SSID
		respItem["auth_type"] = item.AuthType
		respItem["passphrase"] = item.Passphrase
		respItem["is_fast_lane_enabled"] = boolPtrToString(item.IsFastLaneEnabled)
		respItem["is_mac_filtering_enabled"] = boolPtrToString(item.IsMacFilteringEnabled)
		respItem["ssid_radio_type"] = item.SSIDRadioType
		respItem["is_broadcast_ssi_d"] = boolPtrToString(item.IsBroadcastSSID)
		respItem["fast_transition"] = item.FastTransition
		respItem["session_time_out_enable"] = boolPtrToString(item.SessionTimeOutEnable)
		respItem["session_time_out"] = item.SessionTimeOut
		respItem["client_exclusion_enable"] = boolPtrToString(item.ClientExclusionEnable)
		respItem["client_exclusion_timeout"] = item.ClientExclusionTimeout
		respItem["basic_service_set_max_idle_enable"] = boolPtrToString(item.BasicServiceSetMaxIDleEnable)
		respItem["basic_service_set_client_idle_timeout"] = item.BasicServiceSetClientIDleTimeout
		respItem["directed_multicast_service_enable"] = boolPtrToString(item.DirectedMulticastServiceEnable)
		respItem["neighbor_list_enable"] = boolPtrToString(item.NeighborListEnable)
		respItem["management_frame_protection_clientprotection"] = item.ManagementFrameProtectionClientprotection
		respItem["nas_options"] = item.NasOptions
		respItem["profile_name"] = item.ProfileName
		respItem["policy_profile_name"] = item.PolicyProfileName
		respItem["aaa_override"] = boolPtrToString(item.AAAOverride)
		respItem["coverage_hole_detection_enable"] = boolPtrToString(item.CoverageHoleDetectionEnable)
		respItem["protected_management_frame"] = item.ProtectedManagementFrame
		respItem["multi_psk_settings"] = flattenWirelessGetSSIDBySiteItemsMultipSKSettings(item.MultipSKSettings)
		respItem["client_rate_limit"] = item.ClientRateLimit
		respItem["rsn_cipher_suite_gcmp256"] = boolPtrToString(item.RsnCipherSuiteGcmp256)
		respItem["rsn_cipher_suite_ccmp256"] = boolPtrToString(item.RsnCipherSuiteCcmp256)
		respItem["rsn_cipher_suite_gcmp128"] = boolPtrToString(item.RsnCipherSuiteGcmp128)
		respItem["rsn_cipher_suite_ccmp128"] = boolPtrToString(item.RsnCipherSuiteCcmp128)
		respItem["ghz6_policy_client_steering"] = boolPtrToString(item.Ghz6PolicyClientSteering)
		respItem["is_auth_key8021x"] = boolPtrToString(item.IsAuthKey8021X)
		respItem["is_auth_key8021x_plus_ft"] = boolPtrToString(item.IsAuthKey8021XPlusFT)
		respItem["is_auth_key8021x_sha256"] = boolPtrToString(item.IsAuthKey8021XSHA256)
		respItem["is_auth_key_sae"] = boolPtrToString(item.IsAuthKeySae)
		respItem["is_auth_key_sae_plus_ft"] = boolPtrToString(item.IsAuthKeySaePlusFT)
		respItem["is_auth_key_psk"] = boolPtrToString(item.IsAuthKeyPSK)
		respItem["is_auth_key_psk_plus_ft"] = boolPtrToString(item.IsAuthKeyPSKPlusFT)
		respItem["is_auth_key_owe"] = boolPtrToString(item.IsAuthKeyOWE)
		respItem["is_auth_key_easy_psk"] = boolPtrToString(item.IsAuthKeyEasyPSK)
		respItem["is_auth_key_psk_sha256"] = boolPtrToString(item.IsAuthKeyPSKSHA256)
		respItem["open_ssid"] = item.OpenSSID
		respItem["is_custom_nas_id_options"] = boolPtrToString(item.IsCustomNasIDOptions)
		respItem["wlan_band_select_enable"] = boolPtrToString(item.WLANBandSelectEnable)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["auth_servers"] = item.AuthServers
		respItem["acct_servers"] = item.AcctServers
		respItem["egress_qos"] = item.EgressQos
		respItem["ingress_qos"] = item.IngressQos
		respItem["inherited_site_id"] = item.InheritedSiteID
		respItem["inherited_site_name"] = item.InheritedSiteName
		respItem["wlan_type"] = item.WLANType
		respItem["l3_auth_type"] = item.L3AuthType
		respItem["auth_server"] = item.AuthServer
		respItem["external_auth_ip_address"] = item.ExternalAuthIPAddress
		respItem["web_passthrough"] = boolPtrToString(item.WebPassthrough)
		respItem["sleeping_client_enable"] = boolPtrToString(item.SleepingClientEnable)
		respItem["sleeping_client_timeout"] = item.SleepingClientTimeout
		respItem["acl_name"] = item.ACLName
		respItem["is_posturing_enabled"] = boolPtrToString(item.IsPosturingEnabled)
		respItem["is_auth_key_suite_b1x"] = boolPtrToString(item.IsAuthKeySuiteB1X)
		respItem["is_auth_key_suite_b1921x"] = boolPtrToString(item.IsAuthKeySuiteB1921X)
		respItem["is_auth_key_sae_ext"] = boolPtrToString(item.IsAuthKeySaeExt)
		respItem["is_auth_key_sae_ext_plus_ft"] = boolPtrToString(item.IsAuthKeySaeExtPlusFT)
		respItem["is_ap_beacon_protection_enabled"] = boolPtrToString(item.IsApBeaconProtectionEnabled)
		respItem["ghz24_policy"] = item.Ghz24Policy
		respItem["cckm_tsf_tolerance"] = item.CckmTsfTolerance
		respItem["is_cckm_enabled"] = boolPtrToString(item.IsCckmEnabled)
		respItem["is_hex"] = boolPtrToString(item.IsHex)
		respItem["is_sensor_pnp"] = boolPtrToString(item.IsSensorPnp)
		respItem["id"] = item.ID
		respItem["is_random_mac_filter_enabled"] = boolPtrToString(item.IsRandomMacFilterEnabled)
		respItem["fast_transition_over_the_distributed_system_enable"] = boolPtrToString(item.FastTransitionOverTheDistributedSystemEnable)
		respItem["inherited_site_name_hierarchy"] = item.InheritedSiteNameHierarchy
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetSSIDBySiteItemsMultipSKSettings(items *[]dnacentersdkgo.ResponseWirelessGetSSIDBySiteResponseMultipSKSettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["priority"] = item.Priority
		respItem["passphrase_type"] = item.PassphraseType
		respItem["passphrase"] = item.Passphrase
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetSSIDByIDItem(item *dnacentersdkgo.ResponseWirelessGetSSIDByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ssid"] = item.SSID
	respItem["auth_type"] = item.AuthType
	respItem["passphrase"] = item.Passphrase
	respItem["is_fast_lane_enabled"] = boolPtrToString(item.IsFastLaneEnabled)
	respItem["is_mac_filtering_enabled"] = boolPtrToString(item.IsMacFilteringEnabled)
	respItem["ssid_radio_type"] = item.SSIDRadioType
	respItem["is_broadcast_ssi_d"] = boolPtrToString(item.IsBroadcastSSID)
	respItem["fast_transition"] = item.FastTransition
	respItem["session_time_out_enable"] = boolPtrToString(item.SessionTimeOutEnable)
	respItem["session_time_out"] = item.SessionTimeOut
	respItem["client_exclusion_enable"] = boolPtrToString(item.ClientExclusionEnable)
	respItem["client_exclusion_timeout"] = item.ClientExclusionTimeout
	respItem["basic_service_set_max_idle_enable"] = boolPtrToString(item.BasicServiceSetMaxIDleEnable)
	respItem["basic_service_set_client_idle_timeout"] = item.BasicServiceSetClientIDleTimeout
	respItem["directed_multicast_service_enable"] = boolPtrToString(item.DirectedMulticastServiceEnable)
	respItem["neighbor_list_enable"] = boolPtrToString(item.NeighborListEnable)
	respItem["management_frame_protection_clientprotection"] = item.ManagementFrameProtectionClientprotection
	respItem["nas_options"] = item.NasOptions
	respItem["profile_name"] = item.ProfileName
	respItem["policy_profile_name"] = item.PolicyProfileName
	respItem["aaa_override"] = boolPtrToString(item.AAAOverride)
	respItem["coverage_hole_detection_enable"] = boolPtrToString(item.CoverageHoleDetectionEnable)
	respItem["protected_management_frame"] = item.ProtectedManagementFrame
	respItem["multi_psk_settings"] = flattenWirelessGetSSIDByIDItemMultipSKSettings(item.MultipSKSettings)
	respItem["client_rate_limit"] = item.ClientRateLimit
	respItem["rsn_cipher_suite_gcmp256"] = boolPtrToString(item.RsnCipherSuiteGcmp256)
	respItem["rsn_cipher_suite_ccmp256"] = boolPtrToString(item.RsnCipherSuiteCcmp256)
	respItem["rsn_cipher_suite_gcmp128"] = boolPtrToString(item.RsnCipherSuiteGcmp128)
	respItem["rsn_cipher_suite_ccmp128"] = boolPtrToString(item.RsnCipherSuiteCcmp128)
	respItem["ghz6_policy_client_steering"] = boolPtrToString(item.Ghz6PolicyClientSteering)
	respItem["is_auth_key8021x"] = boolPtrToString(item.IsAuthKey8021X)
	respItem["is_auth_key8021x_plus_ft"] = boolPtrToString(item.IsAuthKey8021XPlusFT)
	respItem["is_auth_key8021x_sha256"] = boolPtrToString(item.IsAuthKey8021XSHA256)
	respItem["is_auth_key_sae"] = boolPtrToString(item.IsAuthKeySae)
	respItem["is_auth_key_sae_plus_ft"] = boolPtrToString(item.IsAuthKeySaePlusFT)
	respItem["is_auth_key_psk"] = boolPtrToString(item.IsAuthKeyPSK)
	respItem["is_auth_key_psk_plus_ft"] = boolPtrToString(item.IsAuthKeyPSKPlusFT)
	respItem["is_auth_key_owe"] = boolPtrToString(item.IsAuthKeyOWE)
	respItem["is_auth_key_easy_psk"] = boolPtrToString(item.IsAuthKeyEasyPSK)
	respItem["is_auth_key_psk_sha256"] = boolPtrToString(item.IsAuthKeyPSKSHA256)
	respItem["open_ssid"] = item.OpenSSID
	respItem["is_custom_nas_id_options"] = boolPtrToString(item.IsCustomNasIDOptions)
	respItem["wlan_band_select_enable"] = boolPtrToString(item.WLANBandSelectEnable)
	respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
	respItem["auth_servers"] = item.AuthServers
	respItem["acct_servers"] = item.AcctServers
	respItem["egress_qos"] = item.EgressQos
	respItem["ingress_qos"] = item.IngressQos
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName
	respItem["wlan_type"] = item.WLANType
	respItem["l3_auth_type"] = item.L3AuthType
	respItem["auth_server"] = item.AuthServer
	respItem["external_auth_ip_address"] = item.ExternalAuthIPAddress
	respItem["web_passthrough"] = boolPtrToString(item.WebPassthrough)
	respItem["sleeping_client_enable"] = boolPtrToString(item.SleepingClientEnable)
	respItem["sleeping_client_timeout"] = item.SleepingClientTimeout
	respItem["acl_name"] = item.ACLName
	respItem["is_posturing_enabled"] = boolPtrToString(item.IsPosturingEnabled)
	respItem["is_auth_key_suite_b1x"] = boolPtrToString(item.IsAuthKeySuiteB1X)
	respItem["is_auth_key_suite_b1921x"] = boolPtrToString(item.IsAuthKeySuiteB1921X)
	respItem["is_auth_key_sae_ext"] = boolPtrToString(item.IsAuthKeySaeExt)
	respItem["is_auth_key_sae_ext_plus_ft"] = boolPtrToString(item.IsAuthKeySaeExtPlusFT)
	respItem["is_ap_beacon_protection_enabled"] = boolPtrToString(item.IsApBeaconProtectionEnabled)
	respItem["ghz24_policy"] = item.Ghz24Policy
	respItem["cckm_tsf_tolerance"] = item.CckmTsfTolerance
	respItem["is_cckm_enabled"] = boolPtrToString(item.IsCckmEnabled)
	respItem["is_hex"] = boolPtrToString(item.IsHex)
	respItem["is_sensor_pnp"] = boolPtrToString(item.IsSensorPnp)
	respItem["id"] = item.ID
	respItem["is_random_mac_filter_enabled"] = boolPtrToString(item.IsRandomMacFilterEnabled)
	respItem["fast_transition_over_the_distributed_system_enable"] = boolPtrToString(item.FastTransitionOverTheDistributedSystemEnable)
	respItem["inherited_site_name_hierarchy"] = item.InheritedSiteNameHierarchy
	respItem["inherited_site_uui_d"] = item.InheritedSiteUUID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetSSIDByIDItemMultipSKSettings(items *[]dnacentersdkgo.ResponseWirelessGetSSIDByIDResponseMultipSKSettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["priority"] = item.Priority
		respItem["passphrase_type"] = item.PassphraseType
		respItem["passphrase"] = item.Passphrase
		respItems = append(respItems, respItem)
	}
	return respItems
}
