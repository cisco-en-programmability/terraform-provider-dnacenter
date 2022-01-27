package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceGlobalCredentialDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Discovery.

- Deletes global credential for the given ID
`,

		ReadContext: dataSourceGlobalCredentialDeleteRead,
		Schema: map[string]*schema.Schema{
			"global_credential_id": &schema.Schema{
				Description: `globalCredentialId path parameter. ID of global-credential
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalCredentialDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vGlobalCredentialID := d.Get("global_credential_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeleteGlobalCredentialsByID")
		vvGlobalCredentialID := vGlobalCredentialID.(string)

		response1, restyResp1, err := client.Discovery.DeleteGlobalCredentialsByID(vvGlobalCredentialID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteGlobalCredentialsByID", err,
				"Failure at DeleteGlobalCredentialsByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryDeleteGlobalCredentialsByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeleteGlobalCredentialsByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryDeleteGlobalCredentialsByIDItem(item *dnacentersdkgo.ResponseDiscoveryDeleteGlobalCredentialsByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
