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

// resourceAction
func resourceAssuranceIssuesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Issues.

- Returns all details of each issue along with suggested actions for given set of filters specified in request body. If
there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to
24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml
`,

		CreateContext: resourceAssuranceIssuesQueryCreate,
		ReadContext:   resourceAssuranceIssuesQueryRead,
		DeleteContext: resourceAssuranceIssuesQueryDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_language": &schema.Schema{
							Description: `Accept-Language header parameter. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"operator": &schema.Schema{
													Description: `Operator`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"logical_operator": &schema.Schema{
										Description: `Logical Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"category": &schema.Schema{
										Description: `Category`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"entity_id": &schema.Schema{
										Description: `Entity Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"entity_type": &schema.Schema{
										Description: `Entity Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"first_occurred_time": &schema.Schema{
										Description: `First Occurred Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"is_global": &schema.Schema{
										Description: `Is Global`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"issue_id": &schema.Schema{
										Description: `Issue Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"most_recent_occurred_time": &schema.Schema{
										Description: `Most Recent Occurred Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"notes": &schema.Schema{
										Description: `Notes`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"priority": &schema.Schema{
										Description: `Priority`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy_id": &schema.Schema{
										Description: `Site Hierarchy Id`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_name": &schema.Schema{
										Description: `Site Name`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"status": &schema.Schema{
										Description: `Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"suggested_actions": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Description: `Message`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"steps": &schema.Schema{
													Description: `Steps`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"summary": &schema.Schema{
										Description: `Summary`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"updated_by": &schema.Schema{
										Description: `Updated By`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"updated_time": &schema.Schema{
										Description: `Updated Time`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceAssuranceIssuesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vAcceptLanguage := resourceItem["accept_language"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFilters(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams{}

	headerParams1.AcceptLanguage = vAcceptLanguage.(string)

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Issues.GetTheDetailsOfIssuesForGivenSetOfFilters(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetTheDetailsOfIssuesForGivenSetOfFilters", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetTheDetailsOfIssuesForGivenSetOfFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceAssuranceIssuesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssuranceIssuesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters {
	request := dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters {
	request := []dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters{}
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
		i := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters {
	request := dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters {
	request := []dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters{}
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
		i := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters {
	request := dnacentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItems(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["summary"] = item.Summary
		respItem["priority"] = item.Priority
		respItem["severity"] = item.Severity
		respItem["device_type"] = item.DeviceType
		respItem["category"] = item.Category
		respItem["entity_type"] = item.EntityType
		respItem["entity_id"] = item.EntityID
		respItem["first_occurred_time"] = item.FirstOccurredTime
		respItem["most_recent_occurred_time"] = item.MostRecentOccurredTime
		respItem["status"] = item.Status
		respItem["is_global"] = boolPtrToString(item.IsGlobal)
		respItem["updated_by"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsUpdatedBy(item.UpdatedBy)
		respItem["updated_time"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsUpdatedTime(item.UpdatedTime)
		respItem["notes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsNotes(item.Notes)
		respItem["site_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteID(item.SiteID)
		respItem["site_hierarchy_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteHierarchyID(item.SiteHierarchyID)
		respItem["site_name"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteName(item.SiteName)
		respItem["site_hierarchy"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteHierarchy(item.SiteHierarchy)
		respItem["suggested_actions"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSuggestedActions(item.SuggestedActions)
		respItem["additional_attributes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsAdditionalAttributes(item.AdditionalAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsUpdatedBy(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedBy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsUpdatedTime(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsNotes(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseNotes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteID(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteHierarchyID(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchyID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteName(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSiteHierarchy(item *dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSuggestedActions(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsSuggestedActionsSteps(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActionsSteps) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersItemsAdditionalAttributes(items *[]dnacentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseAdditionalAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
