package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

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
				Description: `deviceUuid path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"compliance_status": &schema.Schema{
							Description: `Compliance Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_uuid": &schema.Schema{
							Description: `Device Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"schedule_time": &schema.Schema{
							Description: `Schedule Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
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
		log.Printf("[DEBUG] Selected method 1: DeviceComplianceStatus")
		vvDeviceUUID := vDeviceUUID.(string)

		response1, restyResp1, err := client.Compliance.DeviceComplianceStatus(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceComplianceStatus", err,
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
	respItem["message"] = item.Message
	respItem["schedule_time"] = item.ScheduleTime
	respItem["last_update_time"] = item.LastUpdateTime
	return []map[string]interface{}{
		respItem,
	}
}
