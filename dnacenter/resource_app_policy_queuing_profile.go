package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppPolicyQueuingProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing custom application queuing profile

- Create new custom application queuing profile

- Delete existing custom application policy queuing profile by id
`,

		CreateContext: resourceAppPolicyQueuingProfileCreate,
		ReadContext:   resourceAppPolicyQueuingProfileRead,
		UpdateContext: resourceAppPolicyQueuingProfileUpdate,
		DeleteContext: resourceAppPolicyQueuingProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"clause": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interface_speed_bandwidth_clauses": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Optional: true,
															},
															"traffic_class": &schema.Schema{
																Description: `Traffic Class
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
									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
									},
									"tc_dscp_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Free test description
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Description: `Id of Queueing profile
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Description: `Queueing profile name
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

func resourceAppPolicyQueuingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationPolicyQueuingProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfile", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationPolicyQueuingProfile")
		queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams{}

		if okName {
			queryParams1.Name = vName
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetApplicationPolicyQueuingProfile", err,
				"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsApplicationPolicyGetApplicationPolicyQueuingProfile(m, response1, nil)
		item1, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetApplicationPolicyQueuingProfile response", err,
				"Failure when searching item from GetApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenApplicationPolicyGetApplicationPolicyQueuingProfileByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyQueuingProfile search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAppPolicyQueuingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams
	queryParams1.Name = vName
	item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetApplicationPolicyQueuingProfile", err,
			"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx, "parameters", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ApplicationPolicy.UpdateApplicationPolicyQueuingProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err, restyResp1.String(),
					"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err,
				"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
	}

	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName := resourceMap["name"]

	queryParams1 := dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams
	queryParams1.Name = vName
	item, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetApplicationPolicyQueuingProfile", err,
			"Failure at GetApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsApplicationPolicyGetApplicationPolicyQueuingProfile(m, getResp1, nil)
		item1, err := searchApplicationPolicyGetApplicationPolicyQueuingProfile(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationPolicyQueuingProfile(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationPolicyQueuingProfile", err, restyResp1.String(),
				"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationPolicyQueuingProfile", err,
			"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfile{}
	if v := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileItemClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
	if v := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClause(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := []dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileItemClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings {
	request := dnacentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplicationPolicyQueuingProfile(m interface{}, queryParams dnacentersdkgo.GetApplicationPolicyQueuingProfileQueryParams) (*dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicyQueuingProfile, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicyQueuingProfile
	var ite *dnacentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfile
	ite, _, err = client.ApplicationPolicy.GetApplicationPolicyQueuingProfile(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemApplicationPolicyGetApplicationPolicyQueuingProfile
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
