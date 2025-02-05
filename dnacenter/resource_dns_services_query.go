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
func resourceDNSServicesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Retrieves the list of DNS Services and offers complex filtering and sorting capabilities. For detailed information
about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml
`,

		CreateContext: resourceDNSServicesQueryCreate,
		ReadContext:   resourceDNSServicesQueryRead,
		DeleteContext: resourceDNSServicesQueryDelete,
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
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

									"device_family": &schema.Schema{
										Description: `Device Family`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_id": &schema.Schema{
										Description: `Device Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_name": &schema.Schema{
										Description: `Device Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_site_hierarchy": &schema.Schema{
										Description: `Device Site Hierarchy`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_site_hierarchy_id": &schema.Schema{
										Description: `Device Site Hierarchy Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_site_id": &schema.Schema{
										Description: `Device Site Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"failed_transactions": &schema.Schema{
										Description: `Failed Transactions`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"failures": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"failed_transactions": &schema.Schema{
													Description: `Failed Transactions`,
													Type:        schema.TypeInt,
													ForceNew:    true,
													Computed:    true,
												},
												"failure_description": &schema.Schema{
													Description: `Failure Description`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"failure_response_code": &schema.Schema{
													Description: `Failure Response Code`,
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
									"latency": &schema.Schema{
										Description: `Latency`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"server_ip": &schema.Schema{
										Description: `Server Ip`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"successful_transactions": &schema.Schema{
										Description: `Successful Transactions`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"transactions": &schema.Schema{
										Description: `Transactions`,
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
					},
				},
			},
		},
	}
}

func resourceDNSServicesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Devices.RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(request1, &headerParams1)

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

	vItems1 := flattenDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceDNSServicesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDNSServicesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters {
	request := dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters {
	request := []dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters{}
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
		i := expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters {
	request := dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage {
	request := dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy {
	request := []dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy{}
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
		i := expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDNSServicesQueryRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy {
	request := dnacentersdkgo.RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersItems(items *[]dnacentersdkgo.ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["server_ip"] = item.ServerIP
		respItem["device_id"] = item.DeviceID
		respItem["device_name"] = item.DeviceName
		respItem["device_family"] = item.DeviceFamily
		respItem["device_site_hierarchy"] = item.DeviceSiteHierarchy
		respItem["device_site_id"] = item.DeviceSiteID
		respItem["device_site_hierarchy_id"] = item.DeviceSiteHierarchyID
		respItem["transactions"] = item.Transactions
		respItem["failed_transactions"] = item.FailedTransactions
		respItem["failures"] = flattenDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersItemsFailures(item.Failures)
		respItem["successful_transactions"] = item.SuccessfulTransactions
		respItem["latency"] = item.Latency
		respItem["ssid"] = item.SSID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersItemsFailures(items *[]dnacentersdkgo.ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponseFailures) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["failure_response_code"] = item.FailureResponseCode
		respItem["failure_description"] = item.FailureDescription
		respItem["failed_transactions"] = item.FailedTransactions
		respItems = append(respItems, respItem)
	}
	return respItems
}
