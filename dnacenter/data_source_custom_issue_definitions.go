package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Retrieve the existing syslog-based custom issue definitions. The supported filters are id, name, profileId,
definition enable status, priority, severity, facility and mnemonic. The issue definition configurations may vary across
profiles, hence specifying the profile Id in the query parameter is important and the default profile is global.

  The assurance profile definitions can be obtain via the API endpoint: /api/v1/siteprofile?namespace=assurance. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceCustomIssueDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"facility": &schema.Schema{
				Description: `facility query parameter. The syslog facility name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. The custom issue definition identifier and unique identifier across the profile.Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_enabled": &schema.Schema{
				Description: `isEnabled query parameter. The enable status of the custom issue definition, either true or false.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mnemonic": &schema.Schema{
				Description: `mnemonic query parameter. The syslog mnemonic name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The list of UDI issue names
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Description: `priority query parameter. The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Description: `profileId query parameter. The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_time": &schema.Schema{
							Description: `Created Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_deletable": &schema.Schema{
							Description: `Is Deletable`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Is Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_notification_enabled": &schema.Schema{
							Description: `Is Notification Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"profile_id": &schema.Schema{
							Description: `Profile Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration_in_minutes": &schema.Schema{
										Description: `Duration In Minutes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"occurrences": &schema.Schema{
										Description: `Occurrences`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"pattern": &schema.Schema{
										Description: `Pattern`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"severity": &schema.Schema{
										Description: `Severity`,
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

						"trigger_id": &schema.Schema{
							Description: `Trigger Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCustomIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vProfileID, okProfileID := d.GetOk("profile_id")
	vName, okName := d.GetOk("name")
	vPriority, okPriority := d.GetOk("priority")
	vIsEnabled, okIsEnabled := d.GetOk("is_enabled")
	vSeverity, okSeverity := d.GetOk("severity")
	vFacility, okFacility := d.GetOk("facility")
	vMnemonic, okMnemonic := d.GetOk("mnemonic")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters")
		queryParams1 := dnacentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okProfileID {
			queryParams1.ProfileID = vProfileID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okIsEnabled {
			queryParams1.IsEnabled = vIsEnabled.(bool)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(float64)
		}
		if okFacility {
			queryParams1.Facility = vFacility.(string)
		}
		if okMnemonic {
			queryParams1.Mnemonic = vMnemonic.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Issues.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters", err,
				"Failure at GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItems(items *[]dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["profile_id"] = item.ProfileID
		respItem["trigger_id"] = item.TriggerID
		respItem["rules"] = flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItemsRules(item.Rules)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["priority"] = item.Priority
		respItem["is_deletable"] = boolPtrToString(item.IsDeletable)
		respItem["is_notification_enabled"] = boolPtrToString(item.IsNotificationEnabled)
		respItem["created_time"] = item.CreatedTime
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItemsRules(items *[]dnacentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["severity"] = item.Severity
		respItem["facility"] = item.Facility
		respItem["mnemonic"] = item.Mnemonic
		respItem["pattern"] = item.Pattern
		respItem["occurrences"] = item.Occurrences
		respItem["duration_in_minutes"] = item.DurationInMinutes
		respItems = append(respItems, respItem)
	}
	return respItems
}
