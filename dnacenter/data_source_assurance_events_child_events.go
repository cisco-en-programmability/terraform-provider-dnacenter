package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceEventsChildEvents() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Wireless client event could have child events and this API can be used to fetch the same using parent event *id* as
the input. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceEvents-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceEventsChildEventsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Unique identifier for the event
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"details": &schema.Schema{
							Description: `Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_category": &schema.Schema{
							Description: `Failure Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reason_code": &schema.Schema{
							Description: `Reason Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reason_description": &schema.Schema{
							Description: `Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"result_status": &schema.Schema{
							Description: `Result Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sub_reason_description": &schema.Schema{
							Description: `Sub Reason Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"subreason_code": &schema.Schema{
							Description: `Subreason Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"wireless_event_type": &schema.Schema{
							Description: `Wireless Event Type`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAssuranceEventsChildEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetListOfChildEventsForTheGivenWirelessClientEvent")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams{}

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Devices.GetListOfChildEventsForTheGivenWirelessClientEvent(vvID, &headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetListOfChildEventsForTheGivenWirelessClientEvent", err,
				"Failure at GetListOfChildEventsForTheGivenWirelessClientEvent, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetListOfChildEventsForTheGivenWirelessClientEventItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfChildEventsForTheGivenWirelessClientEvent response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetListOfChildEventsForTheGivenWirelessClientEventItems(items *[]dnacentersdkgo.ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["timestamp"] = item.Timestamp
		respItem["wireless_event_type"] = item.WirelessEventType
		respItem["details"] = item.Details
		respItem["reason_code"] = item.ReasonCode
		respItem["subreason_code"] = item.SubreasonCode
		respItem["result_status"] = item.ResultStatus
		respItem["reason_description"] = item.ReasonDescription
		respItem["sub_reason_description"] = item.SubReasonDescription
		respItem["failure_category"] = item.FailureCategory
		respItems = append(respItems, respItem)
	}
	return respItems
}
