package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesNetworkDeviceID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get security advisory network device for the security advisory by network device id
`,

		ReadContext: dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesNetworkDeviceIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the security advisory
`,
				Type:     schema.TypeString,
				Required: true,
			},
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

						"advisory_count": &schema.Schema{
							Description: `Number of advisories to which the network device is vulnerable
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
							Description: `'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scan_status": &schema.Schema{
							Description: `'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
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

func dataSourceSecurityAdvisoriesResultsAdvisoriesIDNetworkDevicesNetworkDeviceIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vNetworkDeviceID := d.Get("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID")
		vvID := vID.(string)
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response1, restyResp1, err := client.Compliance.GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID(vvID, vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID", err,
				"Failure at GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDItem(item *dnacentersdkgo.ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["advisory_count"] = item.AdvisoryCount
	respItem["scan_mode"] = item.ScanMode
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = item.Comments
	respItem["last_scan_time"] = item.LastScanTime
	return []map[string]interface{}{
		respItem,
	}
}
