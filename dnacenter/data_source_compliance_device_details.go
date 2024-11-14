package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

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
				Description: `complianceStatus query parameter. Specify "Compliance status(es)" in commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. Specify "Compliance type(s)" in commas. The Compliance type can be 'NETWORK_PROFILE', 'IMAGE', 'FABRIC', 'APPLICATION_VISIBILITY', 'FABRIC', RUNNING_CONFIG', 'NETWORK_SETTINGS', 'WORKFLOW' , 'EoX'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `deviceUuid query parameter. Comma separated "Device Id(s)"
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

						"category": &schema.Schema{
							Description: `category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EoX' , 'NETWORK_SETTINGS'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"compliance_type": &schema.Schema{
							Description: `Compliance type corresponds to a tile on the UI. Will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EoX.
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

						"display_name": &schema.Schema{
							Description: `User friendly name for the configuration.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_sync_time": &schema.Schema{
							Description: `Timestamp when the status changed from different value to the current value.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Timestamp when the latest compliance checks ran.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"state": &schema.Schema{
							Description: `State of latest compliance check for the complianceType. Will be one of SUCCESS, FAILED, or IN_PROGRESS.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Current status of compliance for the complianceType. Will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
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
		log.Printf("[DEBUG] Selected method: GetComplianceDetail")
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
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceDetail(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetComplianceDetail", err,
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
