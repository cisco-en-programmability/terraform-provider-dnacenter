package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSystemIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Get all system issue defintions. The supported filters are id, name, profileId and definition enable status. An issue
trigger definition can be different across the profile and device type. So, *profileId* and *deviceType* in the query
param is important and default is global profile and all device type. For detailed information about the usage of the
API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-
api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml

- Get system issue defintion for the given id. Definition includes all properties from IssueTriggerDefinition schema by
default. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSystemIssueDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. These are the attributes supported in system issue definitions response. By default, all properties are sent in response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_type": &schema.Schema{
				Description: `deviceType query parameter. These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Issue trigger definition id.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"issue_enabled": &schema.Schema{
				Description: `issueEnabled query parameter. The enablement status of the issue definition, either true or false.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The list of system defined issue names. (Ex."BGP_Down")
Examples:
name=BGP_Down (single entity uuid requested)
name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
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
				Description: `priority query parameter. Issue priority, possible values are P1, P2, P3, P4.
*P1*: A critical issue that needs immediate attention and can have a wide impact on network operations.
*P2*: A major issue that can potentially impact multiple devices or clients.
*P3*: A minor issue that has a localized or minimal impact.
*P4*: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Description: `profileId query parameter. The profile identier to fetch the profile associated issue defintions. The default is *global*. Please refer Network design profiles documentation for more details.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category_name": &schema.Schema{
							Description: `Category Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_priority": &schema.Schema{
							Description: `Default Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"definition_status": &schema.Schema{
							Description: `Definition Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"issue_enabled": &schema.Schema{
							Description: `Issue Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeString,
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

						"synchronize_to_health_threshold": &schema.Schema{
							Description: `Synchronize To Health Threshold`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"threshold_value": &schema.Schema{
							Description: `Threshold Value`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category_name": &schema.Schema{
							Description: `Category Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_priority": &schema.Schema{
							Description: `Default Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"definition_status": &schema.Schema{
							Description: `Definition Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_type": &schema.Schema{
							Description: `Device Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"issue_enabled": &schema.Schema{
							Description: `Issue Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeString,
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

						"synchronize_to_health_threshold": &schema.Schema{
							Description: `Synchronize To Health Threshold`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"threshold_value": &schema.Schema{
							Description: `Threshold Value`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vProfileID, okProfileID := d.GetOk("profile_id")
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vPriority, okPriority := d.GetOk("priority")
	vIssueEnabled, okIssueEnabled := d.GetOk("issue_enabled")
	vAttribute, okAttribute := d.GetOk("attribute")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")

	method1 := []bool{okDeviceType, okProfileID, okID, okName, okPriority, okIssueEnabled, okAttribute, okOffset, okLimit, okSortBy, okOrder, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsAllIssueTriggerDefinitionsForGivenFilters")

		headerParams1 := dnacentersdkgo.ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams{}
		queryParams1 := dnacentersdkgo.ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams{}

		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		if okProfileID {
			queryParams1.ProfileID = vProfileID.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okIssueEnabled {
			queryParams1.IssueEnabled = vIssueEnabled.(bool)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
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
		if okXCaLLERID {
			headerParams1.XCaLLERID = vXCaLLERID.(string)
		}

		response1, restyResp1, err := client.Issues.ReturnsAllIssueTriggerDefinitionsForGivenFilters(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReturnsAllIssueTriggerDefinitionsForGivenFilters", err,
				"Failure at ReturnsAllIssueTriggerDefinitionsForGivenFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsAllIssueTriggerDefinitionsForGivenFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetIssueTriggerDefinitionForGivenID")
		vvID := vID.(string)

		headerParams2 := dnacentersdkgo.GetIssueTriggerDefinitionForGivenIDHeaderParams{}

		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		response2, restyResp2, err := client.Issues.GetIssueTriggerDefinitionForGivenID(vvID, &headerParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetIssueTriggerDefinitionForGivenID", err,
				"Failure at GetIssueTriggerDefinitionForGivenID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIssuesGetIssueTriggerDefinitionForGivenIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIssueTriggerDefinitionForGivenID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersItems(items *[]dnacentersdkgo.ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["display_name"] = item.DisplayName
		respItem["description"] = item.Description
		respItem["priority"] = item.Priority
		respItem["default_priority"] = item.DefaultPriority
		respItem["device_type"] = item.DeviceType
		respItem["issue_enabled"] = boolPtrToString(item.IssueEnabled)
		respItem["profile_id"] = item.ProfileID
		respItem["definition_status"] = item.DefinitionStatus
		respItem["category_name"] = item.CategoryName
		respItem["synchronize_to_health_threshold"] = boolPtrToString(item.SynchronizeToHealthThreshold)
		respItem["threshold_value"] = item.ThresholdValue
		respItem["last_modified"] = item.LastModified
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetIssueTriggerDefinitionForGivenIDItem(item *dnacentersdkgo.ResponseIssuesGetIssueTriggerDefinitionForGivenIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["display_name"] = item.DisplayName
	respItem["description"] = item.Description
	respItem["priority"] = item.Priority
	respItem["default_priority"] = item.DefaultPriority
	respItem["device_type"] = item.DeviceType
	respItem["issue_enabled"] = boolPtrToString(item.IssueEnabled)
	respItem["profile_id"] = item.ProfileID
	respItem["definition_status"] = item.DefinitionStatus
	respItem["category_name"] = item.CategoryName
	respItem["synchronize_to_health_threshold"] = boolPtrToString(item.SynchronizeToHealthThreshold)
	respItem["threshold_value"] = item.ThresholdValue
	respItem["last_modified"] = item.LastModified
	return []map[string]interface{}{
		respItem,
	}
}
