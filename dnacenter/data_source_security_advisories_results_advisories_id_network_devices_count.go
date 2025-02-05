package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of security advisory network devices for the security advisory
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the security advisory
`,
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vScanMode, okScanMode := d.GetOk("scan_mode")
	vScanStatus, okScanStatus := d.GetOk("scan_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okScanMode {
			queryParams1.ScanMode = vScanMode.(string)
		}
		if okScanStatus {
			queryParams1.ScanStatus = vScanStatus.(string)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory", err,
				"Failure at GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryItem(item *dnacentersdkgo.ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
