package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessRfProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Wireless.

- Create or Update RF profile

- Delete RF profile
`,

		CreateContext: resourceWirelessRfProfileCreate,
		ReadContext:   resourceWirelessRfProfileRead,
		UpdateContext: resourceWirelessRfProfileUpdate,
		DeleteContext: resourceWirelessRfProfileDelete,
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

						"channel_width": &schema.Schema{
							Description: `Channel Width
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_rf_profile": &schema.Schema{
							Description: `is Default Rf Profile
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_brown_field": &schema.Schema{
							Description: `Enable Brown Field
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_custom": &schema.Schema{
							Description: `Enable Custom
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type_a": &schema.Schema{
							Description: `Enable Radio Type A
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type_b": &schema.Schema{
							Description: `Enable Radio Type B
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_radio_type_c": &schema.Schema{
							Description: `Enable Radio Type C (6GHz)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `RF Profile Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates (Default : "6,9,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates (Default: "6,12,24")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold  ( (Default: -70)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: "AUTO")
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
										Description: `Data Rates  (Default: "9,11,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates  (Default: "12")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Min Power Level  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold   (Default: -70)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "9,11,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold (Default: "AUTO")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"radio_type_c_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates  (Default: "6,9,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates  (Default: "6,12,24")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Min Power Level  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold   (Default: -70)
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: "AUTO")
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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

						"channel_width": &schema.Schema{
							Description: `Channel Width
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_rf_profile": &schema.Schema{
							Description: `is Default Rf Profile
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_brown_field": &schema.Schema{
							Description: `Enable Brown Field
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_custom": &schema.Schema{
							Description: `Enable Custom
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type_a": &schema.Schema{
							Description: `Enable Radio Type A
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type_b": &schema.Schema{
							Description: `Enable Radio Type B
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"enable_radio_type_c": &schema.Schema{
							Description: `Enable Radio Type C (6GHz)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Description: `RF Profile Name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates (Default : "6,9,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates (Default: "6,12,24")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold  ( (Default: -70)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: "AUTO")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

									"data_rates": &schema.Schema{
										Description: `Data Rates  (Default: "9,11,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates  (Default: "12")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Min Power Level  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold   (Default: -70)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "9,11,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold (Default: "AUTO")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"radio_type_c_properties": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"data_rates": &schema.Schema{
										Description: `Data Rates  (Default: "6,9,12,18,24,36,48,54")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory Data Rates  (Default: "6,12,24")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_power_level": &schema.Schema{
										Description: `Max Power Level  (Default: 30)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"min_power_level": &schema.Schema{
										Description: `Min Power Level  (Default: -10)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"parent_profile": &schema.Schema{
										Description: `Parent Profile (Default : CUSTOM)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"power_threshold_v1": &schema.Schema{
										Description: `Power Threshold   (Default: -70)
`,
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"radio_channels": &schema.Schema{
										Description: `Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rx_sop_threshold": &schema.Schema{
										Description: `Rx Sop Threshold  (Default: "AUTO")
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"rf_profile_name": &schema.Schema{
							Description: `rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessRfProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessRfProfileCreateOrUpdateRfProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vRfProfileName := resourceItem["name"]
	vvRfProfileName := interfaceToString(vRfProfileName)

	queryParams1 := dnacentersdkgo.RetrieveRfProfilesQueryParams{}
	queryParams1.RfProfileName = vvRfProfileName
	getResponse2, err := searchWirelessRetrieveRfProfiles(m, queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvRfProfileName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessRfProfileRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Wireless.CreateOrUpdateRfProfile(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateOrUpdateRfProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateOrUpdateRfProfile", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing CreateOrUpdateRfProfile", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvRfProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessRfProfileRead(ctx, d, m)
}

func resourceWirelessRfProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vRfProfileName, okRfProfileName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RetrieveRfProfiles")
		queryParams1 := dnacentersdkgo.RetrieveRfProfilesQueryParams{}

		if okRfProfileName {
			queryParams1.RfProfileName = vRfProfileName
		}

		response1, restyResp1, _ := client.Wireless.RetrieveRfProfiles(&queryParams1)

		// if err != nil {
		// 	diags = append(diags, diagErrorWithAlt(
		// 		"Failure when executing RetrieveRfProfiles", err,
		// 		"Failure at RetrieveRfProfiles, unexpected response", ""))
		// 	return diags
		// }
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveRfProfilesItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveRfProfiles search response",
				err))
			return diags
		}
		vParameters1 := flattenWirelessRetrieveRfProfilesItem(response1)
		log.Printf("[DEBUG] parameters set sent => %v", responseInterfaceToString(vParameters1))
		if err := d.Set("parameters", vParameters1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveRfProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessRfProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	if d.HasChange("parameters") {
		// log.Printf("[DEBUG] Name used for update operation %v", queryParams1)
		request1 := expandRequestWirelessRfProfileCreateOrUpdateRfProfile(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Wireless.CreateOrUpdateRfProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing CreateOrUpdateRfProfile", err, restyResp1.String(),
					"Failure at CreateOrUpdateRfProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateOrUpdateRfProfile", err,
				"Failure at CreateOrUpdateRfProfile, unexpected response", ""))
			return diags
		}
		executionID := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionID)
		time.Sleep(5 * time.Second)
		if executionID != "" {
			response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionID)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetBusinessAPIExecutionDetails", err,
					"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
				return diags
			}
			for statusIsPending(response2.Status) {
				time.Sleep(10 * time.Second)
				response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionID)
				if err != nil || response2 == nil {
					if restyResp1 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if statusIsFailure(response2.Status) {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing CreateOrUpdateRfProfile", err))
				return diags
			}
		}
	}

	return diags
}

func resourceWirelessRfProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vRfProfileName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.RetrieveRfProfilesQueryParams{}
	queryParams1.RfProfileName = vRfProfileName
	item, err := searchWirelessRetrieveRfProfiles(m, queryParams1)
	var vvRfProfileName string
	if err != nil || item == nil {
		return diags
	}

	vvRfProfileName = queryParams1.RfProfileName
	response1, restyResp1, err := client.Wireless.DeleteRfProfiles(vvRfProfileName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteRfProfiles", err, restyResp1.String(),
				"Failure at DeleteRfProfiles, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteRfProfiles", err,
			"Failure at DeleteRfProfiles, unexpected response", ""))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteRfProfiles", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessRfProfileCreateOrUpdateRfProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfile {
	request := dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".channel_width")))) {
		request.ChannelWidth = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_custom")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_custom")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_custom")))) {
		request.EnableCustom = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_brown_field")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_brown_field")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_brown_field")))) {
		request.EnableBrownField = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_a_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_a_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_a_properties")))) {
		request.RadioTypeAProperties = expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeAProperties(ctx, key+".radio_type_a_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_b_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_b_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_b_properties")))) {
		request.RadioTypeBProperties = expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeBProperties(ctx, key+".radio_type_b_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_type_c_properties")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_type_c_properties")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_type_c_properties")))) {
		request.RadioTypeCProperties = expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeCProperties(ctx, key+".radio_type_c_properties.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_radio_type_c")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_radio_type_c")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_radio_type_c")))) {
		request.EnableRadioTypeC = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeAProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties {
	request := dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties{}
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
		request.PowerThreshold = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeBProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties {
	request := dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties{}
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
		request.PowerThreshold = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessRfProfileCreateOrUpdateRfProfileRadioTypeCProperties(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties {
	request := dnacentersdkgo.RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rx_sop_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rx_sop_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rx_sop_threshold")))) {
		request.RxSopThreshold = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_power_level")))) {
		request.MinPowerLevel = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_power_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_power_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_power_level")))) {
		request.MaxPowerLevel = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".power_threshold_v1")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".power_threshold_v1")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".power_threshold_v1")))) {
		request.PowerThreshold = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessRetrieveRfProfiles(m interface{}, queryParams dnacentersdkgo.RetrieveRfProfilesQueryParams) (*dnacentersdkgo.ResponseWirelessRetrieveRfProfiles, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessRetrieveRfProfiles
	var ite *dnacentersdkgo.ResponseWirelessRetrieveRfProfiles
	ite, _, err = client.Wireless.RetrieveRfProfiles(&queryParams)
	if err != nil {
		return foundItem, err
	}

	if ite == nil {
		return foundItem, err
	}
	if ite == nil {
		return nil, err
	}

	// Call get by _ method and set value to foundItem and return
	if ite.Name == queryParams.RfProfileName {
		var getItem *dnacentersdkgo.ResponseWirelessRetrieveRfProfiles
		getItem = ite
		foundItem = getItem
		return foundItem, err
	}

	return foundItem, err
}
