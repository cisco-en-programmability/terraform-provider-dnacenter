package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseLastOperationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Retrieves the status of the last system licensing operation.
`,

		ReadContext: dataSourceLicenseLastOperationStatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"error_code": &schema.Schema{
							Description: `An error code if in case this task has failed in its execution
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_reason": &schema.Schema{
							Description: `A textual description indicating the reason why a task has failed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `The ID of this task
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_error": &schema.Schema{
							Description: `A boolean indicating if this task has ended with or without error. true indicates a failure, whereas false indicates a success.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update": &schema.Schema{
							Description: `A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Summarizes the status of a task
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

func dataSourceLicenseLastOperationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SystemLicensingLastOperationStatus")

		response1, restyResp1, err := client.Licenses.SystemLicensingLastOperationStatus()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemLicensingLastOperationStatus", err,
				"Failure at SystemLicensingLastOperationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesSystemLicensingLastOperationStatusItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemLicensingLastOperationStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesSystemLicensingLastOperationStatusItem(item *dnacentersdkgo.ResponseLicensesSystemLicensingLastOperationStatusResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["status"] = item.Status
	respItem["is_error"] = boolPtrToString(item.IsError)
	respItem["failure_reason"] = item.FailureReason
	respItem["error_code"] = item.ErrorCode
	respItem["last_update"] = item.LastUpdate
	return []map[string]interface{}{
		respItem,
	}
}
