package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIseIntegrationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on System Settings.

- API to check Cisco ISE server integration status.
`,

		ReadContext: dataSourceIseIntegrationStatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_server_setting_id": &schema.Schema{
							Description: `Cisco ISE Server setting identifier (E.g. 867e46c9-f8f5-40b1-8de2-62f7744f75f6)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"overall_error_message": &schema.Schema{
							Description: `Cisco ISE Server integration failure message
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"overall_status": &schema.Schema{
							Description: `Cisco ISE Server integration status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"steps": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cert_accepted_by_user": &schema.Schema{
										Description: `If user accept Cisco ISE Server certificate, value will be true otherwise it will be false
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"step_description": &schema.Schema{
										Description: `Cisco ISE Server step description
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"step_id": &schema.Schema{
										Description: `Cisco ISE Server integration step identifier (E.g. 1)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"step_name": &schema.Schema{
										Description: `Cisco ISE Server integration step name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"step_order": &schema.Schema{
										Description: `Cisco ISE Server integration step order (E.g. 1)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"step_status": &schema.Schema{
										Description: `Cisco ISE Server integration step status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"step_time": &schema.Schema{
										Description: `Last updated epoc time  by the step (E.g. 1677745739314)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIseIntegrationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CiscoIseServerIntegrationStatus")

		response1, restyResp1, err := client.SystemSettings.CiscoIseServerIntegrationStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CiscoIseServerIntegrationStatus", err,
				"Failure at CiscoIseServerIntegrationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsCiscoIseServerIntegrationStatusItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CiscoIseServerIntegrationStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSystemSettingsCiscoIseServerIntegrationStatusItem(item *dnacentersdkgo.ResponseSystemSettingsCiscoIseServerIntegrationStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_server_setting_id"] = item.AAAServerSettingID
	respItem["overall_status"] = item.OverallStatus
	respItem["overall_error_message"] = item.OverallErrorMessage
	respItem["steps"] = flattenSystemSettingsCiscoIseServerIntegrationStatusItemSteps(item.Steps)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSystemSettingsCiscoIseServerIntegrationStatusItemSteps(items *[]dnacentersdkgo.ResponseSystemSettingsCiscoIseServerIntegrationStatusSteps) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["step_id"] = item.StepID
		respItem["step_order"] = item.StepOrder
		respItem["step_name"] = item.StepName
		respItem["step_description"] = item.StepDescription
		respItem["step_status"] = item.StepStatus
		respItem["cert_accepted_by_user"] = boolPtrToString(item.CertAcceptedByUser)
		respItem["step_time"] = item.StepTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
