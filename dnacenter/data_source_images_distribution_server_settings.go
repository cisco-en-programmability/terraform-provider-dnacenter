package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImagesDistributionServerSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- Retrieve the list of remote image distribution servers. There can be up to two remote servers.Product always acts as
local distribution server, and it is not part of this API response.
`,

		ReadContext: dataSourceImagesDistributionServerSettingsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Unique identifier for the server
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"port_number": &schema.Schema{
							Description: `Port number
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"root_location": &schema.Schema{
							Description: `Server root location
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"server_address": &schema.Schema{
							Description: `FQDN or IP address of the server
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Description: `Server username
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceImagesDistributionServerSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveImageDistributionServers")

		response1, restyResp1, err := client.SoftwareImageManagementSwim.RetrieveImageDistributionServers()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveImageDistributionServers", err,
				"Failure at RetrieveImageDistributionServers, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimRetrieveImageDistributionServersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveImageDistributionServers response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimRetrieveImageDistributionServersItems(items *[]dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["username"] = item.Username
		respItem["server_address"] = item.ServerAddress
		respItem["port_number"] = item.PortNumber
		respItem["root_location"] = item.RootLocation
		respItems = append(respItems, respItem)
	}
	return respItems
}
