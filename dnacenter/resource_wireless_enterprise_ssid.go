package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessEnterpriseSSID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- Creates enterprise SSID

- Update enterprise SSID

- Deletes given enterprise SSID
`,

		CreateContext: resourceWirelessEnterpriseSSIDCreate,
		ReadContext:   resourceWirelessEnterpriseSSIDRead,
		UpdateContext: resourceWirelessEnterpriseSSIDUpdate,
		DeleteContext: resourceWirelessEnterpriseSSIDDelete,
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

						"group_uuid": &schema.Schema{
							Description: `Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_server": &schema.Schema{
										Description: `Auth Server
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_broadcast_ssi_d": &schema.Schema{
										Description: `Enable Broadcast SSID
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fast_lane": &schema.Schema{
										Description: `Enable Fast Lane
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_mac_filtering": &schema.Schema{
										Description: `Enable MAC Filtering
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fast_transition": &schema.Schema{
										Description: `Fast Transition
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_enabled": &schema.Schema{
										Description: `Is Enabled
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_fabric": &schema.Schema{
										Description: `Is Fabric
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"passphrase": &schema.Schema{
										Description: `Passphrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_policy": &schema.Schema{
										Description: `Radio Policy
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"security_level": &schema.Schema{
										Description: `Security Level
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"traffic_type": &schema.Schema{
										Description: `Traffic Type
`,
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

						"version": &schema.Schema{
							Description: `Version
`,
							Type:     schema.TypeInt,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"basic_service_set_client_idle_timeout": &schema.Schema{
							Description: `Basic Service Set Client Idle Timeout
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_exclusion_timeout": &schema.Schema{
							Description: `Client Exclusion Timeout
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"enable_basic_service_set_max_idle": &schema.Schema{
							Description: `Enable Basic Service Set Max Idle 
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_broadcast_ssi_d": &schema.Schema{
							Description: `Enable Broadcast SSID
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_client_exclusion": &schema.Schema{
							Description: `Enable Client Exclusion
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_directed_multicast_service": &schema.Schema{
							Description: `Enable Directed Multicast Service
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_fast_lane": &schema.Schema{
							Description: `Enable Fast Lane
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_mac_filtering": &schema.Schema{
							Description: `Enable MAC Filtering
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_neighbor_list": &schema.Schema{
							Description: `Enable Neighbor List
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_session_time_out": &schema.Schema{
							Description: `Enable Session Timeout
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"fast_transition": &schema.Schema{
							Description: `Fast Transition
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"mfp_client_protection": &schema.Schema{
							Description: `Management Frame Protection Client
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Enter SSID Name
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"passphrase": &schema.Schema{
							Description: `Pass Phrase (Only applicable for SSID with PERSONAL security level)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"radio_policy": &schema.Schema{
							Description: `Radio Policy. Allowed values are 'Dual band operation (2.4GHz and 5GHz)', 'Dual band operation with band select', '5GHz only', '2.4GHz only'.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_level": &schema.Schema{
							Description: `Security Level
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_time_out": &schema.Schema{
							Description: `Session Time Out
`,
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ssid_name": &schema.Schema{
							Description: `ssidName path parameter. Enter the SSID name to be deleted
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"traffic_type": &schema.Schema{
							Description: `Traffic Type
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessEnterpriseSSIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSID(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vSSIDName := resourceItem["name"]
	vvSSIDName := interfaceToString(vSSIDName)

	queryParams1 := dnacentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vvSSIDName
	getResponse2, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvSSIDName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Wireless.CreateEnterpriseSSID(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEnterpriseSSID", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEnterpriseSSID", err))
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing CreateEnterpriseSSID", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvSSIDName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
}

func resourceWirelessEnterpriseSSIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName, okSSIDName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEnterpriseSSID")
		queryParams1 := dnacentersdkgo.GetEnterpriseSSIDQueryParams{}

		if okSSIDName {
			queryParams1.SSIDName = vSSIDName
		}

		response1, restyResp1, err := client.Wireless.GetEnterpriseSSID(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			/*diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEnterpriseSSID", err,
				"Failure at GetEnterpriseSSID, unexpected response", ""))
			return diags*/
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenWirelessGetEnterpriseSSIDItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEnterpriseSSID search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessEnterpriseSSIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vSSIDName
	item, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetEnterpriseSSID", err,
			"Failure at GetEnterpriseSSID, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", queryParams1)
		request1 := expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if request1 != nil {
			request1.Name = vSSIDName
		}
		response1, restyResp1, err := client.Wireless.UpdateEnterpriseSSID(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEnterpriseSSID", err, restyResp1.String(),
					"Failure at UpdateEnterpriseSSID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEnterpriseSSID", err,
				"Failure at UpdateEnterpriseSSID, unexpected response", ""))
			return diags
		}
		executionId := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionId)
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing UpdateEnterpriseSSID", err))
			return diags
		}
	}

	return resourceWirelessEnterpriseSSIDRead(ctx, d, m)
}

func resourceWirelessEnterpriseSSIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSSIDName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.GetEnterpriseSSIDQueryParams{}
	queryParams1.SSIDName = vSSIDName
	var vvSSIDName string
	item, err := searchWirelessGetEnterpriseSSID(m, queryParams1)
	if err != nil || item == nil {
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetEnterpriseSSID", err,
		"Failure at GetEnterpriseSSID, unexpected response", ""))*/
		return diags
	}

	vvSSIDName = queryParams1.SSIDName
	response1, restyResp1, err := client.Wireless.DeleteEnterpriseSSID(vvSSIDName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteEnterpriseSSID", err, restyResp1.String(),
				"Failure at DeleteEnterpriseSSID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteEnterpriseSSID", err,
			"Failure at DeleteEnterpriseSSID, unexpected response", ""))
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteEnterpriseSSID", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessEnterpriseSSIDCreateEnterpriseSSID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateEnterpriseSSID {
	request := dnacentersdkgo.RequestWirelessCreateEnterpriseSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_level")))) {
		request.SecurityLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fast_lane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fast_lane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fast_lane")))) {
		request.EnableFastLane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_mac_filtering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_mac_filtering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_mac_filtering")))) {
		request.EnableMacFiltering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_policy")))) {
		request.RadioPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) {
		request.EnableBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_session_time_out")))) {
		request.EnableSessionTimeOut = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out")))) {
		request.SessionTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_client_exclusion")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_client_exclusion")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_client_exclusion")))) {
		request.EnableClientExclusion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) {
		request.ClientExclusionTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_basic_service_set_max_idle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) {
		request.EnableBasicServiceSetMaxIDle = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_client_idle_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) {
		request.BasicServiceSetClientIDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_directed_multicast_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) {
		request.EnableDirectedMulticastService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_neighbor_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_neighbor_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_neighbor_list")))) {
		request.EnableNeighborList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mfp_client_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mfp_client_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mfp_client_protection")))) {
		request.MfpClientProtection = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessEnterpriseSSIDUpdateEnterpriseSSID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateEnterpriseSSID {
	request := dnacentersdkgo.RequestWirelessUpdateEnterpriseSSID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_level")))) {
		request.SecurityLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".passphrase")))) {
		request.Passphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fast_lane")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fast_lane")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fast_lane")))) {
		request.EnableFastLane = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_mac_filtering")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_mac_filtering")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_mac_filtering")))) {
		request.EnableMacFiltering = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radio_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radio_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radio_policy")))) {
		request.RadioPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_broadcast_ssi_d")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_broadcast_ssi_d")))) {
		request.EnableBroadcastSSID = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fast_transition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fast_transition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fast_transition")))) {
		request.FastTransition = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_session_time_out")))) {
		request.EnableSessionTimeOut = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".session_time_out")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".session_time_out")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".session_time_out")))) {
		request.SessionTimeOut = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_client_exclusion")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_client_exclusion")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_client_exclusion")))) {
		request.EnableClientExclusion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_exclusion_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_exclusion_timeout")))) {
		request.ClientExclusionTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_basic_service_set_max_idle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_basic_service_set_max_idle")))) {
		request.EnableBasicServiceSetMaxIDle = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".basic_service_set_client_idle_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".basic_service_set_client_idle_timeout")))) {
		request.BasicServiceSetClientIDleTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_directed_multicast_service")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_directed_multicast_service")))) {
		request.EnableDirectedMulticastService = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_neighbor_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_neighbor_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_neighbor_list")))) {
		request.EnableNeighborList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mfp_client_protection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mfp_client_protection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mfp_client_protection")))) {
		request.MfpClientProtection = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchWirelessGetEnterpriseSSID(m interface{}, queryParams dnacentersdkgo.GetEnterpriseSSIDQueryParams) (*dnacentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails
	var ite *dnacentersdkgo.ResponseWirelessGetEnterpriseSSID
	ite, _, err = client.Wireless.GetEnterpriseSSID(&queryParams)
	if err != nil {
		return foundItem, err
	}

	if ite == nil {
		return foundItem, err
	}

	items := ite

	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		itemsCopy2 := *item.SSIDDetails
		for _, item := range itemsCopy2 {
			if item.Name == queryParams.SSIDName {
				var getItem *dnacentersdkgo.ResponseItemWirelessGetEnterpriseSSIDSSIDDetails
				getItem = &item
				foundItem = getItem
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
