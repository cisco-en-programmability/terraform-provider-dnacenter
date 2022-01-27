package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSubscriptionDetailsEmail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of subscription details for specified connectorType
`,

		ReadContext: dataSourceEventSubscriptionDetailsEmailRead,
		Schema: map[string]*schema.Schema{
			"connector_type": &schema.Schema{
				Description: `connectorType query parameter. Connector Type [EMAIL]
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": &schema.Schema{
				Description: `instanceId query parameter. Instance Id of the specific configuration
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of the specific configuration
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

						"from_email_address": &schema.Schema{
							Description: `From Email Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"subject": &schema.Schema{
							Description: `Subject`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"to_email_addresses": &schema.Schema{
							Description: `To Email Addresses`,
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
	}
}

func dataSourceEventSubscriptionDetailsEmailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vConnectorType := d.Get("connector_type")
	vName, okName := d.GetOk("name")
	vInstanceID, okInstanceID := d.GetOk("instance_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEmailSubscriptionDetails")
		queryParams1 := dnacentersdkgo.GetEmailSubscriptionDetailsQueryParams{}

		queryParams1.ConnectorType = vConnectorType.(string)

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetEmailSubscriptionDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEmailSubscriptionDetails", err,
				"Failure at GetEmailSubscriptionDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetEmailSubscriptionDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEmailSubscriptionDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetEmailSubscriptionDetailsItems(items *dnacentersdkgo.ResponseEventManagementGetEmailSubscriptionDetails) []map[string]interface{} {
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
		respItem["from_email_address"] = item.FromEmailAddress
		respItem["to_email_addresses"] = item.ToEmailAddresses
		respItem["subject"] = item.Subject
		respItems = append(respItems, respItem)
	}
	return respItems
}
