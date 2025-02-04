package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCiscoImcsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Cisco IMC.

- This data source retrieves the Cisco Integrated Management Controller (IMC) configuration for a Catalyst Center node,
identified by the specified ID.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By
providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health
status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in
the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-
management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where
Cisco IMC configuration is not supported by the deployment, these APIs will respond with a *404 Not Found* status code.
`,

		ReadContext: dataSourceCiscoImcsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The unique identifier for this Cisco IMC configuration
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceCiscoImcsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCiscoIMCConfigurationForACatalystCenterNode")
		vvID := vID.(string)

		response1, restyResp1, err := client.CiscoIMC.RetrievesTheCiscoIMCConfigurationForACatalystCenterNode(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCiscoIMCConfigurationForACatalystCenterNode", err,
				"Failure at RetrievesTheCiscoIMCConfigurationForACatalystCenterNode, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCiscoIMCConfigurationForACatalystCenterNode response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeItem(item *dnacentersdkgo.ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["node_id"] = item.NodeID
	respItem["ip_address"] = item.IPAddress
	respItem["username"] = item.Username
	return []map[string]interface{}{
		respItem,
	}
}
