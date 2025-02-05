package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventWebhook() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get Webhook Destination
`,

		ReadContext: dataSourceEventWebhookRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of webhook configuration's to limit in the resultset whose default value 10
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The number of webhook configuration's to offset in the resultset whose default value 0
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
			"webhook_ids": &schema.Schema{
				Description: `webhookIds query parameter. List of webhook configurations
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"error_message": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"errors": &schema.Schema{
										Description: `Errors`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"status_message": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

												"default_value": &schema.Schema{
													Description: `Default Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"encrypt": &schema.Schema{
													Description: `Encrypt`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"is_proxy_route": &schema.Schema{
										Description: `Is Proxy Route`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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

									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"trust_cert": &schema.Schema{
										Description: `Trust Cert`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"url": &schema.Schema{
										Description: `Url`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"webhook_id": &schema.Schema{
										Description: `Webhook Id`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceEventWebhookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vWebhookIDs, okWebhookIDs := d.GetOk("webhook_ids")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWebhookDestination")
		queryParams1 := dnacentersdkgo.GetWebhookDestinationQueryParams{}

		if okWebhookIDs {
			queryParams1.WebhookIDs = vWebhookIDs.(string)
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

		response1, restyResp1, err := client.EventManagement.GetWebhookDestination(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWebhookDestination", err,
				"Failure at GetWebhookDestination, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEventManagementGetWebhookDestinationItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWebhookDestination response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetWebhookDestinationItem(item *dnacentersdkgo.ResponseEventManagementGetWebhookDestination) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_message"] = flattenEventManagementGetWebhookDestinationItemErrorMessage(item.ErrorMessage)
	respItem["api_status"] = item.APIStatus
	respItem["status_message"] = flattenEventManagementGetWebhookDestinationItemStatusMessage(item.StatusMessage)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEventManagementGetWebhookDestinationItemErrorMessage(item *dnacentersdkgo.ResponseEventManagementGetWebhookDestinationErrorMessage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["errors"] = item.Errors

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetWebhookDestinationItemStatusMessage(items *[]dnacentersdkgo.ResponseEventManagementGetWebhookDestinationStatusMessage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItem["webhook_id"] = item.WebhookID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["url"] = item.URL
		respItem["method"] = item.Method
		respItem["trust_cert"] = boolPtrToString(item.TrustCert)
		respItem["headers"] = flattenEventManagementGetWebhookDestinationItemStatusMessageHeaders(item.Headers)
		respItem["is_proxy_route"] = boolPtrToString(item.IsProxyRoute)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetWebhookDestinationItemStatusMessageHeaders(items *[]dnacentersdkgo.ResponseEventManagementGetWebhookDestinationStatusMessageHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItem["default_value"] = item.DefaultValue
		respItem["encrypt"] = boolPtrToString(item.Encrypt)
		respItems = append(respItems, respItem)
	}
	return respItems
}
