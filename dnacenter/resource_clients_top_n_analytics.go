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
func resourceClientsTopNAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Clients.

- Retrieves the top N analytics data related to clients based on the provided input data. This data source action
facilitates obtaining insights into the top-performing or most impacted clients. For detailed information about the
usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml
`,

		CreateContext: resourceClientsTopNAnalyticsCreate,
		ReadContext:   resourceClientsTopNAnalyticsRead,
		DeleteContext: resourceClientsTopNAnalyticsDelete,
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
										Type:        schema.TypeInt,
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
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cursor": &schema.Schema{
										Description: `Cursor`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"limit": &schema.Schema{
										Description: `Limit`,
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
						"top_n": &schema.Schema{
							Description: `Top N`,
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

func resourceClientsTopNAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClients(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Clients.RetrievesTheTopNAnalyticsDataRelatedToClients(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing RetrievesTheTopNAnalyticsDataRelatedToClients", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheTopNAnalyticsDataRelatedToClients response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceClientsTopNAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceClientsTopNAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClients(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClients {
	request := dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClients{}
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters {
	request := []dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters{}
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
		i := expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters {
	request := dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes {
	request := []dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes{}
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
		i := expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes {
	request := dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage {
	request := dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cursor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cursor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cursor")))) {
		request.Cursor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy {
	request := []dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy{}
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
		i := expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestClientsTopNAnalyticsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy {
	request := dnacentersdkgo.RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItems(items *[]dnacentersdkgo.ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItemsAttributes(items *[]dnacentersdkgo.ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAttributes) []map[string]interface{} {
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

func flattenClientsRetrievesTheTopNAnalyticsDataRelatedToClientsItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAggregateAttributes) []map[string]interface{} {
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
