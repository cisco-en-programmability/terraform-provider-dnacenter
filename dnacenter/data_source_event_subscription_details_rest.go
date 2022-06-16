package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSubscriptionDetailsRest() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of subscription details for specified connectorType
`,

		ReadContext: dataSourceEventSubscriptionDetailsRestRead,
		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Description: `instanceId query parameter. Instance Id of the specific configuration
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of Rest/Webhook Subscription detail's to limit in the resultset whose default value 10
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of the specific configuration
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The number of Rest/Webhook Subscription detail's to offset in the resultset whose default value 0
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
							Description: `Path Params`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"query_params": &schema.Schema{
							Description: `Query Params`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
	}
}

func dataSourceEventSubscriptionDetailsRestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vInstanceID, okInstanceID := d.GetOk("instance_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRestWebhookSubscriptionDetails")
		queryParams1 := dnacentersdkgo.GetRestWebhookSubscriptionDetailsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
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

		response1, restyResp1, err := client.EventManagement.GetRestWebhookSubscriptionDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRestWebhookSubscriptionDetails", err,
				"Failure at GetRestWebhookSubscriptionDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetRestWebhookSubscriptionDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRestWebhookSubscriptionDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetRestWebhookSubscriptionDetailsItems(items *dnacentersdkgo.ResponseEventManagementGetRestWebhookSubscriptionDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_id"] = item.InstanceID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["connector_type"] = item.ConnectorType
		respItem["url"] = item.URL
		respItem["method"] = item.Method
		respItem["trust_cert"] = item.TrustCert
		respItem["headers"] = flattenEventManagementGetRestWebhookSubscriptionDetailsItemsHeaders(item.Headers)
		respItem["query_params"] = item.QueryParams
		respItem["path_params"] = item.PathParams
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetRestWebhookSubscriptionDetailsItemsHeaders(items *[]dnacentersdkgo.ResponseItemEventManagementGetRestWebhookSubscriptionDetailsHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
