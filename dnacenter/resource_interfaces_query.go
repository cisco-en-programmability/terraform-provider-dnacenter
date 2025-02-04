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
func resourceInterfacesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default
value: name.
The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper,
interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
interfaces-1.0.2-resolved.yaml
`,

		CreateContext: resourceInterfacesQueryCreate,
		ReadContext:   resourceInterfacesQueryRead,
		DeleteContext: resourceInterfacesQueryDelete,
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
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"admin_status": &schema.Schema{
										Description: `Admin Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
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
												"values": &schema.Schema{
													Type:     schema.TypeList,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Description: `Key`,
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
											},
										},
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"duplex_config": &schema.Schema{
										Description: `Duplex Config`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"duplex_oper": &schema.Schema{
										Description: `Duplex Oper`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"interface_if_index": &schema.Schema{
										Description: `Interface If Index`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"interface_type": &schema.Schema{
										Description: `Interface Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv4_address": &schema.Schema{
										Description: `Ipv4 Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"ipv6_address_list": &schema.Schema{
										Description: `Ipv6 Address List`,
										Type:        schema.TypeList,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"is_l3_interface": &schema.Schema{
										Description: `Is L3 Interface`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"is_wan": &schema.Schema{
										Description: `Is Wan`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"mac_addr": &schema.Schema{
										Description: `Mac Addr`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"media_type": &schema.Schema{
										Description: `Media Type`,
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
									"network_device_id": &schema.Schema{
										Description: `Network Device Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"network_device_ip_address": &schema.Schema{
										Description: `Network Device Ip Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"network_device_mac_address": &schema.Schema{
										Description: `Network Device Mac Address`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"oper_status": &schema.Schema{
										Description: `Oper Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"peer_stack_member": &schema.Schema{
										Description: `Peer Stack Member`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"peer_stack_port": &schema.Schema{
										Description: `Peer Stack Port`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"port_channel_id": &schema.Schema{
										Description: `Port Channel Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"port_mode": &schema.Schema{
										Description: `Port Mode`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"port_type": &schema.Schema{
										Description: `Port Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_discards": &schema.Schema{
										Description: `Rx Discards`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_error": &schema.Schema{
										Description: `Rx Error`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_rate": &schema.Schema{
										Description: `Rx Rate`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_utilization": &schema.Schema{
										Description: `Rx Utilization`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy_id": &schema.Schema{
										Description: `Site Hierarchy Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_name": &schema.Schema{
										Description: `Site Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"speed": &schema.Schema{
										Description: `Speed`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"stack_port_type": &schema.Schema{
										Description: `Stack Port Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_discards": &schema.Schema{
										Description: `Tx Discards`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_error": &schema.Schema{
										Description: `Tx Error`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_rate": &schema.Schema{
										Description: `Tx Rate`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_utilization": &schema.Schema{
										Description: `Tx Utilization`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
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
						"views": &schema.Schema{
							Description: `Views`,
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
		},
	}
}

func resourceInterfacesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceInterfacesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceInterfacesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters{}
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
		i := expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters{}
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
		request.Value = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue {
	var request dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters{}
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
		i := expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters{}
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
		request.Value = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue {
	var request dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes{}
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
		i := expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy {
	request := []dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy{}
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
		i := expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestInterfacesQueryGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy {
	request := dnacentersdkgo.RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItems(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["admin_status"] = item.AdminStatus
		respItem["description"] = item.Description
		respItem["duplex_config"] = item.DuplexConfig
		respItem["duplex_oper"] = item.DuplexOper
		respItem["interface_if_index"] = item.InterfaceIfIndex
		respItem["interface_type"] = item.InterfaceType
		respItem["ipv4_address"] = item.IPv4Address
		respItem["ipv6_address_list"] = item.IPv6AddressList
		respItem["is_l3_interface"] = boolPtrToString(item.IsL3Interface)
		respItem["is_wan"] = boolPtrToString(item.IsWan)
		respItem["mac_addr"] = item.MacAddr
		respItem["media_type"] = item.MediaType
		respItem["name"] = item.Name
		respItem["oper_status"] = item.OperStatus
		respItem["peer_stack_member"] = item.PeerStackMember
		respItem["peer_stack_port"] = item.PeerStackPort
		respItem["port_channel_id"] = item.PortChannelID
		respItem["port_mode"] = item.PortMode
		respItem["port_type"] = item.PortType
		respItem["rx_discards"] = item.RxDiscards
		respItem["rx_error"] = item.RxError
		respItem["rx_rate"] = item.RxRate
		respItem["rx_utilization"] = item.RxUtilization
		respItem["speed"] = item.Speed
		respItem["stack_port_type"] = item.StackPortType
		respItem["timestamp"] = item.Timestamp
		respItem["tx_discards"] = item.TxDiscards
		respItem["tx_error"] = item.TxError
		respItem["tx_rate"] = item.TxRate
		respItem["tx_utilization"] = item.TxUtilization
		respItem["vlan_id"] = item.VLANID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["network_device_ip_address"] = item.NetworkDeviceIPAddress
		respItem["network_device_mac_address"] = item.NetworkDeviceMacAddress
		respItem["site_name"] = item.SiteName
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["site_hierarchy_id"] = item.SiteHierarchyID
		respItem["aggregate_attributes"] = flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributes(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["values"] = flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributesValues(item.Values)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsItemsAggregateAttributesValues(items *[]dnacentersdkgo.ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributesValues) []map[string]interface{} {
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
