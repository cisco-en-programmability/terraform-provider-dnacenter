package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsRfProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all RF Profiles

- This data source allows the user to get a RF Profile by RF Profile ID
`,

		ReadContext: dataSourceWirelessSettingsRfProfilesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. RF Profile ID
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

									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
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

									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
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
								},
							},
						},

						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

			"items": &schema.Schema{
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

									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
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

									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
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
								},
							},
						},

						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
		},
	}
}

func dataSourceWirelessSettingsRfProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRfProfiles")
		queryParams1 := dnacentersdkgo.GetRfProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetRfProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetRfProfiles", err,
				"Failure at GetRfProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetRfProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRfProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetRfProfileByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Wireless.GetRfProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetRfProfileByID", err,
				"Failure at GetRfProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetRfProfileByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRfProfileByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetRfProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGetRfProfilesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["rf_profile_name"] = item.RfProfileName
		respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
		respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
		respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
		respItem["enable_radio_type6_g_hz"] = boolPtrToString(item.EnableRadioType6GHz)
		respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
		respItem["radio_type_a_properties"] = flattenWirelessGetRfProfilesItemsRadioTypeAProperties(item.RadioTypeAProperties)
		respItem["radio_type_b_properties"] = flattenWirelessGetRfProfilesItemsRadioTypeBProperties(item.RadioTypeBProperties)
		respItem["radio_type6_g_hz_properties"] = flattenWirelessGetRfProfilesItemsRadioType6GHzProperties(item.RadioType6GHzProperties)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetRfProfilesItemsRadioTypeAProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioTypeAProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["channel_width"] = item.ChannelWidth
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesItemsRadioTypeBProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioTypeBProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesItemsRadioType6GHzProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioType6GHzProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["enable_standard_power_service"] = boolPtrToString(item.EnableStandardPowerService)
	respItem["multi_bssid_properties"] = flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidProperties(item.MultiBssidProperties)
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["min_dbs_width"] = item.MinDbsWidth
	respItem["max_dbs_width"] = item.MaxDbsWidth

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_parameters"] = flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item.Dot11AxParameters)
	respItem["dot11be_parameters"] = flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item.Dot11BeParameters)
	respItem["target_wake_time"] = boolPtrToString(item.TargetWakeTime)
	respItem["twt_broadcast_support"] = boolPtrToString(item.TwtBroadcastSupport)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item *dnacentersdkgo.ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
	respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItem(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rf_profile_name"] = item.RfProfileName
	respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
	respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
	respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
	respItem["enable_radio_type6_g_hz"] = boolPtrToString(item.EnableRadioType6GHz)
	respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
	respItem["radio_type_a_properties"] = flattenWirelessGetRfProfileByIDItemRadioTypeAProperties(item.RadioTypeAProperties)
	respItem["radio_type_b_properties"] = flattenWirelessGetRfProfileByIDItemRadioTypeBProperties(item.RadioTypeBProperties)
	respItem["radio_type6_g_hz_properties"] = flattenWirelessGetRfProfileByIDItemRadioType6GHzProperties(item.RadioType6GHzProperties)
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetRfProfileByIDItemRadioTypeAProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioTypeAProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["channel_width"] = item.ChannelWidth
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItemRadioTypeBProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioTypeBProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItemRadioType6GHzProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioType6GHzProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThreshold
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["enable_standard_power_service"] = boolPtrToString(item.EnableStandardPowerService)
	respItem["multi_bssid_properties"] = flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidProperties(item.MultiBssidProperties)
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["min_dbs_width"] = item.MinDbsWidth
	respItem["max_dbs_width"] = item.MaxDbsWidth

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidProperties(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_parameters"] = flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item.Dot11AxParameters)
	respItem["dot11be_parameters"] = flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item.Dot11BeParameters)
	respItem["target_wake_time"] = boolPtrToString(item.TargetWakeTime)
	respItem["twt_broadcast_support"] = boolPtrToString(item.TwtBroadcastSupport)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDItemRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item *dnacentersdkgo.ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
	respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)

	return []map[string]interface{}{
		respItem,
	}

}
