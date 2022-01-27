package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- Delete the Wireless Profile from DNAC whose name is provided.

- Updates the wireless Network Profile with updated details provided. All sites to be present in the network profile
should be provided.

- Creates Wireless Network Profile on DNAC and associates sites and SSIDs to it.
`,

		CreateContext: resourceWirelessProfileCreate,
		ReadContext:   resourceWirelessProfileRead,
		UpdateContext: resourceWirelessProfileUpdate,
		DeleteContext: resourceWirelessProfileDelete,
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

						"profile_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sites": &schema.Schema{
										Description: `array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_fabric": &schema.Schema{
													Description: `true if fabric is enabled else false
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"flex_connect": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_flex_connect": &schema.Schema{
																Description: `true if flex connect is enabled else false
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"local_to_vlan": &schema.Schema{
																Description: `Local To VLAN ID
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_name": &schema.Schema{
													Description: `Interface Name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `SSID Name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `SSID Type(enum: Enterprise/Guest)
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

						"profile_details": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Profile Name
`,
										Type:     schema.TypeString,
										Required: true,
									},
									"sites": &schema.Schema{
										Description: `array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"]) 
`,
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_fabric": &schema.Schema{
													Description: `true is ssid is fabric else false
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
												},
												"flex_connect": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_flex_connect": &schema.Schema{
																Description: `true if flex connect is enabled else false
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
															},
															"local_to_vlan": &schema.Schema{
																Description: `Local To Vlan
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"interface_name": &schema.Schema{
													Description: `Interface Name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": &schema.Schema{
													Description: `Ssid Name
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": &schema.Schema{
													Description: `Ssid Type(enum: Enterprise/Guest)
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"wireless_profile_name": &schema.Schema{
							Description: `wirelessProfileName path parameter. Wireless Profile Name
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

func resourceWirelessProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessProfileCreateWirelessProfile(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vvName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.profile_details"); ok {
			if _, ok := d.GetOk("parameters.0.profile_details.0"); ok {
				if v, ok := d.GetOk("parameters.0.profile_details.0.name"); ok {
					vvName = interfaceToString(v)
				}
			}
		}
	}

	queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}
	queryParams1.ProfileName = vvName
	getResponse2, err := searchWirelessGetWirelessProfile(m, queryParams1)
	if getResponse2 != nil {
		log.Printf("[DEBUG] searchWirelessGetWirelessProfile result %v", responseInterfaceToString(*getResponse2))
	}
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessProfileRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Wireless.CreateWirelessProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateWirelessProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateWirelessProfile", err))
		return diags
	}
	executionID := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionID)
	if executionID != "" {
		time.Sleep(5 * time.Second)
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing CreateWirelessProfile", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessProfileRead(ctx, d, m)
}

func resourceWirelessProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vProfileName, okProfileName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetWirelessProfile")
		queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}

		if okProfileName {
			queryParams1.ProfileName = vProfileName
		}

		response1, restyResp1, err := client.Wireless.GetWirelessProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			//diags = append(diags, diagErrorWithAlt(
			//	"Failure when executing GetWirelessProfile", err,
			//	"Failure at GetWirelessProfile, unexpected response", ""))
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenWirelessGetWirelessProfileItems(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfile search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceID := d.Id()
	//resourceMap := separateResourceID(resourceID)
	vvName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.profile_details"); ok {
			if _, ok := d.GetOk("parameters.0.profile_details.0"); ok {
				if v, ok := d.GetOk("parameters.0.profile_details.0.name"); ok {
					vvName = interfaceToString(v)
				}
			}
		}
	}
	queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}
	queryParams1.ProfileName = vvName
	item, err := searchWirelessGetWirelessProfile(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetWirelessProfile", err,
			"Failure at GetWirelessProfile, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", queryParams1)
		request1 := expandRequestWirelessProfileUpdateWirelessProfile(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.Wireless.UpdateWirelessProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateWirelessProfile", err, restyResp1.String(),
					"Failure at UpdateWirelessProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateWirelessProfile", err,
				"Failure at UpdateWirelessProfile, unexpected response", ""))
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
			for response2.Status == "IN_PROGRESS" {
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
			if response2.Status == "FAILURE" {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing UpdateWirelessProfile", err))
				return diags
			}
		}
	}

	return resourceWirelessProfileRead(ctx, d, m)
}

func resourceWirelessProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	//resourceID := d.Id()
	//resourceMap := separateResourceID(resourceID)
	vvName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.profile_details"); ok {
			if _, ok := d.GetOk("parameters.0.profile_details.0"); ok {
				if v, ok := d.GetOk("parameters.0.profile_details.0.name"); ok {
					vvName = interfaceToString(v)
				}
			}
		}
	}
	queryParams1 := dnacentersdkgo.GetWirelessProfileQueryParams{}
	queryParams1.ProfileName = vvName
	item, err := searchWirelessGetWirelessProfile(m, queryParams1)
	var vvWirelessProfileName string
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetWirelessProfile", err,
			"Failure at GetWirelessProfile, unexpected response", ""))
		return diags
	}

	vvWirelessProfileName = queryParams1.ProfileName

	response1, restyResp1, err := client.Wireless.DeleteWirelessProfile(vvWirelessProfileName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteWirelessProfile", err, restyResp1.String(),
				"Failure at DeleteWirelessProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteWirelessProfile", err,
			"Failure at DeleteWirelessProfile, unexpected response", ""))
		return diags
	}
	executionID := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionID)
	if executionID != "" {
		time.Sleep(5 * time.Second)
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteWirelessProfile", err))
			return diags
		}
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessProfileCreateWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfile {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfile{}
	request.ProfileDetails = expandRequestWirelessProfileCreateWirelessProfileProfileDetails(ctx, key+".profile_details.0", d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetails {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sites")))) {
		request.Sites = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails {
	request := []dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_flex_connect")))) {
		request.EnableFlexConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_to_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_to_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_to_vlan")))) {
		request.LocalToVLAN = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfile {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfile{}
	request.ProfileDetails = expandRequestWirelessProfileUpdateWirelessProfileProfileDetails(ctx, key+".profile_details.0", d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetails {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sites")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sites")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sites")))) {
		request.Sites = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails {
	request := []dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProfileUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_flex_connect")))) {
		request.EnableFlexConnect = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_to_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_to_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_to_vlan")))) {
		request.LocalToVLAN = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchWirelessGetWirelessProfile(m interface{}, queryParams dnacentersdkgo.GetWirelessProfileQueryParams) (*dnacentersdkgo.ResponseItemWirelessGetWirelessProfile, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemWirelessGetWirelessProfile
	var ite *dnacentersdkgo.ResponseWirelessGetWirelessProfile
	ite, _, err = client.Wireless.GetWirelessProfile(&queryParams)
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
		if item.ProfileDetails.Name == queryParams.ProfileName {
			var getItem *dnacentersdkgo.ResponseItemWirelessGetWirelessProfile
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
