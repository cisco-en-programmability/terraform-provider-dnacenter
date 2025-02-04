package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceByID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return compliance status of a device.
`,

		ReadContext: dataSourceComplianceDeviceByIDRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. Device Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"compliance_status": &schema.Schema{
							Description: `Current compliance status of the device that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_uuid": &schema.Schema{
							Description: `UUID of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Timestamp when the latest compliance checks ran.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"schedule_time": &schema.Schema{
							Description: `Timestamp when the next compliance checks will run.
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

func dataSourceComplianceDeviceByIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DeviceComplianceStatus")
		vvDeviceUUID := vDeviceUUID.(string)

		response1, restyResp1, err := client.Compliance.DeviceComplianceStatus(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 DeviceComplianceStatus", err,
				"Failure at DeviceComplianceStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceDeviceComplianceStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeviceComplianceStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceDeviceComplianceStatusItem(item *dnacentersdkgo.ResponseComplianceDeviceComplianceStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_uuid"] = item.DeviceUUID
	respItem["compliance_status"] = item.ComplianceStatus
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["schedule_time"] = item.ScheduleTime
	return []map[string]interface{}{
		respItem,
	}
}
