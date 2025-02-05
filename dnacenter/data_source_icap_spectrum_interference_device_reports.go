package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSpectrumInterferenceDeviceReports() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Retrieves the spectrum interference devices reports sent by WLC for provided AP Mac. For detailed information about
the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSpectrumInterferenceDeviceReportsRead,
		Schema: map[string]*schema.Schema{
			"ap_mac": &schema.Schema{
				Description: `apMac query parameter. The base ethernet macAddress of the access point
`,
				Type:     schema.TypeString,
				Required: true,
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

						"band_width_k_hz": &schema.Schema{
							Description: `Band Width K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"central_frequency_k_hz": &schema.Schema{
							Description: `Central Frequency K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"detected_channels": &schema.Schema{
							Description: `Detected Channels`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duty_cycle": &schema.Schema{
							Description: `Duty Cycle`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"high_end_frequency_k_hz": &schema.Schema{
							Description: `High End Frequency K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"low_end_frequency_k_hz": &schema.Schema{
							Description: `Low End Frequency K Hz`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"power_dbm": &schema.Schema{
							Description: `Power Dbm`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"severity_index": &schema.Schema{
							Description: `Severity Index`,
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

func dataSourceIcapSpectrumInterferenceDeviceReportsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vApMac := d.Get("ap_mac")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vTimeSortOrder, okTimeSortOrder := d.GetOk("time_sort_order")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac")

		headerParams1 := dnacentersdkgo.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacHeaderParams{}
		queryParams1 := dnacentersdkgo.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		queryParams1.ApMac = vApMac.(string)

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

		response1, restyResp1, err := client.Sensors.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac", err,
				"Failure at RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacItems(items *[]dnacentersdkgo.ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ap_mac"] = item.ApMac
		respItem["central_frequency_k_hz"] = item.CentralFrequencyKHz
		respItem["band_width_k_hz"] = item.BandWidthKHz
		respItem["low_end_frequency_k_hz"] = item.LowEndFrequencyKHz
		respItem["high_end_frequency_k_hz"] = item.HighEndFrequencyKHz
		respItem["power_dbm"] = item.PowerDbm
		respItem["band"] = item.Band
		respItem["duty_cycle"] = item.DutyCycle
		respItem["timestamp"] = item.Timestamp
		respItem["device_type"] = item.DeviceType
		respItem["severity_index"] = item.SeverityIndex
		respItem["detected_channels"] = item.DetectedChannels
		respItems = append(respItems, respItem)
	}
	return respItems
}
