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

func resourceSiteKpiSummariesTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Sites.

- Submits the task to get site analytics trend data for a given site. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		CreateContext: resourceSiteKpiSummariesTrendAnalyticsCreate,
		ReadContext:   resourceSiteKpiSummariesTrendAnalyticsRead,
		UpdateContext: resourceSiteKpiSummariesTrendAnalyticsUpdate,
		DeleteContext: resourceSiteKpiSummariesTrendAnalyticsDelete,
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
						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
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

						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
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
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"offset": &schema.Schema{
										Description: `Offset`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"timestamp_order": &schema.Schema{
										Description: `Timestamp Order`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"trend_interval": &schema.Schema{
							Description: `Trend Interval`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteKpiSummariesTrendAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendData(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTaskID := resourceItem["task_id"]
	vvTaskID := interfaceToString(vTaskID)
	vXCaLLERID := resourceItem["xca_lle_rid"]
	vvXCaLLERID := interfaceToString(vXCaLLERID)
	queryParamImport := dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams{}
	queryParamImport.TaskID = vvTaskID
	headerParams1 := dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDHeaderParams{}
	headerParams1.XCaLLERID = vvXCaLLERID
	item2, err := searchSitesGetSiteAnalyticsTrendDataForTheGivenTaskID(m, queryParamImport, headerParams1)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["task_id"] = vvTaskID
		resourceMap["xca_lle_rid"] = vvXCaLLERID
		d.SetId(joinResourceID(resourceMap))
		return resourceSiteKpiSummariesTrendAnalyticsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sites.SubmitRequestForSiteAnalyticsTrendData(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing SubmitRequestForSiteAnalyticsTrendData", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForSiteAnalyticsTrendData", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForSiteAnalyticsTrendData", err))
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
				"Failure when executing SubmitRequestForSiteAnalyticsTrendData", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams{}
	queryParamValidate.TaskID = vvTaskID
	item3, err := searchSitesGetSiteAnalyticsTrendDataForTheGivenTaskID(m, queryParamValidate, headerParams1)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SubmitRequestForSiteAnalyticsTrendData", err,
			"Failure at SubmitRequestForSiteAnalyticsTrendData, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["task_id"] = vvTaskID
	resourceMap["xca_lle_rid"] = vvXCaLLERID
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteKpiSummariesTrendAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTaskID := resourceMap["task_id"]

	vXCaLLERID := resourceMap["xca_lle_rid"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAnalyticsTrendDataForTheGivenTaskID")

		headerParams1 := dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDHeaderParams{}
		queryParams1 := dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams{}

		queryParams1.TaskID = vTaskID

		headerParams1.XCaLLERID = vXCaLLERID

		response1, restyResp1, err := client.Sites.GetSiteAnalyticsTrendDataForTheGivenTaskID(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchSitesGetSiteAnalyticsTrendDataForTheGivenTaskID(m, queryParams1, headerParams1)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAnalyticsTrendDataForTheGivenTaskID search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDByIDItem(item *dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["timestamp"] = item.Timestamp
	respItem["attributes"] = flattenSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDByIDItemAttributes(item.Attributes)
	return []map[string]interface{}{
		respItem,
	}
}
func flattenSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDByIDItemAttributes(items *[]dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponseAttributes) []map[string]interface{} {
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

func resourceSiteKpiSummariesTrendAnalyticsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSiteKpiSummariesTrendAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesTrendAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SiteKpiSummariesTrendAnalytics on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendData(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendData {
	request := dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval")))) {
		request.TrendInterval = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataPage(ctx, key+".page.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters {
	request := []dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters{}
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
		i := expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters {
	request := dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters{}
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

func expandRequestSiteKpiSummariesTrendAnalyticsSubmitRequestForSiteAnalyticsTrendDataPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataPage {
	request := dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsTrendDataPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp_order")))) {
		request.TimestampOrder = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSitesGetSiteAnalyticsTrendDataForTheGivenTaskID(m interface{}, queryParams dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams, header dnacentersdkgo.GetSiteAnalyticsTrendDataForTheGivenTaskIDHeaderParams) (*dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse
	var ite *dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskID
	ite, _, err = client.Sites.GetSiteAnalyticsTrendDataForTheGivenTaskID(&header, &queryParams)
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if "" == queryParams.TaskID {
			var getItem *dnacentersdkgo.ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
