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
func resourceInterfacesIDTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- The Trend analytcis data for the interface, identified by its instanceUuid, in the specified time range. The data is
grouped based on the trend time Interval, other input parameters like attributes and aggregate attributes. The default
time interval range is 3 hours when start and endTime is not provided.
The field trendIntervalInMinutes is requiered and either the attributes or the aggregateAttributes fields is also
required.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
interfaces-2.0.0-resolved.yaml
`,

		CreateContext: resourceInterfacesIDTrendAnalyticsCreate,
		ReadContext:   resourceInterfacesIDTrendAnalyticsRead,
		DeleteContext: resourceInterfacesIDTrendAnalyticsDelete,
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
							Description: `id path parameter. The interface instance Uuid
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

									"aggregate_attributes": &schema.Schema{
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
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
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
						"timestamp_order": &schema.Schema{
							Description: `Timestamp Order`,
							Type:        schema.TypeString,
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

func resourceInterfacesIDTrendAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Devices.TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(vvID, request1)

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

	vItems1 := flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceInterfacesIDTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceInterfacesIDTrendAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval_in_minutes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval_in_minutes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval_in_minutes")))) {
		request.TrendIntervalInMinutes = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp_order")))) {
		request.TimestampOrder = interfaceToString(v)
	}
	return &request
}

func expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters {
	request := []dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters{}
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
		i := expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters{}
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

func expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes{}
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
		i := expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesIDTrendAnalyticsTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItems(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItemsAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAttributes) []map[string]interface{} {
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

func flattenDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
