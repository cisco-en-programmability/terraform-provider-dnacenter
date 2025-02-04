package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntentCustomIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Get the custom issue definition for the given custom issue definition Id. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIntentCustomIssueDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Get the custom issue definition for the given custom issue definition Id.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceIntentCustomIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID")
		vvID := vID.(string)

		headerParams1 := dnacentersdkgo.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDHeaderParams{}

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		response1, restyResp1, err := client.Issues.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID(vvID, &headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID", err,
				"Failure at GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIssuesGetTheIntentCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetTheIntentCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItem(item *dnacentersdkgo.ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["profile_id"] = item.ProfileID
	respItem["trigger_id"] = item.TriggerID
	respItem["rules"] = flattenIssuesGetTheIntentCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItemRules(item.Rules)
	respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
	respItem["priority"] = item.Priority
	respItem["is_deletable"] = boolPtrToString(item.IsDeletable)
	respItem["is_notification_enabled"] = boolPtrToString(item.IsNotificationEnabled)
	respItem["created_time"] = item.CreatedTime
	respItem["last_updated_time"] = item.LastUpdatedTime
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIssuesGetTheIntentCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItemRules(items *[]dnacentersdkgo.ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponseRules) []map[string]interface{} {
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
