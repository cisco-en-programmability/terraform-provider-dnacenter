package dnacenter

import (
	"context"
	"errors"
	"log"
	"reflect"
	"time"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessSettingsRfProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- This resource allows the user to create a custom RF Profile

- This resource allows the user to delete a custom RF Profile

- This resource allows the user to update a custom RF Profile
`,

		CreateContext: resourceWirelessSettingsRfProfilesCreate,
		ReadContext:   resourceWirelessSettingsRfProfilesRead,
		UpdateContext: resourceWirelessSettingsRfProfilesUpdate,
		DeleteContext: resourceWirelessSettingsRfProfilesDelete,
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

						"default_rf_profile": &schema.Schema{
							Description: `True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_custom": &schema.Schema{
							Description: `True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type6_g_hz": &schema.Schema{
							Description: `True if 6 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type_a": &schema.Schema{
							Description: `True if 5 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type_b": &schema.Schema{
							Description: `True if 2.4 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `RF Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"radio_type6_g_hz_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"broadcast_probe_response_interval": &schema.Schema{
										Description: `Broadcast Probe Response Interval of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"discovery_frames6_g_hz": &schema.Schema{
										Description: `Discovery Frames of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_standard_power_service": &schema.Schema{
										Description: `True if Standard Power Service is enabled, else False
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_reset_count": &schema.Schema{
													Description: `Client Reset Count of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"client_utilization_threshold": &schema.Schema{
													Description: `Client Utilization Threshold of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_dbs_width": &schema.Schema{
										Description: `Maximum DBS Width (Permissible Values: 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_dbs_width": &schema.Schema{
										Description: `Minimum DBS Width ( Permissible values : 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"multi_bssid_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dot11be_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ofdma_multi_ru": &schema.Schema{
																Description: `OFDMA Multi-RU
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"target_wake_time": &schema.Schema{
													Description: `Target Wake Time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"twt_broadcast_support": &schema.Schema{
													Description: `TWT Broadcast Support
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"psc_enforcing_enabled": &schema.Schema{
										Description: `PSC Enforcing Enable for 6 GHz radio band
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"channel_width": &schema.Schema{
										Description: `Channel Width
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_aware": &schema.Schema{
													Description: `Client Aware of 5 GHz radio band
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"client_reset": &schema.Schema{
													Description: `Client Reset(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"client_select": &schema.Schema{
													Description: `Client Select(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"zero_wait_dfs_enable": &schema.Schema{
										Description: `Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name
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

						"default_rf_profile": &schema.Schema{
							Description: `True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type6_g_hz": &schema.Schema{
							Description: `True if 6 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type_a": &schema.Schema{
							Description: `True if 5 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type_b": &schema.Schema{
							Description: `True if 2.4 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. RF Profile ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"radio_type6_g_hz_properties": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"broadcast_probe_response_interval": &schema.Schema{
										Description: `Broadcast Probe Response Interval of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"discovery_frames6_g_hz": &schema.Schema{
										Description: `Discovery Frames of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enable_standard_power_service": &schema.Schema{
										Description: `True if Standard Power Service is enabled, else False
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_reset_count": &schema.Schema{
													Description: `Client Reset Count of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"client_utilization_threshold": &schema.Schema{
													Description: `Client Utilization Threshold of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_dbs_width": &schema.Schema{
										Description: `Maximum DBS Width (Permissible Values:20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_dbs_width": &schema.Schema{
										Description: `Minimum DBS Width (Permissible Values:20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"multi_bssid_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
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
												"dot11be_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"ofdma_multi_ru": &schema.Schema{
																Description: `OFDMA Multi-RU
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																Computed:     true,
															},
															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
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
												"target_wake_time": &schema.Schema{
													Description: `Target Wake Time
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"twt_broadcast_support": &schema.Schema{
													Description: `TWT Broadcast Support
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
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"psc_enforcing_enabled": &schema.Schema{
										Description: `PSC Enforcing Enable for 6 GHz radio band
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"channel_width": &schema.Schema{
										Description: `Channel Width
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_aware": &schema.Schema{
													Description: `Client Aware of 5 GHz radio band
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"client_reset": &schema.Schema{
													Description: `Client Reset(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"client_select": &schema.Schema{
													Description: `Client Select(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"zero_wait_dfs_enable": &schema.Schema{
										Description: `Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions
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
						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_rates": &schema.Schema{
										Description: `Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent profile of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessSettingsRfProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsRfProfilesCreateRfProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["rf_profile_name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.Wireless.GetRfProfileByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsRfProfilesRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.GetRfProfilesQueryParams{}

		response2, err := searchWirelessGetRfProfiles(m, queryParamImport, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response2.ID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessSettingsRfProfilesRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Wireless.CreateRfProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRfProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRfProfile", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateRfProfile", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateRfProfile", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetRfProfilesQueryParams{}
	item3, err := searchWirelessGetRfProfiles(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateRfProfile", err,
			"Failure at CreateRfProfile, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsRfProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsRfProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRfProfileByID")
		vvID := vID

		response1, restyResp1, err := client.Wireless.GetRfProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		// Review flatten function used
		vItem1 := flattenWirelessGetRfProfileByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRfProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessSettingsRfProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestWirelessSettingsRfProfilesUpdateRfProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateRfProfile(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRfProfile", err, restyResp1.String(),
					"Failure at UpdateRfProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRfProfile", err,
				"Failure at UpdateRfProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateRfProfile", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateRfProfile", err1))
				return diags
			}
		}

	}

	return resourceWirelessSettingsRfProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsRfProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Wireless.DeleteRfProfile(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteRfProfile", err, restyResp1.String(),
				"Failure at DeleteRfProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteRfProfile", err,
			"Failure at DeleteRfProfile, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteRfProfile", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteRfProfile", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessSettingsRfProfilesCreateRfProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfile {
	request := dnacentersdkgo.RequestWirelessCreateRfProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile_name")))) {
		request.RfProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_rf_profile")))) {
		request.DefaultRfProfile = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type_a")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type_a")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type_a")))) {
		request.EnableRadioTypeA = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type_b")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type_b")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type_b")))) {
		request.EnableRadioTypeB = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type6_g_hz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type6_g_hz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type6_g_hz")))) {
		request.EnableRadioType6GHz = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_a_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_a_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_a_properties")))) {
		request.RadioTypeAProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAProperties(ctx, key+".radio_type_a_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_b_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_b_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_b_properties")))) {
		request.RadioTypeBProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBProperties(ctx, key+".radio_type_b_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type6_g_hz_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type6_g_hz_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type6_g_hz_properties")))) {
		request.RadioType6GHzProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzProperties(ctx, key+".radio_type6_g_hz_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_width")))) {
		request.ChannelWidth = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preamble_puncture")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preamble_puncture")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preamble_puncture")))) {
		request.PreamblePuncture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".zero_wait_dfs_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".zero_wait_dfs_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".zero_wait_dfs_enable")))) {
		request.ZeroWaitDfsEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fra_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fra_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fra_properties")))) {
		request.FraProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesFraProperties(ctx, key+".fra_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesFraProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesFraProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesFraProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_aware")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_aware")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_aware")))) {
		request.ClientAware = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_select")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_select")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_select")))) {
		request.ClientSelect = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_reset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_reset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_reset")))) {
		request.ClientReset = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_standard_power_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_standard_power_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_standard_power_service")))) {
		request.EnableStandardPowerService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multi_bssid_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multi_bssid_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multi_bssid_properties")))) {
		request.MultiBssidProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties(ctx, key+".multi_bssid_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preamble_puncture")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preamble_puncture")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preamble_puncture")))) {
		request.PreamblePuncture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_dbs_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_dbs_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_dbs_width")))) {
		request.MinDbsWidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_dbs_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_dbs_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_dbs_width")))) {
		request.MaxDbsWidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psc_enforcing_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psc_enforcing_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psc_enforcing_enabled")))) {
		request.PscEnforcingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_frames6_g_hz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_frames6_g_hz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_frames6_g_hz")))) {
		request.DiscoveryFrames6GHz = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".broadcast_probe_response_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".broadcast_probe_response_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".broadcast_probe_response_interval")))) {
		request.BroadcastProbeResponseInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fra_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fra_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fra_properties")))) {
		request.FraProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesFraProperties(ctx, key+".fra_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_parameters")))) {
		request.Dot11AxParameters = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(ctx, key+".dot11ax_parameters.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11be_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11be_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11be_parameters")))) {
		request.Dot11BeParameters = expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(ctx, key+".dot11be_parameters.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_wake_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_wake_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_wake_time")))) {
		request.TargetWakeTime = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".twt_broadcast_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".twt_broadcast_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".twt_broadcast_support")))) {
		request.TwtBroadcastSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_multi_ru")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) {
		request.OfdmaMultiRu = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesFraProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesFraProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesFraProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_reset_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_reset_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_reset_count")))) {
		request.ClientResetCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_utilization_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_utilization_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_utilization_threshold")))) {
		request.ClientUtilizationThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfile {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile_name")))) {
		request.RfProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_rf_profile")))) {
		request.DefaultRfProfile = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type_a")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type_a")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type_a")))) {
		request.EnableRadioTypeA = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type_b")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type_b")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type_b")))) {
		request.EnableRadioTypeB = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type6_g_hz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type6_g_hz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type6_g_hz")))) {
		request.EnableRadioType6GHz = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_a_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_a_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_a_properties")))) {
		request.RadioTypeAProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAProperties(ctx, key+".radio_type_a_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_b_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_b_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_b_properties")))) {
		request.RadioTypeBProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBProperties(ctx, key+".radio_type_b_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type6_g_hz_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type6_g_hz_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type6_g_hz_properties")))) {
		request.RadioType6GHzProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzProperties(ctx, key+".radio_type6_g_hz_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_width")))) {
		request.ChannelWidth = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preamble_puncture")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preamble_puncture")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preamble_puncture")))) {
		request.PreamblePuncture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".zero_wait_dfs_enable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".zero_wait_dfs_enable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".zero_wait_dfs_enable")))) {
		request.ZeroWaitDfsEnable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fra_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fra_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fra_properties")))) {
		request.FraProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesFraProperties(ctx, key+".fra_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesFraProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesFraProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesFraProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_aware")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_aware")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_aware")))) {
		request.ClientAware = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_select")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_select")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_select")))) {
		request.ClientSelect = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_reset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_reset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_reset")))) {
		request.ClientReset = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_profile")))) {
		request.ParentProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_channels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_channels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_channels")))) {
		request.RadioChannels = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_rates")))) {
		request.DataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mandatory_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mandatory_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mandatory_data_rates")))) {
		request.MandatoryDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_standard_power_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_standard_power_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_standard_power_service")))) {
		request.EnableStandardPowerService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multi_bssid_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multi_bssid_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multi_bssid_properties")))) {
		request.MultiBssidProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties(ctx, key+".multi_bssid_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".preamble_puncture")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".preamble_puncture")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".preamble_puncture")))) {
		request.PreamblePuncture = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_dbs_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_dbs_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_dbs_width")))) {
		request.MinDbsWidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_dbs_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_dbs_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_dbs_width")))) {
		request.MaxDbsWidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_rx_sop_threshold")))) {
		request.CustomRxSopThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_radio_clients")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_radio_clients")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_radio_clients")))) {
		request.MaxRadioClients = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psc_enforcing_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psc_enforcing_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psc_enforcing_enabled")))) {
		request.PscEnforcingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".discovery_frames6_g_hz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".discovery_frames6_g_hz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".discovery_frames6_g_hz")))) {
		request.DiscoveryFrames6GHz = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".broadcast_probe_response_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".broadcast_probe_response_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".broadcast_probe_response_interval")))) {
		request.BroadcastProbeResponseInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fra_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fra_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fra_properties")))) {
		request.FraProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesFraProperties(ctx, key+".fra_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".coverage_hole_detection_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".coverage_hole_detection_properties")))) {
		request.CoverageHoleDetectionProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties(ctx, key+".coverage_hole_detection_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".spatial_reuse_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".spatial_reuse_properties")))) {
		request.SpatialReuseProperties = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties(ctx, key+".spatial_reuse_properties.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_parameters")))) {
		request.Dot11AxParameters = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(ctx, key+".dot11ax_parameters.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11be_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11be_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11be_parameters")))) {
		request.Dot11BeParameters = expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(ctx, key+".dot11be_parameters.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_wake_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_wake_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_wake_time")))) {
		request.TargetWakeTime = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".twt_broadcast_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".twt_broadcast_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".twt_broadcast_support")))) {
		request.TwtBroadcastSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_down_link")))) {
		request.OfdmaDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_up_link")))) {
		request.OfdmaUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_up_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_up_link")))) {
		request.MuMimoUpLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mu_mimo_down_link")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mu_mimo_down_link")))) {
		request.MuMimoDownLink = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ofdma_multi_ru")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ofdma_multi_ru")))) {
		request.OfdmaMultiRu = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesFraProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesFraProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesFraProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_reset_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_reset_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_reset_count")))) {
		request.ClientResetCount = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_utilization_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_utilization_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_utilization_threshold")))) {
		request.ClientUtilizationThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_client_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_client_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_client_level")))) {
		request.ChdClientLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_data_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_data_rssi_threshold")))) {
		request.ChdDataRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_voice_rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_voice_rssi_threshold")))) {
		request.ChdVoiceRssiThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".chd_exception_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".chd_exception_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".chd_exception_level")))) {
		request.ChdExceptionLevel = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsRfProfilesUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties {
	request := dnacentersdkgo.RequestWirelessUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect")))) {
		request.Dot11AxNonSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_non_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_non_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxNonSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect")))) {
		request.Dot11AxSrgObssPacketDetect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_min_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_min_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMinThreshold = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11ax_srg_obss_packet_detect_max_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11ax_srg_obss_packet_detect_max_threshold")))) {
		request.Dot11AxSrgObssPacketDetectMaxThreshold = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetRfProfiles(m interface{}, queryParams dnacentersdkgo.GetRfProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGetRfProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGetRfProfilesResponse

	queryParams.Offset = 1
	nResponse, _, err := client.Wireless.GetRfProfiles(nil)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.RfProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Wireless.GetRfProfiles(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
