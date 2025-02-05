package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfilesIDPolicyTagsPolicyTagID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This endpoint retrieves the details of a specific *Policy Tag* associated with a given *Wireless Profile*. This data
source requires the *id* of the *Wireless Profile* and the *policyTagId* of the *Policy Tag*.
`,

		ReadContext: dataSourceWirelessProfilesIDPolicyTagsPolicyTagIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless Profile Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_tag_id": &schema.Schema{
				Description: `policyTagId path parameter. Policy Tag Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_zones": &schema.Schema{
							Description: `Ap Zones`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"policy_tag_id": &schema.Schema{
							Description: `Policy Tag Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"policy_tag_name": &schema.Schema{
							Description: `Policy Tag Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_ids": &schema.Schema{
							Description: `Site Ids`,
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

func dataSourceWirelessProfilesIDPolicyTagsPolicyTagIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vPolicyTagID := d.Get("policy_tag_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveASpecificPolicyTagForAWirelessProfile")
		vvID := vID.(string)
		vvPolicyTagID := vPolicyTagID.(string)

		response1, restyResp1, err := client.Wireless.RetrieveASpecificPolicyTagForAWirelessProfile(vvID, vvPolicyTagID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveASpecificPolicyTagForAWirelessProfile", err,
				"Failure at RetrieveASpecificPolicyTagForAWirelessProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveASpecificPolicyTagForAWirelessProfileItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveASpecificPolicyTagForAWirelessProfile response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveASpecificPolicyTagForAWirelessProfileItem(item *dnacentersdkgo.ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_ids"] = item.SiteIDs
	respItem["policy_tag_name"] = item.PolicyTagName
	respItem["ap_zones"] = item.ApZones
	respItem["policy_tag_id"] = item.PolicyTagID
	return []map[string]interface{}{
		respItem,
	}
}
