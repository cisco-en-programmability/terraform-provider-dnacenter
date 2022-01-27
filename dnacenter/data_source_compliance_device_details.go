package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return Compliance Detail
`,

		ReadContext: dataSourceComplianceDeviceDetailsRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. complianceType can have any value among 'NETWORK_PROFILE', 'IMAGE', 'APPLICATION_VISIBILITY', 'FABRIC', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"compliance_type": &schema.Schema{
							Description: `Compliance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_uuid": &schema.Schema{
							Description: `Device Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_sync_time": &schema.Schema{
							Description: `Last Sync Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"state": &schema.Schema{
							Description: `State`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceComplianceDeviceDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceType, okComplianceType := d.GetOk("compliance_type")
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")
	vDeviceUUID, okDeviceUUID := d.GetOk("device_uuid")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetComplianceDetail")
		queryParams1 := dnacentersdkgo.GetComplianceDetailQueryParams{}

		if okComplianceType {
			queryParams1.ComplianceType = vComplianceType.(string)
		}
		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}
		if okDeviceUUID {
			queryParams1.DeviceUUID = vDeviceUUID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetComplianceDetail", err,
				"Failure at GetComplianceDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetComplianceDetailItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceDetail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceDetailItems(items *[]dnacentersdkgo.ResponseComplianceGetComplianceDetailResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["compliance_type"] = item.ComplianceType
		respItem["last_sync_time"] = item.LastSyncTime
		respItem["device_uuid"] = item.DeviceUUID
		respItem["display_name"] = item.DisplayName
		respItem["status"] = item.Status
		respItem["category"] = item.Category
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["state"] = item.State
		respItems = append(respItems, respItem)
	}
	return respItems
}
