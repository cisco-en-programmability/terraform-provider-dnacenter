package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Issues.

- Create a new custom issue definition using the provided input request data. The unique identifier for this issue
definition is id. Please note that the issue names cannot be duplicated. The definition is based on the syslog. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml

- Updates an existing custom issue definition based on the provided Id. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml

- Deletes an existing custom issue definition based on the Id. Only the Global profile issue has the access to delete
the issue definition, so no profile id is required. For detailed information about the usage of the API, please refer to
the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml
`,

		CreateContext: resourceCustomIssueDefinitionsCreate,
		ReadContext:   resourceCustomIssueDefinitionsRead,
		UpdateContext: resourceCustomIssueDefinitionsUpdate,
		DeleteContext: resourceCustomIssueDefinitionsDelete,
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

						"created_time": &schema.Schema{
							Description: `Created Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_deletable": &schema.Schema{
							Description: `Is Deletable`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": &schema.Schema{
							Description: `Is Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_notification_enabled": &schema.Schema{
							Description: `Is Notification Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"profile_id": &schema.Schema{
							Description: `Profile Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration_in_minutes": &schema.Schema{
										Description: `Duration In Minutes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"occurrences": &schema.Schema{
										Description: `Occurrences`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"pattern": &schema.Schema{
										Description: `Pattern`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"trigger_id": &schema.Schema{
							Description: `Trigger Id`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Description`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. The custom issue definition Identifier
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"is_enabled": &schema.Schema{
							Description: `Is Enabled`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"is_notification_enabled": &schema.Schema{
							Description: `Is Notification Enabled`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration_in_minutes": &schema.Schema{
										Description: `Duration In Minutes`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"occurrences": &schema.Schema{
										Description: `Occurrences`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"pattern": &schema.Schema{
										Description: `Pattern`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
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

func resourceCustomIssueDefinitionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitions(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := dnacentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		resourceMap["id"] = item2.Name
		d.SetId(joinResourceID(resourceMap))
		return resourceCustomIssueDefinitionsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Issues.CreatesANewUserDefinedIssueDefinitions(request1, nil)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesANewUserDefinedIssueDefinitions", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesANewUserDefinedIssueDefinitions", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreatesANewUserDefinedIssueDefinitions", err,
			"Failure at CreatesANewUserDefinedIssueDefinitions, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceCustomIssueDefinitionsRead(ctx, d, m)
}

func resourceCustomIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters")
		queryParams1 := dnacentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams{}
		queryParams1.ID = vID
		queryParams1.Name = vName
		item1, err := searchIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(m, queryParams1, vID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse{
			*item1,
		}
		vItem1 := flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceCustomIssueDefinitionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Issues.UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID", err, restyResp1.String(),
					"Failure at UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID", err,
				"Failure at UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceCustomIssueDefinitionsRead(ctx, d, m)
}

func resourceCustomIssueDefinitionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	restyResp1, err := client.Issues.DeletesAnExistingCustomIssueDefinition(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesAnExistingCustomIssueDefinition", err, restyResp1.String(),
				"Failure at DeletesAnExistingCustomIssueDefinition, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesAnExistingCustomIssueDefinition", err,
			"Failure at DeletesAnExistingCustomIssueDefinition, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitions(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitions {
	request := dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitionsRulesArray(ctx, key+".rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_notification_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_notification_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_notification_enabled")))) {
		request.IsNotificationEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitionsRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules {
	request := []dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules{}
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
		i := expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitionsRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCustomIssueDefinitionsCreatesANewUserDefinedIssueDefinitionsRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules {
	request := dnacentersdkgo.RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severity")))) {
		request.Severity = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".facility")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".facility")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".facility")))) {
		request.Facility = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mnemonic")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mnemonic")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mnemonic")))) {
		request.Mnemonic = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pattern")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pattern")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pattern")))) {
		request.Pattern = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".occurrences")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".occurrences")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".occurrences")))) {
		request.Occurrences = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration_in_minutes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration_in_minutes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration_in_minutes")))) {
		request.DurationInMinutes = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID {
	request := dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rules")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rules")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rules")))) {
		request.Rules = expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRulesArray(ctx, key+".rules", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_enabled")))) {
		request.IsEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_notification_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_notification_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_notification_enabled")))) {
		request.IsNotificationEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRulesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules {
	request := []dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules{}
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
		i := expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCustomIssueDefinitionsUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules {
	request := dnacentersdkgo.RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severity")))) {
		request.Severity = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".facility")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".facility")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".facility")))) {
		request.Facility = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mnemonic")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mnemonic")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mnemonic")))) {
		request.Mnemonic = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pattern")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pattern")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pattern")))) {
		request.Pattern = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".occurrences")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".occurrences")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".occurrences")))) {
		request.Occurrences = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration_in_minutes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration_in_minutes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration_in_minutes")))) {
		request.DurationInMinutes = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(m interface{}, queryParams dnacentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams, vID string) (*dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse
	var ite *dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters

	ite, _, err = client.Issues.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err
	}
	itemsCopy := *ite.Response
	if itemsCopy == nil {
		return foundItem, err
	}
	for _, item := range itemsCopy {
		if item.Name == queryParams.Name {
			foundItem = &item
			return foundItem, err
		}
	}
	return foundItem, err
}
