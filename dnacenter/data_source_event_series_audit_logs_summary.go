package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSeriesAuditLogsSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get Audit Log Summary from the Event-Hub
`,

		ReadContext: dataSourceEventSeriesAuditLogsSummaryRead,
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
			"is_parent_only": &schema.Schema{
				Description: `isParentOnly query parameter. Parameter to filter parent only audit-logs.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_system_events": &schema.Schema{
				Description: `isSystemEvents query parameter. Parameter to filter system generated audit-logs.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Audit Log notification event name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_instance_id": &schema.Schema{
				Description: `parentInstanceId query parameter. Parent Audit Log record's instanceID.
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

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"max_timestamp": &schema.Schema{
							Description: `Max Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"min_timestamp": &schema.Schema{
							Description: `Min Timestamp`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventSeriesAuditLogsSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vParentInstanceID, okParentInstanceID := d.GetOk("parent_instance_id")
	vIsParentOnly, okIsParentOnly := d.GetOk("is_parent_only")
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
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetAuditLogSummary")
		queryParams1 := dnacentersdkgo.GetAuditLogSummaryQueryParams{}

		if okParentInstanceID {
			queryParams1.ParentInstanceID = vParentInstanceID.(string)
		}
		if okIsParentOnly {
			queryParams1.IsParentOnly = vIsParentOnly.(bool)
		}
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
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}

		response1, restyResp1, err := client.EventManagement.GetAuditLogSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetAuditLogSummary", err,
				"Failure at GetAuditLogSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetAuditLogSummaryItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuditLogSummary response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetAuditLogSummaryItems(items *dnacentersdkgo.ResponseEventManagementGetAuditLogSummary) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["count"] = item.Count
		respItem["max_timestamp"] = item.MaxTimestamp
		respItem["min_timestamp"] = item.MinTimestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}
