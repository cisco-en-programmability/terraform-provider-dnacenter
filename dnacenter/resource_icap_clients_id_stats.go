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
func resourceIcapClientsIDStats() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Retrieves the time series statistics of a specific client by applying complex filters. If startTime and endTime are
not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer
to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapClientsIDStatsCreate,
		ReadContext:   resourceIcapClientsIDStatsRead,
		DeleteContext: resourceIcapClientsIDStatsDelete,
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
							Description: `id path parameter. id is the client mac address. It can be specified in one of the notational conventions 01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive
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
									"client_ip": &schema.Schema{
										Description: `Client Ip`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"client_mac": &schema.Schema{
										Description: `Client Mac`,
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
									"radio_id": &schema.Schema{
										Description: `Radio Id`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rssi": &schema.Schema{
										Description: `Rssi`,
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
									"rx_ctrl_packets": &schema.Schema{
										Description: `Rx Ctrl Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_data_packets": &schema.Schema{
										Description: `Rx Data Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_mgmt_packets": &schema.Schema{
										Description: `Rx Mgmt Packets`,
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
									"rx_rate": &schema.Schema{
										Description: `Rx Rate`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"rx_retries": &schema.Schema{
										Description: `Rx Retries`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"ssid": &schema.Schema{
										Description: `Ssid`,
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
									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_ctrl_packets": &schema.Schema{
										Description: `Tx Ctrl Packets`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_data_packets": &schema.Schema{
										Description: `Tx Data Packets`,
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
									"tx_rate": &schema.Schema{
										Description: `Tx Rate`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"tx_unicast_data_packets": &schema.Schema{
										Description: `Tx Unicast Data Packets`,
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

func resourceIcapClientsIDStatsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	vvID := vID.(string)
	request1 := expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Sensors.RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(vvID, request1, &headerParams1)

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

	vItems1 := flattenSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceIcapClientsIDStatsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapClientsIDStatsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters {
	request := []dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters{}
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
		i := expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters{}
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

func expandRequestIcapClientsIDStatsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage {
	request := dnacentersdkgo.RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage{}
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

func flattenSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["client_mac"] = item.ClientMac
		respItem["ap_mac"] = item.ApMac
		respItem["radio_id"] = item.RadioID
		respItem["timestamp"] = item.Timestamp
		respItem["band"] = item.Band
		respItem["ssid"] = item.SSID
		respItem["rssi"] = item.Rssi
		respItem["snr"] = item.Snr
		respItem["tx_bytes"] = item.TxBytes
		respItem["rx_bytes"] = item.RxBytes
		respItem["rx_packets"] = item.RxPackets
		respItem["tx_packets"] = item.TxPackets
		respItem["rx_mgmt_packets"] = item.RxMgmtPackets
		respItem["tx_mgmt_packets"] = item.TxMgmtPackets
		respItem["rx_data_packets"] = item.RxDataPackets
		respItem["tx_data_packets"] = item.TxDataPackets
		respItem["tx_unicast_data_packets"] = item.TxUnicastDataPackets
		respItem["rx_ctrl_packets"] = item.RxCtrlPackets
		respItem["tx_ctrl_packets"] = item.TxCtrlPackets
		respItem["rx_retries"] = item.RxRetries
		respItem["rx_rate"] = item.RxRate
		respItem["tx_rate"] = item.TxRate
		respItem["client_ip"] = item.ClientIP
		respItems = append(respItems, respItem)
	}
	return respItems
}
