package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessAccessPointsFactoryResetRequestStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source returns each AP Factory Reset initiation status.
`,

		ReadContext: dataSourceWirelessAccessPointsFactoryResetRequestStatusRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId query parameter. provide the task id which is returned in the response of ap factory reset post api
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_response_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_factory_reset_status": &schema.Schema{
										Description: `AP factory reset status, "Success" or "Failure" or "In Progress"
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ap_name": &schema.Schema{
										Description: `Access Point name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ethernet_mac_address": &schema.Schema{
										Description: `AP Ethernet Mac Address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Reason for failure if the factory reset status is "Failure"
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_mac_address": &schema.Schema{
										Description: `AP Radio Mac Address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wlc_ip": &schema.Schema{
							Description: `Wireless Controller IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlc_name": &schema.Schema{
							Description: `Wireless Controller name
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

func dataSourceWirelessAccessPointsFactoryResetRequestStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointsFactoryResetStatus")
		queryParams1 := dnacentersdkgo.GetAccessPointsFactoryResetStatusQueryParams{}

		queryParams1.TaskID = vTaskID.(string)

		response1, restyResp1, err := client.Wireless.GetAccessPointsFactoryResetStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointsFactoryResetStatus", err,
				"Failure at GetAccessPointsFactoryResetStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetAccessPointsFactoryResetStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointsFactoryResetStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointsFactoryResetStatusItems(items *[]dnacentersdkgo.ResponseWirelessGetAccessPointsFactoryResetStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["wlc_ip"] = item.WlcIP
		respItem["wlc_name"] = item.WlcName
		respItem["ap_response_info_list"] = flattenWirelessGetAccessPointsFactoryResetStatusItemsApResponseInfoList(item.ApResponseInfoList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointsFactoryResetStatusItemsApResponseInfoList(items *[]dnacentersdkgo.ResponseWirelessGetAccessPointsFactoryResetStatusResponseApResponseInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_name"] = item.ApName
		respItem["ap_factory_reset_status"] = item.ApFactoryResetStatus
		respItem["failure_reason"] = item.FailureReason
		respItem["radio_mac_address"] = item.RadioMacAddress
		respItem["ethernet_mac_address"] = item.EthernetMacAddress
		respItems = append(respItems, respItem)
	}
	return respItems
}
