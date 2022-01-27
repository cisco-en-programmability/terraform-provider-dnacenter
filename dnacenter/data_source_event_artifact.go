package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventArtifact() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of artifacts based on provided offset and limit
`,

		ReadContext: dataSourceEventArtifactRead,
		Schema: map[string]*schema.Schema{
			"event_ids": &schema.Schema{
				Description: `eventIds query parameter. List of eventIds
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. # of records to return in result set
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Record start offset
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. sorting order (asc/desc)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"search": &schema.Schema{
				Description: `search query parameter. findd matches in name, description, eventId, type, category
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort by field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Description: `tags query parameter. Tags defined
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"artifact_id": &schema.Schema{
							Description: `Artifact Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"cisco_dna_event_link": &schema.Schema{
							Description: `Cisco D N A Event Link`,
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

						"event_payload": &schema.Schema{
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

									"details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_ip": &schema.Schema{
													Description: `Device Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"message": &schema.Schema{
													Description: `Message`,
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

						"event_templates": &schema.Schema{
							Description: `Event Templates`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"is_private": &schema.Schema{
							Description: `Is Private`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_template_enabled": &schema.Schema{
							Description: `Is Template Enabled`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_tenant_aware": &schema.Schema{
							Description: `Is Tenant Aware`,
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

						"note": &schema.Schema{
							Description: `Note`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sub_domain": &schema.Schema{
							Description: `Sub Domain`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"supported_connector_types": &schema.Schema{
							Description: `Supported Connector Types`,
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

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
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

func dataSourceEventArtifactRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEventIDs, okEventIDs := d.GetOk("event_ids")
	vTags, okTags := d.GetOk("tags")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSearch, okSearch := d.GetOk("search")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEventArtifacts")
		queryParams1 := dnacentersdkgo.GetEventArtifactsQueryParams{}

		if okEventIDs {
			queryParams1.EventIDs = vEventIDs.(string)
		}
		if okTags {
			queryParams1.Tags = vTags.(string)
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
		if okSearch {
			queryParams1.Search = vSearch.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetEventArtifacts(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEventArtifacts", err,
				"Failure at GetEventArtifacts, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetEventArtifactsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEventArtifacts response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetEventArtifactsItems(items *dnacentersdkgo.ResponseEventManagementGetEventArtifacts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["version"] = item.Version
		respItem["artifact_id"] = item.ArtifactID
		respItem["namespace"] = item.Namespace
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["domain"] = item.Domain
		respItem["sub_domain"] = item.SubDomain
		respItem["tags"] = item.Tags
		respItem["is_template_enabled"] = boolPtrToString(item.IsTemplateEnabled)
		respItem["cisco_dna_event_link"] = item.CiscoDnaEventLink
		respItem["note"] = item.Note
		respItem["is_private"] = boolPtrToString(item.IsPrivate)
		respItem["event_payload"] = flattenEventManagementGetEventArtifactsItemsEventPayload(item.EventPayload)
		respItem["event_templates"] = flattenEventManagementGetEventArtifactsItemsEventTemplates(item.EventTemplates)
		respItem["is_tenant_aware"] = boolPtrToString(item.IsTenantAware)
		respItem["supported_connector_types"] = item.SupportedConnectorTypes
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetEventArtifactsItemsEventPayload(item *dnacentersdkgo.ResponseItemEventManagementGetEventArtifactsEventPayload) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["event_id"] = item.EventID
	respItem["version"] = item.Version
	respItem["category"] = item.Category
	respItem["type"] = item.Type
	respItem["source"] = item.Source
	respItem["severity"] = item.Severity
	respItem["details"] = flattenEventManagementGetEventArtifactsItemsEventPayloadDetails(item.Details)
	respItem["additional_details"] = flattenEventManagementGetEventArtifactsItemsEventPayloadAdditionalDetails(item.AdditionalDetails)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetEventArtifactsItemsEventPayloadDetails(item *dnacentersdkgo.ResponseItemEventManagementGetEventArtifactsEventPayloadDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_ip"] = item.DeviceIP
	respItem["message"] = item.Message

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetEventArtifactsItemsEventPayloadAdditionalDetails(item *dnacentersdkgo.ResponseItemEventManagementGetEventArtifactsEventPayloadAdditionalDetails) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenEventManagementGetEventArtifactsItemsEventTemplates(items *[]dnacentersdkgo.ResponseItemEventManagementGetEventArtifactsEventTemplates) []interface{} {
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
