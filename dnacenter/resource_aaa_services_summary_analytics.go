package dnacenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAAAServicesSummaryAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the summary analytics data related to AAA Services based on given filters and group by field. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml
`,

		CreateContext: resourceAAAServicesSummaryAnalyticsCreate,
		ReadContext:   resourceAAAServicesSummaryAnalyticsRead,
		DeleteContext: resourceAAAServicesSummaryAnalyticsDelete,
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

						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Computed:    true,
									},
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
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													Computed:    true,
												},
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
													Type:        schema.TypeString,
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
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
										Description: `Filters`,
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
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
										Type:        schema.TypeString, //TEST,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"group_by": &schema.Schema{
							Description: `Group By`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"offset": &schema.Schema{
										Description: `Offset`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"order": &schema.Schema{
													Description: `Order`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
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

func resourceAAAServicesSummaryAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Devices.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(request1, &headerParams1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItem1 := flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceAAAServicesSummaryAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAAAServicesSummaryAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters {
	request := dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters {
	request := []dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters {
	request := dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue {
	var request dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage {
	request := dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy {
	request := []dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy {
	request := dnacentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItem(item *dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroups(item.Groups)
	respItem["attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemAttributes(item.Attributes)
	respItem["aggregate_attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemAggregateAttributes(item.AggregateAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroups(items *[]dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroupsAttributes(items *[]dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAttributes) []map[string]interface{} {
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

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemGroupsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["function"] = item.Function
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemAttributes(items *[]dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes) []map[string]interface{} {
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

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersItemAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["function"] = item.Function
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
