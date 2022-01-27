package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIssuesEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Enriches a given network issue context (an issue id or end user’s Mac Address) with details about the issue(s),
impacted hosts and suggested actions for remediation
`,

		ReadContext: dataSourceIssuesEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. Issue enrichment details can be fetched based on either Issue ID or Client MAC address. This parameter value must either be issue_id/mac_address
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_value": &schema.Schema{
				Description: `entity_value header parameter. Contains the actual value for the entity type that has been defined
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"issue": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"impacted_hosts": &schema.Schema{
										Description: `Impacted Hosts`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"issue_category": &schema.Schema{
										Description: `Issue Category`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_description": &schema.Schema{
										Description: `Issue Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_entity": &schema.Schema{
										Description: `Issue Entity`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_entity_value": &schema.Schema{
										Description: `Issue Entity Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_id": &schema.Schema{
										Description: `Issue Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_name": &schema.Schema{
										Description: `Issue Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_priority": &schema.Schema{
										Description: `Issue Priority`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_severity": &schema.Schema{
										Description: `Issue Severity`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_source": &schema.Schema{
										Description: `Issue Source`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_summary": &schema.Schema{
										Description: `Issue Summary`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_timestamp": &schema.Schema{
										Description: `Issue Timestamp`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"suggested_actions": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Description: `Message`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"steps": &schema.Schema{
													Description: `Steps`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIssuesEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetIssueEnrichmentDetails")

		headerParams1 := dnacentersdkgo.GetIssueEnrichmentDetailsHeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		response1, restyResp1, err := client.Issues.GetIssueEnrichmentDetails(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetIssueEnrichmentDetails", err,
				"Failure at GetIssueEnrichmentDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenIssuesGetIssueEnrichmentDetailsItem(response1.IssueDetails)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIssueEnrichmentDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetIssueEnrichmentDetailsItem(item *dnacentersdkgo.ResponseIssuesGetIssueEnrichmentDetailsIssueDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["issue"] = flattenIssuesGetIssueEnrichmentDetailsItemIssue(item.Issue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIssuesGetIssueEnrichmentDetailsItemIssue(items *[]dnacentersdkgo.ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["issue_source"] = item.IssueSource
		respItem["issue_category"] = item.IssueCategory
		respItem["issue_name"] = item.IssueName
		respItem["issue_description"] = item.IssueDescription
		respItem["issue_entity"] = item.IssueEntity
		respItem["issue_entity_value"] = item.IssueEntityValue
		respItem["issue_severity"] = item.IssueSeverity
		respItem["issue_priority"] = item.IssuePriority
		respItem["issue_summary"] = item.IssueSummary
		respItem["issue_timestamp"] = item.IssueTimestamp
		respItem["suggested_actions"] = flattenIssuesGetIssueEnrichmentDetailsItemIssueSuggestedActions(item.SuggestedActions)
		respItem["impacted_hosts"] = flattenIssuesGetIssueEnrichmentDetailsItemIssueImpactedHosts(item.ImpactedHosts)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetIssueEnrichmentDetailsItemIssueSuggestedActions(items *[]dnacentersdkgo.ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenIssuesGetIssueEnrichmentDetailsItemIssueSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetIssueEnrichmentDetailsItemIssueSuggestedActionsSteps(items *[]dnacentersdkgo.ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps) []interface{} {
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

func flattenIssuesGetIssueEnrichmentDetailsItemIssueImpactedHosts(items *[]dnacentersdkgo.ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueImpactedHosts) []interface{} {
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
