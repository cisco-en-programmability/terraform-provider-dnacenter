package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

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
				Description: `complianceStatus query parameter. Compliance status can be have value among 'COMPLIANT','NON_COMPLIANT','IN_PROGRESS', 'ERROR'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `deviceUuid query parameter. Comma separated deviceUuids
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of records to be retrieved
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
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

func dataSourceComplianceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")
	vDeviceUUID, okDeviceUUID := d.GetOk("device_uuid")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetComplianceStatus")
		queryParams1 := dnacentersdkgo.GetComplianceStatusQueryParams{}

		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}
		if okDeviceUUID {
			queryParams1.DeviceUUID = vDeviceUUID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetComplianceStatus", err,
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
