package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCiscoImcs() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Cisco IMC.

- This data source retrieves the configurations of the Cisco Integrated Management Controller (IMC) that have been added
to the Catalyst Center nodes.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By
providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health
status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in
the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-
management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where
Cisco IMC configuration is not supported by the deployment, these APIs will respond with a *404 Not Found* status code.
`,

		ReadContext: dataSourceCiscoImcsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The unique identifier for this Cisco IMC configuration
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP address of the Cisco IMC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"node_id": &schema.Schema{
							Description: `The UUID that represents the Catalyst Center node. Its value can be obtained from the *id* attribute of the response of the */dna/intent/api/v1/nodes-config* API.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Description: `Username of the Cisco IMC
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

func dataSourceCiscoImcsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesCiscoIMCConfigurationsForCatalystCenterNodes")

		response1, restyResp1, err := client.CiscoIMC.RetrievesCiscoIMCConfigurationsForCatalystCenterNodes()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesCiscoIMCConfigurationsForCatalystCenterNodes", err,
				"Failure at RetrievesCiscoIMCConfigurationsForCatalystCenterNodes, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesCiscoIMCConfigurationsForCatalystCenterNodes response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesItems(items *[]dnacentersdkgo.ResponseCiscoIMCRetrievesCiscoIMCConfigurationsForCatalystCenterNodesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["node_id"] = item.NodeID
		respItem["ip_address"] = item.IPAddress
		respItem["username"] = item.Username
		respItems = append(respItems, respItem)
	}
	return respItems
}
