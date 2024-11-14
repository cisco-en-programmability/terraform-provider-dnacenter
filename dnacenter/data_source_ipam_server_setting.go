package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamServerSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on System Settings.

- Retrieves configuration details of the external IPAM server.  If an external IPAM server has not been created, this
resource will return a '404' response.
`,

		ReadContext: dataSourceIPamServerSettingRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"provider": &schema.Schema{
							Description: `Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"server_name": &schema.Schema{
							Description: `A descriptive name of this external server, used for identification purposes`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"server_url": &schema.Schema{
							Description: `The URL of this external server`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"state": &schema.Schema{
							Description: `State of the the external IPAM.* OK indicates success of most recent periodic communication check with external IPAM.* CRITICAL indicates failure of most recent attempt to communicate with the external IPAM.* SYNCHRONIZING indicates that the process of synchronizing the external IPAM database with the local IPAM database is running and all other IPAM processes will be blocked until the completes.* DISCONNECTED indicates the external IPAM is no longer being used.`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"user_name": &schema.Schema{
							Description: `The external IPAM server login username`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"view": &schema.Schema{
							Description: `The view under which pools are created in the external IPAM server.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIPamServerSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesConfigurationDetailsOfTheExternalIPAMServer")

		response1, restyResp1, err := client.SystemSettings.RetrievesConfigurationDetailsOfTheExternalIPAMServer()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesConfigurationDetailsOfTheExternalIPAMServer", err,
				"Failure at RetrievesConfigurationDetailsOfTheExternalIPAMServer, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesConfigurationDetailsOfTheExternalIPAMServer response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerItem(item *dnacentersdkgo.ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["provider"] = item.Provider
	respItem["server_name"] = item.ServerName
	respItem["server_url"] = item.ServerURL
	respItem["state"] = item.State
	respItem["user_name"] = item.UserName
	respItem["view"] = item.View
	return []map[string]interface{}{
		respItem,
	}
}
