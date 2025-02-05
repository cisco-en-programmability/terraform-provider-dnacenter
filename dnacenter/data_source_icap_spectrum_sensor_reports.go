package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSpectrumSensorReports() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves the spectrum sensor reports sent by WLC for provided AP Mac. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSpectrumSensorReportsRead,
		Schema: map[string]*schema.Schema{
			"ap_mac": &schema.Schema{
				Description: `apMac query parameter. The base ethernet macAddress of the access point
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"data_type": &schema.Schema{
				Description: `dataType query parameter. Data type reported by the sensor
|Data Type | Description | | --| --| | *0* | Duty Cycle | | *1* | Max Power | | *2* | Average Power | | *3* | Max Power in dBm with adjusted base of +48 | | *4* | Average Power in dBm with adjusted base of +48 |
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"time_sort_order": &schema.Schema{
				Description: `timeSortOrder query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_mac": &schema.Schema{
							Description: `Ap Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"band": &schema.Schema{
							Description: `Band`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"central_frequency_k_hz": &schema.Schema{
							Description: `Central Frequency K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"channels": &schema.Schema{
							Description: `Channels`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						"data": &schema.Schema{
							Description: `Data`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeFloat,
							},
						},

						"data_avg": &schema.Schema{
							Description: `Data Avg`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"data_max": &schema.Schema{
							Description: `Data Max`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"data_min": &schema.Schema{
							Description: `Data Min`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"data_size": &schema.Schema{
							Description: `Data Size`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"data_type": &schema.Schema{
							Description: `Data Type`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"data_units": &schema.Schema{
							Description: `Data Units`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"span_k_hz": &schema.Schema{
							Description: `Span K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSpectrumSensorReportsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vApMac := d.Get("ap_mac")
	vDataType, okDataType := d.GetOk("data_type")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vTimeSortOrder, okTimeSortOrder := d.GetOk("time_sort_order")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac")

		headerParams1 := dnacentersdkgo.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		queryParams1.ApMac = vApMac.(string)

		if okDataType {
			queryParams1.DataType = vDataType.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okTimeSortOrder {
			queryParams1.TimeSortOrder = vTimeSortOrder.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Sensors.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac", err,
				"Failure at RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["span_k_hz"] = item.SpanKHz
		respItem["data_type"] = item.DataType
		respItem["ap_mac"] = item.ApMac
		respItem["data_avg"] = item.DataAvg
		respItem["data_min"] = item.DataMin
		respItem["data_max"] = item.DataMax
		respItem["data_units"] = item.DataUnits
		respItem["central_frequency_k_hz"] = item.CentralFrequencyKHz
		respItem["band"] = item.Band
		respItem["timestamp"] = item.Timestamp
		respItem["data"] = item.Data
		respItem["data_size"] = item.DataSize
		respItem["channels"] = item.Channels
		respItems = append(respItems, respItem)
	}
	return respItems
}
