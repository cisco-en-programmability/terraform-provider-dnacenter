package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return compliance status of device(s).
`,

		ReadContext: dataSourceComplianceDeviceRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `deviceUuid query parameter. Comma separated 'Device Ids'
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"compliance_status": &schema.Schema{
							Description: `Current compliance status for the compliance type that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
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

						"message": &schema.Schema{
							Description: `Additional message of compliance status for the compliance type.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"schedule_time": &schema.Schema{
							Description: `Timestamp when compliance is scheduled to run.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceComplianceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")
	vDeviceUUID, okDeviceUUID := d.GetOk("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetComplianceStatus")
		queryParams1 := dnacentersdkgo.GetComplianceStatusQueryParams{}

		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}
		if okDeviceUUID {
			queryParams1.DeviceUUID = vDeviceUUID.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetComplianceStatus", err,
				"Failure at GetComplianceStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetComplianceStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceStatusItems(items *[]dnacentersdkgo.ResponseComplianceGetComplianceStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_uuid"] = item.DeviceUUID
		respItem["compliance_status"] = item.ComplianceStatus
		respItem["message"] = item.Message
		respItem["schedule_time"] = item.ScheduleTime
		respItem["last_update_time"] = item.LastUpdateTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
