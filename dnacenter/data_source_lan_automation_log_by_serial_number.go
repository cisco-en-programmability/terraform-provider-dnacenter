package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLanAutomationLogBySerialNumber() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation session logs for individual devices based on the given LAN Automation
session id and device serial number.
`,

		ReadContext: dataSourceLanAutomationLogBySerialNumberRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. LAN Automation session identifier.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"log_level": &schema.Schema{
				Description: `logLevel query parameter. Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber path parameter. Device serial number.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"logs": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"log_level": &schema.Schema{
										Description: `Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"record": &schema.Schema{
										Description: `Detailed log message.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"time_stamp": &schema.Schema{
										Description: `Time at which the log message is created.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"nw_orch_id": &schema.Schema{
							Description: `LAN Automation session identifier.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Device serial number for which the log messages are associated.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLanAutomationLogBySerialNumberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vSerialNumber := d.Get("serial_number")
	vLogLevel, okLogLevel := d.GetOk("log_level")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomationLogsForIndividualDevices")
		vvID := vID.(string)
		vvSerialNumber := vSerialNumber.(string)
		queryParams1 := dnacentersdkgo.LanAutomationLogsForIndividualDevicesQueryParams{}

		if okLogLevel {
			queryParams1.LogLevel = vLogLevel.(string)
		}

		response1, restyResp1, err := client.LanAutomation.LanAutomationLogsForIndividualDevices(vvID, vvSerialNumber, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LanAutomationLogsForIndividualDevices", err,
				"Failure at LanAutomationLogsForIndividualDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLanAutomationLanAutomationLogsForIndividualDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationLogsForIndividualDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLanAutomationLanAutomationLogsForIndividualDevicesItems(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nw_orch_id"] = item.NwOrchID
		respItem["logs"] = flattenLanAutomationLanAutomationLogsForIndividualDevicesItemsLogs(item.Logs)
		respItem["serial_number"] = item.SerialNumber
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationLogsForIndividualDevicesItemsLogs(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponseLogs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["log_level"] = item.LogLevel
		respItem["time_stamp"] = item.TimeStamp
		respItem["record"] = item.Record
		respItems = append(respItems, respItem)
	}
	return respItems
}
