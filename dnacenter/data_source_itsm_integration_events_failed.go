package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceItsmIntegrationEventsFailed() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM.

- Used to retrieve the list of integration events that failed to create tickets in ITSM
`,

		ReadContext: dataSourceItsmIntegrationEventsFailedRead,
		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Description: `instanceId query parameter. Instance Id of the failed event as in the Runtime Dashboard
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

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"domain": &schema.Schema{
							Description: `Domain`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"enrichment_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"error_code": &schema.Schema{
										Description: `Error Code`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"error_description": &schema.Schema{
										Description: `Error Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"event_status": &schema.Schema{
										Description: `Event Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"response_received_from_itsmsystem": &schema.Schema{
										Description: `Response Received From ITSMSystem`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"event_id": &schema.Schema{
							Description: `Event Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"source": &schema.Schema{
							Description: `Source`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sub_domain": &schema.Schema{
							Description: `Sub Domain`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceItsmIntegrationEventsFailedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInstanceID, okInstanceID := d.GetOk("instance_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetFailedItsmEvents")
		queryParams1 := dnacentersdkgo.GetFailedItsmEventsQueryParams{}

		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
		}

		response1, restyResp1, err := client.Itsm.GetFailedItsmEvents(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetFailedItsmEvents", err,
				"Failure at GetFailedItsmEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenItsmGetFailedItsmEventsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFailedItsmEvents response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmGetFailedItsmEventsItems(items *dnacentersdkgo.ResponseItsmGetFailedItsmEvents) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_id"] = item.InstanceID
		respItem["event_id"] = item.EventID
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["category"] = item.Category
		respItem["domain"] = item.Domain
		respItem["sub_domain"] = item.SubDomain
		respItem["severity"] = item.Severity
		respItem["source"] = item.Source
		respItem["timestamp"] = item.Timestamp
		respItem["enrichment_info"] = flattenItsmGetFailedItsmEventsItemsEnrichmentInfo(item.EnrichmentInfo)
		respItem["description"] = item.Description
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenItsmGetFailedItsmEventsItemsEnrichmentInfo(item *dnacentersdkgo.ResponseItemItsmGetFailedItsmEventsEnrichmentInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["event_status"] = item.EventStatus
	respItem["error_code"] = item.ErrorCode
	respItem["error_description"] = item.ErrorDescription
	respItem["response_received_from_itsmsystem"] = flattenItsmGetFailedItsmEventsItemsEnrichmentInfoResponseReceivedFromITSmsystem(item.ResponseReceivedFromITSmsystem)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenItsmGetFailedItsmEventsItemsEnrichmentInfoResponseReceivedFromITSmsystem(item *dnacentersdkgo.ResponseItemItsmGetFailedItsmEventsEnrichmentInfoResponseReceivedFromITSmsystem) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
