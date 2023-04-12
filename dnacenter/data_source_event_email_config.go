package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventEmailConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get Email Destination
`,

		ReadContext: dataSourceEventEmailConfigRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"email_config_id": &schema.Schema{
							Description: `UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"from_email": &schema.Schema{
							Description: `From Email`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"primary_smt_p_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"security": &schema.Schema{
										Description: `Security`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"user_name": &schema.Schema{
										Description: `User Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"secondary_smt_p_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"security": &schema.Schema{
										Description: `Security`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"user_name": &schema.Schema{
										Description: `User Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"subject": &schema.Schema{
							Description: `Subject`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"to_email": &schema.Schema{
							Description: `To Email`,
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

func dataSourceEventEmailConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEmailDestination")

		response1, restyResp1, err := client.EventManagement.GetEmailDestination()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEmailDestination", err,
				"Failure at GetEmailDestination, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetEmailDestinationItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEmailDestination response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetEmailDestinationItems(items *dnacentersdkgo.ResponseEventManagementGetEmailDestination) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["email_config_id"] = item.EmailConfigID
		respItem["primary_smt_p_config"] = flattenEventManagementGetEmailDestinationItemsPrimarySmtpConfig(item.PrimarySmtpConfig)
		respItem["secondary_smt_p_config"] = flattenEventManagementGetEmailDestinationItemsSecondarySmtpConfig(item.SecondarySmtpConfig)
		respItem["from_email"] = item.FromEmail
		respItem["to_email"] = item.ToEmail
		respItem["subject"] = item.Subject
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetEmailDestinationItemsPrimarySmtpConfig(item *dnacentersdkgo.ResponseItemEventManagementGetEmailDestinationPrimarySmtpConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["port"] = item.Port
	respItem["user_name"] = item.UserName
	respItem["password"] = item.Password
	respItem["security"] = item.Security

	return []map[string]interface{}{
		respItem,
	}

}

func flattenEventManagementGetEmailDestinationItemsSecondarySmtpConfig(item *dnacentersdkgo.ResponseItemEventManagementGetEmailDestinationSecondarySmtpConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_name"] = item.HostName
	respItem["port"] = item.Port
	respItem["user_name"] = item.UserName
	respItem["password"] = item.Password
	respItem["security"] = item.Security

	return []map[string]interface{}{
		respItem,
	}

}
