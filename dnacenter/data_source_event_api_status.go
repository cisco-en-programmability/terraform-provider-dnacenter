package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventAPIStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get the Status of events API calls with provided executionId as mandatory path parameter
`,

		ReadContext: dataSourceEventAPIStatusRead,
		Schema: map[string]*schema.Schema{
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. Execution ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_status": &schema.Schema{
							Description: `Api Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"error_message": &schema.Schema{
							Description: `Error Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status_message": &schema.Schema{
							Description: `Status Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventAPIStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vExecutionID := d.Get("execution_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetStatusAPIForEvents")
		vvExecutionID := vExecutionID.(string)

		response1, restyResp1, err := client.EventManagement.GetStatusAPIForEvents(vvExecutionID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetStatusAPIForEvents", err,
				"Failure at GetStatusAPIForEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEventManagementGetStatusAPIForEventsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetStatusAPIForEvents response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetStatusAPIForEventsItem(item *dnacentersdkgo.ResponseEventManagementGetStatusAPIForEvents) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_message"] = flattenEventManagementGetStatusAPIForEventsItemErrorMessage(item.ErrorMessage)
	respItem["api_status"] = item.APIStatus
	respItem["status_message"] = item.StatusMessage
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEventManagementGetStatusAPIForEventsItemErrorMessage(item *dnacentersdkgo.ResponseEventManagementGetStatusAPIForEventsErrorMessage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
