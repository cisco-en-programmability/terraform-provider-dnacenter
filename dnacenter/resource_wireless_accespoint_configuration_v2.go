package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessAccespointConfigurationV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- User can configure multiple access points with required options using this intent API
`,

		CreateContext: resourceWirelessAccespointConfigurationV2Create,
		ReadContext:   resourceWirelessAccespointConfigurationV2Read,
		DeleteContext: resourceWirelessAccespointConfigurationV2Delete,
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

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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
						"admin_status": &schema.Schema{
							Description: `Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"ap_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_name": &schema.Schema{
										Description: `The current host name of the access point.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ap_name_new": &schema.Schema{
										Description: `The modified hostname of the access point.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `The ethernet MAC address of the access point.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"ap_mode": &schema.Schema{
							Description: `Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"clean_air_s_i24": &schema.Schema{
							Description: `Configure clean air status for radios that are in 2.4 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"clean_air_s_i5": &schema.Schema{
							Description: `Configure clean air status for radios that are in 5 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"clean_air_s_i6": &schema.Schema{
							Description: `Configure clean air status for radios that are in 6 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_admin_status": &schema.Schema{
							Description: `To change the access point's admin status, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_ap_mode": &schema.Schema{
							Description: `To change the access point's mode, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_clean_air_s_i24_ghz": &schema.Schema{
							Description: `To change the clean air status for radios that are in 2.4 Ghz band, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_clean_air_s_i5_ghz": &schema.Schema{
							Description: `To change the clean air status for radios that are in 5 Ghz band, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_clean_air_s_i6_ghz": &schema.Schema{
							Description: `To change the clean air status for radios that are in 6 Ghz band, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_failover_priority": &schema.Schema{
							Description: `To change the access point's failover priority, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_hacontroller": &schema.Schema{
							Description: `To change the access point's HA controller, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_led_brightness_level": &schema.Schema{
							Description: `To change the access point's LED brightness level, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_led_status": &schema.Schema{
							Description: `To change the access point's LED status, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"configure_location": &schema.Schema{
							Description: `To change the access point's location, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"failover_priority": &schema.Schema{
							Description: `Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"is_assigned_site_as_location": &schema.Schema{
							Description: `To configure the access point's location as the site assigned to the access point, set this parameter's value to "true".
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"led_brightness_level": &schema.Schema{
							Description: `Configure the access point's LED brightness level by setting a value between 1 and 8.
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"led_status": &schema.Schema{
							Description: `Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"location": &schema.Schema{
							Description: `Configure the access point's location.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"primary_controller_name": &schema.Schema{
							Description: `Configure the hostname for an access point's primary controller.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"primary_ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": &schema.Schema{
										Description: `Configure the IP address for an access point's primary controller.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"radio_configurations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"admin_status": &schema.Schema{
										Description: `Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"antenna_cable_name": &schema.Schema{
										Description: `Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"antenna_gain": &schema.Schema{
										Description: `Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"antenna_pattern_name": &schema.Schema{
										Description: `Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"cable_loss": &schema.Schema{
										Description: `Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).
`,
										Type:     schema.TypeFloat,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"channel_assignment_mode": &schema.Schema{
										Description: `Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"channel_number": &schema.Schema{
										Description: `Configure the channel number on the specified radio for an access point.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"channel_width": &schema.Schema{
										Description: `Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"configure_admin_status": &schema.Schema{
										Description: `To change the admin status on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_antenna_cable": &schema.Schema{
										Description: `To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_antenna_pattern_name": &schema.Schema{
										Description: `To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_channel": &schema.Schema{
										Description: `To change the channel on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_channel_width": &schema.Schema{
										Description: `To change the channel width on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_power": &schema.Schema{
										Description: `To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"configure_radio_role_assignment": &schema.Schema{
										Description: `To change the radio role on the specified radio for an access point, set this parameter's value to "true".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"power_assignment_mode": &schema.Schema{
										Description: `Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"powerlevel": &schema.Schema{
										Description: `Configure the power level on the specified radio for an access point by setting a value between 1 and 8.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"radio_band": &schema.Schema{
										Description: `Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"radio_role_assignment": &schema.Schema{
										Description: `Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"radio_type": &schema.Schema{
										Description: `Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"secondary_controller_name": &schema.Schema{
							Description: `Configure the hostname for an access point's secondary controller.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"secondary_ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": &schema.Schema{
										Description: `Configure the IP address for an access point's secondary controller.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"tertiary_controller_name": &schema.Schema{
							Description: `Configure the hostname for an access point's tertiary controller.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"tertiary_ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": &schema.Schema{
										Description: `Configure the IP address for an access point's tertiary controller.
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
				},
			},
		},
	}
}

func resourceWirelessAccespointConfigurationV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.ConfigureAccessPointsV2(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ConfigureAccessPointsV2", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ConfigureAccessPointsV2", err1))
			return diags
		}
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenWirelessConfigureAccessPointsV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ConfigureAccessPointsV2 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceWirelessAccespointConfigurationV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessAccespointConfigurationV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2 {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_list")))) {
		request.ApList = expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2ApListArray(ctx, key+".ap_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_admin_status")))) {
		request.ConfigureAdminStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_status")))) {
		request.AdminStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_ap_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_ap_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_ap_mode")))) {
		request.ConfigureApMode = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_mode")))) {
		request.ApMode = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_failover_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_failover_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_failover_priority")))) {
		request.ConfigureFailoverPriority = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failover_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failover_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failover_priority")))) {
		request.FailoverPriority = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_led_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_led_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_led_status")))) {
		request.ConfigureLedStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".led_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".led_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".led_status")))) {
		request.LedStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_led_brightness_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_led_brightness_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_led_brightness_level")))) {
		request.ConfigureLedBrightnessLevel = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".led_brightness_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".led_brightness_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".led_brightness_level")))) {
		request.LedBrightnessLevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_location")))) {
		request.ConfigureLocation = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location")))) {
		request.Location = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_hacontroller")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_hacontroller")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_hacontroller")))) {
		request.ConfigureHAController = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_controller_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_controller_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_controller_name")))) {
		request.PrimaryControllerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_ip_address")))) {
		request.PrimaryIPAddress = expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2PrimaryIPAddress(ctx, key+".primary_ip_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_controller_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_controller_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_controller_name")))) {
		request.SecondaryControllerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_ip_address")))) {
		request.SecondaryIPAddress = expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2SecondaryIPAddress(ctx, key+".secondary_ip_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tertiary_controller_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tertiary_controller_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tertiary_controller_name")))) {
		request.TertiaryControllerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tertiary_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tertiary_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tertiary_ip_address")))) {
		request.TertiaryIPAddress = expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2TertiaryIPAddress(ctx, key+".tertiary_ip_address.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_configurations")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_configurations")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_configurations")))) {
		request.RadioConfigurations = expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2RadioConfigurationsArray(ctx, key+".radio_configurations", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_clean_air_s_i24_ghz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_clean_air_s_i24_ghz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_clean_air_s_i24_ghz")))) {
		request.ConfigureCleanAirSI24Ghz = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clean_air_s_i24")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clean_air_s_i24")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clean_air_s_i24")))) {
		request.CleanAirSI24 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_clean_air_s_i5_ghz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_clean_air_s_i5_ghz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_clean_air_s_i5_ghz")))) {
		request.ConfigureCleanAirSI5Ghz = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clean_air_s_i5")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clean_air_s_i5")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clean_air_s_i5")))) {
		request.CleanAirSI5 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_clean_air_s_i6_ghz")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_clean_air_s_i6_ghz")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_clean_air_s_i6_ghz")))) {
		request.ConfigureCleanAirSI6Ghz = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clean_air_s_i6")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clean_air_s_i6")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clean_air_s_i6")))) {
		request.CleanAirSI6 = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_assigned_site_as_location")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_assigned_site_as_location")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_assigned_site_as_location")))) {
		request.IsAssignedSiteAsLocation = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2ApListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessConfigureAccessPointsV2ApList {
	request := []dnacentersdkgo.RequestWirelessConfigureAccessPointsV2ApList{}
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
		i := expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2ApList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2ApList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2ApList {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2ApList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_name")))) {
		request.ApName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_name_new")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_name_new")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_name_new")))) {
		request.ApNameNew = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2PrimaryIPAddress(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2PrimaryIPAddress {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2PrimaryIPAddress{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2SecondaryIPAddress(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2SecondaryIPAddress {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2SecondaryIPAddress{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2TertiaryIPAddress(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2TertiaryIPAddress {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2TertiaryIPAddress{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2RadioConfigurationsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessConfigureAccessPointsV2RadioConfigurations {
	request := []dnacentersdkgo.RequestWirelessConfigureAccessPointsV2RadioConfigurations{}
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
		i := expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2RadioConfigurations(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessAccespointConfigurationV2ConfigureAccessPointsV2RadioConfigurations(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessConfigureAccessPointsV2RadioConfigurations {
	request := dnacentersdkgo.RequestWirelessConfigureAccessPointsV2RadioConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_radio_role_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_radio_role_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_radio_role_assignment")))) {
		request.ConfigureRadioRoleAssignment = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_role_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_role_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_role_assignment")))) {
		request.RadioRoleAssignment = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_band")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_band")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_band")))) {
		request.RadioBand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_admin_status")))) {
		request.ConfigureAdminStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_status")))) {
		request.AdminStatus = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_antenna_pattern_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_antenna_pattern_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_antenna_pattern_name")))) {
		request.ConfigureAntennaPatternName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".antenna_pattern_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".antenna_pattern_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".antenna_pattern_name")))) {
		request.AntennaPatternName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".antenna_gain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".antenna_gain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".antenna_gain")))) {
		request.AntennaGain = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_antenna_cable")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_antenna_cable")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_antenna_cable")))) {
		request.ConfigureAntennaCable = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".antenna_cable_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".antenna_cable_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".antenna_cable_name")))) {
		request.AntennaCableName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cable_loss")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cable_loss")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cable_loss")))) {
		request.CableLoss = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_channel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_channel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_channel")))) {
		request.ConfigureChannel = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_assignment_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_assignment_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_assignment_mode")))) {
		request.ChannelAssignmentMode = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_number")))) {
		request.ChannelNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_channel_width")))) {
		request.ConfigureChannelWidth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_width")))) {
		request.ChannelWidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".configure_power")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".configure_power")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".configure_power")))) {
		request.ConfigurePower = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_assignment_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_assignment_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_assignment_mode")))) {
		request.PowerAssignmentMode = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".powerlevel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".powerlevel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".powerlevel")))) {
		request.Powerlevel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type")))) {
		request.RadioType = interfaceToIntPtr(v)
	}
	return &request
}

func flattenWirelessConfigureAccessPointsV2Item(item *dnacentersdkgo.ResponseWirelessConfigureAccessPointsV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
