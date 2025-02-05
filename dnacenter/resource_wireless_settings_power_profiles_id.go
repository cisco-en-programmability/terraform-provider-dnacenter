package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessSettingsPowerProfilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read, update and delete operations on Wireless.

- This resource allows the user to delete an Power Profile by specifying the Power Profile ID.

- This resource allows the user to update a custom power Profile
`,

		CreateContext: resourceWirelessSettingsPowerProfilesIDCreate,
		ReadContext:   resourceWirelessSettingsPowerProfilesIDRead,
		UpdateContext: resourceWirelessSettingsPowerProfilesIDUpdate,
		DeleteContext: resourceWirelessSettingsPowerProfilesIDDelete,
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
										Description: `Sequential Ordered List of rules for Power Profile.
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
							Description: `Description of the Power Profile. Max length is 32 characters
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Power Profile Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"profile_name": &schema.Schema{
							Description: `Name of the Power Profile. Max length is 32 characters
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
										Description: `Interface ID
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_type": &schema.Schema{
										Description: `Interface Type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameter_type": &schema.Schema{
										Description: `Parameter Type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parameter_value": &schema.Schema{
										Description: `Parameter Value
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

func resourceWirelessSettingsPowerProfilesIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceWirelessSettingsPowerProfilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPowerProfileByID")
		vvID := vID

		response1, restyResp1, err := client.Wireless.GetPowerProfileByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetPowerProfileByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPowerProfileByID response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceWirelessSettingsPowerProfilesIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByID(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Wireless.UpdatePowerProfileByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePowerProfileByID", err, restyResp1.String(),
					"Failure at UpdatePowerProfileByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePowerProfileByID", err,
				"Failure at UpdatePowerProfileByID, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatePowerProfileByID", err))
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
					"Failure when executing UpdatePowerProfileByID", err1))
				return diags
			}
		}

	}

	return resourceWirelessSettingsPowerProfilesIDRead(ctx, d, m)
}

func resourceWirelessSettingsPowerProfilesIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Wireless.DeletePowerProfileByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePowerProfileByID", err, restyResp1.String(),
				"Failure at DeletePowerProfileByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePowerProfileByID", err,
			"Failure at DeletePowerProfileByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletePowerProfileByID", err))
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
				"Failure when executing DeletePowerProfileByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdatePowerProfileByID {
	request := dnacentersdkgo.RequestWirelessUpdatePowerProfileByID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByIDRulesArray(ctx, key+".rules", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByIDRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessUpdatePowerProfileByIDRules {
	request := []dnacentersdkgo.RequestWirelessUpdatePowerProfileByIDRules{}
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
		i := expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByIDRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsPowerProfilesIDUpdatePowerProfileByIDRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessUpdatePowerProfileByIDRules {
	request := dnacentersdkgo.RequestWirelessUpdatePowerProfileByIDRules{}
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
