package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNoticesIDNetworkDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get field notice network devices for the notice
`,

		ReadContext: dataSourceFieldNoticesResultsNoticesIDNetworkDevicesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the field notice
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Default value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_status": &schema.Schema{
				Description: `scanStatus query parameter. Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
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

func dataSourceFieldNoticesResultsNoticesIDNetworkDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vScanStatus, okScanStatus := d.GetOk("scan_status")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFieldNoticeNetworkDevicesForTheNotice")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetFieldNoticeNetworkDevicesForTheNoticeQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okScanStatus {
			queryParams1.ScanStatus = vScanStatus.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Compliance.GetFieldNoticeNetworkDevicesForTheNotice(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFieldNoticeNetworkDevicesForTheNotice", err,
				"Failure at GetFieldNoticeNetworkDevicesForTheNotice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetFieldNoticeNetworkDevicesForTheNoticeItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFieldNoticeNetworkDevicesForTheNotice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetFieldNoticeNetworkDevicesForTheNoticeItems(items *[]dnacentersdkgo.ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["notice_count"] = item.NoticeCount
		respItem["potential_notice_count"] = item.PotentialNoticeCount
		respItem["scan_status"] = item.ScanStatus
		respItem["comments"] = item.Comments
		respItem["last_scan_time"] = item.LastScanTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
