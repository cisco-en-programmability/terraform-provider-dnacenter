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

func resourceSiteKpiSummariesTopNAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Sites.

- Gets the Top N entites related based on site analytics for a given kpi type. For detailed information about the usage
of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		CreateContext: resourceSiteKpiSummariesTopNAnalyticsCreate,
		ReadContext:   resourceSiteKpiSummariesTopNAnalyticsRead,
		UpdateContext: resourceSiteKpiSummariesTopNAnalyticsUpdate,
		DeleteContext: resourceSiteKpiSummariesTopNAnalyticsDelete,
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

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Id`,
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

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"group_by": &schema.Schema{
							Description: `Group By`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"task_id": &schema.Schema{
							Description: `task id`,
							Type:        schema.TypeString,
							Required:    true,
						}, "xca_lle_rid": &schema.Schema{
							Description: `xca lle rid`,
							Type:        schema.TypeString,
							Required:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"top_n": &schema.Schema{
							Description: `Top N`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteKpiSummariesTopNAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalytics(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTaskID := resourceItem["task_id"]
	vvTaskID := interfaceToString(vTaskID)
	vXCaLLERID := resourceItem["xca_lle_rid"]
	vvXCaLLERID := interfaceToString(vXCaLLERID)
	queryParamImport := dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams{}
	queryParamImport.TaskID = vvTaskID
	headerParams1 := dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDHeaderParams{}
	headerParams1.XCaLLERID = vvXCaLLERID
	item2, err := searchSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(m, queryParamImport, headerParams1)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["task_id"] = vvTaskID
		resourceMap["xca_lle_rid"] = vvXCaLLERID
		d.SetId(joinResourceID(resourceMap))
		return resourceSiteKpiSummariesTopNAnalyticsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sites.SubmitRequestForTopNEntitiesRelatedToSiteAnalytics(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing SubmitRequestForTopNEntitiesRelatedToSiteAnalytics", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForTopNEntitiesRelatedToSiteAnalytics", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForTopNEntitiesRelatedToSiteAnalytics", err))
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
				"Failure when executing SubmitRequestForTopNEntitiesRelatedToSiteAnalytics", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams{}
	queryParamValidate.TaskID = vvTaskID
	item3, err := searchSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(m, queryParamValidate, headerParams1)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SubmitRequestForTopNEntitiesRelatedToSiteAnalytics", err,
			"Failure at SubmitRequestForTopNEntitiesRelatedToSiteAnalytics, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["task_id"] = vvTaskID
	resourceMap["xca_lle_rid"] = vvXCaLLERID
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteKpiSummariesTopNAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesTopNAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTaskID := resourceMap["task_id"]

	vXCaLLERID := resourceMap["xca_lle_rid"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID")

		headerParams1 := dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDHeaderParams{}
		queryParams1 := dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams{}

		queryParams1.TaskID = vTaskID

		headerParams1.XCaLLERID = vXCaLLERID

		response1, restyResp1, err := client.Sites.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(m, queryParams1, headerParams1)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDByIDItem(item *dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["attributes"] = flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDByIDItemAttributes(item.Attributes)
	return []map[string]interface{}{
		respItem,
	}
}
func flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDByIDItemAttributes(items *[]dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponseAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func resourceSiteKpiSummariesTopNAnalyticsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSiteKpiSummariesTopNAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesTopNAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SiteKpiSummariesTopNAnalytics on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalytics(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics {
	request := dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_n")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_n")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_n")))) {
		request.TopN = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFiltersArray(ctx, key+".filters", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters {
	request := []dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters{}
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
		i := expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesTopNAnalyticsSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters {
	request := dnacentersdkgo.RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(m interface{}, queryParams dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams, header dnacentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDHeaderParams) (*dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse
	var ite *dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID
	ite, _, err = client.Sites.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(&header, &queryParams)
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.ID == queryParams.TaskID {
			var getItem *dnacentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
