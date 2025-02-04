package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get network bug device by device id
`,

		ReadContext: dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bug_count": &schema.Schema{
										Description: `Number of bugs to which the network device is vulnerable
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"comments": &schema.Schema{
										Description: `More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_scan_time": &schema.Schema{
										Description: `Time at which the device was scanned
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"network_device_id": &schema.Schema{
										Description: `Id of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"scan_mode": &schema.Schema{
										Description: `'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"scan_status": &schema.Schema{
										Description: `'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkBugDeviceByDeviceID")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Compliance.GetNetworkBugDeviceByDeviceID(vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkBugDeviceByDeviceID", err,
				"Failure at GetNetworkBugDeviceByDeviceID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetNetworkBugDeviceByDeviceIDItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkBugDeviceByDeviceID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetNetworkBugDeviceByDeviceIDItems(items *dnacentersdkgo.ResponseComplianceGetNetworkBugDeviceByDeviceID) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["response"] = flattenComplianceGetNetworkBugDeviceByDeviceIDItemsResponse(item.Response)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceGetNetworkBugDeviceByDeviceIDItemsResponse(item *dnacentersdkgo.ResponseItemComplianceGetNetworkBugDeviceByDeviceIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["bug_count"] = item.BugCount
	respItem["scan_mode"] = item.ScanMode
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = item.Comments
	respItem["last_scan_time"] = item.LastScanTime

	return []map[string]interface{}{
		respItem,
	}

}
