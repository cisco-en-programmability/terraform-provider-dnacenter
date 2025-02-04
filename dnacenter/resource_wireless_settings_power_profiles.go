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

func resourceWirelessSettingsPowerProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Wireless.

- This resource allows the user to create a custom Power Profile.
`,

		CreateContext: resourceWirelessSettingsPowerProfilesCreate,
		ReadContext:   resourceWirelessSettingsPowerProfilesRead,
		UpdateContext: resourceWirelessSettingsPowerProfilesUpdate,
		DeleteContext: resourceWirelessSettingsPowerProfilesDelete,
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

						"description": &schema.Schema{
							Description: `The description of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Unique Identifier of the power profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Description: `The Name of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_id": &schema.Schema{
										Description: `Interface Id for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"interface_type": &schema.Schema{
										Description: `Interface Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_type": &schema.Schema{
										Description: `Parameter Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"parameter_value": &schema.Schema{
										Description: `Parameter Value for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"sequence": &schema.Schema{
										Description: `The sequence of the power profile rule.
`,
										Type:     schema.TypeInt,
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

						"description": &schema.Schema{
							Description: `Description of the Power Profile. Max allowed characters is 128
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"profile_name": &schema.Schema{
							Description: `Name of the Power Profile. Max allowed characters is 128
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_id": &schema.Schema{
										Description: `Interface Id for the rule.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_type": &schema.Schema{
										Description: `Interface Type for the rule.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameter_type": &schema.Schema{
										Description: `Parameter Type for the rule.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameter_value": &schema.Schema{
										Description: `Parameter Value for the rule.
`,
										Type:     schema.TypeString,
										Optional: true,
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

func resourceWirelessSettingsPowerProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsPowerProfilesCreatePowerProfile(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vProfileName := resourceItem["profile_name"]
	vvProfileName := interfaceToString(vProfileName)

	queryParamImport := dnacentersdkgo.GetPowerProfilesQueryParams{}
	queryParamImport.ProfileName = vvProfileName
	item2, err := searchWirelessGetPowerProfiles(m, queryParamImport, vvProfileName)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["profile_name"] = item2.ProfileName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessSettingsPowerProfilesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Wireless.CreatePowerProfile(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatePowerProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatePowerProfile", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreatePowerProfile", err))
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
				"Failure when executing CreatePowerProfile", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetPowerProfilesQueryParams{}
	queryParamValidate.ProfileName = vvProfileName
	item3, err := searchWirelessGetPowerProfiles(m, queryParamValidate, vvProfileName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreatePowerProfile", err,
			"Failure at CreatePowerProfile, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["profile_name"] = item3.ProfileName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsPowerProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsPowerProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvProfileName := resourceMap["profile_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPowerProfiles")
		queryParams1 := dnacentersdkgo.GetPowerProfilesQueryParams{}
		queryParams1.ProfileName = vvProfileName
		response1, restyResp1, err := client.Wireless.GetPowerProfiles(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		item1, err := searchWirelessGetPowerProfiles(m, queryParams1, vvProfileName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenWirelessGetPowerProfilesByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPowerProfiles search response",
				err))
			return diags
		}

	}
	return diags
}
func flattenWirelessGetPowerProfilesByIDItem(item *dnacentersdkgo.ResponseWirelessGetPowerProfilesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["profileName"] = item.ProfileName
	respItem["description"] = item.Description
	respItem["rules"] = flattenWirelessGetPowerProfilesByIDItemRules(item.Rules)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetPowerProfilesByIDItemRules(items *[]dnacentersdkgo.ResponseWirelessGetPowerProfilesResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}

	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sequence"] = item.Sequence
		respItem["interfaceType"] = item.InterfaceType
		respItem["interfaceId"] = item.InterfaceID
		respItem["parameterType"] = item.ParameterType
		respItem["parameterValue"] = item.ParameterValue
		respItems = append(respItems, respItem)
	}
	return respItems
}
func resourceWirelessSettingsPowerProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessSettingsPowerProfilesRead(ctx, d, m)
}

func resourceWirelessSettingsPowerProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete WirelessSettingsPowerProfiles on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestWirelessSettingsPowerProfilesCreatePowerProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreatePowerProfile {
	request := dnacentersdkgo.RequestWirelessCreatePowerProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestWirelessSettingsPowerProfilesCreatePowerProfileRulesArray(ctx, key+".rules", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsPowerProfilesCreatePowerProfileRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessCreatePowerProfileRules {
	request := []dnacentersdkgo.RequestWirelessCreatePowerProfileRules{}
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
		i := expandRequestWirelessSettingsPowerProfilesCreatePowerProfileRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsPowerProfilesCreatePowerProfileRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreatePowerProfileRules {
	request := dnacentersdkgo.RequestWirelessCreatePowerProfileRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_type")))) {
		request.InterfaceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_type")))) {
		request.ParameterType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_value")))) {
		request.ParameterValue = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetPowerProfiles(m interface{}, queryParams dnacentersdkgo.GetPowerProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseWirelessGetPowerProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseWirelessGetPowerProfilesResponse
	var ite *dnacentersdkgo.ResponseWirelessGetPowerProfiles
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Wireless.GetPowerProfiles(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Wireless.GetPowerProfiles(&queryParams)
		}
		return nil, err
	} else if queryParams.ProfileName != "" {
		ite, _, err = client.Wireless.GetPowerProfiles(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.ProfileName == queryParams.ProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
