package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSNMPConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get SNMP Destination
`,

		ReadContext: dataSourceEventSNMPConfigRead,
		Schema: map[string]*schema.Schema{
			"config_id": &schema.Schema{
				Description: `configId query parameter. List of SNMP configurations
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of SNMP configuration's to limit in the resultset whose default value 10
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The number of SNMP configuration's to offset in the resultset whose default value 0
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

									"auth_password": &schema.Schema{
										Description: `Auth Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"community": &schema.Schema{
										Description: `Community`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"config_id": &schema.Schema{
										Description: `Config Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ip_address": &schema.Schema{
										Description: `Ip Address`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"privacy_password": &schema.Schema{
										Description: `Privacy Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snmp_auth_type": &schema.Schema{
										Description: `Snmp Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snmp_mode": &schema.Schema{
										Description: `Snmp Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snmp_privacy_type": &schema.Schema{
										Description: `Snmp Privacy Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"snmp_version": &schema.Schema{
										Description: `Snmp Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"user_name": &schema.Schema{
										Description: `User Name`,
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
				},
			},
		},
	}
}

func dataSourceEventSNMPConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vConfigID, okConfigID := d.GetOk("config_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSNMPDestination")
		queryParams1 := dnacentersdkgo.GetSNMPDestinationQueryParams{}

		if okConfigID {
			queryParams1.ConfigID = vConfigID.(string)
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

		response1, restyResp1, err := client.EventManagement.GetSNMPDestination(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSNMPDestination", err,
				"Failure at GetSNMPDestination, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEventManagementGetSNMPDestinationItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSNMPDestination response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetSNMPDestinationItem(item *dnacentersdkgo.ResponseEventManagementGetSNMPDestination) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_message"] = flattenEventManagementGetSNMPDestinationItemErrorMessage(item.ErrorMessage)
	respItem["api_status"] = item.APIStatus
	respItem["status_message"] = flattenEventManagementGetSNMPDestinationItemStatusMessage(item.StatusMessage)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEventManagementGetSNMPDestinationItemErrorMessage(item *dnacentersdkgo.ResponseEventManagementGetSNMPDestinationErrorMessage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["errors"] = item.Errors

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetSNMPDestinationItemStatusMessage(items *[]dnacentersdkgo.ResponseEventManagementGetSNMPDestinationStatusMessage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItem["config_id"] = item.ConfigID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["ip_address"] = item.IPAddress
		respItem["port"] = item.Port
		respItem["snmp_version"] = item.SNMPVersion
		respItem["community"] = item.Community
		respItem["user_name"] = item.UserName
		respItem["snmp_mode"] = item.SNMPMode
		respItem["snmp_auth_type"] = item.SNMPAuthType
		respItem["auth_password"] = item.AuthPassword
		respItem["snmp_privacy_type"] = item.SNMPPrivacyType
		respItem["privacy_password"] = item.PrivacyPassword
		respItems = append(respItems, respItem)
	}
	return respItems
}
