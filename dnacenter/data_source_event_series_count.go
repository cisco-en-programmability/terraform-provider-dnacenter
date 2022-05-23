package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSeriesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get the Count of Published Notifications
`,

		ReadContext: dataSourceEventSeriesCountRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Description: `category query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"domain": &schema.Schema{
				Description: `domain query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End Time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"event_ids": &schema.Schema{
				Description: `eventIds query parameter. The registered EventId should be provided
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"source": &schema.Schema{
				Description: `source query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start Time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sub_domain": &schema.Schema{
				Description: `subDomain query parameter. Sub Domain
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventSeriesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEventIDs, okEventIDs := d.GetOk("event_ids")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vCategory, okCategory := d.GetOk("category")
	vType, okType := d.GetOk("type")
	vSeverity, okSeverity := d.GetOk("severity")
	vDomain, okDomain := d.GetOk("domain")
	vSubDomain, okSubDomain := d.GetOk("sub_domain")
	vSource, okSource := d.GetOk("source")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CountOfNotifications")
		queryParams1 := dnacentersdkgo.CountOfNotificationsQueryParams{}

		if okEventIDs {
			queryParams1.EventIDs = vEventIDs.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okCategory {
			queryParams1.Category = vCategory.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
		}
		if okDomain {
			queryParams1.Domain = vDomain.(string)
		}
		if okSubDomain {
			queryParams1.SubDomain = vSubDomain.(string)
		}
		if okSource {
			queryParams1.Source = vSource.(string)
		}

		response1, restyResp1, err := client.EventManagement.CountOfNotifications(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CountOfNotifications", err,
				"Failure at CountOfNotifications, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEventManagementCountOfNotificationsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountOfNotifications response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementCountOfNotificationsItem(item *dnacentersdkgo.ResponseEventManagementCountOfNotifications) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
