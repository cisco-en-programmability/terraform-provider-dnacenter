package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessSettingsApProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Wireless.

- This resource allows the user to create a custom AP Profile.
`,

		CreateContext: resourceWirelessSettingsApProfilesCreate,
		ReadContext:   resourceWirelessSettingsApProfilesRead,
		UpdateContext: resourceWirelessSettingsApProfilesUpdate,
		DeleteContext: resourceWirelessSettingsApProfilesDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_power_profile_name": &schema.Schema{
							Description: `Name of the existing AP power profile.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_profile_name": &schema.Schema{
							Description: `Name of the Access Point profile. Max length is 32 characters.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"awips_enabled": &schema.Schema{
							Description: `Indicates if AWIPS is enabled on the AP.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"awips_forensic_enabled": &schema.Schema{
							Description: `Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"calendar_power_profiles": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scheduler_date": &schema.Schema{
													Description: `Start and End date of the duration setting, applicable for MONTHLY schedulers.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"scheduler_day": &schema.Schema{
													Description: `Applies every week on the selected days
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"scheduler_end_time": &schema.Schema{
													Description: `End time of the duration setting.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"scheduler_start_time": &schema.Schema{
													Description: `Start time of the duration setting.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"power_profile_name": &schema.Schema{
										Description: `Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scheduler_type": &schema.Schema{
										Description: `Type of the scheduler.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"client_limit": &schema.Schema{
							Description: `Number of clients. Value should be between 0-1200.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"country_code": &schema.Schema{
							Description: `Country Code`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description of the AP profile. Max length is 241 characters
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"management_setting": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_type": &schema.Schema{
										Description: `Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"cdp_state": &schema.Schema{
										Description: `Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"dot1x_password": &schema.Schema{
										Description: `Password for 802.1X authentication. AP dot1x password length should not exceed 120.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dot1x_username": &schema.Schema{
										Description: `Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"management_enable_password": &schema.Schema{
										Description: `Enable password for managing the AP. Length must be 8-120 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"management_password": &schema.Schema{
										Description: `Management password for the AP. Length must be 8-120 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"management_user_name": &schema.Schema{
										Description: `Management username must have a minimum of 1 character and a maximum of 32 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_enabled": &schema.Schema{
										Description: `Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"telnet_enabled": &schema.Schema{
										Description: `Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.
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
						"mesh_enabled": &schema.Schema{
							Description: `This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"mesh_setting": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"backhaul_client_access": &schema.Schema{
										Description: `Indicates if backhaul client access is enabled on the AP.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"bridge_group_name": &schema.Schema{
										Description: `Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ghz24_backhaul_data_rates": &schema.Schema{
										Description: `2.4GHz backhaul data rates.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ghz5_backhaul_data_rates": &schema.Schema{
										Description: `5GHz backhaul data rates.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range": &schema.Schema{
										Description: `Range of the mesh network. Value should be between 150-132000
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"rap_downlink_backhaul": &schema.Schema{
										Description: `Type of downlink backhaul used.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"pmf_denial_enabled": &schema.Schema{
							Description: `Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"remote_worker_enabled": &schema.Schema{
							Description: `Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"rogue_detection_setting": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rogue_detection": &schema.Schema{
										Description: `Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"rogue_detection_min_rssi": &schema.Schema{
										Description: `Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"rogue_detection_report_interval": &schema.Schema{
										Description: `Report interval for rogue detection. Value should be in range 10 and 300.
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"rogue_detection_transient_interval": &schema.Schema{
										Description: `Transient interval for rogue detection. Value should be 0 or from 120 to 1800.
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"time_zone": &schema.Schema{
							Description: `In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_zone_offset_hour": &schema.Schema{
							Description: `Enter the hour value (HH). The valid range is from -12 through 14.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"time_zone_offset_minutes": &schema.Schema{
							Description: `Enter the minute value (MM). The valid range is from 0 through 59.
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

func resourceWirelessSettingsApProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsApProfilesCreateApProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vApProfileName := resourceItem["ap_profile_name"]
	vvApProfileName := interfaceToString(vApProfileName)

	queryParamImport := dnacentersdkgo.GetApProfilesQueryParams{}
	queryParamImport.ApProfileName = vvApProfileName
	item2, err := searchWirelessGetApProfiles(m, queryParamImport, vvApProfileName)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["ap_profile_name"] = item2.ApProfileName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessSettingsApProfilesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Wireless.CreateApProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApProfile", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApProfile", err))
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
				"Failure when executing CreateApProfile", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetApProfilesQueryParams{}
	queryParamValidate.ApProfileName = vvApProfileName
	item3, err := searchWirelessGetApProfiles(m, queryParamValidate, vvApProfileName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApProfile", err,
			"Failure at CreateApProfile, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["ap_profile_name"] = item3.ApProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsApProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsApProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvApProfileName := resourceMap["ap_profile_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApProfiles")
		queryParams1 := dnacentersdkgo.GetApProfilesQueryParams{}
		queryParams1.ApProfileName = vvApProfileName
		response1, restyResp1, err := client.Wireless.GetApProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchWirelessGetApProfiles(m, queryParams1, vvApProfileName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenWirelessGetApProfilesByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenWirelessGetApProfilesByIDItem(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["apProfileName"] = item.ApProfileName
	respItem["description"] = item.Description
	respItem["remoteWorkerEnabled"] = item.RemoteWorkerEnabled
	respItem["managementSetting"] = flattenWirelessGetApProfilesByIDItemManagementSetting(item.ManagementSetting)
	respItem["awipsEnabled"] = item.AwipsEnabled
	respItem["awipsForensicEnabled"] = item.AwipsForensicEnabled
	respItem["rogueDetectionSetting"] = flattenWirelessGetApProfilesByIDItemRogueDetectionSetting(item.RogueDetectionSetting)
	respItem["pmfDenialEnabled"] = item.PmfDenialEnabled
	respItem["meshEnabled"] = item.MeshEnabled
	respItem["meshSetting"] = flattenWirelessGetApProfilesByIDItemMeshSetting(item.MeshSetting)
	respItem["apPowerProfileName"] = item.ApPowerProfileName
	respItem["calendarPowerProfiles"] = flattenWirelessGetApProfilesByIDItemCalendarPowerProfiles(item.CalendarPowerProfiles)
	respItem["countryCode"] = item.CountryCode
	respItem["timeZone"] = item.TimeZone
	respItem["timeZoneOffsetHour"] = item.TimeZoneOffsetHour
	respItem["timeZoneOffsetMinutes"] = item.TimeZoneOffsetMinutes
	respItem["clientLimit"] = item.ClientLimit
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApProfilesByIDItemManagementSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseManagementSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authType"] = item.AuthType
	respItem["dot1xUsername"] = item.Dot1XUsername
	respItem["dot1xPassword"] = item.Dot1XPassword
	respItem["sshEnabled"] = item.SSHEnabled
	respItem["telnetEnabled"] = item.TelnetEnabled
	respItem["managementUserName"] = item.ManagementUserName
	respItem["managementPassword"] = item.ManagementPassword
	respItem["managementEnablePassword"] = item.ManagementEnablePassword
	respItem["cdpState"] = item.CdpState
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApProfilesByIDItemRogueDetectionSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseRogueDetectionSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rogueDetection"] = item.RogueDetection
	respItem["rogueDetectionMinRssi"] = item.RogueDetectionMinRssi
	respItem["rogueDetectionTransientInterval"] = item.RogueDetectionTransientInterval
	respItem["rogueDetectionReportInterval"] = item.RogueDetectionReportInterval
	return []map[string]interface{}{
		respItem,
	}
}
func flattenWirelessGetApProfilesByIDItemMeshSetting(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseMeshSetting) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bridgeGroupName"] = item.BridgeGroupName
	respItem["backhaulClientAccess"] = item.BackhaulClientAccess
	respItem["range"] = item.Range
	respItem["ghz5BackhaulDataRates"] = item.Ghz5BackhaulDataRates
	respItem["ghz24BackhaulDataRates"] = item.Ghz24BackhaulDataRates
	respItem["rapDownlinkBackhaul"] = item.RapDownlinkBackhaul
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApProfilesByIDItemCalendarPowerProfiles(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseCalendarPowerProfiles) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["powerProfileName"] = item.PowerProfileName
	respItem["schedulerType"] = item.SchedulerType
	respItem["Duration"] = flattenWirelessGetApProfilesByIDItemCalendarPowerProfilesDuration(item.Duration)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApProfilesByIDItemCalendarPowerProfilesDuration(item *dnacentersdkgo.ResponseWirelessGetApProfilesResponseCalendarPowerProfilesDuration) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["schedulerStartTime"] = item.SchedulerStartTime
	respItem["schedulerEndTime"] = item.SchedulerEndTime
	respItem["schedulerDay"] = item.SchedulerDay
	respItem["Duration"] = item.SchedulerDate
	return []map[string]interface{}{
		respItem,
	}
}

func resourceWirelessSettingsApProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessSettingsApProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsApProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete WirelessSettingsApProfiles on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestWirelessSettingsApProfilesCreateApProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfile {
	request := dnacentersdkgo.RequestWirelessCreateApProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_profile_name")))) {
		request.ApProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_worker_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_worker_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_worker_enabled")))) {
		request.RemoteWorkerEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_setting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_setting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_setting")))) {
		request.ManagementSetting = expandRequestWirelessSettingsApProfilesCreateApProfileManagementSetting(ctx, key+".management_setting.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".awips_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".awips_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".awips_enabled")))) {
		request.AwipsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".awips_forensic_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".awips_forensic_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".awips_forensic_enabled")))) {
		request.AwipsForensicEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rogue_detection_setting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rogue_detection_setting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rogue_detection_setting")))) {
		request.RogueDetectionSetting = expandRequestWirelessSettingsApProfilesCreateApProfileRogueDetectionSetting(ctx, key+".rogue_detection_setting.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pmf_denial_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pmf_denial_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pmf_denial_enabled")))) {
		request.PmfDenialEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mesh_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mesh_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mesh_enabled")))) {
		request.MeshEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mesh_setting")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mesh_setting")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mesh_setting")))) {
		request.MeshSetting = expandRequestWirelessSettingsApProfilesCreateApProfileMeshSetting(ctx, key+".mesh_setting.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_power_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_power_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_power_profile_name")))) {
		request.ApPowerProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".calendar_power_profiles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".calendar_power_profiles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".calendar_power_profiles")))) {
		request.CalendarPowerProfiles = expandRequestWirelessSettingsApProfilesCreateApProfileCalendarPowerProfiles(ctx, key+".calendar_power_profiles.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country_code")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country_code")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country_code")))) {
		request.CountryCode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_zone")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_zone")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_zone")))) {
		request.TimeZone = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_zone_offset_hour")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_zone_offset_hour")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_zone_offset_hour")))) {
		request.TimeZoneOffsetHour = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_zone_offset_minutes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_zone_offset_minutes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_zone_offset_minutes")))) {
		request.TimeZoneOffsetMinutes = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_limit")))) {
		request.ClientLimit = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApProfilesCreateApProfileManagementSetting(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfileManagementSetting {
	request := dnacentersdkgo.RequestWirelessCreateApProfileManagementSetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot1x_username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot1x_username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot1x_username")))) {
		request.Dot1XUsername = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot1x_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot1x_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot1x_password")))) {
		request.Dot1XPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssh_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssh_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssh_enabled")))) {
		request.SSHEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".telnet_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".telnet_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".telnet_enabled")))) {
		request.TelnetEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_user_name")))) {
		request.ManagementUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_password")))) {
		request.ManagementPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_enable_password")))) {
		request.ManagementEnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cdp_state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cdp_state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cdp_state")))) {
		request.CdpState = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApProfilesCreateApProfileRogueDetectionSetting(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfileRogueDetectionSetting {
	request := dnacentersdkgo.RequestWirelessCreateApProfileRogueDetectionSetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rogue_detection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rogue_detection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rogue_detection")))) {
		request.RogueDetection = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rogue_detection_min_rssi")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rogue_detection_min_rssi")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rogue_detection_min_rssi")))) {
		request.RogueDetectionMinRssi = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rogue_detection_transient_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rogue_detection_transient_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rogue_detection_transient_interval")))) {
		request.RogueDetectionTransientInterval = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rogue_detection_report_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rogue_detection_report_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rogue_detection_report_interval")))) {
		request.RogueDetectionReportInterval = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApProfilesCreateApProfileMeshSetting(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfileMeshSetting {
	request := dnacentersdkgo.RequestWirelessCreateApProfileMeshSetting{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bridge_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bridge_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bridge_group_name")))) {
		request.BridgeGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".backhaul_client_access")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".backhaul_client_access")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".backhaul_client_access")))) {
		request.BackhaulClientAccess = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz5_backhaul_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz5_backhaul_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz5_backhaul_data_rates")))) {
		request.Ghz5BackhaulDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ghz24_backhaul_data_rates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ghz24_backhaul_data_rates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ghz24_backhaul_data_rates")))) {
		request.Ghz24BackhaulDataRates = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rap_downlink_backhaul")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rap_downlink_backhaul")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rap_downlink_backhaul")))) {
		request.RapDownlinkBackhaul = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApProfilesCreateApProfileCalendarPowerProfiles(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfileCalendarPowerProfiles {
	request := dnacentersdkgo.RequestWirelessCreateApProfileCalendarPowerProfiles{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_profile_name")))) {
		request.PowerProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_type")))) {
		request.SchedulerType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration")))) {
		request.Duration = expandRequestWirelessSettingsApProfilesCreateApProfileCalendarPowerProfilesDuration(ctx, key+".duration.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApProfilesCreateApProfileCalendarPowerProfilesDuration(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApProfileCalendarPowerProfilesDuration {
	request := dnacentersdkgo.RequestWirelessCreateApProfileCalendarPowerProfilesDuration{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_start_time")))) {
		request.SchedulerStartTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_end_time")))) {
		request.SchedulerEndTime = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_day")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_day")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_day")))) {
		request.SchedulerDay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scheduler_date")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scheduler_date")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scheduler_date")))) {
		request.SchedulerDate = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetApProfiles(m interface{}, queryParams dnacentersdkgo.GetApProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGetApProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGetApProfilesResponse
	var ite *dnacentersdkgo.ResponseWirelessGetApProfiles
	if vID != "" {
		queryParams.Offset = strconv.Itoa(1)
		nResponse, _, err := client.Wireless.GetApProfiles(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			offset, err := strconv.Atoi(queryParams.Offset)
			if err != nil {
				return foundItem, err
			}
			offset += maxPageSize
			queryParams.Limit = strconv.Itoa(maxPageSize)
			queryParams.Offset += strconv.Itoa(offset)
			nResponse, _, err = client.Wireless.GetApProfiles(&queryParams)
		}
		return nil, err
	} else if queryParams.ApProfileName != "" {
		ite, _, err = client.Wireless.GetApProfiles(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.ApProfileName == queryParams.ApProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
