package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSeries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get the list of Published Notifications
`,

		ReadContext: dataSourceEventSeriesRead,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. # of records
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Start Offset
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Ascending/Descending order [asc/desc]
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort By column
`,
				Type:     schema.TypeString,
				Optional: true,
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

						"event_hierarchy": &schema.Schema{
							Description: `Event Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"name_space": &schema.Schema{
							Description: `Name Space`,
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
							Type:        schema.TypeString,
							Computed:    true,
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

func dataSourceEventSeriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetNotifications")
		queryParams1 := dnacentersdkgo.GetNotificationsQueryParams{}

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

		response1, restyResp1, err := client.EventManagement.GetNotifications(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNotifications", err,
				"Failure at GetNotifications, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetNotificationsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNotifications response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetNotificationsItems(items *dnacentersdkgo.ResponseEventManagementGetNotifications) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["event_id"] = item.EventID
		respItem["instance_id"] = item.InstanceID
		respItem["name_space"] = item.NameSpace
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["version"] = item.Version
		respItem["category"] = item.Category
		respItem["domain"] = item.Domain
		respItem["sub_domain"] = item.SubDomain
		respItem["type"] = item.Type
		respItem["severity"] = item.Severity
		respItem["source"] = item.Source
		respItem["timestamp"] = item.Timestamp
		respItem["details"] = item.Details
		respItem["event_hierarchy"] = item.EventHierarchy
		respItems = append(respItems, respItem)
	}
	return respItems
}
