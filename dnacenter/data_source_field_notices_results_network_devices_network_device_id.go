package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get field notice network device by device id
`,

		ReadContext: dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceIDRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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

						"notice_count": &schema.Schema{
							Description: `Number of field notices to which the network device is vulnerable
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"potential_notice_count": &schema.Schema{
							Description: `Number of potential field notices to which the network device is vulnerable
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"scan_status": &schema.Schema{
							Description: `'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
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

func dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFieldNoticeNetworkDeviceByDeviceID")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Compliance.GetFieldNoticeNetworkDeviceByDeviceID(vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFieldNoticeNetworkDeviceByDeviceID", err,
				"Failure at GetFieldNoticeNetworkDeviceByDeviceID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetFieldNoticeNetworkDeviceByDeviceIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFieldNoticeNetworkDeviceByDeviceID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetFieldNoticeNetworkDeviceByDeviceIDItem(item *dnacentersdkgo.ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["notice_count"] = item.NoticeCount
	respItem["potential_notice_count"] = item.PotentialNoticeCount
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = item.Comments
	respItem["last_scan_time"] = item.LastScanTime
	return []map[string]interface{}{
		respItem,
	}
}
