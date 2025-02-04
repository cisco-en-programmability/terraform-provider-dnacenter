package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesTemplateIDNetworkProfilesForSitesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Retrieves the count of network profiles that a CLI template has been attached to by the template ID.
`,

		ReadContext: dataSourceTemplatesTemplateIDNetworkProfilesForSitesCountRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Description: `templateId path parameter. The *id* of the template, retrievable from *GET /intent/api/v1/templates*
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplatesTemplateIDNetworkProfilesForSitesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveCountOfNetworkProfilesAttachedToACLITemplate")
		vvTemplateID := vTemplateID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.RetrieveCountOfNetworkProfilesAttachedToACLITemplate(vvTemplateID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveCountOfNetworkProfilesAttachedToACLITemplate", err,
				"Failure at RetrieveCountOfNetworkProfilesAttachedToACLITemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplateItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveCountOfNetworkProfilesAttachedToACLITemplate response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplateItem(item *dnacentersdkgo.ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
