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

func resourceSiteKpiSummariesSummaryAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Sites.

- Submits the task to get summary analytics data for a given site. For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		CreateContext: resourceSiteKpiSummariesSummaryAnalyticsCreate,
		ReadContext:   resourceSiteKpiSummariesSummaryAnalyticsRead,
		UpdateContext: resourceSiteKpiSummariesSummaryAnalyticsUpdate,
		DeleteContext: resourceSiteKpiSummariesSummaryAnalyticsDelete,
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
						},
						"xca_lle_rid": &schema.Schema{
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
						"start_time": &schema.Schema{
							Description: `Start Time`,
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

func resourceSiteKpiSummariesSummaryAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryData(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTaskID := resourceItem["task_id"]
	vvTaskID := interfaceToString(vTaskID)
	vXCaLLERID := resourceItem["xca_lle_rid"]
	vvXCaLLERID := interfaceToString(vXCaLLERID)
	queryParamImport := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams{}
	queryParamImport.TaskID = vvTaskID
	item2, _, err := client.Sites.GetSiteAnalyticsSummaryDataForTheGivenTaskID(nil, &queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["task_id"] = vvTaskID
		resourceMap["xca_lle_rid"] = vvXCaLLERID
		d.SetId(joinResourceID(resourceMap))
		return resourceSiteKpiSummariesSummaryAnalyticsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sites.SubmitRequestForSiteAnalyticsSummaryData(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing SubmitRequestForSiteAnalyticsSummaryData", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForSiteAnalyticsSummaryData", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing SubmitRequestForSiteAnalyticsSummaryData", err))
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
				"Failure when executing SubmitRequestForSiteAnalyticsSummaryData", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams{}
	queryParamValidate.TaskID = vvTaskID
	item3, _, err := client.Sites.GetSiteAnalyticsSummaryDataForTheGivenTaskID(nil, &queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing SubmitRequestForSiteAnalyticsSummaryData", err,
			"Failure at SubmitRequestForSiteAnalyticsSummaryData, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["task_id"] = vvTaskID
	resourceMap["xca_lle_rid"] = vvXCaLLERID
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteKpiSummariesSummaryAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesSummaryAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vTaskID := resourceMap["task_id"]

	vXCaLLERID := resourceMap["xca_lle_rid"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteAnalyticsSummaryDataForTheGivenTaskID")

		headerParams1 := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDHeaderParams{}
		queryParams1 := dnacentersdkgo.GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams{}

		queryParams1.TaskID = vTaskID

		headerParams1.XCaLLERID = vXCaLLERID

		response1, restyResp1, err := client.Sites.GetSiteAnalyticsSummaryDataForTheGivenTaskID(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteAnalyticsSummaryDataForTheGivenTaskID response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSiteKpiSummariesSummaryAnalyticsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSiteKpiSummariesSummaryAnalyticsRead(ctx, d, m)
}

func resourceSiteKpiSummariesSummaryAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SiteKpiSummariesSummaryAnalytics on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryData(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryData {
	request := dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryData{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryDataFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryDataFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters {
	request := []dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters{}
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
		i := expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryDataFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSiteKpiSummariesSummaryAnalyticsSubmitRequestForSiteAnalyticsSummaryDataFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters {
	request := dnacentersdkgo.RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters{}
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
