package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSubscriptionRest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of Rest/Webhook Subscriptions's based on provided offset and limit
`,

		ReadContext: dataSourceEventSubscriptionRestRead,
		Schema: map[string]*schema.Schema{
			"event_ids": &schema.Schema{
				Description: `eventIds query parameter. List of subscriptions related to the respective eventIds (Comma separated event ids)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of Subscriptions's to limit in the resultset whose default value 10
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The number of Subscriptions's to offset in the resultset whose default value 0
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"categories": &schema.Schema{
										Description: `Categories`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"domains_subdomains": &schema.Schema{
										Description: `Domains Subdomains`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"event_ids": &schema.Schema{
										Description: `Event Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"others": &schema.Schema{
										Description: `Others`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"severities": &schema.Schema{
										Description: `Severities`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"sources": &schema.Schema{
										Description: `Sources`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"types": &schema.Schema{
										Description: `Types`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"is_private": &schema.Schema{
							Description: `Is Private`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connector_type": &schema.Schema{
										Description: `Connector Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"base_path": &schema.Schema{
													Description: `Base Path`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"body": &schema.Schema{
													Description: `Body`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"connect_timeout": &schema.Schema{
													Description: `Connect Timeout`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"connector_type": &schema.Schema{
													Description: `Connector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"headers": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"method": &schema.Schema{
													Description: `Method`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"path_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"query_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},

												"read_timeout": &schema.Schema{
													Description: `Read Timeout`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"resource": &schema.Schema{
													Description: `Resource`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"trust_cert": &schema.Schema{
													Description: `Trust Cert`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"subscription_id": &schema.Schema{
							Description: `Subscription Id`,
							Type:        schema.TypeString,
							Computed:    true,
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

func dataSourceEventSubscriptionRestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vEventIDs, okEventIDs := d.GetOk("event_ids")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRestWebhookEventSubscriptions")
		queryParams1 := dnacentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}

		if okEventIDs {
			queryParams1.EventIDs = vEventIDs.(string)
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

		response1, restyResp1, err := client.EventManagement.GetRestWebhookEventSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestWebhookEventSubscriptions", err,
				"Failure at GetRestWebhookEventSubscriptions, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetRestWebhookEventSubscriptionsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestWebhookEventSubscriptions response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItems(items *dnacentersdkgo.ResponseEventManagementGetRestWebhookEventSubscriptions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["version"] = item.Version
		respItem["subscription_id"] = item.SubscriptionID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["subscription_endpoints"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpoints(item.SubscriptionEndpoints)
		respItem["filter"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsFilter(item.Filter)
		respItem["is_private"] = boolPtrToString(item.IsPrivate)
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpoints(items *[]dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpoints) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_id"] = item.InstanceID
		respItem["subscription_details"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(item.SubscriptionDetails)
		respItem["connector_type"] = item.ConnectorType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetails(item *dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connector_type"] = item.ConnectorType
	respItem["instance_id"] = item.InstanceID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["url"] = item.URL
	respItem["base_path"] = item.BasePath
	respItem["resource"] = item.Resource
	respItem["method"] = item.Method
	respItem["trust_cert"] = item.TrustCert
	respItem["headers"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsHeaders(item.Headers)
	respItem["query_params"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsQueryParams(item.QueryParams)
	respItem["path_params"] = flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsPathParams(item.PathParams)
	respItem["body"] = item.Body
	respItem["connect_timeout"] = item.ConnectTimeout
	respItem["read_timeout"] = item.ReadTimeout

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsHeaders(items *[]dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["string"] = item.String
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsQueryParams(items *[]dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsQueryParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["string"] = item.String
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsSubscriptionEndpointsSubscriptionDetailsPathParams(items *[]dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsPathParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["string"] = item.String
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookEventSubscriptionsItemsFilter(item *dnacentersdkgo.ResponseItemEventManagementGetRestWebhookEventSubscriptionsFilter) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["event_ids"] = item.EventIDs
	respItem["others"] = item.Others
	respItem["domains_subdomains"] = item.DomainsSubdomains
	respItem["types"] = item.Types
	respItem["categories"] = item.Categories
	respItem["severities"] = item.Severities
	respItem["sources"] = item.Sources

	return []map[string]interface{}{
		respItem,
	}

}
