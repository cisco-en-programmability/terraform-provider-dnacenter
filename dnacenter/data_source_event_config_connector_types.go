package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventConfigConnectorTypes() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get the list of connector types
`,

		ReadContext: dataSourceEventConfigConnectorTypesRead,
		Schema: map[string]*schema.Schema{

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

						"display_name": &schema.Schema{
							Description: `Display Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_custom_connector": &schema.Schema{
							Description: `Is Custom Connector`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_default_supported": &schema.Schema{
							Description: `Is Default Supported`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventConfigConnectorTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConnectorTypes")

		response1, restyResp1, err := client.EventManagement.GetConnectorTypes()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConnectorTypes", err,
				"Failure at GetConnectorTypes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEventManagementGetConnectorTypesItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConnectorTypes response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementGetConnectorTypesItems(items *dnacentersdkgo.ResponseEventManagementGetConnectorTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["connector_type"] = item.ConnectorType
		respItem["display_name"] = item.DisplayName
		respItem["is_default_supported"] = boolPtrToString(item.IsDefaultSupported)
		respItem["is_custom_connector"] = boolPtrToString(item.IsCustomConnector)
		respItems = append(respItems, respItem)
	}
	return respItems
}
