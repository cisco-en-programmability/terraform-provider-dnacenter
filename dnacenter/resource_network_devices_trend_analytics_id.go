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
func resourceNetworkDevicesTrendAnalyticsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- The Trend analytics data for the network Device in the specified time range. The data is grouped based on the trend
time Interval, other input parameters like attribute and aggregate attributes. For detailed information about the usage
of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		CreateContext: resourceNetworkDevicesTrendAnalyticsIDCreate,
		ReadContext:   resourceNetworkDevicesTrendAnalyticsIDRead,
		DeleteContext: resourceNetworkDevicesTrendAnalyticsIDDelete,
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
							Description: `id path parameter. The device Uuid
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
													Type:        schema.TypeFloat,
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
																Type:        schema.TypeFloat,
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
										Type:        schema.TypeFloat,
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
						"trend_interval_in_minutes": &schema.Schema{
							Description: `Trend Interval In Minutes`,
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

func resourceNetworkDevicesTrendAnalyticsIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(vvID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesTrendAnalyticsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesTrendAnalyticsIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval_in_minutes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval_in_minutes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval_in_minutes")))) {
		request.TrendIntervalInMinutes = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters {
	request := []dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters{}
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
		i := expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters{}
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
		request.Value = expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue {
	var request dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes{}
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
		i := expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsIDTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage{}
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

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItems(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsAggregateAttributes(item.AggregateAttributes)
		respItem["groups"] = flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroups(item.Groups)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAttributes) []map[string]interface{} {
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

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAggregateAttributes) []map[string]interface{} {
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

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroups(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroupsAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAttributes) []map[string]interface{} {
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

func flattenDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeItemsGroupsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAggregateAttributes) []map[string]interface{} {
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
