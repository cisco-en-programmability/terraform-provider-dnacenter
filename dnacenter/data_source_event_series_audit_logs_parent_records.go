package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSeriesAuditLogsParentRecords() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get Parent Audit Log Event instances from the Event-Hub
`,

		ReadContext: dataSourceEventSeriesAuditLogsParentRecordsRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Description: `category query parameter. Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"context": &schema.Schema{
				Description: `context query parameter. Audit Log notification's event correlationId.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Description: `description query parameter. String full/partial search (Provided input string is case insensitively matched for records).
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. Audit Log notification's deviceId.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Description: `domain query parameter. Audit Log notification's event domain.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"event_hierarchy": &schema.Schema{
				Description: `eventHierarchy query parameter. Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" Delimiter for hierarchy separation is ".".
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_id": &schema.Schema{
				Description: `eventId query parameter. Audit Log notification's event ID. 
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": &schema.Schema{
				Description: `instanceId query parameter. InstanceID of the Audit Log.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_system_events": &schema.Schema{
				Description: `isSystemEvents query parameter. Parameter to filter system generated audit-logs.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of Audit Log records to be returned per page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Audit Log notification event name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Position of a particular Audit Log record in the data. 
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order of the sorted Audit Log records. Default value is desc by timestamp. Supported values: asc, desc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Audit Log notification's siteId.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort the Audit Logs by certain fields. Supported values are event notification header attributes.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Description: `source query parameter. Audit Log notification's event source.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sub_domain": &schema.Schema{
				Description: `subDomain query parameter. Audit Log notification's event sub-domain.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Description: `userId query parameter. Audit Log notification's event userId.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_details": &schema.Schema{
							Description: `Additional Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"child_count": &schema.Schema{
							Description: `Child Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"cisco_dna_event_link": &schema.Schema{
							Description: `Cisco Dna Event Link`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"context": &schema.Schema{
							Description: `Context`,
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

						"i18n": &schema.Schema{
							Description: `I18n`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"message_params": &schema.Schema{
							Description: `Message Params`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network": &schema.Schema{
							Description: `Network`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"note": &schema.Schema{
							Description: `Note`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"parent_instance_id": &schema.Schema{
							Description: `Parent Instance Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"severity": &schema.Schema{
							Description: `Severity`,
							Type:        schema.TypeInt,
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

						"tags": &schema.Schema{
							Description: `Tags`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"tnt_id": &schema.Schema{
							Description: `Tnt Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"user_id": &schema.Schema{
							Description: `User Id`,
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

func dataSourceEventSeriesAuditLogsParentRecordsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInstanceID, okInstanceID := d.GetOk("instance_id")
	vName, okName := d.GetOk("name")
	vEventID, okEventID := d.GetOk("event_id")
	vCategory, okCategory := d.GetOk("category")
	vSeverity, okSeverity := d.GetOk("severity")
	vDomain, okDomain := d.GetOk("domain")
	vSubDomain, okSubDomain := d.GetOk("sub_domain")
	vSource, okSource := d.GetOk("source")
	vUserID, okUserID := d.GetOk("user_id")
	vContext, okContext := d.GetOk("context")
	vEventHierarchy, okEventHierarchy := d.GetOk("event_hierarchy")
	vSiteID, okSiteID := d.GetOk("site_id")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vIsSystemEvents, okIsSystemEvents := d.GetOk("is_system_events")
	vDescription, okDescription := d.GetOk("description")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAuditLogParentRecords")
		queryParams1 := dnacentersdkgo.GetAuditLogParentRecordsQueryParams{}

		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okEventID {
			queryParams1.EventID = vEventID.(string)
		}
		if okCategory {
			queryParams1.Category = vCategory.(string)
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
		if okUserID {
			queryParams1.UserID = vUserID.(string)
		}
		if okContext {
			queryParams1.Context = vContext.(string)
		}
		if okEventHierarchy {
			queryParams1.EventHierarchy = vEventHierarchy.(string)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okIsSystemEvents {
			queryParams1.IsSystemEvents = vIsSystemEvents.(bool)
		}
		if okDescription {
			queryParams1.Description = vDescription.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetAuditLogParentRecords(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuditLogParentRecords", err,
				"Failure at GetAuditLogParentRecords, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetAuditLogParentRecordsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuditLogParentRecords response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetAuditLogParentRecordsItems(items *dnacentersdkgo.ResponseEventManagementGetAuditLogParentRecords) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["version"] = item.Version
		respItem["instance_id"] = item.InstanceID
		respItem["event_id"] = item.EventID
		respItem["namespace"] = item.Namespace
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["type"] = item.Type
		respItem["category"] = item.Category
		respItem["domain"] = item.Domain
		respItem["sub_domain"] = item.SubDomain
		respItem["severity"] = item.Severity
		respItem["source"] = item.Source
		respItem["timestamp"] = item.Timestamp
		respItem["tags"] = flattenEventManagementGetAuditLogParentRecordsItemsTags(item.Tags)
		respItem["details"] = flattenEventManagementGetAuditLogParentRecordsItemsDetails(item.Details)
		respItem["cisco_dna_event_link"] = item.CiscoDnaEventLink
		respItem["note"] = item.Note
		respItem["tnt_id"] = item.TntID
		respItem["context"] = item.Context
		respItem["user_id"] = item.UserID
		respItem["i18n"] = item.I18N
		respItem["event_hierarchy"] = item.EventHierarchy
		respItem["message"] = item.Message
		respItem["message_params"] = item.MessageParams
		respItem["additional_details"] = flattenEventManagementGetAuditLogParentRecordsItemsAdditionalDetails(item.AdditionalDetails)
		respItem["parent_instance_id"] = item.ParentInstanceID
		respItem["network"] = item.Network
		respItem["child_count"] = item.ChildCount
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetAuditLogParentRecordsItemsTags(items *[]dnacentersdkgo.ResponseItemEventManagementGetAuditLogParentRecordsTags) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenEventManagementGetAuditLogParentRecordsItemsDetails(item *dnacentersdkgo.ResponseItemEventManagementGetAuditLogParentRecordsDetails) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEventManagementGetAuditLogParentRecordsItemsAdditionalDetails(item *dnacentersdkgo.ResponseItemEventManagementGetAuditLogParentRecordsAdditionalDetails) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
