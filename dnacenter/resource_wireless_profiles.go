package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Wireless.

- This resource allows the user to create a Wireless Network Profile

- This resource allows the user to update a Wireless Network Profile by ID

- This resource allows the user to delete Wireless Network Profile by ID
`,

		CreateContext: resourceWirelessProfilesCreate,
		ReadContext:   resourceWirelessProfilesRead,
		UpdateContext: resourceWirelessProfilesUpdate,
		DeleteContext: resourceWirelessProfilesDelete,
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

						"id": &schema.Schema{
							Description: `Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
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
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
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
									"policy_profile_name": &schema.Schema{
										Description: `Policy Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Profile Name
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

						"id": &schema.Schema{
							Description: `id path parameter. Wireless Profile Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_flex_connect": &schema.Schema{
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"interface_name": &schema.Schema{
										Description: `Interface Name. Default Value: management
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Network Profile Name
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

func resourceWirelessProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessProfilesCreateWirelessProfileConnectivity(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["wireless_profile_name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.Wireless.GetWirelessProfileByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessProfilesRead(ctx, d, m)
		}
	} else {
		queryParamImport := dnacentersdkgo.GetWirelessProfilesQueryParams{}

		response2, err := searchWirelessGetWirelessProfiles(m, queryParamImport, vvName)
		if response2 != nil && err == nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = response2.ID
			d.SetId(joinResourceID(resourceMap))
			return resourceWirelessProfilesRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.Wireless.CreateWirelessProfileConnectivity(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateWirelessProfile2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateWirelessProfile2", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateWirelessProfile2", err))
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
				"Failure when executing CreateWirelessProfile2", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetWirelessProfilesQueryParams{}
	item3, err := searchWirelessGetWirelessProfiles(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateWirelessProfile2", err,
			"Failure at CreateWirelessProfile2, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessProfilesRead(ctx, d, m)
}

func resourceWirelessProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfileByID")
		vvID := vID

		response1, restyResp1, err := client.Wireless.GetWirelessProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		// Review flatten function used
		vItem1 := flattenWirelessGetWirelessProfileByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestWirelessProfilesUpdateWirelessProfileConnectivity(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdateWirelessProfileConnectivity(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateWirelessProfile2", err, restyResp1.String(),
					"Failure at UpdateWirelessProfile2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateWirelessProfile2", err,
				"Failure at UpdateWirelessProfile2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateWirelessProfile2", err))
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
					"Failure when executing UpdateWirelessProfile2", err1))
				return diags
			}
		}

	}

	return resourceWirelessProfilesRead(ctx, d, m)
}

func resourceWirelessProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Wireless.DeleteWirelessProfileConnectivity(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteWirelessProfile2", err, restyResp1.String(),
				"Failure at DeleteWirelessProfile2, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteWirelessProfile2", err,
			"Failure at DeleteWirelessProfile2, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteWirelessProfile2", err))
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
				"Failure when executing DeleteWirelessProfile2", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessProfilesCreateWirelessProfileConnectivity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivity {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wireless_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wireless_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wireless_profile_name")))) {
		request.WirelessProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetails {
	request := []dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetails{}
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
		i := expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetails {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_name")))) {
		request.SSIDName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_profile_name")))) {
		request.WLANProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11be_profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11be_profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11be_profile_id")))) {
		request.Dot11BeProfileID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesCreateWirelessProfileConnectivitySSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessCreateWirelessProfileConnectivitySSIDDetailsFlexConnect{}
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

func expandRequestWirelessProfilesUpdateWirelessProfileConnectivity(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivity {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wireless_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wireless_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wireless_profile_name")))) {
		request.WirelessProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails {
	request := []dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails{}
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
		i := expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_name")))) {
		request.SSIDName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".flex_connect")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".flex_connect")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".flex_connect")))) {
		request.FlexConnect = expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect(ctx, key+".flex_connect.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_fabric")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_fabric")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_fabric")))) {
		request.EnableFabric = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_profile_name")))) {
		request.WLANProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot11be_profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot11be_profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot11be_profile_id")))) {
		request.Dot11BeProfileID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessProfilesUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect {
	request := dnacentersdkgo.RequestWirelessUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect{}
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

func searchWirelessGetWirelessProfiles(m interface{}, queryParams dnacentersdkgo.GetWirelessProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGetWirelessProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGetWirelessProfilesResponse

	queryParams.Offset = 1
	nResponse, _, err := client.Wireless.GetWirelessProfiles(nil)
	maxPageSize := len(*nResponse.Response)
	for len(*nResponse.Response) > 0 {
		time.Sleep(15 * time.Second)
		for _, item := range *nResponse.Response {
			if vID == item.WirelessProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		queryParams.Limit = float64(maxPageSize)
		queryParams.Offset += float64(maxPageSize)
		nResponse, _, err = client.Wireless.GetWirelessProfiles(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return nil, err

}
