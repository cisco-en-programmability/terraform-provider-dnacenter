package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesIDsPerDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Security Advisories.

- Retrieves advisory device details for a device
`,

		ReadContext: dataSourceSecurityAdvisoriesIDsPerDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device instance UUID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advisory_ids": &schema.Schema{
							Description: `Advisories detected on the network device
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"comments": &schema.Schema{
							Description: `More details about the scan status. Ie:- if the scan status is failed, comments will give the reason for failure
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `Network device ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hidden_advisory_count": &schema.Schema{
							Description: `Number of advisories detected on the network device that were suppressed by the user
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"last_scan_time": &schema.Schema{
							Description: `Time at which the network device was scanned. The representation is unix time.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"scan_mode": &schema.Schema{
							Description: `Criteria on which the network device was scanned
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scan_status": &schema.Schema{
							Description: `Status of the scan performed on the network device
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

func dataSourceSecurityAdvisoriesIDsPerDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdvisoryDeviceDetail")
		vvDeviceID := vDeviceID.(string)

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoryDeviceDetail(vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoryDeviceDetail", err,
				"Failure at GetAdvisoryDeviceDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSecurityAdvisoriesGetAdvisoryDeviceDetailItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoryDeviceDetail response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoryDeviceDetailItem(item *dnacentersdkgo.ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_id"] = item.DeviceID
	respItem["advisory_ids"] = item.AdvisoryIDs
	respItem["hidden_advisory_count"] = item.HiddenAdvisoryCount
	respItem["scan_mode"] = item.ScanMode
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = item.Comments
	respItem["last_scan_time"] = item.LastScanTime
	return []map[string]interface{}{
		respItem,
	}
}
