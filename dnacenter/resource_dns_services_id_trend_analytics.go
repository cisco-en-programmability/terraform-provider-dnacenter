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
func resourceDNSServicesIDTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the trend analytics data related to a particular DNS Service matching the id. For detailed information about the
usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml
`,

		CreateContext: resourceDNSServicesIDTrendAnalyticsCreate,
		ReadContext:   resourceDNSServicesIDTrendAnalyticsRead,
		DeleteContext: resourceDNSServicesIDTrendAnalyticsDelete,
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
						"id": &schema.Schema{
							Description: `id path parameter. Unique id of the DNS Service. It is the combination of DNS Server IP (*serverIp*) and Device UUID (*deviceId*) separated by underscore (*_*). Example: If *serverIp* is *10.76.81.33* and *deviceId* is *6bef213c-19ca-4170-8375-b694e251101c*, then the *id* would be *10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c*
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
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
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"groups": &schema.Schema{
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
																Type:        schema.TypeString,
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
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
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
									"timestamp_order": &schema.Schema{
										Description: `Timestamp Order`,
										Type:        schema.TypeString,
										Optional:    true,
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
						"trend_interval": &schema.Schema{
							Description: `Trend Interval`,
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

func resourceDNSServicesIDTrendAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	vvID := vID.(string)
	request1 := expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Devices.GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(vvID, request1, &headerParams1)

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

	vItems1 := flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceDNSServicesIDTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDNSServicesIDTrendAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService {
	request := dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval")))) {
		request.TrendInterval = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters {
	request := []dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters{}
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
		i := expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters {
	request := dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters{}
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
		request.Value = expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue {
	var request dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes{}
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
		i := expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestDNSServicesIDTrendAnalyticsGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage {
	request := dnacentersdkgo.RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp_order")))) {
		request.TimestampOrder = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItems(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["groups"] = flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroups(item.Groups)
		respItem["attributes"] = flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroups(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroupsAttributes(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAttributes) []map[string]interface{} {
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

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsGroupsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes) []map[string]interface{} {
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

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsAttributes(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAttributes) []map[string]interface{} {
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

func flattenDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAggregateAttributes) []map[string]interface{} {
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
