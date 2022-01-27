package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventSubscriptionDetailsSyslog() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Gets the list of subscription details for specified connectorType
`,

		ReadContext: dataSourceEventSubscriptionDetailsSyslogRead,
		Schema: map[string]*schema.Schema{
			"connector_type": &schema.Schema{
				Description: `connectorType query parameter. Connector Type [SYSLOG]
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

						"syslog_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"host": &schema.Schema{
										Description: `Host`,
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
										Type:        schema.TypeString,
										Computed:    true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
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

func dataSourceEventSubscriptionDetailsSyslogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vConnectorType := d.Get("connector_type")
	vName, okName := d.GetOk("name")
	vInstanceID, okInstanceID := d.GetOk("instance_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSyslogSubscriptionDetails")
		queryParams1 := dnacentersdkgo.GetSyslogSubscriptionDetailsQueryParams{}

		queryParams1.ConnectorType = vConnectorType.(string)

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okInstanceID {
			queryParams1.InstanceID = vInstanceID.(string)
		}

		response1, restyResp1, err := client.EventManagement.GetSyslogSubscriptionDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSyslogSubscriptionDetails", err,
				"Failure at GetSyslogSubscriptionDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetSyslogSubscriptionDetailsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSyslogSubscriptionDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetSyslogSubscriptionDetailsItems(items *dnacentersdkgo.ResponseEventManagementGetSyslogSubscriptionDetails) []map[string]interface{} {
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
		respItem["syslog_config"] = flattenEventManagementGetSyslogSubscriptionDetailsItemsSyslogConfig(item.SyslogConfig)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEventManagementGetSyslogSubscriptionDetailsItemsSyslogConfig(item *dnacentersdkgo.ResponseItemEventManagementGetSyslogSubscriptionDetailsSyslogConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["host"] = item.Host
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol

	return []map[string]interface{}{
		respItem,
	}

}
