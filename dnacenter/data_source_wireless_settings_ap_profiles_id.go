package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsApProfilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get a AP Profile by AP Profile ID that captured in wireless settings design
`,

		ReadContext: dataSourceWirelessSettingsApProfilesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Ap Profile ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_power_profile_name": &schema.Schema{
							Description: `Name of the existing AP power profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_profile_name": &schema.Schema{
							Description: `Name of the Access Point profile. Max length is 32 characters.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"awips_enabled": &schema.Schema{
							Description: `Indicates if AWIPS is enabled on the AP.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"awips_forensic_enabled": &schema.Schema{
							Description: `Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"calendar_power_profiles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scheduler_date": &schema.Schema{
													Description: `Start and End date of the duration setting, applicable for MONTHLY schedulers.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"scheduler_day": &schema.Schema{
													Description: `Applies every week on the selected days
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"scheduler_end_time": &schema.Schema{
													Description: `End time of the duration setting.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"scheduler_start_time": &schema.Schema{
													Description: `Start time of the duration setting.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"power_profile_name": &schema.Schema{
										Description: `Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"scheduler_type": &schema.Schema{
										Description: `Type of the scheduler.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"client_limit": &schema.Schema{
							Description: `Number of clients. Value should be between 0-1200.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"country_code": &schema.Schema{
							Description: `Country Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description of the AP profile. Max length is 241 characters
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `AP Profile unique ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_setting": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_type": &schema.Schema{
										Description: `Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"cdp_state": &schema.Schema{
										Description: `Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot1x_password": &schema.Schema{
										Description: `Password for 802.1X authentication. AP dot1x password length should not exceed 120.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot1x_username": &schema.Schema{
										Description: `Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"management_enable_password": &schema.Schema{
										Description: `Enable password for managing the AP. Length must be 8-120 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"management_password": &schema.Schema{
										Description: `Management password for the AP. Length must be 8-120 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"management_user_name": &schema.Schema{
										Description: `Management username must have a minimum of 1 character and a maximum of 32 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssh_enabled": &schema.Schema{
										Description: `Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"telnet_enabled": &schema.Schema{
										Description: `Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"mesh_enabled": &schema.Schema{
							Description: `This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mesh_setting": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"backhaul_client_access": &schema.Schema{
										Description: `Indicates if backhaul client access is enabled on the AP.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bridge_group_name": &schema.Schema{
										Description: `Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ghz24_backhaul_data_rates": &schema.Schema{
										Description: `2.4GHz backhaul data rates.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ghz5_backhaul_data_rates": &schema.Schema{
										Description: `5GHz backhaul data rates.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"range": &schema.Schema{
										Description: `Range of the mesh network. Value should be between 150-132000
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"rap_downlink_backhaul": &schema.Schema{
										Description: `Type of downlink backhaul used.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"pmf_denial_enabled": &schema.Schema{
							Description: `Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"remote_worker_enabled": &schema.Schema{
							Description: `Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rogue_detection_setting": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rogue_detection": &schema.Schema{
										Description: `Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rogue_detection_min_rssi": &schema.Schema{
										Description: `Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"rogue_detection_report_interval": &schema.Schema{
										Description: `Report interval for rogue detection. Value should be in range 10 and 300.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"rogue_detection_transient_interval": &schema.Schema{
										Description: `Transient interval for rogue detection. Value should be 0 or from 120 to 1800.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"time_zone": &schema.Schema{
							Description: `In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"time_zone_offset_hour": &schema.Schema{
							Description: `Enter the hour value (HH). The valid range is from -12 through 14.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"time_zone_offset_minutes": &schema.Schema{
							Description: `Enter the minute value (MM). The valid range is from 0 through 59.
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

func dataSourceWirelessSettingsApProfilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApProfileByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Wireless.GetApProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApProfileByID", err,
				"Failure at GetApProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetApProfileByIDItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApProfileByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApProfileByIDItems(items *[]dnacentersdkgo.ResponseWirelessGetApProfileByIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ap_profile_name"] = item.ApProfileName
		respItem["description"] = item.Description
		respItem["remote_worker_enabled"] = boolPtrToString(item.RemoteWorkerEnabled)
		respItem["management_setting"] = flattenWirelessGetApProfileByIDItemsManagementSetting(item.ManagementSetting)
		respItem["awips_enabled"] = boolPtrToString(item.AwipsEnabled)
		respItem["awips_forensic_enabled"] = boolPtrToString(item.AwipsForensicEnabled)
		respItem["rogue_detection_setting"] = flattenWirelessGetApProfileByIDItemsRogueDetectionSetting(item.RogueDetectionSetting)
		respItem["pmf_denial_enabled"] = boolPtrToString(item.PmfDenialEnabled)
		respItem["mesh_enabled"] = boolPtrToString(item.MeshEnabled)
		respItem["mesh_setting"] = flattenWirelessGetApProfileByIDItemsMeshSetting(item.MeshSetting)
		respItem["ap_power_profile_name"] = item.ApPowerProfileName
		respItem["calendar_power_profiles"] = flattenWirelessGetApProfileByIDItemsCalendarPowerProfiles(item.CalendarPowerProfiles)
		respItem["country_code"] = item.CountryCode
		respItem["time_zone"] = item.TimeZone
		respItem["time_zone_offset_hour"] = item.TimeZoneOffsetHour
		respItem["time_zone_offset_minutes"] = item.TimeZoneOffsetMinutes
		respItem["client_limit"] = item.ClientLimit
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetApProfileByIDItemsManagementSetting(item *dnacentersdkgo.ResponseWirelessGetApProfileByIDResponseManagementSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["auth_type"] = item.AuthType
	respItem["dot1x_username"] = item.Dot1XUsername
	respItem["dot1x_password"] = item.Dot1XPassword
	respItem["ssh_enabled"] = boolPtrToString(item.SSHEnabled)
	respItem["telnet_enabled"] = boolPtrToString(item.TelnetEnabled)
	respItem["management_user_name"] = item.ManagementUserName
	respItem["management_password"] = item.ManagementPassword
	respItem["management_enable_password"] = item.ManagementEnablePassword
	respItem["cdp_state"] = boolPtrToString(item.CdpState)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApProfileByIDItemsRogueDetectionSetting(item *dnacentersdkgo.ResponseWirelessGetApProfileByIDResponseRogueDetectionSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rogue_detection"] = boolPtrToString(item.RogueDetection)
	respItem["rogue_detection_min_rssi"] = item.RogueDetectionMinRssi
	respItem["rogue_detection_transient_interval"] = item.RogueDetectionTransientInterval
	respItem["rogue_detection_report_interval"] = item.RogueDetectionReportInterval

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApProfileByIDItemsMeshSetting(item *dnacentersdkgo.ResponseWirelessGetApProfileByIDResponseMeshSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bridge_group_name"] = item.BridgeGroupName
	respItem["backhaul_client_access"] = boolPtrToString(item.BackhaulClientAccess)
	respItem["range"] = item.Range
	respItem["ghz5_backhaul_data_rates"] = item.Ghz5BackhaulDataRates
	respItem["ghz24_backhaul_data_rates"] = item.Ghz24BackhaulDataRates
	respItem["rap_downlink_backhaul"] = item.RapDownlinkBackhaul

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApProfileByIDItemsCalendarPowerProfiles(item *dnacentersdkgo.ResponseWirelessGetApProfileByIDResponseCalendarPowerProfiles) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["power_profile_name"] = item.PowerProfileName
	respItem["scheduler_type"] = item.SchedulerType
	respItem["duration"] = flattenWirelessGetApProfileByIDItemsCalendarPowerProfilesDuration(item.Duration)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApProfileByIDItemsCalendarPowerProfilesDuration(item *dnacentersdkgo.ResponseWirelessGetApProfileByIDResponseCalendarPowerProfilesDuration) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["scheduler_start_time"] = item.SchedulerStartTime
	respItem["scheduler_end_time"] = item.SchedulerEndTime
	respItem["scheduler_day"] = item.SchedulerDay
	respItem["scheduler_date"] = item.SchedulerDate

	return []map[string]interface{}{
		respItem,
	}

}
