package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsApProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get AP profiles that are captured in wireless settings design.
`,

		ReadContext: dataSourceWirelessSettingsApProfilesRead,
		Schema: map[string]*schema.Schema{
			"ap_profile_name": &schema.Schema{
				Description: `apProfileName query parameter. Employ this query parameter to obtain the details of the apProfiles corresponding to the provided apProfileName.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeString,
				Optional: true,
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
										Description: `Name of the existing AP power profile to be mapped to the calendar power profile. The following API is used create AP power profile. API-/intent/api/v1/wirelessSettings/powerProfiles
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
										Description: `Indicates if rogue detection is enabled. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters
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
							Description: `Time zone of the AP.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"time_zone_offset_hour": &schema.Schema{
							Description: `Hour 'Delta from Controller' for the time zone. The value should be between -12 and 14.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"time_zone_offset_minutes": &schema.Schema{
							Description: `Minute 'Delta from Controller' for the time zone. Value should be between 0 to 59.
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

func dataSourceWirelessSettingsApProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vApProfileName, okApProfileName := d.GetOk("ap_profile_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApProfiles")
		queryParams1 := dnacentersdkgo.GetApProfilesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okApProfileName {
			queryParams1.ApProfileName = vApProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetApProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApProfiles", err,
				"Failure at GetApProfiles, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetApProfilesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApProfiles response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApProfilesItems(items *[]dnacentersdkgo.ResponseWirelessGetApProfilesResponse) []map[string]interface{} {
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
		respItem["management_setting"] = flattenWirelessGetApProfilesItemsManagementSetting(item.ManagementSetting)
		respItem["awips_enabled"] = boolPtrToString(item.AwipsEnabled)
		respItem["awips_forensic_enabled"] = boolPtrToString(item.AwipsForensicEnabled)
		respItem["rogue_detection_setting"] = flattenWirelessGetApProfilesItemsRogueDetectionSetting(item.RogueDetectionSetting)
		respItem["pmf_denial_enabled"] = boolPtrToString(item.PmfDenialEnabled)
		respItem["mesh_enabled"] = boolPtrToString(item.MeshEnabled)
		respItem["mesh_setting"] = flattenWirelessGetApProfilesItemsMeshSetting(item.MeshSetting)
		respItem["ap_power_profile_name"] = item.ApPowerProfileName
		respItem["calendar_power_profiles"] = flattenWirelessGetApProfilesItemsCalendarPowerProfiles(item.CalendarPowerProfiles)
		respItem["country_code"] = item.CountryCode
		respItem["time_zone"] = item.TimeZone
		respItem["time_zone_offset_hour"] = item.TimeZoneOffsetHour
		respItem["time_zone_offset_minutes"] = item.TimeZoneOffsetMinutes
		respItem["client_limit"] = item.ClientLimit
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetApProfilesItemsManagementSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseManagementSetting) []map[string]interface{} {
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

func flattenWirelessGetApProfilesItemsRogueDetectionSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseRogueDetectionSetting) []map[string]interface{} {
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

func flattenWirelessGetApProfilesItemsMeshSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseMeshSetting) []map[string]interface{} {
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

func flattenWirelessGetApProfilesItemsCalendarPowerProfiles(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseCalendarPowerProfiles) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["power_profile_name"] = item.PowerProfileName
	respItem["scheduler_type"] = item.SchedulerType
	respItem["duration"] = flattenWirelessGetApProfilesItemsCalendarPowerProfilesDuration(item.Duration)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApProfilesItemsCalendarPowerProfilesDuration(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseCalendarPowerProfilesDuration) []map[string]interface{} {
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
