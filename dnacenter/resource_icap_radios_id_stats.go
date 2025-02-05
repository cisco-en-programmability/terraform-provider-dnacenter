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
func resourceIcapRadiosIDStats() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Retrieves the time series statistics of a specific radio by applying complex filters. If startTime and endTime are not
provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to
the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapRadiosIDStatsCreate,
		ReadContext:   resourceIcapRadiosIDStatsRead,
		DeleteContext: resourceIcapRadiosIDStatsDelete,
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
							Description: `id path parameter. id is the composite key made of AP Base Ethernet macAddress and Radio Slot Id. Format apMac_RadioId
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
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac": &schema.Schema{
										Description: `Ap Mac`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"band": &schema.Schema{
										Description: `Band`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"channel": &schema.Schema{
										Description: `Channel`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"channel_width": &schema.Schema{
										Description: `Channel Width`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"max_tx_power": &schema.Schema{
										Description: `Max Tx Power`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"noise_floor": &schema.Schema{
										Description: `Noise Floor`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"non_wifi_utilization": &schema.Schema{
										Description: `Non Wifi Utilization`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"radio_id": &schema.Schema{
										Description: `Radio Id`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_bytes": &schema.Schema{
										Description: `Rx Bytes`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_errors": &schema.Schema{
										Description: `Rx Errors`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_in_bss_utilization": &schema.Schema{
										Description: `Rx In B S S Utilization`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_mgmt_packets": &schema.Schema{
										Description: `Rx Mgmt Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_other_bss_utilization": &schema.Schema{
										Description: `Rx Other B S S Utilization`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_packets": &schema.Schema{
										Description: `Rx Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_errors": &schema.Schema{
										Description: `Tx Errors`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_mgmt_packets": &schema.Schema{
										Description: `Tx Mgmt Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_packets": &schema.Schema{
										Description: `Tx Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_power": &schema.Schema{
										Description: `Tx Power`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_utilization": &schema.Schema{
										Description: `Tx Utilization`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"utilization": &schema.Schema{
										Description: `Utilization`,
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
									"time_sort_order": &schema.Schema{
										Description: `Time Sort Order`,
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
					},
				},
			},
		},
	}
}

func resourceIcapRadiosIDStatsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	vvID := vID.(string)
	request1 := expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Sensors.RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(vvID, request1, &headerParams1)

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

	vItems1 := flattenSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceIcapRadiosIDStatsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapRadiosIDStatsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters {
	request := []dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters{}
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
		i := expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters{}
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

func expandRequestIcapRadiosIDStatsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_sort_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_sort_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_sort_order")))) {
		request.TimeSortOrder = interfaceToString(v)
	}
	return &request
}

func flattenSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["ap_mac"] = item.ApMac
		respItem["radio_id"] = item.RadioID
		respItem["band"] = item.Band
		respItem["utilization"] = item.Utilization
		respItem["non_wifi_utilization"] = item.NonWifiUtilization
		respItem["rx_other_bss_utilization"] = item.RxOtherBSSUtilization
		respItem["rx_in_bss_utilization"] = item.RxInBSSUtilization
		respItem["tx_utilization"] = item.TxUtilization
		respItem["noise_floor"] = item.NoiseFloor
		respItem["channel"] = item.Channel
		respItem["channel_width"] = item.ChannelWidth
		respItem["tx_power"] = item.TxPower
		respItem["max_tx_power"] = item.MaxTxPower
		respItem["tx_bytes"] = item.TxBytes
		respItem["rx_bytes"] = item.RxBytes
		respItem["rx_packets"] = item.RxPackets
		respItem["tx_packets"] = item.TxPackets
		respItem["rx_mgmt_packets"] = item.RxMgmtPackets
		respItem["tx_mgmt_packets"] = item.TxMgmtPackets
		respItem["rx_errors"] = item.RxErrors
		respItem["tx_errors"] = item.TxErrors
		respItems = append(respItems, respItem)
	}
	return respItems
}
