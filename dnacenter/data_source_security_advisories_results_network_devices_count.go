package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of security advisory network devices
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"advisory_count": &schema.Schema{
				Description: `advisoryCount query parameter. Return network devices with advisoryCount greater than this advisoryCount
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
			"scan_mode": &schema.Schema{
				Description: `scanMode query parameter. Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_status": &schema.Schema{
				Description: `scanStatus query parameter. Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityAdvisoriesResultsNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vScanMode, okScanMode := d.GetOk("scan_mode")
	vScanStatus, okScanStatus := d.GetOk("scan_status")
	vAdvisoryCount, okAdvisoryCount := d.GetOk("advisory_count")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfSecurityAdvisoryNetworkDevices")
		queryParams1 := dnacentersdkgo.GetCountOfSecurityAdvisoryNetworkDevicesQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okScanMode {
			queryParams1.ScanMode = vScanMode.(string)
		}
		if okScanStatus {
			queryParams1.ScanStatus = vScanStatus.(string)
		}
		if okAdvisoryCount {
			queryParams1.AdvisoryCount = vAdvisoryCount.(float64)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfSecurityAdvisoryNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfSecurityAdvisoryNetworkDevices", err,
				"Failure at GetCountOfSecurityAdvisoryNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfSecurityAdvisoryNetworkDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfSecurityAdvisoryNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfSecurityAdvisoryNetworkDevicesItem(item *dnacentersdkgo.ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
