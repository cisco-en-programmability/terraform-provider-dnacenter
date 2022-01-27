package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEvent() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of registered Events with provided eventIds or tags as mandatory
`,

		ReadContext: dataSourceEventRead,
		Schema: map[string]*schema.Schema{
			"event_id": &schema.Schema{
				Description: `eventId query parameter. The registered EventId should be provided
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of Registries to limit in the resultset whose default value 10
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The number of Registries to offset in the resultset whose default value 0
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. SortBy field name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Description: `tags query parameter. The registered Tags should be provided
`,
				Type:     schema.TypeString,
				Required: true,
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

						"details": &schema.Schema{
							Description: `Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"domain": &schema.Schema{
							Description: `Domain`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"event_id": &schema.Schema{
							Description: `Event Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name_space": &schema.Schema{
							Description: `Name Space`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"sub_domain": &schema.Schema{
							Description: `Sub Domain`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"subscription_types": &schema.Schema{
							Description: `Subscription Types`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"tags": &schema.Schema{
							Description: `Tags`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEventID, okEventID := d.GetOk("event_id")
	vTags := d.Get("tags")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEvents")
		queryParams1 := dnacentersdkgo.GetEventsQueryParams{}

		if okEventID {
			queryParams1.EventID = vEventID.(string)
		}
		queryParams1.Tags = vTags.(string)

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetEvents(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEvents", err,
				"Failure at GetEvents, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetEventsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEvents response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetEventsItems(items *dnacentersdkgo.ResponseEventManagementGetEvents) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["event_id"] = item.EventID
		respItem["name_space"] = item.NameSpace
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["version"] = item.Version
		respItem["category"] = item.Category
		respItem["domain"] = item.Domain
		respItem["sub_domain"] = item.SubDomain
		respItem["type"] = item.Type
		respItem["tags"] = item.Tags
		respItem["severity"] = item.Severity
		respItem["details"] = flattenEventManagementGetEventsItemsDetails(item.Details)
		respItem["subscription_types"] = item.SubscriptionTypes
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetEventsItemsDetails(item *dnacentersdkgo.ResponseItemEventManagementGetEventsDetails) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
