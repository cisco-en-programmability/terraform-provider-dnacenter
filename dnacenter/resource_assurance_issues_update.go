package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAssuranceIssuesUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Issues.

- Updates selected fields in the given issue. Currently the only field that can be updated is 'notes' field. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
IssuesLifecycle-1.0.0-resolved.yaml
`,

		CreateContext: resourceAssuranceIssuesUpdateCreate,
		ReadContext:   resourceAssuranceIssuesUpdateRead,
		DeleteContext: resourceAssuranceIssuesUpdateDelete,
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

						"additional_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"entity_id": &schema.Schema{
							Description: `Entity Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"entity_type": &schema.Schema{
							Description: `Entity Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"first_occurred_time": &schema.Schema{
							Description: `First Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"is_global": &schema.Schema{
							Description: `Is Global`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"issue_id": &schema.Schema{
							Description: `Issue Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"most_recent_occurred_time": &schema.Schema{
							Description: `Most Recent Occurred Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"notes": &schema.Schema{
							Description: `Notes`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchy Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_name": &schema.Schema{
							Description: `Site Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"suggested_actions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"message": &schema.Schema{
										Description: `Message`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"summary": &schema.Schema{
							Description: `Summary`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_by": &schema.Schema{
							Description: `Updated By`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_time": &schema.Schema{
							Description: `Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
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
						"accept_language": &schema.Schema{
							Description: `Accept-Language header parameter. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. The issue Uuid
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
						"notes": &schema.Schema{
							Description: `Notes`,
							Type:        schema.TypeString,
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

func resourceAssuranceIssuesUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vAcceptLanguage := resourceItem["accept_language"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	vvID := vID.(string)
	request1 := expandRequestAssuranceIssuesUpdateUpdateTheGivenIssueByUpdatingSelectedFields(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams{}

	headerParams1.AcceptLanguage = vAcceptLanguage.(string)

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Issues.UpdateTheGivenIssueByUpdatingSelectedFields(vvID, request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateTheGivenIssueByUpdatingSelectedFields", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateTheGivenIssueByUpdatingSelectedFields response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceAssuranceIssuesUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssuranceIssuesUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssuranceIssuesUpdateUpdateTheGivenIssueByUpdatingSelectedFields(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFields {
	request := dnacentersdkgo.RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFields{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".notes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".notes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".notes")))) {
		request.Notes = interfaceToString(v)
	}
	return &request
}

func flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItem(item *dnacentersdkgo.ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["updated_by"] = item.UpdatedBy
	respItem["updated_time"] = item.UpdatedTime
	respItem["notes"] = item.Notes
	respItem["site_id"] = item.SiteID
	respItem["site_hierarchy_id"] = item.SiteHierarchyID
	respItem["site_name"] = item.SiteName
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["suggested_actions"] = flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItemSuggestedActions(item.SuggestedActions)
	respItem["additional_attributes"] = flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItemAdditionalAttributes(item.AdditionalAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItemSuggestedActions(items *[]dnacentersdkgo.ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsItemAdditionalAttributes(items *[]dnacentersdkgo.ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseAdditionalAttributes) []map[string]interface{} {
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
