package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaExtranetPolicies() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of extranet policies that match the provided query parameters.
`,

		ReadContext: dataSourceSdaExtranetPoliciesRead,
		Schema: map[string]*schema.Schema{
			"extranet_policy_name": &schema.Schema{
				Description: `extranetPolicyName query parameter. Name of the extranet policy.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"extranet_policy_name": &schema.Schema{
							Description: `Name of the extranet policy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_ids": &schema.Schema{
							Description: `IDs of the fabric sites associated with this extranet policy.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `ID of the extranet policy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"provider_virtual_network_name": &schema.Schema{
							Description: `Name of the provider virtual network.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"subscriber_virtual_network_names": &schema.Schema{
							Description: `Name of the subscriber virtual network names.
`,
							Type:     schema.TypeList,
							Computed: true,
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

func dataSourceSdaExtranetPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vExtranetPolicyName, okExtranetPolicyName := d.GetOk("extranet_policy_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExtranetPolicies")
		queryParams1 := dnacentersdkgo.GetExtranetPoliciesQueryParams{}

		if okExtranetPolicyName {
			queryParams1.ExtranetPolicyName = vExtranetPolicyName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetExtranetPolicies(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetExtranetPolicies", err,
				"Failure at GetExtranetPolicies, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetExtranetPoliciesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExtranetPolicies response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetExtranetPoliciesItems(items *[]dnacentersdkgo.ResponseSdaGetExtranetPoliciesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["extranet_policy_name"] = item.ExtranetPolicyName
		respItem["fabric_ids"] = item.FabricIDs
		respItem["provider_virtual_network_name"] = item.ProviderVirtualNetworkName
		respItem["subscriber_virtual_network_names"] = item.SubscriberVirtualNetworkNames
		respItems = append(respItems, respItem)
	}
	return respItems
}
